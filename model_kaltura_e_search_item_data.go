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

// KalturaESearchItemData `abstract`
type KalturaESearchItemData struct {
	Highlight []KalturaESearchHighlight `json:"highlight,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
}

// NewKalturaESearchItemData instantiates a new KalturaESearchItemData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchItemData() *KalturaESearchItemData {
	this := KalturaESearchItemData{}
	return &this
}

// NewKalturaESearchItemDataWithDefaults instantiates a new KalturaESearchItemData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchItemDataWithDefaults() *KalturaESearchItemData {
	this := KalturaESearchItemData{}
	return &this
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *KalturaESearchItemData) GetHighlight() []KalturaESearchHighlight {
	if o == nil || o.Highlight == nil {
		var ret []KalturaESearchHighlight
		return ret
	}
	return o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchItemData) GetHighlightOk() ([]KalturaESearchHighlight, bool) {
	if o == nil || o.Highlight == nil {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *KalturaESearchItemData) HasHighlight() bool {
	if o != nil && o.Highlight != nil {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given []KalturaESearchHighlight and assigns it to the Highlight field.
func (o *KalturaESearchItemData) SetHighlight(v []KalturaESearchHighlight) {
	o.Highlight = v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaESearchItemData) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchItemData) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaESearchItemData) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaESearchItemData) SetObjectType(v string) {
	o.ObjectType = &v
}

func (o KalturaESearchItemData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Highlight != nil {
		toSerialize["highlight"] = o.Highlight
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchItemData struct {
	value *KalturaESearchItemData
	isSet bool
}

func (v NullableKalturaESearchItemData) Get() *KalturaESearchItemData {
	return v.value
}

func (v *NullableKalturaESearchItemData) Set(val *KalturaESearchItemData) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchItemData) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchItemData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchItemData(val *KalturaESearchItemData) *NullableKalturaESearchItemData {
	return &NullableKalturaESearchItemData{value: val, isSet: true}
}

func (v NullableKalturaESearchItemData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchItemData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

