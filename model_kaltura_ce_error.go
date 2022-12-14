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

// KalturaCEError struct for KalturaCEError
type KalturaCEError struct {
	Browser *string `json:"browser,omitempty"`
	CeAdminEmail *string `json:"ceAdminEmail,omitempty"`
	Data *string `json:"data,omitempty"`
	Description *string `json:"description,omitempty"`
	// `readOnly`
	Id *string `json:"id,omitempty"`
	PartnerId *int32 `json:"partnerId,omitempty"`
	PhpVersion *string `json:"phpVersion,omitempty"`
	ServerIp *string `json:"serverIp,omitempty"`
	ServerOs *string `json:"serverOs,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewKalturaCEError instantiates a new KalturaCEError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaCEError() *KalturaCEError {
	this := KalturaCEError{}
	return &this
}

// NewKalturaCEErrorWithDefaults instantiates a new KalturaCEError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaCEErrorWithDefaults() *KalturaCEError {
	this := KalturaCEError{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *KalturaCEError) GetBrowser() string {
	if o == nil || o.Browser == nil {
		var ret string
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetBrowserOk() (*string, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *KalturaCEError) HasBrowser() bool {
	if o != nil && o.Browser != nil {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given string and assigns it to the Browser field.
func (o *KalturaCEError) SetBrowser(v string) {
	o.Browser = &v
}

// GetCeAdminEmail returns the CeAdminEmail field value if set, zero value otherwise.
func (o *KalturaCEError) GetCeAdminEmail() string {
	if o == nil || o.CeAdminEmail == nil {
		var ret string
		return ret
	}
	return *o.CeAdminEmail
}

// GetCeAdminEmailOk returns a tuple with the CeAdminEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetCeAdminEmailOk() (*string, bool) {
	if o == nil || o.CeAdminEmail == nil {
		return nil, false
	}
	return o.CeAdminEmail, true
}

// HasCeAdminEmail returns a boolean if a field has been set.
func (o *KalturaCEError) HasCeAdminEmail() bool {
	if o != nil && o.CeAdminEmail != nil {
		return true
	}

	return false
}

// SetCeAdminEmail gets a reference to the given string and assigns it to the CeAdminEmail field.
func (o *KalturaCEError) SetCeAdminEmail(v string) {
	o.CeAdminEmail = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *KalturaCEError) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *KalturaCEError) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *KalturaCEError) SetData(v string) {
	o.Data = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaCEError) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaCEError) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaCEError) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaCEError) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaCEError) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaCEError) SetId(v string) {
	o.Id = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaCEError) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaCEError) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaCEError) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPhpVersion returns the PhpVersion field value if set, zero value otherwise.
func (o *KalturaCEError) GetPhpVersion() string {
	if o == nil || o.PhpVersion == nil {
		var ret string
		return ret
	}
	return *o.PhpVersion
}

// GetPhpVersionOk returns a tuple with the PhpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetPhpVersionOk() (*string, bool) {
	if o == nil || o.PhpVersion == nil {
		return nil, false
	}
	return o.PhpVersion, true
}

// HasPhpVersion returns a boolean if a field has been set.
func (o *KalturaCEError) HasPhpVersion() bool {
	if o != nil && o.PhpVersion != nil {
		return true
	}

	return false
}

// SetPhpVersion gets a reference to the given string and assigns it to the PhpVersion field.
func (o *KalturaCEError) SetPhpVersion(v string) {
	o.PhpVersion = &v
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise.
func (o *KalturaCEError) GetServerIp() string {
	if o == nil || o.ServerIp == nil {
		var ret string
		return ret
	}
	return *o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetServerIpOk() (*string, bool) {
	if o == nil || o.ServerIp == nil {
		return nil, false
	}
	return o.ServerIp, true
}

// HasServerIp returns a boolean if a field has been set.
func (o *KalturaCEError) HasServerIp() bool {
	if o != nil && o.ServerIp != nil {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given string and assigns it to the ServerIp field.
func (o *KalturaCEError) SetServerIp(v string) {
	o.ServerIp = &v
}

// GetServerOs returns the ServerOs field value if set, zero value otherwise.
func (o *KalturaCEError) GetServerOs() string {
	if o == nil || o.ServerOs == nil {
		var ret string
		return ret
	}
	return *o.ServerOs
}

// GetServerOsOk returns a tuple with the ServerOs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetServerOsOk() (*string, bool) {
	if o == nil || o.ServerOs == nil {
		return nil, false
	}
	return o.ServerOs, true
}

// HasServerOs returns a boolean if a field has been set.
func (o *KalturaCEError) HasServerOs() bool {
	if o != nil && o.ServerOs != nil {
		return true
	}

	return false
}

// SetServerOs gets a reference to the given string and assigns it to the ServerOs field.
func (o *KalturaCEError) SetServerOs(v string) {
	o.ServerOs = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaCEError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaCEError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaCEError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KalturaCEError) SetType(v string) {
	o.Type = &v
}

func (o KalturaCEError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Browser != nil {
		toSerialize["browser"] = o.Browser
	}
	if o.CeAdminEmail != nil {
		toSerialize["ceAdminEmail"] = o.CeAdminEmail
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.PhpVersion != nil {
		toSerialize["phpVersion"] = o.PhpVersion
	}
	if o.ServerIp != nil {
		toSerialize["serverIp"] = o.ServerIp
	}
	if o.ServerOs != nil {
		toSerialize["serverOs"] = o.ServerOs
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaCEError struct {
	value *KalturaCEError
	isSet bool
}

func (v NullableKalturaCEError) Get() *KalturaCEError {
	return v.value
}

func (v *NullableKalturaCEError) Set(val *KalturaCEError) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaCEError) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaCEError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaCEError(val *KalturaCEError) *NullableKalturaCEError {
	return &NullableKalturaCEError{value: val, isSet: true}
}

func (v NullableKalturaCEError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaCEError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


