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

// KalturaClipAttributes Clip operation attributes
type KalturaClipAttributes struct {
	KalturaOperationAttributes
}

// NewKalturaClipAttributes instantiates a new KalturaClipAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaClipAttributes() *KalturaClipAttributes {
	this := KalturaClipAttributes{}
	return &this
}

// NewKalturaClipAttributesWithDefaults instantiates a new KalturaClipAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaClipAttributesWithDefaults() *KalturaClipAttributes {
	this := KalturaClipAttributes{}
	return &this
}

func (o KalturaClipAttributes) MarshalJSON() ([]byte, error) {
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

type NullableKalturaClipAttributes struct {
	value *KalturaClipAttributes
	isSet bool
}

func (v NullableKalturaClipAttributes) Get() *KalturaClipAttributes {
	return v.value
}

func (v *NullableKalturaClipAttributes) Set(val *KalturaClipAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaClipAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaClipAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaClipAttributes(val *KalturaClipAttributes) *NullableKalturaClipAttributes {
	return &NullableKalturaClipAttributes{value: val, isSet: true}
}

func (v NullableKalturaClipAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaClipAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


