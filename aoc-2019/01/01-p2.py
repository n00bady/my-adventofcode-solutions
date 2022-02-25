
# Advent of Code day 01 part 2
# fuel req. = mass / 3 round down - 2
# fuel for fuel = fuel / 3 round down - 2 until fuel for fuel / 3 is negative 

f = open("./aoc-input01p1.txt")
truetotal = 0

for line in f:
    fuel = int(line)
    total_fuel = 0
    while fuel > 5:             # 5/3 - 2 < 0 this is the limit where you don't need more fuel for the fuel
        fuel = fuel // 3 - 2
        total_fuel += fuel
    truetotal += total_fuel
print(truetotal)
f.close()
