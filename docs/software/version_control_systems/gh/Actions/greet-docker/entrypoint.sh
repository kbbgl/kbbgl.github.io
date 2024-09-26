#!/bin/sh -l

# https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions
# This is how we set a debug msg
echo "::debug ::Debug Message"
echo "::warning ::Warning Message"
echo "::error ::Error Message"

# obfuscate
echo "::add-mask ::$1"

# output
time=$(date)
echo "::set-output name=time::$time"

# log group
echo "::group::Expandable logs"
echo "Some output"
echo "Some output"
echo "Some output"
echo "::endgroup::"

# set env
# https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#setting-an-environment-variable

echo "Hello $1"