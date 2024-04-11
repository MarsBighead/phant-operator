
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
SKIP_CRD_FILES=(
    "pingcap.com_dataresources.yaml"
)
cd $PRO_ROOT


source hack/lib.sh
CONTROLLER_GEN=${OUTPUT_BIN}/controller-gen
hack__ensure_controller_gen

echo "Generating CRDs ..."

API_PACKAGES="phant-operator/pkg/apis/phant/v1alpha1/..."
CRD_OUTPUT_DIR=${PRO_ROOT}/manifests/crd
CRD_OPTIONS="preserveUnknownFields=false,allowDangerousTypes=true,maxDescLen=0"
echo "${CRD_OUTPUT_DIR}"
# generate CRDs
${CONTROLLER_GEN} \
    crd:crdVersions=v1beta1,${CRD_OPTIONS} \
    paths=${API_PACKAGES} \
    output:crd:dir=${CRD_OUTPUT_DIR}/v1beta1
${CONTROLLER_GEN} \
    crd:crdVersions=v1,${CRD_OPTIONS} \
    paths=${API_PACKAGES} \
    output:crd:dir=${CRD_OUTPUT_DIR}/v1

for file in ${SKIP_CRD_FILES[@]}; do
    rm -f ${CRD_OUTPUT_DIR}/v1beta1/${file}
    rm -f ${CRD_OUTPUT_DIR}/v1/${file}
done