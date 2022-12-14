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

// KalturaAccessControlLimitDeliveryProfilesAction struct for KalturaAccessControlLimitDeliveryProfilesAction
type KalturaAccessControlLimitDeliveryProfilesAction struct {
	KalturaRuleAction
}

// NewKalturaAccessControlLimitDeliveryProfilesAction instantiates a new KalturaAccessControlLimitDeliveryProfilesAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAccessControlLimitDeliveryProfilesAction() *KalturaAccessControlLimitDeliveryProfilesAction {
	this := KalturaAccessControlLimitDeliveryProfilesAction{}
	return &this
}

// NewKalturaAccessControlLimitDeliveryProfilesActionWithDefaults instantiates a new KalturaAccessControlLimitDeliveryProfilesAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAccessControlLimitDeliveryProfilesActionWithDefaults() *KalturaAccessControlLimitDeliveryProfilesAction {
	this := KalturaAccessControlLimitDeliveryProfilesAction{}
	return &this
}

func (o KalturaAccessControlLimitDeliveryProfilesAction) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAccessControlLimitDeliveryProfilesAction struct {
	value *KalturaAccessControlLimitDeliveryProfilesAction
	isSet bool
}

func (v NullableKalturaAccessControlLimitDeliveryProfilesAction) Get() *KalturaAccessControlLimitDeliveryProfilesAction {
	return v.value
}

func (v *NullableKalturaAccessControlLimitDeliveryProfilesAction) Set(val *KalturaAccessControlLimitDeliveryProfilesAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAccessControlLimitDeliveryProfilesAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAccessControlLimitDeliveryProfilesAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAccessControlLimitDeliveryProfilesAction(val *KalturaAccessControlLimitDeliveryProfilesAction) *NullableKalturaAccessControlLimitDeliveryProfilesAction {
	return &NullableKalturaAccessControlLimitDeliveryProfilesAction{value: val, isSet: true}
}

func (v NullableKalturaAccessControlLimitDeliveryProfilesAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAccessControlLimitDeliveryProfilesAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


