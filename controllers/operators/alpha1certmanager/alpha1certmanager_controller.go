/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package alpha1certmanager

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	issuerKindAnnotation string = "cert-manager.io/issuer-kind"
)

// PodRefreshReconciler reconciles a Secret object
type Alpha1CertManagerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	record.EventRecorder
}

// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets/status,verbs=get;update;patch;
// +kubebuilder:rbac:groups=apps,resources=daemonsets,verbs=list;update;watch;
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=list;update;watch;
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=list;update;watch;

// Reconcile watches for secrets and if a secret is a certmanager secret, it checks for deployments, statefulsets,
// and daemonsets that may be using the secret and triggers a re-rollout of those objects.
func (r *Alpha1CertManagerReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("secret", req.NamespacedName)

	r.Log.Info("##### DEBUG ##### testing alpha1certmanager controller")

	return ctrl.Result{}, nil
}

// SetupWithManager configures a controller owned by the manager mgr.
func (r *Alpha1CertManagerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Secret{}).
		WithEventFilter(predicate.And(
			predicate.ResourceVersionChangedPredicate{},
			isCertManagerIssuedTLSPredicate{},
		)).
		Complete(r)
}

// isCertManagerIssuedTLSPredicate implements a predicate verifying secret object meta indicates a
// cert-manager issued TLS certificate. This only applies to Create and Update events. Deletes
// and Generics should not make it to the work queue.
type isCertManagerIssuedTLSPredicate struct{}

// Update implements default UpdateEvent filter for validating if object has the
// `cert-manager.io/issuer-kind` which helps identify that it is cert-manager issued.
// Intended to be used with secrets.
func (isCertManagerIssuedTLSPredicate) Update(e event.UpdateEvent) bool {
	a := e.MetaNew.GetAnnotations()
	_, isCertManagerIssued := a[issuerKindAnnotation]
	return isCertManagerIssued
}

// Create implements default CreateEvent filter for validating if object has the
// `cert-manager.io/issuer-kind` which helps identify that it is cert-manager issued.
// Intended to be used with secrets.
func (isCertManagerIssuedTLSPredicate) Create(e event.CreateEvent) bool {
	a := e.Meta.GetAnnotations()
	_, isCertManagerIssued := a[issuerKindAnnotation]
	return isCertManagerIssued
}

func (isCertManagerIssuedTLSPredicate) Delete(e event.DeleteEvent) bool {
	return false
}

func (isCertManagerIssuedTLSPredicate) Generic(e event.GenericEvent) bool {
	return false
}
