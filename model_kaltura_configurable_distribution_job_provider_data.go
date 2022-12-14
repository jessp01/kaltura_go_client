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

// KalturaConfigurableDistributionJobProviderData `abstract`
type KalturaConfigurableDistributionJobProviderData struct {
	KalturaDistributionJobProviderData
}

// NewKalturaConfigurableDistributionJobProviderData instantiates a new KalturaConfigurableDistributionJobProviderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaConfigurableDistributionJobProviderData() *KalturaConfigurableDistributionJobProviderData {
	this := KalturaConfigurableDistributionJobProviderData{}
	return &this
}

// NewKalturaConfigurableDistributionJobProviderDataWithDefaults instantiates a new KalturaConfigurableDistributionJobProviderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaConfigurableDistributionJobProviderDataWithDefaults() *KalturaConfigurableDistributionJobProviderData {
	this := KalturaConfigurableDistributionJobProviderData{}
	return &this
}

func (o KalturaConfigurableDistributionJobProviderData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaConfigurableDistributionJobProviderData struct {
	value *KalturaConfigurableDistributionJobProviderData
	isSet bool
}

func (v NullableKalturaConfigurableDistributionJobProviderData) Get() *KalturaConfigurableDistributionJobProviderData {
	return v.value
}

func (v *NullableKalturaConfigurableDistributionJobProviderData) Set(val *KalturaConfigurableDistributionJobProviderData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaConfigurableDistributionJobProviderData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaConfigurableDistributionJobProviderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaConfigurableDistributionJobProviderData(val *KalturaConfigurableDistributionJobProviderData) *NullableKalturaConfigurableDistributionJobProviderData {
	return &NullableKalturaConfigurableDistributionJobProviderData{value: val, isSet: true}
}

func (v NullableKalturaConfigurableDistributionJobProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaConfigurableDistributionJobProviderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


