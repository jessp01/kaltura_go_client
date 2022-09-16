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

// InteractivityGetRequest struct for InteractivityGetRequest
type InteractivityGetRequest struct {
	DataFilter *KalturaInteractivityDataFilter `json:"dataFilter,omitempty"`
	EntryId string `json:"entryId"`
}

// NewInteractivityGetRequest instantiates a new InteractivityGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInteractivityGetRequest(entryId string) *InteractivityGetRequest {
	this := InteractivityGetRequest{}
	this.EntryId = entryId
	return &this
}

// NewInteractivityGetRequestWithDefaults instantiates a new InteractivityGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInteractivityGetRequestWithDefaults() *InteractivityGetRequest {
	this := InteractivityGetRequest{}
	return &this
}

// GetDataFilter returns the DataFilter field value if set, zero value otherwise.
func (o *InteractivityGetRequest) GetDataFilter() KalturaInteractivityDataFilter {
	if o == nil || o.DataFilter == nil {
		var ret KalturaInteractivityDataFilter
		return ret
	}
	return *o.DataFilter
}

// GetDataFilterOk returns a tuple with the DataFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InteractivityGetRequest) GetDataFilterOk() (*KalturaInteractivityDataFilter, bool) {
	if o == nil || o.DataFilter == nil {
		return nil, false
	}
	return o.DataFilter, true
}

// HasDataFilter returns a boolean if a field has been set.
func (o *InteractivityGetRequest) HasDataFilter() bool {
	if o != nil && o.DataFilter != nil {
		return true
	}

	return false
}

// SetDataFilter gets a reference to the given KalturaInteractivityDataFilter and assigns it to the DataFilter field.
func (o *InteractivityGetRequest) SetDataFilter(v KalturaInteractivityDataFilter) {
	o.DataFilter = &v
}

// GetEntryId returns the EntryId field value
func (o *InteractivityGetRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *InteractivityGetRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *InteractivityGetRequest) SetEntryId(v string) {
	o.EntryId = v
}

func (o InteractivityGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataFilter != nil {
		toSerialize["dataFilter"] = o.DataFilter
	}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	return json.Marshal(toSerialize)
}

type NullableInteractivityGetRequest struct {
	value *InteractivityGetRequest
	isSet bool
}

func (v NullableInteractivityGetRequest) Get() *InteractivityGetRequest {
	return v.value
}

func (v *NullableInteractivityGetRequest) Set(val *InteractivityGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInteractivityGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInteractivityGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInteractivityGetRequest(val *InteractivityGetRequest) *NullableInteractivityGetRequest {
	return &NullableInteractivityGetRequest{value: val, isSet: true}
}

func (v NullableInteractivityGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInteractivityGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

