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

// DataAddRequest struct for DataAddRequest
type DataAddRequest struct {
	DataEntry KalturaDataEntry `json:"dataEntry"`
}

// NewDataAddRequest instantiates a new DataAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAddRequest(dataEntry KalturaDataEntry) *DataAddRequest {
	this := DataAddRequest{}
	this.DataEntry = dataEntry
	return &this
}

// NewDataAddRequestWithDefaults instantiates a new DataAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAddRequestWithDefaults() *DataAddRequest {
	this := DataAddRequest{}
	return &this
}

// GetDataEntry returns the DataEntry field value
func (o *DataAddRequest) GetDataEntry() KalturaDataEntry {
	if o == nil {
		var ret KalturaDataEntry
		return ret
	}

	return o.DataEntry
}

// GetDataEntryOk returns a tuple with the DataEntry field value
// and a boolean to check if the value has been set.
func (o *DataAddRequest) GetDataEntryOk() (*KalturaDataEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataEntry, true
}

// SetDataEntry sets field value
func (o *DataAddRequest) SetDataEntry(v KalturaDataEntry) {
	o.DataEntry = v
}

func (o DataAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dataEntry"] = o.DataEntry
	}
	return json.Marshal(toSerialize)
}

type NullableDataAddRequest struct {
	value *DataAddRequest
	isSet bool
}

func (v NullableDataAddRequest) Get() *DataAddRequest {
	return v.value
}

func (v *NullableDataAddRequest) Set(val *DataAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAddRequest(val *DataAddRequest) *NullableDataAddRequest {
	return &NullableDataAddRequest{value: val, isSet: true}
}

func (v NullableDataAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


