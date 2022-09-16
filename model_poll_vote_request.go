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

// PollVoteRequest struct for PollVoteRequest
type PollVoteRequest struct {
	AnswerIds string `json:"answerIds"`
	PollId string `json:"pollId"`
	UserId string `json:"userId"`
}

// NewPollVoteRequest instantiates a new PollVoteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollVoteRequest(answerIds string, pollId string, userId string) *PollVoteRequest {
	this := PollVoteRequest{}
	this.AnswerIds = answerIds
	this.PollId = pollId
	this.UserId = userId
	return &this
}

// NewPollVoteRequestWithDefaults instantiates a new PollVoteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollVoteRequestWithDefaults() *PollVoteRequest {
	this := PollVoteRequest{}
	return &this
}

// GetAnswerIds returns the AnswerIds field value
func (o *PollVoteRequest) GetAnswerIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnswerIds
}

// GetAnswerIdsOk returns a tuple with the AnswerIds field value
// and a boolean to check if the value has been set.
func (o *PollVoteRequest) GetAnswerIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnswerIds, true
}

// SetAnswerIds sets field value
func (o *PollVoteRequest) SetAnswerIds(v string) {
	o.AnswerIds = v
}

// GetPollId returns the PollId field value
func (o *PollVoteRequest) GetPollId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PollId
}

// GetPollIdOk returns a tuple with the PollId field value
// and a boolean to check if the value has been set.
func (o *PollVoteRequest) GetPollIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PollId, true
}

// SetPollId sets field value
func (o *PollVoteRequest) SetPollId(v string) {
	o.PollId = v
}

// GetUserId returns the UserId field value
func (o *PollVoteRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PollVoteRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PollVoteRequest) SetUserId(v string) {
	o.UserId = v
}

func (o PollVoteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["answerIds"] = o.AnswerIds
	}
	if true {
		toSerialize["pollId"] = o.PollId
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullablePollVoteRequest struct {
	value *PollVoteRequest
	isSet bool
}

func (v NullablePollVoteRequest) Get() *PollVoteRequest {
	return v.value
}

func (v *NullablePollVoteRequest) Set(val *PollVoteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePollVoteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePollVoteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollVoteRequest(val *PollVoteRequest) *NullablePollVoteRequest {
	return &NullablePollVoteRequest{value: val, isSet: true}
}

func (v NullablePollVoteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollVoteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


