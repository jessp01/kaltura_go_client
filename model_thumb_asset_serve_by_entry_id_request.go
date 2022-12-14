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

// ThumbAssetServeByEntryIdRequest struct for ThumbAssetServeByEntryIdRequest
type ThumbAssetServeByEntryIdRequest struct {
	EntryId string `json:"entryId"`
	ThumbParamId *int32 `json:"thumbParamId,omitempty"`
}

// NewThumbAssetServeByEntryIdRequest instantiates a new ThumbAssetServeByEntryIdRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThumbAssetServeByEntryIdRequest(entryId string) *ThumbAssetServeByEntryIdRequest {
	this := ThumbAssetServeByEntryIdRequest{}
	this.EntryId = entryId
	return &this
}

// NewThumbAssetServeByEntryIdRequestWithDefaults instantiates a new ThumbAssetServeByEntryIdRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThumbAssetServeByEntryIdRequestWithDefaults() *ThumbAssetServeByEntryIdRequest {
	this := ThumbAssetServeByEntryIdRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *ThumbAssetServeByEntryIdRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *ThumbAssetServeByEntryIdRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *ThumbAssetServeByEntryIdRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetThumbParamId returns the ThumbParamId field value if set, zero value otherwise.
func (o *ThumbAssetServeByEntryIdRequest) GetThumbParamId() int32 {
	if o == nil || o.ThumbParamId == nil {
		var ret int32
		return ret
	}
	return *o.ThumbParamId
}

// GetThumbParamIdOk returns a tuple with the ThumbParamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThumbAssetServeByEntryIdRequest) GetThumbParamIdOk() (*int32, bool) {
	if o == nil || o.ThumbParamId == nil {
		return nil, false
	}
	return o.ThumbParamId, true
}

// HasThumbParamId returns a boolean if a field has been set.
func (o *ThumbAssetServeByEntryIdRequest) HasThumbParamId() bool {
	if o != nil && o.ThumbParamId != nil {
		return true
	}

	return false
}

// SetThumbParamId gets a reference to the given int32 and assigns it to the ThumbParamId field.
func (o *ThumbAssetServeByEntryIdRequest) SetThumbParamId(v int32) {
	o.ThumbParamId = &v
}

func (o ThumbAssetServeByEntryIdRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if o.ThumbParamId != nil {
		toSerialize["thumbParamId"] = o.ThumbParamId
	}
	return json.Marshal(toSerialize)
}

type NullableThumbAssetServeByEntryIdRequest struct {
	value *ThumbAssetServeByEntryIdRequest
	isSet bool
}

func (v NullableThumbAssetServeByEntryIdRequest) Get() *ThumbAssetServeByEntryIdRequest {
	return v.value
}

func (v *NullableThumbAssetServeByEntryIdRequest) Set(val *ThumbAssetServeByEntryIdRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableThumbAssetServeByEntryIdRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableThumbAssetServeByEntryIdRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThumbAssetServeByEntryIdRequest(val *ThumbAssetServeByEntryIdRequest) *NullableThumbAssetServeByEntryIdRequest {
	return &NullableThumbAssetServeByEntryIdRequest{value: val, isSet: true}
}

func (v NullableThumbAssetServeByEntryIdRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThumbAssetServeByEntryIdRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


