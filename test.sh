#!/bin/bash
set -e
clear

verbose=""

echo $@

if [[ "$@" == "-v" ]]
then
  verbose="-v"
fi

function cleanup {
    echo "Cleanup resources..."
    docker-compose down
    find ./sql_scripts/sqlite -name *.sqlite* -delete || true
}
# defer cleanup, so it will be executed even after premature exit
trap cleanup EXIT

docker-compose up -d
sleep 10 # Ensure mysql is online

go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest

function test {
  echo "!!! Testing $1"
  export SODA_DIALECT=$1
  soda drop -e $SODA_DIALECT
  soda create -e $SODA_DIALECT
  soda migrate -e $SODA_DIALECT -p ./testdata/migrations
  go test -tags sqlite -count=1 $verbose $(go list ./... | grep -v /vendor/)
  echo "!!! Resetting $1"
  soda drop -e $SODA_DIALECT
  soda create -e $SODA_DIALECT
  echo "!!! Running e2e tests $1"
  cd testdata/e2e; go test -tags sqlite,e2e -count=1 $verbose ./...; cd ..
}

test "sqlite"
test "postgres"
test "cockroach"
test "mysql"

# Does not appear to be implemented in pop:
# test "sqlserver"
