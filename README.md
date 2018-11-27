> WIP

# droplets

[![GoDoc](https://godoc.org/github.com/spy16/droplets?status.svg)](https://godoc.org/github.com/spy16/droplets) [![Go Report Card](https://goreportcard.com/badge/github.com/spy16/droplets)](https://goreportcard.com/report/github.com/spy16/droplets)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fspy16%2Fdroplets.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fspy16%2Fdroplets?ref=badge_shield)

A platform for Gophers similar to the awesome [Golang News](http://golangnews.com).

## Why ?

Droplets is NOT built because there is no such platform nor is it built to solve problems
with existing platforms.

New gophers often struggle with deciding how to structure their applications and also miss certain
important conventions (e.g., `Accept Interfaces, Return Structs`).

Droplets is built to showcase:

1. A manageable [Project Layout](https://github.com/golang-standards/project-layout/)
2. Application of [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) and [EffectiveGo](https://golang.org/doc/effective_go.html)
3. Usage of [Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)


## Organization

### Directory Structure

Directory structure is based on [Project Layout](https://github.com/golang-standards/project-layout/).

#### 1. `internal/`

- contains non-reusable parts of the project
- directories inside `internal/` are organized as per [Clean Architecture](http://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
    - `domain/` contains different core entity definitions and represents the entities layer.
    - `usecases/` contains different busniess logic built around entities and represents the usecases layer.
    - `delivery/` exposes the usecases as API or web app and represents interface-adapter layer.
    - `stores/` provides storage functions for domain entities and is also part of interface-adapter layer.

#### 2. `pkg/`

- contains re-usable parts of the project
- these packages can be directly imported in other projects without being dependent on logic specific to `droplets` project.
- some of the packages included:
    - `logger` - provides logging functions.
    - `graceful` - provides a server wrapper with graceful shutdown enabled.
    - `middlewares` - provides generic middlewares for use in REST or HTTP handlers

#### 3. `web/`

- contains web assets such as css, images, templates etc.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fspy16%2Fdroplets.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fspy16%2Fdroplets?ref=badge_large)