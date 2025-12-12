# Daisy

Daisy is a lightweight + zero dependency middleware chainer for `net/http`
Syntax is meant to be extremely readable, following a first-in first-executed pattern

## Installation

```bash
go get github.com/17ColinMiPerry/daisy
```

## Why Daisy?

Vanilla Go uses nesting, which can become hard to read:

```Go
// hard to read. inside-out execution order
http.Handle("/", MiddlewareThree(MiddlewareTwo(MiddlewareOne(myHandler))))
```

Go + Daisy is readable and easy to use:

```Go
// easy to read. reads left-to-right
http.Handle("/", daisy.chain(MiddlewareOne, MiddlewareTwo, MiddlewareThree).To(myHandler))
```
