#!/usr/bin/env bash

set -eo pipefail

# Store exit value
exit_val=0

# collect aggregate profile
echo "mode: set" > aggregate.out
for dir in $(find pkg -maxdepth 2 -type d ); do
  if ls ${dir}/*.go &> /dev/null; then
    # If the exit value is non-zero, then store it. This wil allow us to test
    # everything before failing
    go test -v -coverprofile=profile.out github.com/miroswan/mesops/${dir} || {
      exit_val=$?
    }
    if [ -f profile.out ]; then
      cat profile.out | grep -v "mode: set" >> aggregate.out
    fi
  fi
done

# Exit with failure if any tests failed
if [ $exit_val -ne 0 ]; then
  exit $exit_val
fi

# post it
$HOME/gopath/bin/goveralls    \
  -coverprofile=aggregate.out \
  -service=travis-ci          \
  -repotoken $COVERALLS_TOKEN
