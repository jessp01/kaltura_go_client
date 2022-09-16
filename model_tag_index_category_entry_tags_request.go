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

// TagIndexCategoryEntryTagsRequest struct for TagIndexCategoryEntryTagsRequest
type TagIndexCategoryEntryTagsRequest struct {
	CategoryId int32 `json:"categoryId"`
	PcToDecrement string `json:"pcToDecrement"`
	PcToIncrement string `json:"pcToIncrement"`
}

// NewTagIndexCategoryEntryTagsRequest instantiates a new TagIndexCategoryEntryTagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagIndexCategoryEntryTagsRequest(categoryId int32, pcToDecrement string, pcToIncrement string) *TagIndexCategoryEntryTagsRequest {
	this := TagIndexCategoryEntryTagsRequest{}
	this.CategoryId = categoryId
	this.PcToDecrement = pcToDecrement
	this.PcToIncrement = pcToIncrement
	return &this
}

// NewTagIndexCategoryEntryTagsRequestWithDefaults instantiates a new TagIndexCategoryEntryTagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagIndexCategoryEntryTagsRequestWithDefaults() *TagIndexCategoryEntryTagsRequest {
	this := TagIndexCategoryEntryTagsRequest{}
	return &this
}

// GetCategoryId returns the CategoryId field value
func (o *TagIndexCategoryEntryTagsRequest) GetCategoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *TagIndexCategoryEntryTagsRequest) GetCategoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *TagIndexCategoryEntryTagsRequest) SetCategoryId(v int32) {
	o.CategoryId = v
}

// GetPcToDecrement returns the PcToDecrement field value
func (o *TagIndexCategoryEntryTagsRequest) GetPcToDecrement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PcToDecrement
}

// GetPcToDecrementOk returns a tuple with the PcToDecrement field value
// and a boolean to check if the value has been set.
func (o *TagIndexCategoryEntryTagsRequest) GetPcToDecrementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PcToDecrement, true
}

// SetPcToDecrement sets field value
func (o *TagIndexCategoryEntryTagsRequest) SetPcToDecrement(v string) {
	o.PcToDecrement = v
}

// GetPcToIncrement returns the PcToIncrement field value
func (o *TagIndexCategoryEntryTagsRequest) GetPcToIncrement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PcToIncrement
}

// GetPcToIncrementOk returns a tuple with the PcToIncrement field value
// and a boolean to check if the value has been set.
func (o *TagIndexCategoryEntryTagsRequest) GetPcToIncrementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PcToIncrement, true
}

// SetPcToIncrement sets field value
func (o *TagIndexCategoryEntryTagsRequest) SetPcToIncrement(v string) {
	o.PcToIncrement = v
}

func (o TagIndexCategoryEntryTagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["categoryId"] = o.CategoryId
	}
	if true {
		toSerialize["pcToDecrement"] = o.PcToDecrement
	}
	if true {
		toSerialize["pcToIncrement"] = o.PcToIncrement
	}
	return json.Marshal(toSerialize)
}

type NullableTagIndexCategoryEntryTagsRequest struct {
	value *TagIndexCategoryEntryTagsRequest
	isSet bool
}

func (v NullableTagIndexCategoryEntryTagsRequest) Get() *TagIndexCategoryEntryTagsRequest {
	return v.value
}

func (v *NullableTagIndexCategoryEntryTagsRequest) Set(val *TagIndexCategoryEntryTagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTagIndexCategoryEntryTagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTagIndexCategoryEntryTagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagIndexCategoryEntryTagsRequest(val *TagIndexCategoryEntryTagsRequest) *NullableTagIndexCategoryEntryTagsRequest {
	return &NullableTagIndexCategoryEntryTagsRequest{value: val, isSet: true}
}

func (v NullableTagIndexCategoryEntryTagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagIndexCategoryEntryTagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


