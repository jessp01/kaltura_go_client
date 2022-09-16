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

// KalturaESearchOrderByItem `abstract`
type KalturaESearchOrderByItem struct {
	ObjectType *string `json:"objectType,omitempty"`
	// Enum Type: `KalturaESearchSortOrder`
	SortOrder *string `json:"sortOrder,omitempty"`
}

// NewKalturaESearchOrderByItem instantiates a new KalturaESearchOrderByItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchOrderByItem() *KalturaESearchOrderByItem {
	this := KalturaESearchOrderByItem{}
	return &this
}

// NewKalturaESearchOrderByItemWithDefaults instantiates a new KalturaESearchOrderByItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchOrderByItemWithDefaults() *KalturaESearchOrderByItem {
	this := KalturaESearchOrderByItem{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaESearchOrderByItem) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchOrderByItem) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaESearchOrderByItem) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaESearchOrderByItem) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *KalturaESearchOrderByItem) GetSortOrder() string {
	if o == nil || o.SortOrder == nil {
		var ret string
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchOrderByItem) GetSortOrderOk() (*string, bool) {
	if o == nil || o.SortOrder == nil {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *KalturaESearchOrderByItem) HasSortOrder() bool {
	if o != nil && o.SortOrder != nil {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given string and assigns it to the SortOrder field.
func (o *KalturaESearchOrderByItem) SetSortOrder(v string) {
	o.SortOrder = &v
}

func (o KalturaESearchOrderByItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.SortOrder != nil {
		toSerialize["sortOrder"] = o.SortOrder
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchOrderByItem struct {
	value *KalturaESearchOrderByItem
	isSet bool
}

func (v NullableKalturaESearchOrderByItem) Get() *KalturaESearchOrderByItem {
	return v.value
}

func (v *NullableKalturaESearchOrderByItem) Set(val *KalturaESearchOrderByItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchOrderByItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchOrderByItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchOrderByItem(val *KalturaESearchOrderByItem) *NullableKalturaESearchOrderByItem {
	return &NullableKalturaESearchOrderByItem{value: val, isSet: true}
}

func (v NullableKalturaESearchOrderByItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchOrderByItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

