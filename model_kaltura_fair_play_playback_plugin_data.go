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

// KalturaFairPlayPlaybackPluginData struct for KalturaFairPlayPlaybackPluginData
type KalturaFairPlayPlaybackPluginData struct {
	KalturaDrmPlaybackPluginData
}

// NewKalturaFairPlayPlaybackPluginData instantiates a new KalturaFairPlayPlaybackPluginData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFairPlayPlaybackPluginData() *KalturaFairPlayPlaybackPluginData {
	this := KalturaFairPlayPlaybackPluginData{}
	return &this
}

// NewKalturaFairPlayPlaybackPluginDataWithDefaults instantiates a new KalturaFairPlayPlaybackPluginData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFairPlayPlaybackPluginDataWithDefaults() *KalturaFairPlayPlaybackPluginData {
	this := KalturaFairPlayPlaybackPluginData{}
	return &this
}

func (o KalturaFairPlayPlaybackPluginData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaDrmPlaybackPluginData, errKalturaDrmPlaybackPluginData := json.Marshal(o.KalturaDrmPlaybackPluginData)
	if errKalturaDrmPlaybackPluginData != nil {
		return []byte{}, errKalturaDrmPlaybackPluginData
	}
	errKalturaDrmPlaybackPluginData = json.Unmarshal([]byte(serializedKalturaDrmPlaybackPluginData), &toSerialize)
	if errKalturaDrmPlaybackPluginData != nil {
		return []byte{}, errKalturaDrmPlaybackPluginData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFairPlayPlaybackPluginData struct {
	value *KalturaFairPlayPlaybackPluginData
	isSet bool
}

func (v NullableKalturaFairPlayPlaybackPluginData) Get() *KalturaFairPlayPlaybackPluginData {
	return v.value
}

func (v *NullableKalturaFairPlayPlaybackPluginData) Set(val *KalturaFairPlayPlaybackPluginData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFairPlayPlaybackPluginData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFairPlayPlaybackPluginData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFairPlayPlaybackPluginData(val *KalturaFairPlayPlaybackPluginData) *NullableKalturaFairPlayPlaybackPluginData {
	return &NullableKalturaFairPlayPlaybackPluginData{value: val, isSet: true}
}

func (v NullableKalturaFairPlayPlaybackPluginData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFairPlayPlaybackPluginData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


