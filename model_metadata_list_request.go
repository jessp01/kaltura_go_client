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

// MetadataListRequest struct for MetadataListRequest
type MetadataListRequest struct {
	Filter *KalturaMetadataFilter `json:"filter,omitempty"`
	Pager *KalturaFilterPager `json:"pager,omitempty"`
}

// NewMetadataListRequest instantiates a new MetadataListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataListRequest() *MetadataListRequest {
	this := MetadataListRequest{}
	return &this
}

// NewMetadataListRequestWithDefaults instantiates a new MetadataListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataListRequestWithDefaults() *MetadataListRequest {
	this := MetadataListRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *MetadataListRequest) GetFilter() KalturaMetadataFilter {
	if o == nil || o.Filter == nil {
		var ret KalturaMetadataFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataListRequest) GetFilterOk() (*KalturaMetadataFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *MetadataListRequest) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given KalturaMetadataFilter and assigns it to the Filter field.
func (o *MetadataListRequest) SetFilter(v KalturaMetadataFilter) {
	o.Filter = &v
}

// GetPager returns the Pager field value if set, zero value otherwise.
func (o *MetadataListRequest) GetPager() KalturaFilterPager {
	if o == nil || o.Pager == nil {
		var ret KalturaFilterPager
		return ret
	}
	return *o.Pager
}

// GetPagerOk returns a tuple with the Pager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataListRequest) GetPagerOk() (*KalturaFilterPager, bool) {
	if o == nil || o.Pager == nil {
		return nil, false
	}
	return o.Pager, true
}

// HasPager returns a boolean if a field has been set.
func (o *MetadataListRequest) HasPager() bool {
	if o != nil && o.Pager != nil {
		return true
	}

	return false
}

// SetPager gets a reference to the given KalturaFilterPager and assigns it to the Pager field.
func (o *MetadataListRequest) SetPager(v KalturaFilterPager) {
	o.Pager = &v
}

func (o MetadataListRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.Pager != nil {
		toSerialize["pager"] = o.Pager
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataListRequest struct {
	value *MetadataListRequest
	isSet bool
}

func (v NullableMetadataListRequest) Get() *MetadataListRequest {
	return v.value
}

func (v *NullableMetadataListRequest) Set(val *MetadataListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataListRequest(val *MetadataListRequest) *NullableMetadataListRequest {
	return &NullableMetadataListRequest{value: val, isSet: true}
}

func (v NullableMetadataListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


