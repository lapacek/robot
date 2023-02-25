# Robot

We build robot system on the lego mindstorms ev3 platform.
Project start with the simplest components and build up to a complete robot.

## Components

Components are the building blocks of the robot system. Each component is a separate program that can be run independently of the other components. The components communicate with each other using the message bus.

name | description
--- | ---
[tracker](cmd/robot/cmd/tracker/README.md) | Robot movement platform control.

## Build the project

### Install dependencies: 

*(Silicon **M1**, **x86_64** Intel macs)*

```bash
# missing sdl2
$ brew install sdl2

# missing pkg-config
$ brew install pkg-config
```

*(Fedora 37 **x86_64** Linux)*

```bash
# missing sdl2
$ sudo dnf install SDL2-devel
```

### Download and compile the project:

```bash
# dowload the project
$ git clone github.com/lapacek/robot.git

# change directory to the project
$ cd robot

# build the project
$ make build

# result is in the ./build directory
```

## Run unit tests

```bash
# execute all tests
$ make test
```

## Run program

You can run the project from the build directory or install it and run from project directory.
Program is called `robot` and take name of system argument

```bash
# run the project after build
$ ./build/robot

# install the project
$ make install

# run the project after install
$ ./robot
```

Program usage example:

```bash
# print help
$ ./robot --help

# list all system components
$ ./robot --list

# run the project with name of system component
$ ./robot {component}
```

## Deployment

Read about a deployment of system components on target platforms [here](doc/Deployment.md).
