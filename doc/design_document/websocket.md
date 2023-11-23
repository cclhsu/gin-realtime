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
// WebsocketClientControllerInterface.go
type WebsocketClientControllerInterface interface {
    Connect(url string) error
    SendMessage(message string) error
    CloseConnection() error
}
```

```golang
// WebsocketClientServiceInterface.go
type WebsocketClientServiceInterface interface {
    Initialize() error
    OnMessageReceived(callback func(message string)) error
    OnClose(callback func()) error
}
```

```golang
// WebsocketServerControllerInterface.go
type WebsocketServerControllerInterface interface {
    StartServer(port int) error
    BroadcastMessage(message string) error
    StopServer() error
}
```

```golang
// WebsocketServerServiceInterface.go
type WebsocketServerServiceInterface interface {
    Initialize() error
    OnClientConnect(callback func(client *websocket.Conn)) error
    OnClientDisconnect(callback func(client *websocket.Conn)) error
    OnMessageReceived(callback func(client *websocket.Conn, message string)) error
}
```

---

## Reference

- [](<URL>)
- [Company/Project](<https://{{ GITHUB_PROJECT }}.io/>)
- [Documentation](<https://{{ GITHUB_PROJECT }}.io/doc>)
- [Github](<https://github.com/{{ GITHUB_USER }}/{{ GITHUB_PROJECT }}>)
- [Wikipedia](<https://en.wikipedia.org/wiki/{{ TOPIC }}>)
- https://gist.github.com/crosstyan/47e7d3fa1b9e4716c0d6c76760a4a70c
- https://lwebapp.com/en/post/go-websocket-chat-server
- https://lwebapp.com/en/post/go-websocket-simple-server
- https://hoohoo.top/blog/20220320172715-go-websocket/
- https://juejin.cn/post/7103737973782511646
- https://websocketking.com/
- https://betterprogramming.pub/building-web-chat-with-go-and-websockets-312f459c001a
- https://cloud.tencent.com/developer/article/1774583

---
