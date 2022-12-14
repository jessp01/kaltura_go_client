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

// KalturaValue `abstract`  A representation to return an array of values
type KalturaValue struct {
	Description *string `json:"description,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaValue instantiates a new KalturaValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaValue() *KalturaValue {
	this := KalturaValue{}
	return &this
}

// NewKalturaValueWithDefaults instantiates a new KalturaValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaValueWithDefaults() *KalturaValue {
	this := KalturaValue{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaValue) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaValue) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaValue) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaValue) SetDescription(v string) {
	o.Description = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaValue) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaValue) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaValue) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaValue) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaValue struct {
	value *KalturaValue
	isSet bool
}

func (v NullableKalturaValue) Get() *KalturaValue {
	return v.value
}

func (v *NullableKalturaValue) Set(val *KalturaValue) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaValue) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaValue(val *KalturaValue) *NullableKalturaValue {
	return &NullableKalturaValue{value: val, isSet: true}
}

func (v NullableKalturaValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


