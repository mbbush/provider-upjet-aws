#!/bin/bash

XP_PROVIDER_PATH=`pwd`

cd $TF_PROVIDER_PATH/internal/service
git grep '// @FrameworkResource(' | sed -n 's|\([^/]*\)/\(\w*\)\.go:// @FrameworkResource([ "]*\(aws_[a-z0-9_]*\).*|\1 \2 \3|p' > $XP_PROVIDER_PATH/hack/framework-resources.lst
git grep '// @SDKResource(' | sed -n 's|\([^/]*\)/\(\w*\)\.go:// @SDKResource([ "]*\(aws_[a-z0-9_]*\).*|\1 \2 \3|p' > $XP_PROVIDER_PATH/hack/sdk-resources.lst