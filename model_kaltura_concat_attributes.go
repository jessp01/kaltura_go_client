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

// KalturaConcatAttributes Concat operation attributes
type KalturaConcatAttributes struct {
	KalturaOperationAttributes
}

// NewKalturaConcatAttributes instantiates a new KalturaConcatAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaConcatAttributes() *KalturaConcatAttributes {
	this := KalturaConcatAttributes{}
	return &this
}

// NewKalturaConcatAttributesWithDefaults instantiates a new KalturaConcatAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaConcatAttributesWithDefaults() *KalturaConcatAttributes {
	this := KalturaConcatAttributes{}
	return &this
}

func (o KalturaConcatAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaOperationAttributes, errKalturaOperationAttributes := json.Marshal(o.KalturaOperationAttributes)
	if errKalturaOperationAttributes != nil {
		return []byte{}, errKalturaOperationAttributes
	}
	errKalturaOperationAttributes = json.Unmarshal([]byte(serializedKalturaOperationAttributes), &toSerialize)
	if errKalturaOperationAttributes != nil {
		return []byte{}, errKalturaOperationAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaConcatAttributes struct {
	value *KalturaConcatAttributes
	isSet bool
}

func (v NullableKalturaConcatAttributes) Get() *KalturaConcatAttributes {
	return v.value
}

func (v *NullableKalturaConcatAttributes) Set(val *KalturaConcatAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaConcatAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaConcatAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaConcatAttributes(val *KalturaConcatAttributes) *NullableKalturaConcatAttributes {
	return &NullableKalturaConcatAttributes{value: val, isSet: true}
}

func (v NullableKalturaConcatAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaConcatAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


