#!/bin/bash

SPECS_LOCATION=$(mktemp -d specs.XXXX)

if test -z "$1"; then
    SPECS_REF="master"
else
    SPECS_REF="$1"
fi

echo "travis_fold:start:generate_specs@$SPECS_REF"

set -ev

git clone --depth 1 --branch $SPECS_REF -- https://github.com/Azure/azure-rest-api-specs.git $SPECS_LOCATION

for f in $(find $SPECS_LOCATION -iname README.md | grep -v node_modules); do 
    autorest --go $f
done

rm -rf $SPECS_LOCATION

echo "travis_fold:end:generate_specs@$SPECS_REF"
