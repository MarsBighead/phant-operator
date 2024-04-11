
#!/usr/bin/env bash

# Authors: Marsbighead <duanhmhy@126.com>
#
# Copyright (c) 2024 Marsbighead
#
# Permission is hereby granted, free of charge, to any person
# obtaining a copy of this software and associated documentation
# files (the "Software"), to deal in the Software without
# restriction, including without limitation the rights to use,
# copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the
# Software is furnished to do so, subject to the following
# conditions:
#
# The above copyright notice and this permission notice shall be
# included in all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
# EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
# OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
# HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
# WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
# OTHER DEALINGS IN THE SOFTWARE.

set -o errexit
set -o nounset
set -o pipefail

PRO_ROOT=$(unset CDPATH && cd $(dirname "${BASH_SOURCE[0]}")/.. && pwd)
cd $PRO_ROOT
source hack/lib.sh
hack__ensure_codegen

# `--output-base $ROOT` will output generated code to current dir
GOBIN=$OUTPUT_BIN bash $PRO_ROOT/hack/generate-groups.sh "deepcopy,client,informer,lister" \
    phant-operator/pkg/client\
    phant-operator/pkg/apis \
    phant:v1alpha1 \
    --output-base $PRO_ROOT \
    --go-header-file ./hack/boilerplate/boilerplate.generatego.txt

# then we merge generated code with our code base and clean up
cp -r phant-operator/pkg $PRO_ROOT && rm -rf phant-operator