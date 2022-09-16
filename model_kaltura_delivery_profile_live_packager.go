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

// KalturaDeliveryProfileLivePackager struct for KalturaDeliveryProfileLivePackager
type KalturaDeliveryProfileLivePackager struct {
	KalturaDeliveryProfile
}

// NewKalturaDeliveryProfileLivePackager instantiates a new KalturaDeliveryProfileLivePackager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileLivePackager() *KalturaDeliveryProfileLivePackager {
	this := KalturaDeliveryProfileLivePackager{}
	return &this
}

// NewKalturaDeliveryProfileLivePackagerWithDefaults instantiates a new KalturaDeliveryProfileLivePackager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileLivePackagerWithDefaults() *KalturaDeliveryProfileLivePackager {
	this := KalturaDeliveryProfileLivePackager{}
	return &this
}

func (o KalturaDeliveryProfileLivePackager) MarshalJSON() ([]byte, error) {
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

type NullableKalturaDeliveryProfileLivePackager struct {
	value *KalturaDeliveryProfileLivePackager
	isSet bool
}

func (v NullableKalturaDeliveryProfileLivePackager) Get() *KalturaDeliveryProfileLivePackager {
	return v.value
}

func (v *NullableKalturaDeliveryProfileLivePackager) Set(val *KalturaDeliveryProfileLivePackager) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileLivePackager) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileLivePackager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileLivePackager(val *KalturaDeliveryProfileLivePackager) *NullableKalturaDeliveryProfileLivePackager {
	return &NullableKalturaDeliveryProfileLivePackager{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileLivePackager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileLivePackager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

