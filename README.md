# sqlclass

[![CI](https://github.com/Orfeo42/sqlclass/actions/workflows/ci.yml/badge.svg)](https://github.com/Orfeo42/sqlclass/actions/workflows/ci.yml)

`sqlclass` reports database error branches that raise through anything other than the project's classifying query-error wrapper.

An error branch guarding a database call raises through anything other than the project's query-error wrapper, collapsing duplicate key, foreign key violation, deadlock and data errors into one opaque class. Findings are prefixed `SQLCLASS:`. A guard that branches on one of the wrapper's `Is*` predicates counts as classifying. Database calls nested inside a function literal are ignored, since their error is already propagated out through the closure's own return.

This analyzer encodes a project's own error-wrapping vocabulary: the package that classifies SQL errors, the package that holds the domain error wrappers and sentinels, the dao package and the assert helpers. The names are settings, with defaults taken from the project it was written for. It is only useful where that vocabulary matches, and it is meant to be used as a golangci-lint module plugin rather than as a general-purpose linter. It ships no standalone binary and reads syntax only, so it costs no type-checking.

## Plugin usage

Build a custom golangci-lint with this module compiled in, using a `.custom-gcl.yml`:

```yaml
version: v2.14.0
name: custom-gcl
plugins:
  - module: github.com/Orfeo42/sqlclass
    import: github.com/Orfeo42/sqlclass/plugin
    version: latest
```

Then enable it in `.golangci.yml`:

```yaml
linters:
  enable:
    - sqlclass
  settings:
    custom:
      sqlclass:
        type: module
        description: reports database error branches that raise through anything other than the classifying query-error wrapper
        settings:
          exclude-paths: internal/dao/,_mock.go,_test.go
          sql-utility-paths: src/utility/sql_utility/
```

## Settings

| setting | default | meaning |
| --- | --- | --- |
| `exclude-paths` | — | Comma-separated path fragments; a file whose slash-normalized path contains one is never reported. |
| `sql-utility-paths` | — | Comma-separated path fragments; a file matching one is skipped, so the wrapper's own package does not report itself. |
| `sql-utility-package` | `sql_utility` | Package selector whose `WrapQueryError`/`Classify`/`Is*` predicates count as classifying. |
| `domain-package` | `domain` | Package selector for the `Error`/`WrapError` wrappers and the `Err*` sentinels. |
| `dao-package` | `dao` | Package selector whose calls count as database calls. |
| `assert-package` | `utility` | Package selector for `AssertError`/`AssertErrorWithCode`. |

## Origin

Split from github.com/Orfeo42/nilfield at commit 01c7880, where it shipped inside the droppederr analyzer together with DROPPED (now github.com/Orfeo42/droppederr).

## License

MIT, see `LICENSE`. See `NOTICE.md` for the licensing of a compiled `custom-gcl` binary.
