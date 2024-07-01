#!/usr/bin/env bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

run_linters() {
    if ! golangci-lint run; then
        echo "golangci-lint failed!"
        exit 1
    fi

    echo "golangci-lint passed!"
}

run_unit_tests() {
    echo "running unit tests on ${SCRIPT_DIR}"
    if ! (
        cd "${SCRIPT_DIR}"
        go test -v -race -coverpkg=./... -coverprofile=coverage.txt ./...
    ); then
        echo "unit tests failed!"
        exit 1
    fi
}

run_coverage() {
    coverage=$(go tool cover -func coverage.txt | grep total | grep -Eo '[0-9]+\.[0-9]+')
    echo "Total Coverage: $coverage%"
}

run_all() {
    run_linters

    run_unit_tests

    run_coverage

    echo "All validations passed!"
}

run_all