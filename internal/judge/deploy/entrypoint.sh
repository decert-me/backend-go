#!/bin/bash

nohup anvil > /dev/null 2>&1 &

./decert-judge > log.log 2>&1