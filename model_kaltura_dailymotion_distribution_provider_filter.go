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

// KalturaDailymotionDistributionProviderFilter struct for KalturaDailymotionDistributionProviderFilter
type KalturaDailymotionDistributionProviderFilter struct {
	KalturaDailymotionDistributionProviderBaseFilter
}

// NewKalturaDailymotionDistributionProviderFilter instantiates a new KalturaDailymotionDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDailymotionDistributionProviderFilter() *KalturaDailymotionDistributionProviderFilter {
	this := KalturaDailymotionDistributionProviderFilter{}
	return &this
}

// NewKalturaDailymotionDistributionProviderFilterWithDefaults instantiates a new KalturaDailymotionDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDailymotionDistributionProviderFilterWithDefaults() *KalturaDailymotionDistributionProviderFilter {
	this := KalturaDailymotionDistributionProviderFilter{}
	return &this
}

func (o KalturaDailymotionDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDailymotionDistributionProviderBaseFilter, errKalturaDailymotionDistributionProviderBaseFilter := json.Marshal(o.KalturaDailymotionDistributionProviderBaseFilter)
	if errKalturaDailymotionDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaDailymotionDistributionProviderBaseFilter
	}
	errKalturaDailymotionDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaDailymotionDistributionProviderBaseFilter), &toSerialize)
	if errKalturaDailymotionDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaDailymotionDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDailymotionDistributionProviderFilter struct {
	value *KalturaDailymotionDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaDailymotionDistributionProviderFilter) Get() *KalturaDailymotionDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaDailymotionDistributionProviderFilter) Set(val *KalturaDailymotionDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDailymotionDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDailymotionDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDailymotionDistributionProviderFilter(val *KalturaDailymotionDistributionProviderFilter) *NullableKalturaDailymotionDistributionProviderFilter {
	return &NullableKalturaDailymotionDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaDailymotionDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDailymotionDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


