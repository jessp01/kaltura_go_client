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

// KalturaESearchAggregationItem `abstract`
type KalturaESearchAggregationItem struct {
	ObjectType *string `json:"objectType,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewKalturaESearchAggregationItem instantiates a new KalturaESearchAggregationItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchAggregationItem() *KalturaESearchAggregationItem {
	this := KalturaESearchAggregationItem{}
	return &this
}

// NewKalturaESearchAggregationItemWithDefaults instantiates a new KalturaESearchAggregationItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchAggregationItemWithDefaults() *KalturaESearchAggregationItem {
	this := KalturaESearchAggregationItem{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaESearchAggregationItem) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchAggregationItem) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaESearchAggregationItem) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaESearchAggregationItem) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *KalturaESearchAggregationItem) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchAggregationItem) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *KalturaESearchAggregationItem) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *KalturaESearchAggregationItem) SetSize(v int32) {
	o.Size = &v
}

func (o KalturaESearchAggregationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchAggregationItem struct {
	value *KalturaESearchAggregationItem
	isSet bool
}

func (v NullableKalturaESearchAggregationItem) Get() *KalturaESearchAggregationItem {
	return v.value
}

func (v *NullableKalturaESearchAggregationItem) Set(val *KalturaESearchAggregationItem) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchAggregationItem) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchAggregationItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchAggregationItem(val *KalturaESearchAggregationItem) *NullableKalturaESearchAggregationItem {
	return &NullableKalturaESearchAggregationItem{value: val, isSet: true}
}

func (v NullableKalturaESearchAggregationItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchAggregationItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

