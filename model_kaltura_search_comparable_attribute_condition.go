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

// KalturaSearchComparableAttributeCondition `abstract`
type KalturaSearchComparableAttributeCondition struct {
	KalturaAttributeCondition
}

// NewKalturaSearchComparableAttributeCondition instantiates a new KalturaSearchComparableAttributeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaSearchComparableAttributeCondition() *KalturaSearchComparableAttributeCondition {
	this := KalturaSearchComparableAttributeCondition{}
	return &this
}

// NewKalturaSearchComparableAttributeConditionWithDefaults instantiates a new KalturaSearchComparableAttributeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaSearchComparableAttributeConditionWithDefaults() *KalturaSearchComparableAttributeCondition {
	this := KalturaSearchComparableAttributeCondition{}
	return &this
}

func (o KalturaSearchComparableAttributeCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaAttributeCondition, errKalturaAttributeCondition := json.Marshal(o.KalturaAttributeCondition)
	if errKalturaAttributeCondition != nil {
		return []byte{}, errKalturaAttributeCondition
	}
	errKalturaAttributeCondition = json.Unmarshal([]byte(serializedKalturaAttributeCondition), &toSerialize)
	if errKalturaAttributeCondition != nil {
		return []byte{}, errKalturaAttributeCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaSearchComparableAttributeCondition struct {
	value *KalturaSearchComparableAttributeCondition
	isSet bool
}

func (v NullableKalturaSearchComparableAttributeCondition) Get() *KalturaSearchComparableAttributeCondition {
	return v.value
}

func (v *NullableKalturaSearchComparableAttributeCondition) Set(val *KalturaSearchComparableAttributeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaSearchComparableAttributeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaSearchComparableAttributeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaSearchComparableAttributeCondition(val *KalturaSearchComparableAttributeCondition) *NullableKalturaSearchComparableAttributeCondition {
	return &NullableKalturaSearchComparableAttributeCondition{value: val, isSet: true}
}

func (v NullableKalturaSearchComparableAttributeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaSearchComparableAttributeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


