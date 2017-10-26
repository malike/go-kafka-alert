
# Basic Overview

A Go application that feeds of data from Apache Kafka to send SMS,EMAIL or connects via webhook.

1. **Notification Service** : It can  be used as is for just notification.By reacting to events pushed to Apache Kafka.

2. **Custom Watcher** : Where it would work with [elasticsearch kafka watch]( https://malike.github.io/elasticsearch-kafka-watch/) to send notification once there's _hit_ in elasticsearch.

3. **Scheduled Reports** : Uses [elasticsearch report engine](http://malike.github.io/elasticsearch-report-engine) to send scheduled reports as PDF,HTML or CSV by email.


#### SMS

Connects via Twilio to send sms messages


#### Email
<br/>


#### API (Webhook)
<br/>


# Setup

#### Configuration 

The app is meant to be a light-weight application.  Find a [sample configuration](https://github.com/malike/go-kafka-alert/blob/master/configuration.json) file,which is kept in memory, to get app running:


           {
             "workers": 4,
             "logFileLocation": "/var/log",
             "log": true,
             "kafkaConfig": {
               "bootstrapServers": "localhost:2181",
               "kafkaTopic": "go-kafka-event-stream",
               "kafkaTopicConfig": "earliest",
               "kafkaGroupId": "consumerGroupOne",
               "kafkaTimeout": 5000
             },
             "smsConfig": {
               "twilioAccountId": "Malike",
               "twilioAuthToken": "Malike",
               "smsSender": "+15005550006"
             },
             "emailConfig": {
               "smtpServerHost": "smtp.gmail.com",
               "tls": true,
               "smtpServerPort": 465,
               "emailSender": "Sender",
               "emailFrom": "youreamail@gmail.com",
               "emailAuthUserName": "youreamail@gmail.com",
               "emailAuthPassword": "xxxxxx"
             },
             "dbConfig": {
               "mongoHost": "localhost",
               "mongoPort": 27017,
               "mongoDBUsername": "",
               "mongoDBPassword": "",
               "mongoDB": "go_kafka_alert",
               "collection": "message"
             },
             "templates": {
               "APPFLAG_SMS": "User {{.UnmappedData.UserName}} has failed to execute service {{.UnmappedData.ServiceName}} {{.UnmappedData.FailureCount}} times in the past {{.UnmappedData.FailureDuration}} minutes",
               "SERVICEHEALTH_SMS": "Service {{.UnmappedData.ServiceName}} has failed execution {{.UnmappedData.FailureCount}} in the past {{.UnmappedData.FailureDuration}} minutes",
               "SUBSCRIPTION_SMS": "Hello {{.UnmappedData.Name}}, Thanks for subscribing to {{.UnmappedData.ItemName}}",
               "SUBSCRIPTION_EMAIL": "<html><head></head><body> Hello {{.UnmappedData.Name}}, Thanks for subscribing to {{.UnmappedData.ItemName}} </body></html>",
               "REPORTATTACHED_EMAIL": "<html><head></head><body> Hello {{.UnmappedData.Name}}, Find attached report for {{.UnmappedData.ItemName}} </body></html>",
               "REPORTEMBEDED_EMAIL": "{{.UnmappedData.Content}}"
             }
           }


<br/>

**i. kafkaConfig**
[Apache Kafka]() configuration. Note you can comma separate the value for `bootstrapServers` nodes if you have multiple nodes. 
Example `127.0.0.1:2181,127.0.0.2:2181`. 
For the other [Apache Kafka]() configurations I'm assuming you already know how what they mean. Read the Apache Kafka docs if you want to know more. The project uses the [go kafka library](https://github.com/confluentinc/confluent-kafka-go) by Confluent. 
<br/>


**ii. smsConfig**
This is where configuration for your [twilio account](https://www.twilio.com/) are. This would enable sending SMS notifications. The project uses the [twilio sms config](https://github.com/sfreiberg/gotwilio).
<br/>

**iii. emailConfig**
This is where configuration for your [email smtp]() would be. This would enable sending EMAIL notifications. It uses [http://gopkg.in/gomail.v2](http://gopkg.in/gomail.v2)
<br/>

**iv. dbConfig**
Messages sent out are stored for auditing purposes. Together with the response from twilio or your smtp gateway. This configuration stores them in MongoDB. Uses [this](gopkg.in/mgo.v2/bson) mongodb library for Go.
<br/>

**iv. templates**
These are the messaging templates configured for all the alert types. Follow [this](https://gohugo.io/templates/introduction/) to learn how to create your templates. The templates are stored as maps to use *_O(1)_* when finding a template. The key of the map follows this convention `{{EventType}}`+`_`+`{{Delivery Channel}}`. This means an SMS for EventType, SUBSCRIPTION would be `SUBSCRIPTION_SMS` 
<br/>


#### Use Case 1 : Notification Service
<br/>

#### Use Case 2 : Custom Watcher For ElasticSearch
<br/>

#### Use Case 3 : Scheduled Reports For ElasticSearch Data

   **i. Embedded Reports**
    <br/>

   **ii. CSV/PDF Attached Reports**
    <br/>

  
# Download
| Version  |
| -------- |
| [0.1-Prelease Tag]()   |


## Contribute

Contributions are always welcome!
Please read the [contribution guidelines](CONTRIBUTING.md) first.

## Code of Conduct

Please read [this](CODE_OF_CONDUCT.md).

## License

[GNU General Public License v3.0](https://github.com/malike/go-kafka-alert/blob/master/LICENSE)





