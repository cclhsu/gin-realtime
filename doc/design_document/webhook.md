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
// WebhookClientControllerInterface.go
type WebhookClientControllerInterface interface {
    SendRequest(url string, data interface{}) error
    SetHeaders(headers map[string]string) error
    SetTimeout(timeout int) error
    CloseConnection() error
}
```

```golang
// WebhookClientServiceInterface.go
type WebhookClientServiceInterface interface {
    Initialize() error
    OnRequestSent(callback func(response interface{})) error
    OnClose(callback func()) error
}
```

```golang
// WebhookServerControllerInterface.go
type WebhookServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    SetRequestHandler(handler func(request interface{})) error
}
```

```golang
// WebhookServerServiceInterface.go
type WebhookServerServiceInterface interface {
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
