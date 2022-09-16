# KalturaBatchJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchJob** | Pointer to [**KalturaBatchJob**](KalturaBatchJob.md) |  | [optional] 
**ChildBatchJobs** | Pointer to [**[]KalturaBatchJob**](KalturaBatchJob.md) |  | [optional] 

## Methods

### NewKalturaBatchJobResponse

`func NewKalturaBatchJobResponse() *KalturaBatchJobResponse`

NewKalturaBatchJobResponse instantiates a new KalturaBatchJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaBatchJobResponseWithDefaults

`func NewKalturaBatchJobResponseWithDefaults() *KalturaBatchJobResponse`

NewKalturaBatchJobResponseWithDefaults instantiates a new KalturaBatchJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchJob

`func (o *KalturaBatchJobResponse) GetBatchJob() KalturaBatchJob`

GetBatchJob returns the BatchJob field if non-nil, zero value otherwise.

### GetBatchJobOk

`func (o *KalturaBatchJobResponse) GetBatchJobOk() (*KalturaBatchJob, bool)`

GetBatchJobOk returns a tuple with the BatchJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchJob

`func (o *KalturaBatchJobResponse) SetBatchJob(v KalturaBatchJob)`

SetBatchJob sets BatchJob field to given value.

### HasBatchJob

`func (o *KalturaBatchJobResponse) HasBatchJob() bool`

HasBatchJob returns a boolean if a field has been set.

### GetChildBatchJobs

`func (o *KalturaBatchJobResponse) GetChildBatchJobs() []KalturaBatchJob`

GetChildBatchJobs returns the ChildBatchJobs field if non-nil, zero value otherwise.

### GetChildBatchJobsOk

`func (o *KalturaBatchJobResponse) GetChildBatchJobsOk() (*[]KalturaBatchJob, bool)`

GetChildBatchJobsOk returns a tuple with the ChildBatchJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildBatchJobs

`func (o *KalturaBatchJobResponse) SetChildBatchJobs(v []KalturaBatchJob)`

SetChildBatchJobs sets ChildBatchJobs field to given value.

### HasChildBatchJobs

`func (o *KalturaBatchJobResponse) HasChildBatchJobs() bool`

HasChildBatchJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


