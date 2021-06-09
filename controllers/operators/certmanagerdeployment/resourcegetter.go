package certmanagerdeployment

import operatorsv1alpha1 "github.com/komish/cmd-operator-dev/apis/operators/v1alpha1"

// ResourceGetter facilitates getting various owned resources expected by
// a CertManagerDeployment CR.
type ResourceGetter struct {
	CustomResource operatorsv1alpha1.CertManagerDeployment
}
