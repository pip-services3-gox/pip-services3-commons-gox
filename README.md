# <img src="https://uploads-ssl.webflow.com/5ea5d3315186cf5ec60c3ee4/5edf1c94ce4c859f2b188094_logo.svg" alt="Pip.Services Logo" width="200"> <br/> Portable Abstractions and Patterns for Golang

This module is a part of the [Pip.Services](http://pip.services.org) polyglot microservices toolkit.
It provides a set of basic patterns used in microservices or backend services.
Also the module implemenets a reasonably thin abstraction layer over most fundamental functions across
all languages supported by the toolkit to facilitate symmetric implementation.

The module contains the following packages:

- [**Commands**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/commands) - commanding and eventing patterns
- [**Config**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/config) - configuration framework
- [**Convert**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/convert) - soft value converters
- [**Data**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/data) - data patterns
- [**Errors**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/errors) - application errors
- [**Random**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/random) - random data generators
- [**Refer**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/refer) - locator (IoC) pattern
- [**Reflect**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/reflect) - reflection framework
- [**Run**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/run) - execution framework
- [**Validate**](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox/validate) - validation framework

<a name="links"></a> Quick links:

* [Configuration Pattern](https://www.pipservices.org/recipies/configuration) 
* [Locator Pattern](https://www.pipservices.org/recipies/references)
* [Component Lifecycle](https://www.pipservices.org/recipies/component-lifecycle)
* [Components with Active Logic](https://www.pipservices.org/recipies/active-logic)
* [Data Patterns](https://www.pipservices.org/recipies/memory-persistence)
* [API Reference](https://godoc.org/github.com/pip-services3-gox/pip-services3-commons-gox)
* [Change Log](CHANGELOG.md)
* [Get Help](https://www.pipservices.org/community/help)
* [Contribute](https://www.pipservices.org/community/contribute)


## Use

Get the package from the Github repository:
```bash
go get -u github.com/pip-services3-gox/pip-services3-commons-gox@latest
```

## Develop

For development you shall install the following prerequisites:
* Golang v1.12+
* Visual Studio Code or another IDE of your choice
* Docker
* Git

Run automated tests:
```bash
go test -v ./test/...
```

Generate API documentation:
```bash
./docgen.ps1
```

Before committing changes run dockerized test as:
```bash
./test.ps1
./clear.ps1
```

## Contacts

The library is created and maintained by **Sergey Seroukhov**.

The documentation is written by **Danyil Tretiakov** and **Levichev Dmitry**.
