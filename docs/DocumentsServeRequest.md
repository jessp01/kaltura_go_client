# DocumentsServeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryId** | **string** |  | 
**FlavorAssetId** | Pointer to **string** |  | [optional] 
**ForceProxy** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewDocumentsServeRequest

`func NewDocumentsServeRequest(entryId string, ) *DocumentsServeRequest`

NewDocumentsServeRequest instantiates a new DocumentsServeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentsServeRequestWithDefaults

`func NewDocumentsServeRequestWithDefaults() *DocumentsServeRequest`

NewDocumentsServeRequestWithDefaults instantiates a new DocumentsServeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryId

`func (o *DocumentsServeRequest) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *DocumentsServeRequest) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *DocumentsServeRequest) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.


### GetFlavorAssetId

`func (o *DocumentsServeRequest) GetFlavorAssetId() string`

GetFlavorAssetId returns the FlavorAssetId field if non-nil, zero value otherwise.

### GetFlavorAssetIdOk

`func (o *DocumentsServeRequest) GetFlavorAssetIdOk() (*string, bool)`

GetFlavorAssetIdOk returns a tuple with the FlavorAssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorAssetId

`func (o *DocumentsServeRequest) SetFlavorAssetId(v string)`

SetFlavorAssetId sets FlavorAssetId field to given value.

### HasFlavorAssetId

`func (o *DocumentsServeRequest) HasFlavorAssetId() bool`

HasFlavorAssetId returns a boolean if a field has been set.

### GetForceProxy

`func (o *DocumentsServeRequest) GetForceProxy() bool`

GetForceProxy returns the ForceProxy field if non-nil, zero value otherwise.

### GetForceProxyOk

`func (o *DocumentsServeRequest) GetForceProxyOk() (*bool, bool)`

GetForceProxyOk returns a tuple with the ForceProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceProxy

`func (o *DocumentsServeRequest) SetForceProxy(v bool)`

SetForceProxy sets ForceProxy field to given value.

### HasForceProxy

`func (o *DocumentsServeRequest) HasForceProxy() bool`

HasForceProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


