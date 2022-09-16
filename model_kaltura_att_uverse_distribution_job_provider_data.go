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

// KalturaAttUverseDistributionJobProviderData struct for KalturaAttUverseDistributionJobProviderData
type KalturaAttUverseDistributionJobProviderData struct {
	KalturaConfigurableDistributionJobProviderData
}

// NewKalturaAttUverseDistributionJobProviderData instantiates a new KalturaAttUverseDistributionJobProviderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAttUverseDistributionJobProviderData() *KalturaAttUverseDistributionJobProviderData {
	this := KalturaAttUverseDistributionJobProviderData{}
	return &this
}

// NewKalturaAttUverseDistributionJobProviderDataWithDefaults instantiates a new KalturaAttUverseDistributionJobProviderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAttUverseDistributionJobProviderDataWithDefaults() *KalturaAttUverseDistributionJobProviderData {
	this := KalturaAttUverseDistributionJobProviderData{}
	return &this
}

func (o KalturaAttUverseDistributionJobProviderData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConfigurableDistributionJobProviderData, errKalturaConfigurableDistributionJobProviderData := json.Marshal(o.KalturaConfigurableDistributionJobProviderData)
	if errKalturaConfigurableDistributionJobProviderData != nil {
		return []byte{}, errKalturaConfigurableDistributionJobProviderData
	}
	errKalturaConfigurableDistributionJobProviderData = json.Unmarshal([]byte(serializedKalturaConfigurableDistributionJobProviderData), &toSerialize)
	if errKalturaConfigurableDistributionJobProviderData != nil {
		return []byte{}, errKalturaConfigurableDistributionJobProviderData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAttUverseDistributionJobProviderData struct {
	value *KalturaAttUverseDistributionJobProviderData
	isSet bool
}

func (v NullableKalturaAttUverseDistributionJobProviderData) Get() *KalturaAttUverseDistributionJobProviderData {
	return v.value
}

func (v *NullableKalturaAttUverseDistributionJobProviderData) Set(val *KalturaAttUverseDistributionJobProviderData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAttUverseDistributionJobProviderData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAttUverseDistributionJobProviderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAttUverseDistributionJobProviderData(val *KalturaAttUverseDistributionJobProviderData) *NullableKalturaAttUverseDistributionJobProviderData {
	return &NullableKalturaAttUverseDistributionJobProviderData{value: val, isSet: true}
}

func (v NullableKalturaAttUverseDistributionJobProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAttUverseDistributionJobProviderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

