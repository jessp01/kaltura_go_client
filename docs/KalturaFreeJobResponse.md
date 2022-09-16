# KalturaFreeJobResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Job** | Pointer to [**KalturaBatchJob**](KalturaBatchJob.md) |  | [optional] 
**JobType** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaBatchJobType&#x60; | [optional] [readonly] 
**QueueSize** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaFreeJobResponse

`func NewKalturaFreeJobResponse() *KalturaFreeJobResponse`

NewKalturaFreeJobResponse instantiates a new KalturaFreeJobResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaFreeJobResponseWithDefaults

`func NewKalturaFreeJobResponseWithDefaults() *KalturaFreeJobResponse`

NewKalturaFreeJobResponseWithDefaults instantiates a new KalturaFreeJobResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJob

`func (o *KalturaFreeJobResponse) GetJob() KalturaBatchJob`

GetJob returns the Job field if non-nil, zero value otherwise.

### GetJobOk

`func (o *KalturaFreeJobResponse) GetJobOk() (*KalturaBatchJob, bool)`

GetJobOk returns a tuple with the Job field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJob

`func (o *KalturaFreeJobResponse) SetJob(v KalturaBatchJob)`

SetJob sets Job field to given value.

### HasJob

`func (o *KalturaFreeJobResponse) HasJob() bool`

HasJob returns a boolean if a field has been set.

### GetJobType

`func (o *KalturaFreeJobResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KalturaFreeJobResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KalturaFreeJobResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KalturaFreeJobResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetQueueSize

`func (o *KalturaFreeJobResponse) GetQueueSize() int32`

GetQueueSize returns the QueueSize field if non-nil, zero value otherwise.

### GetQueueSizeOk

`func (o *KalturaFreeJobResponse) GetQueueSizeOk() (*int32, bool)`

GetQueueSizeOk returns a tuple with the QueueSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueSize

`func (o *KalturaFreeJobResponse) SetQueueSize(v int32)`

SetQueueSize sets QueueSize field to given value.

### HasQueueSize

`func (o *KalturaFreeJobResponse) HasQueueSize() bool`

HasQueueSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


