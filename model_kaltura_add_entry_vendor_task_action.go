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

// KalturaAddEntryVendorTaskAction struct for KalturaAddEntryVendorTaskAction
type KalturaAddEntryVendorTaskAction struct {
	KalturaRuleAction
}

// NewKalturaAddEntryVendorTaskAction instantiates a new KalturaAddEntryVendorTaskAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAddEntryVendorTaskAction() *KalturaAddEntryVendorTaskAction {
	this := KalturaAddEntryVendorTaskAction{}
	return &this
}

// NewKalturaAddEntryVendorTaskActionWithDefaults instantiates a new KalturaAddEntryVendorTaskAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAddEntryVendorTaskActionWithDefaults() *KalturaAddEntryVendorTaskAction {
	this := KalturaAddEntryVendorTaskAction{}
	return &this
}

func (o KalturaAddEntryVendorTaskAction) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAddEntryVendorTaskAction struct {
	value *KalturaAddEntryVendorTaskAction
	isSet bool
}

func (v NullableKalturaAddEntryVendorTaskAction) Get() *KalturaAddEntryVendorTaskAction {
	return v.value
}

func (v *NullableKalturaAddEntryVendorTaskAction) Set(val *KalturaAddEntryVendorTaskAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAddEntryVendorTaskAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAddEntryVendorTaskAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAddEntryVendorTaskAction(val *KalturaAddEntryVendorTaskAction) *NullableKalturaAddEntryVendorTaskAction {
	return &NullableKalturaAddEntryVendorTaskAction{value: val, isSet: true}
}

func (v NullableKalturaAddEntryVendorTaskAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAddEntryVendorTaskAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


