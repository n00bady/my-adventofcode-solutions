#!/usr/bin/env python3

f = open("./Day_01_Input.txt", "r")
lst=[int(line) for line in f.read().split()]

print("-----Part 1------")
for line in lst:
    for l in lst:
        if line+l == 2020:
            print("!!!!!!!!!!!!!!!")
            print("FOUND IT!!!")
            print(line, "+", l, "= 2020")
            print(line, "*", l, "=", line*l)
            print("!!!!!!!!!!!!!!!")
            break

print("-----Part 2------")
for line in lst:
    for l in lst:
        for i in lst:
            if line+l+i == 2020:
                print("**********")
                print("FOUNT IT!!!")
                print(line, "+", l, "+", i, "= 2020")
                print(line, "*", l, "*", i, "= ", line*l*i)
                print("**********")
                break
                

f.close()
