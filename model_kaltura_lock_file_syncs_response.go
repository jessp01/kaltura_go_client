/*
Kaltura VPaaS

The Kaltura VPaaS API

API version: 18.12.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// KalturaLockFileSyncsResponse struct for KalturaLockFileSyncsResponse
type KalturaLockFileSyncsResponse struct {
	BaseUrl *string `json:"baseUrl,omitempty"`
	DcSecret *string `json:"dcSecret,omitempty"`
	FileSyncs []KalturaFileSync `json:"fileSyncs,omitempty"`
	LimitReached *bool `json:"limitReached,omitempty"`
}

// NewKalturaLockFileSyncsResponse instantiates a new KalturaLockFileSyncsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLockFileSyncsResponse() *KalturaLockFileSyncsResponse {
	this := KalturaLockFileSyncsResponse{}
	return &this
}

// NewKalturaLockFileSyncsResponseWithDefaults instantiates a new KalturaLockFileSyncsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLockFileSyncsResponseWithDefaults() *KalturaLockFileSyncsResponse {
	this := KalturaLockFileSyncsResponse{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *KalturaLockFileSyncsResponse) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaLockFileSyncsResponse) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *KalturaLockFileSyncsResponse) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *KalturaLockFileSyncsResponse) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetDcSecret returns the DcSecret field value if set, zero value otherwise.
func (o *KalturaLockFileSyncsResponse) GetDcSecret() string {
	if o == nil || o.DcSecret == nil {
		var ret string
		return ret
	}
	return *o.DcSecret
}

// GetDcSecretOk returns a tuple with the DcSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaLockFileSyncsResponse) GetDcSecretOk() (*string, bool) {
	if o == nil || o.DcSecret == nil {
		return nil, false
	}
	return o.DcSecret, true
}

// HasDcSecret returns a boolean if a field has been set.
func (o *KalturaLockFileSyncsResponse) HasDcSecret() bool {
	if o != nil && o.DcSecret != nil {
		return true
	}

	return false
}

// SetDcSecret gets a reference to the given string and assigns it to the DcSecret field.
func (o *KalturaLockFileSyncsResponse) SetDcSecret(v string) {
	o.DcSecret = &v
}

// GetFileSyncs returns the FileSyncs field value if set, zero value otherwise.
func (o *KalturaLockFileSyncsResponse) GetFileSyncs() []KalturaFileSync {
	if o == nil || o.FileSyncs == nil {
		var ret []KalturaFileSync
		return ret
	}
	return o.FileSyncs
}

// GetFileSyncsOk returns a tuple with the FileSyncs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaLockFileSyncsResponse) GetFileSyncsOk() ([]KalturaFileSync, bool) {
	if o == nil || o.FileSyncs == nil {
		return nil, false
	}
	return o.FileSyncs, true
}

// HasFileSyncs returns a boolean if a field has been set.
func (o *KalturaLockFileSyncsResponse) HasFileSyncs() bool {
	if o != nil && o.FileSyncs != nil {
		return true
	}

	return false
}

// SetFileSyncs gets a reference to the given []KalturaFileSync and assigns it to the FileSyncs field.
func (o *KalturaLockFileSyncsResponse) SetFileSyncs(v []KalturaFileSync) {
	o.FileSyncs = v
}

// GetLimitReached returns the LimitReached field value if set, zero value otherwise.
func (o *KalturaLockFileSyncsResponse) GetLimitReached() bool {
	if o == nil || o.LimitReached == nil {
		var ret bool
		return ret
	}
	return *o.LimitReached
}

// GetLimitReachedOk returns a tuple with the LimitReached field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaLockFileSyncsResponse) GetLimitReachedOk() (*bool, bool) {
	if o == nil || o.LimitReached == nil {
		return nil, false
	}
	return o.LimitReached, true
}

// HasLimitReached returns a boolean if a field has been set.
func (o *KalturaLockFileSyncsResponse) HasLimitReached() bool {
	if o != nil && o.LimitReached != nil {
		return true
	}

	return false
}

// SetLimitReached gets a reference to the given bool and assigns it to the LimitReached field.
func (o *KalturaLockFileSyncsResponse) SetLimitReached(v bool) {
	o.LimitReached = &v
}

func (o KalturaLockFileSyncsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseUrl != nil {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if o.DcSecret != nil {
		toSerialize["dcSecret"] = o.DcSecret
	}
	if o.FileSyncs != nil {
		toSerialize["fileSyncs"] = o.FileSyncs
	}
	if o.LimitReached != nil {
		toSerialize["limitReached"] = o.LimitReached
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLockFileSyncsResponse struct {
	value *KalturaLockFileSyncsResponse
	isSet bool
}

func (v NullableKalturaLockFileSyncsResponse) Get() *KalturaLockFileSyncsResponse {
	return v.value
}

func (v *NullableKalturaLockFileSyncsResponse) Set(val *KalturaLockFileSyncsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLockFileSyncsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLockFileSyncsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLockFileSyncsResponse(val *KalturaLockFileSyncsResponse) *NullableKalturaLockFileSyncsResponse {
	return &NullableKalturaLockFileSyncsResponse{value: val, isSet: true}
}

func (v NullableKalturaLockFileSyncsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLockFileSyncsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


