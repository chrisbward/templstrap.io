

# TemplStrap.io - PROJECT IS IN ALPHA, DO NOT USE FOR PRODUCTION

A component framework for Templ that aligns to the Twitter Bootstrap design system.

The project is currently in iterative testing with real applications.

## Goals

- To have parity with bootstrap design system
- htmx support
- Tailwind support
- Supporting website with component examples w/ code

## Todo

PROJECT IS NOT FINISHED, DO NOT USE FOR PRODUCTION

Code in the repo is liable to change, some components are incomplete 
and will need further testing

components to finish;
- collapse
- listgroup
- modal 
- placeholders
- popovers
- toasts
- tooltips 

## Quickstart

### Templ Installation
Install `v0.3.1001` runtime with the following;
```bash
go install github.com/a-h/templ/cmd/templ@v0.3.1001
```

### Using TemplStrap.io

On your project, include the CLI template engine;
```bash
go get -tool github.com/a-h/templ/cmd/templ@v0.3.1001
```

Add the following to your Makefile;
```make
templ:
    @go tool templ generate
```

If you wish to generate templates, execute `make templ` in your project folder.
