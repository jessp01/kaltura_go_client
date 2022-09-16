# FileSyncLockFileSyncsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**KalturaFileSyncFilter**](KalturaFileSyncFilter.md) |  | 
**LockExpiryTimeout** | **int32** |  | 
**MaxCount** | **int32** |  | 
**RelativeTimeLimit** | **int32** |  | 
**RelativeTimeRange** | **int32** |  | 
**WorkerId** | **int32** |  | 

## Methods

### NewFileSyncLockFileSyncsRequest

`func NewFileSyncLockFileSyncsRequest(filter KalturaFileSyncFilter, lockExpiryTimeout int32, maxCount int32, relativeTimeLimit int32, relativeTimeRange int32, workerId int32, ) *FileSyncLockFileSyncsRequest`

NewFileSyncLockFileSyncsRequest instantiates a new FileSyncLockFileSyncsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileSyncLockFileSyncsRequestWithDefaults

`func NewFileSyncLockFileSyncsRequestWithDefaults() *FileSyncLockFileSyncsRequest`

NewFileSyncLockFileSyncsRequestWithDefaults instantiates a new FileSyncLockFileSyncsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *FileSyncLockFileSyncsRequest) GetFilter() KalturaFileSyncFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *FileSyncLockFileSyncsRequest) GetFilterOk() (*KalturaFileSyncFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *FileSyncLockFileSyncsRequest) SetFilter(v KalturaFileSyncFilter)`

SetFilter sets Filter field to given value.


### GetLockExpiryTimeout

`func (o *FileSyncLockFileSyncsRequest) GetLockExpiryTimeout() int32`

GetLockExpiryTimeout returns the LockExpiryTimeout field if non-nil, zero value otherwise.

### GetLockExpiryTimeoutOk

`func (o *FileSyncLockFileSyncsRequest) GetLockExpiryTimeoutOk() (*int32, bool)`

GetLockExpiryTimeoutOk returns a tuple with the LockExpiryTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockExpiryTimeout

`func (o *FileSyncLockFileSyncsRequest) SetLockExpiryTimeout(v int32)`

SetLockExpiryTimeout sets LockExpiryTimeout field to given value.


### GetMaxCount

`func (o *FileSyncLockFileSyncsRequest) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *FileSyncLockFileSyncsRequest) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *FileSyncLockFileSyncsRequest) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.


### GetRelativeTimeLimit

`func (o *FileSyncLockFileSyncsRequest) GetRelativeTimeLimit() int32`

GetRelativeTimeLimit returns the RelativeTimeLimit field if non-nil, zero value otherwise.

### GetRelativeTimeLimitOk

`func (o *FileSyncLockFileSyncsRequest) GetRelativeTimeLimitOk() (*int32, bool)`

GetRelativeTimeLimitOk returns a tuple with the RelativeTimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTimeLimit

`func (o *FileSyncLockFileSyncsRequest) SetRelativeTimeLimit(v int32)`

SetRelativeTimeLimit sets RelativeTimeLimit field to given value.


### GetRelativeTimeRange

`func (o *FileSyncLockFileSyncsRequest) GetRelativeTimeRange() int32`

GetRelativeTimeRange returns the RelativeTimeRange field if non-nil, zero value otherwise.

### GetRelativeTimeRangeOk

`func (o *FileSyncLockFileSyncsRequest) GetRelativeTimeRangeOk() (*int32, bool)`

GetRelativeTimeRangeOk returns a tuple with the RelativeTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTimeRange

`func (o *FileSyncLockFileSyncsRequest) SetRelativeTimeRange(v int32)`

SetRelativeTimeRange sets RelativeTimeRange field to given value.


### GetWorkerId

`func (o *FileSyncLockFileSyncsRequest) GetWorkerId() int32`

GetWorkerId returns the WorkerId field if non-nil, zero value otherwise.

### GetWorkerIdOk

`func (o *FileSyncLockFileSyncsRequest) GetWorkerIdOk() (*int32, bool)`

GetWorkerIdOk returns a tuple with the WorkerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerId

`func (o *FileSyncLockFileSyncsRequest) SetWorkerId(v int32)`

SetWorkerId sets WorkerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


