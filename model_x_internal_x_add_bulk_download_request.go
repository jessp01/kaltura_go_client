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

// XInternalXAddBulkDownloadRequest struct for XInternalXAddBulkDownloadRequest
type XInternalXAddBulkDownloadRequest struct {
	EntryIds string `json:"entryIds"`
	FlavorParamsId *string `json:"flavorParamsId,omitempty"`
}

// NewXInternalXAddBulkDownloadRequest instantiates a new XInternalXAddBulkDownloadRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXInternalXAddBulkDownloadRequest(entryIds string) *XInternalXAddBulkDownloadRequest {
	this := XInternalXAddBulkDownloadRequest{}
	this.EntryIds = entryIds
	return &this
}

// NewXInternalXAddBulkDownloadRequestWithDefaults instantiates a new XInternalXAddBulkDownloadRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXInternalXAddBulkDownloadRequestWithDefaults() *XInternalXAddBulkDownloadRequest {
	this := XInternalXAddBulkDownloadRequest{}
	return &this
}

// GetEntryIds returns the EntryIds field value
func (o *XInternalXAddBulkDownloadRequest) GetEntryIds() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryIds
}

// GetEntryIdsOk returns a tuple with the EntryIds field value
// and a boolean to check if the value has been set.
func (o *XInternalXAddBulkDownloadRequest) GetEntryIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryIds, true
}

// SetEntryIds sets field value
func (o *XInternalXAddBulkDownloadRequest) SetEntryIds(v string) {
	o.EntryIds = v
}

// GetFlavorParamsId returns the FlavorParamsId field value if set, zero value otherwise.
func (o *XInternalXAddBulkDownloadRequest) GetFlavorParamsId() string {
	if o == nil || o.FlavorParamsId == nil {
		var ret string
		return ret
	}
	return *o.FlavorParamsId
}

// GetFlavorParamsIdOk returns a tuple with the FlavorParamsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XInternalXAddBulkDownloadRequest) GetFlavorParamsIdOk() (*string, bool) {
	if o == nil || o.FlavorParamsId == nil {
		return nil, false
	}
	return o.FlavorParamsId, true
}

// HasFlavorParamsId returns a boolean if a field has been set.
func (o *XInternalXAddBulkDownloadRequest) HasFlavorParamsId() bool {
	if o != nil && o.FlavorParamsId != nil {
		return true
	}

	return false
}

// SetFlavorParamsId gets a reference to the given string and assigns it to the FlavorParamsId field.
func (o *XInternalXAddBulkDownloadRequest) SetFlavorParamsId(v string) {
	o.FlavorParamsId = &v
}

func (o XInternalXAddBulkDownloadRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryIds"] = o.EntryIds
	}
	if o.FlavorParamsId != nil {
		toSerialize["flavorParamsId"] = o.FlavorParamsId
	}
	return json.Marshal(toSerialize)
}

type NullableXInternalXAddBulkDownloadRequest struct {
	value *XInternalXAddBulkDownloadRequest
	isSet bool
}

func (v NullableXInternalXAddBulkDownloadRequest) Get() *XInternalXAddBulkDownloadRequest {
	return v.value
}

func (v *NullableXInternalXAddBulkDownloadRequest) Set(val *XInternalXAddBulkDownloadRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableXInternalXAddBulkDownloadRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableXInternalXAddBulkDownloadRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXInternalXAddBulkDownloadRequest(val *XInternalXAddBulkDownloadRequest) *NullableXInternalXAddBulkDownloadRequest {
	return &NullableXInternalXAddBulkDownloadRequest{value: val, isSet: true}
}

func (v NullableXInternalXAddBulkDownloadRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXInternalXAddBulkDownloadRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


