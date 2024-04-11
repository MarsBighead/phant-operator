#!/usr/bin/env bash

if [ -z "$PRO_ROOT" ]; then
    echo "error: PRO_ROOT should be initialized"
    exit 1
fi

OS=$(go env GOOS)
ARCH=$(go env GOARCH)
OUTPUT=${PRO_ROOT}/output
OUTPUT_BIN=${OUTPUT}/bin

KUBECTL_VERSION=${KUBECTL_VERSION:-1.20.2}
KUBECTL_BIN=$OUTPUT_BIN/kubectl
HELM_BIN=$OUTPUT_BIN/helm
DOCS_BIN=$OUTPUT_BIN/gen-crd-api-reference-docs
# K8S_VERSION should keep the same version with package k8s.io/* in go.mod
K8S_VERSION=${K8S_VERSION:-0.29.3}



echo $OUTPUT_BIN
test -d "$OUTPUT_BIN" || mkdir -p "$OUTPUT_BIN"

function hack__ensure_codegen() {
    echo "Installing codegen..."
    GOBIN=$OUTPUT_BIN go install k8s.io/code-generator/cmd/{defaulter-gen,client-gen,lister-gen,informer-gen,deepcopy-gen}@v$K8S_VERSION
}

function hack__ensure_openapi() {
    echo "Installing openpi_gen..."
    GOBIN=$OUTPUT_BIN go install k8s.io/code-generator/cmd/openapi-gen@v$K8S_VERSION
}


function hack__ensure_gen_crd_api_references_docs() {
    echo "Installing gen_crd_api_references_docs..."
    GOBIN=$OUTPUT_BIN go install github.com/xhebox/gen-crd-api-reference-docs@e46d84594a6d158ec7123ff05acd57acf62e140f
}

function hack__ensure_controller_gen() {
    echo "Installing controller_gen..."
    GOBIN=$OUTPUT_BIN go install sigs.k8s.io/controller-tools/cmd/controller-gen@v0.6.2
}
