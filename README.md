# Workers

[![Build Status](http://github.dronehippie.de/api/badges/webhippie/workers/status.svg)](http://github.dronehippie.de/webhippie/workers)
[![Go Doc](https://godoc.org/github.com/webhippie/workers?status.svg)](http://godoc.org/github.com/webhippie/workers)
[![Go Report](https://goreportcard.com/badge/github.com/webhippie/workers)](https://goreportcard.com/report/github.com/webhippie/workers)
[![Sourcegraph](https://sourcegraph.com/github.com/webhippie/workers/-/badge.svg)](https://sourcegraph.com/github.com/webhippie/workers?badge)
[![Join the chat at https://gitter.im/webhippie/general](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/webhippie/general)
[![Stories in Ready](https://badge.waffle.io/webhippie/workers.svg?label=ready&title=Ready)](http://waffle.io/webhippie/workers)

TBD


## Usage

TBD


## Development

Make sure you have a working Go environment, for further reference or a guide take a look at the [install instructions](http://golang.org/doc/install.html). As this project relies on vendoring of the dependencies and we are not exporting `GO15VENDOREXPERIMENT=1` within our makefile you have to use a Go version `>= 1.6`. It is also possible to just simply execute the `go get -d github.com/webhippie/workers` command, but we prefer to use our `Makefile`:

```bash
go get -d github.com/webhippie/workers
cd $GOPATH/src/github.com/webhippie/workers
make retool sync clean test
```


## Contributing

Fork -> Patch -> Push -> Pull Request


## Authors

* [Thomas Boerger](https://github.com/tboerger)


## License

Apache-2.0


## Copyright

```
Copyright (c) 2017 Thomas Boerger <http://www.webhippie.de>
```
