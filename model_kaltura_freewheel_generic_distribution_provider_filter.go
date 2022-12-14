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

// KalturaFreewheelGenericDistributionProviderFilter struct for KalturaFreewheelGenericDistributionProviderFilter
type KalturaFreewheelGenericDistributionProviderFilter struct {
	KalturaFreewheelGenericDistributionProviderBaseFilter
}

// NewKalturaFreewheelGenericDistributionProviderFilter instantiates a new KalturaFreewheelGenericDistributionProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFreewheelGenericDistributionProviderFilter() *KalturaFreewheelGenericDistributionProviderFilter {
	this := KalturaFreewheelGenericDistributionProviderFilter{}
	return &this
}

// NewKalturaFreewheelGenericDistributionProviderFilterWithDefaults instantiates a new KalturaFreewheelGenericDistributionProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFreewheelGenericDistributionProviderFilterWithDefaults() *KalturaFreewheelGenericDistributionProviderFilter {
	this := KalturaFreewheelGenericDistributionProviderFilter{}
	return &this
}

func (o KalturaFreewheelGenericDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaFreewheelGenericDistributionProviderBaseFilter, errKalturaFreewheelGenericDistributionProviderBaseFilter := json.Marshal(o.KalturaFreewheelGenericDistributionProviderBaseFilter)
	if errKalturaFreewheelGenericDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaFreewheelGenericDistributionProviderBaseFilter
	}
	errKalturaFreewheelGenericDistributionProviderBaseFilter = json.Unmarshal([]byte(serializedKalturaFreewheelGenericDistributionProviderBaseFilter), &toSerialize)
	if errKalturaFreewheelGenericDistributionProviderBaseFilter != nil {
		return []byte{}, errKalturaFreewheelGenericDistributionProviderBaseFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFreewheelGenericDistributionProviderFilter struct {
	value *KalturaFreewheelGenericDistributionProviderFilter
	isSet bool
}

func (v NullableKalturaFreewheelGenericDistributionProviderFilter) Get() *KalturaFreewheelGenericDistributionProviderFilter {
	return v.value
}

func (v *NullableKalturaFreewheelGenericDistributionProviderFilter) Set(val *KalturaFreewheelGenericDistributionProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFreewheelGenericDistributionProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFreewheelGenericDistributionProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFreewheelGenericDistributionProviderFilter(val *KalturaFreewheelGenericDistributionProviderFilter) *NullableKalturaFreewheelGenericDistributionProviderFilter {
	return &NullableKalturaFreewheelGenericDistributionProviderFilter{value: val, isSet: true}
}

func (v NullableKalturaFreewheelGenericDistributionProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFreewheelGenericDistributionProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


