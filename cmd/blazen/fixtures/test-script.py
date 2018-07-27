#!/usr/bin/env python3


import logging
import sys
from random import randint

def main():
    i = input()
    print("starting")
    print(i)
    exit_code = randint(0, 1)
    print("exit with: ", exit_code)
    raise SyntaxError()
    sys.exit(exit_code)

if __name__ == '__main__':
    sys.exit(main())