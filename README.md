# go-athndl

Package **athndl** provides an implementation of **BlueSky**'s **AT-Protocol** **handle** , for the Go programming language.

I.e., **handle*s* are the things that look like these:

* `@reiver.bsky.social`
* `@pfrazee.com`

Although technically, the **handle** part of those does not include the `@`.
I.e.:

* `reiver.bsky.social`
* `pfrazee.com`

**AT-Protocol** **handles** are defined here:
https://atproto.com/specs/handle

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-athndl

[![GoDoc](https://godoc.org/github.com/reiver/go-athndl?status.svg)](https://godoc.org/github.com/reiver/go-athndl)

## Import

To import package **athndl** use `import` code like the follownig:
```
import "github.com/reiver/go-athndl"
```

## Installation

To install package **athndl** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-athndl
```

## Author

Package **athndl** was written by [Charles Iliya Krempeaux](http://reiver.link)
