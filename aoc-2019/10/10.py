#!/usr/bin/env python3

# AoC day 10
# So many asteroids @_@

import sys
import math

def countinlos(station, asteroids, size):
    detected = set()
    for asteroid in asteroids:
        if asteroid != station:
            dx, dy = asteroid[0] - station[0], asteroid[1] - station[1]
            g = abs(math.gcd(dx, dy))
            reduced = (dx//g, dy//g)
            detected.add(reduced)
    return detected

with open(sys.argv[1]) as f:
    content = f.readlines()
print(content)
rawasteroids = []
rawasteroids = [s.strip() for s in content]
size =(len(rawasteroids), len(rawasteroids[0]))
print('\n')
print(rawasteroids)
print('\n')
print(size)

asteroids = set()
for x in range(size[0]):
    for y in range(size[1]):
        if rawasteroids[x][y] == '#':
            asteroids.add((x,y))
print('\n')
print(asteroids)

stationCounts =[]
for station in asteroids:
    inlos = countinlos(station, asteroids, size)
    stationCounts.append((len(inlos), station, inlos))
    stationCounts.sort(reverse=True)
    amtinlos, station, inlos = stationCounts[0]
print(amtinlos)
