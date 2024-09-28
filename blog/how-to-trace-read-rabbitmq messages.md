---
slug: how-to-trace-read-rabbitmq messages
title: How To Trace/Read RabbitMQ Messages
description: Fill me up!
authors: [kbbgl]
tags: [debugging,docker,kubernetes,rabbitmq,trace]
image_url: https://tilsupport.files.wordpress.com/2021/04/1_unyl-2r54_7anewqv0cvxa.png
---

## Introduction

This week, I needed to debug a production issue where one of the critical ReactJS applications happened to load exactly after 60 seconds.

<!-- truncate -->

When I opened the browser development tools network tab, I saw that the request was stuck in Pending state indicating that it was waiting for the server to respond.

As I am already quite familiar with the relevant GraphQL server which works as a backend for this ReactJS application, I knew that the constant 60 seconds response time was no coincidence. It was a hardcoded RabbitMQ RPC (Remote Procedure Call) timeout which is configured systemwide. This means that any microservice in the cluster has 60 seconds to respond to the RPC, and if it doesn’t do so by that interval, the message is dropped. In order to further debug the issue, I needed to somehow figure out which RPC call was stuck as neither the RabbitMQ or the GraphQL services DEBUG/TRACE logs had printed information about these RPCs.

![broker](https://www.cloudamqp.com/img/blog/workflow-rabbitmq.png)

Microservices use the RabbitMQ message broker to transfer messages between them. These messages, when consumed by the destination service, will [`ACK`](https://developer.mozilla.org/en-US/docs/Glossary/TCP_handshake) with a response message to the service which requested the information. Once the reply was accepted by the requesting service, the message is deleted from the queue.

![queue](https://www.cloudamqp.com/img/blog/exchanges-bidings-routing-keys.png)

In some cases, such as investigation of this issue, we would need to be able to review which messages were published and consumed in the RabbitMQ message broker prior to deletion.

This tutorial will explain how to enable message tracing to be able to review the messages published and consumed. It was created because I found it very frustrating that the official RabbitMQ documentation for this feature and its relevant plugin were unclear, not up-to-date and not procedural.

## Disclaimer for Performance Degradation

Enabling message tracing is a very costly operation and increases the processing and memory consumption of the [`erl`](https://www.erlang.org/doc/man/erl.html) process/RabbitMQ container.

You might need to increase the [RabbitMQ memory watermark](https://www.rabbitmq.com/memory.html#threshold) and/or the RabbitMQ container memory/CPU limits in the environment where you’re debugging. It is mandatory to clean up thereafter (see last section on how to do so).

The environment where this issue was occurring was in an Ubuntu server running a Kubernetes 3-node cluster where the RabbitMQ container was running as a ClusterIP Service (which means it’s only exposed to internally). With some minor modifications, it could also be tweaked to be relevant for Windows environments.

Follow [these instruction on how to set up `rabbitmqadmin` for Windows environments](https://www.rabbitmq.com/management-cli.html).

## Configuration (Optional)

We can create a configuration file which we will use to authenticate the commands we’ll send to the RabbitMQ service. This will save us some time as we won’t need to specify mandatory arguments for every command we type and can use the configuration file instead which will include these arguments.
One of the mandatory arguments is the IP where the RabbitMQ service is running. We can retrieve the IP by running the following command:

```bash
kubectl get endpoints rabbitmq-ha -n prod -o=jsonpath='{.subsets[0].addresses[0].ip}'
 
192.168.1.15
```

Let’s create a configuration file:

```bash
cat <<EOT >> /tmp/rmq.conf
# Name of instance we want to connect to
[instance]
 
# Host/IP where the RabbitMQ server is running
hostname = 192.168.1.15
 
# RabbitMQ credentials
username = YOUR_UN
password = YOUR_PW
EOT
nano 
```

Inside the text editor, specify the following:

Notice that the file follows the [Python INI configuration](https://docs.python.org/3/library/configparser.html#supported-ini-file-structure) format. This is because `rabbitmqadmin` is written in Python.

## Enabling Tracing Plugin

RabbitMQ ships by default with a plugin called `rabbitmq_tracing`.
To enable it, we run the following command:

```bash
kubectl exec rabbitmq-ha-0 -c rabbitmq-ha -n prod -- rabbitmq-plugins enable rabbitmq_tracing
```

We can verify that the plugin was enabled by running:

```bash
kubectl exec rabbitmq-ha-0 -c rabbitmq-ha -n prod -- rabbitmq-plugins list
```

Enabling the `rabbitmq_tracing` plugin creates a new exchange named `amq.rabbitmq.trace` where all messages sent to the vhost will be forwarded to.

## Getting `rabbitmqadmin` Binary

In order to read the forwarded messages, we would need to create a new queue where all messages will be sent to and bind this new queue to the `amq.rabbitmq.trace` exchange.

The fastest way to do this is using the `rabbitmqadmin` command line tool.

We can download it from different locations but the best place is from the container itself as it will pair with the exact version of the RabbitMQ server running in the container. To copy the binary to the host machine, we run the following command:

```bash
kubectl cp rabbitmq-ha-0:/var/lib/rabbitmq/mnesia/rabbit@rabbitmq-ha-0.rabbitmq-ha-discovery.prod.svc.cluster.local-plugins-expand/rabbitmq_management-$VERSION/priv/www/cli/rabbitmqadmin -c rabbitmq-ha ~/rabbitmqadmin
```

This will download the `rabbitmqadmin` command line tool to the home (~) directory.
It’s worth to keep in mind that the path might be different in other deployments as this tutorial was tested with a specific RabbitMQ version (represented by `$VERSION`).

To use the `rabbitmqadmin` tool, we need to either (a) specify the configuration file path created in the [Configuration section](#configuration-optional) or (b) the command line arguments for the host, user and password alongside the rest of the commands.

Option (a) will be used for the remainder of the document but if you insist on not creating a configuration file, you can also run the commands in the following format:

```bash
ip=$(kubectl get endpoints -n prod rabbitmq-ha -o=jsonpath='{.subsets[0].addresses[0].ip}')
rabbitmqadmin -H $ip -u YOUR_USERNAME -p YOUR_PASSWORD ...[rest of commands]
```

## Creating RabbitMQ Resources and Bindings

Now that we have the tool to use to interact with the RabbitMQ cluster, we need to do two things before we’re able to read the RPC messages:

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance declare queue name=debug
```

2. Bind the debug queue to the exchange `amq.rabbitmq.trace`. To do this, run the following command:

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance declare binding source=amq.rabbitmq.trace destination=debug routing_key=#
```

The `routing_key=#` parameter ensures that we forward all published, delivered and consumed messages across queues and exchanges to the bound debug queue.

## Reading Messages

Now that all messages are set up to be forwarded to queue debug, we can see how many there are. In the example below we have 317 messages:

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance get queue=debug
+-------------+--------------------+---------------+---------+---------------+------------------+-------------+
| routing_key |      exchange      | message_count | payload | payload_bytes | payload_encoding | redelivered |
+-------------+--------------------+---------------+---------+---------------+------------------+-------------+
| publish.    | amq.rabbitmq.trace | 317           | {}      | 2             | string           | False       |
+-------------+--------------------+---------------+---------+---------------+------------------+-------------+
```

We can also read the actual payload. Here we specify to read the first 5 (includes only the headers, not actual data):

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance get queue=debug count=5
 
+--------------------------------------+--------------------+---------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+---------------+------------------+-------------+
|             routing_key              |      exchange      | message_count |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    payload                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | payload_bytes | payload_encoding | redelivered |
+--------------------------------------+--------------------+---------------+------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+---------------+------------------+-------------+
|...
```

A better way to see the messages would be to connect to the queue as a consumer. We can do this faily easily by using the [`pika` Python library](https://pika.readthedocs.io/en/stable/).
We first need to install it:

```bash
pip install pika
```

Then we can use the following script `receive.py` to connect to it:

```python
#!/usr/bin/env python
import pika, sys, os
 
def main():
    cred = pika.PlainCredentials(username=YOUR_USER, password=YOUR_PASSWORD)
    connection = pika.BlockingConnection(pika.ConnectionParameters(host='192.168.1.15', credentials=cred))
    channel = connection.channel()
 
    channel.queue_declare(queue='debug', durable=True)
 
    def callback(ch, method, properties, body):
         
        print(" [x] Received method '{}', body {}".format(method.routing_key, body.decode()))
 
    channel.basic_consume(queue='debug', on_message_callback=callback, auto_ack=True)
 
    print(' [*] Waiting for messages. To exit press CTRL+C')
    channel.start_consuming()
 
if __name__ == '__main__':
    try:
        main()
    except KeyboardInterrupt:
        print('Interrupted')
        try:
            sys.exit(0)
        except SystemExit:
            os._exit(0)
```

Keep in mind that you need to change the `YOUR_USER` and `YOUR_PASSWORD` values (as well as the `host=IP`) according to your credentials/address.

We can then run the script to begin consuming the messages:

```bash
chmod u+x receive.py
./receive.py
 
[x] Received u'{...}'
```

This is the stage where we would use the payload to troubleshoot our issue. In the specific case I was debugging, we found that the environment had an RPC timeout because of a very large database query to retrieve data for thousands of resources created in the application. We recommended a scaling solution to resolve it.

## Cleanup

Creating the debug queue without consumers means that the messages will get stuck there forever. This could lead to increase in server/container resources, specifically RAM, disk and CPU and in turn, to severe application performance degradation. Thererfore, it’s required to erase all created resources before terminating the session.

We first need to remove the binding between the exchange and the queue to stop forwarding messages to the queue. To do this, we run the following command:

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance delete binding source=amq.rabbitmq.trace destination=debug destination_type=queue properties_key="#"
binding deleted
```

Next step is to purge (i.e. delete/remove) all the messages in the debug queue:

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance purge queue name=debug
queue purged
```

Delete the debug queue:

```bash
rabbitmqadmin -c /tmp/rmq.conf -N instance delete queue name=debug
queue deleted
```

And finally, disable the `rabbitmq_tracing` plugin to remove the `amq.rabbitmq.trace` exchange:

```bash
kubectl exec rabbitmq-ha-0 -c rabbitmq-ha -n prod -- rabbitmq-plugins disable rabbitmq_tracing
```
