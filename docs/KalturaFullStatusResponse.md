# KalturaFullStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueuesStatus** | Pointer to [**[]KalturaBatchQueuesStatus**](KalturaBatchQueuesStatus.md) |  | [optional] 
**Schedulers** | Pointer to [**[]KalturaScheduler**](KalturaScheduler.md) |  | [optional] 

## Methods

### NewKalturaFullStatusResponse

`func NewKalturaFullStatusResponse() *KalturaFullStatusResponse`

NewKalturaFullStatusResponse instantiates a new KalturaFullStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaFullStatusResponseWithDefaults

`func NewKalturaFullStatusResponseWithDefaults() *KalturaFullStatusResponse`

NewKalturaFullStatusResponseWithDefaults instantiates a new KalturaFullStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueuesStatus

`func (o *KalturaFullStatusResponse) GetQueuesStatus() []KalturaBatchQueuesStatus`

GetQueuesStatus returns the QueuesStatus field if non-nil, zero value otherwise.

### GetQueuesStatusOk

`func (o *KalturaFullStatusResponse) GetQueuesStatusOk() (*[]KalturaBatchQueuesStatus, bool)`

GetQueuesStatusOk returns a tuple with the QueuesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuesStatus

`func (o *KalturaFullStatusResponse) SetQueuesStatus(v []KalturaBatchQueuesStatus)`

SetQueuesStatus sets QueuesStatus field to given value.

### HasQueuesStatus

`func (o *KalturaFullStatusResponse) HasQueuesStatus() bool`

HasQueuesStatus returns a boolean if a field has been set.

### GetSchedulers

`func (o *KalturaFullStatusResponse) GetSchedulers() []KalturaScheduler`

GetSchedulers returns the Schedulers field if non-nil, zero value otherwise.

### GetSchedulersOk

`func (o *KalturaFullStatusResponse) GetSchedulersOk() (*[]KalturaScheduler, bool)`

GetSchedulersOk returns a tuple with the Schedulers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulers

`func (o *KalturaFullStatusResponse) SetSchedulers(v []KalturaScheduler)`

SetSchedulers sets Schedulers field to given value.

### HasSchedulers

`func (o *KalturaFullStatusResponse) HasSchedulers() bool`

HasSchedulers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


