# `wasi:clocks` bindings for Go

🕒 Centralized bindings for [`wasi:clocks`](https://github.com/WebAssembly/wasi-clocks) interfaces

<table align=center>
<td>

```
.
└── internal/
    ├── octocat/
    │   └── my-app/
    │       └── my-interface/
    │           └── ...
    └── wasi/
        ├── clocks/ 👈 Replaces this folder
        │   ...
```

</table>

✂️ Reduces the size of generated bindings folder by centralizing bindings

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get github.com/jcbhmr/go-wasi-clocks
```

⚠️ The latest version is v0.2.7. You probably want v0.2.0. Use @v0.2.0 to select it.

```sh
go get github.com/jcbhmr/go-wasi-clocks@v0.2.0
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)
![WebAssembly](https://img.shields.io/badge/WebAssembly-654FF0?style=for-the-badge&logo=WebAssembly&logoColor=FFFFFF)

```go
//go:generate go tool wit-bindgen-go generate --out ./internal/ ./wit/
//go:generate rm -rf ./internal/wasi/io/
//go:generate go tool jet -g "*.go" "<your-package-root>/internal/wasi/io/" "github.com/jcbhmr/go-wasi-clocks/" ./internal/
```

```json
{
    "go.buildTags": "wasip2"
}
```

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)
![WebAssembly](https://img.shields.io/badge/WebAssembly-654FF0?style=for-the-badge&logo=WebAssembly&logoColor=FFFFFF)
