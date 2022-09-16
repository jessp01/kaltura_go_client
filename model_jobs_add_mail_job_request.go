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

// JobsAddMailJobRequest struct for JobsAddMailJobRequest
type JobsAddMailJobRequest struct {
	MailJobData KalturaMailJobData `json:"mailJobData"`
}

// NewJobsAddMailJobRequest instantiates a new JobsAddMailJobRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobsAddMailJobRequest(mailJobData KalturaMailJobData) *JobsAddMailJobRequest {
	this := JobsAddMailJobRequest{}
	this.MailJobData = mailJobData
	return &this
}

// NewJobsAddMailJobRequestWithDefaults instantiates a new JobsAddMailJobRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobsAddMailJobRequestWithDefaults() *JobsAddMailJobRequest {
	this := JobsAddMailJobRequest{}
	return &this
}

// GetMailJobData returns the MailJobData field value
func (o *JobsAddMailJobRequest) GetMailJobData() KalturaMailJobData {
	if o == nil {
		var ret KalturaMailJobData
		return ret
	}

	return o.MailJobData
}

// GetMailJobDataOk returns a tuple with the MailJobData field value
// and a boolean to check if the value has been set.
func (o *JobsAddMailJobRequest) GetMailJobDataOk() (*KalturaMailJobData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MailJobData, true
}

// SetMailJobData sets field value
func (o *JobsAddMailJobRequest) SetMailJobData(v KalturaMailJobData) {
	o.MailJobData = v
}

func (o JobsAddMailJobRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["mailJobData"] = o.MailJobData
	}
	return json.Marshal(toSerialize)
}

type NullableJobsAddMailJobRequest struct {
	value *JobsAddMailJobRequest
	isSet bool
}

func (v NullableJobsAddMailJobRequest) Get() *JobsAddMailJobRequest {
	return v.value
}

func (v *NullableJobsAddMailJobRequest) Set(val *JobsAddMailJobRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableJobsAddMailJobRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableJobsAddMailJobRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobsAddMailJobRequest(val *JobsAddMailJobRequest) *NullableJobsAddMailJobRequest {
	return &NullableJobsAddMailJobRequest{value: val, isSet: true}
}

func (v NullableJobsAddMailJobRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobsAddMailJobRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


