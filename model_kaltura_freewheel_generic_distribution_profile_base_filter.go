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

// KalturaFreewheelGenericDistributionProfileBaseFilter `abstract`
type KalturaFreewheelGenericDistributionProfileBaseFilter struct {
	KalturaConfigurableDistributionProfileFilter
}

// NewKalturaFreewheelGenericDistributionProfileBaseFilter instantiates a new KalturaFreewheelGenericDistributionProfileBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFreewheelGenericDistributionProfileBaseFilter() *KalturaFreewheelGenericDistributionProfileBaseFilter {
	this := KalturaFreewheelGenericDistributionProfileBaseFilter{}
	return &this
}

// NewKalturaFreewheelGenericDistributionProfileBaseFilterWithDefaults instantiates a new KalturaFreewheelGenericDistributionProfileBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFreewheelGenericDistributionProfileBaseFilterWithDefaults() *KalturaFreewheelGenericDistributionProfileBaseFilter {
	this := KalturaFreewheelGenericDistributionProfileBaseFilter{}
	return &this
}

func (o KalturaFreewheelGenericDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
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

type NullableKalturaFreewheelGenericDistributionProfileBaseFilter struct {
	value *KalturaFreewheelGenericDistributionProfileBaseFilter
	isSet bool
}

func (v NullableKalturaFreewheelGenericDistributionProfileBaseFilter) Get() *KalturaFreewheelGenericDistributionProfileBaseFilter {
	return v.value
}

func (v *NullableKalturaFreewheelGenericDistributionProfileBaseFilter) Set(val *KalturaFreewheelGenericDistributionProfileBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFreewheelGenericDistributionProfileBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFreewheelGenericDistributionProfileBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFreewheelGenericDistributionProfileBaseFilter(val *KalturaFreewheelGenericDistributionProfileBaseFilter) *NullableKalturaFreewheelGenericDistributionProfileBaseFilter {
	return &NullableKalturaFreewheelGenericDistributionProfileBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaFreewheelGenericDistributionProfileBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFreewheelGenericDistributionProfileBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


