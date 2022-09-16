# ConversionProfileUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversionProfile** | [**KalturaConversionProfile**](KalturaConversionProfile.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewConversionProfileUpdateRequest

`func NewConversionProfileUpdateRequest(conversionProfile KalturaConversionProfile, id int32, ) *ConversionProfileUpdateRequest`

NewConversionProfileUpdateRequest instantiates a new ConversionProfileUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionProfileUpdateRequestWithDefaults

`func NewConversionProfileUpdateRequestWithDefaults() *ConversionProfileUpdateRequest`

NewConversionProfileUpdateRequestWithDefaults instantiates a new ConversionProfileUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversionProfile

`func (o *ConversionProfileUpdateRequest) GetConversionProfile() KalturaConversionProfile`

GetConversionProfile returns the ConversionProfile field if non-nil, zero value otherwise.

### GetConversionProfileOk

`func (o *ConversionProfileUpdateRequest) GetConversionProfileOk() (*KalturaConversionProfile, bool)`

GetConversionProfileOk returns a tuple with the ConversionProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionProfile

`func (o *ConversionProfileUpdateRequest) SetConversionProfile(v KalturaConversionProfile)`

SetConversionProfile sets ConversionProfile field to given value.


### GetId

`func (o *ConversionProfileUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversionProfileUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversionProfileUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


