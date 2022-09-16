# KalturaSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthData** | Pointer to **string** |  | [optional] 
**ExtraData** | Pointer to **string** | Use this field to pass dynamic data for searching  For example - if you set this field to \&quot;mymovies_$partner_id\&quot;  The $partner_id will be automatically replcaed with your real partner Id | [optional] 
**KeyWords** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **int32** | Enum Type: &#x60;KalturaMediaType&#x60; | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**SearchSource** | Pointer to **int32** | Enum Type: &#x60;KalturaSearchProviderType&#x60; | [optional] 

## Methods

### NewKalturaSearch

`func NewKalturaSearch() *KalturaSearch`

NewKalturaSearch instantiates a new KalturaSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaSearchWithDefaults

`func NewKalturaSearchWithDefaults() *KalturaSearch`

NewKalturaSearchWithDefaults instantiates a new KalturaSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthData

`func (o *KalturaSearch) GetAuthData() string`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *KalturaSearch) GetAuthDataOk() (*string, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *KalturaSearch) SetAuthData(v string)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *KalturaSearch) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.

### GetExtraData

`func (o *KalturaSearch) GetExtraData() string`

GetExtraData returns the ExtraData field if non-nil, zero value otherwise.

### GetExtraDataOk

`func (o *KalturaSearch) GetExtraDataOk() (*string, bool)`

GetExtraDataOk returns a tuple with the ExtraData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraData

`func (o *KalturaSearch) SetExtraData(v string)`

SetExtraData sets ExtraData field to given value.

### HasExtraData

`func (o *KalturaSearch) HasExtraData() bool`

HasExtraData returns a boolean if a field has been set.

### GetKeyWords

`func (o *KalturaSearch) GetKeyWords() string`

GetKeyWords returns the KeyWords field if non-nil, zero value otherwise.

### GetKeyWordsOk

`func (o *KalturaSearch) GetKeyWordsOk() (*string, bool)`

GetKeyWordsOk returns a tuple with the KeyWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyWords

`func (o *KalturaSearch) SetKeyWords(v string)`

SetKeyWords sets KeyWords field to given value.

### HasKeyWords

`func (o *KalturaSearch) HasKeyWords() bool`

HasKeyWords returns a boolean if a field has been set.

### GetMediaType

`func (o *KalturaSearch) GetMediaType() int32`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *KalturaSearch) GetMediaTypeOk() (*int32, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *KalturaSearch) SetMediaType(v int32)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *KalturaSearch) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaSearch) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaSearch) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaSearch) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaSearch) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetSearchSource

`func (o *KalturaSearch) GetSearchSource() int32`

GetSearchSource returns the SearchSource field if non-nil, zero value otherwise.

### GetSearchSourceOk

`func (o *KalturaSearch) GetSearchSourceOk() (*int32, bool)`

GetSearchSourceOk returns a tuple with the SearchSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchSource

`func (o *KalturaSearch) SetSearchSource(v int32)`

SetSearchSource sets SearchSource field to given value.

### HasSearchSource

`func (o *KalturaSearch) HasSearchSource() bool`

HasSearchSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


