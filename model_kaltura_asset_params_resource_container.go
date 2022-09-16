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

// KalturaAssetParamsResourceContainer struct for KalturaAssetParamsResourceContainer
type KalturaAssetParamsResourceContainer struct {
	KalturaResource
}

// NewKalturaAssetParamsResourceContainer instantiates a new KalturaAssetParamsResourceContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAssetParamsResourceContainer() *KalturaAssetParamsResourceContainer {
	this := KalturaAssetParamsResourceContainer{}
	return &this
}

// NewKalturaAssetParamsResourceContainerWithDefaults instantiates a new KalturaAssetParamsResourceContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetParamsResourceContainerWithDefaults() *KalturaAssetParamsResourceContainer {
	this := KalturaAssetParamsResourceContainer{}
	return &this
}

func (o KalturaAssetParamsResourceContainer) MarshalJSON() ([]byte, error) {
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

type NullableKalturaAssetParamsResourceContainer struct {
	value *KalturaAssetParamsResourceContainer
	isSet bool
}

func (v NullableKalturaAssetParamsResourceContainer) Get() *KalturaAssetParamsResourceContainer {
	return v.value
}

func (v *NullableKalturaAssetParamsResourceContainer) Set(val *KalturaAssetParamsResourceContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAssetParamsResourceContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAssetParamsResourceContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAssetParamsResourceContainer(val *KalturaAssetParamsResourceContainer) *NullableKalturaAssetParamsResourceContainer {
	return &NullableKalturaAssetParamsResourceContainer{value: val, isSet: true}
}

func (v NullableKalturaAssetParamsResourceContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAssetParamsResourceContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

