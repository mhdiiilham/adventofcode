#!/bin/sh

# THIS IS A VERY SIMPLE SHELL SCRIPT HELPER.
# THIS SCRIPT ONLY USE TO OPEN THE LATEST ADVENT OF CODE CHALLENGE.
# GOOD LUCK!!!
# PS: I only tried this on MacOS

YEAR=$(TZ="EST" date +%Y)
DATE=$(TZ="EST" date +%d)
DAY=${DATE#"0"}
URL="https://adventofcode.com/$YEAR/day/$DAY"

open $URL
