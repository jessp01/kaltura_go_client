# MetadataInvalidateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Version** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetadataInvalidateRequest

`func NewMetadataInvalidateRequest(id int32, ) *MetadataInvalidateRequest`

NewMetadataInvalidateRequest instantiates a new MetadataInvalidateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataInvalidateRequestWithDefaults

`func NewMetadataInvalidateRequestWithDefaults() *MetadataInvalidateRequest`

NewMetadataInvalidateRequestWithDefaults instantiates a new MetadataInvalidateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataInvalidateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataInvalidateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataInvalidateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVersion

`func (o *MetadataInvalidateRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetadataInvalidateRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetadataInvalidateRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetadataInvalidateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


