#!/usr/bin/env python3

# Advent of code day 05
# Intcode computor v2.0 !
# It took more time than I would like to admit especially for the 2nd part and the
# code is very spaghetti I have seen faster and more interesting ways to do this
# but this is all me I got very little help from others.

# What I learned here is that if I used some library or implement some proper
# checks on my own I would probably finished this several hours ago o_o .
# Well at least I managed to do it ! Enjoy my spaghet!


# Function to check for what parameter modes to use 
def param_mode(icode, opcode, i):
    opcode = list(str(opcode))
    paramode = (['0'] * (5-len(opcode))) + opcode
# Had to check if I check for param modes for non existant parameters otherwise
# in some cases I got out of index usually in the examples because of how small
# they were
    if i + 1 < len(icode):
        if paramode[2] == '0':
            param1 = icode[i+1]
        else:
            param1 = i+1
    else:
        param1 = 0
    if i + 2 < len(icode):
        if paramode[1] == '0':
            param2 = icode[i+2]
        else:
            param2 = i+2
    else:
        param2 = 0
    if i + 3 < len(icode):
        if paramode[0] == '0':
            param3 = icode[i+3]
        else:
            param3 = i+3
    else:
        param3 = 0

    return param1, param2, param3

# This is the main function all the operation checks and their calculations
# happen here
def IntingComputor(inputprogram):
    with open(inputprogram, 'r') as f:
        i = 0
        # Strip input and split it on commas then turn the list to integers
        for line in f:
            icode = line.strip().split(',')
            icode = list(map(int, icode))
        print('Running program: ', icode)
        while i < len(icode):
            opcode = str(icode[i])
            op = opcode[len(opcode)-1]
# I think it's better to first check for op 99
            # If opearation is HALT then stop and return
            if op == '9':
                return icode[0]

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
                icode[param1] = int(input("Enter input: >> "))
                i += 2
            # Output 
            elif op == '4':
                print('Diagnostic output: ', icode[param1])
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
            # Anything else is an error :/
            else:
                print(icode[i])
                return 'ERROR!'
            
res = IntingComputor('./input05.txt')
#print('Result: ', res)
