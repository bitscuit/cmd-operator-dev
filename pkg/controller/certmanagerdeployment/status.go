package certmanagerdeployment

import (
	"context"

	"github.com/go-logr/logr"
	redhatv1alpha1 "github.com/komish/certmanager-operator/pkg/apis/redhat/v1alpha1"
	"github.com/komish/certmanager-operator/pkg/controller/certmanagerdeployment/componentry"
	appsv1 "k8s.io/api/apps/v1"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// deploymentState is a type to help facilitate reading the current state of existing deployments
// in the cluster.
type deploymentState struct {
	// Count is the number of deployments
	count                 int
	availableMatchesReady []bool
	readyMatchesDesired   []bool
}

// allAvailableAreReady return true if all deployments add to the deploymentState struct have matching
// available replicas (as determined by their status) and matching ready replicas (as determined by their status).
func (ds *deploymentState) allAvailableAreReady() bool {
	res := true
	for _, v := range ds.availableMatchesReady {
		if !v {
			res = false
		}
	}

	return res
}

// readyCountMatchesDesiredCount return true if all the deployments added to the deploymentState struct have matching
// ready replicas (as determined by their status) and matching desired replicas (as determined by their spec).
func (ds *deploymentState) readyCountMatchesDesiredCount() bool {
	res := true
	for _, v := range ds.readyMatchesDesired {
		if !v {
			res = false
		}
	}

	return res
}

// deploymentCountMatchesCountOfStoredStates will return true if the number of stored states match the expected number of deployments
// to be evaluated as a part of this struct (stored in count).
func (ds *deploymentState) deploymentCountMatchesCountOfStoredStates() bool {
	return len(ds.availableMatchesReady) == len(ds.readyMatchesDesired) && ds.count == len(ds.availableMatchesReady)
}

// crdState is a type to help facilitate reading the current state of existing CRDs
// in the cluster.
type crdState struct {
	// Count is the number of CRDs
	count             int
	crdIsEstablished  []bool
	crdNameIsAccepted []bool
}

// allAreEstablished return true if all CRD added to the crdState struct have a status.condition
// Established and the status of that condition is true.
func (cs *crdState) allAreEstablished() bool {
	res := true
	for _, v := range cs.crdIsEstablished {
		if !v {
			res = false
		}
	}

	return res
}

// allNamesAreAccepted returns true if all the CRDs added to the crdState struct have a status.condition
// NameAccepted and the status of that condition is true.
func (cs *crdState) allNamesAreAccepted() bool {
	res := true
	for _, v := range cs.crdNameIsAccepted {
		if !v {
			res = false
		}
	}

	return res
}

// crdCountMatchesCountOfStoredStates will return true if the number of stored states match the expected number of CRDs
// to be evaluated as a part of this struct (stored in count).
func (cs *crdState) crdCountMatchesCountOfStoredStates() bool {
	return len(cs.crdIsEstablished) == len(cs.crdNameIsAccepted) && cs.count == len(cs.crdIsEstablished)
}

// getUninitializedCertManagerDeploymentStatus returns a CertManagerDeploymentStatus with unknown values
// to be modified and added to the API.
func getUninitializedCertManagerDeploymentStatus() *redhatv1alpha1.CertManagerDeploymentStatus {
	return &redhatv1alpha1.CertManagerDeploymentStatus{
		Version:            "unknown",
		Phase:              "unknown",
		DeploymentsHealthy: false,
		CRDsHealthy:        false,
	}
}

// deploymentsAreReady evaluates the status fields in existingDeploys, and return true
// if all existingDeploys are in an acceptable state.
func deploymentsAreReady(existingDeploys []*appsv1.Deployment) bool {
	state, ok := deploymentState{count: len(existingDeploys)}, false
	for _, deploy := range existingDeploys {
		state.availableMatchesReady = append(state.availableMatchesReady, (deploy.Status.AvailableReplicas == deploy.Status.ReadyReplicas))
		state.readyMatchesDesired = append(state.readyMatchesDesired, (*deploy.Spec.Replicas == deploy.Status.ReadyReplicas))
	}

	if state.allAvailableAreReady() &&
		state.readyCountMatchesDesiredCount() &&
		state.deploymentCountMatchesCountOfStoredStates() {
		ok = true
	}

	return ok
}

// crdsAreReady evaluates the status fields in existingCRDs and returns true
// if all existingCRDs are in an acceptable state.
func crdsAreReady(existingCRDs []*apiextv1beta1.CustomResourceDefinition) bool {
	state, ok := crdState{count: len(existingCRDs)}, false
	for _, crd := range existingCRDs {
		// search specifically for conditions we care about - NamesAccepted and Established
		// it is possible that NamesAccepted or Established is not available in the condition list
		// in which case we want to return false.
		accepted, established := false, false
		conditions := crd.Status.Conditions
		for _, condition := range conditions {
			switch condition.Type {
			case "Established":
				if condition.Status == apiextv1beta1.ConditionTrue {
					established = true
				}
			case "NamesAccepted":
				if condition.Status == apiextv1beta1.ConditionTrue {
					accepted = true
				}
			}
		}
		state.crdIsEstablished = append(state.crdIsEstablished, established)
		state.crdNameIsAccepted = append(state.crdNameIsAccepted, accepted)
	}
	if state.allAreEstablished() &&
		state.allNamesAreAccepted() &&
		state.crdCountMatchesCountOfStoredStates() {
		ok = true
	}

	return ok
}

// reconcileStatusVersion is a subreconciliation function called by ReconcileStatus that injects the version
// of cert-manager desired into the status field.
func (r *ReconcileCertManagerDeployment) reconcileStatusVersion(inStatus *redhatv1alpha1.CertManagerDeploymentStatus, vers string) *redhatv1alpha1.CertManagerDeploymentStatus {
	inStatus.Version = vers
	// TODO() check deployed versions of things to make sure they reflect the right version.
	// requires updating all types that we create with some form of identifier stating which
	// version is deployed for that CR.
	return inStatus
}

// reconcileStatusDeploymentsHealthy updated the DeploymentsHealthy status field. This
// checks if the deployments expected for the CR instance exists per the API and are
// ready and available with the expected replica count.
func (r *ReconcileCertManagerDeployment) reconcileStatusDeploymentsHealthy(
	inStatus *redhatv1alpha1.CertManagerDeploymentStatus,
	rg ResourceGetter,
	reqLogger logr.Logger) *redhatv1alpha1.CertManagerDeploymentStatus {

	existingDeploys, ok := queryAPIForExpectedDeployments(r.client, rg, reqLogger)
	if existingDeploys == nil {
		reqLogger.Info("unable to determine status of expected deployments for this instance")
	}

	if ok {
		inStatus.DeploymentsHealthy = deploymentsAreReady(existingDeploys)
	}

	return inStatus
}

// reconcileStatusCRDsHealthy updated the DeploymentsHealthy status field. This
// checks if the deployments expected for the CR instance exists per the API and are
// ready and available with the expected replica count.
func (r *ReconcileCertManagerDeployment) reconcileStatusCRDsHealthy(
	inStatus *redhatv1alpha1.CertManagerDeploymentStatus,
	rg ResourceGetter,
	reqLogger logr.Logger) *redhatv1alpha1.CertManagerDeploymentStatus {

	existingCRDs, ok := queryAPIForExpectedCRDs(r.client, rg, reqLogger)
	if existingCRDs == nil {
		reqLogger.Info("unable to determine status of expected CRDs for this instance")
	}

	if ok {
		inStatus.CRDsHealthy = crdsAreReady(existingCRDs)
	}

	return inStatus
}

// queryAPIForExpectedDeployments will check that the deployments expected for a given instance actually
// exist in the API. will return ok as true when all were found, and false if not. Will return the
// deployment slice as nil if an error other than IsNotfound was encountered trying to obtain the data, as well
// as return ok as false.
func queryAPIForExpectedDeployments(client client.Client, rg ResourceGetter, reqLogger logr.Logger) ([]*appsv1.Deployment, bool) {
	foundDeployments := make([]*appsv1.Deployment, 0)
	var ok bool

	expectedDeployments := rg.GetDeployments()
	for _, deploy := range expectedDeployments {
		receiver := appsv1.Deployment{}
		err := client.Get(context.TODO(), types.NamespacedName{Name: deploy.GetName(), Namespace: deploy.GetNamespace()}, &receiver)
		if err != nil {
			if errors.IsNotFound(err) {
				// if we got an IsNotFound error, we later make sure that ok is false.
				continue
			} else {
				reqLogger.Error(err, "unable to query for existing deployment")
				return nil, false
			}
		}

		foundDeployments = append(foundDeployments, &receiver)
	}

	// evaluate if we found what we expected
	ok = len(foundDeployments) == len(expectedDeployments)

	return foundDeployments, ok
}

// queryAPIForExpectedCRDs will check that the CRDs expected for a given instance actually
// exist in the API. will return ok as true when all were found, and false if not. Will return the
// CRD slice as nil if an error other than IsNotfound was encountered trying to obtain the data, as well as
// set return ok as false.
func queryAPIForExpectedCRDs(client client.Client, rg ResourceGetter, reqLogger logr.Logger) ([]*apiextv1beta1.CustomResourceDefinition, bool) {
	foundCRDs := make([]*apiextv1beta1.CustomResourceDefinition, 0)
	var ok bool

	expectedCRDs, err := rg.GetCRDs()
	if err != nil {
		// GetCRDs returns an error in case it can't read the CRD from the filesystem.
		// we have to handle it.
		return nil, false
	}

	for _, crd := range expectedCRDs {
		receiver := apiextv1beta1.CustomResourceDefinition{}
		err := client.Get(context.TODO(), types.NamespacedName{Name: crd.GetName()}, &receiver)
		if err != nil {
			if errors.IsNotFound(err) {
				// if we got an IsNotFound error, we later make sure that ok is false.
				continue
			} else {
				reqLogger.Error(err, "unable to query for existing custom resource definitions")
				return nil, false
			}
		}

		foundCRDs = append(foundCRDs, &receiver)
	}

	// evaluate if we found what we expected
	ok = len(foundCRDs) == len(expectedCRDs)

	return foundCRDs, ok
}

// reconcileStatusPhase will update the status phase indicator based on the discovered status of deployments
// and CRDs. This must run after DeploymentsHealthy and CRDsHealthy have been updated by the status reconciler.
func (r *ReconcileCertManagerDeployment) reconcileStatusPhase(inStatus *redhatv1alpha1.CertManagerDeploymentStatus) *redhatv1alpha1.CertManagerDeploymentStatus {

	if inStatus.CRDsHealthy && inStatus.DeploymentsHealthy {
		inStatus.Phase = componentry.StatusPhaseRunning
	} else if !inStatus.CRDsHealthy || !inStatus.DeploymentsHealthy {
		inStatus.Phase = componentry.StatusPhasePending
	}

	return inStatus
}

// func reconcileStatusDeploymentConditions() {}
// func reconcileStatusCRDConditions() {}
