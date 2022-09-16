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

// KalturaDeliveryProfile struct for KalturaDeliveryProfile
type KalturaDeliveryProfile struct {
	// `readOnly`  Creation time as Unix timestamp (In seconds)
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// The description of the Delivery
	Description *string `json:"description,omitempty"`
	// Extra query string parameters that should be added to the url
	ExtraParams *string `json:"extraParams,omitempty"`
	// `readOnly`  the host part of the url
	HostName *string `json:"hostName,omitempty"`
	// `readOnly`  The id of the Delivery
	Id *int32 `json:"id,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`  True if this is the systemwide default for the protocol
	IsDefault *int32 `json:"isDefault,omitempty"`
	// Comma separated list of supported media protocols. f.i. rtmpe
	MediaProtocols *string `json:"mediaProtocols,omitempty"`
	// The name of the Delivery
	Name *string `json:"name,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// `readOnly`  the object from which this object was cloned (or 0)
	ParentId *int32 `json:"parentId,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// priority used for ordering similar delivery profiles
	Priority *int32 `json:"priority,omitempty"`
	Recognizer *KalturaUrlRecognizer `json:"recognizer,omitempty"`
	// Enum Type: `KalturaDeliveryStatus`
	Status *int32 `json:"status,omitempty"`
	// Enum Type: `KalturaPlaybackProtocol`
	StreamerType *string `json:"streamerType,omitempty"`
	SupplementaryAssetsFilter *KalturaAssetFilter `json:"supplementaryAssetsFilter,omitempty"`
	// System name of the delivery
	SystemName *string `json:"systemName,omitempty"`
	Tokenizer *KalturaUrlTokenizer `json:"tokenizer,omitempty"`
	// Enum Type: `KalturaDeliveryProfileType`  Delivery type
	Type *string `json:"type,omitempty"`
	// `readOnly`  Update time as Unix timestamp (In seconds)
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewKalturaDeliveryProfile instantiates a new KalturaDeliveryProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaDeliveryProfile() *KalturaDeliveryProfile {
	this := KalturaDeliveryProfile{}
	return &this
}

// NewKalturaDeliveryProfileWithDefaults instantiates a new KalturaDeliveryProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaDeliveryProfileWithDefaults() *KalturaDeliveryProfile {
	this := KalturaDeliveryProfile{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaDeliveryProfile) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaDeliveryProfile) SetDescription(v string) {
	o.Description = &v
}

// GetExtraParams returns the ExtraParams field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetExtraParams() string {
	if o == nil || o.ExtraParams == nil {
		var ret string
		return ret
	}
	return *o.ExtraParams
}

// GetExtraParamsOk returns a tuple with the ExtraParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetExtraParamsOk() (*string, bool) {
	if o == nil || o.ExtraParams == nil {
		return nil, false
	}
	return o.ExtraParams, true
}

// HasExtraParams returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasExtraParams() bool {
	if o != nil && o.ExtraParams != nil {
		return true
	}

	return false
}

// SetExtraParams gets a reference to the given string and assigns it to the ExtraParams field.
func (o *KalturaDeliveryProfile) SetExtraParams(v string) {
	o.ExtraParams = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *KalturaDeliveryProfile) SetHostName(v string) {
	o.HostName = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaDeliveryProfile) SetId(v int32) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetIsDefault() int32 {
	if o == nil || o.IsDefault == nil {
		var ret int32
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetIsDefaultOk() (*int32, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given int32 and assigns it to the IsDefault field.
func (o *KalturaDeliveryProfile) SetIsDefault(v int32) {
	o.IsDefault = &v
}

// GetMediaProtocols returns the MediaProtocols field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetMediaProtocols() string {
	if o == nil || o.MediaProtocols == nil {
		var ret string
		return ret
	}
	return *o.MediaProtocols
}

// GetMediaProtocolsOk returns a tuple with the MediaProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetMediaProtocolsOk() (*string, bool) {
	if o == nil || o.MediaProtocols == nil {
		return nil, false
	}
	return o.MediaProtocols, true
}

// HasMediaProtocols returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasMediaProtocols() bool {
	if o != nil && o.MediaProtocols != nil {
		return true
	}

	return false
}

// SetMediaProtocols gets a reference to the given string and assigns it to the MediaProtocols field.
func (o *KalturaDeliveryProfile) SetMediaProtocols(v string) {
	o.MediaProtocols = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaDeliveryProfile) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaDeliveryProfile) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetParentId() int32 {
	if o == nil || o.ParentId == nil {
		var ret int32
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetParentIdOk() (*int32, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given int32 and assigns it to the ParentId field.
func (o *KalturaDeliveryProfile) SetParentId(v int32) {
	o.ParentId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaDeliveryProfile) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetPriority() int32 {
	if o == nil || o.Priority == nil {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetPriorityOk() (*int32, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasPriority() bool {
	if o != nil && o.Priority != nil {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *KalturaDeliveryProfile) SetPriority(v int32) {
	o.Priority = &v
}

// GetRecognizer returns the Recognizer field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetRecognizer() KalturaUrlRecognizer {
	if o == nil || o.Recognizer == nil {
		var ret KalturaUrlRecognizer
		return ret
	}
	return *o.Recognizer
}

// GetRecognizerOk returns a tuple with the Recognizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetRecognizerOk() (*KalturaUrlRecognizer, bool) {
	if o == nil || o.Recognizer == nil {
		return nil, false
	}
	return o.Recognizer, true
}

// HasRecognizer returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasRecognizer() bool {
	if o != nil && o.Recognizer != nil {
		return true
	}

	return false
}

// SetRecognizer gets a reference to the given KalturaUrlRecognizer and assigns it to the Recognizer field.
func (o *KalturaDeliveryProfile) SetRecognizer(v KalturaUrlRecognizer) {
	o.Recognizer = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaDeliveryProfile) SetStatus(v int32) {
	o.Status = &v
}

// GetStreamerType returns the StreamerType field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetStreamerType() string {
	if o == nil || o.StreamerType == nil {
		var ret string
		return ret
	}
	return *o.StreamerType
}

// GetStreamerTypeOk returns a tuple with the StreamerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetStreamerTypeOk() (*string, bool) {
	if o == nil || o.StreamerType == nil {
		return nil, false
	}
	return o.StreamerType, true
}

// HasStreamerType returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasStreamerType() bool {
	if o != nil && o.StreamerType != nil {
		return true
	}

	return false
}

// SetStreamerType gets a reference to the given string and assigns it to the StreamerType field.
func (o *KalturaDeliveryProfile) SetStreamerType(v string) {
	o.StreamerType = &v
}

// GetSupplementaryAssetsFilter returns the SupplementaryAssetsFilter field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetSupplementaryAssetsFilter() KalturaAssetFilter {
	if o == nil || o.SupplementaryAssetsFilter == nil {
		var ret KalturaAssetFilter
		return ret
	}
	return *o.SupplementaryAssetsFilter
}

// GetSupplementaryAssetsFilterOk returns a tuple with the SupplementaryAssetsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetSupplementaryAssetsFilterOk() (*KalturaAssetFilter, bool) {
	if o == nil || o.SupplementaryAssetsFilter == nil {
		return nil, false
	}
	return o.SupplementaryAssetsFilter, true
}

// HasSupplementaryAssetsFilter returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasSupplementaryAssetsFilter() bool {
	if o != nil && o.SupplementaryAssetsFilter != nil {
		return true
	}

	return false
}

// SetSupplementaryAssetsFilter gets a reference to the given KalturaAssetFilter and assigns it to the SupplementaryAssetsFilter field.
func (o *KalturaDeliveryProfile) SetSupplementaryAssetsFilter(v KalturaAssetFilter) {
	o.SupplementaryAssetsFilter = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaDeliveryProfile) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTokenizer returns the Tokenizer field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetTokenizer() KalturaUrlTokenizer {
	if o == nil || o.Tokenizer == nil {
		var ret KalturaUrlTokenizer
		return ret
	}
	return *o.Tokenizer
}

// GetTokenizerOk returns a tuple with the Tokenizer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetTokenizerOk() (*KalturaUrlTokenizer, bool) {
	if o == nil || o.Tokenizer == nil {
		return nil, false
	}
	return o.Tokenizer, true
}

// HasTokenizer returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasTokenizer() bool {
	if o != nil && o.Tokenizer != nil {
		return true
	}

	return false
}

// SetTokenizer gets a reference to the given KalturaUrlTokenizer and assigns it to the Tokenizer field.
func (o *KalturaDeliveryProfile) SetTokenizer(v KalturaUrlTokenizer) {
	o.Tokenizer = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *KalturaDeliveryProfile) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaDeliveryProfile) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *KalturaDeliveryProfile) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaDeliveryProfile) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *KalturaDeliveryProfile) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *KalturaDeliveryProfile) SetUrl(v string) {
	o.Url = &v
}

func (o KalturaDeliveryProfile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExtraParams != nil {
		toSerialize["extraParams"] = o.ExtraParams
	}
	if o.HostName != nil {
		toSerialize["hostName"] = o.HostName
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.MediaProtocols != nil {
		toSerialize["mediaProtocols"] = o.MediaProtocols
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Recognizer != nil {
		toSerialize["recognizer"] = o.Recognizer
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StreamerType != nil {
		toSerialize["streamerType"] = o.StreamerType
	}
	if o.SupplementaryAssetsFilter != nil {
		toSerialize["supplementaryAssetsFilter"] = o.SupplementaryAssetsFilter
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.Tokenizer != nil {
		toSerialize["tokenizer"] = o.Tokenizer
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaDeliveryProfile struct {
	value *KalturaDeliveryProfile
	isSet bool
}

func (v NullableKalturaDeliveryProfile) Get() *KalturaDeliveryProfile {
	return v.value
}

func (v *NullableKalturaDeliveryProfile) Set(val *KalturaDeliveryProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaDeliveryProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaDeliveryProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaDeliveryProfile(val *KalturaDeliveryProfile) *NullableKalturaDeliveryProfile {
	return &NullableKalturaDeliveryProfile{value: val, isSet: true}
}

func (v NullableKalturaDeliveryProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaDeliveryProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


