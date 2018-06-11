# [flink](https://flink.apache.org/introduction.html)

## Execute sample of [quickstart](https://ci.apache.org/projects/flink/flink-docs-release-1.5/quickstart/setup_quickstart.html) 

```bash
localhost:flink-1.5.0 zhuyu$ ./bin/flink run examples/streaming/SocketWindowWordCount.jar --port 9000
WARNING: An illegal reflective access operation has occurred
WARNING: Illegal reflective access by org.apache.flink.shaded.netty4.io.netty.util.internal.PlatformDependent0 (file:/Users/zhuyu/workspace/tools/flink-1.5.0/lib/flink-dist_2.11-1.5.0.jar) to field java.nio.Buffer.address
WARNING: Please consider reporting this to the maintainers of org.apache.flink.shaded.netty4.io.netty.util.internal.PlatformDependent0
WARNING: Use --illegal-access=warn to enable warnings of further illegal reflective access operations
WARNING: All illegal access operations will be denied in a future release
Starting execution of program

------------------------------------------------------------
 The program finished with the following exception:

org.apache.flink.client.program.ProgramInvocationException: Could not retrieve the execution result.
	at org.apache.flink.client.program.rest.RestClusterClient.submitJob(RestClusterClient.java:258)
	at org.apache.flink.client.program.ClusterClient.run(ClusterClient.java:464)
	at org.apache.flink.streaming.api.environment.StreamContextEnvironment.execute(StreamContextEnvironment.java:66)
	at org.apache.flink.streaming.examples.socket.SocketWindowWordCount.main(SocketWindowWordCount.java:92)
	at java.base/jdk.internal.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
	at java.base/jdk.internal.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
	at java.base/jdk.internal.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
	at java.base/java.lang.reflect.Method.invoke(Method.java:564)
	at org.apache.flink.client.program.PackagedProgram.callMainMethod(PackagedProgram.java:528)
	at org.apache.flink.client.program.PackagedProgram.invokeInteractiveModeForExecution(PackagedProgram.java:420)
	at org.apache.flink.client.program.ClusterClient.run(ClusterClient.java:404)
	at org.apache.flink.client.cli.CliFrontend.executeProgram(CliFrontend.java:781)
	at org.apache.flink.client.cli.CliFrontend.runProgram(CliFrontend.java:275)
	at org.apache.flink.client.cli.CliFrontend.run(CliFrontend.java:210)
	at org.apache.flink.client.cli.CliFrontend.parseParameters(CliFrontend.java:1020)
	at org.apache.flink.client.cli.CliFrontend.lambda$main$9(CliFrontend.java:1096)
	at org.apache.flink.runtime.security.NoOpSecurityContext.runSecured(NoOpSecurityContext.java:30)
	at org.apache.flink.client.cli.CliFrontend.main(CliFrontend.java:1096)
Caused by: org.apache.flink.runtime.concurrent.FutureUtils$RetryException: Could not complete the operation. Exception is not retryable.
	at org.apache.flink.runtime.concurrent.FutureUtils.lambda$retryOperationWithDelay$5(FutureUtils.java:214)
	at java.base/java.util.concurrent.CompletableFuture.uniWhenComplete(CompletableFuture.java:859)
	at java.base/java.util.concurrent.CompletableFuture$UniWhenComplete.tryFire(CompletableFuture.java:837)
	at java.base/java.util.concurrent.CompletableFuture.postComplete(CompletableFuture.java:506)
	at java.base/java.util.concurrent.CompletableFuture.postFire(CompletableFuture.java:610)
	at java.base/java.util.concurrent.CompletableFuture$UniCompose.tryFire(CompletableFuture.java:1085)
	at java.base/java.util.concurrent.CompletableFuture$Completion.run(CompletableFuture.java:478)
	at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1135)
	at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:635)
	at java.base/java.lang.Thread.run(Thread.java:844)
Caused by: java.util.concurrent.CompletionException: org.apache.flink.runtime.rest.util.RestClientException: [Internal server error.]
	at java.base/java.util.concurrent.CompletableFuture.encodeRelay(CompletableFuture.java:367)
	at java.base/java.util.concurrent.CompletableFuture.completeRelay(CompletableFuture.java:376)
	at java.base/java.util.concurrent.CompletableFuture$UniCompose.tryFire(CompletableFuture.java:1074)
	... 4 more
Caused by: org.apache.flink.runtime.rest.util.RestClientException: [Internal server error.]
	at org.apache.flink.runtime.rest.RestClient.parseResponse(RestClient.java:225)
	at org.apache.flink.runtime.rest.RestClient.lambda$submitRequest$3(RestClient.java:209)
	at java.base/java.util.concurrent.CompletableFuture$UniCompose.tryFire(CompletableFuture.java:1072)
	... 4 more
```

Trace `tail -f log/flink-zhuyu-standalonesession-*.log`

```log
2018-06-09 17:55:56,521 INFO  org.apache.flink.runtime.dispatcher.StandaloneDispatcher      - Submitting job 3811e996a347c41b11f3b182a2ebbe24 (Socket Window WordCount).
2018-06-09 17:55:56,522 INFO  org.apache.flink.runtime.rpc.akka.AkkaRpcService              - Starting RPC endpoint for org.apache.flink.runtime.jobmaster.JobMaster at akka://flink/user/jobmanager_7 .
2018-06-09 17:55:56,522 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Initializing job Socket Window WordCount (3811e996a347c41b11f3b182a2ebbe24).
2018-06-09 17:55:56,522 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Using restart strategy NoRestartStrategy for Socket Window WordCount (3811e996a347c41b11f3b182a2ebbe24).
2018-06-09 17:55:56,523 INFO  org.apache.flink.runtime.rpc.akka.AkkaRpcService              - Starting RPC endpoint for org.apache.flink.runtime.jobmaster.slotpool.SlotPool at akka://flink/user/f2e7f062-495c-4f62-a982-8b687d41b4a8 .
2018-06-09 17:55:56,523 INFO  org.apache.flink.runtime.executiongraph.ExecutionGraph        - Job recovers via failover strategy: full graph restart
2018-06-09 17:55:56,524 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Running initialization on master for job Socket Window WordCount (3811e996a347c41b11f3b182a2ebbe24).
2018-06-09 17:55:56,524 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Successfully ran initialization on master in 0 ms.
2018-06-09 17:55:56,524 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - No state backend has been configured, using default (Memory / JobManager) MemoryStateBackend (data in heap memory / checkpoints to JobManager) (checkpoints: 'null', savepoints: 'null', asynchronous: TRUE, maxStateSize: 5242880)
2018-06-09 17:55:56,524 INFO  org.apache.flink.runtime.jobmaster.JobManagerRunner           - JobManager runner for job Socket Window WordCount (3811e996a347c41b11f3b182a2ebbe24) was granted leadership with session id 00000000-0000-0000-0000-000000000000 at akka.tcp://flink@localhost:6123/user/jobmanager_7.
2018-06-09 17:55:56,525 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Starting execution of job Socket Window WordCount (3811e996a347c41b11f3b182a2ebbe24)
2018-06-09 17:55:56,525 INFO  org.apache.flink.runtime.executiongraph.ExecutionGraph        - Job Socket Window WordCount (3811e996a347c41b11f3b182a2ebbe24) switched from state CREATED to RUNNING.
2018-06-09 17:55:56,525 INFO  org.apache.flink.runtime.executiongraph.ExecutionGraph        - Source: Socket Stream -> Flat Map (1/1) (77523c41b4983bc01d8279e14e17c189) switched from CREATED to SCHEDULED.
2018-06-09 17:55:56,525 INFO  org.apache.flink.runtime.executiongraph.ExecutionGraph        - Window(TumblingProcessingTimeWindows(5000), ProcessingTimeTrigger, ReduceFunction$1, PassThroughWindowFunction) -> Sink: Print to Std. Out (1/1) (224044f4edb733fcf2d0cb9fc6a721ff) switched from CREATED to SCHEDULED.
2018-06-09 17:55:56,525 INFO  org.apache.flink.runtime.jobmaster.slotpool.SlotPool          - Cannot serve slot request, no ResourceManager connected. Adding as pending request 180fd0c6e23fd6362ea087706b9af902
2018-06-09 17:55:56,525 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Connecting to ResourceManager akka.tcp://flink@localhost:6123/user/resourcemanager(00000000000000000000000000000000)
2018-06-09 17:55:56,526 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Resolved ResourceManager address, beginning registration
2018-06-09 17:55:56,526 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - Registration at ResourceManager attempt 1 (timeout=100ms)
2018-06-09 17:55:56,526 INFO  org.apache.flink.runtime.resourcemanager.StandaloneResourceManager  - Registering job manager 00000000000000000000000000000000@akka.tcp://flink@localhost:6123/user/jobmanager_7 for job 3811e996a347c41b11f3b182a2ebbe24.
2018-06-09 17:55:56,526 INFO  org.apache.flink.runtime.resourcemanager.StandaloneResourceManager  - Registered job manager 00000000000000000000000000000000@akka.tcp://flink@localhost:6123/user/jobmanager_7 for job 3811e996a347c41b11f3b182a2ebbe24.
2018-06-09 17:55:56,527 INFO  org.apache.flink.runtime.jobmaster.JobMaster                  - JobManager successfully registered at ResourceManager, leader id: 00000000000000000000000000000000.
2018-06-09 17:55:56,527 INFO  org.apache.flink.runtime.jobmaster.slotpool.SlotPool          - Requesting slot with profile ResourceProfile{cpuCores=-1.0, heapMemoryInMB=-1, directMemoryInMB=0, nativeMemoryInMB=0, networkMemoryInMB=0} from resource manager (request = 180fd0c6e23fd6362ea087706b9af902).
2018-06-09 17:55:56,527 INFO  org.apache.flink.runtime.resourcemanager.StandaloneResourceManager  - Request slot with profile ResourceProfile{cpuCores=-1.0, heapMemoryInMB=-1, directMemoryInMB=0, nativeMemoryInMB=0, networkMemoryInMB=0} for job 3811e996a347c41b11f3b182a2ebbe24 with allocation id 1a214c9058910bf3ce068ea47a73b86f.
2018-06-09 17:55:56,531 ERROR org.apache.flink.runtime.rest.handler.job.JobExecutionResultHandler  - Request processing failed.
java.lang.NoClassDefFoundError: javax/xml/bind/DatatypeConverter
	at org.apache.flink.api.common.JobID.fromHexString(JobID.java:110)
	at org.apache.flink.runtime.rest.messages.JobIDPathParameter.convertFromString(JobIDPathParameter.java:36)
	at org.apache.flink.runtime.rest.messages.JobIDPathParameter.convertFromString(JobIDPathParameter.java:26)
	at org.apache.flink.runtime.rest.messages.MessageParameter.resolveFromString(MessageParameter.java:76)
	at org.apache.flink.runtime.rest.handler.HandlerRequest.<init>(HandlerRequest.java:59)
	at org.apache.flink.runtime.rest.AbstractHandler.respondAsLeader(AbstractHandler.java:155)
	at org.apache.flink.runtime.rest.handler.RedirectHandler.lambda$null$0(RedirectHandler.java:139)
	at java.base/java.util.concurrent.CompletableFuture.uniWhenComplete(CompletableFuture.java:859)
	at java.base/java.util.concurrent.CompletableFuture$UniWhenComplete.tryFire(CompletableFuture.java:837)
	at java.base/java.util.concurrent.CompletableFuture$Completion.run(CompletableFuture.java:478)
	at org.apache.flink.shaded.netty4.io.netty.util.concurrent.SingleThreadEventExecutor.runAllTasks(SingleThreadEventExecutor.java:357)
	at org.apache.flink.shaded.netty4.io.netty.channel.nio.NioEventLoop.run(NioEventLoop.java:357)
	at org.apache.flink.shaded.netty4.io.netty.util.concurrent.SingleThreadEventExecutor$2.run(SingleThreadEventExecutor.java:111)
	at org.apache.flink.shaded.netty4.io.netty.util.concurrent.DefaultThreadFactory$DefaultRunnableDecorator.run(DefaultThreadFactory.java:137)
	at java.base/java.lang.Thread.run(Thread.java:844)
```

This problem is about java's Version. flink needs java-1.8

```
// https://ci.apache.org/projects/flink/flink-docs-release-1.5/quickstart/setup_quickstart.html

Flink runs on Linux, Mac OS X, and Windows. To be able to run Flink, the only requirement is to have a working Java 8.x installation. 

```

Reinstall java 1.8 on MacOS

```
https://stackoverflow.com/questions/24342886/how-to-install-java-8-on-mac

```

And restart flink & sample. It's OK.

```bash
localhost:flink-1.5.0 zhuyu$ java -version
java version "1.8.0_172"
Java(TM) SE Runtime Environment (build 1.8.0_172-b11)
Java HotSpot(TM) 64-Bit Server VM (build 25.172-b11, mixed mode)
localhost:flink-1.5.0 zhuyu$ ./bin/flink run examples/streaming/SocketWindowWordCount.jar --port 9000
Starting execution of program
```