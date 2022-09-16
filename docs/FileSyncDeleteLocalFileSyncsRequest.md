# FileSyncDeleteLocalFileSyncsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**KalturaFileSyncFilter**](KalturaFileSyncFilter.md) |  | 
**LockExpiryTimeout** | **int32** |  | 
**RelativeTimeDeletionLimit** | **int32** |  | 
**RelativeTimeRange** | **int32** |  | 
**WorkerId** | **int32** |  | 

## Methods

### NewFileSyncDeleteLocalFileSyncsRequest

`func NewFileSyncDeleteLocalFileSyncsRequest(filter KalturaFileSyncFilter, lockExpiryTimeout int32, relativeTimeDeletionLimit int32, relativeTimeRange int32, workerId int32, ) *FileSyncDeleteLocalFileSyncsRequest`

NewFileSyncDeleteLocalFileSyncsRequest instantiates a new FileSyncDeleteLocalFileSyncsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSyncDeleteLocalFileSyncsRequestWithDefaults

`func NewFileSyncDeleteLocalFileSyncsRequestWithDefaults() *FileSyncDeleteLocalFileSyncsRequest`

NewFileSyncDeleteLocalFileSyncsRequestWithDefaults instantiates a new FileSyncDeleteLocalFileSyncsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetFilter() KalturaFileSyncFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetFilterOk() (*KalturaFileSyncFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FileSyncDeleteLocalFileSyncsRequest) SetFilter(v KalturaFileSyncFilter)`

SetFilter sets Filter field to given value.


### GetLockExpiryTimeout

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetLockExpiryTimeout() int32`

GetLockExpiryTimeout returns the LockExpiryTimeout field if non-nil, zero value otherwise.

### GetLockExpiryTimeoutOk

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetLockExpiryTimeoutOk() (*int32, bool)`

GetLockExpiryTimeoutOk returns a tuple with the LockExpiryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiryTimeout

`func (o *FileSyncDeleteLocalFileSyncsRequest) SetLockExpiryTimeout(v int32)`

SetLockExpiryTimeout sets LockExpiryTimeout field to given value.


### GetRelativeTimeDeletionLimit

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeDeletionLimit() int32`

GetRelativeTimeDeletionLimit returns the RelativeTimeDeletionLimit field if non-nil, zero value otherwise.

### GetRelativeTimeDeletionLimitOk

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeDeletionLimitOk() (*int32, bool)`

GetRelativeTimeDeletionLimitOk returns a tuple with the RelativeTimeDeletionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTimeDeletionLimit

`func (o *FileSyncDeleteLocalFileSyncsRequest) SetRelativeTimeDeletionLimit(v int32)`

SetRelativeTimeDeletionLimit sets RelativeTimeDeletionLimit field to given value.


### GetRelativeTimeRange

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeRange() int32`

GetRelativeTimeRange returns the RelativeTimeRange field if non-nil, zero value otherwise.

### GetRelativeTimeRangeOk

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetRelativeTimeRangeOk() (*int32, bool)`

GetRelativeTimeRangeOk returns a tuple with the RelativeTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTimeRange

`func (o *FileSyncDeleteLocalFileSyncsRequest) SetRelativeTimeRange(v int32)`

SetRelativeTimeRange sets RelativeTimeRange field to given value.


### GetWorkerId

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *FileSyncDeleteLocalFileSyncsRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *FileSyncDeleteLocalFileSyncsRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


