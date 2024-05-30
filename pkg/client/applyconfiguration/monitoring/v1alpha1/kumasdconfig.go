// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/client/applyconfiguration/monitoring/v1"
	corev1 "k8s.io/api/core/v1"
)

// KumaSDConfigApplyConfiguration represents an declarative configuration of the KumaSDConfig type for use
// with apply.
type KumaSDConfigApplyConfiguration struct {
	Server                                     *string      `json:"server,omitempty"`
	ClientID                                   *string      `json:"clientID,omitempty"`
	RefreshInterval                            *v1.Duration `json:"refreshInterval,omitempty"`
	FetchTimeout                               *v1.Duration `json:"fetchTimeout,omitempty"`
	monitoringv1.ProxyConfigApplyConfiguration `json:",inline"`
	TLSConfig                                  *monitoringv1.SafeTLSConfigApplyConfiguration     `json:"tlsConfig,omitempty"`
	BasicAuth                                  *monitoringv1.BasicAuthApplyConfiguration         `json:"basicAuth,omitempty"`
	Authorization                              *monitoringv1.SafeAuthorizationApplyConfiguration `json:"authorization,omitempty"`
	OAuth2                                     *monitoringv1.OAuth2ApplyConfiguration            `json:"oauth2,omitempty"`
	FollowRedirects                            *bool                                             `json:"followRedirects,omitempty"`
	EnableHTTP2                                *bool                                             `json:"enableHTTP2,omitempty"`
}

// KumaSDConfigApplyConfiguration constructs an declarative configuration of the KumaSDConfig type for use with
// apply.
func KumaSDConfig() *KumaSDConfigApplyConfiguration {
	return &KumaSDConfigApplyConfiguration{}
}

// WithServer sets the Server field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Server field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithServer(value string) *KumaSDConfigApplyConfiguration {
	b.Server = &value
	return b
}

// WithClientID sets the ClientID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientID field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithClientID(value string) *KumaSDConfigApplyConfiguration {
	b.ClientID = &value
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithRefreshInterval(value v1.Duration) *KumaSDConfigApplyConfiguration {
	b.RefreshInterval = &value
	return b
}

// WithFetchTimeout sets the FetchTimeout field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FetchTimeout field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithFetchTimeout(value v1.Duration) *KumaSDConfigApplyConfiguration {
	b.FetchTimeout = &value
	return b
}

// WithProxyURL sets the ProxyURL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyURL field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithProxyURL(value string) *KumaSDConfigApplyConfiguration {
	b.ProxyURL = &value
	return b
}

// WithNoProxy sets the NoProxy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NoProxy field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithNoProxy(value string) *KumaSDConfigApplyConfiguration {
	b.NoProxy = &value
	return b
}

// WithProxyFromEnvironment sets the ProxyFromEnvironment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProxyFromEnvironment field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithProxyFromEnvironment(value bool) *KumaSDConfigApplyConfiguration {
	b.ProxyFromEnvironment = &value
	return b
}

// WithProxyConnectHeader puts the entries into the ProxyConnectHeader field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the ProxyConnectHeader field,
// overwriting an existing map entries in ProxyConnectHeader field with the same key.
func (b *KumaSDConfigApplyConfiguration) WithProxyConnectHeader(entries map[string][]corev1.SecretKeySelector) *KumaSDConfigApplyConfiguration {
	if b.ProxyConnectHeader == nil && len(entries) > 0 {
		b.ProxyConnectHeader = make(map[string][]corev1.SecretKeySelector, len(entries))
	}
	for k, v := range entries {
		b.ProxyConnectHeader[k] = v
	}
	return b
}

// WithTLSConfig sets the TLSConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TLSConfig field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithTLSConfig(value *monitoringv1.SafeTLSConfigApplyConfiguration) *KumaSDConfigApplyConfiguration {
	b.TLSConfig = value
	return b
}

// WithBasicAuth sets the BasicAuth field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BasicAuth field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithBasicAuth(value *monitoringv1.BasicAuthApplyConfiguration) *KumaSDConfigApplyConfiguration {
	b.BasicAuth = value
	return b
}

// WithAuthorization sets the Authorization field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Authorization field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithAuthorization(value *monitoringv1.SafeAuthorizationApplyConfiguration) *KumaSDConfigApplyConfiguration {
	b.Authorization = value
	return b
}

// WithOAuth2 sets the OAuth2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OAuth2 field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithOAuth2(value *monitoringv1.OAuth2ApplyConfiguration) *KumaSDConfigApplyConfiguration {
	b.OAuth2 = value
	return b
}

// WithFollowRedirects sets the FollowRedirects field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FollowRedirects field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithFollowRedirects(value bool) *KumaSDConfigApplyConfiguration {
	b.FollowRedirects = &value
	return b
}

// WithEnableHTTP2 sets the EnableHTTP2 field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableHTTP2 field is set to the value of the last call.
func (b *KumaSDConfigApplyConfiguration) WithEnableHTTP2(value bool) *KumaSDConfigApplyConfiguration {
	b.EnableHTTP2 = &value
	return b
}
