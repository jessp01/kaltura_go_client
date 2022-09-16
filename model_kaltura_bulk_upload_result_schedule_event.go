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

// KalturaBulkUploadResultScheduleEvent struct for KalturaBulkUploadResultScheduleEvent
type KalturaBulkUploadResultScheduleEvent struct {
	KalturaBulkUploadResult
}

// NewKalturaBulkUploadResultScheduleEvent instantiates a new KalturaBulkUploadResultScheduleEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaBulkUploadResultScheduleEvent() *KalturaBulkUploadResultScheduleEvent {
	this := KalturaBulkUploadResultScheduleEvent{}
	return &this
}

// NewKalturaBulkUploadResultScheduleEventWithDefaults instantiates a new KalturaBulkUploadResultScheduleEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaBulkUploadResultScheduleEventWithDefaults() *KalturaBulkUploadResultScheduleEvent {
	this := KalturaBulkUploadResultScheduleEvent{}
	return &this
}

func (o KalturaBulkUploadResultScheduleEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedKalturaBulkUploadResult, errKalturaBulkUploadResult := json.Marshal(o.KalturaBulkUploadResult)
	if errKalturaBulkUploadResult != nil {
		return []byte{}, errKalturaBulkUploadResult
	}
	errKalturaBulkUploadResult = json.Unmarshal([]byte(serializedKalturaBulkUploadResult), &toSerialize)
	if errKalturaBulkUploadResult != nil {
		return []byte{}, errKalturaBulkUploadResult
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaBulkUploadResultScheduleEvent struct {
	value *KalturaBulkUploadResultScheduleEvent
	isSet bool
}

func (v NullableKalturaBulkUploadResultScheduleEvent) Get() *KalturaBulkUploadResultScheduleEvent {
	return v.value
}

func (v *NullableKalturaBulkUploadResultScheduleEvent) Set(val *KalturaBulkUploadResultScheduleEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaBulkUploadResultScheduleEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaBulkUploadResultScheduleEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaBulkUploadResultScheduleEvent(val *KalturaBulkUploadResultScheduleEvent) *NullableKalturaBulkUploadResultScheduleEvent {
	return &NullableKalturaBulkUploadResultScheduleEvent{value: val, isSet: true}
}

func (v NullableKalturaBulkUploadResultScheduleEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaBulkUploadResultScheduleEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


