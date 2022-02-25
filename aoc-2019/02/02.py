#!/usr/bin/env python3

# Advent of code day 02 part 1 and part 2
# Intcode computor !

# Part 1
def part1(x, y):
    with open('./input02.txt', 'r') as f:
        done = False
        i = 0
        # Strip input and split it on commas then turn the list to integers
        for line in f:
            icode = line.strip().split(',')
            icode = list(map(int, icode))
        icode[1] = x
        icode[2] = y
        while not done:
            if icode[i] == 1:       # opcode 1 is for addition 
                icode[icode[i+3]] = icode[icode[i+1]] + icode[icode[i+2]]
                i += 4              # moving our instruction pointer
            elif icode[i] == 2:     # similar with the addition we execute the multyplication
                icode[icode[i+3]] = icode[icode[i+1]] * icode[icode[i+2]]
                i += 4              # again need to move our instruction pointer to the next instruction
            elif icode[i] == 99:    # 99 opcode means the program's end
                done = True
                return icode[0]

# Part 2 
def part2(output):
    # This is for the noun 0-99
    for n in range(100):
        # This is for the verb 0-99
        for v in range(100):
            if part1(n, v) == output:
                return (n, v)

print(part1(12, 2))  
print(part2(19690720))
ans2 = 100 * part2(19690720)[0] + part2(19690720)[1]
print(ans2)
