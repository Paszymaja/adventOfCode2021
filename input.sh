#!/bin/bash
for i in {1..10}
do
    mkdir day$i
    curl -b $(cat .session) https://adventofcode.com/2021/day/$i/input > day$i/input.txt
done
