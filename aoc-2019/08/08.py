#!/usr/bin/env python3

# Advent of Code day 08
# Too many pixelz
# numpy is just too convenient and still don't fully understant it ¯\_(ツ)_/¯
# I bet there is a way easier way to make an actual img out of this but I am too
# lazy to search for it now...

import sys
import numpy as np

# Checking if the image isn't corrupted as per part 1
def check(data):
    # Make an nparray with the layer that has the least 0s and then calculate
    # the sum of the 1s and 2s multiply them and return them.
    least_zer0s = min(data, key=lambda layer:np.sum(layer == '0'))
    return (np.sum(least_zer0s == '1') * np.sum(least_zer0s == '2'))

# decoding the data as per part 2
def decode(data):
    decoded = data[0]
    for layer in data:
        decoded = np.where( decoded != '2', decoded, layer)
    decoded = np.where(decoded == '1', '█', ' ')
    return decoded

# Printing some ascii art :)
def displayimg(decoded):
    for r in decoded:
        print(*r)

# Reading the input file as command line argument and then putting the contets
# in a list wich we later shape in numpy array 6x25
with open(sys.argv[1], 'r') as f:
    for line in f:
        imgdata = list(line.strip())
shaped_data = np.array(imgdata).reshape(-1, 6, 25)

# Part 1 
print('Image check: ', check(shaped_data), end='\n------------------\n\n')

# Part 2
displayimg(decode(shaped_data))

# Something says to me that we gona need this piece of code in a future puzzle.
