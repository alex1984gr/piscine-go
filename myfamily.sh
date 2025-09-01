#!/bin/bash
if [ -z "$HERI_ID" ]; then
echo "error : HERO_ID is not set."
if curl -s https://platform.zone01.gr/assets/superhero/all.json \
| jq -r ".[] | select(.id == $HERO_ID) | .relatives"