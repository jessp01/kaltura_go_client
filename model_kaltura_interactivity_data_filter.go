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

// KalturaInteractivityDataFilter struct for KalturaInteractivityDataFilter
type KalturaInteractivityDataFilter struct {
	InteractionFilter *KalturaInteractivityInteractionFilter `json:"interactionFilter,omitempty"`
	NodeFilter *KalturaInteractivityNodeFilter `json:"nodeFilter,omitempty"`
	RootFilter *KalturaInteractivityRootFilter `json:"rootFilter,omitempty"`
}

// NewKalturaInteractivityDataFilter instantiates a new KalturaInteractivityDataFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaInteractivityDataFilter() *KalturaInteractivityDataFilter {
	this := KalturaInteractivityDataFilter{}
	return &this
}

// NewKalturaInteractivityDataFilterWithDefaults instantiates a new KalturaInteractivityDataFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaInteractivityDataFilterWithDefaults() *KalturaInteractivityDataFilter {
	this := KalturaInteractivityDataFilter{}
	return &this
}

// GetInteractionFilter returns the InteractionFilter field value if set, zero value otherwise.
func (o *KalturaInteractivityDataFilter) GetInteractionFilter() KalturaInteractivityInteractionFilter {
	if o == nil || o.InteractionFilter == nil {
		var ret KalturaInteractivityInteractionFilter
		return ret
	}
	return *o.InteractionFilter
}

// GetInteractionFilterOk returns a tuple with the InteractionFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaInteractivityDataFilter) GetInteractionFilterOk() (*KalturaInteractivityInteractionFilter, bool) {
	if o == nil || o.InteractionFilter == nil {
		return nil, false
	}
	return o.InteractionFilter, true
}

// HasInteractionFilter returns a boolean if a field has been set.
func (o *KalturaInteractivityDataFilter) HasInteractionFilter() bool {
	if o != nil && o.InteractionFilter != nil {
		return true
	}

	return false
}

// SetInteractionFilter gets a reference to the given KalturaInteractivityInteractionFilter and assigns it to the InteractionFilter field.
func (o *KalturaInteractivityDataFilter) SetInteractionFilter(v KalturaInteractivityInteractionFilter) {
	o.InteractionFilter = &v
}

// GetNodeFilter returns the NodeFilter field value if set, zero value otherwise.
func (o *KalturaInteractivityDataFilter) GetNodeFilter() KalturaInteractivityNodeFilter {
	if o == nil || o.NodeFilter == nil {
		var ret KalturaInteractivityNodeFilter
		return ret
	}
	return *o.NodeFilter
}

// GetNodeFilterOk returns a tuple with the NodeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaInteractivityDataFilter) GetNodeFilterOk() (*KalturaInteractivityNodeFilter, bool) {
	if o == nil || o.NodeFilter == nil {
		return nil, false
	}
	return o.NodeFilter, true
}

// HasNodeFilter returns a boolean if a field has been set.
func (o *KalturaInteractivityDataFilter) HasNodeFilter() bool {
	if o != nil && o.NodeFilter != nil {
		return true
	}

	return false
}

// SetNodeFilter gets a reference to the given KalturaInteractivityNodeFilter and assigns it to the NodeFilter field.
func (o *KalturaInteractivityDataFilter) SetNodeFilter(v KalturaInteractivityNodeFilter) {
	o.NodeFilter = &v
}

// GetRootFilter returns the RootFilter field value if set, zero value otherwise.
func (o *KalturaInteractivityDataFilter) GetRootFilter() KalturaInteractivityRootFilter {
	if o == nil || o.RootFilter == nil {
		var ret KalturaInteractivityRootFilter
		return ret
	}
	return *o.RootFilter
}

// GetRootFilterOk returns a tuple with the RootFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaInteractivityDataFilter) GetRootFilterOk() (*KalturaInteractivityRootFilter, bool) {
	if o == nil || o.RootFilter == nil {
		return nil, false
	}
	return o.RootFilter, true
}

// HasRootFilter returns a boolean if a field has been set.
func (o *KalturaInteractivityDataFilter) HasRootFilter() bool {
	if o != nil && o.RootFilter != nil {
		return true
	}

	return false
}

// SetRootFilter gets a reference to the given KalturaInteractivityRootFilter and assigns it to the RootFilter field.
func (o *KalturaInteractivityDataFilter) SetRootFilter(v KalturaInteractivityRootFilter) {
	o.RootFilter = &v
}

func (o KalturaInteractivityDataFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InteractionFilter != nil {
		toSerialize["interactionFilter"] = o.InteractionFilter
	}
	if o.NodeFilter != nil {
		toSerialize["nodeFilter"] = o.NodeFilter
	}
	if o.RootFilter != nil {
		toSerialize["rootFilter"] = o.RootFilter
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaInteractivityDataFilter struct {
	value *KalturaInteractivityDataFilter
	isSet bool
}

func (v NullableKalturaInteractivityDataFilter) Get() *KalturaInteractivityDataFilter {
	return v.value
}

func (v *NullableKalturaInteractivityDataFilter) Set(val *KalturaInteractivityDataFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaInteractivityDataFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaInteractivityDataFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaInteractivityDataFilter(val *KalturaInteractivityDataFilter) *NullableKalturaInteractivityDataFilter {
	return &NullableKalturaInteractivityDataFilter{value: val, isSet: true}
}

func (v NullableKalturaInteractivityDataFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaInteractivityDataFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


