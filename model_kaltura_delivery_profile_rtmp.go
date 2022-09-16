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

// KalturaDeliveryProfileRtmp struct for KalturaDeliveryProfileRtmp
type KalturaDeliveryProfileRtmp struct {
	KalturaDeliveryProfile
}

// NewKalturaDeliveryProfileRtmp instantiates a new KalturaDeliveryProfileRtmp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileRtmp() *KalturaDeliveryProfileRtmp {
	this := KalturaDeliveryProfileRtmp{}
	return &this
}

// NewKalturaDeliveryProfileRtmpWithDefaults instantiates a new KalturaDeliveryProfileRtmp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileRtmpWithDefaults() *KalturaDeliveryProfileRtmp {
	this := KalturaDeliveryProfileRtmp{}
	return &this
}

func (o KalturaDeliveryProfileRtmp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDeliveryProfile, errKalturaDeliveryProfile := json.Marshal(o.KalturaDeliveryProfile)
	if errKalturaDeliveryProfile != nil {
		return []byte{}, errKalturaDeliveryProfile
	}
	errKalturaDeliveryProfile = json.Unmarshal([]byte(serializedKalturaDeliveryProfile), &toSerialize)
	if errKalturaDeliveryProfile != nil {
		return []byte{}, errKalturaDeliveryProfile
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfileRtmp struct {
	value *KalturaDeliveryProfileRtmp
	isSet bool
}

func (v NullableKalturaDeliveryProfileRtmp) Get() *KalturaDeliveryProfileRtmp {
	return v.value
}

func (v *NullableKalturaDeliveryProfileRtmp) Set(val *KalturaDeliveryProfileRtmp) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileRtmp) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileRtmp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileRtmp(val *KalturaDeliveryProfileRtmp) *NullableKalturaDeliveryProfileRtmp {
	return &NullableKalturaDeliveryProfileRtmp{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileRtmp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileRtmp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

