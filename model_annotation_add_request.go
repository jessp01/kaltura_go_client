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

// AnnotationAddRequest struct for AnnotationAddRequest
type AnnotationAddRequest struct {
	Annotation KalturaCuePoint `json:"annotation"`
}

// NewAnnotationAddRequest instantiates a new AnnotationAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnotationAddRequest(annotation KalturaCuePoint) *AnnotationAddRequest {
	this := AnnotationAddRequest{}
	this.Annotation = annotation
	return &this
}

// NewAnnotationAddRequestWithDefaults instantiates a new AnnotationAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnotationAddRequestWithDefaults() *AnnotationAddRequest {
	this := AnnotationAddRequest{}
	return &this
}

// GetAnnotation returns the Annotation field value
func (o *AnnotationAddRequest) GetAnnotation() KalturaCuePoint {
	if o == nil {
		var ret KalturaCuePoint
		return ret
	}

	return o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value
// and a boolean to check if the value has been set.
func (o *AnnotationAddRequest) GetAnnotationOk() (*KalturaCuePoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotation, true
}

// SetAnnotation sets field value
func (o *AnnotationAddRequest) SetAnnotation(v KalturaCuePoint) {
	o.Annotation = v
}

func (o AnnotationAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["annotation"] = o.Annotation
	}
	return json.Marshal(toSerialize)
}

type NullableAnnotationAddRequest struct {
	value *AnnotationAddRequest
	isSet bool
}

func (v NullableAnnotationAddRequest) Get() *AnnotationAddRequest {
	return v.value
}

func (v *NullableAnnotationAddRequest) Set(val *AnnotationAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnotationAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnotationAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnotationAddRequest(val *AnnotationAddRequest) *NullableAnnotationAddRequest {
	return &NullableAnnotationAddRequest{value: val, isSet: true}
}

func (v NullableAnnotationAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnotationAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


