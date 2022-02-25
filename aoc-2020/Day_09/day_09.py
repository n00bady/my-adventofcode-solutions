#!/usr/bin/env python3

import sys

f = open('./Day_09_Input.txt', 'r')
nums = [int(line.strip()) for line in f]
i = 25

while i < len(nums):
    # check_current_number
    checkrange = nums[i-25: i]
    preamble = [j + o for j in checkrange for o in checkrange[1:]]
    if nums[i] not in preamble:
        print("First number encountered that's not in the preamble..: ", nums[i])
        sys.exit()
    i += 1

