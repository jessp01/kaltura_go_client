# MetadataBatchGetExclusiveTransformMetadataJobsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**KalturaBatchJobFilter**](KalturaBatchJobFilter.md) |  | [optional] 
**LockKey** | [**KalturaExclusiveLockKey**](KalturaExclusiveLockKey.md) |  | 
**MaxExecutionTime** | **int32** |  | 
**MaxJobToPull** | Pointer to **int32** |  | [optional] 
**NumberOfJobs** | **int32** |  | 

## Methods

### NewMetadataBatchGetExclusiveTransformMetadataJobsRequest

`func NewMetadataBatchGetExclusiveTransformMetadataJobsRequest(lockKey KalturaExclusiveLockKey, maxExecutionTime int32, numberOfJobs int32, ) *MetadataBatchGetExclusiveTransformMetadataJobsRequest`

NewMetadataBatchGetExclusiveTransformMetadataJobsRequest instantiates a new MetadataBatchGetExclusiveTransformMetadataJobsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataBatchGetExclusiveTransformMetadataJobsRequestWithDefaults

`func NewMetadataBatchGetExclusiveTransformMetadataJobsRequestWithDefaults() *MetadataBatchGetExclusiveTransformMetadataJobsRequest`

NewMetadataBatchGetExclusiveTransformMetadataJobsRequestWithDefaults instantiates a new MetadataBatchGetExclusiveTransformMetadataJobsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetFilter() KalturaBatchJobFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetFilterOk() (*KalturaBatchJobFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) SetFilter(v KalturaBatchJobFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetLockKey

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetLockKey() KalturaExclusiveLockKey`

GetLockKey returns the LockKey field if non-nil, zero value otherwise.

### GetLockKeyOk

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetLockKeyOk() (*KalturaExclusiveLockKey, bool)`

GetLockKeyOk returns a tuple with the LockKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockKey

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) SetLockKey(v KalturaExclusiveLockKey)`

SetLockKey sets LockKey field to given value.


### GetMaxExecutionTime

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetMaxExecutionTime() int32`

GetMaxExecutionTime returns the MaxExecutionTime field if non-nil, zero value otherwise.

### GetMaxExecutionTimeOk

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetMaxExecutionTimeOk() (*int32, bool)`

GetMaxExecutionTimeOk returns a tuple with the MaxExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecutionTime

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) SetMaxExecutionTime(v int32)`

SetMaxExecutionTime sets MaxExecutionTime field to given value.


### GetMaxJobToPull

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetMaxJobToPull() int32`

GetMaxJobToPull returns the MaxJobToPull field if non-nil, zero value otherwise.

### GetMaxJobToPullOk

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetMaxJobToPullOk() (*int32, bool)`

GetMaxJobToPullOk returns a tuple with the MaxJobToPull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxJobToPull

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) SetMaxJobToPull(v int32)`

SetMaxJobToPull sets MaxJobToPull field to given value.

### HasMaxJobToPull

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) HasMaxJobToPull() bool`

HasMaxJobToPull returns a boolean if a field has been set.

### GetNumberOfJobs

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetNumberOfJobs() int32`

GetNumberOfJobs returns the NumberOfJobs field if non-nil, zero value otherwise.

### GetNumberOfJobsOk

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) GetNumberOfJobsOk() (*int32, bool)`

GetNumberOfJobsOk returns a tuple with the NumberOfJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfJobs

`func (o *MetadataBatchGetExclusiveTransformMetadataJobsRequest) SetNumberOfJobs(v int32)`

SetNumberOfJobs sets NumberOfJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


