
# Interfaces

Following are some best practices for using interfaces:

1. Define small interfaces with well defined scope
   - Single-method interfaces are ideal (e.g. `io.Reader`, `io.Writer` etc.)
   - [Bigger the interface, weaker the abstraction - Go Proverbs by Rob Pike](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=5m17s)
2. Accept interfaces, return structs
   - Interfaces should be defined where they are used [Read More](#where-should-i-define-the-interface-)

## Where should I define the interface ?

From [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#interfaces):

> Go interfaces generally belong in the package that uses values of the interface type, not the
> package that implements those values.

Interfaces are contracts that should be used to define the minimal requirement of a client
to execute its functionality. In other words, client defines what it needs and not the
implementor. So, interfaces should generally be defined on the client side. This is also inline
with the [Interface Segregation Principle](https://en.wikipedia.org/wiki/Interface_segregation_principle)
from [SOLID](https://en.wikipedia.org/wiki/SOLID) principles.

A **bad** pattern that shows up quite a lot:

```go
package producer

func NewThinger() Thinger {
    return defaultThinger{ … }
}

type Thinger interface {
    Thing() bool
}


type defaultThinger struct{ … }
func (t defaultThinger) Thing() bool { … }
```

### Why is this bad?

Go uses [Structural Type System](https://en.wikipedia.org/wiki/Structural_type_system) as opposed to
[Nominal Type System](https://en.wikipedia.org/wiki/Nominal_type_system) used in other static languages
like `Java`, `C#` etc. This simply means that a type `MyType` does not need to add `implements Doer` clause
to be compatible with an interface `Doer`. `MyType` is compatible with `Doer` interface if it has all the
methods defined in `Doer`.

Read following articles for more information:

1. https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a
2. https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8

This also provides an interesting power to Go interfaces. Clients are truly free to define interfaces when they
need to. For example consider the following function:

```go
func writeData(f *os.File, data string) {
    f.Write([]byte(data))
}
```

Let's assume after sometime a new feature requirement which requires us to write to a tcp connection. One
thing we could do is define a new function:

```go
func writeDataToTCPCon(con *net.TCPConn, data string) {
    con.Write([]byte(data))
}
```

But this approach is tedious and will grow out of control quickly as new requirements are added. Also, different
writers cannot be injected into other entities easily. But instead, you can simply refactor the `writeData` function
as below:

```go
type writer interface {
    Write([]byte) (int, error)
}

func writeData(wr writer, data string) {
    wr.Write([]byte(data))
}
```

Refactored `writeData` will continue to work with our existing code that is passing `*os.File` since it
implements `writer`. In addition, `writeData` function can now accept anything that implements `writer`
which includes `os.File`, `net.TCPConn`, `http.ResponseWriter` etc. (And every single Go entity in the
**entire world** that has a method `Write([]byte) (int, error)`)

Note that, this pattern is *not possible in other languages*. Because, after refactoring `writeData` to
accept a new interface `writer`, you need to refactor all the classes you want to use with `writeData` to
have `implements writer` in their declarations.

Another advantage is that client is free to define the subset of features it requires instead of accepting
more than it needs.


