# GOZILA

This is an attempt to create a bootstraped project for use
with go templ, with lit components.

## Motivation

The motivation behind this project is to create something that run in a go backend
and serves a mostly static templates with some dynamic components.
The advantage of this approach to say something like a go backend + a SSR application,
is that in this setup you do not have to proxy the client requests from your websites server to
an external api. You skip that cost and you can even add a DB directly to your backend.

to run this project you need to have installed the go templ cli tool:

[go templ](https://github.com/a-h/templ)

you need to install rollup globally

```bash
npm i -g rollup
```

## How to use

Make sure you have the go templ cli tool installed.

```bash
make run-both
```
