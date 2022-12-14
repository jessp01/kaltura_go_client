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

// ScheduleResourceDeleteRequest struct for ScheduleResourceDeleteRequest
type ScheduleResourceDeleteRequest struct {
	ScheduleResourceId int32 `json:"scheduleResourceId"`
}

// NewScheduleResourceDeleteRequest instantiates a new ScheduleResourceDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleResourceDeleteRequest(scheduleResourceId int32) *ScheduleResourceDeleteRequest {
	this := ScheduleResourceDeleteRequest{}
	this.ScheduleResourceId = scheduleResourceId
	return &this
}

// NewScheduleResourceDeleteRequestWithDefaults instantiates a new ScheduleResourceDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleResourceDeleteRequestWithDefaults() *ScheduleResourceDeleteRequest {
	this := ScheduleResourceDeleteRequest{}
	return &this
}

// GetScheduleResourceId returns the ScheduleResourceId field value
func (o *ScheduleResourceDeleteRequest) GetScheduleResourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScheduleResourceId
}

// GetScheduleResourceIdOk returns a tuple with the ScheduleResourceId field value
// and a boolean to check if the value has been set.
func (o *ScheduleResourceDeleteRequest) GetScheduleResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleResourceId, true
}

// SetScheduleResourceId sets field value
func (o *ScheduleResourceDeleteRequest) SetScheduleResourceId(v int32) {
	o.ScheduleResourceId = v
}

func (o ScheduleResourceDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scheduleResourceId"] = o.ScheduleResourceId
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleResourceDeleteRequest struct {
	value *ScheduleResourceDeleteRequest
	isSet bool
}

func (v NullableScheduleResourceDeleteRequest) Get() *ScheduleResourceDeleteRequest {
	return v.value
}

func (v *NullableScheduleResourceDeleteRequest) Set(val *ScheduleResourceDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleResourceDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleResourceDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleResourceDeleteRequest(val *ScheduleResourceDeleteRequest) *NullableScheduleResourceDeleteRequest {
	return &NullableScheduleResourceDeleteRequest{value: val, isSet: true}
}

func (v NullableScheduleResourceDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleResourceDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


