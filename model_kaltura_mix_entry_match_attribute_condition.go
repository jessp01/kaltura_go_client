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

// KalturaMixEntryMatchAttributeCondition Auto-generated class.  Used to search KalturaMixEntry attributes. Use KalturaMixEntryMatchAttribute enum to provide attribute name. /
type KalturaMixEntryMatchAttributeCondition struct {
	KalturaSearchMatchAttributeCondition
}

// NewKalturaMixEntryMatchAttributeCondition instantiates a new KalturaMixEntryMatchAttributeCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaMixEntryMatchAttributeCondition() *KalturaMixEntryMatchAttributeCondition {
	this := KalturaMixEntryMatchAttributeCondition{}
	return &this
}

// NewKalturaMixEntryMatchAttributeConditionWithDefaults instantiates a new KalturaMixEntryMatchAttributeCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaMixEntryMatchAttributeConditionWithDefaults() *KalturaMixEntryMatchAttributeCondition {
	this := KalturaMixEntryMatchAttributeCondition{}
	return &this
}

func (o KalturaMixEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
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

type NullableKalturaMixEntryMatchAttributeCondition struct {
	value *KalturaMixEntryMatchAttributeCondition
	isSet bool
}

func (v NullableKalturaMixEntryMatchAttributeCondition) Get() *KalturaMixEntryMatchAttributeCondition {
	return v.value
}

func (v *NullableKalturaMixEntryMatchAttributeCondition) Set(val *KalturaMixEntryMatchAttributeCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaMixEntryMatchAttributeCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaMixEntryMatchAttributeCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaMixEntryMatchAttributeCondition(val *KalturaMixEntryMatchAttributeCondition) *NullableKalturaMixEntryMatchAttributeCondition {
	return &NullableKalturaMixEntryMatchAttributeCondition{value: val, isSet: true}
}

func (v NullableKalturaMixEntryMatchAttributeCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaMixEntryMatchAttributeCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


