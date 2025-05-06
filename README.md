# go-athandle

Package **athandle** provides an implementation of **BlueSky**'s **AT-Protocol** **handle** , for the Go programming language.

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

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-athandle

[![GoDoc](https://godoc.org/github.com/reiver/go-athandle?status.svg)](https://godoc.org/github.com/reiver/go-athandle)

## Examples

To normalize a **handle** do something similar to:

```golang
import "github.com/reiver/go-athandle"

// ...

var handle string = "Example.COM"

normalizedHandle := athandle.Normalize(handle)

// normalizedHandle == "example.com"
```

To validate a **handle** do something similar to:

```golang
import "github.com/reiver/go-athandle"

// ...

var handle string = "Example.COM"

err := athandle.Validate(handle)

// returns an error if the handle is invalid
```

## Import

To import package **athandle** use `import` code like the follownig:
```
import "github.com/reiver/go-athandle"
```

## Installation

To install package **athandle** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-athandle
```

## Author

Package **athandle** was written by [Charles Iliya Krempeaux](http://reiver.link)
