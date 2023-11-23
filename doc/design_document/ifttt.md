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
// IftttClientControllerInterface.go
type IftttClientControllerInterface interface {
    TriggerEvent(event string, data interface{}) error
    SetWebhookKey(apiKey string) error
}
```

```golang
// IftttClientServiceInterface.go
type IftttClientServiceInterface interface {
    Initialize() error
    OnEventTriggered(callback func(response interface{})) error
}
```

```golang
// IftttServerControllerInterface.go
type IftttServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    SetEventHandler(handler func(event string, data interface{})) error
}
```

```golang
// IftttServerServiceInterface.go
type IftttServerServiceInterface interface {
    Initialize() error
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
