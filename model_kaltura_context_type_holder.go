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

// KalturaContextTypeHolder struct for KalturaContextTypeHolder
type KalturaContextTypeHolder struct {
	ObjectType *string `json:"objectType,omitempty"`
	// Enum Type: `KalturaContextType`  The type of the condition context
	Type *string `json:"type,omitempty"`
}

// NewKalturaContextTypeHolder instantiates a new KalturaContextTypeHolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaContextTypeHolder() *KalturaContextTypeHolder {
	this := KalturaContextTypeHolder{}
	return &this
}

// NewKalturaContextTypeHolderWithDefaults instantiates a new KalturaContextTypeHolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaContextTypeHolderWithDefaults() *KalturaContextTypeHolder {
	this := KalturaContextTypeHolder{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaContextTypeHolder) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaContextTypeHolder) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaContextTypeHolder) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaContextTypeHolder) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaContextTypeHolder) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaContextTypeHolder) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaContextTypeHolder) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KalturaContextTypeHolder) SetType(v string) {
	o.Type = &v
}

func (o KalturaContextTypeHolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaContextTypeHolder struct {
	value *KalturaContextTypeHolder
	isSet bool
}

func (v NullableKalturaContextTypeHolder) Get() *KalturaContextTypeHolder {
	return v.value
}

func (v *NullableKalturaContextTypeHolder) Set(val *KalturaContextTypeHolder) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaContextTypeHolder) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaContextTypeHolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaContextTypeHolder(val *KalturaContextTypeHolder) *NullableKalturaContextTypeHolder {
	return &NullableKalturaContextTypeHolder{value: val, isSet: true}
}

func (v NullableKalturaContextTypeHolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaContextTypeHolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


