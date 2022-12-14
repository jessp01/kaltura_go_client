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

// KalturaPlayableEntryMatchAttributeCondition Auto-generated class.  Used to search KalturaPlayableEntry attributes. Use KalturaPlayableEntryMatchAttribute enum to provide attribute name. /
type KalturaPlayableEntryMatchAttributeCondition struct {
	KalturaSearchMatchAttributeCondition
}

// NewKalturaPlayableEntryMatchAttributeCondition instantiates a new KalturaPlayableEntryMatchAttributeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPlayableEntryMatchAttributeCondition() *KalturaPlayableEntryMatchAttributeCondition {
	this := KalturaPlayableEntryMatchAttributeCondition{}
	return &this
}

// NewKalturaPlayableEntryMatchAttributeConditionWithDefaults instantiates a new KalturaPlayableEntryMatchAttributeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPlayableEntryMatchAttributeConditionWithDefaults() *KalturaPlayableEntryMatchAttributeCondition {
	this := KalturaPlayableEntryMatchAttributeCondition{}
	return &this
}

func (o KalturaPlayableEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaSearchMatchAttributeCondition, errKalturaSearchMatchAttributeCondition := json.Marshal(o.KalturaSearchMatchAttributeCondition)
	if errKalturaSearchMatchAttributeCondition != nil {
		return []byte{}, errKalturaSearchMatchAttributeCondition
	}
	errKalturaSearchMatchAttributeCondition = json.Unmarshal([]byte(serializedKalturaSearchMatchAttributeCondition), &toSerialize)
	if errKalturaSearchMatchAttributeCondition != nil {
		return []byte{}, errKalturaSearchMatchAttributeCondition
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPlayableEntryMatchAttributeCondition struct {
	value *KalturaPlayableEntryMatchAttributeCondition
	isSet bool
}

func (v NullableKalturaPlayableEntryMatchAttributeCondition) Get() *KalturaPlayableEntryMatchAttributeCondition {
	return v.value
}

func (v *NullableKalturaPlayableEntryMatchAttributeCondition) Set(val *KalturaPlayableEntryMatchAttributeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPlayableEntryMatchAttributeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPlayableEntryMatchAttributeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPlayableEntryMatchAttributeCondition(val *KalturaPlayableEntryMatchAttributeCondition) *NullableKalturaPlayableEntryMatchAttributeCondition {
	return &NullableKalturaPlayableEntryMatchAttributeCondition{value: val, isSet: true}
}

func (v NullableKalturaPlayableEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPlayableEntryMatchAttributeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


