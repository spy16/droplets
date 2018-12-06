> WIP

# droplets

[![GoDoc](https://godoc.org/github.com/spy16/droplets?status.svg)](https://godoc.org/github.com/spy16/droplets) [![Go Report Card](https://goreportcard.com/badge/github.com/spy16/droplets)](https://goreportcard.com/report/github.com/spy16/droplets)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fspy16%2Fdroplets.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fspy16%2Fdroplets?ref=badge_shield)

A platform for Gophers similar to the awesome [Golang News](http://golangnews.com).

## Why?

Droplets is NOT built because there is no such platform nor is it built to solve problems
with existing platforms.

New gophers often struggle with deciding how to structure their applications and also miss certain
important conventions (e.g., `Accept Interfaces, Return Structs`).

Droplets is built to showcase:

1. Application of [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) and [EffectiveGo](https://golang.org/doc/effective_go.html)
2. Usage of [Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


## Organization

### Directory Structure

Directory structure is based on [Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).


#### 1. `domain/`

- represents the `entities` layer from the Clean Architecture
- contains different core entity definitions and core validation logic
- this package **strictly** cannot have direct dependency on external packages

### 2. `usecases/`

- represents the `usecases` layer from the Clean Architecture
- Usecases layer builds different business oriented usecases using the entities provided by `entities` layer
- Any real use case would also need external entities such as persistence, external service integration etc.
  But this layer also **strictly** cannot have direct dependency on external packages. This crossing of boundaries
  is done through interfaces.

### 3. `interfaces/`

- represents the `interface-adapter` layer from the Clean Architecture
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

- contains re-usable packages that is safe to be imported in other projects
- this package should not import anything from `domain`, `interfaces`, `usecases` or their sub-packages


#### 5. `web/`

- `web/` is **NOT** a Go package
- contains web assets such as css, images, templates etc.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fspy16%2Fdroplets.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fspy16%2Fdroplets?ref=badge_large)
