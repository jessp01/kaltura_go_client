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

// MetadataProfileRevertRequest struct for MetadataProfileRevertRequest
type MetadataProfileRevertRequest struct {
	Id int32 `json:"id"`
	ToVersion int32 `json:"toVersion"`
}

// NewMetadataProfileRevertRequest instantiates a new MetadataProfileRevertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataProfileRevertRequest(id int32, toVersion int32) *MetadataProfileRevertRequest {
	this := MetadataProfileRevertRequest{}
	this.Id = id
	this.ToVersion = toVersion
	return &this
}

// NewMetadataProfileRevertRequestWithDefaults instantiates a new MetadataProfileRevertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataProfileRevertRequestWithDefaults() *MetadataProfileRevertRequest {
	this := MetadataProfileRevertRequest{}
	return &this
}

// GetId returns the Id field value
func (o *MetadataProfileRevertRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetadataProfileRevertRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetadataProfileRevertRequest) SetId(v int32) {
	o.Id = v
}

// GetToVersion returns the ToVersion field value
func (o *MetadataProfileRevertRequest) GetToVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ToVersion
}

// GetToVersionOk returns a tuple with the ToVersion field value
// and a boolean to check if the value has been set.
func (o *MetadataProfileRevertRequest) GetToVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToVersion, true
}

// SetToVersion sets field value
func (o *MetadataProfileRevertRequest) SetToVersion(v int32) {
	o.ToVersion = v
}

func (o MetadataProfileRevertRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["toVersion"] = o.ToVersion
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataProfileRevertRequest struct {
	value *MetadataProfileRevertRequest
	isSet bool
}

func (v NullableMetadataProfileRevertRequest) Get() *MetadataProfileRevertRequest {
	return v.value
}

func (v *NullableMetadataProfileRevertRequest) Set(val *MetadataProfileRevertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataProfileRevertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataProfileRevertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataProfileRevertRequest(val *MetadataProfileRevertRequest) *NullableMetadataProfileRevertRequest {
	return &NullableMetadataProfileRevertRequest{value: val, isSet: true}
}

func (v NullableMetadataProfileRevertRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataProfileRevertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


