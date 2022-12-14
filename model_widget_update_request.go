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

// WidgetUpdateRequest struct for WidgetUpdateRequest
type WidgetUpdateRequest struct {
	Id string `json:"id"`
	Widget KalturaWidget `json:"widget"`
}

// NewWidgetUpdateRequest instantiates a new WidgetUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWidgetUpdateRequest(id string, widget KalturaWidget) *WidgetUpdateRequest {
	this := WidgetUpdateRequest{}
	this.Id = id
	this.Widget = widget
	return &this
}

// NewWidgetUpdateRequestWithDefaults instantiates a new WidgetUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWidgetUpdateRequestWithDefaults() *WidgetUpdateRequest {
	this := WidgetUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *WidgetUpdateRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WidgetUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WidgetUpdateRequest) SetId(v string) {
	o.Id = v
}

// GetWidget returns the Widget field value
func (o *WidgetUpdateRequest) GetWidget() KalturaWidget {
	if o == nil {
		var ret KalturaWidget
		return ret
	}

	return o.Widget
}

// GetWidgetOk returns a tuple with the Widget field value
// and a boolean to check if the value has been set.
func (o *WidgetUpdateRequest) GetWidgetOk() (*KalturaWidget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Widget, true
}

// SetWidget sets field value
func (o *WidgetUpdateRequest) SetWidget(v KalturaWidget) {
	o.Widget = v
}

func (o WidgetUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["widget"] = o.Widget
	}
	return json.Marshal(toSerialize)
}

type NullableWidgetUpdateRequest struct {
	value *WidgetUpdateRequest
	isSet bool
}

func (v NullableWidgetUpdateRequest) Get() *WidgetUpdateRequest {
	return v.value
}

func (v *NullableWidgetUpdateRequest) Set(val *WidgetUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetUpdateRequest(val *WidgetUpdateRequest) *NullableWidgetUpdateRequest {
	return &NullableWidgetUpdateRequest{value: val, isSet: true}
}

func (v NullableWidgetUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


