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

// BaseEntryCloneRequest struct for BaseEntryCloneRequest
type BaseEntryCloneRequest struct {
	CloneOptions []KalturaBaseEntryCloneOptionItem `json:"cloneOptions,omitempty"`
	EntryId string `json:"entryId"`
}

// NewBaseEntryCloneRequest instantiates a new BaseEntryCloneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntryCloneRequest(entryId string) *BaseEntryCloneRequest {
	this := BaseEntryCloneRequest{}
	this.EntryId = entryId
	return &this
}

// NewBaseEntryCloneRequestWithDefaults instantiates a new BaseEntryCloneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntryCloneRequestWithDefaults() *BaseEntryCloneRequest {
	this := BaseEntryCloneRequest{}
	return &this
}

// GetCloneOptions returns the CloneOptions field value if set, zero value otherwise.
func (o *BaseEntryCloneRequest) GetCloneOptions() []KalturaBaseEntryCloneOptionItem {
	if o == nil || o.CloneOptions == nil {
		var ret []KalturaBaseEntryCloneOptionItem
		return ret
	}
	return o.CloneOptions
}

// GetCloneOptionsOk returns a tuple with the CloneOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntryCloneRequest) GetCloneOptionsOk() ([]KalturaBaseEntryCloneOptionItem, bool) {
	if o == nil || o.CloneOptions == nil {
		return nil, false
	}
	return o.CloneOptions, true
}

// HasCloneOptions returns a boolean if a field has been set.
func (o *BaseEntryCloneRequest) HasCloneOptions() bool {
	if o != nil && o.CloneOptions != nil {
		return true
	}

	return false
}

// SetCloneOptions gets a reference to the given []KalturaBaseEntryCloneOptionItem and assigns it to the CloneOptions field.
func (o *BaseEntryCloneRequest) SetCloneOptions(v []KalturaBaseEntryCloneOptionItem) {
	o.CloneOptions = v
}

// GetEntryId returns the EntryId field value
func (o *BaseEntryCloneRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *BaseEntryCloneRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *BaseEntryCloneRequest) SetEntryId(v string) {
	o.EntryId = v
}

func (o BaseEntryCloneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloneOptions != nil {
		toSerialize["cloneOptions"] = o.CloneOptions
	}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEntryCloneRequest struct {
	value *BaseEntryCloneRequest
	isSet bool
}

func (v NullableBaseEntryCloneRequest) Get() *BaseEntryCloneRequest {
	return v.value
}

func (v *NullableBaseEntryCloneRequest) Set(val *BaseEntryCloneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntryCloneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntryCloneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntryCloneRequest(val *BaseEntryCloneRequest) *NullableBaseEntryCloneRequest {
	return &NullableBaseEntryCloneRequest{value: val, isSet: true}
}

func (v NullableBaseEntryCloneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntryCloneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


