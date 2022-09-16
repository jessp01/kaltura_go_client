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

// KalturaMatchCondition `abstract`
type KalturaMatchCondition struct {
	KalturaCondition
}

// NewKalturaMatchCondition instantiates a new KalturaMatchCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMatchCondition() *KalturaMatchCondition {
	this := KalturaMatchCondition{}
	return &this
}

// NewKalturaMatchConditionWithDefaults instantiates a new KalturaMatchCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMatchConditionWithDefaults() *KalturaMatchCondition {
	this := KalturaMatchCondition{}
	return &this
}

func (o KalturaMatchCondition) MarshalJSON() ([]byte, error) {
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

type NullableKalturaMatchCondition struct {
	value *KalturaMatchCondition
	isSet bool
}

func (v NullableKalturaMatchCondition) Get() *KalturaMatchCondition {
	return v.value
}

func (v *NullableKalturaMatchCondition) Set(val *KalturaMatchCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMatchCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMatchCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMatchCondition(val *KalturaMatchCondition) *NullableKalturaMatchCondition {
	return &NullableKalturaMatchCondition{value: val, isSet: true}
}

func (v NullableKalturaMatchCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMatchCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

