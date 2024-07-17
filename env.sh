#!/bin/bash
# 生成构建时的环境变量
export VERSION=$(git describe --tags --abbrev=0)
export COMMIT_HASH=$(git rev-parse HEAD)
export BRANCH_NAME=$(git rev-parse --abbrev-ref HEAD)
export BUILD_TIME=$(date "+%Y-%m-%d %H:%M:%S")
export LAST_COMMIT_LOG=$(git log -1 --pretty=%B)
export GO_VERSION=$(go version | awk '{print $3}') # 例如 "go1.15.2"
export GIT_SHA1=$(git show-ref --head --hash=8 2> /dev/null | head -n1 || echo 00000000)
export GIT_DIRTY=$(git diff --no-ext-diff 2> /dev/null | wc -l)
export BUILD_ID=$(uname -n)-$(date +%s)
if [ -n "$SOURCE_DATE_EPOCH" ]; then
  export BUILD_ID=$(date -u -d "@$SOURCE_DATE_EPOCH" +%s 2>/dev/null || date -u -r "$SOURCE_DATE_EPOCH" +%s 2>/dev/null || date -u +%s)
fi