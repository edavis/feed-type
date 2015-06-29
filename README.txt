# feed-type

feed-type is a small, command-line utility that parses an XML URL and
prints out whether it's an RSS or Atom feed.

It's the end result of trying to learn more about XML parsing, type
assertions, and command line flags in Go.

## Usage

$ feed-type -url http://hnrss.org/newest
rss
$ feed-type -url https://groups.google.com/forum/feed/golang-nuts/topics/atom.xml
atom

## Installation

$ go get github.com/edavis/feed-type
