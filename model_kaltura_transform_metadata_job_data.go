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

// KalturaTransformMetadataJobData struct for KalturaTransformMetadataJobData
type KalturaTransformMetadataJobData struct {
	KalturaJobData
}

// NewKalturaTransformMetadataJobData instantiates a new KalturaTransformMetadataJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaTransformMetadataJobData() *KalturaTransformMetadataJobData {
	this := KalturaTransformMetadataJobData{}
	return &this
}

// NewKalturaTransformMetadataJobDataWithDefaults instantiates a new KalturaTransformMetadataJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaTransformMetadataJobDataWithDefaults() *KalturaTransformMetadataJobData {
	this := KalturaTransformMetadataJobData{}
	return &this
}

func (o KalturaTransformMetadataJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaJobData, errKalturaJobData := json.Marshal(o.KalturaJobData)
	if errKalturaJobData != nil {
		return []byte{}, errKalturaJobData
	}
	errKalturaJobData = json.Unmarshal([]byte(serializedKalturaJobData), &toSerialize)
	if errKalturaJobData != nil {
		return []byte{}, errKalturaJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaTransformMetadataJobData struct {
	value *KalturaTransformMetadataJobData
	isSet bool
}

func (v NullableKalturaTransformMetadataJobData) Get() *KalturaTransformMetadataJobData {
	return v.value
}

func (v *NullableKalturaTransformMetadataJobData) Set(val *KalturaTransformMetadataJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaTransformMetadataJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaTransformMetadataJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaTransformMetadataJobData(val *KalturaTransformMetadataJobData) *NullableKalturaTransformMetadataJobData {
	return &NullableKalturaTransformMetadataJobData{value: val, isSet: true}
}

func (v NullableKalturaTransformMetadataJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaTransformMetadataJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


