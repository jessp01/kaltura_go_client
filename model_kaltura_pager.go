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

// KalturaPager The KalturaPager object enables paging management to be applied upon service list/search actions.
type KalturaPager struct {
	ObjectType *string `json:"objectType,omitempty"`
	// The page number for which {pageSize} of objects should be retrieved (Default is 1).
	PageIndex *int32 `json:"pageIndex,omitempty"`
	// The number of objects to retrieve. (Default is 30, maximum page size is 500).
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewKalturaPager instantiates a new KalturaPager object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaPager() *KalturaPager {
	this := KalturaPager{}
	return &this
}

// NewKalturaPagerWithDefaults instantiates a new KalturaPager object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaPagerWithDefaults() *KalturaPager {
	this := KalturaPager{}
	return &this
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaPager) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPager) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaPager) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaPager) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPageIndex returns the PageIndex field value if set, zero value otherwise.
func (o *KalturaPager) GetPageIndex() int32 {
	if o == nil || o.PageIndex == nil {
		var ret int32
		return ret
	}
	return *o.PageIndex
}

// GetPageIndexOk returns a tuple with the PageIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPager) GetPageIndexOk() (*int32, bool) {
	if o == nil || o.PageIndex == nil {
		return nil, false
	}
	return o.PageIndex, true
}

// HasPageIndex returns a boolean if a field has been set.
func (o *KalturaPager) HasPageIndex() bool {
	if o != nil && o.PageIndex != nil {
		return true
	}

	return false
}

// SetPageIndex gets a reference to the given int32 and assigns it to the PageIndex field.
func (o *KalturaPager) SetPageIndex(v int32) {
	o.PageIndex = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *KalturaPager) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaPager) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *KalturaPager) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *KalturaPager) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o KalturaPager) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PageIndex != nil {
		toSerialize["pageIndex"] = o.PageIndex
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaPager struct {
	value *KalturaPager
	isSet bool
}

func (v NullableKalturaPager) Get() *KalturaPager {
	return v.value
}

func (v *NullableKalturaPager) Set(val *KalturaPager) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaPager) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaPager) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaPager(val *KalturaPager) *NullableKalturaPager {
	return &NullableKalturaPager{value: val, isSet: true}
}

func (v NullableKalturaPager) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaPager) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


