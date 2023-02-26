# Deployment

- [ev3 brick](#ev3-brick)

## EV3 brick

The EV3 brick is powered by the ARM926EJ-S 32-bit (ARMv5Tej) processor.

* [https://en.wikipedia.org/wiki/ARM9](https://en.wikipedia.org/wiki/ARM9)

### Prepare your brick for deployment with `ev3dev` project:

* [www,ev3dev.org - prepare brick](https://www.ev3dev.org/docs/getting-started/)
* [www.ev3dev.org - connect brick](https://www.ev3dev.org/docs/tutorials/connecting-to-ev3dev-with-ssh/)
* [www.ev3dev.org - connect internet](https://www.ev3dev.org/docs/tutorials/connecting-to-the-internet-via-usb/)
* [www.ev3dev.org - connect Wi-Fi](https://www.ev3dev.org/docs/tutorials/setting-up-wifi-using-the-command-line/)

### Build project for the EV3 brick:

```bash
# cross compile project for the EV3 brick
$ make build-ev3

# result is in the ./build directory
```

### Install program to the brick:

```bash
# set executable permission
$ chmod +x ./build/robot

# copy program to the brick after build
$ scp ./build/robot robot@ev3dev.local:/home/robot

# set executable permission
$ chmod +x ./scripts/run_tracker.sh

# copy starting script to the brick
$ scp ./scripts/run_tracker.sh robot@ev3dev.local:/home/robot
```

### Run program on the brick via ssh:

```bash
# connect to the brick
$ ssh robot@ev3dev.local

# change directory to the home directory
$ cd /home/robot

# run program
$ ./robot tracker

# or run starting script
$ ./run_tracker.sh
```
