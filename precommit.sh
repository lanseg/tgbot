#!/usr/bin/bash 

set -euo pipefail

if command -v gofmt
then
  find -iname '*go' -print -exec gofmt -s -w {} \;
else
  echo "Not doing gofmt (gofmt not found)"
fi

python_files="$( find -iname '*py' )"

if command -v black
then
  black $python_files
else
  echo "Not doing autoformat (black not found)"
fi

if command -v pylint
then 
  pylint --argument-naming-style camelCase \
         --attr-naming-style camelCase \
         --variable-naming-style camelCase \
         --function-naming-style camelCase \
         $python_files
else
  echo "Not doing static analysis (pylint not found)"
fi
