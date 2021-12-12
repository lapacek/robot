# go-ev3-dualshock3

The EV3 robot controlled by Dualshock3.

## motivation

For kids.

## development

### build

```bash
$ git clone github.com/lapacek/go-ev3-dualshock-3
$ cd go-ev3-dualshock-3
```

```bash
$ go mod download
$ go build -x cmd/main.go
```

#### macos 10.15 Catalina

* sdl2 is required

```bash
$ brew install sdl2
```

## manual testing

```bash
make test
```

