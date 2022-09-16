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

// EntryDistributionSubmitAddRequest struct for EntryDistributionSubmitAddRequest
type EntryDistributionSubmitAddRequest struct {
	Id int32 `json:"id"`
	SubmitWhenReady *bool `json:"submitWhenReady,omitempty"`
}

// NewEntryDistributionSubmitAddRequest instantiates a new EntryDistributionSubmitAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntryDistributionSubmitAddRequest(id int32) *EntryDistributionSubmitAddRequest {
	this := EntryDistributionSubmitAddRequest{}
	this.Id = id
	var submitWhenReady bool = false
	this.SubmitWhenReady = &submitWhenReady
	return &this
}

// NewEntryDistributionSubmitAddRequestWithDefaults instantiates a new EntryDistributionSubmitAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntryDistributionSubmitAddRequestWithDefaults() *EntryDistributionSubmitAddRequest {
	this := EntryDistributionSubmitAddRequest{}
	var submitWhenReady bool = false
	this.SubmitWhenReady = &submitWhenReady
	return &this
}

// GetId returns the Id field value
func (o *EntryDistributionSubmitAddRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntryDistributionSubmitAddRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntryDistributionSubmitAddRequest) SetId(v int32) {
	o.Id = v
}

// GetSubmitWhenReady returns the SubmitWhenReady field value if set, zero value otherwise.
func (o *EntryDistributionSubmitAddRequest) GetSubmitWhenReady() bool {
	if o == nil || o.SubmitWhenReady == nil {
		var ret bool
		return ret
	}
	return *o.SubmitWhenReady
}

// GetSubmitWhenReadyOk returns a tuple with the SubmitWhenReady field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntryDistributionSubmitAddRequest) GetSubmitWhenReadyOk() (*bool, bool) {
	if o == nil || o.SubmitWhenReady == nil {
		return nil, false
	}
	return o.SubmitWhenReady, true
}

// HasSubmitWhenReady returns a boolean if a field has been set.
func (o *EntryDistributionSubmitAddRequest) HasSubmitWhenReady() bool {
	if o != nil && o.SubmitWhenReady != nil {
		return true
	}

	return false
}

// SetSubmitWhenReady gets a reference to the given bool and assigns it to the SubmitWhenReady field.
func (o *EntryDistributionSubmitAddRequest) SetSubmitWhenReady(v bool) {
	o.SubmitWhenReady = &v
}

func (o EntryDistributionSubmitAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.SubmitWhenReady != nil {
		toSerialize["submitWhenReady"] = o.SubmitWhenReady
	}
	return json.Marshal(toSerialize)
}

type NullableEntryDistributionSubmitAddRequest struct {
	value *EntryDistributionSubmitAddRequest
	isSet bool
}

func (v NullableEntryDistributionSubmitAddRequest) Get() *EntryDistributionSubmitAddRequest {
	return v.value
}

func (v *NullableEntryDistributionSubmitAddRequest) Set(val *EntryDistributionSubmitAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEntryDistributionSubmitAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEntryDistributionSubmitAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntryDistributionSubmitAddRequest(val *EntryDistributionSubmitAddRequest) *NullableEntryDistributionSubmitAddRequest {
	return &NullableEntryDistributionSubmitAddRequest{value: val, isSet: true}
}

func (v NullableEntryDistributionSubmitAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntryDistributionSubmitAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


