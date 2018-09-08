#! /usr/bin/python

import argparse
from random import randrange

DEAD = '0'
ALIVE = '1'


def main():
    parser = argparse.ArgumentParser(description="Generate lines from elementary cellular automaton.")
    parser.add_argument("-r", "--rule_number",
                        action="store", nargs="?", default=110, type=int,
                        help="Rule number as integer", metavar="N"
                        )
    parser.add_argument("-i", "--iterations",
                        action="store", nargs="?", default=20, type=int,
                        help="Number of iterations", metavar="N"
                        )
    parser.add_argument("-c", "--cells",
                        action="store", nargs="?", default=10, type=int,
                        help="Number of cells", metavar="N"
                        )
    args = parser.parse_args()
    print(args)
    rule_number = args.rule_number
    iterations = args.iterations
    cells = args.cells
    current_generation = ""
    rule_set = build_ruleset(rule_number)
    print(rule_set)
    # Generate first cell line */
    for cell in range(0, cells):
        current_generation += str(randrange(2))
    print(current_generation)
    for generation in range(1, cells):
        current_generation = update(current_generation, rule_set)
        print(current_generation)


def string_to_int(string):
    return int(string, 2)


def build_ruleset(rule_number):
    ruleset = {}
    binary_rule = "{0:0>8b}".format(rule_number)
    position = 7
    for status in binary_rule:
        ruleset["{0:0>3b}".format(position)] = status
        position -= 1
    return ruleset


def update(cell_line, ruleset):
    new_line = ""
    for position in range(0, len(cell_line)):
        left = position - 1
        right = position + 1
        if right >= len(cell_line):
            right = 0
        pattern = cell_line[left] + cell_line[position] + cell_line[right]
        new_line += ruleset[pattern]
    return new_line

if __name__ == "__main__":
    main()


