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
// SockIOClientControllerInterface.go
type SocketIOClientControllerInterface interface {
    ConnectToServer(serverURL string) error
    EmitEvent(event string, data interface{}) error
    OnEvent(event string, callback func(data interface{})) error
    Disconnect() error
}
```

```golang
// SockIOClientServiceInterface.go
type SocketIOClientServiceInterface interface {
    Initialize() error
    OnConnected(callback func()) error
    OnDisconnected(callback func()) error
    OnEventReceived(event string, callback func(data interface{})) error
}
```

```golang
// SockIOServerControllerInterface.go
type SocketIOServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    OnConnection(callback func(socket interface{})) error
}
```

```golang
// SockIOServerServiceInterface.go
type SocketIOServerServiceInterface interface {
    Initialize() error
    OnClientConnected(callback func(socket interface{})) error
    OnClientDisconnected(callback func(socket interface{})) error
    OnEventReceived(event string, callback func(socket interface{}, data interface{})) error
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
