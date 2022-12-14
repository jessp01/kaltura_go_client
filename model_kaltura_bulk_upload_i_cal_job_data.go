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

// KalturaBulkUploadICalJobData Represents the Bulk upload job data for iCal bulk upload
type KalturaBulkUploadICalJobData struct {
	KalturaBulkUploadScheduleEventJobData
}

// NewKalturaBulkUploadICalJobData instantiates a new KalturaBulkUploadICalJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadICalJobData() *KalturaBulkUploadICalJobData {
	this := KalturaBulkUploadICalJobData{}
	return &this
}

// NewKalturaBulkUploadICalJobDataWithDefaults instantiates a new KalturaBulkUploadICalJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadICalJobDataWithDefaults() *KalturaBulkUploadICalJobData {
	this := KalturaBulkUploadICalJobData{}
	return &this
}

func (o KalturaBulkUploadICalJobData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBulkUploadScheduleEventJobData, errKalturaBulkUploadScheduleEventJobData := json.Marshal(o.KalturaBulkUploadScheduleEventJobData)
	if errKalturaBulkUploadScheduleEventJobData != nil {
		return []byte{}, errKalturaBulkUploadScheduleEventJobData
	}
	errKalturaBulkUploadScheduleEventJobData = json.Unmarshal([]byte(serializedKalturaBulkUploadScheduleEventJobData), &toSerialize)
	if errKalturaBulkUploadScheduleEventJobData != nil {
		return []byte{}, errKalturaBulkUploadScheduleEventJobData
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBulkUploadICalJobData struct {
	value *KalturaBulkUploadICalJobData
	isSet bool
}

func (v NullableKalturaBulkUploadICalJobData) Get() *KalturaBulkUploadICalJobData {
	return v.value
}

func (v *NullableKalturaBulkUploadICalJobData) Set(val *KalturaBulkUploadICalJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadICalJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadICalJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadICalJobData(val *KalturaBulkUploadICalJobData) *NullableKalturaBulkUploadICalJobData {
	return &NullableKalturaBulkUploadICalJobData{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadICalJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadICalJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


