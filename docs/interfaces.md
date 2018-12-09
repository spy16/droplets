# Interfaces

## Where should I define the interface ?

From [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments#interfaces):

> Go interfaces generally belong in the package that uses values of the interface type, not the
> package that implements those values.

Interfaces are contracts that should be used to define the minimal requirement of a client
to execute its functionality. In other words, client defines what it needs and not the
implementor. So, interfaces should always be defined on the client side. This is also inline
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

Refer following articles for more information:

1. https://medium.com/@cep21/preemptive-interface-anti-pattern-in-go-54c18ac0668a
2. https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8
