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

// KalturaCompareCondition `abstract`
type KalturaCompareCondition struct {
	KalturaCondition
}

// NewKalturaCompareCondition instantiates a new KalturaCompareCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCompareCondition() *KalturaCompareCondition {
	this := KalturaCompareCondition{}
	return &this
}

// NewKalturaCompareConditionWithDefaults instantiates a new KalturaCompareCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCompareConditionWithDefaults() *KalturaCompareCondition {
	this := KalturaCompareCondition{}
	return &this
}

func (o KalturaCompareCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaCondition, errKalturaCondition := json.Marshal(o.KalturaCondition)
	if errKalturaCondition != nil {
		return []byte{}, errKalturaCondition
	}
	errKalturaCondition = json.Unmarshal([]byte(serializedKalturaCondition), &toSerialize)
	if errKalturaCondition != nil {
		return []byte{}, errKalturaCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCompareCondition struct {
	value *KalturaCompareCondition
	isSet bool
}

func (v NullableKalturaCompareCondition) Get() *KalturaCompareCondition {
	return v.value
}

func (v *NullableKalturaCompareCondition) Set(val *KalturaCompareCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCompareCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCompareCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCompareCondition(val *KalturaCompareCondition) *NullableKalturaCompareCondition {
	return &NullableKalturaCompareCondition{value: val, isSet: true}
}

func (v NullableKalturaCompareCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCompareCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


