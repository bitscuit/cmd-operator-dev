#!/usr/bin/env bash

BASE_DIR=${1-.}
OPTS=${2}

# echo $OPTS

. ${BASE_DIR}/demo-magic.sh ${OPTS}

pei "oc apply -f ${BASE_DIR}/demo_yamls/ss-issuer.yaml"

wait

pei "oc apply -f ${BASE_DIR}/demo_yamls/ca-cert.yaml"

wait

pei "oc apply -f ${BASE_DIR}/demo_yamls/cr.yaml"
