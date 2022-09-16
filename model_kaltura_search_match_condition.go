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

// KalturaSearchMatchCondition struct for KalturaSearchMatchCondition
type KalturaSearchMatchCondition struct {
	KalturaSearchCondition
}

// NewKalturaSearchMatchCondition instantiates a new KalturaSearchMatchCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSearchMatchCondition() *KalturaSearchMatchCondition {
	this := KalturaSearchMatchCondition{}
	return &this
}

// NewKalturaSearchMatchConditionWithDefaults instantiates a new KalturaSearchMatchCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSearchMatchConditionWithDefaults() *KalturaSearchMatchCondition {
	this := KalturaSearchMatchCondition{}
	return &this
}

func (o KalturaSearchMatchCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearchCondition, errKalturaSearchCondition := json.Marshal(o.KalturaSearchCondition)
	if errKalturaSearchCondition != nil {
		return []byte{}, errKalturaSearchCondition
	}
	errKalturaSearchCondition = json.Unmarshal([]byte(serializedKalturaSearchCondition), &toSerialize)
	if errKalturaSearchCondition != nil {
		return []byte{}, errKalturaSearchCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSearchMatchCondition struct {
	value *KalturaSearchMatchCondition
	isSet bool
}

func (v NullableKalturaSearchMatchCondition) Get() *KalturaSearchMatchCondition {
	return v.value
}

func (v *NullableKalturaSearchMatchCondition) Set(val *KalturaSearchMatchCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSearchMatchCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSearchMatchCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSearchMatchCondition(val *KalturaSearchMatchCondition) *NullableKalturaSearchMatchCondition {
	return &NullableKalturaSearchMatchCondition{value: val, isSet: true}
}

func (v NullableKalturaSearchMatchCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSearchMatchCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


