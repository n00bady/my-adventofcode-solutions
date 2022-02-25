#!/usr/bin/env python3 

import itertools
import intcomputor

n_amps = 5
lstoutputs = []

for settings in itertools.permutations(range(5), n_amps):
    lastoutput = 0
    print('Trying with settings: ', settings)
    for amp in range(n_amps):
        lastoutput = intcomputor.main(settings[amp], lastoutput)
        lstoutputs.append(str(lastoutput))
    print('Output is: ', lastoutput)
print('The max output to thrusters is: ', max(list(map(int, lstoutputs))))
