# Displaying Animations to Indicate Long-Running Tasks

```bash
#!/bin/bash
while true;
do
    # Frame #1
    printf "\r< Loading..." 
    sleep 0.5
    # Frame #2 
    printf "\r> Loading..." 
    sleep 0.5 
done
```


We can add more frames to the animation and display it until a specific time-consuming task completes with the following Bash script.

```bash
#!/bin/bash
sleep 5 &
pid=$!
frames="/ | \\ -"
while kill -0 $pid 2&>1 > /dev/null;
do
    for frame in $frames;
    do
        printf "\r$frame Loading..." 
        sleep 0.5
    done
done
printf "\n"
```