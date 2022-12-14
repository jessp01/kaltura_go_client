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

// KalturaBulkUploadXmlJobData Represents the Bulk upload job data for xml bulk upload
type KalturaBulkUploadXmlJobData struct {
	KalturaBulkUploadJobData
}

// NewKalturaBulkUploadXmlJobData instantiates a new KalturaBulkUploadXmlJobData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadXmlJobData() *KalturaBulkUploadXmlJobData {
	this := KalturaBulkUploadXmlJobData{}
	return &this
}

// NewKalturaBulkUploadXmlJobDataWithDefaults instantiates a new KalturaBulkUploadXmlJobData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadXmlJobDataWithDefaults() *KalturaBulkUploadXmlJobData {
	this := KalturaBulkUploadXmlJobData{}
	return &this
}

func (o KalturaBulkUploadXmlJobData) MarshalJSON() ([]byte, error) {
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

type NullableKalturaBulkUploadXmlJobData struct {
	value *KalturaBulkUploadXmlJobData
	isSet bool
}

func (v NullableKalturaBulkUploadXmlJobData) Get() *KalturaBulkUploadXmlJobData {
	return v.value
}

func (v *NullableKalturaBulkUploadXmlJobData) Set(val *KalturaBulkUploadXmlJobData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadXmlJobData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadXmlJobData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadXmlJobData(val *KalturaBulkUploadXmlJobData) *NullableKalturaBulkUploadXmlJobData {
	return &NullableKalturaBulkUploadXmlJobData{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadXmlJobData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadXmlJobData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


