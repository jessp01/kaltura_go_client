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

// KalturaESearchBaseFilter `abstract`
type KalturaESearchBaseFilter struct {
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaESearchBaseFilter instantiates a new KalturaESearchBaseFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchBaseFilter() *KalturaESearchBaseFilter {
	this := KalturaESearchBaseFilter{}
	return &this
}

// NewKalturaESearchBaseFilterWithDefaults instantiates a new KalturaESearchBaseFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchBaseFilterWithDefaults() *KalturaESearchBaseFilter {
	this := KalturaESearchBaseFilter{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaESearchBaseFilter) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchBaseFilter) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaESearchBaseFilter) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaESearchBaseFilter) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaESearchBaseFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchBaseFilter struct {
	value *KalturaESearchBaseFilter
	isSet bool
}

func (v NullableKalturaESearchBaseFilter) Get() *KalturaESearchBaseFilter {
	return v.value
}

func (v *NullableKalturaESearchBaseFilter) Set(val *KalturaESearchBaseFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchBaseFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchBaseFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchBaseFilter(val *KalturaESearchBaseFilter) *NullableKalturaESearchBaseFilter {
	return &NullableKalturaESearchBaseFilter{value: val, isSet: true}
}

func (v NullableKalturaESearchBaseFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchBaseFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


