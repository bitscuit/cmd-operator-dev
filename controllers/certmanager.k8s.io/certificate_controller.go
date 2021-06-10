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

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	v1certmgr "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1"
	v1certmgrmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	certmanagerk8siov1alpha1 "github.com/komish/cmd-operator-dev/apis/certmanager.k8s.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateReconciler reconciles a Certificate object
type CertificateReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=certmanager.k8s.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=certmanager.k8s.io,resources=certificates/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cert-manager.io,resources=certificates,verbs=get;list;watch;create;update;patch;delete

func (r *CertificateReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("certificate", req.NamespacedName)

	r.Log.Info("##### DEBUG ##### Reconciling alpha1 Certificate")

	r.Log.Info("Get Certificate fields")

	alpha1Cert := &certmanagerk8siov1alpha1.Certificate{}
	err := r.Get(context.TODO(), req.NamespacedName, alpha1Cert)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile req.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the req.
		return reconcile.Result{}, err
	}

	r.Log.Info("Converting to v1 Certificate")

	certificate := v1certmgr.Certificate{
		TypeMeta: metav1.TypeMeta{Kind: "Certificate", APIVersion: "cert-manager.io/v1"},
		ObjectMeta: metav1.ObjectMeta{
			Name:      alpha1Cert.Name,
			Namespace: alpha1Cert.Namespace,
			Annotations: map[string]string{
				"ibm-cert-manager-operator-generated": "true",
			},
		},
		Spec: v1certmgr.CertificateSpec{
			CommonName: alpha1Cert.Spec.CommonName,
			Duration:   &metav1.Duration{Duration: alpha1Cert.Spec.Duration.Duration},
			IssuerRef: v1certmgrmeta.ObjectReference{
				Name:  alpha1Cert.Spec.IssuerRef.Name,
				Kind:  alpha1Cert.Spec.IssuerRef.Kind,
				Group: alpha1Cert.Spec.IssuerRef.Group,
			},
			IsCA:        alpha1Cert.Spec.IsCA,
			RenewBefore: &metav1.Duration{Duration: alpha1Cert.Spec.RenewBefore.Duration},
			SecretName:  alpha1Cert.Spec.CommonName,
		},
	}
	if err := r.Client.Create(context.TODO(), &certificate); err != nil {
		r.Log.Error(err, "##### DEBUG ##### Failed to create v1 Certificate")
	}

	r.Log.Info("##### DEBUG ##### Finished reconciling alpha1 Certificate")

	return ctrl.Result{}, nil
}

func (r *CertificateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&certmanagerk8siov1alpha1.Certificate{}).
		Complete(r)
}
