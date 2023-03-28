#!/usr/bin/env python3

import sys
import random

length = int(sys.argv[1])
count = int(sys.argv[2])


def generate_random_string(length):
	chars = '`1234567890-=qwertyuiop[]\asdfghjkl;zxcvbnm,./~!@#$%^&*()_+QWERTYUIOP{}|ASDFGHJKL:"ZXCVBNM<>?'
	rand_string = ''.join(random.choice(chars) for i in range(length))
	return rand_string

for i in range(count):
	print(generate_random_string(length))
