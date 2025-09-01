#!/bin/bash
if [ -z "$HERO_ID" ]; then
  echo "Error: HERO_ID is not set."
  exit 1
fi
curl -s https://platform.zone01.gr/assets/superhero/all.json \
  | jq -r ".[] | select(.id == $HERO_ID) | .relatives"
