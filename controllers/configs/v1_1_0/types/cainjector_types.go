package types

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CertManagerCAInjectorConfig struct {
	metav1.TypeMeta `json:",inline"`

	Flags CertManagerCAInjectorFlags `json:"flags"`
}

type CertManagerCAInjectorFlags struct {
	loggingFlags
	Kubeconfig              string `json:"kubeconfig"`
	Master                  string `json:"master"`
	Namespace               string `json:"namespace"`
	LeaderElect             bool   `json:"leader-elect"`
	LeaderElectionNamespace string `json:"leader-election-namespace"`
}
