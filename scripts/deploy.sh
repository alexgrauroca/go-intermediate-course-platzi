#!/bin/bash
parent_path=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )
cd "${parent_path/scripts/}"
git reset --hard
git clean -fd
git pull
go mod vendor
