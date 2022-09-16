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

// BaseEntryListFlagsRequest struct for BaseEntryListFlagsRequest
type BaseEntryListFlagsRequest struct {
	EntryId string `json:"entryId"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewBaseEntryListFlagsRequest instantiates a new BaseEntryListFlagsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseEntryListFlagsRequest(entryId string) *BaseEntryListFlagsRequest {
	this := BaseEntryListFlagsRequest{}
	this.EntryId = entryId
	return &this
}

// NewBaseEntryListFlagsRequestWithDefaults instantiates a new BaseEntryListFlagsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseEntryListFlagsRequestWithDefaults() *BaseEntryListFlagsRequest {
	this := BaseEntryListFlagsRequest{}
	return &this
}

// GetEntryId returns the EntryId field value
func (o *BaseEntryListFlagsRequest) GetEntryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value
// and a boolean to check if the value has been set.
func (o *BaseEntryListFlagsRequest) GetEntryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntryId, true
}

// SetEntryId sets field value
func (o *BaseEntryListFlagsRequest) SetEntryId(v string) {
	o.EntryId = v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *BaseEntryListFlagsRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseEntryListFlagsRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *BaseEntryListFlagsRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *BaseEntryListFlagsRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o BaseEntryListFlagsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["entryId"] = o.EntryId
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableBaseEntryListFlagsRequest struct {
	value *BaseEntryListFlagsRequest
	isSet bool
}

func (v NullableBaseEntryListFlagsRequest) Get() *BaseEntryListFlagsRequest {
	return v.value
}

func (v *NullableBaseEntryListFlagsRequest) Set(val *BaseEntryListFlagsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseEntryListFlagsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseEntryListFlagsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseEntryListFlagsRequest(val *BaseEntryListFlagsRequest) *NullableBaseEntryListFlagsRequest {
	return &NullableBaseEntryListFlagsRequest{value: val, isSet: true}
}

func (v NullableBaseEntryListFlagsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseEntryListFlagsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


