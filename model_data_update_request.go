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

// DataUpdateRequest struct for DataUpdateRequest
type DataUpdateRequest struct {
	DocumentEntry KalturaDataEntry `json:"documentEntry"`
	EntryId string `json:"entryId"`
}

// NewDataUpdateRequest instantiates a new DataUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataUpdateRequest(documentEntry KalturaDataEntry, entryId string) *DataUpdateRequest {
	this := DataUpdateRequest{}
	this.DocumentEntry = documentEntry
	this.EntryId = entryId
	return &this
}

// NewDataUpdateRequestWithDefaults instantiates a new DataUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataUpdateRequestWithDefaults() *DataUpdateRequest {
	this := DataUpdateRequest{}
	return &this
}

// GetDocumentEntry returns the DocumentEntry field value
func (o *DataUpdateRequest) GetDocumentEntry() KalturaDataEntry {
	if o == nil {
		var ret KalturaDataEntry
		return ret
	}

	return o.DocumentEntry
}

// GetDocumentEntryOk returns a tuple with the DocumentEntry field value
// and a boolean to check if the value has been set.
func (o *DataUpdateRequest) GetDocumentEntryOk() (*KalturaDataEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentEntry, true
}

// SetDocumentEntry sets field value
func (o *DataUpdateRequest) SetDocumentEntry(v KalturaDataEntry) {
	o.DocumentEntry = v
}

// GetEntryId returns the EntryId field value
func (o *DataUpdateRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *DataUpdateRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *DataUpdateRequest) SetEntryId(v string) {
	o.EntryId = v
}

func (o DataUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["documentEntry"] = o.DocumentEntry
	}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	return json.Marshal(toSerialize)
}

type NullableDataUpdateRequest struct {
	value *DataUpdateRequest
	isSet bool
}

func (v NullableDataUpdateRequest) Get() *DataUpdateRequest {
	return v.value
}

func (v *NullableDataUpdateRequest) Set(val *DataUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataUpdateRequest(val *DataUpdateRequest) *NullableDataUpdateRequest {
	return &NullableDataUpdateRequest{value: val, isSet: true}
}

func (v NullableDataUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


