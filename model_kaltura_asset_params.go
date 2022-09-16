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

// KalturaAssetParams struct for KalturaAssetParams
type KalturaAssetParams struct {
	// `readOnly`  Creation date as Unix timestamp (In seconds)
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// The description of the Flavor Params
	Description *string `json:"description,omitempty"`
	// `readOnly`  The id of the Flavor Params
	Id *int32 `json:"id,omitempty"`
	// `readOnly`  Enum Type: `KalturaNullableBoolean`  True if those Flavor Params are part of system defaults
	IsSystemDefault *int32 `json:"isSystemDefault,omitempty"`
	// Enum Type: `KalturaMediaParserType`  Media parser type to be used for post-conversion validation
	MediaParserType *string `json:"mediaParserType,omitempty"`
	// The name of the Flavor Params
	Name *string `json:"name,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// Comma seperated ids of remote storage profiles that the flavor distributed to, the distribution done by the conversion engine
	RemoteStorageProfileIds *int32 `json:"remoteStorageProfileIds,omitempty"`
	RequiredPermissions []KalturaString `json:"requiredPermissions,omitempty"`
	// Comma seperated ids of source flavor params this flavor is created from
	SourceAssetParamsIds *string `json:"sourceAssetParamsIds,omitempty"`
	// Id of remote storage profile that used to get the source, zero indicates Kaltura data center
	SourceRemoteStorageProfileId *int32 `json:"sourceRemoteStorageProfileId,omitempty"`
	// System name of the Flavor Params
	SystemName *string `json:"systemName,omitempty"`
	// The Flavor Params tags are used to identify the flavor for different usage (e.g. web, hd, mobile)
	Tags *string `json:"tags,omitempty"`
}

// NewKalturaAssetParams instantiates a new KalturaAssetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAssetParams() *KalturaAssetParams {
	this := KalturaAssetParams{}
	return &this
}

// NewKalturaAssetParamsWithDefaults instantiates a new KalturaAssetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetParamsWithDefaults() *KalturaAssetParams {
	this := KalturaAssetParams{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaAssetParams) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaAssetParams) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaAssetParams) SetId(v int32) {
	o.Id = &v
}

// GetIsSystemDefault returns the IsSystemDefault field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetIsSystemDefault() int32 {
	if o == nil || o.IsSystemDefault == nil {
		var ret int32
		return ret
	}
	return *o.IsSystemDefault
}

// GetIsSystemDefaultOk returns a tuple with the IsSystemDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetIsSystemDefaultOk() (*int32, bool) {
	if o == nil || o.IsSystemDefault == nil {
		return nil, false
	}
	return o.IsSystemDefault, true
}

// HasIsSystemDefault returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasIsSystemDefault() bool {
	if o != nil && o.IsSystemDefault != nil {
		return true
	}

	return false
}

// SetIsSystemDefault gets a reference to the given int32 and assigns it to the IsSystemDefault field.
func (o *KalturaAssetParams) SetIsSystemDefault(v int32) {
	o.IsSystemDefault = &v
}

// GetMediaParserType returns the MediaParserType field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetMediaParserType() string {
	if o == nil || o.MediaParserType == nil {
		var ret string
		return ret
	}
	return *o.MediaParserType
}

// GetMediaParserTypeOk returns a tuple with the MediaParserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetMediaParserTypeOk() (*string, bool) {
	if o == nil || o.MediaParserType == nil {
		return nil, false
	}
	return o.MediaParserType, true
}

// HasMediaParserType returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasMediaParserType() bool {
	if o != nil && o.MediaParserType != nil {
		return true
	}

	return false
}

// SetMediaParserType gets a reference to the given string and assigns it to the MediaParserType field.
func (o *KalturaAssetParams) SetMediaParserType(v string) {
	o.MediaParserType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaAssetParams) SetName(v string) {
	o.Name = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaAssetParams) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaAssetParams) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetRemoteStorageProfileIds returns the RemoteStorageProfileIds field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetRemoteStorageProfileIds() int32 {
	if o == nil || o.RemoteStorageProfileIds == nil {
		var ret int32
		return ret
	}
	return *o.RemoteStorageProfileIds
}

// GetRemoteStorageProfileIdsOk returns a tuple with the RemoteStorageProfileIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetRemoteStorageProfileIdsOk() (*int32, bool) {
	if o == nil || o.RemoteStorageProfileIds == nil {
		return nil, false
	}
	return o.RemoteStorageProfileIds, true
}

// HasRemoteStorageProfileIds returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasRemoteStorageProfileIds() bool {
	if o != nil && o.RemoteStorageProfileIds != nil {
		return true
	}

	return false
}

// SetRemoteStorageProfileIds gets a reference to the given int32 and assigns it to the RemoteStorageProfileIds field.
func (o *KalturaAssetParams) SetRemoteStorageProfileIds(v int32) {
	o.RemoteStorageProfileIds = &v
}

// GetRequiredPermissions returns the RequiredPermissions field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetRequiredPermissions() []KalturaString {
	if o == nil || o.RequiredPermissions == nil {
		var ret []KalturaString
		return ret
	}
	return o.RequiredPermissions
}

// GetRequiredPermissionsOk returns a tuple with the RequiredPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetRequiredPermissionsOk() ([]KalturaString, bool) {
	if o == nil || o.RequiredPermissions == nil {
		return nil, false
	}
	return o.RequiredPermissions, true
}

// HasRequiredPermissions returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasRequiredPermissions() bool {
	if o != nil && o.RequiredPermissions != nil {
		return true
	}

	return false
}

// SetRequiredPermissions gets a reference to the given []KalturaString and assigns it to the RequiredPermissions field.
func (o *KalturaAssetParams) SetRequiredPermissions(v []KalturaString) {
	o.RequiredPermissions = v
}

// GetSourceAssetParamsIds returns the SourceAssetParamsIds field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetSourceAssetParamsIds() string {
	if o == nil || o.SourceAssetParamsIds == nil {
		var ret string
		return ret
	}
	return *o.SourceAssetParamsIds
}

// GetSourceAssetParamsIdsOk returns a tuple with the SourceAssetParamsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetSourceAssetParamsIdsOk() (*string, bool) {
	if o == nil || o.SourceAssetParamsIds == nil {
		return nil, false
	}
	return o.SourceAssetParamsIds, true
}

// HasSourceAssetParamsIds returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasSourceAssetParamsIds() bool {
	if o != nil && o.SourceAssetParamsIds != nil {
		return true
	}

	return false
}

// SetSourceAssetParamsIds gets a reference to the given string and assigns it to the SourceAssetParamsIds field.
func (o *KalturaAssetParams) SetSourceAssetParamsIds(v string) {
	o.SourceAssetParamsIds = &v
}

// GetSourceRemoteStorageProfileId returns the SourceRemoteStorageProfileId field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetSourceRemoteStorageProfileId() int32 {
	if o == nil || o.SourceRemoteStorageProfileId == nil {
		var ret int32
		return ret
	}
	return *o.SourceRemoteStorageProfileId
}

// GetSourceRemoteStorageProfileIdOk returns a tuple with the SourceRemoteStorageProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetSourceRemoteStorageProfileIdOk() (*int32, bool) {
	if o == nil || o.SourceRemoteStorageProfileId == nil {
		return nil, false
	}
	return o.SourceRemoteStorageProfileId, true
}

// HasSourceRemoteStorageProfileId returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasSourceRemoteStorageProfileId() bool {
	if o != nil && o.SourceRemoteStorageProfileId != nil {
		return true
	}

	return false
}

// SetSourceRemoteStorageProfileId gets a reference to the given int32 and assigns it to the SourceRemoteStorageProfileId field.
func (o *KalturaAssetParams) SetSourceRemoteStorageProfileId(v int32) {
	o.SourceRemoteStorageProfileId = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaAssetParams) SetSystemName(v string) {
	o.SystemName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *KalturaAssetParams) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAssetParams) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *KalturaAssetParams) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *KalturaAssetParams) SetTags(v string) {
	o.Tags = &v
}

func (o KalturaAssetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsSystemDefault != nil {
		toSerialize["isSystemDefault"] = o.IsSystemDefault
	}
	if o.MediaParserType != nil {
		toSerialize["mediaParserType"] = o.MediaParserType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.RemoteStorageProfileIds != nil {
		toSerialize["remoteStorageProfileIds"] = o.RemoteStorageProfileIds
	}
	if o.RequiredPermissions != nil {
		toSerialize["requiredPermissions"] = o.RequiredPermissions
	}
	if o.SourceAssetParamsIds != nil {
		toSerialize["sourceAssetParamsIds"] = o.SourceAssetParamsIds
	}
	if o.SourceRemoteStorageProfileId != nil {
		toSerialize["sourceRemoteStorageProfileId"] = o.SourceRemoteStorageProfileId
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAssetParams struct {
	value *KalturaAssetParams
	isSet bool
}

func (v NullableKalturaAssetParams) Get() *KalturaAssetParams {
	return v.value
}

func (v *NullableKalturaAssetParams) Set(val *KalturaAssetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAssetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAssetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAssetParams(val *KalturaAssetParams) *NullableKalturaAssetParams {
	return &NullableKalturaAssetParams{value: val, isSet: true}
}

func (v NullableKalturaAssetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAssetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

