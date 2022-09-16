# KalturaTransformMetadataResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LowerVersionCount** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Objects** | Pointer to [**[]KalturaMetadata**](KalturaMetadata.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 

## Methods

### NewKalturaTransformMetadataResponse

`func NewKalturaTransformMetadataResponse() *KalturaTransformMetadataResponse`

NewKalturaTransformMetadataResponse instantiates a new KalturaTransformMetadataResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaTransformMetadataResponseWithDefaults

`func NewKalturaTransformMetadataResponseWithDefaults() *KalturaTransformMetadataResponse`

NewKalturaTransformMetadataResponseWithDefaults instantiates a new KalturaTransformMetadataResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLowerVersionCount

`func (o *KalturaTransformMetadataResponse) GetLowerVersionCount() int32`

GetLowerVersionCount returns the LowerVersionCount field if non-nil, zero value otherwise.

### GetLowerVersionCountOk

`func (o *KalturaTransformMetadataResponse) GetLowerVersionCountOk() (*int32, bool)`

GetLowerVersionCountOk returns a tuple with the LowerVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerVersionCount

`func (o *KalturaTransformMetadataResponse) SetLowerVersionCount(v int32)`

SetLowerVersionCount sets LowerVersionCount field to given value.

### HasLowerVersionCount

`func (o *KalturaTransformMetadataResponse) HasLowerVersionCount() bool`

HasLowerVersionCount returns a boolean if a field has been set.

### GetObjects

`func (o *KalturaTransformMetadataResponse) GetObjects() []KalturaMetadata`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *KalturaTransformMetadataResponse) GetObjectsOk() (*[]KalturaMetadata, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *KalturaTransformMetadataResponse) SetObjects(v []KalturaMetadata)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *KalturaTransformMetadataResponse) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetTotalCount

`func (o *KalturaTransformMetadataResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *KalturaTransformMetadataResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *KalturaTransformMetadataResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *KalturaTransformMetadataResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


