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
- [Deploy elasticsearch](#deploy-elasticsearch)
- [Deploy kibana](#deploy-kibana)
- [Deploy logstash](#deploy-logstash)
- [Build service](#build-service)
- [Run service](#run-service)
- [Intreface](#intreface)
- [Reference](#reference)

---

## [](<URL>)

## Install

```bash

```

## Deploy elasticsearch

```bash

```

## Deploy kibana

```bash

```

## Deploy logstash

```bash

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
// ElasticsearchClientControllerInterface.go
type GrpcClientControllerInterface interface {
    CallMethod(methodName string, request interface{}) error
    SetEndpoint(endpoint string) error
}
```

```golang
// ElasticsearchClientServiceInterface.go
type GrpcClientServiceInterface interface {
    Initialize() error
    OnResponseReceived(callback func(response interface{})) error
}
```

```golang
// ElasticsearchServerControllerInterface.go
type GrpcServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    AddService(service interface{}) error
}
```

```golang
// ElasticsearchServerServiceInterface.go
type GrpcServerServiceInterface interface {
    Initialize() error
    OnRequestReceived(callback func(request interface{})) error
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
