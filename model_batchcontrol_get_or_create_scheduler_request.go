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

// BatchcontrolGetOrCreateSchedulerRequest struct for BatchcontrolGetOrCreateSchedulerRequest
type BatchcontrolGetOrCreateSchedulerRequest struct {
	Scheduler KalturaScheduler `json:"scheduler"`
}

// NewBatchcontrolGetOrCreateSchedulerRequest instantiates a new BatchcontrolGetOrCreateSchedulerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchcontrolGetOrCreateSchedulerRequest(scheduler KalturaScheduler) *BatchcontrolGetOrCreateSchedulerRequest {
	this := BatchcontrolGetOrCreateSchedulerRequest{}
	this.Scheduler = scheduler
	return &this
}

// NewBatchcontrolGetOrCreateSchedulerRequestWithDefaults instantiates a new BatchcontrolGetOrCreateSchedulerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchcontrolGetOrCreateSchedulerRequestWithDefaults() *BatchcontrolGetOrCreateSchedulerRequest {
	this := BatchcontrolGetOrCreateSchedulerRequest{}
	return &this
}

// GetScheduler returns the Scheduler field value
func (o *BatchcontrolGetOrCreateSchedulerRequest) GetScheduler() KalturaScheduler {
	if o == nil {
		var ret KalturaScheduler
		return ret
	}

	return o.Scheduler
}

// GetSchedulerOk returns a tuple with the Scheduler field value
// and a boolean to check if the value has been set.
func (o *BatchcontrolGetOrCreateSchedulerRequest) GetSchedulerOk() (*KalturaScheduler, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheduler, true
}

// SetScheduler sets field value
func (o *BatchcontrolGetOrCreateSchedulerRequest) SetScheduler(v KalturaScheduler) {
	o.Scheduler = v
}

func (o BatchcontrolGetOrCreateSchedulerRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["scheduler"] = o.Scheduler
	}
	return json.Marshal(toSerialize)
}

type NullableBatchcontrolGetOrCreateSchedulerRequest struct {
	value *BatchcontrolGetOrCreateSchedulerRequest
	isSet bool
}

func (v NullableBatchcontrolGetOrCreateSchedulerRequest) Get() *BatchcontrolGetOrCreateSchedulerRequest {
	return v.value
}

func (v *NullableBatchcontrolGetOrCreateSchedulerRequest) Set(val *BatchcontrolGetOrCreateSchedulerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchcontrolGetOrCreateSchedulerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchcontrolGetOrCreateSchedulerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchcontrolGetOrCreateSchedulerRequest(val *BatchcontrolGetOrCreateSchedulerRequest) *NullableBatchcontrolGetOrCreateSchedulerRequest {
	return &NullableBatchcontrolGetOrCreateSchedulerRequest{value: val, isSet: true}
}

func (v NullableBatchcontrolGetOrCreateSchedulerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchcontrolGetOrCreateSchedulerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


