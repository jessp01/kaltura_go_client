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

// PexipListRoomsRequest struct for PexipListRoomsRequest
type PexipListRoomsRequest struct {
	ActiveOnly *bool `json:"activeOnly,omitempty"`
	Offset *int32 `json:"offset,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewPexipListRoomsRequest instantiates a new PexipListRoomsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPexipListRoomsRequest() *PexipListRoomsRequest {
	this := PexipListRoomsRequest{}
	var activeOnly bool = false
	this.ActiveOnly = &activeOnly
	return &this
}

// NewPexipListRoomsRequestWithDefaults instantiates a new PexipListRoomsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPexipListRoomsRequestWithDefaults() *PexipListRoomsRequest {
	this := PexipListRoomsRequest{}
	var activeOnly bool = false
	this.ActiveOnly = &activeOnly
	return &this
}

// GetActiveOnly returns the ActiveOnly field value if set, zero value otherwise.
func (o *PexipListRoomsRequest) GetActiveOnly() bool {
	if o == nil || o.ActiveOnly == nil {
		var ret bool
		return ret
	}
	return *o.ActiveOnly
}

// GetActiveOnlyOk returns a tuple with the ActiveOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PexipListRoomsRequest) GetActiveOnlyOk() (*bool, bool) {
	if o == nil || o.ActiveOnly == nil {
		return nil, false
	}
	return o.ActiveOnly, true
}

// HasActiveOnly returns a boolean if a field has been set.
func (o *PexipListRoomsRequest) HasActiveOnly() bool {
	if o != nil && o.ActiveOnly != nil {
		return true
	}

	return false
}

// SetActiveOnly gets a reference to the given bool and assigns it to the ActiveOnly field.
func (o *PexipListRoomsRequest) SetActiveOnly(v bool) {
	o.ActiveOnly = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *PexipListRoomsRequest) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PexipListRoomsRequest) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *PexipListRoomsRequest) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *PexipListRoomsRequest) SetOffset(v int32) {
	o.Offset = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *PexipListRoomsRequest) GetPageSize() int32 {
	if o == nil || o.PageSize == nil {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PexipListRoomsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *PexipListRoomsRequest) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *PexipListRoomsRequest) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o PexipListRoomsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveOnly != nil {
		toSerialize["activeOnly"] = o.ActiveOnly
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.PageSize != nil {
		toSerialize["pageSize"] = o.PageSize
	}
	return json.Marshal(toSerialize)
}

type NullablePexipListRoomsRequest struct {
	value *PexipListRoomsRequest
	isSet bool
}

func (v NullablePexipListRoomsRequest) Get() *PexipListRoomsRequest {
	return v.value
}

func (v *NullablePexipListRoomsRequest) Set(val *PexipListRoomsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePexipListRoomsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePexipListRoomsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePexipListRoomsRequest(val *PexipListRoomsRequest) *NullablePexipListRoomsRequest {
	return &NullablePexipListRoomsRequest{value: val, isSet: true}
}

func (v NullablePexipListRoomsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePexipListRoomsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


