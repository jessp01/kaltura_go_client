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

// KalturaTVComDistributionProviderBaseFilter `abstract`
type KalturaTVComDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaTVComDistributionProviderBaseFilter instantiates a new KalturaTVComDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTVComDistributionProviderBaseFilter() *KalturaTVComDistributionProviderBaseFilter {
	this := KalturaTVComDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaTVComDistributionProviderBaseFilterWithDefaults instantiates a new KalturaTVComDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTVComDistributionProviderBaseFilterWithDefaults() *KalturaTVComDistributionProviderBaseFilter {
	this := KalturaTVComDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaTVComDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaTVComDistributionProviderBaseFilter struct {
	value *KalturaTVComDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaTVComDistributionProviderBaseFilter) Get() *KalturaTVComDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaTVComDistributionProviderBaseFilter) Set(val *KalturaTVComDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTVComDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTVComDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTVComDistributionProviderBaseFilter(val *KalturaTVComDistributionProviderBaseFilter) *NullableKalturaTVComDistributionProviderBaseFilter {
	return &NullableKalturaTVComDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaTVComDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTVComDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


