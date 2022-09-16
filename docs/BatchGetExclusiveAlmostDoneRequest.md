# BatchGetExclusiveAlmostDoneRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaBatchJobFilter**](KalturaBatchJobFilter.md) |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 
**MaxExecutionTime** | **int32** |  | 
**NumberOfJobs** | **int32** |  | 

## Methods

### NewBatchGetExclusiveAlmostDoneRequest

`func NewBatchGetExclusiveAlmostDoneRequest(lockKey KalturaExclusiveLockKey, maxExecutionTime int32, numberOfJobs int32, ) *BatchGetExclusiveAlmostDoneRequest`

NewBatchGetExclusiveAlmostDoneRequest instantiates a new BatchGetExclusiveAlmostDoneRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGetExclusiveAlmostDoneRequestWithDefaults

`func NewBatchGetExclusiveAlmostDoneRequestWithDefaults() *BatchGetExclusiveAlmostDoneRequest`

NewBatchGetExclusiveAlmostDoneRequestWithDefaults instantiates a new BatchGetExclusiveAlmostDoneRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *BatchGetExclusiveAlmostDoneRequest) GetFilter() KalturaBatchJobFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *BatchGetExclusiveAlmostDoneRequest) GetFilterOk() (*KalturaBatchJobFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *BatchGetExclusiveAlmostDoneRequest) SetFilter(v KalturaBatchJobFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *BatchGetExclusiveAlmostDoneRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetJobType

`func (o *BatchGetExclusiveAlmostDoneRequest) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *BatchGetExclusiveAlmostDoneRequest) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *BatchGetExclusiveAlmostDoneRequest) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *BatchGetExclusiveAlmostDoneRequest) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetLockKey

`func (o *BatchGetExclusiveAlmostDoneRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *BatchGetExclusiveAlmostDoneRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *BatchGetExclusiveAlmostDoneRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.


### GetMaxExecutionTime

`func (o *BatchGetExclusiveAlmostDoneRequest) GetMaxExecutionTime() int32`

GetMaxExecutionTime returns the MaxExecutionTime field if non-nil, zero value otherwise.

### GetMaxExecutionTimeOk

`func (o *BatchGetExclusiveAlmostDoneRequest) GetMaxExecutionTimeOk() (*int32, bool)`

GetMaxExecutionTimeOk returns a tuple with the MaxExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutionTime

`func (o *BatchGetExclusiveAlmostDoneRequest) SetMaxExecutionTime(v int32)`

SetMaxExecutionTime sets MaxExecutionTime field to given value.


### GetNumberOfJobs

`func (o *BatchGetExclusiveAlmostDoneRequest) GetNumberOfJobs() int32`

GetNumberOfJobs returns the NumberOfJobs field if non-nil, zero value otherwise.

### GetNumberOfJobsOk

`func (o *BatchGetExclusiveAlmostDoneRequest) GetNumberOfJobsOk() (*int32, bool)`

GetNumberOfJobsOk returns a tuple with the NumberOfJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfJobs

`func (o *BatchGetExclusiveAlmostDoneRequest) SetNumberOfJobs(v int32)`

SetNumberOfJobs sets NumberOfJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


