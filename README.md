# robot

We build robot system on the lego mindstorms ev3 platform.
Project start with the simplest components and build up to a complete robot. 

## system

```bash
```

## arch

### design decisions

```bash
```

## test

```bash
# execute all tests
$ make test
```

## build

### MacOS 13.01 Ventura

```bash
# dowload the project
$ git clone github.com/lapacek/robot.git

# change directory to the project
$ cd robot

# build the project
$ make build

# result is in the ./build directory
```

## run

#### MacOS 13.01 Ventura

Works on Silicon **M1**, **x86_64** Intel macs.

```bash
# missing sdl2
$ brew install sdl2

# missing pkg-config
$ brew install pkg-config
```

#### Fedora 37 Linux

Works on **x86_64** Intel machines.

```bash
# missing SDL2-devel
$ sudo dnf install SDL2-devel
```

#### Lego Mindstorms EV3

```bash
```