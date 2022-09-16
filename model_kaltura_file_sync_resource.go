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

// KalturaFileSyncResource Used to ingest media that is already ingested to Kaltura system as a different file in the past, the new created flavor asset will be ready immediately using a file sync of link type that will point to the existing file sync.
type KalturaFileSyncResource struct {
	KalturaContentResource
}

// NewKalturaFileSyncResource instantiates a new KalturaFileSyncResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFileSyncResource() *KalturaFileSyncResource {
	this := KalturaFileSyncResource{}
	return &this
}

// NewKalturaFileSyncResourceWithDefaults instantiates a new KalturaFileSyncResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFileSyncResourceWithDefaults() *KalturaFileSyncResource {
	this := KalturaFileSyncResource{}
	return &this
}

func (o KalturaFileSyncResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaContentResource, errKalturaContentResource := json.Marshal(o.KalturaContentResource)
	if errKalturaContentResource != nil {
		return []byte{}, errKalturaContentResource
	}
	errKalturaContentResource = json.Unmarshal([]byte(serializedKalturaContentResource), &toSerialize)
	if errKalturaContentResource != nil {
		return []byte{}, errKalturaContentResource
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFileSyncResource struct {
	value *KalturaFileSyncResource
	isSet bool
}

func (v NullableKalturaFileSyncResource) Get() *KalturaFileSyncResource {
	return v.value
}

func (v *NullableKalturaFileSyncResource) Set(val *KalturaFileSyncResource) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFileSyncResource) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFileSyncResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFileSyncResource(val *KalturaFileSyncResource) *NullableKalturaFileSyncResource {
	return &NullableKalturaFileSyncResource{value: val, isSet: true}
}

func (v NullableKalturaFileSyncResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFileSyncResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


