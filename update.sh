#! /bin/sh

# This file is just a shorthand for the command below

make update && make build && pkill logbot && ./bin/logbot
