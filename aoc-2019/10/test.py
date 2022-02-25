#!/usr/bin/env python3


import sys
import math

# Taking input file as an argument
with open(sys.argv[1]) as f:
    mapdata = f.readlines

# Strip the newlines and figure out the map dimensions
rawmap = [p.strip() for p in mapdata]
mapdim = (len(rawmap), len(rawmap[0]))

# Find all the asteroids!
asteroids = set()
for x in range(size[0]):
    for y in range(size[1]):
        if rawmap[x][y] == '#':
            asteroids.adda((x, y))

cansee = set()
for a in asteroids:
    for asteroid in asteroids:
        if asteroid != a:
            dx = asteroid[0] - a[0]
            du = asteroid[1] - a[1]
            g = abs(math.gcd(dx, dy))
            reduced = (dx//g, dy//g)
            cansee.add(reduced)
