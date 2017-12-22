#!/usr/bin/env bash

set -eo pipefail

help(){
  cat <<EOF
This script will help you clean the proto files. First, fix the paths for the
includes by starting from the GOPATH all the way up to the file. Next, call
this script by passing the name of the proto file (without the extension) to the
-p flag.
EOF
}

parse_args(){
  while getopts p:h flag $@; do
    case $flag in
      ( p ) PROTO=$OPTARG  ;;
      ( h ) help && exit   ;;
      ( * ) help && exit 1 ;;
    esac
  done
}

main(){
  parse_args $@
  [[ "x${PROTO}x" != "xx" ]] || {
    help && exit 1
  }
  local proto_file=github.com/miroswan/mesops/pkg/v1/$PROTO/$PROTO.proto
  pushd $GOPATH/src > /dev/null 2>&1
    declare source_no_ext
    while read -r t; do
      while read -r p; do
        local source_base="$(basename $(echo "$p" | awk '{ print $1 }'))"
        source_no_ext=${source_base%.*}
        break
      done < <(find github.com/miroswan/mesops/pkg/v1 -name "*proto" -exec grep -in "message $t" {} +)
      local tmp=$(mktemp)
      sed "s/ $t / $source_no_ext.$t /g" "$proto_file" > $tmp
      mv $tmp $proto_file
    done < <(
      pushd $GOPATH/src > /dev/null 2>&1 &&
        protoc -I . --go_out=. github.com/miroswan/mesops/pkg/v1/$PROTO/$PROTO.proto 2>&1 |
          grep defined                                                                    |
          awk '{ print $2 }'                                                              |
          tr -d '"' &&
      popd > /dev/null 2>&1
    )
  popd > /dev/null
  echo "Processing on $PROTO.proto is complete"
}

if [[ $BASH_SOURCE == $0 ]]; then
  main $@
fi
