#!/bin/bash

set -ue

SCRIPT_DIR="$( cd "$( dirname "$0" )" && pwd )"

: "$YOUR_NAME"

printf "Name: %s\n" "$YOUR_NAME"
printf "Script Dir: %s\n" "$SCRIPT_DIR"