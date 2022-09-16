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

// KalturaGenericDistributionJobProviderData struct for KalturaGenericDistributionJobProviderData
type KalturaGenericDistributionJobProviderData struct {
	KalturaDistributionJobProviderData
}

// NewKalturaGenericDistributionJobProviderData instantiates a new KalturaGenericDistributionJobProviderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaGenericDistributionJobProviderData() *KalturaGenericDistributionJobProviderData {
	this := KalturaGenericDistributionJobProviderData{}
	return &this
}

// NewKalturaGenericDistributionJobProviderDataWithDefaults instantiates a new KalturaGenericDistributionJobProviderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaGenericDistributionJobProviderDataWithDefaults() *KalturaGenericDistributionJobProviderData {
	this := KalturaGenericDistributionJobProviderData{}
	return &this
}

func (o KalturaGenericDistributionJobProviderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionJobProviderData, errKalturaDistributionJobProviderData := json.Marshal(o.KalturaDistributionJobProviderData)
	if errKalturaDistributionJobProviderData != nil {
		return []byte{}, errKalturaDistributionJobProviderData
	}
	errKalturaDistributionJobProviderData = json.Unmarshal([]byte(serializedKalturaDistributionJobProviderData), &toSerialize)
	if errKalturaDistributionJobProviderData != nil {
		return []byte{}, errKalturaDistributionJobProviderData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaGenericDistributionJobProviderData struct {
	value *KalturaGenericDistributionJobProviderData
	isSet bool
}

func (v NullableKalturaGenericDistributionJobProviderData) Get() *KalturaGenericDistributionJobProviderData {
	return v.value
}

func (v *NullableKalturaGenericDistributionJobProviderData) Set(val *KalturaGenericDistributionJobProviderData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaGenericDistributionJobProviderData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaGenericDistributionJobProviderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaGenericDistributionJobProviderData(val *KalturaGenericDistributionJobProviderData) *NullableKalturaGenericDistributionJobProviderData {
	return &NullableKalturaGenericDistributionJobProviderData{value: val, isSet: true}
}

func (v NullableKalturaGenericDistributionJobProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaGenericDistributionJobProviderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


