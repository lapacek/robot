# robot

We build robot system on the lego mindstorms ev3 platform.
Project start with the simplest components and build up to a complete robot.

## architecture

### components

Components are the building blocks of the robot system. Each component is a separate program that can be run independently of the other components. The components communicate with each other using the message bus.

name | description
--- | ---
[tracker](cmd/robot/cmd/tracker/Readme.md) | Robot movement platform control.

## automation

### build

Install dependencies: *(Silicon **M1**, **x86_64** Intel macs)*

```bash
# missing sdl2
$ brew install sdl2

# missing pkg-config
$ brew install pkg-config
```

Install dependencies: *(Fedora 37 **x86_64** Linux)*

```bash
# missing sdl2
$ sudo dnf install SDL2-devel
```

Build the project:

```bash
# dowload the project
$ git clone github.com/lapacek/robot.git

# change directory to the project
$ cd robot

# build the project
$ make build

# result is in the ./build directory
```

### test

Test the project:

```bash
# execute all tests
$ make test
```

### run

You can run the project from the build directory or install it and run from project directory.
Program is called `robot` and take name of system argument 

Run the project:

```bash
# build the project after clone repo
$ make build

# run the project
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

## deployment

Read about a deployment of system components on target platforms [here](doc/Deployment.md).
