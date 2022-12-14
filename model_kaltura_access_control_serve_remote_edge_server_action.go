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

// KalturaAccessControlServeRemoteEdgeServerAction struct for KalturaAccessControlServeRemoteEdgeServerAction
type KalturaAccessControlServeRemoteEdgeServerAction struct {
	KalturaRuleAction
}

// NewKalturaAccessControlServeRemoteEdgeServerAction instantiates a new KalturaAccessControlServeRemoteEdgeServerAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAccessControlServeRemoteEdgeServerAction() *KalturaAccessControlServeRemoteEdgeServerAction {
	this := KalturaAccessControlServeRemoteEdgeServerAction{}
	return &this
}

// NewKalturaAccessControlServeRemoteEdgeServerActionWithDefaults instantiates a new KalturaAccessControlServeRemoteEdgeServerAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAccessControlServeRemoteEdgeServerActionWithDefaults() *KalturaAccessControlServeRemoteEdgeServerAction {
	this := KalturaAccessControlServeRemoteEdgeServerAction{}
	return &this
}

func (o KalturaAccessControlServeRemoteEdgeServerAction) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAccessControlServeRemoteEdgeServerAction struct {
	value *KalturaAccessControlServeRemoteEdgeServerAction
	isSet bool
}

func (v NullableKalturaAccessControlServeRemoteEdgeServerAction) Get() *KalturaAccessControlServeRemoteEdgeServerAction {
	return v.value
}

func (v *NullableKalturaAccessControlServeRemoteEdgeServerAction) Set(val *KalturaAccessControlServeRemoteEdgeServerAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAccessControlServeRemoteEdgeServerAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAccessControlServeRemoteEdgeServerAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAccessControlServeRemoteEdgeServerAction(val *KalturaAccessControlServeRemoteEdgeServerAction) *NullableKalturaAccessControlServeRemoteEdgeServerAction {
	return &NullableKalturaAccessControlServeRemoteEdgeServerAction{value: val, isSet: true}
}

func (v NullableKalturaAccessControlServeRemoteEdgeServerAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAccessControlServeRemoteEdgeServerAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


