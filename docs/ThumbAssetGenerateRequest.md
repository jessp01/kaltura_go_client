# ThumbAssetGenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**SourceAssetId** | Pointer to **string** |  | [optional] 
**ThumbParams** | [**KalturaThumbParams**](KalturaThumbParams.md) |  | 

## Methods

### NewThumbAssetGenerateRequest

`func NewThumbAssetGenerateRequest(entryId string, thumbParams KalturaThumbParams, ) *ThumbAssetGenerateRequest`

NewThumbAssetGenerateRequest instantiates a new ThumbAssetGenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThumbAssetGenerateRequestWithDefaults

`func NewThumbAssetGenerateRequestWithDefaults() *ThumbAssetGenerateRequest`

NewThumbAssetGenerateRequestWithDefaults instantiates a new ThumbAssetGenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *ThumbAssetGenerateRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *ThumbAssetGenerateRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *ThumbAssetGenerateRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetSourceAssetId

`func (o *ThumbAssetGenerateRequest) GetSourceAssetId() string`

GetSourceAssetId returns the SourceAssetId field if non-nil, zero value otherwise.

### GetSourceAssetIdOk

`func (o *ThumbAssetGenerateRequest) GetSourceAssetIdOk() (*string, bool)`

GetSourceAssetIdOk returns a tuple with the SourceAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAssetId

`func (o *ThumbAssetGenerateRequest) SetSourceAssetId(v string)`

SetSourceAssetId sets SourceAssetId field to given value.

### HasSourceAssetId

`func (o *ThumbAssetGenerateRequest) HasSourceAssetId() bool`

HasSourceAssetId returns a boolean if a field has been set.

### GetThumbParams

`func (o *ThumbAssetGenerateRequest) GetThumbParams() KalturaThumbParams`

GetThumbParams returns the ThumbParams field if non-nil, zero value otherwise.

### GetThumbParamsOk

`func (o *ThumbAssetGenerateRequest) GetThumbParamsOk() (*KalturaThumbParams, bool)`

GetThumbParamsOk returns a tuple with the ThumbParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbParams

`func (o *ThumbAssetGenerateRequest) SetThumbParams(v KalturaThumbParams)`

SetThumbParams sets ThumbParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


