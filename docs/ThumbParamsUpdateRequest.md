# ThumbParamsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**ThumbParams** | [**KalturaThumbParams**](KalturaThumbParams.md) |  | 

## Methods

### NewThumbParamsUpdateRequest

`func NewThumbParamsUpdateRequest(id int32, thumbParams KalturaThumbParams, ) *ThumbParamsUpdateRequest`

NewThumbParamsUpdateRequest instantiates a new ThumbParamsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbParamsUpdateRequestWithDefaults

`func NewThumbParamsUpdateRequestWithDefaults() *ThumbParamsUpdateRequest`

NewThumbParamsUpdateRequestWithDefaults instantiates a new ThumbParamsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThumbParamsUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThumbParamsUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThumbParamsUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetThumbParams

`func (o *ThumbParamsUpdateRequest) GetThumbParams() KalturaThumbParams`

GetThumbParams returns the ThumbParams field if non-nil, zero value otherwise.

### GetThumbParamsOk

`func (o *ThumbParamsUpdateRequest) GetThumbParamsOk() (*KalturaThumbParams, bool)`

GetThumbParamsOk returns a tuple with the ThumbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbParams

`func (o *ThumbParamsUpdateRequest) SetThumbParams(v KalturaThumbParams)`

SetThumbParams sets ThumbParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


