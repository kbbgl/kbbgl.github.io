# Process Management

## Finding processes

```bash
ps aux | grep someprocess
```

```bash
top
```

Press `?` or `H` while running `top` to get a list of options

## Changing process priority

The `nice` command is used to the prioritize the process to the kernel.
The values range from `-19 to +19` where a high value is being very _nice_ => lower priority so:

```plaintext
-19 == very nice == lowest priority
+19 == not nice  == highest priority
```

A child process inherits the priority from the parent.

To change the priority after it was initially modified, we would use the `renice` command:

```bash
# Increase `nice` value by 10
nice -n -10 /bin/slowprocess

# Decrease value by 10
nice -n 10 /bin/slowprocess

# Change the priority of process IDs 987 and 32, 
# and all processes owned by users daemon and root, 
# to be one greater (+1, one increment "nicer") than its current value.
renice +1 987 -u daemon root -p 32
```

Useful settings for priority are:

* 20: the affected processes will run only when nothing else in the system needs the resources.
* 0: the default.
* any negative value: will make things go very fast, at the expense of other processes.

## Killing Processes

| Signal Name | Number | Description |
|-|-|-|
| `SIGHUP` | 1 | This is known as the Hangup (HUP) signal. It stops the designated process and restarts it with the same PID. |
| `SIGINT` | 2 | This is the Interrupt (INT) signal. It is a weak kill signal that isn’t guaranteed to work, but it works in most cases. |
| `SIGQUIT` | 3 | This is known as the core dump . It terminates the process and saves the process information in memory, and then it saves this information in the current working directory to a file named core . |
| `SIGTERM` | 15 | This is the Termination (TERM) signal. It is the kill command’s default kill signal. |
| `SIGKILL` | 9 | This is the absolute kill signal. It forces the process to stop by sending the process’s resources to a special device, `/dev/null` . |

To kill a process:

```bash
# Kill specific process ID
kill -9 6996

# Kill process by name
killall -9 zombieprocess
```

## Run process in the background

Using `&` at the end of the command

```bash
aprocess 1234 &
```

## Moving a process to the foreground

```bash
fg 1234
```

## Scheduling processes

We can use two tools:

* `at` - a daemon process useful for one-time job to be run in the future.

```bash
at 7:20am 
at> /root/myscanningscript

at now + 20 minutes
at> /root/myscanningscript

at 7:20pm 06/25/2019
```

* `crond`- more suited for every day/week type of jobs.
