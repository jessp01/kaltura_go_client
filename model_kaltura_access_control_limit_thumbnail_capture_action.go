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

// KalturaAccessControlLimitThumbnailCaptureAction struct for KalturaAccessControlLimitThumbnailCaptureAction
type KalturaAccessControlLimitThumbnailCaptureAction struct {
	KalturaRuleAction
}

// NewKalturaAccessControlLimitThumbnailCaptureAction instantiates a new KalturaAccessControlLimitThumbnailCaptureAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAccessControlLimitThumbnailCaptureAction() *KalturaAccessControlLimitThumbnailCaptureAction {
	this := KalturaAccessControlLimitThumbnailCaptureAction{}
	return &this
}

// NewKalturaAccessControlLimitThumbnailCaptureActionWithDefaults instantiates a new KalturaAccessControlLimitThumbnailCaptureAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAccessControlLimitThumbnailCaptureActionWithDefaults() *KalturaAccessControlLimitThumbnailCaptureAction {
	this := KalturaAccessControlLimitThumbnailCaptureAction{}
	return &this
}

func (o KalturaAccessControlLimitThumbnailCaptureAction) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAccessControlLimitThumbnailCaptureAction struct {
	value *KalturaAccessControlLimitThumbnailCaptureAction
	isSet bool
}

func (v NullableKalturaAccessControlLimitThumbnailCaptureAction) Get() *KalturaAccessControlLimitThumbnailCaptureAction {
	return v.value
}

func (v *NullableKalturaAccessControlLimitThumbnailCaptureAction) Set(val *KalturaAccessControlLimitThumbnailCaptureAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAccessControlLimitThumbnailCaptureAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAccessControlLimitThumbnailCaptureAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAccessControlLimitThumbnailCaptureAction(val *KalturaAccessControlLimitThumbnailCaptureAction) *NullableKalturaAccessControlLimitThumbnailCaptureAction {
	return &NullableKalturaAccessControlLimitThumbnailCaptureAction{value: val, isSet: true}
}

func (v NullableKalturaAccessControlLimitThumbnailCaptureAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAccessControlLimitThumbnailCaptureAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

