# unic

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/unic)](https://goreportcard.com/report/github.com/donatj/unic)
[![GoDoc](https://godoc.org/github.com/donatj/unic?status.svg)](https://godoc.org/github.com/donatj/unic)

Works like UNIX `sort | uniq` except you don't have to call `sort` first.

Works by using Cuckoo Filters - See: https://github.com/seiflotfy/cuckoofilter

## Advantages over `sort | uniq`

### Quicker output, lower memory footprint

`sort` by definitions needs to buffer the entire input before it can begin outputing **anything**. This can use a lot of memory and prevents anything from getting output until the initial process completes.

`unic` uses probabalistic filters (Cuckoo) to determine if the input has been seen before, and can begin output after the first line of input.

### Original item order is kept

Given the list `3 1 2 1 2 3`, compare `sort|uniq` 's output

```bash
$ echo '3\n1\n2\n1\n2\n3' | sort | uniq
1
2
3
```

to `unic`

```bash
echo '3\n1\n2\n1\n2\n3' | unic
3
1
2
```

## Disadvantages

### Probabilistic Filtering

As `unic` works with Cuckoo Filters, there is a very small probability a line will be wrongly marked duplicate. Lines will **never** be incorrectly marked as unique due to the nature of the filter.

In cases where a false positive cannot ever be tolerated, `unic` **should not** be used.

### Not compatible with all of `uniq`'s flags

`unic` by nature does not buffer; thus some of `uniq`'s flags cannot be implemented.

In these cases, you should use `uniq`.

## Installing

### Binaries

See: [releases](https://github.com/donatj/unic/releases)

### Compile

```bash
$ go get -u -v github.com/donatj/unic
```
