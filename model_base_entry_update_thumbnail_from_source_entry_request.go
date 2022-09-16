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

// BaseEntryUpdateThumbnailFromSourceEntryRequest struct for BaseEntryUpdateThumbnailFromSourceEntryRequest
type BaseEntryUpdateThumbnailFromSourceEntryRequest struct {
	EntryId string `json:"entryId"`
	SourceEntryId string `json:"sourceEntryId"`
	TimeOffset int32 `json:"timeOffset"`
}

// NewBaseEntryUpdateThumbnailFromSourceEntryRequest instantiates a new BaseEntryUpdateThumbnailFromSourceEntryRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntryUpdateThumbnailFromSourceEntryRequest(entryId string, sourceEntryId string, timeOffset int32) *BaseEntryUpdateThumbnailFromSourceEntryRequest {
	this := BaseEntryUpdateThumbnailFromSourceEntryRequest{}
	this.EntryId = entryId
	this.SourceEntryId = sourceEntryId
	this.TimeOffset = timeOffset
	return &this
}

// NewBaseEntryUpdateThumbnailFromSourceEntryRequestWithDefaults instantiates a new BaseEntryUpdateThumbnailFromSourceEntryRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntryUpdateThumbnailFromSourceEntryRequestWithDefaults() *BaseEntryUpdateThumbnailFromSourceEntryRequest {
	this := BaseEntryUpdateThumbnailFromSourceEntryRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetSourceEntryId returns the SourceEntryId field value
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) GetSourceEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceEntryId
}

// GetSourceEntryIdOk returns a tuple with the SourceEntryId field value
// and a boolean to check if the value has been set.
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) GetSourceEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceEntryId, true
}

// SetSourceEntryId sets field value
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) SetSourceEntryId(v string) {
	o.SourceEntryId = v
}

// GetTimeOffset returns the TimeOffset field value
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) GetTimeOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimeOffset
}

// GetTimeOffsetOk returns a tuple with the TimeOffset field value
// and a boolean to check if the value has been set.
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) GetTimeOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOffset, true
}

// SetTimeOffset sets field value
func (o *BaseEntryUpdateThumbnailFromSourceEntryRequest) SetTimeOffset(v int32) {
	o.TimeOffset = v
}

func (o BaseEntryUpdateThumbnailFromSourceEntryRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["sourceEntryId"] = o.SourceEntryId
	}
	if true {
		toSerialize["timeOffset"] = o.TimeOffset
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEntryUpdateThumbnailFromSourceEntryRequest struct {
	value *BaseEntryUpdateThumbnailFromSourceEntryRequest
	isSet bool
}

func (v NullableBaseEntryUpdateThumbnailFromSourceEntryRequest) Get() *BaseEntryUpdateThumbnailFromSourceEntryRequest {
	return v.value
}

func (v *NullableBaseEntryUpdateThumbnailFromSourceEntryRequest) Set(val *BaseEntryUpdateThumbnailFromSourceEntryRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntryUpdateThumbnailFromSourceEntryRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntryUpdateThumbnailFromSourceEntryRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntryUpdateThumbnailFromSourceEntryRequest(val *BaseEntryUpdateThumbnailFromSourceEntryRequest) *NullableBaseEntryUpdateThumbnailFromSourceEntryRequest {
	return &NullableBaseEntryUpdateThumbnailFromSourceEntryRequest{value: val, isSet: true}
}

func (v NullableBaseEntryUpdateThumbnailFromSourceEntryRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntryUpdateThumbnailFromSourceEntryRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

