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
// RedisClientControllerInterface.go
type RedisClientControllerInterface interface {
    Set(key string, value string) error
    Get(key string) error
    Del(key string) error
    Disconnect() error
}
```

```golang
// RedisClientServiceInterface.go
type RedisClientServiceInterface interface {
    Connect() error
    OnValueReceived(callback func(key string, value string)) error
    OnDisconnected(callback func()) error
}
```

```golang
// RedisServerControllerInterface.go
type RedisServerControllerInterface interface {
    StartServer() error
    StopServer() error
    Set(key string, value string) error
    Get(key string) error
    Del(key string) error
}
```

```golang
// RedisServerServiceInterface.go
type RedisServerServiceInterface interface {
    Initialize() error
    OnValueSet(callback func(key string, value string)) error
    OnValueDeleted(callback func(key string)) error
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
