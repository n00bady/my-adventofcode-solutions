#!/usr/bin/env python3

f = open('./test_input.txt', 'r')

arr = [c.strip() for c in f.readlines()] 
print(arr)
print(len(arr[0]))
for i in range(len(arr)):
    for j in range(len(arr[i])):
        if arr[i][j] == 'L':
            arr[i][j].replace('L', '#')
print()
print(arr)

