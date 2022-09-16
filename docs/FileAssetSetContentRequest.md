# FileAssetSetContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentResource** | [**KalturaContentResource**](KalturaContentResource.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewFileAssetSetContentRequest

`func NewFileAssetSetContentRequest(contentResource KalturaContentResource, id int32, ) *FileAssetSetContentRequest`

NewFileAssetSetContentRequest instantiates a new FileAssetSetContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFileAssetSetContentRequestWithDefaults

`func NewFileAssetSetContentRequestWithDefaults() *FileAssetSetContentRequest`

NewFileAssetSetContentRequestWithDefaults instantiates a new FileAssetSetContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentResource

`func (o *FileAssetSetContentRequest) GetContentResource() KalturaContentResource`

GetContentResource returns the ContentResource field if non-nil, zero value otherwise.

### GetContentResourceOk

`func (o *FileAssetSetContentRequest) GetContentResourceOk() (*KalturaContentResource, bool)`

GetContentResourceOk returns a tuple with the ContentResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentResource

`func (o *FileAssetSetContentRequest) SetContentResource(v KalturaContentResource)`

SetContentResource sets ContentResource field to given value.


### GetId

`func (o *FileAssetSetContentRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FileAssetSetContentRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FileAssetSetContentRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


