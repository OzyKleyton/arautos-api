#!/bin/sh
set -e

echo "Starting api with params $@"
if [ "$1" != "" ]; then
    exec /bin/api $@
else
    exec /bin/api
fi