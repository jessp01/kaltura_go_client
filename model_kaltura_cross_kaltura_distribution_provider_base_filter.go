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

// KalturaCrossKalturaDistributionProviderBaseFilter `abstract`
type KalturaCrossKalturaDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaCrossKalturaDistributionProviderBaseFilter instantiates a new KalturaCrossKalturaDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCrossKalturaDistributionProviderBaseFilter() *KalturaCrossKalturaDistributionProviderBaseFilter {
	this := KalturaCrossKalturaDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaCrossKalturaDistributionProviderBaseFilterWithDefaults instantiates a new KalturaCrossKalturaDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCrossKalturaDistributionProviderBaseFilterWithDefaults() *KalturaCrossKalturaDistributionProviderBaseFilter {
	this := KalturaCrossKalturaDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaCrossKalturaDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDistributionProviderFilter, errKalturaDistributionProviderFilter := json.Marshal(o.KalturaDistributionProviderFilter)
	if errKalturaDistributionProviderFilter != nil {
		return []byte{}, errKalturaDistributionProviderFilter
	}
	errKalturaDistributionProviderFilter = json.Unmarshal([]byte(serializedKalturaDistributionProviderFilter), &toSerialize)
	if errKalturaDistributionProviderFilter != nil {
		return []byte{}, errKalturaDistributionProviderFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCrossKalturaDistributionProviderBaseFilter struct {
	value *KalturaCrossKalturaDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaCrossKalturaDistributionProviderBaseFilter) Get() *KalturaCrossKalturaDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaCrossKalturaDistributionProviderBaseFilter) Set(val *KalturaCrossKalturaDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCrossKalturaDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCrossKalturaDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCrossKalturaDistributionProviderBaseFilter(val *KalturaCrossKalturaDistributionProviderBaseFilter) *NullableKalturaCrossKalturaDistributionProviderBaseFilter {
	return &NullableKalturaCrossKalturaDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaCrossKalturaDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCrossKalturaDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


