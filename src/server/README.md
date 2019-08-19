# Server for LAR Project
[![pipeline status](https://gitlab.com/lar-project/server/badges/master/pipeline.svg)](https://gitlab.com/lar-project/server/commits/master)

![](https://gitlab.com/lar-project/lar-project/uploads/2658b3e37c3e396e2df56a2163c3ab22/Untitled_Diagram.png)

## Before Buid

**Note: Go >= 1.9.0 is required.**

There are some pakages required for build:
- go
- git
- make
- glibc (or musl-dev)
- dep
- upx (for production only)

> `dep` is official Go dependency management tool. It can be installed via `go get -u -v github.com/golang/dep/cmd/dep`


## Build

- Install dependency packages:

``` bash
make requires
```

- Compile to executable file:

```bash
make
```
> Executable file will be put in `build` directory.

- Build for production:

```bash
make production
```
> It will utilze UPX to compress and pack executable file. If UPX is not installed in system, it will simply copy the build result.

#to build in current development environment
#enter in bash terminal:

#    go build -o larserver.exe -ldflags "-s -w -X main.mode=debug"