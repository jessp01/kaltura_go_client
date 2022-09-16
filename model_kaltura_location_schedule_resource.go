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

// KalturaLocationScheduleResource struct for KalturaLocationScheduleResource
type KalturaLocationScheduleResource struct {
	KalturaScheduleResource
}

// NewKalturaLocationScheduleResource instantiates a new KalturaLocationScheduleResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaLocationScheduleResource() *KalturaLocationScheduleResource {
	this := KalturaLocationScheduleResource{}
	return &this
}

// NewKalturaLocationScheduleResourceWithDefaults instantiates a new KalturaLocationScheduleResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaLocationScheduleResourceWithDefaults() *KalturaLocationScheduleResource {
	this := KalturaLocationScheduleResource{}
	return &this
}

func (o KalturaLocationScheduleResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaScheduleResource, errKalturaScheduleResource := json.Marshal(o.KalturaScheduleResource)
	if errKalturaScheduleResource != nil {
		return []byte{}, errKalturaScheduleResource
	}
	errKalturaScheduleResource = json.Unmarshal([]byte(serializedKalturaScheduleResource), &toSerialize)
	if errKalturaScheduleResource != nil {
		return []byte{}, errKalturaScheduleResource
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaLocationScheduleResource struct {
	value *KalturaLocationScheduleResource
	isSet bool
}

func (v NullableKalturaLocationScheduleResource) Get() *KalturaLocationScheduleResource {
	return v.value
}

func (v *NullableKalturaLocationScheduleResource) Set(val *KalturaLocationScheduleResource) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaLocationScheduleResource) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaLocationScheduleResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaLocationScheduleResource(val *KalturaLocationScheduleResource) *NullableKalturaLocationScheduleResource {
	return &NullableKalturaLocationScheduleResource{value: val, isSet: true}
}

func (v NullableKalturaLocationScheduleResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaLocationScheduleResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


