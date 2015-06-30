# feed-type

feed-type is a small, command-line utility that loads an XML URL and
prints out whether it's an RSS or Atom feed.

It's the end result of trying to learn more about XML parsing, type
assertions, and command line flags in Go.

## Usage

```shell
$ feed-type -url http://hnrss.org/newest
rss
$ feed-type -url http://blog.golang.org/feed.atom
atom
```

## Installation

```shell
$ go get github.com/edavis/feed-type
```
