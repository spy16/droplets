
# Directory Structure

Directory structure is based on [Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).

![Clean Architecture](http://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

Most important part of Clean Architecture is the **Dependency Rule**:

> **source code dependencies can only point inwards**

### 1. `domain/`

Package `domain` represents the `entities` layer from the Clean Architecture. Entities in this layer are
*least likely to change when something external changes*. For example, the rules inside `Validate` methods
are designed in such a way that the these are *absolute requirement* for the entity to belong to the domain.
For example, a `User` object without `Name` does not belong to `domain` of Droplets.

In other words, any feature or business requirement that comes in later, these definitions would still not
change unless the requirement is leading to a domain change.

This package **strictly cannot** have direct dependency on external packages. It can use built-in types,
functions and standard library types/functions.

> `domain` package makes one exception and imports `github.com/spy16/droplets/pkg/errors`. This is because
> of errors are values and are a basic requirement across the application. In other words, the `errors` package
> is used in place of `errors` package from the standard library.

### 2. `usecases/`

Directory (not a package) `usecases` represents the `Use Cases` layer from the Clean Architecture. It encapsulates
and implements all of the use cases of the system by directing the `entities` layer. This layer is expected to change
when a new use case or business requirement is presented.

In Droplets, Use cases are separated as packages based on the entity they primarily operate on. (e.g. `usecases/users` etc.)

Any real use case would also need external entities such as persistence, external service integration etc.
But this layer also **strictly** cannot have direct dependency on external packages. This crossing of boundaries
is done through interfaces.

For example, `users.Registrar` provides functions for registering users which requires storage functionality. But
this cannot directly import a `mongo` or `sql` driver and implement storage functions. It can also not import an
adapter from the `interfaces` package directly both of which would violate `Dependency Rule`. So instead, an interface
`users.Store` is defined which is expected to injected when calling `NewRegistrar`.

**Why is `Store` interface defined in `users` package?**

See [Interfaces](interfaces.md) for conventions around interfaces.


### 3. `interfaces/`

> Should not be confused with Go `interface` keyword.

- Represents the `interface-adapter` layer from the Clean Architecture
- This is the layer that cares about the external world (i.e, external dependencies).
- Interfacing includes:
    - Exposing `usecases` as API (e.g., RPC, GraphQL, REST etc.)
    - Presenting `usecases` to end-user (e.g., GUI, WebApp etc.)
    - Persistence logic (e.g., cache, datastores etc.)
    - Integrating an external service required by `usecases`
- Packages inside this are organized in 2 ways:
    1. Based on the medium they use (e.g., `rest`, `web` etc.)
    2. Based on the external dependency they use (e.g., `mongo`, `redis` etc.)

### 4. `pkg/`

- Contains re-usable packages that is safe to be imported in other projects
- This package should not import anything from `domain`, `interfaces`, `usecases` or their sub-packages


### 5. `web/`

- `web/` is **NOT** a Go package
- Contains web assets such as css, images, templates etc.