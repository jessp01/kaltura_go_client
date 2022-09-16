# FileAssetUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileAsset** | [**KalturaFileAsset**](KalturaFileAsset.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewFileAssetUpdateRequest

`func NewFileAssetUpdateRequest(fileAsset KalturaFileAsset, id int32, ) *FileAssetUpdateRequest`

NewFileAssetUpdateRequest instantiates a new FileAssetUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAssetUpdateRequestWithDefaults

`func NewFileAssetUpdateRequestWithDefaults() *FileAssetUpdateRequest`

NewFileAssetUpdateRequestWithDefaults instantiates a new FileAssetUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileAsset

`func (o *FileAssetUpdateRequest) GetFileAsset() KalturaFileAsset`

GetFileAsset returns the FileAsset field if non-nil, zero value otherwise.

### GetFileAssetOk

`func (o *FileAssetUpdateRequest) GetFileAssetOk() (*KalturaFileAsset, bool)`

GetFileAssetOk returns a tuple with the FileAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileAsset

`func (o *FileAssetUpdateRequest) SetFileAsset(v KalturaFileAsset)`

SetFileAsset sets FileAsset field to given value.


### GetId

`func (o *FileAssetUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileAssetUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileAssetUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


