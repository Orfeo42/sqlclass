# Notice

The `sqlclass` analyzer source in this repository is MIT-licensed, see `LICENSE`.

This repository does not publish a binary. The analyzer is meant to be compiled into a custom golangci-lint build through `golangci-lint custom` and a `.custom-gcl.yml` that references this module.

golangci-lint is licensed under GPL-3.0. A `custom-gcl` binary statically links golangci-lint with this plugin, so the combined binary is a derivative work and is distributed under GPL-3.0 by whoever builds and distributes it.

Corresponding source:

- this repository, `github.com/Orfeo42/sqlclass`, for the plugin
- `github.com/golangci/golangci-lint`, for golangci-lint itself
