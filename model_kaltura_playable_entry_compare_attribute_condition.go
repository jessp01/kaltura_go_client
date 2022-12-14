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

// KalturaPlayableEntryCompareAttributeCondition Auto-generated class.  Used to search KalturaPlayableEntry attributes. Use KalturaPlayableEntryCompareAttribute enum to provide attribute name. /
type KalturaPlayableEntryCompareAttributeCondition struct {
	KalturaSearchComparableAttributeCondition
}

// NewKalturaPlayableEntryCompareAttributeCondition instantiates a new KalturaPlayableEntryCompareAttributeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayableEntryCompareAttributeCondition() *KalturaPlayableEntryCompareAttributeCondition {
	this := KalturaPlayableEntryCompareAttributeCondition{}
	return &this
}

// NewKalturaPlayableEntryCompareAttributeConditionWithDefaults instantiates a new KalturaPlayableEntryCompareAttributeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayableEntryCompareAttributeConditionWithDefaults() *KalturaPlayableEntryCompareAttributeCondition {
	this := KalturaPlayableEntryCompareAttributeCondition{}
	return &this
}

func (o KalturaPlayableEntryCompareAttributeCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearchComparableAttributeCondition, errKalturaSearchComparableAttributeCondition := json.Marshal(o.KalturaSearchComparableAttributeCondition)
	if errKalturaSearchComparableAttributeCondition != nil {
		return []byte{}, errKalturaSearchComparableAttributeCondition
	}
	errKalturaSearchComparableAttributeCondition = json.Unmarshal([]byte(serializedKalturaSearchComparableAttributeCondition), &toSerialize)
	if errKalturaSearchComparableAttributeCondition != nil {
		return []byte{}, errKalturaSearchComparableAttributeCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlayableEntryCompareAttributeCondition struct {
	value *KalturaPlayableEntryCompareAttributeCondition
	isSet bool
}

func (v NullableKalturaPlayableEntryCompareAttributeCondition) Get() *KalturaPlayableEntryCompareAttributeCondition {
	return v.value
}

func (v *NullableKalturaPlayableEntryCompareAttributeCondition) Set(val *KalturaPlayableEntryCompareAttributeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayableEntryCompareAttributeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayableEntryCompareAttributeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayableEntryCompareAttributeCondition(val *KalturaPlayableEntryCompareAttributeCondition) *NullableKalturaPlayableEntryCompareAttributeCondition {
	return &NullableKalturaPlayableEntryCompareAttributeCondition{value: val, isSet: true}
}

func (v NullableKalturaPlayableEntryCompareAttributeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayableEntryCompareAttributeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


