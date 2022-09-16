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

// KalturaFtpDistributionProfileFilter struct for KalturaFtpDistributionProfileFilter
type KalturaFtpDistributionProfileFilter struct {
	KalturaFtpDistributionProfileBaseFilter
}

// NewKalturaFtpDistributionProfileFilter instantiates a new KalturaFtpDistributionProfileFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFtpDistributionProfileFilter() *KalturaFtpDistributionProfileFilter {
	this := KalturaFtpDistributionProfileFilter{}
	return &this
}

// NewKalturaFtpDistributionProfileFilterWithDefaults instantiates a new KalturaFtpDistributionProfileFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFtpDistributionProfileFilterWithDefaults() *KalturaFtpDistributionProfileFilter {
	this := KalturaFtpDistributionProfileFilter{}
	return &this
}

func (o KalturaFtpDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFtpDistributionProfileBaseFilter, errKalturaFtpDistributionProfileBaseFilter := json.Marshal(o.KalturaFtpDistributionProfileBaseFilter)
	if errKalturaFtpDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaFtpDistributionProfileBaseFilter
	}
	errKalturaFtpDistributionProfileBaseFilter = json.Unmarshal([]byte(serializedKalturaFtpDistributionProfileBaseFilter), &toSerialize)
	if errKalturaFtpDistributionProfileBaseFilter != nil {
		return []byte{}, errKalturaFtpDistributionProfileBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFtpDistributionProfileFilter struct {
	value *KalturaFtpDistributionProfileFilter
	isSet bool
}

func (v NullableKalturaFtpDistributionProfileFilter) Get() *KalturaFtpDistributionProfileFilter {
	return v.value
}

func (v *NullableKalturaFtpDistributionProfileFilter) Set(val *KalturaFtpDistributionProfileFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFtpDistributionProfileFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFtpDistributionProfileFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFtpDistributionProfileFilter(val *KalturaFtpDistributionProfileFilter) *NullableKalturaFtpDistributionProfileFilter {
	return &NullableKalturaFtpDistributionProfileFilter{value: val, isSet: true}
}

func (v NullableKalturaFtpDistributionProfileFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFtpDistributionProfileFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

