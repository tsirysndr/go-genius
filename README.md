<h1 align="center">Welcome to go-genius 👋 *** Work In Progress ***</h1>
<p align="center">
  <a href="https://github.com/tsirysndr/go-genius/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-genius.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-genius">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-genius">
  <img alt="GitHub closed pull requests" src="https://img.shields.io/github/issues-pr-closed-raw/tsirysndr/go-genius">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-genius">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-genius">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-genius">
  <a href="https://github.com/tsirysndr/go-genius/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>

go-genius is a Go client library for accessing the [Genius API](https://docs.genius.com/)

## 🚚 Install

```sh
go get github.com/tsirysndr/go-genius
```

## 🚀 Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-genius"
```

Construct a new Genius client, then use the various services on the client to access different parts of the Genius API. For example:

```Go
client := genius.NewClient("<YOUR ACCESS TOKEN>")
res, _ := client.Search.Get("Kendrick Lamar")
hits, _ := json.Marshal(res)
fmt.Println(string(hits))
```

## Coverage

Currently the following services are supported:

- [x] Search
- [x] Get artist
- [ ] Get documents (songs) for the artist specified
- [x] Get referents
- [ ] Get annotation
- [x] Get song
- [ ] Web Pages

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!
