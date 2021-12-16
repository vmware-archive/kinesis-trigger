## WARNING: Kubeless is no longer actively maintained by VMware.

VMware has made the difficult decision to stop driving this project and therefore we will no longer actively respond to issues or pull requests. If you would like to take over maintaining this project independently from VMware, please let us know so we can add a link to your forked project here.

Thank You.

# Kinesis-trigger

A Kubeless _Trigger_ represents an event source the functions can be associated with it. When an event occurs in the event source, Kubeless will ensure that the associated functions are invoked. __Kinesis-trigger__ addon to Kubeless adds support for a AWS Kinesis data streaming platform as trigger to Kubeless. An AWS Kinesis stream can be associated with one or more Kubeless functions. Kubeless functions associated with a Kinesis stream are triggerd as and when records get pubslished to the stream.

Please refer to the [documentation](https://kubeless.io/docs/streaming-functions/) on how to use AWS Kinesis stream as trigger with Kubeless.
