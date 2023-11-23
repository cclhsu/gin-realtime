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

## Intreface

```golang
// WebRTCClientControllerInterface.go
type WebRTCClientControllerInterface interface {
    CreatePeerConnection(config RTCConfiguration) error
    CreateDataChannel(label string, options RTCDataChannelInit) error
    AddIceCandidate(candidate RTCIceCandidateInit) error
    SetLocalDescription(description RTCSessionDescriptionInit) error
    SetRemoteDescription(description RTCSessionDescriptionInit) error
    CloseConnection() error
}
```

```golang
// WebRTCClientServiceInterface.go
type WebRTCClientServiceInterface interface {
    Initialize() error
    OnDataChannelCreated(callback func(channel RTCDataChannel) error
    OnIceCandidate(callback func(candidate RTCIceCandidate) error
    OnConnectionClosed(callback func() error
}
```

```golang
// WebRTCServerControllerInterface.go
type WebRTCServerControllerInterface interface {
    StartServer(port int) error
    StopServer() error
    CreatePeerConnection(config RTCConfiguration) error
    AddIceCandidate(clientId string, candidate RTCIceCandidateInit) error
    SetLocalDescription(clientId string, description RTCSessionDescriptionInit) error
    SetRemoteDescription(clientId string, description RTCSessionDescriptionInit) error
}
```

```golang
// WebRTCServerServiceInterface.go
type WebRTCServerServiceInterface interface {
    Initialize() error
    OnClientConnected(callback func(clientId string) error
    OnClientDisconnected(callback func(clientId string) error
    OnDataChannelCreated(callback func(clientId string, channel RTCDataChannel) error
    OnIceCandidate(callback func(clientId string, candidate RTCIceCandidate) error
    OnConnectionClosed(callback func(clientId string) error
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
