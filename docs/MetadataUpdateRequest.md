# MetadataUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Version** | Pointer to **int32** |  | [optional] 
**XmlData** | Pointer to **string** |  | [optional] 

## Methods

### NewMetadataUpdateRequest

`func NewMetadataUpdateRequest(id int32, ) *MetadataUpdateRequest`

NewMetadataUpdateRequest instantiates a new MetadataUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataUpdateRequestWithDefaults

`func NewMetadataUpdateRequestWithDefaults() *MetadataUpdateRequest`

NewMetadataUpdateRequestWithDefaults instantiates a new MetadataUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetadataUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetadataUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetadataUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetVersion

`func (o *MetadataUpdateRequest) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MetadataUpdateRequest) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MetadataUpdateRequest) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MetadataUpdateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetXmlData

`func (o *MetadataUpdateRequest) GetXmlData() string`

GetXmlData returns the XmlData field if non-nil, zero value otherwise.

### GetXmlDataOk

`func (o *MetadataUpdateRequest) GetXmlDataOk() (*string, bool)`

GetXmlDataOk returns a tuple with the XmlData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXmlData

`func (o *MetadataUpdateRequest) SetXmlData(v string)`

SetXmlData sets XmlData field to given value.

### HasXmlData

`func (o *MetadataUpdateRequest) HasXmlData() bool`

HasXmlData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


