# FlavorAssetGetUrlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceProxy** | Pointer to **bool** |  | [optional] [default to false]
**Id** | **string** |  | 
**Options** | Pointer to [**KalturaFlavorAssetUrlOptions**](KalturaFlavorAssetUrlOptions.md) |  | [optional] 
**StorageId** | Pointer to **int32** |  | [optional] 

## Methods

### NewFlavorAssetGetUrlRequest

`func NewFlavorAssetGetUrlRequest(id string, ) *FlavorAssetGetUrlRequest`

NewFlavorAssetGetUrlRequest instantiates a new FlavorAssetGetUrlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorAssetGetUrlRequestWithDefaults

`func NewFlavorAssetGetUrlRequestWithDefaults() *FlavorAssetGetUrlRequest`

NewFlavorAssetGetUrlRequestWithDefaults instantiates a new FlavorAssetGetUrlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceProxy

`func (o *FlavorAssetGetUrlRequest) GetForceProxy() bool`

GetForceProxy returns the ForceProxy field if non-nil, zero value otherwise.

### GetForceProxyOk

`func (o *FlavorAssetGetUrlRequest) GetForceProxyOk() (*bool, bool)`

GetForceProxyOk returns a tuple with the ForceProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceProxy

`func (o *FlavorAssetGetUrlRequest) SetForceProxy(v bool)`

SetForceProxy sets ForceProxy field to given value.

### HasForceProxy

`func (o *FlavorAssetGetUrlRequest) HasForceProxy() bool`

HasForceProxy returns a boolean if a field has been set.

### GetId

`func (o *FlavorAssetGetUrlRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlavorAssetGetUrlRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlavorAssetGetUrlRequest) SetId(v string)`

SetId sets Id field to given value.


### GetOptions

`func (o *FlavorAssetGetUrlRequest) GetOptions() KalturaFlavorAssetUrlOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FlavorAssetGetUrlRequest) GetOptionsOk() (*KalturaFlavorAssetUrlOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FlavorAssetGetUrlRequest) SetOptions(v KalturaFlavorAssetUrlOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FlavorAssetGetUrlRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetStorageId

`func (o *FlavorAssetGetUrlRequest) GetStorageId() int32`

GetStorageId returns the StorageId field if non-nil, zero value otherwise.

### GetStorageIdOk

`func (o *FlavorAssetGetUrlRequest) GetStorageIdOk() (*int32, bool)`

GetStorageIdOk returns a tuple with the StorageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageId

`func (o *FlavorAssetGetUrlRequest) SetStorageId(v int32)`

SetStorageId sets StorageId field to given value.

### HasStorageId

`func (o *FlavorAssetGetUrlRequest) HasStorageId() bool`

HasStorageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


