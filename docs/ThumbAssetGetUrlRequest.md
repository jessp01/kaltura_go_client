# ThumbAssetGetUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**StorageId** | Pointer to **int32** |  | [optional] 
**ThumbParams** | Pointer to [**KalturaThumbParams**](KalturaThumbParams.md) |  | [optional] 

## Methods

### NewThumbAssetGetUrlRequest

`func NewThumbAssetGetUrlRequest(id string, ) *ThumbAssetGetUrlRequest`

NewThumbAssetGetUrlRequest instantiates a new ThumbAssetGetUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbAssetGetUrlRequestWithDefaults

`func NewThumbAssetGetUrlRequestWithDefaults() *ThumbAssetGetUrlRequest`

NewThumbAssetGetUrlRequestWithDefaults instantiates a new ThumbAssetGetUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ThumbAssetGetUrlRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ThumbAssetGetUrlRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ThumbAssetGetUrlRequest) SetId(v string)`

SetId sets Id field to given value.


### GetStorageId

`func (o *ThumbAssetGetUrlRequest) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *ThumbAssetGetUrlRequest) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *ThumbAssetGetUrlRequest) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *ThumbAssetGetUrlRequest) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.

### GetThumbParams

`func (o *ThumbAssetGetUrlRequest) GetThumbParams() KalturaThumbParams`

GetThumbParams returns the ThumbParams field if non-nil, zero value otherwise.

### GetThumbParamsOk

`func (o *ThumbAssetGetUrlRequest) GetThumbParamsOk() (*KalturaThumbParams, bool)`

GetThumbParamsOk returns a tuple with the ThumbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbParams

`func (o *ThumbAssetGetUrlRequest) SetThumbParams(v KalturaThumbParams)`

SetThumbParams sets ThumbParams field to given value.

### HasThumbParams

`func (o *ThumbAssetGetUrlRequest) HasThumbParams() bool`

HasThumbParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


