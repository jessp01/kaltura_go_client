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

// KalturaDoubleClickDistributionProfileFilter struct for KalturaDoubleClickDistributionProfileFilter
type KalturaDoubleClickDistributionProfileFilter struct {
	KalturaDoubleClickDistributionProfileBaseFilter
}

// NewKalturaDoubleClickDistributionProfileFilter instantiates a new KalturaDoubleClickDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDoubleClickDistributionProfileFilter() *KalturaDoubleClickDistributionProfileFilter {
	this := KalturaDoubleClickDistributionProfileFilter{}
	return &this
}

// NewKalturaDoubleClickDistributionProfileFilterWithDefaults instantiates a new KalturaDoubleClickDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDoubleClickDistributionProfileFilterWithDefaults() *KalturaDoubleClickDistributionProfileFilter {
	this := KalturaDoubleClickDistributionProfileFilter{}
	return &this
}

func (o KalturaDoubleClickDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDoubleClickDistributionProfileBaseFilter, errKalturaDoubleClickDistributionProfileBaseFilter := json.Marshal(o.KalturaDoubleClickDistributionProfileBaseFilter)
	if errKalturaDoubleClickDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaDoubleClickDistributionProfileBaseFilter
	}
	errKalturaDoubleClickDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaDoubleClickDistributionProfileBaseFilter), &toSerialize)
	if errKalturaDoubleClickDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaDoubleClickDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDoubleClickDistributionProfileFilter struct {
	value *KalturaDoubleClickDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaDoubleClickDistributionProfileFilter) Get() *KalturaDoubleClickDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaDoubleClickDistributionProfileFilter) Set(val *KalturaDoubleClickDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDoubleClickDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDoubleClickDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDoubleClickDistributionProfileFilter(val *KalturaDoubleClickDistributionProfileFilter) *NullableKalturaDoubleClickDistributionProfileFilter {
	return &NullableKalturaDoubleClickDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaDoubleClickDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDoubleClickDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


