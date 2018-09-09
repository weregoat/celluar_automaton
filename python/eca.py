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
    parser.add_argument("-l", "--live",
                        action="store", nargs="?", default=ALIVE, type=str,
                        help="Character to use for live cells", metavar="C"
                        )
    parser.add_argument("-d", "--dead",
                        action="store", nargs="?", default=DEAD, type=str,
                        help="Character to use for dead cells", metavar="C"
                        )
    parser.add_argument("-f", "--first",
                        action="store", nargs="?", type=str,
                        help="First generation"
                        )
    args = parser.parse_args()
    print(args)
    args_check(args)
    rule_number = args.rule_number
    iterations = args.iterations
    cells = args.cells
    current_generation = args.first
    current_generation = current_generation.replace(args.dead, DEAD)
    current_generation = current_generation.replace(args.live, ALIVE)
    if len(current_generation) < cells:
        current_generation = current_generation.rjust(cells, DEAD)
    rule_set = build_ruleset(rule_number)
    for generation in range(0, iterations):
        print_line(current_generation, args.dead, args.live)
        current_generation = update(current_generation, rule_set)


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


def args_check(args):
    good = True

    if args.dead is None:
        args.dead = DEAD
    if args.live is None:
        args.live = ALIVE
    if args.first is None:
        args.first = generate(args.dead, args.live, args.cells)

    # Check rule between 0 and 255
    if args.rule_number > 255 or args.rule_number < 0:
        good = False
        print("Rule number should be between 0 and 255")
    # Check first generation is consistent with symbols and number of cells
    if len(args.first) > args.cells or len(args.first) == 0:
        good = False
        print("Invalid number of cells for first generation")
    # Check symbols of status have len 1
    if len(args.dead) != 1 or len(args.live) != 1:
        good = False
        print("Symbol for status should be one character long")
    # Check symbols are not the one used already for the opposite status
    if args.dead == ALIVE or args.live == DEAD:
        good = False
        print("Invalid symbol for status")
    # Check first generation is consistent with symbols
    # By removing the dead and alive symbol and making sure the
    # resulting string is empty
    line = args.first
    no_dead = line.replace(args.dead, "")
    no_live = no_dead.replace(args.live, "")
    if len(no_live) != 0:
        good = False
        print("Invalid symbol in first generation")
    if not good:
        exit(99)


def generate(dead, alive, cells):
    first_generation = ""
    # Generate first cell line */
    for cell in range(0, cells):
        status = dead
        if randrange(2) == 1:
            status = alive
        first_generation += status
    return first_generation


def print_line(line, dead, alive):
    line = line.replace(DEAD, dead)
    line = line.replace(ALIVE, alive)
    print(line)


if __name__ == "__main__":
    main()


