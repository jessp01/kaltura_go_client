# FileSyncUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileSync** | [**KalturaFileSync**](KalturaFileSync.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewFileSyncUpdateRequest

`func NewFileSyncUpdateRequest(fileSync KalturaFileSync, id int32, ) *FileSyncUpdateRequest`

NewFileSyncUpdateRequest instantiates a new FileSyncUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSyncUpdateRequestWithDefaults

`func NewFileSyncUpdateRequestWithDefaults() *FileSyncUpdateRequest`

NewFileSyncUpdateRequestWithDefaults instantiates a new FileSyncUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileSync

`func (o *FileSyncUpdateRequest) GetFileSync() KalturaFileSync`

GetFileSync returns the FileSync field if non-nil, zero value otherwise.

### GetFileSyncOk

`func (o *FileSyncUpdateRequest) GetFileSyncOk() (*KalturaFileSync, bool)`

GetFileSyncOk returns a tuple with the FileSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSync

`func (o *FileSyncUpdateRequest) SetFileSync(v KalturaFileSync)`

SetFileSync sets FileSync field to given value.


### GetId

`func (o *FileSyncUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileSyncUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileSyncUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


