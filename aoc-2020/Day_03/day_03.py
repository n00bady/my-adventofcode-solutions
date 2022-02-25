#!/usr/bin/env python3

import sys

f = open('./Day_03_Input.txt', 'r')
readinput=f.read()
trees_count = 0
down = int(sys.argv[1])
right = int(sys.argv[2])

# Put all the text in a list of lists :P  and strip newlines
data = []
for line in readinput.splitlines():
    data.append(list(line.strip()))
f.close()

# Since every line has the same amount of columns it's easier :)
print("-----------------")
print("Our data have")
print("Rows..: ", len(data), "\nCols..: ", len(data[0]))
print("With Row step = ", down, "and Coloumn step = ", right)
print("-----------------")

# Now check all the rows with step one and all the columns with step 3 if they have # trees
j = 0
for i in range(down, len(data), down):
    # wrap around when you get to the end of the column
    if j+right > 30:
        j = j + right - 31
    else:
        j += right
    if data[i][j] == "#":
        trees_count += 1
    
print("Total Trees encountered..: ", trees_count)

