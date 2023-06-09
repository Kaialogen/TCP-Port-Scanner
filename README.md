# TCP-Port-Scanner

Written using Go 1.20.5

## Usage

```shell
$ ./scanner [ip] [tcp/udp]
```

## Building From Source

Since this tool is written in Go you need to install the Go language/compiler/etc. Full details of installation and set up can be found on the Go language website.

### Compiling 

This project incorporates `make` files:

`$ make`

This will make two binaires for both Linux amd64 and Macos arm64. They both need to given permissions before executing:

```shell
$ chmod +x FILE
$ ./FILE
```
