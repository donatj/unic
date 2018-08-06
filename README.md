# unic

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/unic)](https://goreportcard.com/report/github.com/donatj/unic)

Works like UNIX `sort | uniq` except you don't have to call `sort` first, thus you don't lose item order.

Works by using a Cuckoo Filter - See: https://github.com/seiflotfy/cuckoofilter

## Installing

### Binaries

See: [releases](https://github.com/donatj/unic/releases)

### Compile

```bash
$ go get -u -v github.com/donatj/unic
```
