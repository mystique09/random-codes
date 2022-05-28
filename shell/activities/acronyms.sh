#!/bin/bash

WORDS=""

read -p "Enter words: " WORDS

count=0

ACR=""

for WORD in ${WORDS}; do
  W=${WORD:0:1}
  if [[ $count -eq 5 ]]; then
    break
  fi

  ACR="${ACR}${W^}"
  count=$((count+1))
done

echo ${ACR}
