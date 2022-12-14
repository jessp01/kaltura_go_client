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

// KalturaSynacorHboDistributionProfileBaseFilter `abstract`
type KalturaSynacorHboDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaSynacorHboDistributionProfileBaseFilter instantiates a new KalturaSynacorHboDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSynacorHboDistributionProfileBaseFilter() *KalturaSynacorHboDistributionProfileBaseFilter {
	this := KalturaSynacorHboDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaSynacorHboDistributionProfileBaseFilterWithDefaults instantiates a new KalturaSynacorHboDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSynacorHboDistributionProfileBaseFilterWithDefaults() *KalturaSynacorHboDistributionProfileBaseFilter {
	this := KalturaSynacorHboDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaSynacorHboDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaSynacorHboDistributionProfileBaseFilter struct {
	value *KalturaSynacorHboDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaSynacorHboDistributionProfileBaseFilter) Get() *KalturaSynacorHboDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaSynacorHboDistributionProfileBaseFilter) Set(val *KalturaSynacorHboDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSynacorHboDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSynacorHboDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSynacorHboDistributionProfileBaseFilter(val *KalturaSynacorHboDistributionProfileBaseFilter) *NullableKalturaSynacorHboDistributionProfileBaseFilter {
	return &NullableKalturaSynacorHboDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaSynacorHboDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSynacorHboDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


