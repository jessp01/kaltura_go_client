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

// KalturaDailymotionDistributionProfileFilter struct for KalturaDailymotionDistributionProfileFilter
type KalturaDailymotionDistributionProfileFilter struct {
	KalturaDailymotionDistributionProfileBaseFilter
}

// NewKalturaDailymotionDistributionProfileFilter instantiates a new KalturaDailymotionDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDailymotionDistributionProfileFilter() *KalturaDailymotionDistributionProfileFilter {
	this := KalturaDailymotionDistributionProfileFilter{}
	return &this
}

// NewKalturaDailymotionDistributionProfileFilterWithDefaults instantiates a new KalturaDailymotionDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDailymotionDistributionProfileFilterWithDefaults() *KalturaDailymotionDistributionProfileFilter {
	this := KalturaDailymotionDistributionProfileFilter{}
	return &this
}

func (o KalturaDailymotionDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDailymotionDistributionProfileBaseFilter, errKalturaDailymotionDistributionProfileBaseFilter := json.Marshal(o.KalturaDailymotionDistributionProfileBaseFilter)
	if errKalturaDailymotionDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaDailymotionDistributionProfileBaseFilter
	}
	errKalturaDailymotionDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaDailymotionDistributionProfileBaseFilter), &toSerialize)
	if errKalturaDailymotionDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaDailymotionDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDailymotionDistributionProfileFilter struct {
	value *KalturaDailymotionDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaDailymotionDistributionProfileFilter) Get() *KalturaDailymotionDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaDailymotionDistributionProfileFilter) Set(val *KalturaDailymotionDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDailymotionDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDailymotionDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDailymotionDistributionProfileFilter(val *KalturaDailymotionDistributionProfileFilter) *NullableKalturaDailymotionDistributionProfileFilter {
	return &NullableKalturaDailymotionDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaDailymotionDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDailymotionDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

