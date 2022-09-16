# KalturaFileSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Dc** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FileContent** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FileDiscSize** | Pointer to **float32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FileObjectType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaFileSyncObjectType&#x60; | [optional] [readonly] 
**FilePath** | Pointer to **string** |  | [optional] 
**FileRoot** | Pointer to **string** |  | [optional] 
**FileSize** | Pointer to **float32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**FileType** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaFileSyncType&#x60; | [optional] [readonly] 
**FileUrl** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IsCurrentDc** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IsDir** | Pointer to **bool** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LinkCount** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**LinkedId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectSubType** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Original** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**OriginalId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ReadyAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**SrcEncKey** | Pointer to **string** |  | [optional] 
**SrcPath** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Enum Type: &#x60;KalturaFileSyncStatus&#x60; | [optional] 
**StorageClass** | Pointer to **string** |  | [optional] 
**SyncTime** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Version** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaFileSync

`func NewKalturaFileSync() *KalturaFileSync`

NewKalturaFileSync instantiates a new KalturaFileSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaFileSyncWithDefaults

`func NewKalturaFileSyncWithDefaults() *KalturaFileSync`

NewKalturaFileSyncWithDefaults instantiates a new KalturaFileSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaFileSync) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaFileSync) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaFileSync) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaFileSync) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDc

`func (o *KalturaFileSync) GetDc() string`

GetDc returns the Dc field if non-nil, zero value otherwise.

### GetDcOk

`func (o *KalturaFileSync) GetDcOk() (*string, bool)`

GetDcOk returns a tuple with the Dc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDc

`func (o *KalturaFileSync) SetDc(v string)`

SetDc sets Dc field to given value.

### HasDc

`func (o *KalturaFileSync) HasDc() bool`

HasDc returns a boolean if a field has been set.

### GetFileContent

`func (o *KalturaFileSync) GetFileContent() string`

GetFileContent returns the FileContent field if non-nil, zero value otherwise.

### GetFileContentOk

`func (o *KalturaFileSync) GetFileContentOk() (*string, bool)`

GetFileContentOk returns a tuple with the FileContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileContent

`func (o *KalturaFileSync) SetFileContent(v string)`

SetFileContent sets FileContent field to given value.

### HasFileContent

`func (o *KalturaFileSync) HasFileContent() bool`

HasFileContent returns a boolean if a field has been set.

### GetFileDiscSize

`func (o *KalturaFileSync) GetFileDiscSize() float32`

GetFileDiscSize returns the FileDiscSize field if non-nil, zero value otherwise.

### GetFileDiscSizeOk

`func (o *KalturaFileSync) GetFileDiscSizeOk() (*float32, bool)`

GetFileDiscSizeOk returns a tuple with the FileDiscSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileDiscSize

`func (o *KalturaFileSync) SetFileDiscSize(v float32)`

SetFileDiscSize sets FileDiscSize field to given value.

### HasFileDiscSize

`func (o *KalturaFileSync) HasFileDiscSize() bool`

HasFileDiscSize returns a boolean if a field has been set.

### GetFileObjectType

`func (o *KalturaFileSync) GetFileObjectType() string`

GetFileObjectType returns the FileObjectType field if non-nil, zero value otherwise.

### GetFileObjectTypeOk

`func (o *KalturaFileSync) GetFileObjectTypeOk() (*string, bool)`

GetFileObjectTypeOk returns a tuple with the FileObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileObjectType

`func (o *KalturaFileSync) SetFileObjectType(v string)`

SetFileObjectType sets FileObjectType field to given value.

### HasFileObjectType

`func (o *KalturaFileSync) HasFileObjectType() bool`

HasFileObjectType returns a boolean if a field has been set.

### GetFilePath

`func (o *KalturaFileSync) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *KalturaFileSync) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *KalturaFileSync) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *KalturaFileSync) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetFileRoot

`func (o *KalturaFileSync) GetFileRoot() string`

GetFileRoot returns the FileRoot field if non-nil, zero value otherwise.

### GetFileRootOk

`func (o *KalturaFileSync) GetFileRootOk() (*string, bool)`

GetFileRootOk returns a tuple with the FileRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileRoot

`func (o *KalturaFileSync) SetFileRoot(v string)`

SetFileRoot sets FileRoot field to given value.

### HasFileRoot

`func (o *KalturaFileSync) HasFileRoot() bool`

HasFileRoot returns a boolean if a field has been set.

### GetFileSize

`func (o *KalturaFileSync) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *KalturaFileSync) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *KalturaFileSync) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *KalturaFileSync) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetFileType

`func (o *KalturaFileSync) GetFileType() int32`

GetFileType returns the FileType field if non-nil, zero value otherwise.

### GetFileTypeOk

`func (o *KalturaFileSync) GetFileTypeOk() (*int32, bool)`

GetFileTypeOk returns a tuple with the FileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileType

`func (o *KalturaFileSync) SetFileType(v int32)`

SetFileType sets FileType field to given value.

### HasFileType

`func (o *KalturaFileSync) HasFileType() bool`

HasFileType returns a boolean if a field has been set.

### GetFileUrl

`func (o *KalturaFileSync) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *KalturaFileSync) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *KalturaFileSync) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *KalturaFileSync) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetId

`func (o *KalturaFileSync) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaFileSync) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaFileSync) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaFileSync) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCurrentDc

`func (o *KalturaFileSync) GetIsCurrentDc() bool`

GetIsCurrentDc returns the IsCurrentDc field if non-nil, zero value otherwise.

### GetIsCurrentDcOk

`func (o *KalturaFileSync) GetIsCurrentDcOk() (*bool, bool)`

GetIsCurrentDcOk returns a tuple with the IsCurrentDc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrentDc

`func (o *KalturaFileSync) SetIsCurrentDc(v bool)`

SetIsCurrentDc sets IsCurrentDc field to given value.

### HasIsCurrentDc

`func (o *KalturaFileSync) HasIsCurrentDc() bool`

HasIsCurrentDc returns a boolean if a field has been set.

### GetIsDir

`func (o *KalturaFileSync) GetIsDir() bool`

GetIsDir returns the IsDir field if non-nil, zero value otherwise.

### GetIsDirOk

`func (o *KalturaFileSync) GetIsDirOk() (*bool, bool)`

GetIsDirOk returns a tuple with the IsDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDir

`func (o *KalturaFileSync) SetIsDir(v bool)`

SetIsDir sets IsDir field to given value.

### HasIsDir

`func (o *KalturaFileSync) HasIsDir() bool`

HasIsDir returns a boolean if a field has been set.

### GetLinkCount

`func (o *KalturaFileSync) GetLinkCount() int32`

GetLinkCount returns the LinkCount field if non-nil, zero value otherwise.

### GetLinkCountOk

`func (o *KalturaFileSync) GetLinkCountOk() (*int32, bool)`

GetLinkCountOk returns a tuple with the LinkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCount

`func (o *KalturaFileSync) SetLinkCount(v int32)`

SetLinkCount sets LinkCount field to given value.

### HasLinkCount

`func (o *KalturaFileSync) HasLinkCount() bool`

HasLinkCount returns a boolean if a field has been set.

### GetLinkedId

`func (o *KalturaFileSync) GetLinkedId() int32`

GetLinkedId returns the LinkedId field if non-nil, zero value otherwise.

### GetLinkedIdOk

`func (o *KalturaFileSync) GetLinkedIdOk() (*int32, bool)`

GetLinkedIdOk returns a tuple with the LinkedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedId

`func (o *KalturaFileSync) SetLinkedId(v int32)`

SetLinkedId sets LinkedId field to given value.

### HasLinkedId

`func (o *KalturaFileSync) HasLinkedId() bool`

HasLinkedId returns a boolean if a field has been set.

### GetObjectId

`func (o *KalturaFileSync) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *KalturaFileSync) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *KalturaFileSync) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *KalturaFileSync) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetObjectSubType

`func (o *KalturaFileSync) GetObjectSubType() int32`

GetObjectSubType returns the ObjectSubType field if non-nil, zero value otherwise.

### GetObjectSubTypeOk

`func (o *KalturaFileSync) GetObjectSubTypeOk() (*int32, bool)`

GetObjectSubTypeOk returns a tuple with the ObjectSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectSubType

`func (o *KalturaFileSync) SetObjectSubType(v int32)`

SetObjectSubType sets ObjectSubType field to given value.

### HasObjectSubType

`func (o *KalturaFileSync) HasObjectSubType() bool`

HasObjectSubType returns a boolean if a field has been set.

### GetOriginal

`func (o *KalturaFileSync) GetOriginal() int32`

GetOriginal returns the Original field if non-nil, zero value otherwise.

### GetOriginalOk

`func (o *KalturaFileSync) GetOriginalOk() (*int32, bool)`

GetOriginalOk returns a tuple with the Original field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginal

`func (o *KalturaFileSync) SetOriginal(v int32)`

SetOriginal sets Original field to given value.

### HasOriginal

`func (o *KalturaFileSync) HasOriginal() bool`

HasOriginal returns a boolean if a field has been set.

### GetOriginalId

`func (o *KalturaFileSync) GetOriginalId() int32`

GetOriginalId returns the OriginalId field if non-nil, zero value otherwise.

### GetOriginalIdOk

`func (o *KalturaFileSync) GetOriginalIdOk() (*int32, bool)`

GetOriginalIdOk returns a tuple with the OriginalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalId

`func (o *KalturaFileSync) SetOriginalId(v int32)`

SetOriginalId sets OriginalId field to given value.

### HasOriginalId

`func (o *KalturaFileSync) HasOriginalId() bool`

HasOriginalId returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaFileSync) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaFileSync) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaFileSync) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaFileSync) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetReadyAt

`func (o *KalturaFileSync) GetReadyAt() int32`

GetReadyAt returns the ReadyAt field if non-nil, zero value otherwise.

### GetReadyAtOk

`func (o *KalturaFileSync) GetReadyAtOk() (*int32, bool)`

GetReadyAtOk returns a tuple with the ReadyAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyAt

`func (o *KalturaFileSync) SetReadyAt(v int32)`

SetReadyAt sets ReadyAt field to given value.

### HasReadyAt

`func (o *KalturaFileSync) HasReadyAt() bool`

HasReadyAt returns a boolean if a field has been set.

### GetSrcEncKey

`func (o *KalturaFileSync) GetSrcEncKey() string`

GetSrcEncKey returns the SrcEncKey field if non-nil, zero value otherwise.

### GetSrcEncKeyOk

`func (o *KalturaFileSync) GetSrcEncKeyOk() (*string, bool)`

GetSrcEncKeyOk returns a tuple with the SrcEncKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcEncKey

`func (o *KalturaFileSync) SetSrcEncKey(v string)`

SetSrcEncKey sets SrcEncKey field to given value.

### HasSrcEncKey

`func (o *KalturaFileSync) HasSrcEncKey() bool`

HasSrcEncKey returns a boolean if a field has been set.

### GetSrcPath

`func (o *KalturaFileSync) GetSrcPath() string`

GetSrcPath returns the SrcPath field if non-nil, zero value otherwise.

### GetSrcPathOk

`func (o *KalturaFileSync) GetSrcPathOk() (*string, bool)`

GetSrcPathOk returns a tuple with the SrcPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPath

`func (o *KalturaFileSync) SetSrcPath(v string)`

SetSrcPath sets SrcPath field to given value.

### HasSrcPath

`func (o *KalturaFileSync) HasSrcPath() bool`

HasSrcPath returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaFileSync) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaFileSync) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaFileSync) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaFileSync) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageClass

`func (o *KalturaFileSync) GetStorageClass() string`

GetStorageClass returns the StorageClass field if non-nil, zero value otherwise.

### GetStorageClassOk

`func (o *KalturaFileSync) GetStorageClassOk() (*string, bool)`

GetStorageClassOk returns a tuple with the StorageClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClass

`func (o *KalturaFileSync) SetStorageClass(v string)`

SetStorageClass sets StorageClass field to given value.

### HasStorageClass

`func (o *KalturaFileSync) HasStorageClass() bool`

HasStorageClass returns a boolean if a field has been set.

### GetSyncTime

`func (o *KalturaFileSync) GetSyncTime() int32`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *KalturaFileSync) GetSyncTimeOk() (*int32, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *KalturaFileSync) SetSyncTime(v int32)`

SetSyncTime sets SyncTime field to given value.

### HasSyncTime

`func (o *KalturaFileSync) HasSyncTime() bool`

HasSyncTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaFileSync) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaFileSync) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaFileSync) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaFileSync) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaFileSync) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaFileSync) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaFileSync) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaFileSync) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


