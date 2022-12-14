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

// KalturaDoubleClickDistributionProfileBaseFilter `abstract`
type KalturaDoubleClickDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaDoubleClickDistributionProfileBaseFilter instantiates a new KalturaDoubleClickDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDoubleClickDistributionProfileBaseFilter() *KalturaDoubleClickDistributionProfileBaseFilter {
	this := KalturaDoubleClickDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaDoubleClickDistributionProfileBaseFilterWithDefaults instantiates a new KalturaDoubleClickDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDoubleClickDistributionProfileBaseFilterWithDefaults() *KalturaDoubleClickDistributionProfileBaseFilter {
	this := KalturaDoubleClickDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaDoubleClickDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaDoubleClickDistributionProfileBaseFilter struct {
	value *KalturaDoubleClickDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaDoubleClickDistributionProfileBaseFilter) Get() *KalturaDoubleClickDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaDoubleClickDistributionProfileBaseFilter) Set(val *KalturaDoubleClickDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDoubleClickDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDoubleClickDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDoubleClickDistributionProfileBaseFilter(val *KalturaDoubleClickDistributionProfileBaseFilter) *NullableKalturaDoubleClickDistributionProfileBaseFilter {
	return &NullableKalturaDoubleClickDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaDoubleClickDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDoubleClickDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


