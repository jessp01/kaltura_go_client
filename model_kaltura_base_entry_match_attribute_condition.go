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

// KalturaBaseEntryMatchAttributeCondition Auto-generated class.  Used to search KalturaBaseEntry attributes. Use KalturaBaseEntryMatchAttribute enum to provide attribute name. /
type KalturaBaseEntryMatchAttributeCondition struct {
	KalturaSearchMatchAttributeCondition
}

// NewKalturaBaseEntryMatchAttributeCondition instantiates a new KalturaBaseEntryMatchAttributeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBaseEntryMatchAttributeCondition() *KalturaBaseEntryMatchAttributeCondition {
	this := KalturaBaseEntryMatchAttributeCondition{}
	return &this
}

// NewKalturaBaseEntryMatchAttributeConditionWithDefaults instantiates a new KalturaBaseEntryMatchAttributeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBaseEntryMatchAttributeConditionWithDefaults() *KalturaBaseEntryMatchAttributeCondition {
	this := KalturaBaseEntryMatchAttributeCondition{}
	return &this
}

func (o KalturaBaseEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBaseEntryMatchAttributeCondition struct {
	value *KalturaBaseEntryMatchAttributeCondition
	isSet bool
}

func (v NullableKalturaBaseEntryMatchAttributeCondition) Get() *KalturaBaseEntryMatchAttributeCondition {
	return v.value
}

func (v *NullableKalturaBaseEntryMatchAttributeCondition) Set(val *KalturaBaseEntryMatchAttributeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBaseEntryMatchAttributeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBaseEntryMatchAttributeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBaseEntryMatchAttributeCondition(val *KalturaBaseEntryMatchAttributeCondition) *NullableKalturaBaseEntryMatchAttributeCondition {
	return &NullableKalturaBaseEntryMatchAttributeCondition{value: val, isSet: true}
}

func (v NullableKalturaBaseEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBaseEntryMatchAttributeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

