#!/bin/bash

set -e

# Declare some functionality.
function usage
{
    echo "Usage: $0 [-d <proto dir> -p <protoc binary>]" 1>&2
    exit 1
}

# Pre-flight.
PROTOC="$(which protoc)"
PROTO_DIR="."

while getopts ":d:p:" o; do
    case "${o}" in
        d)
            PROTO_DIR="$OPTARG"
            ;;
        p)
            PROTOC="$OPTARG"
            ;;
        *)
            usage
            ;;
    esac
done

if [[ -z "${PROTOC// }" ]]; then
    echo "protoc not not found in path"
    exit
fi

### Pre-flight complete, let's do stuff
"${PROTOC}" -I "${PROTO_DIR}/." "${PROTO_DIR}/"*.proto --gogo_out=plugins=grpc:"${PROTO_DIR}/." --proto_path=../../../:.
