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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ObjectReference struct {
	Name string `json:"name"`
	// +optional
	Kind string `json:"kind,omitempty"`
	// +optional
	Group string `json:"group,omitempty"`
}

// HTTP01SolverConfig contains solver configuration for HTTP01 challenges.
type HTTP01SolverConfig struct {
	// Ingress is the name of an Ingress resource that will be edited to include
	// the ACME HTTP01 'well-known' challenge path in order to solve HTTP01
	// challenges.
	// If this field is specified, 'ingressClass' **must not** be specified.
	// +optional
	Ingress string `json:"ingress,omitempty"`

	// IngressClass is the ingress class that should be set on new ingress
	// resources that are created in order to solve HTTP01 challenges.
	// This field should be used when using an ingress controller such as nginx,
	// which 'flattens' ingress configuration instead of maintaining a 1:1
	// mapping between loadbalancer IP:ingress resources.
	// If this field is not set, and 'ingress' is not set, then ingresses
	// without an ingress class set will be created to solve HTTP01 challenges.
	// If this field is specified, 'ingress' **must not** be specified.
	// +optional
	IngressClass *string `json:"ingressClass,omitempty"`
}

// DNS01SolverConfig contains solver configuration for DNS01 challenges.
type DNS01SolverConfig struct {
	// Provider is the name of the DNS01 challenge provider to use, as configure
	// on the referenced Issuer or ClusterIssuer resource.
	Provider string `json:"provider"`
}

// SolverConfig is a container type holding the configuration for either a
// HTTP01 or DNS01 challenge.
// Only one of HTTP01 or DNS01 should be non-nil.
type SolverConfig struct {
	// HTTP01 contains HTTP01 challenge solving configuration
	// +optional
	HTTP01 *HTTP01SolverConfig `json:"http01,omitempty"`

	// DNS01 contains DNS01 challenge solving configuration
	// +optional
	DNS01 *DNS01SolverConfig `json:"dns01,omitempty"`
}

// DomainSolverConfig contains solver configuration for a set of domains.
type DomainSolverConfig struct {
	// Domains is the list of domains that this SolverConfig applies to.
	Domains []string `json:"domains"`

	// SolverConfig contains the actual solver configuration to use for the
	// provided set of domains.
	SolverConfig `json:",inline"`
}

// ACMECertificateConfig contains the configuration for the ACME certificate provider
type ACMECertificateConfig struct {
	Config []DomainSolverConfig `json:"config"`
}

// KeyUsage specifies valid usage contexts for keys.
// See: https://tools.ietf.org/html/rfc5280#section-4.2.1.3
//      https://tools.ietf.org/html/rfc5280#section-4.2.1.12
// +kubebuilder:validation:Enum="signing";"digital signature";"content commitment";"key encipherment";"key agreement";"data encipherment";"cert sign";"crl sign";"encipher only";"decipher only";"any";"server auth";"client auth";"code signing";"email protection";"s/mime";"ipsec end system";"ipsec tunnel";"ipsec user";"timestamping";"ocsp signing";"microsoft sgc";"netscape sgc"
type KeyUsage string

const (
	UsageSigning            KeyUsage = "signing"
	UsageDigitalSignature   KeyUsage = "digital signature"
	UsageContentCommittment KeyUsage = "content commitment"
	UsageKeyEncipherment    KeyUsage = "key encipherment"
	UsageKeyAgreement       KeyUsage = "key agreement"
	UsageDataEncipherment   KeyUsage = "data encipherment"
	UsageCertSign           KeyUsage = "cert sign"
	UsageCRLSign            KeyUsage = "crl sign"
	UsageEncipherOnly       KeyUsage = "encipher only"
	UsageDecipherOnly       KeyUsage = "decipher only"
	UsageAny                KeyUsage = "any"
	UsageServerAuth         KeyUsage = "server auth"
	UsageClientAuth         KeyUsage = "client auth"
	UsageCodeSigning        KeyUsage = "code signing"
	UsageEmailProtection    KeyUsage = "email protection"
	UsageSMIME              KeyUsage = "s/mime"
	UsageIPsecEndSystem     KeyUsage = "ipsec end system"
	UsageIPsecTunnel        KeyUsage = "ipsec tunnel"
	UsageIPsecUser          KeyUsage = "ipsec user"
	UsageTimestamping       KeyUsage = "timestamping"
	UsageOCSPSigning        KeyUsage = "ocsp signing"
	UsageMicrosoftSGC       KeyUsage = "microsoft sgc"
	UsageNetscapSGC         KeyUsage = "netscape sgc"
)

// +kubebuilder:validation:Enum=rsa;ecdsa
type KeyAlgorithm string

const (
	RSAKeyAlgorithm   KeyAlgorithm = "rsa"
	ECDSAKeyAlgorithm KeyAlgorithm = "ecdsa"
)

// +kubebuilder:validation:Enum=pkcs1;pkcs8
type KeyEncoding string

const (
	PKCS1 KeyEncoding = "pkcs1"
	PKCS8 KeyEncoding = "pkcs8"
)

// CertificateSpec defines the desired state of Certificate
type CertificateSpec struct {
	// CommonName is a common name to be used on the Certificate.
	// If no CommonName is given, then the first entry in DNSNames is used as
	// the CommonName.
	// The CommonName should have a length of 64 characters or fewer to avoid
	// generating invalid CSRs; in order to have longer domain names, set the
	// CommonName (or first DNSNames entry) to have 64 characters or fewer,
	// and then add the longer domain name to DNSNames.
	// +optional
	CommonName string `json:"commonName,omitempty"`

	// Organization is the organization to be used on the Certificate
	// +optional
	Organization []string `json:"organization,omitempty"`

	// Certificate default Duration
	// +optional
	Duration *metav1.Duration `json:"duration,omitempty"`

	// Certificate renew before expiration duration
	// +optional
	RenewBefore *metav1.Duration `json:"renewBefore,omitempty"`

	// DNSNames is a list of subject alt names to be used on the Certificate.
	// If no CommonName is given, then the first entry in DNSNames is used as
	// the CommonName and must have a length of 64 characters or fewer.
	// +optional
	DNSNames []string `json:"dnsNames,omitempty"`

	// IPAddresses is a list of IP addresses to be used on the Certificate
	// +optional
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// SecretName is the name of the secret resource to store this secret in
	SecretName string `json:"secretName"`

	// IssuerRef is a reference to the issuer for this certificate.
	// If the 'kind' field is not set, or set to 'Issuer', an Issuer resource
	// with the given name in the same namespace as the Certificate will be used.
	// If the 'kind' field is set to 'ClusterIssuer', a ClusterIssuer with the
	// provided name will be used.
	// The 'name' field in this stanza is required at all times.
	IssuerRef ObjectReference `json:"issuerRef"`

	// IsCA will mark this Certificate as valid for signing.
	// This implies that the 'cert sign' usage is set
	// +optional
	IsCA bool `json:"isCA,omitempty"`

	// Usages is the set of x509 actions that are enabled for a given key. Defaults are ('digital signature', 'key encipherment') if empty
	// +optional
	Usages []KeyUsage `json:"usages,omitempty"`

	// ACME contains configuration specific to ACME Certificates.
	// Notably, this contains details on how the domain names listed on this
	// Certificate resource should be 'solved', i.e. mapping HTTP01 and DNS01
	// providers to DNS names.
	// +optional
	ACME *ACMECertificateConfig `json:"acme,omitempty"`

	// KeySize is the key bit size of the corresponding private key for this certificate.
	// If provided, value must be between 2048 and 8192 inclusive when KeyAlgorithm is
	// empty or is set to "rsa", and value must be one of (256, 384, 521) when
	// KeyAlgorithm is set to "ecdsa".
	// +optional
	KeySize int `json:"keySize,omitempty"`

	// KeyAlgorithm is the private key algorithm of the corresponding private key
	// for this certificate. If provided, allowed values are either "rsa" or "ecdsa"
	// If KeyAlgorithm is specified and KeySize is not provided,
	// key size of 256 will be used for "ecdsa" key algorithm and
	// key size of 2048 will be used for "rsa" key algorithm.
	// +optional
	KeyAlgorithm KeyAlgorithm `json:"keyAlgorithm,omitempty"`

	// KeyEncoding is the private key cryptography standards (PKCS)
	// for this certificate's private key to be encoded in. If provided, allowed
	// values are "pkcs1" and "pkcs8" standing for PKCS#1 and PKCS#8, respectively.
	// If KeyEncoding is not specified, then PKCS#1 will be used by default.
	KeyEncoding KeyEncoding `json:"keyEncoding,omitempty"`
}

// CertificateStatus defines the observed state of Certificate
type CertificateStatus struct {
	// +optional
	Conditions []CertificateCondition `json:"conditions,omitempty"`

	// +optional
	LastFailureTime *metav1.Time `json:"lastFailureTime,omitempty"`

	// The expiration time of the certificate stored in the secret named
	// by this resource in spec.secretName.
	// +optional
	NotAfter *metav1.Time `json:"notAfter,omitempty"`
}

// CertificateCondition contains condition information for an Certificate.
type CertificateCondition struct {
	// Type of the condition, currently ('Ready').
	Type CertificateConditionType `json:"type"`

	// Status of the condition, one of ('True', 'False', 'Unknown').
	Status ConditionStatus `json:"status"`

	// LastTransitionTime is the timestamp corresponding to the last status
	// change of this condition.
	// +optional
	LastTransitionTime *metav1.Time `json:"lastTransitionTime,omitempty"`

	// Reason is a brief machine readable explanation for the condition's last
	// transition.
	// +optional
	Reason string `json:"reason,omitempty"`

	// Message is a human readable description of the details of the last
	// transition, complementing reason.
	// +optional
	Message string `json:"message,omitempty"`
}

// CertificateConditionType represents an Certificate condition value.
type CertificateConditionType string

// ConditionStatus represents a condition's status.
// +kubebuilder:validation:Enum=True;False;Unknown
type ConditionStatus string

const (
	// CertificateConditionReady indicates that a certificate is ready for use.
	// This is defined as:
	// - The target secret exists
	// - The target secret contains a certificate that has not expired
	// - The target secret contains a private key valid for the certificate
	// - The commonName and dnsNames attributes match those specified on the Certificate
	CertificateConditionReady CertificateConditionType = "Ready"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Certificate is the Schema for the certificates API
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
// +kubebuilder:printcolumn:name="Secret",type="string",JSONPath=".spec.secretName",description=""
// +kubebuilder:printcolumn:name="Issuer",type="string",JSONPath=".spec.issuerRef.name",description="",priority=1
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message",priority=1
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC."
type Certificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateSpec   `json:"spec,omitempty"`
	Status CertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateList contains a list of Certificate
type CertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Certificate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Certificate{}, &CertificateList{})
}
