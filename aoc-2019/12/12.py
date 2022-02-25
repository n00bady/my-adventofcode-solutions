#!/usr/bin/env python3

# AoC day 12
# This is no moon !

io = [[-1, 0, 2], [0, 0, 0]]
europa = [[2, -10, -7], [0, 0, 0]]
ganymede = [[4, -8, 8], [0, 0, 0]]
callisto = [[3, 5, -1], [0, 0, 0]]

def gravity():
    for i in range(3):
        if io[0][i] > europa[0][i]:
            io[1][i] -= 1
            europa[1][i] += 1
        elif io[0][i] < europa[0][i]:
            io[1][i] += 1
            europa[1][i] -= 1
        if ganymede[0][i] > callisto[0][i]:
            ganymede[1][i] -= 1
            callisto[1][i] += 1
        elif ganymede[0][i] < callisto[0][i]:
            ganymede[1][i] += 1
            callisto[1][i] -= 1
def applyvelocity():
    for i in range(3):
        io[0][i] = io[0][i] + io[1][i]
        europa[0][i] = europa[0][i] + europa[1][i]
        ganymede[0][i] = ganymede[0][i] + ganymede[1][i]
        callisto[0][i] = callisto[0][i] + callisto[1][i]
def calcenergy(moon):
    pot = sum(map(abs,moon[0]))
    kin = sum(map(abs, moon[1]))
    return pot * kin

total = 0
for t in range(11):
    gravity()
    applyvelocity()
    print('After step ', t, ':')
    print(io)
    print(europa)
    print(ganymede)
    print(callisto)
total = calcenergy(io) + calcenergy(europa) + calcenergy(ganymede) + calcenergy(callisto)
print(total)
