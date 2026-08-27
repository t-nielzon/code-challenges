package main

/*
Kata: Ball Upwards (6 kyu)

You throw a ball vertically upwards with an initial speed v0 (in km per hour).

The height h of the ball at each time t is given by:
    h = v0 * t - 0.5 * g * t * t
where g is Earth's gravity (g ≈ 9.81 m/s²).

A device is recording at every tenth of second the height of the ball.

For example, with v0 = 15 km/h:
    (0, 0.0), (1, 0.367...), (2, 0.637...), (3, 0.808...), (4, 0.881...) ...
where the first number is time in tenths of a second and the second is height in meters.

Task:
Write a function with parameter v0 (in km per hour) that returns the time
in tenths of a second of the maximum height recorded by the device.

Examples:
    Given v0 = 15 --> should return 4
    Given v0 = 25 --> should return 7

Notes:
    - Convert velocity from km/h to m/s (1 km/h = 5/18 m/s)
    - The maximum height recorded by the device is not necessarily the maximum
      height reached by the ball
*/

func BallUpwards(v0 int) int {
	return 0
}