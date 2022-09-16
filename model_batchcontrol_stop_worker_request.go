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

// BatchcontrolStopWorkerRequest struct for BatchcontrolStopWorkerRequest
type BatchcontrolStopWorkerRequest struct {
	AdminId int32 `json:"adminId"`
	Cause string `json:"cause"`
	WorkerId int32 `json:"workerId"`
}

// NewBatchcontrolStopWorkerRequest instantiates a new BatchcontrolStopWorkerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchcontrolStopWorkerRequest(adminId int32, cause string, workerId int32) *BatchcontrolStopWorkerRequest {
	this := BatchcontrolStopWorkerRequest{}
	this.AdminId = adminId
	this.Cause = cause
	this.WorkerId = workerId
	return &this
}

// NewBatchcontrolStopWorkerRequestWithDefaults instantiates a new BatchcontrolStopWorkerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchcontrolStopWorkerRequestWithDefaults() *BatchcontrolStopWorkerRequest {
	this := BatchcontrolStopWorkerRequest{}
	return &this
}

// GetAdminId returns the AdminId field value
func (o *BatchcontrolStopWorkerRequest) GetAdminId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AdminId
}

// GetAdminIdOk returns a tuple with the AdminId field value
// and a boolean to check if the value has been set.
func (o *BatchcontrolStopWorkerRequest) GetAdminIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdminId, true
}

// SetAdminId sets field value
func (o *BatchcontrolStopWorkerRequest) SetAdminId(v int32) {
	o.AdminId = v
}

// GetCause returns the Cause field value
func (o *BatchcontrolStopWorkerRequest) GetCause() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cause
}

// GetCauseOk returns a tuple with the Cause field value
// and a boolean to check if the value has been set.
func (o *BatchcontrolStopWorkerRequest) GetCauseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cause, true
}

// SetCause sets field value
func (o *BatchcontrolStopWorkerRequest) SetCause(v string) {
	o.Cause = v
}

// GetWorkerId returns the WorkerId field value
func (o *BatchcontrolStopWorkerRequest) GetWorkerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *BatchcontrolStopWorkerRequest) GetWorkerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *BatchcontrolStopWorkerRequest) SetWorkerId(v int32) {
	o.WorkerId = v
}

func (o BatchcontrolStopWorkerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adminId"] = o.AdminId
	}
	if true {
		toSerialize["cause"] = o.Cause
	}
	if true {
		toSerialize["workerId"] = o.WorkerId
	}
	return json.Marshal(toSerialize)
}

type NullableBatchcontrolStopWorkerRequest struct {
	value *BatchcontrolStopWorkerRequest
	isSet bool
}

func (v NullableBatchcontrolStopWorkerRequest) Get() *BatchcontrolStopWorkerRequest {
	return v.value
}

func (v *NullableBatchcontrolStopWorkerRequest) Set(val *BatchcontrolStopWorkerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchcontrolStopWorkerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchcontrolStopWorkerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchcontrolStopWorkerRequest(val *BatchcontrolStopWorkerRequest) *NullableBatchcontrolStopWorkerRequest {
	return &NullableBatchcontrolStopWorkerRequest{value: val, isSet: true}
}

func (v NullableBatchcontrolStopWorkerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchcontrolStopWorkerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


