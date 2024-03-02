#!/usr/bin/env bash

project_path=$(go list -m -f "{{.Dir}}")

echo "${project_path}"