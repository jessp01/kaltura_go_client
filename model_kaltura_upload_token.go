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

// KalturaUploadToken struct for KalturaUploadToken
type KalturaUploadToken struct {
	// `readOnly`  The id of the object this token is attached to.
	AttachedObjectId *string `json:"attachedObjectId,omitempty"`
	// `readOnly`  The type of the object this token is attached to.
	AttachedObjectType *string `json:"attachedObjectType,omitempty"`
	// `insertOnly`  Enum Type: `KalturaNullableBoolean`  autoFinalize - Should the upload be finalized once the file size on disk matches the file size reported when adding the upload token.
	AutoFinalize *int32 `json:"autoFinalize,omitempty"`
	// `readOnly`  Creation date as Unix timestamp (In seconds)
	CreatedAt *int32 `json:"createdAt,omitempty"`
	// `insertOnly`  Name of the file for the upload token, can be empty when the upload token is created and will be updated internally after the file is uploaded
	FileName *string `json:"fileName,omitempty"`
	// `insertOnly`  File size in bytes, can be empty when the upload token is created and will be updated internally after the file is uploaded
	FileSize *float32 `json:"fileSize,omitempty"`
	// `readOnly`  Upload token unique ID
	Id *string `json:"id,omitempty"`
	// `readOnly`  Partner ID of the upload token
	PartnerId *int32 `json:"partnerId,omitempty"`
	// `readOnly`  Enum Type: `KalturaUploadTokenStatus`  Status of the upload token
	Status *int32 `json:"status,omitempty"`
	// `readOnly`  Last update date as Unix timestamp (In seconds)
	UpdatedAt *int32 `json:"updatedAt,omitempty"`
	// `readOnly`  Upload url - to explicitly determine to which domain to address the uploadToken->upload call
	UploadUrl *string `json:"uploadUrl,omitempty"`
	// `readOnly`  Uploaded file size in bytes, can be used to identify how many bytes were uploaded before resuming
	UploadedFileSize *float32 `json:"uploadedFileSize,omitempty"`
	// `readOnly`  User id for the upload token
	UserId *string `json:"userId,omitempty"`
}

// NewKalturaUploadToken instantiates a new KalturaUploadToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaUploadToken() *KalturaUploadToken {
	this := KalturaUploadToken{}
	return &this
}

// NewKalturaUploadTokenWithDefaults instantiates a new KalturaUploadToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaUploadTokenWithDefaults() *KalturaUploadToken {
	this := KalturaUploadToken{}
	return &this
}

// GetAttachedObjectId returns the AttachedObjectId field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetAttachedObjectId() string {
	if o == nil || o.AttachedObjectId == nil {
		var ret string
		return ret
	}
	return *o.AttachedObjectId
}

// GetAttachedObjectIdOk returns a tuple with the AttachedObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetAttachedObjectIdOk() (*string, bool) {
	if o == nil || o.AttachedObjectId == nil {
		return nil, false
	}
	return o.AttachedObjectId, true
}

// HasAttachedObjectId returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasAttachedObjectId() bool {
	if o != nil && o.AttachedObjectId != nil {
		return true
	}

	return false
}

// SetAttachedObjectId gets a reference to the given string and assigns it to the AttachedObjectId field.
func (o *KalturaUploadToken) SetAttachedObjectId(v string) {
	o.AttachedObjectId = &v
}

// GetAttachedObjectType returns the AttachedObjectType field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetAttachedObjectType() string {
	if o == nil || o.AttachedObjectType == nil {
		var ret string
		return ret
	}
	return *o.AttachedObjectType
}

// GetAttachedObjectTypeOk returns a tuple with the AttachedObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetAttachedObjectTypeOk() (*string, bool) {
	if o == nil || o.AttachedObjectType == nil {
		return nil, false
	}
	return o.AttachedObjectType, true
}

// HasAttachedObjectType returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasAttachedObjectType() bool {
	if o != nil && o.AttachedObjectType != nil {
		return true
	}

	return false
}

// SetAttachedObjectType gets a reference to the given string and assigns it to the AttachedObjectType field.
func (o *KalturaUploadToken) SetAttachedObjectType(v string) {
	o.AttachedObjectType = &v
}

// GetAutoFinalize returns the AutoFinalize field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetAutoFinalize() int32 {
	if o == nil || o.AutoFinalize == nil {
		var ret int32
		return ret
	}
	return *o.AutoFinalize
}

// GetAutoFinalizeOk returns a tuple with the AutoFinalize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetAutoFinalizeOk() (*int32, bool) {
	if o == nil || o.AutoFinalize == nil {
		return nil, false
	}
	return o.AutoFinalize, true
}

// HasAutoFinalize returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasAutoFinalize() bool {
	if o != nil && o.AutoFinalize != nil {
		return true
	}

	return false
}

// SetAutoFinalize gets a reference to the given int32 and assigns it to the AutoFinalize field.
func (o *KalturaUploadToken) SetAutoFinalize(v int32) {
	o.AutoFinalize = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetCreatedAt() int32 {
	if o == nil || o.CreatedAt == nil {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetCreatedAtOk() (*int32, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *KalturaUploadToken) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasFileName() bool {
	if o != nil && o.FileName != nil {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *KalturaUploadToken) SetFileName(v string) {
	o.FileName = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetFileSize() float32 {
	if o == nil || o.FileSize == nil {
		var ret float32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetFileSizeOk() (*float32, bool) {
	if o == nil || o.FileSize == nil {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasFileSize() bool {
	if o != nil && o.FileSize != nil {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given float32 and assigns it to the FileSize field.
func (o *KalturaUploadToken) SetFileSize(v float32) {
	o.FileSize = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaUploadToken) SetId(v string) {
	o.Id = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetPartnerId() int32 {
	if o == nil || o.PartnerId == nil {
		var ret int32
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetPartnerIdOk() (*int32, bool) {
	if o == nil || o.PartnerId == nil {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasPartnerId() bool {
	if o != nil && o.PartnerId != nil {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given int32 and assigns it to the PartnerId field.
func (o *KalturaUploadToken) SetPartnerId(v int32) {
	o.PartnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *KalturaUploadToken) SetStatus(v int32) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetUpdatedAt() int32 {
	if o == nil || o.UpdatedAt == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *KalturaUploadToken) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUploadUrl returns the UploadUrl field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetUploadUrl() string {
	if o == nil || o.UploadUrl == nil {
		var ret string
		return ret
	}
	return *o.UploadUrl
}

// GetUploadUrlOk returns a tuple with the UploadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetUploadUrlOk() (*string, bool) {
	if o == nil || o.UploadUrl == nil {
		return nil, false
	}
	return o.UploadUrl, true
}

// HasUploadUrl returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasUploadUrl() bool {
	if o != nil && o.UploadUrl != nil {
		return true
	}

	return false
}

// SetUploadUrl gets a reference to the given string and assigns it to the UploadUrl field.
func (o *KalturaUploadToken) SetUploadUrl(v string) {
	o.UploadUrl = &v
}

// GetUploadedFileSize returns the UploadedFileSize field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetUploadedFileSize() float32 {
	if o == nil || o.UploadedFileSize == nil {
		var ret float32
		return ret
	}
	return *o.UploadedFileSize
}

// GetUploadedFileSizeOk returns a tuple with the UploadedFileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetUploadedFileSizeOk() (*float32, bool) {
	if o == nil || o.UploadedFileSize == nil {
		return nil, false
	}
	return o.UploadedFileSize, true
}

// HasUploadedFileSize returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasUploadedFileSize() bool {
	if o != nil && o.UploadedFileSize != nil {
		return true
	}

	return false
}

// SetUploadedFileSize gets a reference to the given float32 and assigns it to the UploadedFileSize field.
func (o *KalturaUploadToken) SetUploadedFileSize(v float32) {
	o.UploadedFileSize = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *KalturaUploadToken) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaUploadToken) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *KalturaUploadToken) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *KalturaUploadToken) SetUserId(v string) {
	o.UserId = &v
}

func (o KalturaUploadToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttachedObjectId != nil {
		toSerialize["attachedObjectId"] = o.AttachedObjectId
	}
	if o.AttachedObjectType != nil {
		toSerialize["attachedObjectType"] = o.AttachedObjectType
	}
	if o.AutoFinalize != nil {
		toSerialize["autoFinalize"] = o.AutoFinalize
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.FileName != nil {
		toSerialize["fileName"] = o.FileName
	}
	if o.FileSize != nil {
		toSerialize["fileSize"] = o.FileSize
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PartnerId != nil {
		toSerialize["partnerId"] = o.PartnerId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.UploadUrl != nil {
		toSerialize["uploadUrl"] = o.UploadUrl
	}
	if o.UploadedFileSize != nil {
		toSerialize["uploadedFileSize"] = o.UploadedFileSize
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaUploadToken struct {
	value *KalturaUploadToken
	isSet bool
}

func (v NullableKalturaUploadToken) Get() *KalturaUploadToken {
	return v.value
}

func (v *NullableKalturaUploadToken) Set(val *KalturaUploadToken) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaUploadToken) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaUploadToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaUploadToken(val *KalturaUploadToken) *NullableKalturaUploadToken {
	return &NullableKalturaUploadToken{value: val, isSet: true}
}

func (v NullableKalturaUploadToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaUploadToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


