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

// GenericDistributionProviderActionAddResultsTransformRequest struct for GenericDistributionProviderActionAddResultsTransformRequest
type GenericDistributionProviderActionAddResultsTransformRequest struct {
	Id int32 `json:"id"`
	TransformData string `json:"transformData"`
}

// NewGenericDistributionProviderActionAddResultsTransformRequest instantiates a new GenericDistributionProviderActionAddResultsTransformRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericDistributionProviderActionAddResultsTransformRequest(id int32, transformData string) *GenericDistributionProviderActionAddResultsTransformRequest {
	this := GenericDistributionProviderActionAddResultsTransformRequest{}
	this.Id = id
	this.TransformData = transformData
	return &this
}

// NewGenericDistributionProviderActionAddResultsTransformRequestWithDefaults instantiates a new GenericDistributionProviderActionAddResultsTransformRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericDistributionProviderActionAddResultsTransformRequestWithDefaults() *GenericDistributionProviderActionAddResultsTransformRequest {
	this := GenericDistributionProviderActionAddResultsTransformRequest{}
	return &this
}

// GetId returns the Id field value
func (o *GenericDistributionProviderActionAddResultsTransformRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GenericDistributionProviderActionAddResultsTransformRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GenericDistributionProviderActionAddResultsTransformRequest) SetId(v int32) {
	o.Id = v
}

// GetTransformData returns the TransformData field value
func (o *GenericDistributionProviderActionAddResultsTransformRequest) GetTransformData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransformData
}

// GetTransformDataOk returns a tuple with the TransformData field value
// and a boolean to check if the value has been set.
func (o *GenericDistributionProviderActionAddResultsTransformRequest) GetTransformDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransformData, true
}

// SetTransformData sets field value
func (o *GenericDistributionProviderActionAddResultsTransformRequest) SetTransformData(v string) {
	o.TransformData = v
}

func (o GenericDistributionProviderActionAddResultsTransformRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["transformData"] = o.TransformData
	}
	return json.Marshal(toSerialize)
}

type NullableGenericDistributionProviderActionAddResultsTransformRequest struct {
	value *GenericDistributionProviderActionAddResultsTransformRequest
	isSet bool
}

func (v NullableGenericDistributionProviderActionAddResultsTransformRequest) Get() *GenericDistributionProviderActionAddResultsTransformRequest {
	return v.value
}

func (v *NullableGenericDistributionProviderActionAddResultsTransformRequest) Set(val *GenericDistributionProviderActionAddResultsTransformRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericDistributionProviderActionAddResultsTransformRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericDistributionProviderActionAddResultsTransformRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericDistributionProviderActionAddResultsTransformRequest(val *GenericDistributionProviderActionAddResultsTransformRequest) *NullableGenericDistributionProviderActionAddResultsTransformRequest {
	return &NullableGenericDistributionProviderActionAddResultsTransformRequest{value: val, isSet: true}
}

func (v NullableGenericDistributionProviderActionAddResultsTransformRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericDistributionProviderActionAddResultsTransformRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


