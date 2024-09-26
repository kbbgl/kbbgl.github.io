# Run Profiler

## Start app from cmd

```bash
java -XX:+UnlockCommercialFeatures -XX:+FlightRecorder -XX:StartFlightRecording=duration=60s,filename=myrecording.jfr MyApp
```

## Debug running process

```bash
#a 60-second recording on the running Java process with the identifier 5368 and save it to myrecording.jfr in the current directory, use the following:

jcmd 5368 JFR.start duration=60s filename=myrecording.jfr
```

## list all running processes

```bash
jcmd
```
