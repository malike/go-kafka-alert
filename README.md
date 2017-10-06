# go-kafka-alert
A Go application that feeds of data from Apache Kafka to send SMS,EMAIL or connects via webhook.

It can  be used as is for just notification based of events on Apache Kafka or used as

1. Custom Elastic Watcher : Where it would work with [elasticsearch kafka watch](https://github.com/malike/elasticsearch-kafka-watch) to send notification once there's _hit_

2. Report Ops : Uses [elasticsearch report engine](https://github.com/malike/elasticsearch-report-engine) to send scheduled reports as PDF,HTML or CSV by email.



#### SMS

Connects via Twilio to send sms messages


#### Email

