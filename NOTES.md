# NOTES

## joystick handler mvp

```
Get the current state of an axis control on a joystick.

Returns a 16-bit signed integer representing the current position of the axis. The state is a value ranging from -32768 to 32767.

On most modern joysticks the x-axis is usually represented by axis 0 and the y-axis by axis 1. The value returned by axisPosition is a signed integer (-32768 to 32767) representing the current position of the axis. It may be necessary to impose certain tolerances on these values to account for jitter.

Some joysticks use axes 2 and 3 for extra buttons.

See SDL_JoystickGetAxis for C documentation.
```

## right axis x

* minimum -32768
* maximum 32767
* origin [-129, 128]