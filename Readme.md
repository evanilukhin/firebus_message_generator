# Firebus Message Generator

## Description

*FireBus Message Generator* is part of [FireBus](https://github.com/evanilukhin/firebus) ecosystem. It creates a JSON-messages and puts into a Kafka topic.

## Message Structure

- *uuid* - identifier of the messages group;
- *item* - unique number of the message in the group;
- *created_at* - the timestamp is formatted as "15:12:12.124123";

## Usage

```sh
./firebus_message_generator --count=10000000 --kafka=172.21.0.2:9092 --topic=firebus
```
\- generates `10000000` messages and sends into the topic `firebus`
## Parametres
- *count* - count of messages in the group. All messages have one uuid. Default value: `100`;  
- *kafka* - broker address. Default: `0.0.0.0:9092`;
- *topic* - destination topic. Default: `firebus`.
## Dependencies
+ [kafka](github.com/confluentinc/confluent-kafka-go/kafka)
+ [go.uuid](github.com/satori/go.uuid)

## PS

In this realization, delivery all messages is not guaranteed) (Loss near 20%)
