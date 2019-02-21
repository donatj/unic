# unic

[![Go Report Card](https://goreportcard.com/badge/github.com/donatj/unic)](https://goreportcard.com/report/github.com/donatj/unic)

Works like UNIX `sort | uniq` except you don't have to call `sort` first.

Works by using Cuckoo Filters - See: https://github.com/seiflotfy/cuckoofilter

## Advantages over `sort | uniq`

### Quicker Output / Lower Memory Footprint

`sort` by definitions needs to buffer the entire input before it can begin outputing **anything**. This can use a lot of memory and prevents anything from getting output until the initial process completes.

`unic` on the other hand uses probabalistic filters (Cuckoo) to determine if the input has been seen before, and can begin output after the first line of input.

### You don't lose original item order

Compare `sort|uniq`

```bash
$ echo '3\n1\n2\n1\n2\n3' | sort | uniq
1
2
3
```

vs. `unic`

```bash
echo '3\n1\n2\n1\n2\n3' | unic
3
1
2
```

## Disadvantages

### Probabilistic Filtering

As `unic` utilizes Cuckoo Filters

### Not compatbiel with all of `uniq`'s flags

As `unic` does not buffer, not all of `uniq`'s flags can be implemented. This is by design however, and in such cases `uniq` should be used.

## Installing

### Binaries

See: [releases](https://github.com/donatj/unic/releases)

### Compile

```bash
$ go get -u -v github.com/donatj/unic
```
