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

// CategoryIndexRequest struct for CategoryIndexRequest
type CategoryIndexRequest struct {
	Id int32 `json:"id"`
	ShouldUpdate *bool `json:"shouldUpdate,omitempty"`
}

// NewCategoryIndexRequest instantiates a new CategoryIndexRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryIndexRequest(id int32) *CategoryIndexRequest {
	this := CategoryIndexRequest{}
	this.Id = id
	var shouldUpdate bool = true
	this.ShouldUpdate = &shouldUpdate
	return &this
}

// NewCategoryIndexRequestWithDefaults instantiates a new CategoryIndexRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryIndexRequestWithDefaults() *CategoryIndexRequest {
	this := CategoryIndexRequest{}
	var shouldUpdate bool = true
	this.ShouldUpdate = &shouldUpdate
	return &this
}

// GetId returns the Id field value
func (o *CategoryIndexRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryIndexRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryIndexRequest) SetId(v int32) {
	o.Id = v
}

// GetShouldUpdate returns the ShouldUpdate field value if set, zero value otherwise.
func (o *CategoryIndexRequest) GetShouldUpdate() bool {
	if o == nil || o.ShouldUpdate == nil {
		var ret bool
		return ret
	}
	return *o.ShouldUpdate
}

// GetShouldUpdateOk returns a tuple with the ShouldUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryIndexRequest) GetShouldUpdateOk() (*bool, bool) {
	if o == nil || o.ShouldUpdate == nil {
		return nil, false
	}
	return o.ShouldUpdate, true
}

// HasShouldUpdate returns a boolean if a field has been set.
func (o *CategoryIndexRequest) HasShouldUpdate() bool {
	if o != nil && o.ShouldUpdate != nil {
		return true
	}

	return false
}

// SetShouldUpdate gets a reference to the given bool and assigns it to the ShouldUpdate field.
func (o *CategoryIndexRequest) SetShouldUpdate(v bool) {
	o.ShouldUpdate = &v
}

func (o CategoryIndexRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ShouldUpdate != nil {
		toSerialize["shouldUpdate"] = o.ShouldUpdate
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryIndexRequest struct {
	value *CategoryIndexRequest
	isSet bool
}

func (v NullableCategoryIndexRequest) Get() *CategoryIndexRequest {
	return v.value
}

func (v *NullableCategoryIndexRequest) Set(val *CategoryIndexRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryIndexRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryIndexRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryIndexRequest(val *CategoryIndexRequest) *NullableCategoryIndexRequest {
	return &NullableCategoryIndexRequest{value: val, isSet: true}
}

func (v NullableCategoryIndexRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryIndexRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


