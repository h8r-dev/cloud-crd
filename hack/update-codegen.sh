#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# corresponding to go mod init <module>
MODULE=github.com/h8r-dev/cloud-crd
# api package
APIS_PKG=api
# generated output package
OUTPUT_PKG=generated/cloud
# group-version such as foo:v1alpha1
GROUP=cloud
VERSION=v1alpha1
GROUP_VERSION=${GROUP}:${VERSION}

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

GROUP_DIR=${SCRIPT_ROOT}/api/cloud/v1alpha1
rm -rf $GROUP_DIR
cp -r api/v1alpha1 $GROUP_DIR

rm -rf ${OUTPUT_PKG}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
bash "${CODEGEN_PKG}"/generate-groups.sh "client,lister,informer" \
  ${MODULE}/${OUTPUT_PKG} ${MODULE}/${APIS_PKG} \
  ${GROUP_VERSION} \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt \
  # --output-base "${SCRIPT_ROOT}"
#  --output-base "${SCRIPT_ROOT}/../../.." \

mv ${MODULE}/${OUTPUT_PKG} ${OUTPUT_PKG}
