---
slug: gulpjs-task-using-powershell
title: How to Run a Sequential gulp.js Task using PowerShell
description: Fill me up!
authors: [kbbgl]
tags: [gulp,gulp.js,gulpjs,script,scripting,task_scheduler,windows]
---

## Introduction

One of our largest and strategic customer recently requested to be able to execute an automated ETL pipeline from one of their Windows multinode cluster in a sequential manner. The ETL pipeline has basically 3-stages:

1. Initialize an import from their many data sources into a consolidated database on master node.
2. After the import step finishes, copy the flat files representing the consolidated database to the slave nodes.
3. After the copying step is completed, delete the old instance of the database from the slave nodes and launch the newly-copied database.

The customer’s request rose out of a necessity because there was no way to perform this process in a sequential manner out-of-the-box as the product only by offered using scheduled manner. Since the Windows version of this feature was deprecated and no longer maintained, it was left up to Support to provide a crafty solution to this problem as R&D and Product were not willing to put in hours to add this feature.

## Starting Point

Although the product did not supply a way to perform the sequential task out-of-the-box, I did have a pretty solid ground to begin with.
On the master node, where the ETL process needs to initilize from, there were a few Windows Services running which allowed the non-sequential, scheduled build and copy operation to be done.
The first service was a management service written in .NET/C# that managed the consolidated database instances and exposed a CLI to be able to interact with them. Using this CLI, we’re able to start initialize the ETL process and refresh the remote databases after copying the flat files to the slave nodes.
The second service was a NodeJS application which wrapped a [`gulp.js`](https://gulpjs.com/) task initializer. The tasks were responsible for interacting with the first service using the CLI to perform the following operations:

1. Initialize the ETL process.
2. After the ETL process finishes on the master node, synchronously copy the database flat files from the master node to the slave nodes.
3. Send a command to the slave nodes to restart the database instance with the copied flat files.

The way product worked out-of-the-box is to read a configuration file which includes:

a. The cron pattern to initialize the `gulp` task represented by scheduler.schedule, `15 20 1 * *`.
b. The slave nodes’ remote copy location. This is where the flat files copied from the master node will be copied into the slave nodes. Represented by `dbN.slave_path`.
c. The name of the database for which to initialize the ETL process on the master node. Represented by `tasks.some_db_task.build.db`.
The configuration file (very redacted) looks something like this:

```json
{
  "db1": {
      "db_name": "SomeDB",
      "slave_path": "\\\\slave_node_1_hostname\\C:\\path\\to\\SomeDB\\FlatFiles"
    },
    "db1": {
      "db_name": "SomeDB",
      "slave_path": "\\\\slave_node_2_hostname\\C:\\path\\to\\SomeDB\\FlatFiles"
    },
    "scheduler":[  
      {  
         "task":"some_db_task",
         "schedule":"15 20 1 * *",
         "enabled": true
      }
   ],
   "tasks":{  
      "some_db_task":[  
         {  
            "build":{  
               "db":[  
                  "SomeDB"
                  ]
            }
         },
         {  
            "distribute":[  
               "db1.slave_path",
               "db2.slave_path"
               ]
         }
      ]
   }
}
```

I figured, if I could find a way to launch the `gulp` task independently, in the example above `some_db_task`, and develop some logic to initialize it sequentially, I would be able to supply a solution to the customer.
The first step was to understand how `gulp` tasks worked.

## `gulp` Research

After reviewing the source code of the NodeJS service which initializes the `gulp` tasks, I found that there are a few ways to run them:

- We can run the task using

    ```bash
    node.exe /path/to/gulp.js
    ```

- We can specify the task name according to what’s configured:

    ```bash
    node.exe /path/to/gulp.js some_db_task
    ```

- Run it using `gulp`:

    ```bash
    gulp some_db_task
    ```

Unfortunately, option 3 was too cumbersome as it required installing NPM/NodeJS on the master node since the product does not install it by default (it just uses the `node.exe` binary). And since the customer requested a specific DB to run this pipeline on, I decided to discard option 1 and develop a script to run option 2.

As I began testing the procedure, I noticed that the gulp tasks created a forked process by which they would interact with the .NET management service. The problem was, once the command to initialize the ETL process finished executing, the process would hang and would not complete the remote copying operations. I needed to figure out what was causing this hang so I opened my IDE to review the source code behind gulp.
What I found was that after the gulp task stopped its execution, it did not gracefully terminate the process. It can be seen here:

```javascript
gulpInst.on('task_stop', function(e) {
    var time = prettyTime(e.hrDuration);
    gutil.log(
        'Finished', '\'' + chalk.cyan(e.task) + '\'',
        'after', chalk.magenta(time)
    );
});
```

I added the following line to ensure that the process exited when the task stopped:

```javascript
gulpInst.on('task_stop', function(e) {
    var time = prettyTime(e.hrDuration);
    gutil.log(
        'Finished', '\'' + chalk.cyan(e.task) + '\'',
        'after', chalk.magenta(time)
    );
    process.exit(0); // <-------- SHOULD BE HERE!
});
```

After updating the source code, the hanging processes issue was resolved!
Next step was to write a script to run the sequential build.

## `gulp` Powered by PowerShell

My scripting language of choice is bash but since I was working on a Windows environment, without an ability to install 3rd-party tools such as `cygwin` and no access to Python either, I defaulted to PowerShell. The script name is `gulpTaskSequentialWithWait.ps1`:

```powershell
#  -------------------
#  -------------------
#  User configuration 
 
# Configurate number of minutes to wait before each ETL
$waitTimeBetweenBuildsInMinutes = 1
 
# The gulp task name
$taskName = "some_db_task"
 
# log location -default writes to directory where ps script was launched from
$currentLocation = $PSScriptRoot
$logLocation = Join-Path -Path $currentLocation -ChildPath "gulpTaskSequential.log"
Write-Host "Writing logs to '$logLocation'"
 
#  -------------------
#  -------------------
 
 
$workingDir = "C:\path\to\dir"
$env:NODE_ENV = 'production'
$gulpFile = "./node_modules/gulp/bin/gulp.js"
$nodeJsExecutableFilePath = "node.exe"
 
 
# Start gulp task infintely
do {
 
     
    # Create a new process
    $pinfo = New-Object System.Diagnostics.ProcessStartInfo
 
    # Set working dir to Orchestrator folder
    $pinfo.WorkingDirectory = $workingDir
 
    Write-Host "Changed working dir:" $pinfo.WorkingDirectory
 
    # Initialize process arguments and output redirection
    $pinfo.FileName = Join-Path -Path $workingDir -ChildPath $nodeJsExecutableFilePath
    $pinfo.RedirectStandardError = $true
    $pinfo.RedirectStandardOutput = $true
    $pinfo.UseShellExecute = $false
    $pinfo.CreateNoWindow = $true
    $pinfo.Arguments = @($gulpFile, $taskName)
     
    $p = New-Object System.Diagnostics.Process
    $p.StartInfo = $pinfo
    $p.Start() | Out-Null
     
    Write-Host (Get-Date -Format U) "Starting task '$nodeJsExecutableFilePath $gulpFile $taskName'..."
    (Get-Date -Format U) + "Starting task '$nodeJsExecutableFilePath $gulpFile $taskName'..." | Out-File -FilePath $logLocation -Append
 
    $stdout = $p.StandardOutput.ReadToEnd()
    $stderr = $p.StandardError.ReadToEnd()
    $p.WaitForExit()
 
    # Terminate process if non-zero exit code returned from command
    if ($p.ExitCode -ne 0){
 
        (Get-Date -Format U) +  " exit code: " + $p.ExitCode | Out-File -FilePath $logLocation -Append
        Write-Host (Get-Date -Format U) "exit code: " $p.ExitCode
 
        (Get-Date -Format U) +  " stderr: $stderr" | Out-File -FilePath $logLocation -Append
        Write-Host (Get-Date -Format U) "stderr: $stderr"
 
        (Get-Date -Format U) + "Exiting.." | Out-File -FilePath $logLocation -Append
        Write-Host (Get-Date -Format U) "Exiting.."
        exit
    } else {
 
        (Get-Date -Format U) + " stdout: $stdout" | Out-File -FilePath $logLocation -Append
        Write-Host (Get-Date -Format U) "stdout: $stdout"
 
        $sleepTimeInSeconds = $waitTimeBetweenBuildsInMinutes * 60
 
        Write-Host (Get-Date -Format U) " Starting sleep timer: $sleepTimeInSeconds seconds..."
        (Get-Date -Format U) + " Starting sleep timer: $sleepTimeInSeconds seconds..." | Out-File -FilePath $logLocation -Append
 
 
        Start-Sleep -Seconds $sleepTimeInSeconds
 
 
        Write-Host (Get-Date -Format U) "Sleep ended"
        (Get-Date -Format U) + " Sleep ended" | Out-File -FilePath $logLocation -Append
 
    }
} while ($True) 
```

The script basically runs an infinite loop with a configurable timeout (line 6, variable `waitTimeBetweenBuildsInMinutes`) executing the `gulp` task. We can also configure the name of the task in line 9, variable `taskName`.
I needed to ensure that the PowerShell script would be initialized after a reboot so I wrote a Task Scheduler XML to automatically start the script upon server boot:

```markup
<?xml version="1.0" encoding="UTF-16"?>
<Task version="1.2" xmlns="http://schemas.microsoft.com/windows/2004/02/mit/task">
  <RegistrationInfo>
    <Date>2020-01-01T00:00:00.0000000</Date>
    <Author>SOME_DOMAIN\YOUR_USER</Author>                    <!--Change author-->
  </RegistrationInfo>
  <Triggers>
    <BootTrigger>
      <ExecutionTimeLimit>PT12H</ExecutionTimeLimit>
      <Enabled>true</Enabled>
    </BootTrigger>
  </Triggers>
  <Principals>
    <Principal id="Author">
      <UserId>SOME_DOMAIN\YOUR_USER</UserId>                  <!--Change author-->
      <LogonType>Password</LogonType>
      <RunLevel>HighestAvailable</RunLevel>
    </Principal>
  </Principals>
  <Settings>
    <MultipleInstancesPolicy>IgnoreNew</MultipleInstancesPolicy>
    <DisallowStartIfOnBatteries>true</DisallowStartIfOnBatteries>
    <StopIfGoingOnBatteries>true</StopIfGoingOnBatteries>
    <AllowHardTerminate>true</AllowHardTerminate>
    <StartWhenAvailable>false</StartWhenAvailable>
    <RunOnlyIfNetworkAvailable>false</RunOnlyIfNetworkAvailable>
    <IdleSettings>
      <StopOnIdleEnd>true</StopOnIdleEnd>
      <RestartOnIdle>false</RestartOnIdle>
    </IdleSettings>
    <AllowStartOnDemand>true</AllowStartOnDemand>
    <Enabled>true</Enabled>
    <Hidden>false</Hidden>
    <RunOnlyIfIdle>false</RunOnlyIfIdle>
    <WakeToRun>false</WakeToRun>
    <ExecutionTimeLimit>P1D</ExecutionTimeLimit>
    <Priority>7</Priority>
  </Settings>
  <Actions Context="Author">
    <Exec>
      <Command>powershell.exe</Command>
      <Arguments>-ExecutionPolicy Bypass "C:\path\to\gulpTaskSequentialWithWait.ps1"</Arguments> <!--Change script location-->
    </Exec>
  </Actions>
</Task>
```

The Task runs the following commands:

```powershell
powershell.exe -ExecutionPolicy Bypass gulpTaskSequentialWithWait.ps1
```

And ensures that the Task is run on boot:

```markup
<Triggers>
    <BootTrigger>
      <ExecutionTimeLimit>PT12H</ExecutionTimeLimit>
      <Enabled>true</Enabled>
    </BootTrigger>
  </Triggers>
```

Another happy customer!
