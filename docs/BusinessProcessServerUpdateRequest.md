# BusinessProcessServerUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessProcessServer** | [**KalturaBusinessProcessServer**](KalturaBusinessProcessServer.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewBusinessProcessServerUpdateRequest

`func NewBusinessProcessServerUpdateRequest(businessProcessServer KalturaBusinessProcessServer, id int32, ) *BusinessProcessServerUpdateRequest`

NewBusinessProcessServerUpdateRequest instantiates a new BusinessProcessServerUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessProcessServerUpdateRequestWithDefaults

`func NewBusinessProcessServerUpdateRequestWithDefaults() *BusinessProcessServerUpdateRequest`

NewBusinessProcessServerUpdateRequestWithDefaults instantiates a new BusinessProcessServerUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessProcessServer

`func (o *BusinessProcessServerUpdateRequest) GetBusinessProcessServer() KalturaBusinessProcessServer`

GetBusinessProcessServer returns the BusinessProcessServer field if non-nil, zero value otherwise.

### GetBusinessProcessServerOk

`func (o *BusinessProcessServerUpdateRequest) GetBusinessProcessServerOk() (*KalturaBusinessProcessServer, bool)`

GetBusinessProcessServerOk returns a tuple with the BusinessProcessServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessProcessServer

`func (o *BusinessProcessServerUpdateRequest) SetBusinessProcessServer(v KalturaBusinessProcessServer)`

SetBusinessProcessServer sets BusinessProcessServer field to given value.


### GetId

`func (o *BusinessProcessServerUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BusinessProcessServerUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BusinessProcessServerUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


