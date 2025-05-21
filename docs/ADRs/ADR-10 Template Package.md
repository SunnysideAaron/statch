# ADR-10 Template Package

## Status

Accepted

## Context



## Decision

[plush](https://github.com/gobuffalo/plush)

## Why

plush templates are not compiled into the binary. Instead interpreted at runtime.
Since we want users of statch to be able to write their own templates this would
be a good thing.

## Notes

- [gomplate](https://github.com/hairyhenderson/gomplate)
  - seems to have a similar idea of what statch may someday be like. Doesn't
    have sql sources. Con: uses "." context.

## Consequences



## Other Possible Options

- [liquid](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates.
- [plush](https://github.com/gobuffalo/plush)
  - It appears that plush templates are not compiled into the binary. Instead
    interpreted at runtime. Since we want users of statch to be able to write their
    own templates this would be a good thing.
- [templ](https://github.com/a-h/templ) - A HTML templating language that has great developer tooling.
  - Of the choices left templ has the largest user base.
- [Awesome go's list of template engines](https://github.com/avelino/awesome-go?tab=readme-ov-file#template-engines)

## Not an Option

- [ego](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.
  - last release 2021
- [extemplate](https://git.sr.ht/~dvko/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance.
- [fasttemplate](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/).
- [gomponents](https://www.gomponents.com) - HTML 5 components in pure Go, that look something like this: `func(name string) g.Node { return Div(Class("headline"), g.Textf("Hi %v!", name)) }`.
- [got](https://github.com/goradd/got) - A Go code generator inspired by Hero and Fasttemplate. Has include files, custom tag definitions, injected Go code, language translation, and more.
- [goview](https://github.com/foolin/goview) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.
- [hero](https://github.com/shiyanhui/hero)
- [htmgo](https://htmgo.dev) - build simple and scalable systems with go + htmx
- [jet](https://github.com/CloudyKit/jet) - Jet template engine.
  - only sligtly better "." notation than standard template.
- [maroto](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.
- [pongo2](https://github.com/flosch/pongo2) - Django-like template-engine for Go.
  - last release: 2022-06
- [quicktemplate](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.
- [raymond](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go.
  - last release: 2018
- [Razor](https://github.com/sipin/gorazor) - Razor view engine for Golang.
  - last release: 2019
- [Soy](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).
- [sprout](https://github.com/go-sprout/sprout) - Useful template functions for Go templates.
- [tbd](https://github.com/lucasepe/tbd) - A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata.
  - very simple. last release: 2021
- [template](https://pkg.go.dev/text/template).
  - I personally find template gross. The code is ugly and difficult to read.
    "." notation is awful. Hiding where data is coming from. Others may like it
    but I had an immediate visceral reaction against it.