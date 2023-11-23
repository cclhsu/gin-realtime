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
// WebpushClientControllerInterface.go
type WebpushClientControllerInterface interface {
    SubscribeUser(endpoint string, keys interface{}) error
    UnsubscribeUser(endpoint string) error
    SendNotification(subscription interface{}, payload interface{}) error
}
```

```golang
// WebpushClientServiceInterface.go
type WebpushClientServiceInterface interface {
    Initialize() error
    OnNotificationReceived(callback func(notification interface{})) error
}
```

```golang
// WebpushServerControllerInterface.go
type WebpushServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
}
```

```golang
// WebpushServerServiceInterface.go
type WebpushServerServiceInterface interface {
    Initialize() error
    OnUserSubscribed(callback func(subscription interface{})) error
    OnUserUnsubscribed(callback func(endpoint string)) error
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
