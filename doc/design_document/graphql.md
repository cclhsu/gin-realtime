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
// GraphQLClientControllerInterface.go
type GraphQLClientControllerInterface interface {
    Query(query string, variables interface{}) error
    Mutate(mutation string, variables interface{}) error
    SetEndpoint(endpoint string) error
}
```

```golang
// GraphQLClientServiceInterface.go
type GraphQLClientServiceInterface interface {
    Initialize() error
    OnResponseReceived(callback func(response interface{})) error
}
```

```golang
// GraphQLServerControllerInterface.go
type GraphQLServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    SetSchema(schema interface{}) error
}
```

```golang
// GraphQLServerServiceInterface.go
type GraphQLServerServiceInterface interface {
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
