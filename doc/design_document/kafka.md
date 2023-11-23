---
title: '[MyLabel] <MyTitle>'
description: '<MyTemplate> for <MY_PROJECT>'
labels: 'kind/MyLabel'
priority: ''
assignees: ''
status: ''
date: 2022-01-01T00:00:00+08:00
url: ''
doc: ''
github: ''
ticket:
draft: true
---

# {{ TOPIC }} <!-- omit in toc -->

- [](#)
- [Install](#install)
- [Build service](#build-service)
- [Run service](#run-service)
- [Intreface](#intreface)
- [Reference](#reference)

---

## [](<URL>)

## Install

```bash
npm install @nestjs/microservices kafkajs class-validator class-transformer
```

## Build service

```bash

```

## Run service

```bash

```

---

## Intreface

```golang
// KafkaClientControllerInterface.go
type KafkaClientControllerInterface interface {
    ProduceMessage(topic string, message string) error
    ConsumeFromTopic(topic string) error
    CloseConnection() error
}
```

```golang
// KafkaClientServiceInterface.go
type KafkaClientServiceInterface interface {
    Initialize() error
    OnMessageReceived(callback func(message string)) error
    OnClose(callback func()) error
}
```

```golang
// KafkaServerControllerInterface.go
type KafkaServerControllerInterface interface {
    StartServer() error
    StopServer() error
    CreateTopic(topic string) error
}
```

```golang
// KafkaServerServiceInterface.go
type KafkaServerServiceInterface interface {
    Initialize() error
    OnMessageReceived(callback func(topic string, message string)) error
    OnTopicCreated(callback func(topic string)) error
}
```

---

## Reference

- [](<URL>)
- [Company/Project](<https://{{ GITHUB_PROJECT }}.io/>)
- [Documentation](<https://{{ GITHUB_PROJECT }}.io/doc>)
- [Github](<https://github.com/{{ GITHUB_USER }}/{{ GITHUB_PROJECT }}>)
- [Wikipedia](<https://en.wikipedia.org/wiki/{{ TOPIC }}>)

---
