---
slug: os-linux-booting-create-service
title: "How to Create a Service"
authors: [kbbgl]
tags: [os, linux, booting, create_service]
---

# How to Create a Service

1. Create a service file and put in `/etc/init.d/fake.service`:

    ```bash
    #!/bin/bash
    # fake_service
    # Starts up, writes to a dummy file, and exits
    #
    # chkconfig: 35 69 31
    # description: This service doesn't do anything.
    # Source function library

    . /etc/sysconfig/fake_service

    case"$1" in start) 
        
        echo "Running fake_service in start mode..."
        touch /var/lock/subsys/fake_service
        echo "$0start at $(date)" >> /var/log/fake_service.log
        
        if[${VAR1} ="true" ]
        then
            echo"VAR1 set to true" >> /var/log/fake_service.log
        fi
        echo
        ;;
        stop)
        echo "Running the fake_service script in stop mode..."
        echo "$0stop at$(date)" >> /var/log/fake_service.log
        if[${VAR2} ="true" ]
        then
            echo"VAR2 = true" >> /var/log/fake_service.logfirm -f /var/lock/subsys/fake_service
        echo
        ;;
        *)
        echo"Usage: fake_service {start | stop}"
        exit 1
    esac
    exit 0
    ```

1. Create a config file in `/etc/sysconfig/fake.service` (RHEL) or `/etc/default/fake.service` (Debian):

    ```bash
    VAR1="true"
    VAR2="true"
    ```

1. Give permission:

    ```bash
    sudo chmod 755 /etc/init.d/fake.service
    ```

1. Test script

    ```bash
    sudo service fake_service
    sudo service fake_service start
    sudo service fake_service stop
    ```

1. Start servce whenever system starts:

    ```bash
    # Will generate error, next command will add it
    sudo chkconfig --list fake_service
    sudo chkconfig --add fake_service
    sudo chkconfig fake_service on
    sudo chkconfig fake_service off
    ```
