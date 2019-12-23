<h1 align="center">Welcome to go-genius 👋 *** Work In Progress ***</h1>
<p align="center">
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

## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_