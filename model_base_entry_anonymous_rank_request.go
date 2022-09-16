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

// BaseEntryAnonymousRankRequest struct for BaseEntryAnonymousRankRequest
type BaseEntryAnonymousRankRequest struct {
	EntryId string `json:"entryId"`
	Rank int32 `json:"rank"`
}

// NewBaseEntryAnonymousRankRequest instantiates a new BaseEntryAnonymousRankRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntryAnonymousRankRequest(entryId string, rank int32) *BaseEntryAnonymousRankRequest {
	this := BaseEntryAnonymousRankRequest{}
	this.EntryId = entryId
	this.Rank = rank
	return &this
}

// NewBaseEntryAnonymousRankRequestWithDefaults instantiates a new BaseEntryAnonymousRankRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntryAnonymousRankRequestWithDefaults() *BaseEntryAnonymousRankRequest {
	this := BaseEntryAnonymousRankRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *BaseEntryAnonymousRankRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *BaseEntryAnonymousRankRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *BaseEntryAnonymousRankRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetRank returns the Rank field value
func (o *BaseEntryAnonymousRankRequest) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *BaseEntryAnonymousRankRequest) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *BaseEntryAnonymousRankRequest) SetRank(v int32) {
	o.Rank = v
}

func (o BaseEntryAnonymousRankRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["rank"] = o.Rank
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEntryAnonymousRankRequest struct {
	value *BaseEntryAnonymousRankRequest
	isSet bool
}

func (v NullableBaseEntryAnonymousRankRequest) Get() *BaseEntryAnonymousRankRequest {
	return v.value
}

func (v *NullableBaseEntryAnonymousRankRequest) Set(val *BaseEntryAnonymousRankRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntryAnonymousRankRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntryAnonymousRankRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntryAnonymousRankRequest(val *BaseEntryAnonymousRankRequest) *NullableBaseEntryAnonymousRankRequest {
	return &NullableBaseEntryAnonymousRankRequest{value: val, isSet: true}
}

func (v NullableBaseEntryAnonymousRankRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntryAnonymousRankRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


