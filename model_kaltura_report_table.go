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

// KalturaReportTable struct for KalturaReportTable
type KalturaReportTable struct {
	// `readOnly`
	Data *string `json:"data,omitempty"`
	// `readOnly`
	Header *string `json:"header,omitempty"`
	// `readOnly`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewKalturaReportTable instantiates a new KalturaReportTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaReportTable() *KalturaReportTable {
	this := KalturaReportTable{}
	return &this
}

// NewKalturaReportTableWithDefaults instantiates a new KalturaReportTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaReportTableWithDefaults() *KalturaReportTable {
	this := KalturaReportTable{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KalturaReportTable) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportTable) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KalturaReportTable) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *KalturaReportTable) SetData(v string) {
	o.Data = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *KalturaReportTable) GetHeader() string {
	if o == nil || o.Header == nil {
		var ret string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportTable) GetHeaderOk() (*string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *KalturaReportTable) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given string and assigns it to the Header field.
func (o *KalturaReportTable) SetHeader(v string) {
	o.Header = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *KalturaReportTable) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportTable) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *KalturaReportTable) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *KalturaReportTable) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o KalturaReportTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaReportTable struct {
	value *KalturaReportTable
	isSet bool
}

func (v NullableKalturaReportTable) Get() *KalturaReportTable {
	return v.value
}

func (v *NullableKalturaReportTable) Set(val *KalturaReportTable) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaReportTable) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaReportTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaReportTable(val *KalturaReportTable) *NullableKalturaReportTable {
	return &NullableKalturaReportTable{value: val, isSet: true}
}

func (v NullableKalturaReportTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaReportTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

