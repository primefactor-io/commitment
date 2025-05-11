# Commitment

Implementations of various [Commitment Schemes](https://en.wikipedia.org/wiki/Commitment_scheme).

Right now a [Hash-Based Commitment Scheme](https://muens.io/hash-based-commitment-scheme/) using the SHA-256 [hash function](https://muens.io/hash-function/) is available.

## Setup

1. `git clone <url>`
2. `asdf install` (optional)
3. `go test -count 1 -race ./...`

## Useful Commands

```sh
go run <package-path>
go build [<package-path>]

go test [<package-path>][/...] [-v] [-cover] [-race] [-short] [-parallel <number>]
go test -bench=. [<package-path>] [-count <number>] [-benchmem] [-benchtime 2s] [-memprofile <name>]

go test -coverprofile <name> [<package-path>]
go tool cover -html <name>
go tool cover -func <name>

go fmt [<package-path>]

go mod init [<module-path>]
go mod tidy
```

## Useful Resources

- [Go - Learn](https://go.dev/learn)
- [Go - Documentation](https://go.dev/doc)
- [Go - A Tour of Go](https://go.dev/tour)
- [Go - Effective Go](https://go.dev/doc/effective_go)
- [Go - Playground](https://go.dev/play)
- [Go by Example](https://gobyexample.com)
- [100 Go Mistakes and How to Avoid Them](https://100go.co)
