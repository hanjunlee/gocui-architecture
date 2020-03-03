# gocui architecture

## Goal

Benchmarking `k9s` logic and also adapt clean architecture into CUI.

## Structure

The manager in `ui` package has the responsibility for UI stuff and it utilize the `service` pkg to manufacture data stuff.

```shell
├── internal
│   └── ui
│       ├── editor.go
│       └── manager.go
└── pkg
    ├── service
    │   ├── interface.go
    │   └── service.go
...
```
