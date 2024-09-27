# Debug Service from IntelliJ

Based on [Remote debug a Springboot Application on Kuberenetes](https://urosht.dev/remote-debug-a-spring-boot-application-on-kubernetes/)

## On Remote Server

1. Add the following to service configuration:

  ```bash
  # /app/configuration/service/management

  JavaOptions.value: -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005
  ```

1. Restart service:

  ```bash
  kubectl delete pod -l app=management
  ```

1. Expose management:

  ```bash
  kubectl port-forward management-94bfb9df-dsd7k 5005:5005 --address=0.0.0.0
  ```

## In IntelliJ

1. Create debug configuration

  ```text
  Debugger mode: Attach to remote JVM
  Host: IP(!) of remote server
  Port: port debugger wil listen on (default is 5005)
  ```

  In XML:

  ```markup
    <component name="ProjectRunConfigurationManager">
      <configuration default="false" name="k8s_debug" type="Remote">
        <module name="management" />
        <option name="USE_SOCKET_TRANSPORT" value="true" />
        <option name="SERVER_MODE" value="false" />
        <option name="SHMEM_ADDRESS" />
        <option name="HOST" value="10.50.97.236" />
        <option name="PORT" value="5005" />
        <option name="AUTO_RESTART" value="false" />
        <RunnerSettings RunnerId="Debug">
          <option name="DEBUG_PORT" value="5005" />
          <option name="LOCAL" value="false" />
        </RunnerSettings>
        <method v="2" />
      </configuration>
    </component>

  ```

1. Run debug configuration
1. Set breakpoint and debug
