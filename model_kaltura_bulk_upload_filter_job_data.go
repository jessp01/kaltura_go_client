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

// KalturaBulkUploadFilterJobData Represents the Bulk upload job data for filter bulk upload
type KalturaBulkUploadFilterJobData struct {
	KalturaBulkUploadJobData
}

// NewKalturaBulkUploadFilterJobData instantiates a new KalturaBulkUploadFilterJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadFilterJobData() *KalturaBulkUploadFilterJobData {
	this := KalturaBulkUploadFilterJobData{}
	return &this
}

// NewKalturaBulkUploadFilterJobDataWithDefaults instantiates a new KalturaBulkUploadFilterJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadFilterJobDataWithDefaults() *KalturaBulkUploadFilterJobData {
	this := KalturaBulkUploadFilterJobData{}
	return &this
}

func (o KalturaBulkUploadFilterJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBulkUploadJobData, errKalturaBulkUploadJobData := json.Marshal(o.KalturaBulkUploadJobData)
	if errKalturaBulkUploadJobData != nil {
		return []byte{}, errKalturaBulkUploadJobData
	}
	errKalturaBulkUploadJobData = json.Unmarshal([]byte(serializedKalturaBulkUploadJobData), &toSerialize)
	if errKalturaBulkUploadJobData != nil {
		return []byte{}, errKalturaBulkUploadJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBulkUploadFilterJobData struct {
	value *KalturaBulkUploadFilterJobData
	isSet bool
}

func (v NullableKalturaBulkUploadFilterJobData) Get() *KalturaBulkUploadFilterJobData {
	return v.value
}

func (v *NullableKalturaBulkUploadFilterJobData) Set(val *KalturaBulkUploadFilterJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadFilterJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadFilterJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadFilterJobData(val *KalturaBulkUploadFilterJobData) *NullableKalturaBulkUploadFilterJobData {
	return &NullableKalturaBulkUploadFilterJobData{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadFilterJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadFilterJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


