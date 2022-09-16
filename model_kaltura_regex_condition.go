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

// KalturaRegexCondition `abstract`
type KalturaRegexCondition struct {
	KalturaMatchCondition
}

// NewKalturaRegexCondition instantiates a new KalturaRegexCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaRegexCondition() *KalturaRegexCondition {
	this := KalturaRegexCondition{}
	return &this
}

// NewKalturaRegexConditionWithDefaults instantiates a new KalturaRegexCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaRegexConditionWithDefaults() *KalturaRegexCondition {
	this := KalturaRegexCondition{}
	return &this
}

func (o KalturaRegexCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaMatchCondition, errKalturaMatchCondition := json.Marshal(o.KalturaMatchCondition)
	if errKalturaMatchCondition != nil {
		return []byte{}, errKalturaMatchCondition
	}
	errKalturaMatchCondition = json.Unmarshal([]byte(serializedKalturaMatchCondition), &toSerialize)
	if errKalturaMatchCondition != nil {
		return []byte{}, errKalturaMatchCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaRegexCondition struct {
	value *KalturaRegexCondition
	isSet bool
}

func (v NullableKalturaRegexCondition) Get() *KalturaRegexCondition {
	return v.value
}

func (v *NullableKalturaRegexCondition) Set(val *KalturaRegexCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaRegexCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaRegexCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaRegexCondition(val *KalturaRegexCondition) *NullableKalturaRegexCondition {
	return &NullableKalturaRegexCondition{value: val, isSet: true}
}

func (v NullableKalturaRegexCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaRegexCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


