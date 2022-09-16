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

// KalturaGenericDistributionProvider struct for KalturaGenericDistributionProvider
type KalturaGenericDistributionProvider struct {
	KalturaDistributionProvider
}

// NewKalturaGenericDistributionProvider instantiates a new KalturaGenericDistributionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGenericDistributionProvider() *KalturaGenericDistributionProvider {
	this := KalturaGenericDistributionProvider{}
	return &this
}

// NewKalturaGenericDistributionProviderWithDefaults instantiates a new KalturaGenericDistributionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGenericDistributionProviderWithDefaults() *KalturaGenericDistributionProvider {
	this := KalturaGenericDistributionProvider{}
	return &this
}

func (o KalturaGenericDistributionProvider) MarshalJSON() ([]byte, error) {
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

type NullableKalturaGenericDistributionProvider struct {
	value *KalturaGenericDistributionProvider
	isSet bool
}

func (v NullableKalturaGenericDistributionProvider) Get() *KalturaGenericDistributionProvider {
	return v.value
}

func (v *NullableKalturaGenericDistributionProvider) Set(val *KalturaGenericDistributionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGenericDistributionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGenericDistributionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGenericDistributionProvider(val *KalturaGenericDistributionProvider) *NullableKalturaGenericDistributionProvider {
	return &NullableKalturaGenericDistributionProvider{value: val, isSet: true}
}

func (v NullableKalturaGenericDistributionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGenericDistributionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


