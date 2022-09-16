# StorageProfileLockPendingFileSyncsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**KalturaFileSyncFilter**](KalturaFileSyncFilter.md) |  | 
**MaxCount** | **int32** |  | 
**MaxSize** | Pointer to **int32** |  | [optional] 
**StorageProfileId** | **int32** |  | 
**WorkerId** | **int32** |  | 

## Methods

### NewStorageProfileLockPendingFileSyncsRequest

`func NewStorageProfileLockPendingFileSyncsRequest(filter KalturaFileSyncFilter, maxCount int32, storageProfileId int32, workerId int32, ) *StorageProfileLockPendingFileSyncsRequest`

NewStorageProfileLockPendingFileSyncsRequest instantiates a new StorageProfileLockPendingFileSyncsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProfileLockPendingFileSyncsRequestWithDefaults

`func NewStorageProfileLockPendingFileSyncsRequestWithDefaults() *StorageProfileLockPendingFileSyncsRequest`

NewStorageProfileLockPendingFileSyncsRequestWithDefaults instantiates a new StorageProfileLockPendingFileSyncsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *StorageProfileLockPendingFileSyncsRequest) GetFilter() KalturaFileSyncFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *StorageProfileLockPendingFileSyncsRequest) GetFilterOk() (*KalturaFileSyncFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *StorageProfileLockPendingFileSyncsRequest) SetFilter(v KalturaFileSyncFilter)`

SetFilter sets Filter field to given value.


### GetMaxCount

`func (o *StorageProfileLockPendingFileSyncsRequest) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *StorageProfileLockPendingFileSyncsRequest) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *StorageProfileLockPendingFileSyncsRequest) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.


### GetMaxSize

`func (o *StorageProfileLockPendingFileSyncsRequest) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *StorageProfileLockPendingFileSyncsRequest) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *StorageProfileLockPendingFileSyncsRequest) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *StorageProfileLockPendingFileSyncsRequest) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### GetStorageProfileId

`func (o *StorageProfileLockPendingFileSyncsRequest) GetStorageProfileId() int32`

GetStorageProfileId returns the StorageProfileId field if non-nil, zero value otherwise.

### GetStorageProfileIdOk

`func (o *StorageProfileLockPendingFileSyncsRequest) GetStorageProfileIdOk() (*int32, bool)`

GetStorageProfileIdOk returns a tuple with the StorageProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProfileId

`func (o *StorageProfileLockPendingFileSyncsRequest) SetStorageProfileId(v int32)`

SetStorageProfileId sets StorageProfileId field to given value.


### GetWorkerId

`func (o *StorageProfileLockPendingFileSyncsRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *StorageProfileLockPendingFileSyncsRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *StorageProfileLockPendingFileSyncsRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


