# EntryDistributionSubmitAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**SubmitWhenReady** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewEntryDistributionSubmitAddRequest

`func NewEntryDistributionSubmitAddRequest(id int32, ) *EntryDistributionSubmitAddRequest`

NewEntryDistributionSubmitAddRequest instantiates a new EntryDistributionSubmitAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryDistributionSubmitAddRequestWithDefaults

`func NewEntryDistributionSubmitAddRequestWithDefaults() *EntryDistributionSubmitAddRequest`

NewEntryDistributionSubmitAddRequestWithDefaults instantiates a new EntryDistributionSubmitAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntryDistributionSubmitAddRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntryDistributionSubmitAddRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntryDistributionSubmitAddRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSubmitWhenReady

`func (o *EntryDistributionSubmitAddRequest) GetSubmitWhenReady() bool`

GetSubmitWhenReady returns the SubmitWhenReady field if non-nil, zero value otherwise.

### GetSubmitWhenReadyOk

`func (o *EntryDistributionSubmitAddRequest) GetSubmitWhenReadyOk() (*bool, bool)`

GetSubmitWhenReadyOk returns a tuple with the SubmitWhenReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitWhenReady

`func (o *EntryDistributionSubmitAddRequest) SetSubmitWhenReady(v bool)`

SetSubmitWhenReady sets SubmitWhenReady field to given value.

### HasSubmitWhenReady

`func (o *EntryDistributionSubmitAddRequest) HasSubmitWhenReady() bool`

HasSubmitWhenReady returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


