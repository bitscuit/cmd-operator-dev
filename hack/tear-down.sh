#!/usr/bin/env bash

BASE_DIR=${1-.}
OPTS=${2}

# echo $OPTS

. ${BASE_DIR}/demo-magic.sh ${OPTS}

pei "oc delete -f ${BASE_DIR}/demo_yamls/cr.yaml"

wait

pei "oc delete -f ${BASE_DIR}/demo_yamls/ca-cert.yaml"

pei "oc delete -f ${BASE_DIR}/demo_yamls/ss-issuer.yaml"

pei "oc delete cert -n cmd-operator-system test-conversion"

pei "oc delete secret -n cmd-operator-system test-ca-cert-secret"
