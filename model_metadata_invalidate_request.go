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

// MetadataInvalidateRequest struct for MetadataInvalidateRequest
type MetadataInvalidateRequest struct {
	Id int32 `json:"id"`
	Version *int32 `json:"version,omitempty"`
}

// NewMetadataInvalidateRequest instantiates a new MetadataInvalidateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataInvalidateRequest(id int32) *MetadataInvalidateRequest {
	this := MetadataInvalidateRequest{}
	this.Id = id
	return &this
}

// NewMetadataInvalidateRequestWithDefaults instantiates a new MetadataInvalidateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataInvalidateRequestWithDefaults() *MetadataInvalidateRequest {
	this := MetadataInvalidateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *MetadataInvalidateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetadataInvalidateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetadataInvalidateRequest) SetId(v int32) {
	o.Id = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MetadataInvalidateRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataInvalidateRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MetadataInvalidateRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *MetadataInvalidateRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o MetadataInvalidateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataInvalidateRequest struct {
	value *MetadataInvalidateRequest
	isSet bool
}

func (v NullableMetadataInvalidateRequest) Get() *MetadataInvalidateRequest {
	return v.value
}

func (v *NullableMetadataInvalidateRequest) Set(val *MetadataInvalidateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataInvalidateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataInvalidateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataInvalidateRequest(val *MetadataInvalidateRequest) *NullableMetadataInvalidateRequest {
	return &NullableMetadataInvalidateRequest{value: val, isSet: true}
}

func (v NullableMetadataInvalidateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataInvalidateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


