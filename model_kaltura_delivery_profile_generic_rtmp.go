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

// KalturaDeliveryProfileGenericRtmp struct for KalturaDeliveryProfileGenericRtmp
type KalturaDeliveryProfileGenericRtmp struct {
	KalturaDeliveryProfileRtmp
}

// NewKalturaDeliveryProfileGenericRtmp instantiates a new KalturaDeliveryProfileGenericRtmp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfileGenericRtmp() *KalturaDeliveryProfileGenericRtmp {
	this := KalturaDeliveryProfileGenericRtmp{}
	return &this
}

// NewKalturaDeliveryProfileGenericRtmpWithDefaults instantiates a new KalturaDeliveryProfileGenericRtmp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileGenericRtmpWithDefaults() *KalturaDeliveryProfileGenericRtmp {
	this := KalturaDeliveryProfileGenericRtmp{}
	return &this
}

func (o KalturaDeliveryProfileGenericRtmp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDeliveryProfileRtmp, errKalturaDeliveryProfileRtmp := json.Marshal(o.KalturaDeliveryProfileRtmp)
	if errKalturaDeliveryProfileRtmp != nil {
		return []byte{}, errKalturaDeliveryProfileRtmp
	}
	errKalturaDeliveryProfileRtmp = json.Unmarshal([]byte(serializedKalturaDeliveryProfileRtmp), &toSerialize)
	if errKalturaDeliveryProfileRtmp != nil {
		return []byte{}, errKalturaDeliveryProfileRtmp
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfileGenericRtmp struct {
	value *KalturaDeliveryProfileGenericRtmp
	isSet bool
}

func (v NullableKalturaDeliveryProfileGenericRtmp) Get() *KalturaDeliveryProfileGenericRtmp {
	return v.value
}

func (v *NullableKalturaDeliveryProfileGenericRtmp) Set(val *KalturaDeliveryProfileGenericRtmp) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfileGenericRtmp) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfileGenericRtmp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfileGenericRtmp(val *KalturaDeliveryProfileGenericRtmp) *NullableKalturaDeliveryProfileGenericRtmp {
	return &NullableKalturaDeliveryProfileGenericRtmp{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfileGenericRtmp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfileGenericRtmp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

