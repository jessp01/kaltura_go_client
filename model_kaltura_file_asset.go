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

// KalturaFileAsset struct for KalturaFileAsset
type KalturaFileAsset struct {
	// `readOnly`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `insertOnly`  Enum Type: `KalturaFileAssetObjectType`
	FileAssetObjectType *string `json:"fileAssetObjectType,omitempty"`
	FileExt *string `json:"fileExt,omitempty"`
	// `readOnly`
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// `insertOnly`
	ObjectId *string `json:"objectId,omitempty"`
	// `readOnly`
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  Enum Type: `KalturaFileAssetStatus`
	Status *string `json:"status,omitempty"`
	SystemName *string `json:"systemName,omitempty"`
	// `readOnly`
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// `readOnly`
	Version *int32 `json:"version,omitempty"`
}

// NewKalturaFileAsset instantiates a new KalturaFileAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaFileAsset() *KalturaFileAsset {
	this := KalturaFileAsset{}
	return &this
}

// NewKalturaFileAssetWithDefaults instantiates a new KalturaFileAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaFileAssetWithDefaults() *KalturaFileAsset {
	this := KalturaFileAsset{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaFileAsset) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetFileAssetObjectType returns the FileAssetObjectType field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetFileAssetObjectType() string {
	if o == nil || o.FileAssetObjectType == nil {
		var ret string
		return ret
	}
	return *o.FileAssetObjectType
}

// GetFileAssetObjectTypeOk returns a tuple with the FileAssetObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetFileAssetObjectTypeOk() (*string, bool) {
	if o == nil || o.FileAssetObjectType == nil {
		return nil, false
	}
	return o.FileAssetObjectType, true
}

// HasFileAssetObjectType returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasFileAssetObjectType() bool {
	if o != nil && o.FileAssetObjectType != nil {
		return true
	}

	return false
}

// SetFileAssetObjectType gets a reference to the given string and assigns it to the FileAssetObjectType field.
func (o *KalturaFileAsset) SetFileAssetObjectType(v string) {
	o.FileAssetObjectType = &v
}

// GetFileExt returns the FileExt field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetFileExt() string {
	if o == nil || o.FileExt == nil {
		var ret string
		return ret
	}
	return *o.FileExt
}

// GetFileExtOk returns a tuple with the FileExt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetFileExtOk() (*string, bool) {
	if o == nil || o.FileExt == nil {
		return nil, false
	}
	return o.FileExt, true
}

// HasFileExt returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasFileExt() bool {
	if o != nil && o.FileExt != nil {
		return true
	}

	return false
}

// SetFileExt gets a reference to the given string and assigns it to the FileExt field.
func (o *KalturaFileAsset) SetFileExt(v string) {
	o.FileExt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *KalturaFileAsset) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KalturaFileAsset) SetName(v string) {
	o.Name = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *KalturaFileAsset) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaFileAsset) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *KalturaFileAsset) SetStatus(v string) {
	o.Status = &v
}

// GetSystemName returns the SystemName field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetSystemName() string {
	if o == nil || o.SystemName == nil {
		var ret string
		return ret
	}
	return *o.SystemName
}

// GetSystemNameOk returns a tuple with the SystemName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetSystemNameOk() (*string, bool) {
	if o == nil || o.SystemName == nil {
		return nil, false
	}
	return o.SystemName, true
}

// HasSystemName returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasSystemName() bool {
	if o != nil && o.SystemName != nil {
		return true
	}

	return false
}

// SetSystemName gets a reference to the given string and assigns it to the SystemName field.
func (o *KalturaFileAsset) SetSystemName(v string) {
	o.SystemName = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaFileAsset) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *KalturaFileAsset) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaFileAsset) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *KalturaFileAsset) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *KalturaFileAsset) SetVersion(v int32) {
	o.Version = &v
}

func (o KalturaFileAsset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.FileAssetObjectType != nil {
		toSerialize["fileAssetObjectType"] = o.FileAssetObjectType
	}
	if o.FileExt != nil {
		toSerialize["fileExt"] = o.FileExt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SystemName != nil {
		toSerialize["systemName"] = o.SystemName
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaFileAsset struct {
	value *KalturaFileAsset
	isSet bool
}

func (v NullableKalturaFileAsset) Get() *KalturaFileAsset {
	return v.value
}

func (v *NullableKalturaFileAsset) Set(val *KalturaFileAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaFileAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaFileAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaFileAsset(val *KalturaFileAsset) *NullableKalturaFileAsset {
	return &NullableKalturaFileAsset{value: val, isSet: true}
}

func (v NullableKalturaFileAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaFileAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


