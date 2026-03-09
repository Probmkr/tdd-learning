#!/bin/bash
set -e

if [ -z "$1" ]; then
  echo "Usage: $0 <project-name>"
  exit 1
fi

mkdir "$1"

mvn archetype:generate \
  -DgroupId=money \
  -DartifactId="$1" \
  -DarchetypeArtifactId=maven-archetype-quickstart \
  -DarchetypeVersion=1.5 \
  -DinteractiveMode=false \
  -DoutputDirectory=.
