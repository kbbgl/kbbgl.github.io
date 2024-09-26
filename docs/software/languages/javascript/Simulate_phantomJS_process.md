# Simulate PhantomJS Process

```javascript
v
var path = require('path')
const veryOriginalUrl = "http://127.0.0.1:14991/app/reporting?lang=fr-FR#/dashboards/preview/5cc17e8effb09a093ca6e1d4?paperFormat=A3&paperOrientation=landscape&elasticubeBuilt=2020-05-10T19:04:03.493Z&showNarration=undefined";
var phantom = path.join("C:\\Program Files\\app\\app\\galaxy-service\\src\\features\\dashboards\\export\\bin\\", 'phantomjs.exe')
var child_process = require('child_process');


try {
    var
        err,
        warn,
        messageStr = '',
        options = [
            '--web-security=no',
            '--ssl-protocol=any',
            '--ignore-ssl-errors=yes',
            "C:\\Program Files\\app\\app\\galaxy-service\\src\\features\\dashboards\\export\\rasterizeDebug.js",
            veryOriginalUrl,
            path.join("./exports/55a695b0e0076954d6000070.pdf"),
            "A3",
            "landscape",
            "95",
            "asis",
            "undefined",
            "undefined",
            300,
            "Token",
            null,
            undefined,
            undefined,
            "--debug=true"
        ];

    console.info(`Exporting URL "${veryOriginalUrl}"`);
    var phantomProcess = child_process.spawn(phantom, options);
    console.info(`Spawning phantom process with PID: ${phantomProcess.pid}`)

    phantomProcess.stdout.on('data', function(data) {
        console.log(`stdout: ${data}` )
    });

    phantomProcess.stderr.on('data', function(buffer) {
        console.log(`stderr: ${buffer.toString()}` )
        messageStr += buffer.toString();
    });
    phantomProcess.stderr.on('end', function(buffer) {
        if (buffer) {
            messageStr += buffer.toString();
        }

            console.log(`error stream ended: ${messageStr}`)
    });

    phantomProcess.on('error', function(error) {
        if (!err) {
            console.error(error);
        }
        console.error(`Exporting URL "${veryOriginalUrl}" phantomjs child process error: ${error}`);
    });

    phantomProcess.on('exit', function(code) {
        if (code == 0 && !err && !warn) {
            console.log(`program exited with exit code 0`)
        }else{
            console.log(`Exit signal received: ${err || warn}`);
        }
    });

}
catch(e) {
    console.error(`Exporting URL "${veryOriginalUrl}" error: ${e}`);
}
```
