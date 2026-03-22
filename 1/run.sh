#!/bin/bash

set -e

echo "Kompilacja i uruchamianie aplikacji przez Docker..."

docker run --rm \
  -v "$(pwd)":/app \
  -w /app \
  frolvlad/alpine-fpc \
  sh -c "fpc BubbleSortApp.pas && ./BubbleSortApp"

echo "Zakończono."