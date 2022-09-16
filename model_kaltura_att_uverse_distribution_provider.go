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

// KalturaAttUverseDistributionProvider struct for KalturaAttUverseDistributionProvider
type KalturaAttUverseDistributionProvider struct {
	KalturaDistributionProvider
}

// NewKalturaAttUverseDistributionProvider instantiates a new KalturaAttUverseDistributionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttUverseDistributionProvider() *KalturaAttUverseDistributionProvider {
	this := KalturaAttUverseDistributionProvider{}
	return &this
}

// NewKalturaAttUverseDistributionProviderWithDefaults instantiates a new KalturaAttUverseDistributionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttUverseDistributionProviderWithDefaults() *KalturaAttUverseDistributionProvider {
	this := KalturaAttUverseDistributionProvider{}
	return &this
}

func (o KalturaAttUverseDistributionProvider) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAttUverseDistributionProvider struct {
	value *KalturaAttUverseDistributionProvider
	isSet bool
}

func (v NullableKalturaAttUverseDistributionProvider) Get() *KalturaAttUverseDistributionProvider {
	return v.value
}

func (v *NullableKalturaAttUverseDistributionProvider) Set(val *KalturaAttUverseDistributionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttUverseDistributionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttUverseDistributionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttUverseDistributionProvider(val *KalturaAttUverseDistributionProvider) *NullableKalturaAttUverseDistributionProvider {
	return &NullableKalturaAttUverseDistributionProvider{value: val, isSet: true}
}

func (v NullableKalturaAttUverseDistributionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttUverseDistributionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


