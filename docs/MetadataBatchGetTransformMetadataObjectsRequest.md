# MetadataBatchGetTransformMetadataObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestVersion** | **int32** |  | 
**MetadataProfileId** | **int32** |  | 
**Pager** | Pointer to [**KalturaFilterPager**](KalturaFilterPager.md) |  | [optional] 
**SrcVersion** | **int32** |  | 

## Methods

### NewMetadataBatchGetTransformMetadataObjectsRequest

`func NewMetadataBatchGetTransformMetadataObjectsRequest(destVersion int32, metadataProfileId int32, srcVersion int32, ) *MetadataBatchGetTransformMetadataObjectsRequest`

NewMetadataBatchGetTransformMetadataObjectsRequest instantiates a new MetadataBatchGetTransformMetadataObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataBatchGetTransformMetadataObjectsRequestWithDefaults

`func NewMetadataBatchGetTransformMetadataObjectsRequestWithDefaults() *MetadataBatchGetTransformMetadataObjectsRequest`

NewMetadataBatchGetTransformMetadataObjectsRequestWithDefaults instantiates a new MetadataBatchGetTransformMetadataObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestVersion

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetDestVersion() int32`

GetDestVersion returns the DestVersion field if non-nil, zero value otherwise.

### GetDestVersionOk

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetDestVersionOk() (*int32, bool)`

GetDestVersionOk returns a tuple with the DestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestVersion

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) SetDestVersion(v int32)`

SetDestVersion sets DestVersion field to given value.


### GetMetadataProfileId

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetMetadataProfileId() int32`

GetMetadataProfileId returns the MetadataProfileId field if non-nil, zero value otherwise.

### GetMetadataProfileIdOk

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetMetadataProfileIdOk() (*int32, bool)`

GetMetadataProfileIdOk returns a tuple with the MetadataProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataProfileId

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) SetMetadataProfileId(v int32)`

SetMetadataProfileId sets MetadataProfileId field to given value.


### GetPager

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetPager() KalturaFilterPager`

GetPager returns the Pager field if non-nil, zero value otherwise.

### GetPagerOk

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetPagerOk() (*KalturaFilterPager, bool)`

GetPagerOk returns a tuple with the Pager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPager

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) SetPager(v KalturaFilterPager)`

SetPager sets Pager field to given value.

### HasPager

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) HasPager() bool`

HasPager returns a boolean if a field has been set.

### GetSrcVersion

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetSrcVersion() int32`

GetSrcVersion returns the SrcVersion field if non-nil, zero value otherwise.

### GetSrcVersionOk

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) GetSrcVersionOk() (*int32, bool)`

GetSrcVersionOk returns a tuple with the SrcVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcVersion

`func (o *MetadataBatchGetTransformMetadataObjectsRequest) SetSrcVersion(v int32)`

SetSrcVersion sets SrcVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


