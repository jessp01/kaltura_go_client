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

// KalturaZoomRecordingFile struct for KalturaZoomRecordingFile
type KalturaZoomRecordingFile struct {
	DownloadToken *string `json:"downloadToken,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	FileExtension *string `json:"fileExtension,omitempty"`
	// Enum Type: `KalturaRecordingFileType`
	FileType *int32 `json:"fileType,omitempty"`
	Id *string `json:"id,omitempty"`
	RecordingStart *string `json:"recordingStart,omitempty"`
}

// NewKalturaZoomRecordingFile instantiates a new KalturaZoomRecordingFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKalturaZoomRecordingFile() *KalturaZoomRecordingFile {
	this := KalturaZoomRecordingFile{}
	return &this
}

// NewKalturaZoomRecordingFileWithDefaults instantiates a new KalturaZoomRecordingFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKalturaZoomRecordingFileWithDefaults() *KalturaZoomRecordingFile {
	this := KalturaZoomRecordingFile{}
	return &this
}

// GetDownloadToken returns the DownloadToken field value if set, zero value otherwise.
func (o *KalturaZoomRecordingFile) GetDownloadToken() string {
	if o == nil || o.DownloadToken == nil {
		var ret string
		return ret
	}
	return *o.DownloadToken
}

// GetDownloadTokenOk returns a tuple with the DownloadToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomRecordingFile) GetDownloadTokenOk() (*string, bool) {
	if o == nil || o.DownloadToken == nil {
		return nil, false
	}
	return o.DownloadToken, true
}

// HasDownloadToken returns a boolean if a field has been set.
func (o *KalturaZoomRecordingFile) HasDownloadToken() bool {
	if o != nil && o.DownloadToken != nil {
		return true
	}

	return false
}

// SetDownloadToken gets a reference to the given string and assigns it to the DownloadToken field.
func (o *KalturaZoomRecordingFile) SetDownloadToken(v string) {
	o.DownloadToken = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *KalturaZoomRecordingFile) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl == nil {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomRecordingFile) GetDownloadUrlOk() (*string, bool) {
	if o == nil || o.DownloadUrl == nil {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *KalturaZoomRecordingFile) HasDownloadUrl() bool {
	if o != nil && o.DownloadUrl != nil {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *KalturaZoomRecordingFile) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetFileExtension returns the FileExtension field value if set, zero value otherwise.
func (o *KalturaZoomRecordingFile) GetFileExtension() string {
	if o == nil || o.FileExtension == nil {
		var ret string
		return ret
	}
	return *o.FileExtension
}

// GetFileExtensionOk returns a tuple with the FileExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomRecordingFile) GetFileExtensionOk() (*string, bool) {
	if o == nil || o.FileExtension == nil {
		return nil, false
	}
	return o.FileExtension, true
}

// HasFileExtension returns a boolean if a field has been set.
func (o *KalturaZoomRecordingFile) HasFileExtension() bool {
	if o != nil && o.FileExtension != nil {
		return true
	}

	return false
}

// SetFileExtension gets a reference to the given string and assigns it to the FileExtension field.
func (o *KalturaZoomRecordingFile) SetFileExtension(v string) {
	o.FileExtension = &v
}

// GetFileType returns the FileType field value if set, zero value otherwise.
func (o *KalturaZoomRecordingFile) GetFileType() int32 {
	if o == nil || o.FileType == nil {
		var ret int32
		return ret
	}
	return *o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomRecordingFile) GetFileTypeOk() (*int32, bool) {
	if o == nil || o.FileType == nil {
		return nil, false
	}
	return o.FileType, true
}

// HasFileType returns a boolean if a field has been set.
func (o *KalturaZoomRecordingFile) HasFileType() bool {
	if o != nil && o.FileType != nil {
		return true
	}

	return false
}

// SetFileType gets a reference to the given int32 and assigns it to the FileType field.
func (o *KalturaZoomRecordingFile) SetFileType(v int32) {
	o.FileType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *KalturaZoomRecordingFile) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomRecordingFile) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *KalturaZoomRecordingFile) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *KalturaZoomRecordingFile) SetId(v string) {
	o.Id = &v
}

// GetRecordingStart returns the RecordingStart field value if set, zero value otherwise.
func (o *KalturaZoomRecordingFile) GetRecordingStart() string {
	if o == nil || o.RecordingStart == nil {
		var ret string
		return ret
	}
	return *o.RecordingStart
}

// GetRecordingStartOk returns a tuple with the RecordingStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KalturaZoomRecordingFile) GetRecordingStartOk() (*string, bool) {
	if o == nil || o.RecordingStart == nil {
		return nil, false
	}
	return o.RecordingStart, true
}

// HasRecordingStart returns a boolean if a field has been set.
func (o *KalturaZoomRecordingFile) HasRecordingStart() bool {
	if o != nil && o.RecordingStart != nil {
		return true
	}

	return false
}

// SetRecordingStart gets a reference to the given string and assigns it to the RecordingStart field.
func (o *KalturaZoomRecordingFile) SetRecordingStart(v string) {
	o.RecordingStart = &v
}

func (o KalturaZoomRecordingFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DownloadToken != nil {
		toSerialize["downloadToken"] = o.DownloadToken
	}
	if o.DownloadUrl != nil {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if o.FileExtension != nil {
		toSerialize["fileExtension"] = o.FileExtension
	}
	if o.FileType != nil {
		toSerialize["fileType"] = o.FileType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RecordingStart != nil {
		toSerialize["recordingStart"] = o.RecordingStart
	}
	return json.Marshal(toSerialize)
}

type NullableKalturaZoomRecordingFile struct {
	value *KalturaZoomRecordingFile
	isSet bool
}

func (v NullableKalturaZoomRecordingFile) Get() *KalturaZoomRecordingFile {
	return v.value
}

func (v *NullableKalturaZoomRecordingFile) Set(val *KalturaZoomRecordingFile) {
	v.value = val
	v.isSet = true
}

func (v NullableKalturaZoomRecordingFile) IsSet() bool {
	return v.isSet
}

func (v *NullableKalturaZoomRecordingFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKalturaZoomRecordingFile(val *KalturaZoomRecordingFile) *NullableKalturaZoomRecordingFile {
	return &NullableKalturaZoomRecordingFile{value: val, isSet: true}
}

func (v NullableKalturaZoomRecordingFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKalturaZoomRecordingFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


