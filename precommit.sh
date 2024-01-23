#!/usr/bin/bash 

set -euo pipefail

python_files="$( find -iname '*py' )"
black $python_files
pylint --argument-naming-style camelCase \
       --attr-naming-style camelCase \
       --variable-naming-style camelCase \
       --function-naming-style camelCase \
       $python_files