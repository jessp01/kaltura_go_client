# ThumbAssetServeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to [**KalturaThumbnailServeOptions**](KalturaThumbnailServeOptions.md) |  | [optional] 
**ThumbAssetId** | **string** |  | 
**ThumbParams** | Pointer to [**KalturaThumbParams**](KalturaThumbParams.md) |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewThumbAssetServeRequest

`func NewThumbAssetServeRequest(thumbAssetId string, ) *ThumbAssetServeRequest`

NewThumbAssetServeRequest instantiates a new ThumbAssetServeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbAssetServeRequestWithDefaults

`func NewThumbAssetServeRequestWithDefaults() *ThumbAssetServeRequest`

NewThumbAssetServeRequestWithDefaults instantiates a new ThumbAssetServeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *ThumbAssetServeRequest) GetOptions() KalturaThumbnailServeOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *ThumbAssetServeRequest) GetOptionsOk() (*KalturaThumbnailServeOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *ThumbAssetServeRequest) SetOptions(v KalturaThumbnailServeOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *ThumbAssetServeRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetThumbAssetId

`func (o *ThumbAssetServeRequest) GetThumbAssetId() string`

GetThumbAssetId returns the ThumbAssetId field if non-nil, zero value otherwise.

### GetThumbAssetIdOk

`func (o *ThumbAssetServeRequest) GetThumbAssetIdOk() (*string, bool)`

GetThumbAssetIdOk returns a tuple with the ThumbAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbAssetId

`func (o *ThumbAssetServeRequest) SetThumbAssetId(v string)`

SetThumbAssetId sets ThumbAssetId field to given value.


### GetThumbParams

`func (o *ThumbAssetServeRequest) GetThumbParams() KalturaThumbParams`

GetThumbParams returns the ThumbParams field if non-nil, zero value otherwise.

### GetThumbParamsOk

`func (o *ThumbAssetServeRequest) GetThumbParamsOk() (*KalturaThumbParams, bool)`

GetThumbParamsOk returns a tuple with the ThumbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbParams

`func (o *ThumbAssetServeRequest) SetThumbParams(v KalturaThumbParams)`

SetThumbParams sets ThumbParams field to given value.

### HasThumbParams

`func (o *ThumbAssetServeRequest) HasThumbParams() bool`

HasThumbParams returns a boolean if a field has been set.

### GetVersion

`func (o *ThumbAssetServeRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ThumbAssetServeRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ThumbAssetServeRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ThumbAssetServeRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


