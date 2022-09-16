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

// KalturaQuickPlayDistributionProvider struct for KalturaQuickPlayDistributionProvider
type KalturaQuickPlayDistributionProvider struct {
	KalturaDistributionProvider
}

// NewKalturaQuickPlayDistributionProvider instantiates a new KalturaQuickPlayDistributionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaQuickPlayDistributionProvider() *KalturaQuickPlayDistributionProvider {
	this := KalturaQuickPlayDistributionProvider{}
	return &this
}

// NewKalturaQuickPlayDistributionProviderWithDefaults instantiates a new KalturaQuickPlayDistributionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaQuickPlayDistributionProviderWithDefaults() *KalturaQuickPlayDistributionProvider {
	this := KalturaQuickPlayDistributionProvider{}
	return &this
}

func (o KalturaQuickPlayDistributionProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionProvider, errKalturaDistributionProvider := json.Marshal(o.KalturaDistributionProvider)
	if errKalturaDistributionProvider != nil {
		return []byte{}, errKalturaDistributionProvider
	}
	errKalturaDistributionProvider = json.Unmarshal([]byte(serializedKalturaDistributionProvider), &toSerialize)
	if errKalturaDistributionProvider != nil {
		return []byte{}, errKalturaDistributionProvider
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaQuickPlayDistributionProvider struct {
	value *KalturaQuickPlayDistributionProvider
	isSet bool
}

func (v NullableKalturaQuickPlayDistributionProvider) Get() *KalturaQuickPlayDistributionProvider {
	return v.value
}

func (v *NullableKalturaQuickPlayDistributionProvider) Set(val *KalturaQuickPlayDistributionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaQuickPlayDistributionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaQuickPlayDistributionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaQuickPlayDistributionProvider(val *KalturaQuickPlayDistributionProvider) *NullableKalturaQuickPlayDistributionProvider {
	return &NullableKalturaQuickPlayDistributionProvider{value: val, isSet: true}
}

func (v NullableKalturaQuickPlayDistributionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaQuickPlayDistributionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


