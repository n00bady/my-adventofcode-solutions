#!/usr/bin/env python3

data = open('./Day_02_Input.txt', 'r')

# ---- Part 1 ----
valid_pass_count = 0
count = 0

print("Start part 1!")
# Check input line by line and split each string on each line to extract it's values...
for line in data:
    MinMax, letter, pwd = line.split()
    # Remove the - between the 2 numbers
    minmax_nums = MinMax.split('-')
    # count the occurances of the letter after striping the :
    letter_count = pwd.count(letter.strip(':'))
    # if the letter cound is between min max values then add 1 to the valid password counter
    if letter_count >= int(minmax_nums[0]) and letter_count <= int(minmax_nums[1]):
        valid_pass_count += 1

# Result for part 1
print("The number of valid password for Part 1 are..: ", valid_pass_count)
data.close()

data = open('./Day_02_Input.txt', 'r')
# ---- Part 2 ----
valid_pass_count = 0

print("Start part 2!")
for l in data:
    pos, char, pswd = l.split()
    pos_nums = pos.split('-')

# Probably better way to do this in python...............
    # Check each and every letter of the each password 
    for i in range(len(pswd)):
        # if the passwords letter in pos i matches the given char...
        if pswd[i] == char.strip(':'):
            # if the position i also matches one of the given possitions...
            if i == int(pos_nums[0])-1 or i == int(pos_nums[1])-1:
                # and only one of those...
                if pswd[int(pos_nums[0])-1] != pswd[int(pos_nums[1])-1]:
                    # then increase the valid password counter by one.
                    valid_pass_count += 1

# Result for part 2
print("The number of valid password for Part 2 are..: ", valid_pass_count)
data.close()
