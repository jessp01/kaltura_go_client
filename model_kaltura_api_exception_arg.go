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

// KalturaApiExceptionArg struct for KalturaApiExceptionArg
type KalturaApiExceptionArg struct {
	Name *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewKalturaApiExceptionArg instantiates a new KalturaApiExceptionArg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaApiExceptionArg() *KalturaApiExceptionArg {
	this := KalturaApiExceptionArg{}
	return &this
}

// NewKalturaApiExceptionArgWithDefaults instantiates a new KalturaApiExceptionArg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaApiExceptionArgWithDefaults() *KalturaApiExceptionArg {
	this := KalturaApiExceptionArg{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaApiExceptionArg) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaApiExceptionArg) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaApiExceptionArg) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaApiExceptionArg) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KalturaApiExceptionArg) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaApiExceptionArg) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KalturaApiExceptionArg) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *KalturaApiExceptionArg) SetValue(v string) {
	o.Value = &v
}

func (o KalturaApiExceptionArg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaApiExceptionArg struct {
	value *KalturaApiExceptionArg
	isSet bool
}

func (v NullableKalturaApiExceptionArg) Get() *KalturaApiExceptionArg {
	return v.value
}

func (v *NullableKalturaApiExceptionArg) Set(val *KalturaApiExceptionArg) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaApiExceptionArg) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaApiExceptionArg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaApiExceptionArg(val *KalturaApiExceptionArg) *NullableKalturaApiExceptionArg {
	return &NullableKalturaApiExceptionArg{value: val, isSet: true}
}

func (v NullableKalturaApiExceptionArg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaApiExceptionArg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

