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

// BaseEntryAddContentRequest struct for BaseEntryAddContentRequest
type BaseEntryAddContentRequest struct {
	EntryId string `json:"entryId"`
	Resource KalturaResource `json:"resource"`
}

// NewBaseEntryAddContentRequest instantiates a new BaseEntryAddContentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntryAddContentRequest(entryId string, resource KalturaResource) *BaseEntryAddContentRequest {
	this := BaseEntryAddContentRequest{}
	this.EntryId = entryId
	this.Resource = resource
	return &this
}

// NewBaseEntryAddContentRequestWithDefaults instantiates a new BaseEntryAddContentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntryAddContentRequestWithDefaults() *BaseEntryAddContentRequest {
	this := BaseEntryAddContentRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *BaseEntryAddContentRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *BaseEntryAddContentRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *BaseEntryAddContentRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetResource returns the Resource field value
func (o *BaseEntryAddContentRequest) GetResource() KalturaResource {
	if o == nil {
		var ret KalturaResource
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *BaseEntryAddContentRequest) GetResourceOk() (*KalturaResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *BaseEntryAddContentRequest) SetResource(v KalturaResource) {
	o.Resource = v
}

func (o BaseEntryAddContentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if true {
		toSerialize["resource"] = o.Resource
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEntryAddContentRequest struct {
	value *BaseEntryAddContentRequest
	isSet bool
}

func (v NullableBaseEntryAddContentRequest) Get() *BaseEntryAddContentRequest {
	return v.value
}

func (v *NullableBaseEntryAddContentRequest) Set(val *BaseEntryAddContentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntryAddContentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntryAddContentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntryAddContentRequest(val *BaseEntryAddContentRequest) *NullableBaseEntryAddContentRequest {
	return &NullableBaseEntryAddContentRequest{value: val, isSet: true}
}

func (v NullableBaseEntryAddContentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntryAddContentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


