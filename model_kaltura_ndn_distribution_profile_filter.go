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

// KalturaNdnDistributionProfileFilter struct for KalturaNdnDistributionProfileFilter
type KalturaNdnDistributionProfileFilter struct {
	KalturaNdnDistributionProfileBaseFilter
}

// NewKalturaNdnDistributionProfileFilter instantiates a new KalturaNdnDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaNdnDistributionProfileFilter() *KalturaNdnDistributionProfileFilter {
	this := KalturaNdnDistributionProfileFilter{}
	return &this
}

// NewKalturaNdnDistributionProfileFilterWithDefaults instantiates a new KalturaNdnDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaNdnDistributionProfileFilterWithDefaults() *KalturaNdnDistributionProfileFilter {
	this := KalturaNdnDistributionProfileFilter{}
	return &this
}

func (o KalturaNdnDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaNdnDistributionProfileBaseFilter, errKalturaNdnDistributionProfileBaseFilter := json.Marshal(o.KalturaNdnDistributionProfileBaseFilter)
	if errKalturaNdnDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaNdnDistributionProfileBaseFilter
	}
	errKalturaNdnDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaNdnDistributionProfileBaseFilter), &toSerialize)
	if errKalturaNdnDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaNdnDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaNdnDistributionProfileFilter struct {
	value *KalturaNdnDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaNdnDistributionProfileFilter) Get() *KalturaNdnDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaNdnDistributionProfileFilter) Set(val *KalturaNdnDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaNdnDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaNdnDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaNdnDistributionProfileFilter(val *KalturaNdnDistributionProfileFilter) *NullableKalturaNdnDistributionProfileFilter {
	return &NullableKalturaNdnDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaNdnDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaNdnDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

