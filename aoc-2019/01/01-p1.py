
# Advent of Code 2019 day 01 part 1
# mass / 3 round down - 2 = amount of fuel required

f = open("./aoc-input01p1.txt")
total_fuel = 0

for line in f:
    fuel = int(line) // 3 - 2
    total_fuel += fuel
print(total_fuel)
f.close()
