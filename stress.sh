#!/bin/bash

URL="http://localhost:8080/api/test"
TOTAL=15

for i in $(seq 1 $TOTAL); do
  curl -s -o /dev/null -w "req $i -> %{http_code}\n" "$URL" &
done

wait
echo "Stress test finished"
