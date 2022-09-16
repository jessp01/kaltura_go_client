# BatchGetExclusiveJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaBatchJobFilter**](KalturaBatchJobFilter.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 
**MaxExecutionTime** | **int32** |  | 
**MaxJobToPullForCache** | Pointer to **int32** |  | [optional] 
**NumberOfJobs** | **int32** |  | 

## Methods

### NewBatchGetExclusiveJobsRequest

`func NewBatchGetExclusiveJobsRequest(lockKey KalturaExclusiveLockKey, maxExecutionTime int32, numberOfJobs int32, ) *BatchGetExclusiveJobsRequest`

NewBatchGetExclusiveJobsRequest instantiates a new BatchGetExclusiveJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGetExclusiveJobsRequestWithDefaults

`func NewBatchGetExclusiveJobsRequestWithDefaults() *BatchGetExclusiveJobsRequest`

NewBatchGetExclusiveJobsRequestWithDefaults instantiates a new BatchGetExclusiveJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *BatchGetExclusiveJobsRequest) GetFilter() KalturaBatchJobFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BatchGetExclusiveJobsRequest) GetFilterOk() (*KalturaBatchJobFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BatchGetExclusiveJobsRequest) SetFilter(v KalturaBatchJobFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BatchGetExclusiveJobsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetJobType

`func (o *BatchGetExclusiveJobsRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BatchGetExclusiveJobsRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BatchGetExclusiveJobsRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BatchGetExclusiveJobsRequest) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLockKey

`func (o *BatchGetExclusiveJobsRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *BatchGetExclusiveJobsRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *BatchGetExclusiveJobsRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.


### GetMaxExecutionTime

`func (o *BatchGetExclusiveJobsRequest) GetMaxExecutionTime() int32`

GetMaxExecutionTime returns the MaxExecutionTime field if non-nil, zero value otherwise.

### GetMaxExecutionTimeOk

`func (o *BatchGetExclusiveJobsRequest) GetMaxExecutionTimeOk() (*int32, bool)`

GetMaxExecutionTimeOk returns a tuple with the MaxExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutionTime

`func (o *BatchGetExclusiveJobsRequest) SetMaxExecutionTime(v int32)`

SetMaxExecutionTime sets MaxExecutionTime field to given value.


### GetMaxJobToPullForCache

`func (o *BatchGetExclusiveJobsRequest) GetMaxJobToPullForCache() int32`

GetMaxJobToPullForCache returns the MaxJobToPullForCache field if non-nil, zero value otherwise.

### GetMaxJobToPullForCacheOk

`func (o *BatchGetExclusiveJobsRequest) GetMaxJobToPullForCacheOk() (*int32, bool)`

GetMaxJobToPullForCacheOk returns a tuple with the MaxJobToPullForCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobToPullForCache

`func (o *BatchGetExclusiveJobsRequest) SetMaxJobToPullForCache(v int32)`

SetMaxJobToPullForCache sets MaxJobToPullForCache field to given value.

### HasMaxJobToPullForCache

`func (o *BatchGetExclusiveJobsRequest) HasMaxJobToPullForCache() bool`

HasMaxJobToPullForCache returns a boolean if a field has been set.

### GetNumberOfJobs

`func (o *BatchGetExclusiveJobsRequest) GetNumberOfJobs() int32`

GetNumberOfJobs returns the NumberOfJobs field if non-nil, zero value otherwise.

### GetNumberOfJobsOk

`func (o *BatchGetExclusiveJobsRequest) GetNumberOfJobsOk() (*int32, bool)`

GetNumberOfJobsOk returns a tuple with the NumberOfJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfJobs

`func (o *BatchGetExclusiveJobsRequest) SetNumberOfJobs(v int32)`

SetNumberOfJobs sets NumberOfJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


