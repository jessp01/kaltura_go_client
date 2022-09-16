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

// KalturaAssetResource Used to ingest media that is already ingested to Kaltura system as a different flavor asset in the past, the new created flavor asset will be ready immediately using a file sync of link type that will point to the existing file sync of the existing flavor asset.
type KalturaAssetResource struct {
	KalturaContentResource
}

// NewKalturaAssetResource instantiates a new KalturaAssetResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAssetResource() *KalturaAssetResource {
	this := KalturaAssetResource{}
	return &this
}

// NewKalturaAssetResourceWithDefaults instantiates a new KalturaAssetResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetResourceWithDefaults() *KalturaAssetResource {
	this := KalturaAssetResource{}
	return &this
}

func (o KalturaAssetResource) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAssetResource struct {
	value *KalturaAssetResource
	isSet bool
}

func (v NullableKalturaAssetResource) Get() *KalturaAssetResource {
	return v.value
}

func (v *NullableKalturaAssetResource) Set(val *KalturaAssetResource) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAssetResource) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAssetResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAssetResource(val *KalturaAssetResource) *NullableKalturaAssetResource {
	return &NullableKalturaAssetResource{value: val, isSet: true}
}

func (v NullableKalturaAssetResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAssetResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


