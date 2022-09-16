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

// KalturaContentResource `abstract`  Is a unified way to add content to Kaltura whether it's an uploaded file, webcam recording, imported URL or existing file sync.
type KalturaContentResource struct {
	KalturaResource
}

// NewKalturaContentResource instantiates a new KalturaContentResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaContentResource() *KalturaContentResource {
	this := KalturaContentResource{}
	return &this
}

// NewKalturaContentResourceWithDefaults instantiates a new KalturaContentResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaContentResourceWithDefaults() *KalturaContentResource {
	this := KalturaContentResource{}
	return &this
}

func (o KalturaContentResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaResource, errKalturaResource := json.Marshal(o.KalturaResource)
	if errKalturaResource != nil {
		return []byte{}, errKalturaResource
	}
	errKalturaResource = json.Unmarshal([]byte(serializedKalturaResource), &toSerialize)
	if errKalturaResource != nil {
		return []byte{}, errKalturaResource
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaContentResource struct {
	value *KalturaContentResource
	isSet bool
}

func (v NullableKalturaContentResource) Get() *KalturaContentResource {
	return v.value
}

func (v *NullableKalturaContentResource) Set(val *KalturaContentResource) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaContentResource) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaContentResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaContentResource(val *KalturaContentResource) *NullableKalturaContentResource {
	return &NullableKalturaContentResource{value: val, isSet: true}
}

func (v NullableKalturaContentResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaContentResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

