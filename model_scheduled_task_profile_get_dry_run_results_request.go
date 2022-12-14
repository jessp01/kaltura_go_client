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

// ScheduledTaskProfileGetDryRunResultsRequest struct for ScheduledTaskProfileGetDryRunResultsRequest
type ScheduledTaskProfileGetDryRunResultsRequest struct {
	RequestId int32 `json:"requestId"`
}

// NewScheduledTaskProfileGetDryRunResultsRequest instantiates a new ScheduledTaskProfileGetDryRunResultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledTaskProfileGetDryRunResultsRequest(requestId int32) *ScheduledTaskProfileGetDryRunResultsRequest {
	this := ScheduledTaskProfileGetDryRunResultsRequest{}
	this.RequestId = requestId
	return &this
}

// NewScheduledTaskProfileGetDryRunResultsRequestWithDefaults instantiates a new ScheduledTaskProfileGetDryRunResultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledTaskProfileGetDryRunResultsRequestWithDefaults() *ScheduledTaskProfileGetDryRunResultsRequest {
	this := ScheduledTaskProfileGetDryRunResultsRequest{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ScheduledTaskProfileGetDryRunResultsRequest) GetRequestId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskProfileGetDryRunResultsRequest) GetRequestIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ScheduledTaskProfileGetDryRunResultsRequest) SetRequestId(v int32) {
	o.RequestId = v
}

func (o ScheduledTaskProfileGetDryRunResultsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	return json.Marshal(toSerialize)
}

type NullableScheduledTaskProfileGetDryRunResultsRequest struct {
	value *ScheduledTaskProfileGetDryRunResultsRequest
	isSet bool
}

func (v NullableScheduledTaskProfileGetDryRunResultsRequest) Get() *ScheduledTaskProfileGetDryRunResultsRequest {
	return v.value
}

func (v *NullableScheduledTaskProfileGetDryRunResultsRequest) Set(val *ScheduledTaskProfileGetDryRunResultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledTaskProfileGetDryRunResultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledTaskProfileGetDryRunResultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledTaskProfileGetDryRunResultsRequest(val *ScheduledTaskProfileGetDryRunResultsRequest) *NullableScheduledTaskProfileGetDryRunResultsRequest {
	return &NullableScheduledTaskProfileGetDryRunResultsRequest{value: val, isSet: true}
}

func (v NullableScheduledTaskProfileGetDryRunResultsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledTaskProfileGetDryRunResultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


