#!/usr/bin/env python3

import re

f = open('./Day_04_Input.txt', 'r')
# I excluded cid since it's optional
req = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
valid1 = 0 # Counter for part 1
valid2 = 0 # Counter for part 2

# ------ Part 1 ------
# Split input on the two newlines
for line in f.read().split('\n\n'):
    # Replacing the single newlines make it more convinient
    line = line.replace('\n', ' ')
    # If all the requirements are met then increase the valid counter
    # all() and any() are very convinient if you know how to use them
    if all(r+':' in line for r in req):
        valid1 += 1
        # -------- Part 2 --------
        # Make a dictionary with the requirements as keys and the values given in each passport as values
        pport = {k:v for field in line.split(' ') for k, v in [field.split(':')]}
        # If the passport fits all these rules increase the part 2 valid counter
        if (
                int(pport['byr']) >= 1920 and int(pport['byr']) <= 2002 and
                int(pport['iyr']) >= 2010 and int(pport['iyr']) <= 2020 and
                int(pport['eyr']) >= 2020 and int(pport['eyr']) <= 2030 and
                re.match("^(1([5-9][0-9])cm|(59|[6][0-9]|[7][0-6])in)$", pport['hgt']) and
                re.match("#[0-9a-f]{6}", pport['hcl']) and
                re.match("^(amb|blu|brn|gry|grn|hzl|oth)$", pport['ecl']) and
                re.match("^\d{9}$", pport['pid'])
            ): valid2 += 1
                
print("------ Part 1 ------")
print("Total Valid Passports..: ", valid1)
print("\n------ Part 2 ------")
print("Total Valid Passports for part 2..: ", valid2)
f.close()
