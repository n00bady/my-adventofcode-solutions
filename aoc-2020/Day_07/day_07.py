#!/usr/bin/env python3

count = 0
f = open('./Day_07_Input.txt', 'r')

# Making a dictionary from the input with keys the outer bags and values the bags thar are contained.
rulez = {topbag:conbags.strip('\n').split(', ') for line in f for topbag, conbags in [line.split(' bags contain ')]}
# I took more time than I would like to admit to figure out
# how to have only the colors of the bags in values
for r in rulez.keys():
    rulez[r] = [i.split(' ')[1]+' '+i.split(' ')[2] for i in rulez[r]]
# Function to check if there are shiny gold bags in the bag given and if not
# check if any of the bags contain a bag that can contain a shiny gold bag...
# too many bags -_-
def contains_shiny(bag):
    if bag not in rulez:
        return False
    elif 'shiny gold' in rulez[bag]:
        return True
    elif any(contains_shiny(b) for b in rulez[bag]):
        return True
    else:
        return False

# For every bag in the rulez check if it contains a shiny gold bag and increase counter
for bag in rulez:
    if (contains_shiny(bag)):
        count += 1

print(count)
