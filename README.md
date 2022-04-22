[![Go Reference][pkg.go.dev badge]][pkg.go.dev]
[![Go Report Card][goreportcard badge]][goreportcard]

# Base

Base implements Go analogs of the Haskell prelude or [base][base] package.

## Wat?!

The addition of generics in Go 1.18 opened up the possibility of implementing
higher level programming constructs. This library started as an attempt to find
out just how far it could "go".

In some way this enterprise was doomed from the start. Haskell has a highly
expressive type system bedazzled with the standard trappings of ML inspired
languages... Go. Does. Not. Success was not anticipated.

This isn't going to magically turn Go into Haskell (Gaskell?). But if it can
implement even a small part of the Haskell prelude that would be pretty swell.
Possibly some new tricks could be learned along the way (and maybe some things
not to do as well). And those felt like good enough reasons to try.

Besides, it will be such fun!

## Design

> Cthulhu greets you at a gate of the cyclopean city...

### Sum Types

Haskell [data][haskell:data] is represented using a combination of
[type][go:type] structs, interfaces, and receivers. It relies on a trick
involving an interface with an un-exported method.

For example, given this definition for Maybe:

```haskell
data Maybe a = Just a | Nothing
```

The following simulates the general idea:

```go
type Maybe[A any] interface {
  isMaybe()
}

type Just[A any] struct {
  Value A
}

func (j Just[A])isMaybe() {}

type Nothing struct{}

func (n Nothing)isMaybe() {}
```

Because `isMaybe()` is not exported only this package can satisfy it
effectively creating a closed set from the normally open set that interfaces
provide. Only the types `Just[A]` and `Nothing` in this package satisfy the
interface and this fact could be used to exhaustively check the different
constructors with a switch statement.

This pattern is hinted at in the [faq][go:faq:variant_types] and concrete
examples are found in the Go source code for the [ast][go:src:ast:Expr]
package. I found Jerf's blog entry on [sum types][jerf:sum_type] to be
enlightening. It is common enough that a linter exists
[go-sumtype][go-sumtype].

This setup is however not sufficient to ensure everything is well typed.
Consider this:

```go
fmt.Println(Maybe[int](Just[float32]{3}))
```

We would want the compiler to complain that this is invalid, but as far as the
compiler is concerned there's no constraint on the implementation of `Maybe[A]`
being violated and it thinks everything is just awesome.

If we want to ensure that the type parameter is bound and required to match we
need to add it to `isMaybe()`:

```go
type Maybe[A any] interface {
  isMaybe(A)
}
```

This also means we have to add the type paramter to `Nothing` which isn't
exactly a perfect match with the Haskell equivalent, but is better than not
having the type checking.

The full solution looks like:

```go
type Maybe[A any] interface {
	isMaybe(A)
}

type Just[A any] struct {
	Value A
}

func (j Just[A]) isMaybe(_ A) {}

type Nothing[A any] struct{}

func (n Nothing[A]) isMaybe(_ A) {}
```

This now properly [rejects the type mismatch](https://go.dev/play/p/sWIigIgR_yJ):

```go
fmt.Println(Maybe[int](Just[float32]{3}))
```

```
./prog.go:22:25: cannot convert Just[float32]{…} (value of type Just[float32]) to type Maybe[int]:
	Just[float32] does not implement Maybe[int] (wrong type for isMaybe method)
		have isMaybe(_ float32)
		want isMaybe(int)

Go build failed.
```

### Type Classes

### Deriving

### Obey the Laws

---

[base]: https://hackage.haskell.org/package/base-4.16.0.0/docs/index.html
[goreportcard badge]: https://goreportcard.com/badge/github.com/calebcase/base
[goreportcard]: https://goreportcard.com/report/github.com/calebcase/base
[pkg.go.dev badge]: https://pkg.go.dev/badge/github.com/calebcase/base.svg
[pkg.go.dev]: https://pkg.go.dev/github.com/calebcase/base

[haskell:data]: https://wiki.haskell.org/Type#Data_declarations
[haskell:newtype]: https://wiki.haskell.org/Type#Type_and_newtype

[go:type]: https://go.dev/ref/spec#Type_definitions

[go:faq:variant_types]: https://go.dev/doc/faq#variant_types
[go:src:ast:Expr]: https://github.com/golang/go/blob/690ac4071fa3e07113bf371c9e74394ab54d6749/src/go/ast/ast.go#L38-L42

[go-sumtype]: https://github.com/BurntSushi/go-sumtype
[jerf:sum_type]: http://www.jerf.org/iri/post/2917
