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

// KalturaReportExportParams struct for KalturaReportExportParams
type KalturaReportExportParams struct {
	BaseUrl *string `json:"baseUrl,omitempty"`
	RecipientEmail *string `json:"recipientEmail,omitempty"`
	ReportItems []KalturaReportExportItem `json:"reportItems,omitempty"`
	ReportsItemsGroup *string `json:"reportsItemsGroup,omitempty"`
	// Time zone offset in minutes (between client to UTC)
	TimeZoneOffset *int32 `json:"timeZoneOffset,omitempty"`
}

// NewKalturaReportExportParams instantiates a new KalturaReportExportParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaReportExportParams() *KalturaReportExportParams {
	this := KalturaReportExportParams{}
	return &this
}

// NewKalturaReportExportParamsWithDefaults instantiates a new KalturaReportExportParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaReportExportParamsWithDefaults() *KalturaReportExportParams {
	this := KalturaReportExportParams{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise.
func (o *KalturaReportExportParams) GetBaseUrl() string {
	if o == nil || o.BaseUrl == nil {
		var ret string
		return ret
	}
	return *o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportExportParams) GetBaseUrlOk() (*string, bool) {
	if o == nil || o.BaseUrl == nil {
		return nil, false
	}
	return o.BaseUrl, true
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *KalturaReportExportParams) HasBaseUrl() bool {
	if o != nil && o.BaseUrl != nil {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given string and assigns it to the BaseUrl field.
func (o *KalturaReportExportParams) SetBaseUrl(v string) {
	o.BaseUrl = &v
}

// GetRecipientEmail returns the RecipientEmail field value if set, zero value otherwise.
func (o *KalturaReportExportParams) GetRecipientEmail() string {
	if o == nil || o.RecipientEmail == nil {
		var ret string
		return ret
	}
	return *o.RecipientEmail
}

// GetRecipientEmailOk returns a tuple with the RecipientEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportExportParams) GetRecipientEmailOk() (*string, bool) {
	if o == nil || o.RecipientEmail == nil {
		return nil, false
	}
	return o.RecipientEmail, true
}

// HasRecipientEmail returns a boolean if a field has been set.
func (o *KalturaReportExportParams) HasRecipientEmail() bool {
	if o != nil && o.RecipientEmail != nil {
		return true
	}

	return false
}

// SetRecipientEmail gets a reference to the given string and assigns it to the RecipientEmail field.
func (o *KalturaReportExportParams) SetRecipientEmail(v string) {
	o.RecipientEmail = &v
}

// GetReportItems returns the ReportItems field value if set, zero value otherwise.
func (o *KalturaReportExportParams) GetReportItems() []KalturaReportExportItem {
	if o == nil || o.ReportItems == nil {
		var ret []KalturaReportExportItem
		return ret
	}
	return o.ReportItems
}

// GetReportItemsOk returns a tuple with the ReportItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportExportParams) GetReportItemsOk() ([]KalturaReportExportItem, bool) {
	if o == nil || o.ReportItems == nil {
		return nil, false
	}
	return o.ReportItems, true
}

// HasReportItems returns a boolean if a field has been set.
func (o *KalturaReportExportParams) HasReportItems() bool {
	if o != nil && o.ReportItems != nil {
		return true
	}

	return false
}

// SetReportItems gets a reference to the given []KalturaReportExportItem and assigns it to the ReportItems field.
func (o *KalturaReportExportParams) SetReportItems(v []KalturaReportExportItem) {
	o.ReportItems = v
}

// GetReportsItemsGroup returns the ReportsItemsGroup field value if set, zero value otherwise.
func (o *KalturaReportExportParams) GetReportsItemsGroup() string {
	if o == nil || o.ReportsItemsGroup == nil {
		var ret string
		return ret
	}
	return *o.ReportsItemsGroup
}

// GetReportsItemsGroupOk returns a tuple with the ReportsItemsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportExportParams) GetReportsItemsGroupOk() (*string, bool) {
	if o == nil || o.ReportsItemsGroup == nil {
		return nil, false
	}
	return o.ReportsItemsGroup, true
}

// HasReportsItemsGroup returns a boolean if a field has been set.
func (o *KalturaReportExportParams) HasReportsItemsGroup() bool {
	if o != nil && o.ReportsItemsGroup != nil {
		return true
	}

	return false
}

// SetReportsItemsGroup gets a reference to the given string and assigns it to the ReportsItemsGroup field.
func (o *KalturaReportExportParams) SetReportsItemsGroup(v string) {
	o.ReportsItemsGroup = &v
}

// GetTimeZoneOffset returns the TimeZoneOffset field value if set, zero value otherwise.
func (o *KalturaReportExportParams) GetTimeZoneOffset() int32 {
	if o == nil || o.TimeZoneOffset == nil {
		var ret int32
		return ret
	}
	return *o.TimeZoneOffset
}

// GetTimeZoneOffsetOk returns a tuple with the TimeZoneOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaReportExportParams) GetTimeZoneOffsetOk() (*int32, bool) {
	if o == nil || o.TimeZoneOffset == nil {
		return nil, false
	}
	return o.TimeZoneOffset, true
}

// HasTimeZoneOffset returns a boolean if a field has been set.
func (o *KalturaReportExportParams) HasTimeZoneOffset() bool {
	if o != nil && o.TimeZoneOffset != nil {
		return true
	}

	return false
}

// SetTimeZoneOffset gets a reference to the given int32 and assigns it to the TimeZoneOffset field.
func (o *KalturaReportExportParams) SetTimeZoneOffset(v int32) {
	o.TimeZoneOffset = &v
}

func (o KalturaReportExportParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseUrl != nil {
		toSerialize["baseUrl"] = o.BaseUrl
	}
	if o.RecipientEmail != nil {
		toSerialize["recipientEmail"] = o.RecipientEmail
	}
	if o.ReportItems != nil {
		toSerialize["reportItems"] = o.ReportItems
	}
	if o.ReportsItemsGroup != nil {
		toSerialize["reportsItemsGroup"] = o.ReportsItemsGroup
	}
	if o.TimeZoneOffset != nil {
		toSerialize["timeZoneOffset"] = o.TimeZoneOffset
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaReportExportParams struct {
	value *KalturaReportExportParams
	isSet bool
}

func (v NullableKalturaReportExportParams) Get() *KalturaReportExportParams {
	return v.value
}

func (v *NullableKalturaReportExportParams) Set(val *KalturaReportExportParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaReportExportParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaReportExportParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaReportExportParams(val *KalturaReportExportParams) *NullableKalturaReportExportParams {
	return &NullableKalturaReportExportParams{value: val, isSet: true}
}

func (v NullableKalturaReportExportParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaReportExportParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


