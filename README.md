# Legolas

[Legolas](https://github.com/kaleworsley/legolas) is an example of how to structure a web application written in [go](https://golang.org/).

## Dependencies

[go-bindata](https://github.com/jteeuwen/go-bindata) is required to embed the
templates in the binary.

Go dependencies are managed with [dep](https://github.com/golang/dep).

Frontend dependencies are managed with [yarn](https://yarnpkg.com/).

## Building

```
go generate
go install ./...
```

## Usage

```
Usage of legolas:
  -host string
    	http host [LEGOLAS_HOST] (default "127.0.0.1")
  -port string
    	http port [LEGOLAS_PORT] (default "8080")
  -public string
    	path to public directory [LEGOLAS_PUBLIC]
  -templates string
    	path to templates directory [LEGOLAS_TEMPLATES]
```

## License

Copyright 2017, Kale Worsley.

Legolas is made available under the MIT License. See [LICENSE](LICENSE) for details.