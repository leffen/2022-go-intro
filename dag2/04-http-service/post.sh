#!/bin/bash

URL=http://localhost:8080/measurement

curl  -H 'Content-type: application/json' \
      --data @testdata/data.json \
      -v ${URL}
