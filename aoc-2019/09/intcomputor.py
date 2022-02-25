#!usr/bin/env python3

# Advent of Code day 09
# Intcode computor again and again... Don't get me wrong I like the intcomputer
# puzzles I would just prefer if there was more spread out.
# Add support for parameters in relative mode
# Computor's memory should be much bigger than the initial program
# and should have support for large numbers

import sys

# 0: positional, 1: imediate, 2: relative
def param_mode(icode, opcode, i):
    opcode = list(str(opcode))
    paramode = (['0'] * (5-len(opcode))) + opcode
# Had to check if I check for param modes for non existant parameters otherwise
# in some cases I got out of index usually in the examples because of how small
# they were
    if i + 1 < len(icode):
        if paramode[2] == '0':
            param1 = icode[i+1]
        elif paramode[2] == '2':
            param1 = icode[i+1] + relative_base
        else:
            param1 = i+1
    else:
        param1 = 0

    if i + 2 < len(icode):
        if paramode[1] == '0':
            param2 = icode[i+2]
        elif paramode[1] == '2':
            param2 = icode[i+2] + relative_base
        else:
            param2 = i+2
    else:
        param2 = 0

    if i + 3 < len(icode):
        if paramode[0] == '0':
            param3 = icode[i+3]
        elif paramode[0] == '2':
            param3 = icode[i+3] + relative_base
        else:
            param3 = i+3
    else:
        param3 = 0

    return param1, param2, param3

# This is the main function all the operation checks and their calculations
# happen here
def IntingComputor(inputprogram):
    global relative_base
    relative_base = 0
    with open(inputprogram, 'r') as f:
        i = 0
        # Strip input and split it on commas then turn the list to integers
        for line in f:
            icode = line.strip().split(',')
            icode = list(map(int, icode))
        for z in range(1000000):
            icode.append(0)
#        print('Running program: ', icode)
        countinputs = 0
        while i < len(icode):
            opcode = str(icode[i])
            op = opcode[len(opcode)-1]
# I think it's better to first check for op 99
            # If opearation is HALT then stop and return
            if int(opcode)%100 == 99:
                return out

            param1, param2, param3 = param_mode(icode, opcode, i)

            # Addition
            if op == '1':
                icode[param3] = icode[param1] + icode[param2]
                i += 4
            # Multiplication
            elif op == '2':
                icode[param3] = icode[param1] * icode[param2]
                i += 4
            # Input from user
            elif op == '3':
                icode[param1] = int(input("Enter input >> "))
                i += 2
            # Output 
            elif op == '4':
                print('[', i, '] Diagnostic output: ', icode[param1])
                out = icode[param1]
                i += 2
            # Jump if true
            elif op == '5':
                if icode[param1] != 0:
                    i = icode[param2]
                else:
                    i += 3
            # Jump if false
            elif op == '6':
                if icode[param1] == 0:
                    i = icode[param2]
                else:
                    i += 3
            # Less than
            elif op == '7':
                if icode[param1] < icode[param2]:
                    icode[param3] = 1
                    i += 4
                else:
                    icode[param3] = 0
                    i += 4
            # Equals
            elif op == '8':
                if icode[param1] == icode[param2]:
                    icode[param3] = 1
                    i += 4
                else:
                    icode[param3] = 0
                    i += 4
            # Operation to change the relative base for use in parameter mode 2
            elif op == '9':
                relative_base += icode[param1]
                i += 2
            # Anything else is an error :/
            else:
                print(icode[i])
                return 'ERROR!'

def main(inputprogram):
    output = IntingComputor(inputprogram)
    return output

inputprogram = sys.argv[1]
if not sys.argv[1]:
    print('You need to give program file as argument')
else:
    output = main(inputprogram)
    print('The final output is: ', output)
