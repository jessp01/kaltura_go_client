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

// KalturaFreewheelDistributionProviderFilter struct for KalturaFreewheelDistributionProviderFilter
type KalturaFreewheelDistributionProviderFilter struct {
	KalturaFreewheelDistributionProviderBaseFilter
}

// NewKalturaFreewheelDistributionProviderFilter instantiates a new KalturaFreewheelDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFreewheelDistributionProviderFilter() *KalturaFreewheelDistributionProviderFilter {
	this := KalturaFreewheelDistributionProviderFilter{}
	return &this
}

// NewKalturaFreewheelDistributionProviderFilterWithDefaults instantiates a new KalturaFreewheelDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFreewheelDistributionProviderFilterWithDefaults() *KalturaFreewheelDistributionProviderFilter {
	this := KalturaFreewheelDistributionProviderFilter{}
	return &this
}

func (o KalturaFreewheelDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFreewheelDistributionProviderBaseFilter, errKalturaFreewheelDistributionProviderBaseFilter := json.Marshal(o.KalturaFreewheelDistributionProviderBaseFilter)
	if errKalturaFreewheelDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaFreewheelDistributionProviderBaseFilter
	}
	errKalturaFreewheelDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaFreewheelDistributionProviderBaseFilter), &toSerialize)
	if errKalturaFreewheelDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaFreewheelDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFreewheelDistributionProviderFilter struct {
	value *KalturaFreewheelDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaFreewheelDistributionProviderFilter) Get() *KalturaFreewheelDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaFreewheelDistributionProviderFilter) Set(val *KalturaFreewheelDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFreewheelDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFreewheelDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFreewheelDistributionProviderFilter(val *KalturaFreewheelDistributionProviderFilter) *NullableKalturaFreewheelDistributionProviderFilter {
	return &NullableKalturaFreewheelDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaFreewheelDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFreewheelDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


