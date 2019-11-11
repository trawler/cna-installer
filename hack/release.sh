#!/bin/bash -eux

# Build docker container
docker build -t cna_build -f Dockerfile .

# Extract compiled binaries
docker run --rm --name cna_build --entrypoint tar cna_build --create terraform cna-installer | tar --extract --verbose

# See the binaries
ls -la terraform cna-installer
