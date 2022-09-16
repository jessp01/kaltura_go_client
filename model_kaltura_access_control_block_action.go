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

// KalturaAccessControlBlockAction struct for KalturaAccessControlBlockAction
type KalturaAccessControlBlockAction struct {
	KalturaRuleAction
}

// NewKalturaAccessControlBlockAction instantiates a new KalturaAccessControlBlockAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAccessControlBlockAction() *KalturaAccessControlBlockAction {
	this := KalturaAccessControlBlockAction{}
	return &this
}

// NewKalturaAccessControlBlockActionWithDefaults instantiates a new KalturaAccessControlBlockAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAccessControlBlockActionWithDefaults() *KalturaAccessControlBlockAction {
	this := KalturaAccessControlBlockAction{}
	return &this
}

func (o KalturaAccessControlBlockAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaRuleAction, errKalturaRuleAction := json.Marshal(o.KalturaRuleAction)
	if errKalturaRuleAction != nil {
		return []byte{}, errKalturaRuleAction
	}
	errKalturaRuleAction = json.Unmarshal([]byte(serializedKalturaRuleAction), &toSerialize)
	if errKalturaRuleAction != nil {
		return []byte{}, errKalturaRuleAction
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAccessControlBlockAction struct {
	value *KalturaAccessControlBlockAction
	isSet bool
}

func (v NullableKalturaAccessControlBlockAction) Get() *KalturaAccessControlBlockAction {
	return v.value
}

func (v *NullableKalturaAccessControlBlockAction) Set(val *KalturaAccessControlBlockAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAccessControlBlockAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAccessControlBlockAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAccessControlBlockAction(val *KalturaAccessControlBlockAction) *NullableKalturaAccessControlBlockAction {
	return &NullableKalturaAccessControlBlockAction{value: val, isSet: true}
}

func (v NullableKalturaAccessControlBlockAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAccessControlBlockAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

