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

// KalturaVerizonVcastDistributionProfileBaseFilter `abstract`
type KalturaVerizonVcastDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaVerizonVcastDistributionProfileBaseFilter instantiates a new KalturaVerizonVcastDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaVerizonVcastDistributionProfileBaseFilter() *KalturaVerizonVcastDistributionProfileBaseFilter {
	this := KalturaVerizonVcastDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaVerizonVcastDistributionProfileBaseFilterWithDefaults instantiates a new KalturaVerizonVcastDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaVerizonVcastDistributionProfileBaseFilterWithDefaults() *KalturaVerizonVcastDistributionProfileBaseFilter {
	this := KalturaVerizonVcastDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaVerizonVcastDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaConfigurableDistributionProfileFilter, errKalturaConfigurableDistributionProfileFilter := json.Marshal(o.KalturaConfigurableDistributionProfileFilter)
	if errKalturaConfigurableDistributionProfileFilter != nil {
		return []byte{}, errKalturaConfigurableDistributionProfileFilter
	}
	errKalturaConfigurableDistributionProfileFilter = json.Unmarshal([]byte(serializedKalturaConfigurableDistributionProfileFilter), &toSerialize)
	if errKalturaConfigurableDistributionProfileFilter != nil {
		return []byte{}, errKalturaConfigurableDistributionProfileFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaVerizonVcastDistributionProfileBaseFilter struct {
	value *KalturaVerizonVcastDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaVerizonVcastDistributionProfileBaseFilter) Get() *KalturaVerizonVcastDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaVerizonVcastDistributionProfileBaseFilter) Set(val *KalturaVerizonVcastDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaVerizonVcastDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaVerizonVcastDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaVerizonVcastDistributionProfileBaseFilter(val *KalturaVerizonVcastDistributionProfileBaseFilter) *NullableKalturaVerizonVcastDistributionProfileBaseFilter {
	return &NullableKalturaVerizonVcastDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaVerizonVcastDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaVerizonVcastDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

