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

// AnnotationCountRequest struct for AnnotationCountRequest
type AnnotationCountRequest struct {
	Filter *KalturaCuePointFilter `json:"filter,omitempty"`
}

// NewAnnotationCountRequest instantiates a new AnnotationCountRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationCountRequest() *AnnotationCountRequest {
	this := AnnotationCountRequest{}
	return &this
}

// NewAnnotationCountRequestWithDefaults instantiates a new AnnotationCountRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationCountRequestWithDefaults() *AnnotationCountRequest {
	this := AnnotationCountRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *AnnotationCountRequest) GetFilter() KalturaCuePointFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaCuePointFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationCountRequest) GetFilterOk() (*KalturaCuePointFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *AnnotationCountRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaCuePointFilter and assigns it to the Filter field.
func (o *AnnotationCountRequest) SetFilter(v KalturaCuePointFilter) {
	o.Filter = &v
}

func (o AnnotationCountRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableAnnotationCountRequest struct {
	value *AnnotationCountRequest
	isSet bool
}

func (v NullableAnnotationCountRequest) Get() *AnnotationCountRequest {
	return v.value
}

func (v *NullableAnnotationCountRequest) Set(val *AnnotationCountRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationCountRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationCountRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationCountRequest(val *AnnotationCountRequest) *NullableAnnotationCountRequest {
	return &NullableAnnotationCountRequest{value: val, isSet: true}
}

func (v NullableAnnotationCountRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationCountRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


