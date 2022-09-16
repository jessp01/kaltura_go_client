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

// KalturaHuluDistributionJobProviderData struct for KalturaHuluDistributionJobProviderData
type KalturaHuluDistributionJobProviderData struct {
	KalturaConfigurableDistributionJobProviderData
}

// NewKalturaHuluDistributionJobProviderData instantiates a new KalturaHuluDistributionJobProviderData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaHuluDistributionJobProviderData() *KalturaHuluDistributionJobProviderData {
	this := KalturaHuluDistributionJobProviderData{}
	return &this
}

// NewKalturaHuluDistributionJobProviderDataWithDefaults instantiates a new KalturaHuluDistributionJobProviderData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaHuluDistributionJobProviderDataWithDefaults() *KalturaHuluDistributionJobProviderData {
	this := KalturaHuluDistributionJobProviderData{}
	return &this
}

func (o KalturaHuluDistributionJobProviderData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaHuluDistributionJobProviderData struct {
	value *KalturaHuluDistributionJobProviderData
	isSet bool
}

func (v NullableKalturaHuluDistributionJobProviderData) Get() *KalturaHuluDistributionJobProviderData {
	return v.value
}

func (v *NullableKalturaHuluDistributionJobProviderData) Set(val *KalturaHuluDistributionJobProviderData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaHuluDistributionJobProviderData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaHuluDistributionJobProviderData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaHuluDistributionJobProviderData(val *KalturaHuluDistributionJobProviderData) *NullableKalturaHuluDistributionJobProviderData {
	return &NullableKalturaHuluDistributionJobProviderData{value: val, isSet: true}
}

func (v NullableKalturaHuluDistributionJobProviderData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaHuluDistributionJobProviderData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


