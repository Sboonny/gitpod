#!/bin/bash
# Copyright (c) 2021 Gitpod GmbH. All rights reserved.
# Licensed under the GNU Affero General Public License (AGPL).
# See License-AGPL.txt in the project root for license information.

# exclude 'e' (exit on any error)
# there are many test binaries, each can have failures


# shellcheck disable=SC2045
CURRENT=$(pwd)
PATH=$2


cd "${PATH}" || echo "Path invalid ${PATH}"
go test -v ./... "-kubeconfig=$1" -namespace=gitpod -username=gitpod-integration-test 2>&0 -coverprofile=coverage.out
TEST_STATUS=$?
if [ "$TEST_STATUS" -ne "0" ]; then
    echo "Test failed at $(date)"
else
    echo "Test succeeded at $(date)"
fi;

cd "${CURRENT}" || echo "Couldn't move back to test dir"
