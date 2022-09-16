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

// KalturaAsset struct for KalturaAsset
type KalturaAsset struct {
	// Comma separated list of source flavor params ids
	ActualSourceAssetParamsIds *string `json:"actualSourceAssetParamsIds,omitempty"`
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `readOnly`
	DeletedAt *int32 `json:"deletedAt,omitempty"`
	// `readOnly`  System description, error message, warnings and failure cause.
	Description *string `json:"description,omitempty"`
	// `readOnly`  The entry ID of the Flavor Asset
	EntryId *string `json:"entryId,omitempty"`
	// `insertOnly`  The file extension
	FileExt *string `json:"fileExt,omitempty"`
	// `readOnly`  The ID of the Flavor Asset
	Id *string `json:"id,omitempty"`
	ObjectType *string `json:"objectType,omitempty"`
	// Partner private data
	PartnerData *string `json:"partnerData,omitempty"`
	// Partner friendly description
	PartnerDescription *string `json:"partnerDescription,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  The size (in KBytes) of the Flavor Asset
	Size *int32 `json:"size,omitempty"`
	// `readOnly`  The size (in Bytes) of the asset
	SizeInBytes *int32 `json:"sizeInBytes,omitempty"`
	// Tags used to identify the Flavor Asset in various scenarios
	Tags *string `json:"tags,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// `readOnly`  The version of the Flavor Asset
	Version *int32 `json:"version,omitempty"`
}

// NewKalturaAsset instantiates a new KalturaAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaAsset() *KalturaAsset {
	this := KalturaAsset{}
	return &this
}

// NewKalturaAssetWithDefaults instantiates a new KalturaAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaAssetWithDefaults() *KalturaAsset {
	this := KalturaAsset{}
	return &this
}

// GetActualSourceAssetParamsIds returns the ActualSourceAssetParamsIds field value if set, zero value otherwise.
func (o *KalturaAsset) GetActualSourceAssetParamsIds() string {
	if o == nil || o.ActualSourceAssetParamsIds == nil {
		var ret string
		return ret
	}
	return *o.ActualSourceAssetParamsIds
}

// GetActualSourceAssetParamsIdsOk returns a tuple with the ActualSourceAssetParamsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetActualSourceAssetParamsIdsOk() (*string, bool) {
	if o == nil || o.ActualSourceAssetParamsIds == nil {
		return nil, false
	}
	return o.ActualSourceAssetParamsIds, true
}

// HasActualSourceAssetParamsIds returns a boolean if a field has been set.
func (o *KalturaAsset) HasActualSourceAssetParamsIds() bool {
	if o != nil && o.ActualSourceAssetParamsIds != nil {
		return true
	}

	return false
}

// SetActualSourceAssetParamsIds gets a reference to the given string and assigns it to the ActualSourceAssetParamsIds field.
func (o *KalturaAsset) SetActualSourceAssetParamsIds(v string) {
	o.ActualSourceAssetParamsIds = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaAsset) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaAsset) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaAsset) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *KalturaAsset) GetDeletedAt() int32 {
	if o == nil || o.DeletedAt == nil {
		var ret int32
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetDeletedAtOk() (*int32, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *KalturaAsset) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given int32 and assigns it to the DeletedAt field.
func (o *KalturaAsset) SetDeletedAt(v int32) {
	o.DeletedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *KalturaAsset) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *KalturaAsset) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *KalturaAsset) SetDescription(v string) {
	o.Description = &v
}

// GetEntryId returns the EntryId field value if set, zero value otherwise.
func (o *KalturaAsset) GetEntryId() string {
	if o == nil || o.EntryId == nil {
		var ret string
		return ret
	}
	return *o.EntryId
}

// GetEntryIdOk returns a tuple with the EntryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetEntryIdOk() (*string, bool) {
	if o == nil || o.EntryId == nil {
		return nil, false
	}
	return o.EntryId, true
}

// HasEntryId returns a boolean if a field has been set.
func (o *KalturaAsset) HasEntryId() bool {
	if o != nil && o.EntryId != nil {
		return true
	}

	return false
}

// SetEntryId gets a reference to the given string and assigns it to the EntryId field.
func (o *KalturaAsset) SetEntryId(v string) {
	o.EntryId = &v
}

// GetFileExt returns the FileExt field value if set, zero value otherwise.
func (o *KalturaAsset) GetFileExt() string {
	if o == nil || o.FileExt == nil {
		var ret string
		return ret
	}
	return *o.FileExt
}

// GetFileExtOk returns a tuple with the FileExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetFileExtOk() (*string, bool) {
	if o == nil || o.FileExt == nil {
		return nil, false
	}
	return o.FileExt, true
}

// HasFileExt returns a boolean if a field has been set.
func (o *KalturaAsset) HasFileExt() bool {
	if o != nil && o.FileExt != nil {
		return true
	}

	return false
}

// SetFileExt gets a reference to the given string and assigns it to the FileExt field.
func (o *KalturaAsset) SetFileExt(v string) {
	o.FileExt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaAsset) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaAsset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaAsset) SetId(v string) {
	o.Id = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *KalturaAsset) GetObjectType() string {
	if o == nil || o.ObjectType == nil {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetObjectTypeOk() (*string, bool) {
	if o == nil || o.ObjectType == nil {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *KalturaAsset) HasObjectType() bool {
	if o != nil && o.ObjectType != nil {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *KalturaAsset) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetPartnerData returns the PartnerData field value if set, zero value otherwise.
func (o *KalturaAsset) GetPartnerData() string {
	if o == nil || o.PartnerData == nil {
		var ret string
		return ret
	}
	return *o.PartnerData
}

// GetPartnerDataOk returns a tuple with the PartnerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetPartnerDataOk() (*string, bool) {
	if o == nil || o.PartnerData == nil {
		return nil, false
	}
	return o.PartnerData, true
}

// HasPartnerData returns a boolean if a field has been set.
func (o *KalturaAsset) HasPartnerData() bool {
	if o != nil && o.PartnerData != nil {
		return true
	}

	return false
}

// SetPartnerData gets a reference to the given string and assigns it to the PartnerData field.
func (o *KalturaAsset) SetPartnerData(v string) {
	o.PartnerData = &v
}

// GetPartnerDescription returns the PartnerDescription field value if set, zero value otherwise.
func (o *KalturaAsset) GetPartnerDescription() string {
	if o == nil || o.PartnerDescription == nil {
		var ret string
		return ret
	}
	return *o.PartnerDescription
}

// GetPartnerDescriptionOk returns a tuple with the PartnerDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetPartnerDescriptionOk() (*string, bool) {
	if o == nil || o.PartnerDescription == nil {
		return nil, false
	}
	return o.PartnerDescription, true
}

// HasPartnerDescription returns a boolean if a field has been set.
func (o *KalturaAsset) HasPartnerDescription() bool {
	if o != nil && o.PartnerDescription != nil {
		return true
	}

	return false
}

// SetPartnerDescription gets a reference to the given string and assigns it to the PartnerDescription field.
func (o *KalturaAsset) SetPartnerDescription(v string) {
	o.PartnerDescription = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaAsset) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaAsset) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaAsset) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *KalturaAsset) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *KalturaAsset) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *KalturaAsset) SetSize(v int32) {
	o.Size = &v
}

// GetSizeInBytes returns the SizeInBytes field value if set, zero value otherwise.
func (o *KalturaAsset) GetSizeInBytes() int32 {
	if o == nil || o.SizeInBytes == nil {
		var ret int32
		return ret
	}
	return *o.SizeInBytes
}

// GetSizeInBytesOk returns a tuple with the SizeInBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetSizeInBytesOk() (*int32, bool) {
	if o == nil || o.SizeInBytes == nil {
		return nil, false
	}
	return o.SizeInBytes, true
}

// HasSizeInBytes returns a boolean if a field has been set.
func (o *KalturaAsset) HasSizeInBytes() bool {
	if o != nil && o.SizeInBytes != nil {
		return true
	}

	return false
}

// SetSizeInBytes gets a reference to the given int32 and assigns it to the SizeInBytes field.
func (o *KalturaAsset) SetSizeInBytes(v int32) {
	o.SizeInBytes = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *KalturaAsset) GetTags() string {
	if o == nil || o.Tags == nil {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetTagsOk() (*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *KalturaAsset) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *KalturaAsset) SetTags(v string) {
	o.Tags = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaAsset) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaAsset) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaAsset) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KalturaAsset) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaAsset) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KalturaAsset) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *KalturaAsset) SetVersion(v int32) {
	o.Version = &v
}

func (o KalturaAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActualSourceAssetParamsIds != nil {
		toSerialize["actualSourceAssetParamsIds"] = o.ActualSourceAssetParamsIds
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.DeletedAt != nil {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EntryId != nil {
		toSerialize["entryId"] = o.EntryId
	}
	if o.FileExt != nil {
		toSerialize["fileExt"] = o.FileExt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ObjectType != nil {
		toSerialize["objectType"] = o.ObjectType
	}
	if o.PartnerData != nil {
		toSerialize["partnerData"] = o.PartnerData
	}
	if o.PartnerDescription != nil {
		toSerialize["partnerDescription"] = o.PartnerDescription
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.SizeInBytes != nil {
		toSerialize["sizeInBytes"] = o.SizeInBytes
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaAsset struct {
	value *KalturaAsset
	isSet bool
}

func (v NullableKalturaAsset) Get() *KalturaAsset {
	return v.value
}

func (v *NullableKalturaAsset) Set(val *KalturaAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaAsset(val *KalturaAsset) *NullableKalturaAsset {
	return &NullableKalturaAsset{value: val, isSet: true}
}

func (v NullableKalturaAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


