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

// KalturaESearchAggregationBucket struct for KalturaESearchAggregationBucket
type KalturaESearchAggregationBucket struct {
	Count *int32 `json:"count,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewKalturaESearchAggregationBucket instantiates a new KalturaESearchAggregationBucket object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaESearchAggregationBucket() *KalturaESearchAggregationBucket {
	this := KalturaESearchAggregationBucket{}
	return &this
}

// NewKalturaESearchAggregationBucketWithDefaults instantiates a new KalturaESearchAggregationBucket object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaESearchAggregationBucketWithDefaults() *KalturaESearchAggregationBucket {
	this := KalturaESearchAggregationBucket{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *KalturaESearchAggregationBucket) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchAggregationBucket) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *KalturaESearchAggregationBucket) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *KalturaESearchAggregationBucket) SetCount(v int32) {
	o.Count = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *KalturaESearchAggregationBucket) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaESearchAggregationBucket) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *KalturaESearchAggregationBucket) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *KalturaESearchAggregationBucket) SetValue(v string) {
	o.Value = &v
}

func (o KalturaESearchAggregationBucket) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaESearchAggregationBucket struct {
	value *KalturaESearchAggregationBucket
	isSet bool
}

func (v NullableKalturaESearchAggregationBucket) Get() *KalturaESearchAggregationBucket {
	return v.value
}

func (v *NullableKalturaESearchAggregationBucket) Set(val *KalturaESearchAggregationBucket) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaESearchAggregationBucket) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaESearchAggregationBucket) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaESearchAggregationBucket(val *KalturaESearchAggregationBucket) *NullableKalturaESearchAggregationBucket {
	return &NullableKalturaESearchAggregationBucket{value: val, isSet: true}
}

func (v NullableKalturaESearchAggregationBucket) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaESearchAggregationBucket) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

