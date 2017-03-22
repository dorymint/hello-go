#!/bin/sh
set -eu
cd "$(dirname "$(readlink -f "$0")")"
mkfifo ./shellpipe
echo "hello pipe" > ./shellpipe &
sleep 3
cat ./shellpipe
wait
rm ./shellpipe
# EOF
