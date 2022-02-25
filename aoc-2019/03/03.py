#!/usr/bin/env python3

# Advent of Code day 03
# I have no idea what I am doing... Well it was easier than I though maybe I
# should stop doing these late at night and should read the problem more
# focused :|

# Creating a list of all the wires, everyline is a different wire
with open('./input03p1.txt') as f:
    p = f.read().splitlines()

# Create a list of each wire's directions striping the whitespace and spliting
# in commas. We have 2 wires in this part and I am lazy so no N wires :)
wire1 = p[0].strip().split(',')
wire2 = p[1].strip().split(',')

# The function that calculates the path of the wire returns a dict with x,y
# coordinates as keys and number of steps as values
def wirepath(path):
    x, y = 0, 0
    wire = {}
    c = 0
    for i in path:
        dr = i[0]
        # We know that every 1st char in each step is the direction and the rest
        # of the chars are the distance it moves
        dist = int(i[1:]) 
        for m in range(dist):
            c += 1
            if dr == 'R':
                x += 1
            elif dr == 'L':
                x -= 1
            elif dr == 'U':
                y += 1
            elif dr == 'D':
                y -= 1
            wire[x, y] = c
    return wire

# Manhattan distance/taxicab metric/city block distance/L1 distance etc...
# calculation 
def taxicab_distance(intsec):
    distance = sum(map(abs, intsec))
    return distance

path1 = wirepath(wire1)
path2 = wirepath(wire2)

# We find the intersections by making the paths into sets and using the
# intesection operator &
intersections = set(path1) & set(path2)

# Dictionary that has the distances of every intersection as keys and the sum of
# steps it took each wire to get there as values
distances = { taxicab_distance(i): path1[i] + path2[i] for i in intersections }
print('The min distance form origin is: ', min(distances.keys()))
print('The min combined wire lenght is: ', min(distances.values()))
