#!/usr/bin/env bash

BASE_DIR=${1-.}
OPTS=${2}

# echo $OPTS

. ${BASE_DIR}/demo-magic.sh ${OPTS}

pei "oc delete -f ${BASE_DIR}/demo_yamls/cr.yaml"

wait

pei "oc delete -f ${BASE_DIR}/demo_yamls/ca-cert.yaml"

pei "oc delete -f ${BASE_DIR}/demo_yamls/ss-issuer.yaml"

pei "oc get cert -n cmd-operator-system | grep playback | cut -d ' ' -f1 | xargs oc delete -n cmd-operator-system cert"

pei "oc get secret -n cmd-operator-system | grep playback | cut -d ' ' -f1 | xargs oc delete -n cmd-operator-system secret"
