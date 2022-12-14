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

// DataServeRequest struct for DataServeRequest
type DataServeRequest struct {
	EntryId string `json:"entryId"`
	ForceProxy *bool `json:"forceProxy,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewDataServeRequest instantiates a new DataServeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataServeRequest(entryId string) *DataServeRequest {
	this := DataServeRequest{}
	this.EntryId = entryId
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// NewDataServeRequestWithDefaults instantiates a new DataServeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataServeRequestWithDefaults() *DataServeRequest {
	this := DataServeRequest{}
	var forceProxy bool = false
	this.ForceProxy = &forceProxy
	return &this
}

// GetEntryId returns the EntryId field value
func (o *DataServeRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *DataServeRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *DataServeRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetForceProxy returns the ForceProxy field value if set, zero value otherwise.
func (o *DataServeRequest) GetForceProxy() bool {
	if o == nil || o.ForceProxy == nil {
		var ret bool
		return ret
	}
	return *o.ForceProxy
}

// GetForceProxyOk returns a tuple with the ForceProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataServeRequest) GetForceProxyOk() (*bool, bool) {
	if o == nil || o.ForceProxy == nil {
		return nil, false
	}
	return o.ForceProxy, true
}

// HasForceProxy returns a boolean if a field has been set.
func (o *DataServeRequest) HasForceProxy() bool {
	if o != nil && o.ForceProxy != nil {
		return true
	}

	return false
}

// SetForceProxy gets a reference to the given bool and assigns it to the ForceProxy field.
func (o *DataServeRequest) SetForceProxy(v bool) {
	o.ForceProxy = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DataServeRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataServeRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DataServeRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *DataServeRequest) SetVersion(v int32) {
	o.Version = &v
}

func (o DataServeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if o.ForceProxy != nil {
		toSerialize["forceProxy"] = o.ForceProxy
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDataServeRequest struct {
	value *DataServeRequest
	isSet bool
}

func (v NullableDataServeRequest) Get() *DataServeRequest {
	return v.value
}

func (v *NullableDataServeRequest) Set(val *DataServeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDataServeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDataServeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataServeRequest(val *DataServeRequest) *NullableDataServeRequest {
	return &NullableDataServeRequest{value: val, isSet: true}
}

func (v NullableDataServeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataServeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


