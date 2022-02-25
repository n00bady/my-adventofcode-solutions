#!/usr/bin/env python3

from collections import defaultdict

f = open('./Day_10_Input.txt', 'r')

# Make a list of the adapters and add the zero and the max+3 also sort them out
adapters = [int(line.strip()) for line in f]
adapters.append(0)
adapters.sort()
adapters.append(adapters[-1]+3)

# Dict to keep track of the count of diffs
diffs = {1: 0, 3: 0}
i = 1
a = 0

# from 1 to the max of the adapters joltage 
while i <= max(adapters):
    a += 1
    # if that number in in the adapters increase the count of 1 or 2 joltage diffs
    if i in adapters:
        diffs[a] += 1
        # don't forget to zero it
        a = 0
    # step
    i += 1

print("The total differences of 1-jolt multiplied by the total differences of 3-jolts is :\n", diffs[1]*diffs[3])
