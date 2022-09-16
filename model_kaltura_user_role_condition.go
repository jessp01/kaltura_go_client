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

// KalturaUserRoleCondition struct for KalturaUserRoleCondition
type KalturaUserRoleCondition struct {
	KalturaCondition
}

// NewKalturaUserRoleCondition instantiates a new KalturaUserRoleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUserRoleCondition() *KalturaUserRoleCondition {
	this := KalturaUserRoleCondition{}
	return &this
}

// NewKalturaUserRoleConditionWithDefaults instantiates a new KalturaUserRoleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUserRoleConditionWithDefaults() *KalturaUserRoleCondition {
	this := KalturaUserRoleCondition{}
	return &this
}

func (o KalturaUserRoleCondition) MarshalJSON() ([]byte, error) {
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

type NullableKalturaUserRoleCondition struct {
	value *KalturaUserRoleCondition
	isSet bool
}

func (v NullableKalturaUserRoleCondition) Get() *KalturaUserRoleCondition {
	return v.value
}

func (v *NullableKalturaUserRoleCondition) Set(val *KalturaUserRoleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUserRoleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUserRoleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUserRoleCondition(val *KalturaUserRoleCondition) *NullableKalturaUserRoleCondition {
	return &NullableKalturaUserRoleCondition{value: val, isSet: true}
}

func (v NullableKalturaUserRoleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUserRoleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


