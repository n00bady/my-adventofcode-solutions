#!/usr/bin/env python3

# Advent of Code day 04
# MATH! >w<

from collections import Counter

# This is for part 1 and either I am stupid or there is no way to modify it so
# it solves part 2
def checkpass(password):
    c = 0
    # First we check for consecutive duplicate numbers 
    for d in range(len(password)-1):
        if password[d] == password[d+1]:
            c += 1

    # Then we check for accenting order and if it exists in additional to at
    # least 1 consecutive duplicate number then we return True
    if sorted(password) == password and c > 0:
        return True
    else:
        return False

# I had to resort to reddit for this one :( I could not figure it out all on my
# own for whatever reason
def checkpass_2(password):
    # You can easily change the c == 2 to c >= 2 for part 1 
    if sorted(password) == password and any(c == 2 for c in Counter(password).values()):
        return True
    else:
        return False

low = 245318
high = 765747
count = 0
count2 = 0
for t in range(low, high + 1):
    # Turn each number into a list of digits
    digitlst = [int(d) for d in str(t)]
    if checkpass(digitlst) == True:
        count += 1
    if  checkpass_2(digitlst) == True:
        count2 += 1

print('The number of possible passwords for part 1 is: ', count)
print('The number of possible passwords for part 2 is: ', count2)
