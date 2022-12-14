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

// ScheduledTaskProfileRequestDryRunRequest struct for ScheduledTaskProfileRequestDryRunRequest
type ScheduledTaskProfileRequestDryRunRequest struct {
	MaxResults *int32 `json:"maxResults,omitempty"`
	ScheduledTaskProfileId int32 `json:"scheduledTaskProfileId"`
}

// NewScheduledTaskProfileRequestDryRunRequest instantiates a new ScheduledTaskProfileRequestDryRunRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduledTaskProfileRequestDryRunRequest(scheduledTaskProfileId int32) *ScheduledTaskProfileRequestDryRunRequest {
	this := ScheduledTaskProfileRequestDryRunRequest{}
	this.ScheduledTaskProfileId = scheduledTaskProfileId
	return &this
}

// NewScheduledTaskProfileRequestDryRunRequestWithDefaults instantiates a new ScheduledTaskProfileRequestDryRunRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduledTaskProfileRequestDryRunRequestWithDefaults() *ScheduledTaskProfileRequestDryRunRequest {
	this := ScheduledTaskProfileRequestDryRunRequest{}
	return &this
}

// GetMaxResults returns the MaxResults field value if set, zero value otherwise.
func (o *ScheduledTaskProfileRequestDryRunRequest) GetMaxResults() int32 {
	if o == nil || o.MaxResults == nil {
		var ret int32
		return ret
	}
	return *o.MaxResults
}

// GetMaxResultsOk returns a tuple with the MaxResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduledTaskProfileRequestDryRunRequest) GetMaxResultsOk() (*int32, bool) {
	if o == nil || o.MaxResults == nil {
		return nil, false
	}
	return o.MaxResults, true
}

// HasMaxResults returns a boolean if a field has been set.
func (o *ScheduledTaskProfileRequestDryRunRequest) HasMaxResults() bool {
	if o != nil && o.MaxResults != nil {
		return true
	}

	return false
}

// SetMaxResults gets a reference to the given int32 and assigns it to the MaxResults field.
func (o *ScheduledTaskProfileRequestDryRunRequest) SetMaxResults(v int32) {
	o.MaxResults = &v
}

// GetScheduledTaskProfileId returns the ScheduledTaskProfileId field value
func (o *ScheduledTaskProfileRequestDryRunRequest) GetScheduledTaskProfileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScheduledTaskProfileId
}

// GetScheduledTaskProfileIdOk returns a tuple with the ScheduledTaskProfileId field value
// and a boolean to check if the value has been set.
func (o *ScheduledTaskProfileRequestDryRunRequest) GetScheduledTaskProfileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledTaskProfileId, true
}

// SetScheduledTaskProfileId sets field value
func (o *ScheduledTaskProfileRequestDryRunRequest) SetScheduledTaskProfileId(v int32) {
	o.ScheduledTaskProfileId = v
}

func (o ScheduledTaskProfileRequestDryRunRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxResults != nil {
		toSerialize["maxResults"] = o.MaxResults
	}
	if true {
		toSerialize["scheduledTaskProfileId"] = o.ScheduledTaskProfileId
	}
	return json.Marshal(toSerialize)
}

type NullableScheduledTaskProfileRequestDryRunRequest struct {
	value *ScheduledTaskProfileRequestDryRunRequest
	isSet bool
}

func (v NullableScheduledTaskProfileRequestDryRunRequest) Get() *ScheduledTaskProfileRequestDryRunRequest {
	return v.value
}

func (v *NullableScheduledTaskProfileRequestDryRunRequest) Set(val *ScheduledTaskProfileRequestDryRunRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledTaskProfileRequestDryRunRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledTaskProfileRequestDryRunRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledTaskProfileRequestDryRunRequest(val *ScheduledTaskProfileRequestDryRunRequest) *NullableScheduledTaskProfileRequestDryRunRequest {
	return &NullableScheduledTaskProfileRequestDryRunRequest{value: val, isSet: true}
}

func (v NullableScheduledTaskProfileRequestDryRunRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledTaskProfileRequestDryRunRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


