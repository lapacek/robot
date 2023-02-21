## deployment

- [ev3 brick](#ev3-brick)

### ev3 brick

Prepare your brick for deployment with `ev3dev` project:

* [www,ev3dev.org - prepare brick](https://www.ev3dev.org/docs/getting-started/)
* [www.ev3dev.org - connect brick](https://www.ev3dev.org/docs/tutorials/connecting-to-ev3dev-with-ssh/)

Copy program to the brick:

```bash
# copy program to the brick
$ scp ./build/robot robot@ev3dev.local:/home/robot
```

Run program on the brick:

```bash
# go to the home directory on your ev3 brick
$ cd /home/robot

# run program
$ ./robot tracker
```