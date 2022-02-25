#!/user/bin/env python3 

# Advent of Code day 06
# I am too late for this one  :(
# These will be helpfull here: https://docs.python.org/2/library/collections.html#defaultdict-examples
#                              https://www.python.org/doc/essays/graphs/

from collections import defaultdict
from collections import deque

def count_all_orbits(start, count, total):
    total += count
    for sat in sol[start]:
        total = count_all_orbits(sat, count+1, total)
    return total

def find_shortest_path(graph, start, end):
    dist = {start: [start]}
    q = deque(graph)
    while len(q):
        at = q.popleft()
        for next in graph[at]:
            print(next)
            if next not in dist:
                dist[next] = [dist[at], next]
                q.append(next)
    return dist.get(end)

def test(g, s, e, path=[]):
    shortest = None
    if s != e:
        for n in g:
            if s in g[n]:
                path = path + [n]
            elif s == e:
                break
    return path

#def find_shortest(graph, start, end, path=[]):
#    path = path + [start]
#    print(path)
#    if start == end:
#        return path
#    if start not in graph.keys():
#        return None
#    shortest = None
#    for node in graph[start]:
#        if node not in path:
#            newpath = find_shortest(graph, node, end, path)
#            if newpath:
#                if not shortest or len(newpath) < len(shortest):
#                    shortest = newpath
#    return shortest

#def find_all_paths(graph, start, end, path=[]):
#    path = path + [start]
#    if start == end:
#        return [path]
#    if start not in graph.keys():
#        return []
#    paths = []
#    for node in graph[start]:
#        if node not in path:
#            newpaths = find_all_paths(graph, node, end, path)
#            for newpath in newpaths:
#                paths.append(newpath)
#    return paths

sol = defaultdict(list)
with open('./input06.txt') as f:
    for line in f.readlines():
        planet, sat = line.strip().split(')')
        sol[planet].append(sat)
        if sat == 'YOU':
            start = planet
        if sat == 'SAN':
            end = planet

print(count_all_orbits('COM', 0 , 0))
print(start, end)
#print(dict(sol).items())
#print(find_shortest_path(dict(sol), start, end))
print(test(dict(sol), start, end))
