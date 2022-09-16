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

// KalturaBulkUploadScheduleEventCsvJobData Represents the Bulk upload job data for CSV bulk upload
type KalturaBulkUploadScheduleEventCsvJobData struct {
	KalturaBulkUploadScheduleEventJobData
}

// NewKalturaBulkUploadScheduleEventCsvJobData instantiates a new KalturaBulkUploadScheduleEventCsvJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadScheduleEventCsvJobData() *KalturaBulkUploadScheduleEventCsvJobData {
	this := KalturaBulkUploadScheduleEventCsvJobData{}
	return &this
}

// NewKalturaBulkUploadScheduleEventCsvJobDataWithDefaults instantiates a new KalturaBulkUploadScheduleEventCsvJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadScheduleEventCsvJobDataWithDefaults() *KalturaBulkUploadScheduleEventCsvJobData {
	this := KalturaBulkUploadScheduleEventCsvJobData{}
	return &this
}

func (o KalturaBulkUploadScheduleEventCsvJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBulkUploadScheduleEventCsvJobData struct {
	value *KalturaBulkUploadScheduleEventCsvJobData
	isSet bool
}

func (v NullableKalturaBulkUploadScheduleEventCsvJobData) Get() *KalturaBulkUploadScheduleEventCsvJobData {
	return v.value
}

func (v *NullableKalturaBulkUploadScheduleEventCsvJobData) Set(val *KalturaBulkUploadScheduleEventCsvJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadScheduleEventCsvJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadScheduleEventCsvJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadScheduleEventCsvJobData(val *KalturaBulkUploadScheduleEventCsvJobData) *NullableKalturaBulkUploadScheduleEventCsvJobData {
	return &NullableKalturaBulkUploadScheduleEventCsvJobData{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadScheduleEventCsvJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadScheduleEventCsvJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


