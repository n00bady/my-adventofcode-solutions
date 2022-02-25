#!/usr/bin/env python3

f = open('./Day_08_Input.txt', 'r')
# Create a dictionary with the number of line as key and 
# the operation with it's argument as value
inputfile = f.read().splitlines()
instructions = {i:l for i,l in enumerate(inputfile)}

# Make a funtion that pretends to run the code
def runcode(instructions):
    accumulator = 0
    i = 0
    # List to store all the executed lines
    execlines = []
    while True:
        # To make it stop before the the same instruction executes a 2nd time
        # we check it first
        if i in execlines:
            print('insruction executed a 2nd time returning acc')
            return accumulator
        # Then we append the line of the instruction in the executed lines list
        execlines.append(i)
        # And continue according with what every operation does
        if i not in instructions.keys():
            print('i not in keys returning acc')
            return accumulator
        if 'nop' in instructions[i]:
            i += 1
        elif 'acc' in instructions[i]:
            accumulator += int(instructions[i].split(' ')[1])
            i += 1
        elif 'jmp' in instructions[i]:
            i += int(instructions[i].split(' ')[1])
            # We do not need an i += 1 here since we directly modify the contents of i
# I be this would be easier in C 
print(runcode(instructions))
