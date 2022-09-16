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

// CategoryUserIndexRequest struct for CategoryUserIndexRequest
type CategoryUserIndexRequest struct {
	CategoryId int32 `json:"categoryId"`
	ShouldUpdate *bool `json:"shouldUpdate,omitempty"`
	UserId string `json:"userId"`
}

// NewCategoryUserIndexRequest instantiates a new CategoryUserIndexRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryUserIndexRequest(categoryId int32, userId string) *CategoryUserIndexRequest {
	this := CategoryUserIndexRequest{}
	this.CategoryId = categoryId
	var shouldUpdate bool = true
	this.ShouldUpdate = &shouldUpdate
	this.UserId = userId
	return &this
}

// NewCategoryUserIndexRequestWithDefaults instantiates a new CategoryUserIndexRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryUserIndexRequestWithDefaults() *CategoryUserIndexRequest {
	this := CategoryUserIndexRequest{}
	var shouldUpdate bool = true
	this.ShouldUpdate = &shouldUpdate
	return &this
}

// GetCategoryId returns the CategoryId field value
func (o *CategoryUserIndexRequest) GetCategoryId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *CategoryUserIndexRequest) GetCategoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *CategoryUserIndexRequest) SetCategoryId(v int32) {
	o.CategoryId = v
}

// GetShouldUpdate returns the ShouldUpdate field value if set, zero value otherwise.
func (o *CategoryUserIndexRequest) GetShouldUpdate() bool {
	if o == nil || o.ShouldUpdate == nil {
		var ret bool
		return ret
	}
	return *o.ShouldUpdate
}

// GetShouldUpdateOk returns a tuple with the ShouldUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryUserIndexRequest) GetShouldUpdateOk() (*bool, bool) {
	if o == nil || o.ShouldUpdate == nil {
		return nil, false
	}
	return o.ShouldUpdate, true
}

// HasShouldUpdate returns a boolean if a field has been set.
func (o *CategoryUserIndexRequest) HasShouldUpdate() bool {
	if o != nil && o.ShouldUpdate != nil {
		return true
	}

	return false
}

// SetShouldUpdate gets a reference to the given bool and assigns it to the ShouldUpdate field.
func (o *CategoryUserIndexRequest) SetShouldUpdate(v bool) {
	o.ShouldUpdate = &v
}

// GetUserId returns the UserId field value
func (o *CategoryUserIndexRequest) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *CategoryUserIndexRequest) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *CategoryUserIndexRequest) SetUserId(v string) {
	o.UserId = v
}

func (o CategoryUserIndexRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["categoryId"] = o.CategoryId
	}
	if o.ShouldUpdate != nil {
		toSerialize["shouldUpdate"] = o.ShouldUpdate
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryUserIndexRequest struct {
	value *CategoryUserIndexRequest
	isSet bool
}

func (v NullableCategoryUserIndexRequest) Get() *CategoryUserIndexRequest {
	return v.value
}

func (v *NullableCategoryUserIndexRequest) Set(val *CategoryUserIndexRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryUserIndexRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryUserIndexRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryUserIndexRequest(val *CategoryUserIndexRequest) *NullableCategoryUserIndexRequest {
	return &NullableCategoryUserIndexRequest{value: val, isSet: true}
}

func (v NullableCategoryUserIndexRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryUserIndexRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

