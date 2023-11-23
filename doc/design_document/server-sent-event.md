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
// ServerSentEventClientControllerInterface.go
type ServerSentEventClientControllerInterface interface {
    ConnectToServer(url string) error
    SubscribeToEvent(event string) error
    UnsubscribeFromEvent(event string) error
    Disconnect() error
}
```

```golang
// ServerSentEventClientServiceInterface.go
type ServerSentEventClientServiceInterface interface {
    Initialize() error
    OnEventReceived(event string, callback func(data interface{})) error
    OnClose(callback func()) error
}
```

```golang
// ServerSentEventServerControllerInterface.go
type ServerSentEventServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    BroadcastEvent(event string, data interface{}) error
}
```

```golang
// ServerSentEventServerServiceInterface.go
type ServerSentEventServerServiceInterface interface {
    Initialize() error
    OnClientConnected(callback func(client interface{})) error
    OnClientDisconnected(callback func(client interface{})) error
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
