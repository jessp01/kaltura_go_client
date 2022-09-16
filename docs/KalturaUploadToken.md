# KalturaUploadToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedObjectId** | Pointer to **string** | &#x60;readOnly&#x60;  The id of the object this token is attached to. | [optional] [readonly] 
**AttachedObjectType** | Pointer to **string** | &#x60;readOnly&#x60;  The type of the object this token is attached to. | [optional] [readonly] 
**AutoFinalize** | Pointer to **int32** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaNullableBoolean&#x60;  autoFinalize - Should the upload be finalized once the file size on disk matches the file size reported when adding the upload token. | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**FileName** | Pointer to **string** | &#x60;insertOnly&#x60;  Name of the file for the upload token, can be empty when the upload token is created and will be updated internally after the file is uploaded | [optional] 
**FileSize** | Pointer to **float32** | &#x60;insertOnly&#x60;  File size in bytes, can be empty when the upload token is created and will be updated internally after the file is uploaded | [optional] 
**Id** | Pointer to **string** | &#x60;readOnly&#x60;  Upload token unique ID | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60;  Partner ID of the upload token | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaUploadTokenStatus&#x60;  Status of the upload token | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Last update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UploadUrl** | Pointer to **string** | &#x60;readOnly&#x60;  Upload url - to explicitly determine to which domain to address the uploadToken-&gt;upload call | [optional] [readonly] 
**UploadedFileSize** | Pointer to **float32** | &#x60;readOnly&#x60;  Uploaded file size in bytes, can be used to identify how many bytes were uploaded before resuming | [optional] [readonly] 
**UserId** | Pointer to **string** | &#x60;readOnly&#x60;  User id for the upload token | [optional] [readonly] 

## Methods

### NewKalturaUploadToken

`func NewKalturaUploadToken() *KalturaUploadToken`

NewKalturaUploadToken instantiates a new KalturaUploadToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaUploadTokenWithDefaults

`func NewKalturaUploadTokenWithDefaults() *KalturaUploadToken`

NewKalturaUploadTokenWithDefaults instantiates a new KalturaUploadToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedObjectId

`func (o *KalturaUploadToken) GetAttachedObjectId() string`

GetAttachedObjectId returns the AttachedObjectId field if non-nil, zero value otherwise.

### GetAttachedObjectIdOk

`func (o *KalturaUploadToken) GetAttachedObjectIdOk() (*string, bool)`

GetAttachedObjectIdOk returns a tuple with the AttachedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedObjectId

`func (o *KalturaUploadToken) SetAttachedObjectId(v string)`

SetAttachedObjectId sets AttachedObjectId field to given value.

### HasAttachedObjectId

`func (o *KalturaUploadToken) HasAttachedObjectId() bool`

HasAttachedObjectId returns a boolean if a field has been set.

### GetAttachedObjectType

`func (o *KalturaUploadToken) GetAttachedObjectType() string`

GetAttachedObjectType returns the AttachedObjectType field if non-nil, zero value otherwise.

### GetAttachedObjectTypeOk

`func (o *KalturaUploadToken) GetAttachedObjectTypeOk() (*string, bool)`

GetAttachedObjectTypeOk returns a tuple with the AttachedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedObjectType

`func (o *KalturaUploadToken) SetAttachedObjectType(v string)`

SetAttachedObjectType sets AttachedObjectType field to given value.

### HasAttachedObjectType

`func (o *KalturaUploadToken) HasAttachedObjectType() bool`

HasAttachedObjectType returns a boolean if a field has been set.

### GetAutoFinalize

`func (o *KalturaUploadToken) GetAutoFinalize() int32`

GetAutoFinalize returns the AutoFinalize field if non-nil, zero value otherwise.

### GetAutoFinalizeOk

`func (o *KalturaUploadToken) GetAutoFinalizeOk() (*int32, bool)`

GetAutoFinalizeOk returns a tuple with the AutoFinalize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFinalize

`func (o *KalturaUploadToken) SetAutoFinalize(v int32)`

SetAutoFinalize sets AutoFinalize field to given value.

### HasAutoFinalize

`func (o *KalturaUploadToken) HasAutoFinalize() bool`

HasAutoFinalize returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaUploadToken) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaUploadToken) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaUploadToken) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaUploadToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFileName

`func (o *KalturaUploadToken) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KalturaUploadToken) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KalturaUploadToken) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *KalturaUploadToken) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *KalturaUploadToken) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *KalturaUploadToken) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *KalturaUploadToken) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *KalturaUploadToken) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetId

`func (o *KalturaUploadToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaUploadToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaUploadToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaUploadToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaUploadToken) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaUploadToken) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaUploadToken) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaUploadToken) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaUploadToken) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaUploadToken) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaUploadToken) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaUploadToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaUploadToken) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaUploadToken) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaUploadToken) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaUploadToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUploadUrl

`func (o *KalturaUploadToken) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *KalturaUploadToken) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *KalturaUploadToken) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.

### HasUploadUrl

`func (o *KalturaUploadToken) HasUploadUrl() bool`

HasUploadUrl returns a boolean if a field has been set.

### GetUploadedFileSize

`func (o *KalturaUploadToken) GetUploadedFileSize() float32`

GetUploadedFileSize returns the UploadedFileSize field if non-nil, zero value otherwise.

### GetUploadedFileSizeOk

`func (o *KalturaUploadToken) GetUploadedFileSizeOk() (*float32, bool)`

GetUploadedFileSizeOk returns a tuple with the UploadedFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedFileSize

`func (o *KalturaUploadToken) SetUploadedFileSize(v float32)`

SetUploadedFileSize sets UploadedFileSize field to given value.

### HasUploadedFileSize

`func (o *KalturaUploadToken) HasUploadedFileSize() bool`

HasUploadedFileSize returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaUploadToken) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaUploadToken) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaUploadToken) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaUploadToken) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


