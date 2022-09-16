# KalturaDropFolderFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchJobId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**DeletedDropFolderFileId** | Pointer to **int32** |  | [optional] 
**DropFolderId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** | Enum Type: &#x60;KalturaDropFolderFileErrorCode&#x60; | [optional] 
**ErrorDescription** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** | &#x60;insertOnly&#x60; | [optional] 
**FileSize** | Pointer to **float32** |  | [optional] 
**FileSizeLastSetAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ImportEndedAt** | Pointer to **int32** |  | [optional] 
**ImportStartedAt** | Pointer to **int32** |  | [optional] 
**LastModificationTime** | Pointer to **string** |  | [optional] 
**LeadDropFolderFileId** | Pointer to **int32** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**ParsedFlavor** | Pointer to **string** |  | [optional] 
**ParsedSlug** | Pointer to **string** |  | [optional] 
**ParsedUserId** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaDropFolderFileStatus&#x60; | [optional] [readonly] 
**Type** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaDropFolderType&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UploadEndDetectedAt** | Pointer to **int32** |  | [optional] 
**UploadStartDetectedAt** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaDropFolderFile

`func NewKalturaDropFolderFile() *KalturaDropFolderFile`

NewKalturaDropFolderFile instantiates a new KalturaDropFolderFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDropFolderFileWithDefaults

`func NewKalturaDropFolderFileWithDefaults() *KalturaDropFolderFile`

NewKalturaDropFolderFileWithDefaults instantiates a new KalturaDropFolderFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchJobId

`func (o *KalturaDropFolderFile) GetBatchJobId() int32`

GetBatchJobId returns the BatchJobId field if non-nil, zero value otherwise.

### GetBatchJobIdOk

`func (o *KalturaDropFolderFile) GetBatchJobIdOk() (*int32, bool)`

GetBatchJobIdOk returns a tuple with the BatchJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchJobId

`func (o *KalturaDropFolderFile) SetBatchJobId(v int32)`

SetBatchJobId sets BatchJobId field to given value.

### HasBatchJobId

`func (o *KalturaDropFolderFile) HasBatchJobId() bool`

HasBatchJobId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaDropFolderFile) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaDropFolderFile) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaDropFolderFile) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaDropFolderFile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeletedDropFolderFileId

`func (o *KalturaDropFolderFile) GetDeletedDropFolderFileId() int32`

GetDeletedDropFolderFileId returns the DeletedDropFolderFileId field if non-nil, zero value otherwise.

### GetDeletedDropFolderFileIdOk

`func (o *KalturaDropFolderFile) GetDeletedDropFolderFileIdOk() (*int32, bool)`

GetDeletedDropFolderFileIdOk returns a tuple with the DeletedDropFolderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDropFolderFileId

`func (o *KalturaDropFolderFile) SetDeletedDropFolderFileId(v int32)`

SetDeletedDropFolderFileId sets DeletedDropFolderFileId field to given value.

### HasDeletedDropFolderFileId

`func (o *KalturaDropFolderFile) HasDeletedDropFolderFileId() bool`

HasDeletedDropFolderFileId returns a boolean if a field has been set.

### GetDropFolderId

`func (o *KalturaDropFolderFile) GetDropFolderId() int32`

GetDropFolderId returns the DropFolderId field if non-nil, zero value otherwise.

### GetDropFolderIdOk

`func (o *KalturaDropFolderFile) GetDropFolderIdOk() (*int32, bool)`

GetDropFolderIdOk returns a tuple with the DropFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropFolderId

`func (o *KalturaDropFolderFile) SetDropFolderId(v int32)`

SetDropFolderId sets DropFolderId field to given value.

### HasDropFolderId

`func (o *KalturaDropFolderFile) HasDropFolderId() bool`

HasDropFolderId returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaDropFolderFile) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaDropFolderFile) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaDropFolderFile) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaDropFolderFile) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetErrorCode

`func (o *KalturaDropFolderFile) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *KalturaDropFolderFile) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *KalturaDropFolderFile) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *KalturaDropFolderFile) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaDropFolderFile) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaDropFolderFile) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaDropFolderFile) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaDropFolderFile) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetFileName

`func (o *KalturaDropFolderFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *KalturaDropFolderFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *KalturaDropFolderFile) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *KalturaDropFolderFile) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetFileSize

`func (o *KalturaDropFolderFile) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *KalturaDropFolderFile) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *KalturaDropFolderFile) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *KalturaDropFolderFile) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileSizeLastSetAt

`func (o *KalturaDropFolderFile) GetFileSizeLastSetAt() int32`

GetFileSizeLastSetAt returns the FileSizeLastSetAt field if non-nil, zero value otherwise.

### GetFileSizeLastSetAtOk

`func (o *KalturaDropFolderFile) GetFileSizeLastSetAtOk() (*int32, bool)`

GetFileSizeLastSetAtOk returns a tuple with the FileSizeLastSetAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSizeLastSetAt

`func (o *KalturaDropFolderFile) SetFileSizeLastSetAt(v int32)`

SetFileSizeLastSetAt sets FileSizeLastSetAt field to given value.

### HasFileSizeLastSetAt

`func (o *KalturaDropFolderFile) HasFileSizeLastSetAt() bool`

HasFileSizeLastSetAt returns a boolean if a field has been set.

### GetId

`func (o *KalturaDropFolderFile) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaDropFolderFile) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaDropFolderFile) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaDropFolderFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImportEndedAt

`func (o *KalturaDropFolderFile) GetImportEndedAt() int32`

GetImportEndedAt returns the ImportEndedAt field if non-nil, zero value otherwise.

### GetImportEndedAtOk

`func (o *KalturaDropFolderFile) GetImportEndedAtOk() (*int32, bool)`

GetImportEndedAtOk returns a tuple with the ImportEndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportEndedAt

`func (o *KalturaDropFolderFile) SetImportEndedAt(v int32)`

SetImportEndedAt sets ImportEndedAt field to given value.

### HasImportEndedAt

`func (o *KalturaDropFolderFile) HasImportEndedAt() bool`

HasImportEndedAt returns a boolean if a field has been set.

### GetImportStartedAt

`func (o *KalturaDropFolderFile) GetImportStartedAt() int32`

GetImportStartedAt returns the ImportStartedAt field if non-nil, zero value otherwise.

### GetImportStartedAtOk

`func (o *KalturaDropFolderFile) GetImportStartedAtOk() (*int32, bool)`

GetImportStartedAtOk returns a tuple with the ImportStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportStartedAt

`func (o *KalturaDropFolderFile) SetImportStartedAt(v int32)`

SetImportStartedAt sets ImportStartedAt field to given value.

### HasImportStartedAt

`func (o *KalturaDropFolderFile) HasImportStartedAt() bool`

HasImportStartedAt returns a boolean if a field has been set.

### GetLastModificationTime

`func (o *KalturaDropFolderFile) GetLastModificationTime() string`

GetLastModificationTime returns the LastModificationTime field if non-nil, zero value otherwise.

### GetLastModificationTimeOk

`func (o *KalturaDropFolderFile) GetLastModificationTimeOk() (*string, bool)`

GetLastModificationTimeOk returns a tuple with the LastModificationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModificationTime

`func (o *KalturaDropFolderFile) SetLastModificationTime(v string)`

SetLastModificationTime sets LastModificationTime field to given value.

### HasLastModificationTime

`func (o *KalturaDropFolderFile) HasLastModificationTime() bool`

HasLastModificationTime returns a boolean if a field has been set.

### GetLeadDropFolderFileId

`func (o *KalturaDropFolderFile) GetLeadDropFolderFileId() int32`

GetLeadDropFolderFileId returns the LeadDropFolderFileId field if non-nil, zero value otherwise.

### GetLeadDropFolderFileIdOk

`func (o *KalturaDropFolderFile) GetLeadDropFolderFileIdOk() (*int32, bool)`

GetLeadDropFolderFileIdOk returns a tuple with the LeadDropFolderFileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadDropFolderFileId

`func (o *KalturaDropFolderFile) SetLeadDropFolderFileId(v int32)`

SetLeadDropFolderFileId sets LeadDropFolderFileId field to given value.

### HasLeadDropFolderFileId

`func (o *KalturaDropFolderFile) HasLeadDropFolderFileId() bool`

HasLeadDropFolderFileId returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaDropFolderFile) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaDropFolderFile) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaDropFolderFile) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaDropFolderFile) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetParsedFlavor

`func (o *KalturaDropFolderFile) GetParsedFlavor() string`

GetParsedFlavor returns the ParsedFlavor field if non-nil, zero value otherwise.

### GetParsedFlavorOk

`func (o *KalturaDropFolderFile) GetParsedFlavorOk() (*string, bool)`

GetParsedFlavorOk returns a tuple with the ParsedFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedFlavor

`func (o *KalturaDropFolderFile) SetParsedFlavor(v string)`

SetParsedFlavor sets ParsedFlavor field to given value.

### HasParsedFlavor

`func (o *KalturaDropFolderFile) HasParsedFlavor() bool`

HasParsedFlavor returns a boolean if a field has been set.

### GetParsedSlug

`func (o *KalturaDropFolderFile) GetParsedSlug() string`

GetParsedSlug returns the ParsedSlug field if non-nil, zero value otherwise.

### GetParsedSlugOk

`func (o *KalturaDropFolderFile) GetParsedSlugOk() (*string, bool)`

GetParsedSlugOk returns a tuple with the ParsedSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedSlug

`func (o *KalturaDropFolderFile) SetParsedSlug(v string)`

SetParsedSlug sets ParsedSlug field to given value.

### HasParsedSlug

`func (o *KalturaDropFolderFile) HasParsedSlug() bool`

HasParsedSlug returns a boolean if a field has been set.

### GetParsedUserId

`func (o *KalturaDropFolderFile) GetParsedUserId() string`

GetParsedUserId returns the ParsedUserId field if non-nil, zero value otherwise.

### GetParsedUserIdOk

`func (o *KalturaDropFolderFile) GetParsedUserIdOk() (*string, bool)`

GetParsedUserIdOk returns a tuple with the ParsedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedUserId

`func (o *KalturaDropFolderFile) SetParsedUserId(v string)`

SetParsedUserId sets ParsedUserId field to given value.

### HasParsedUserId

`func (o *KalturaDropFolderFile) HasParsedUserId() bool`

HasParsedUserId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaDropFolderFile) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaDropFolderFile) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaDropFolderFile) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaDropFolderFile) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaDropFolderFile) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaDropFolderFile) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaDropFolderFile) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaDropFolderFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *KalturaDropFolderFile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaDropFolderFile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaDropFolderFile) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaDropFolderFile) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaDropFolderFile) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaDropFolderFile) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaDropFolderFile) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaDropFolderFile) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUploadEndDetectedAt

`func (o *KalturaDropFolderFile) GetUploadEndDetectedAt() int32`

GetUploadEndDetectedAt returns the UploadEndDetectedAt field if non-nil, zero value otherwise.

### GetUploadEndDetectedAtOk

`func (o *KalturaDropFolderFile) GetUploadEndDetectedAtOk() (*int32, bool)`

GetUploadEndDetectedAtOk returns a tuple with the UploadEndDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadEndDetectedAt

`func (o *KalturaDropFolderFile) SetUploadEndDetectedAt(v int32)`

SetUploadEndDetectedAt sets UploadEndDetectedAt field to given value.

### HasUploadEndDetectedAt

`func (o *KalturaDropFolderFile) HasUploadEndDetectedAt() bool`

HasUploadEndDetectedAt returns a boolean if a field has been set.

### GetUploadStartDetectedAt

`func (o *KalturaDropFolderFile) GetUploadStartDetectedAt() int32`

GetUploadStartDetectedAt returns the UploadStartDetectedAt field if non-nil, zero value otherwise.

### GetUploadStartDetectedAtOk

`func (o *KalturaDropFolderFile) GetUploadStartDetectedAtOk() (*int32, bool)`

GetUploadStartDetectedAtOk returns a tuple with the UploadStartDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadStartDetectedAt

`func (o *KalturaDropFolderFile) SetUploadStartDetectedAt(v int32)`

SetUploadStartDetectedAt sets UploadStartDetectedAt field to given value.

### HasUploadStartDetectedAt

`func (o *KalturaDropFolderFile) HasUploadStartDetectedAt() bool`

HasUploadStartDetectedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


