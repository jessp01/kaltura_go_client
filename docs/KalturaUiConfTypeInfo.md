# KalturaUiConfTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Directory** | Pointer to **string** | The direcotry this type is saved at | [optional] 
**Filename** | Pointer to **string** | Filename for this UiConf type | [optional] 
**Type** | Pointer to **int32** | Enum Type: &#x60;KalturaUiConfObjType&#x60;  UiConf Type | [optional] 
**Versions** | Pointer to [**[]KalturaString**](KalturaString.md) |  | [optional] 

## Methods

### NewKalturaUiConfTypeInfo

`func NewKalturaUiConfTypeInfo() *KalturaUiConfTypeInfo`

NewKalturaUiConfTypeInfo instantiates a new KalturaUiConfTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaUiConfTypeInfoWithDefaults

`func NewKalturaUiConfTypeInfoWithDefaults() *KalturaUiConfTypeInfo`

NewKalturaUiConfTypeInfoWithDefaults instantiates a new KalturaUiConfTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectory

`func (o *KalturaUiConfTypeInfo) GetDirectory() string`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *KalturaUiConfTypeInfo) GetDirectoryOk() (*string, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *KalturaUiConfTypeInfo) SetDirectory(v string)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *KalturaUiConfTypeInfo) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

### GetFilename

`func (o *KalturaUiConfTypeInfo) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *KalturaUiConfTypeInfo) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *KalturaUiConfTypeInfo) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *KalturaUiConfTypeInfo) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetType

`func (o *KalturaUiConfTypeInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaUiConfTypeInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaUiConfTypeInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaUiConfTypeInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersions

`func (o *KalturaUiConfTypeInfo) GetVersions() []KalturaString`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *KalturaUiConfTypeInfo) GetVersionsOk() (*[]KalturaString, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *KalturaUiConfTypeInfo) SetVersions(v []KalturaString)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *KalturaUiConfTypeInfo) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


