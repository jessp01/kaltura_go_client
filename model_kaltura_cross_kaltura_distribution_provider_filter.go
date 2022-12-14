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

// KalturaCrossKalturaDistributionProviderFilter struct for KalturaCrossKalturaDistributionProviderFilter
type KalturaCrossKalturaDistributionProviderFilter struct {
	KalturaCrossKalturaDistributionProviderBaseFilter
}

// NewKalturaCrossKalturaDistributionProviderFilter instantiates a new KalturaCrossKalturaDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCrossKalturaDistributionProviderFilter() *KalturaCrossKalturaDistributionProviderFilter {
	this := KalturaCrossKalturaDistributionProviderFilter{}
	return &this
}

// NewKalturaCrossKalturaDistributionProviderFilterWithDefaults instantiates a new KalturaCrossKalturaDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCrossKalturaDistributionProviderFilterWithDefaults() *KalturaCrossKalturaDistributionProviderFilter {
	this := KalturaCrossKalturaDistributionProviderFilter{}
	return &this
}

func (o KalturaCrossKalturaDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCrossKalturaDistributionProviderBaseFilter, errKalturaCrossKalturaDistributionProviderBaseFilter := json.Marshal(o.KalturaCrossKalturaDistributionProviderBaseFilter)
	if errKalturaCrossKalturaDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaCrossKalturaDistributionProviderBaseFilter
	}
	errKalturaCrossKalturaDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaCrossKalturaDistributionProviderBaseFilter), &toSerialize)
	if errKalturaCrossKalturaDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaCrossKalturaDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCrossKalturaDistributionProviderFilter struct {
	value *KalturaCrossKalturaDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaCrossKalturaDistributionProviderFilter) Get() *KalturaCrossKalturaDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaCrossKalturaDistributionProviderFilter) Set(val *KalturaCrossKalturaDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCrossKalturaDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCrossKalturaDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCrossKalturaDistributionProviderFilter(val *KalturaCrossKalturaDistributionProviderFilter) *NullableKalturaCrossKalturaDistributionProviderFilter {
	return &NullableKalturaCrossKalturaDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaCrossKalturaDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCrossKalturaDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


