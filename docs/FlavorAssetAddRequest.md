# FlavorAssetAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**FlavorAsset** | [**KalturaFlavorAsset**](KalturaFlavorAsset.md) |  | 

## Methods

### NewFlavorAssetAddRequest

`func NewFlavorAssetAddRequest(entryId string, flavorAsset KalturaFlavorAsset, ) *FlavorAssetAddRequest`

NewFlavorAssetAddRequest instantiates a new FlavorAssetAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlavorAssetAddRequestWithDefaults

`func NewFlavorAssetAddRequestWithDefaults() *FlavorAssetAddRequest`

NewFlavorAssetAddRequestWithDefaults instantiates a new FlavorAssetAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *FlavorAssetAddRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *FlavorAssetAddRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *FlavorAssetAddRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetFlavorAsset

`func (o *FlavorAssetAddRequest) GetFlavorAsset() KalturaFlavorAsset`

GetFlavorAsset returns the FlavorAsset field if non-nil, zero value otherwise.

### GetFlavorAssetOk

`func (o *FlavorAssetAddRequest) GetFlavorAssetOk() (*KalturaFlavorAsset, bool)`

GetFlavorAssetOk returns a tuple with the FlavorAsset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorAsset

`func (o *FlavorAssetAddRequest) SetFlavorAsset(v KalturaFlavorAsset)`

SetFlavorAsset sets FlavorAsset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


