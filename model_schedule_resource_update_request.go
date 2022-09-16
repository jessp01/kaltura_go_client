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

// ScheduleResourceUpdateRequest struct for ScheduleResourceUpdateRequest
type ScheduleResourceUpdateRequest struct {
	ScheduleResource KalturaScheduleResource `json:"scheduleResource"`
	ScheduleResourceId int32 `json:"scheduleResourceId"`
}

// NewScheduleResourceUpdateRequest instantiates a new ScheduleResourceUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleResourceUpdateRequest(scheduleResource KalturaScheduleResource, scheduleResourceId int32) *ScheduleResourceUpdateRequest {
	this := ScheduleResourceUpdateRequest{}
	this.ScheduleResource = scheduleResource
	this.ScheduleResourceId = scheduleResourceId
	return &this
}

// NewScheduleResourceUpdateRequestWithDefaults instantiates a new ScheduleResourceUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleResourceUpdateRequestWithDefaults() *ScheduleResourceUpdateRequest {
	this := ScheduleResourceUpdateRequest{}
	return &this
}

// GetScheduleResource returns the ScheduleResource field value
func (o *ScheduleResourceUpdateRequest) GetScheduleResource() KalturaScheduleResource {
	if o == nil {
		var ret KalturaScheduleResource
		return ret
	}

	return o.ScheduleResource
}

// GetScheduleResourceOk returns a tuple with the ScheduleResource field value
// and a boolean to check if the value has been set.
func (o *ScheduleResourceUpdateRequest) GetScheduleResourceOk() (*KalturaScheduleResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleResource, true
}

// SetScheduleResource sets field value
func (o *ScheduleResourceUpdateRequest) SetScheduleResource(v KalturaScheduleResource) {
	o.ScheduleResource = v
}

// GetScheduleResourceId returns the ScheduleResourceId field value
func (o *ScheduleResourceUpdateRequest) GetScheduleResourceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScheduleResourceId
}

// GetScheduleResourceIdOk returns a tuple with the ScheduleResourceId field value
// and a boolean to check if the value has been set.
func (o *ScheduleResourceUpdateRequest) GetScheduleResourceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduleResourceId, true
}

// SetScheduleResourceId sets field value
func (o *ScheduleResourceUpdateRequest) SetScheduleResourceId(v int32) {
	o.ScheduleResourceId = v
}

func (o ScheduleResourceUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scheduleResource"] = o.ScheduleResource
	}
	if true {
		toSerialize["scheduleResourceId"] = o.ScheduleResourceId
	}
	return json.Marshal(toSerialize)
}

type NullableScheduleResourceUpdateRequest struct {
	value *ScheduleResourceUpdateRequest
	isSet bool
}

func (v NullableScheduleResourceUpdateRequest) Get() *ScheduleResourceUpdateRequest {
	return v.value
}

func (v *NullableScheduleResourceUpdateRequest) Set(val *ScheduleResourceUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleResourceUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleResourceUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleResourceUpdateRequest(val *ScheduleResourceUpdateRequest) *NullableScheduleResourceUpdateRequest {
	return &NullableScheduleResourceUpdateRequest{value: val, isSet: true}
}

func (v NullableScheduleResourceUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleResourceUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


