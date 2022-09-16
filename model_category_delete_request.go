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

// CategoryDeleteRequest struct for CategoryDeleteRequest
type CategoryDeleteRequest struct {
	Id int32 `json:"id"`
	MoveEntriesToParentCategory *int32 `json:"moveEntriesToParentCategory,omitempty"`
}

// NewCategoryDeleteRequest instantiates a new CategoryDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCategoryDeleteRequest(id int32) *CategoryDeleteRequest {
	this := CategoryDeleteRequest{}
	this.Id = id
	return &this
}

// NewCategoryDeleteRequestWithDefaults instantiates a new CategoryDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCategoryDeleteRequestWithDefaults() *CategoryDeleteRequest {
	this := CategoryDeleteRequest{}
	return &this
}

// GetId returns the Id field value
func (o *CategoryDeleteRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CategoryDeleteRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CategoryDeleteRequest) SetId(v int32) {
	o.Id = v
}

// GetMoveEntriesToParentCategory returns the MoveEntriesToParentCategory field value if set, zero value otherwise.
func (o *CategoryDeleteRequest) GetMoveEntriesToParentCategory() int32 {
	if o == nil || o.MoveEntriesToParentCategory == nil {
		var ret int32
		return ret
	}
	return *o.MoveEntriesToParentCategory
}

// GetMoveEntriesToParentCategoryOk returns a tuple with the MoveEntriesToParentCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CategoryDeleteRequest) GetMoveEntriesToParentCategoryOk() (*int32, bool) {
	if o == nil || o.MoveEntriesToParentCategory == nil {
		return nil, false
	}
	return o.MoveEntriesToParentCategory, true
}

// HasMoveEntriesToParentCategory returns a boolean if a field has been set.
func (o *CategoryDeleteRequest) HasMoveEntriesToParentCategory() bool {
	if o != nil && o.MoveEntriesToParentCategory != nil {
		return true
	}

	return false
}

// SetMoveEntriesToParentCategory gets a reference to the given int32 and assigns it to the MoveEntriesToParentCategory field.
func (o *CategoryDeleteRequest) SetMoveEntriesToParentCategory(v int32) {
	o.MoveEntriesToParentCategory = &v
}

func (o CategoryDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.MoveEntriesToParentCategory != nil {
		toSerialize["moveEntriesToParentCategory"] = o.MoveEntriesToParentCategory
	}
	return json.Marshal(toSerialize)
}

type NullableCategoryDeleteRequest struct {
	value *CategoryDeleteRequest
	isSet bool
}

func (v NullableCategoryDeleteRequest) Get() *CategoryDeleteRequest {
	return v.value
}

func (v *NullableCategoryDeleteRequest) Set(val *CategoryDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCategoryDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCategoryDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCategoryDeleteRequest(val *CategoryDeleteRequest) *NullableCategoryDeleteRequest {
	return &NullableCategoryDeleteRequest{value: val, isSet: true}
}

func (v NullableCategoryDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCategoryDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


