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

// KalturaFacebookDistributionProviderBaseFilter `abstract`
type KalturaFacebookDistributionProviderBaseFilter struct {
	KalturaDistributionProviderFilter
}

// NewKalturaFacebookDistributionProviderBaseFilter instantiates a new KalturaFacebookDistributionProviderBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFacebookDistributionProviderBaseFilter() *KalturaFacebookDistributionProviderBaseFilter {
	this := KalturaFacebookDistributionProviderBaseFilter{}
	return &this
}

// NewKalturaFacebookDistributionProviderBaseFilterWithDefaults instantiates a new KalturaFacebookDistributionProviderBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFacebookDistributionProviderBaseFilterWithDefaults() *KalturaFacebookDistributionProviderBaseFilter {
	this := KalturaFacebookDistributionProviderBaseFilter{}
	return &this
}

func (o KalturaFacebookDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaFacebookDistributionProviderBaseFilter struct {
	value *KalturaFacebookDistributionProviderBaseFilter
	isSet bool
}

func (v NullableKalturaFacebookDistributionProviderBaseFilter) Get() *KalturaFacebookDistributionProviderBaseFilter {
	return v.value
}

func (v *NullableKalturaFacebookDistributionProviderBaseFilter) Set(val *KalturaFacebookDistributionProviderBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFacebookDistributionProviderBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFacebookDistributionProviderBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFacebookDistributionProviderBaseFilter(val *KalturaFacebookDistributionProviderBaseFilter) *NullableKalturaFacebookDistributionProviderBaseFilter {
	return &NullableKalturaFacebookDistributionProviderBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaFacebookDistributionProviderBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFacebookDistributionProviderBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


