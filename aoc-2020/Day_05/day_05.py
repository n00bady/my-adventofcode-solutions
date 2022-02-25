#!/usr/bin/env/ python3

f = open('./Day_05_Input.txt', 'r')

# ------ Part 1 ------
# For every line replace the F, R with 1 and B, L with 0 then make them
# integer numbers in binary and put them in a list
seats = [int(line.translate(line.maketrans('FBLR', '0101')), 2) for line in f]
# Find the max element in the list and print it
m = max(seats)
print('The highest seat ID is..: ', m)

# Sort the seat ids
sorted_seats = sorted(seats)
# For every seat id in seats list for the first seat id to the last
# if the id doesn't exists in the seat id list put it in myseat variable
myseat = [s for s in range(seats[0], seats[-1]) if s not in seats]
print(myseat)
f.close()
