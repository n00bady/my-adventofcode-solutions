#!/usr/bin/env python3

f = open('./Day_06_Input.txt', 'r')
# Split the input on the two newlines for each group and then remove 
# the newlines in each group and put them all in a list
groups = [group.replace('\n', '') for group in f.read().split('\n\n')]
# Then for every group on the list create a set of the characters that
# are in it and get the length of that set and sum up all the lengths of the sets
res = sum([len(set(group)) for group in groups])
print("The sum of the count of all positive answered questions is..: ", res)
f.close()

print("------------------"*4)

# For whatever reason I cannot use the previous f so re-initialize it ?
f = open('./Day_06_Input.txt', 'r')
# And I also have to make the same things again only this time I need the newlines
g2 = [gr for gr in f.read().split('\n\n')]
# Then use set.intersection on a set that is created for each group and print the result
res2 = sum([len(set.intersection(*[set(l) for l in g.split()])) for g in g2])
print("The sum of the count of the questions that everyone answered positively is..: ", res2)
f.close()
