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

// KalturaUverseClickToOrderDistributionProvider struct for KalturaUverseClickToOrderDistributionProvider
type KalturaUverseClickToOrderDistributionProvider struct {
	KalturaDistributionProvider
}

// NewKalturaUverseClickToOrderDistributionProvider instantiates a new KalturaUverseClickToOrderDistributionProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUverseClickToOrderDistributionProvider() *KalturaUverseClickToOrderDistributionProvider {
	this := KalturaUverseClickToOrderDistributionProvider{}
	return &this
}

// NewKalturaUverseClickToOrderDistributionProviderWithDefaults instantiates a new KalturaUverseClickToOrderDistributionProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUverseClickToOrderDistributionProviderWithDefaults() *KalturaUverseClickToOrderDistributionProvider {
	this := KalturaUverseClickToOrderDistributionProvider{}
	return &this
}

func (o KalturaUverseClickToOrderDistributionProvider) MarshalJSON() ([]byte, error) {
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

type NullableKalturaUverseClickToOrderDistributionProvider struct {
	value *KalturaUverseClickToOrderDistributionProvider
	isSet bool
}

func (v NullableKalturaUverseClickToOrderDistributionProvider) Get() *KalturaUverseClickToOrderDistributionProvider {
	return v.value
}

func (v *NullableKalturaUverseClickToOrderDistributionProvider) Set(val *KalturaUverseClickToOrderDistributionProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUverseClickToOrderDistributionProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUverseClickToOrderDistributionProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUverseClickToOrderDistributionProvider(val *KalturaUverseClickToOrderDistributionProvider) *NullableKalturaUverseClickToOrderDistributionProvider {
	return &NullableKalturaUverseClickToOrderDistributionProvider{value: val, isSet: true}
}

func (v NullableKalturaUverseClickToOrderDistributionProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUverseClickToOrderDistributionProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


