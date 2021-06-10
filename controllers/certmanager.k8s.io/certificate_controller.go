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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

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

	r.Log.Info("##### DEBUG ##### testing certificate controller")
	certificate := v1certmgr.Certificate{
		TypeMeta:   metav1.TypeMeta{Kind: "Certificate", APIVersion: "cert-manager.io/v1"},
		ObjectMeta: metav1.ObjectMeta{Name: "test-conversion", Namespace: "cmd-operator-system"},
		Spec: v1certmgr.CertificateSpec{
			CommonName: "test-ca-cert",
			Duration:   &metav1.Duration{Duration: time.Hour},
			IssuerRef: v1certmgrmeta.ObjectReference{
				Name:  "demo-ss-issuer",
				Kind:  "Issuer",
				Group: "cert-manager.io",
			},
			IsCA:        true,
			RenewBefore: &metav1.Duration{Duration: time.Minute * 59},
			SecretName:  "test-ca-cert-secret",
		},
	}
	if err := r.Client.Create(context.TODO(), &certificate); err != nil {
		r.Log.Error(err, "##### DEBUG ##### creating v1 Cert")
	}

	return ctrl.Result{}, nil
}

func (r *CertificateReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&certmanagerk8siov1alpha1.Certificate{}).
		Complete(r)
}
