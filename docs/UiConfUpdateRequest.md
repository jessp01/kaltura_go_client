# UiConfUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**UiConf** | [**KalturaUiConf**](KalturaUiConf.md) |  | 

## Methods

### NewUiConfUpdateRequest

`func NewUiConfUpdateRequest(id int32, uiConf KalturaUiConf, ) *UiConfUpdateRequest`

NewUiConfUpdateRequest instantiates a new UiConfUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiConfUpdateRequestWithDefaults

`func NewUiConfUpdateRequestWithDefaults() *UiConfUpdateRequest`

NewUiConfUpdateRequestWithDefaults instantiates a new UiConfUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UiConfUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UiConfUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UiConfUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUiConf

`func (o *UiConfUpdateRequest) GetUiConf() KalturaUiConf`

GetUiConf returns the UiConf field if non-nil, zero value otherwise.

### GetUiConfOk

`func (o *UiConfUpdateRequest) GetUiConfOk() (*KalturaUiConf, bool)`

GetUiConfOk returns a tuple with the UiConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiConf

`func (o *UiConfUpdateRequest) SetUiConf(v KalturaUiConf)`

SetUiConf sets UiConf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


