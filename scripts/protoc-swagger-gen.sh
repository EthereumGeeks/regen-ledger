#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen
proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do

  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 -name 'query.proto')
  if [[ ! -z "$query_file" ]]; then
    protoc  \
    -I "proto" \
    -I "third_party/proto" \
    "$query_file" \
    --swagger_out ./tmp-swagger-gen \
    --swagger_opt logtostderr=true --swagger_opt fqn_for_swagger_name=true --swagger_opt simple_operation_ids=true
  fi
done


# download Cosmos SDK swagger doc
rm -f ./client/docs/swagger.yaml
wget https://raw.githubusercontent.com/cosmos/cosmos-sdk/v0.40.0/client/docs/swagger-ui/swagger.yaml -P ./client/docs

# combine swagger files
# uses nodejs package `swagger-combine`.
# all the individual swagger files need to be configured in `config.json` for merging
swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# clean swagger files
rm -rf ./tmp-swagger-gen