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

// CategoryUserAddRequest struct for CategoryUserAddRequest
type CategoryUserAddRequest struct {
	CategoryUser KalturaCategoryUser `json:"categoryUser"`
}

// NewCategoryUserAddRequest instantiates a new CategoryUserAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryUserAddRequest(categoryUser KalturaCategoryUser) *CategoryUserAddRequest {
	this := CategoryUserAddRequest{}
	this.CategoryUser = categoryUser
	return &this
}

// NewCategoryUserAddRequestWithDefaults instantiates a new CategoryUserAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryUserAddRequestWithDefaults() *CategoryUserAddRequest {
	this := CategoryUserAddRequest{}
	return &this
}

// GetCategoryUser returns the CategoryUser field value
func (o *CategoryUserAddRequest) GetCategoryUser() KalturaCategoryUser {
	if o == nil {
		var ret KalturaCategoryUser
		return ret
	}

	return o.CategoryUser
}

// GetCategoryUserOk returns a tuple with the CategoryUser field value
// and a boolean to check if the value has been set.
func (o *CategoryUserAddRequest) GetCategoryUserOk() (*KalturaCategoryUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryUser, true
}

// SetCategoryUser sets field value
func (o *CategoryUserAddRequest) SetCategoryUser(v KalturaCategoryUser) {
	o.CategoryUser = v
}

func (o CategoryUserAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["categoryUser"] = o.CategoryUser
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryUserAddRequest struct {
	value *CategoryUserAddRequest
	isSet bool
}

func (v NullableCategoryUserAddRequest) Get() *CategoryUserAddRequest {
	return v.value
}

func (v *NullableCategoryUserAddRequest) Set(val *CategoryUserAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryUserAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryUserAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryUserAddRequest(val *CategoryUserAddRequest) *NullableCategoryUserAddRequest {
	return &NullableCategoryUserAddRequest{value: val, isSet: true}
}

func (v NullableCategoryUserAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryUserAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


