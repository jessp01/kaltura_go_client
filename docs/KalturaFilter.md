# KalturaFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedSearch** | Pointer to [**KalturaSearchItem**](KalturaSearchItem.md) |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaFilter

`func NewKalturaFilter() *KalturaFilter`

NewKalturaFilter instantiates a new KalturaFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaFilterWithDefaults

`func NewKalturaFilterWithDefaults() *KalturaFilter`

NewKalturaFilterWithDefaults instantiates a new KalturaFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedSearch

`func (o *KalturaFilter) GetAdvancedSearch() KalturaSearchItem`

GetAdvancedSearch returns the AdvancedSearch field if non-nil, zero value otherwise.

### GetAdvancedSearchOk

`func (o *KalturaFilter) GetAdvancedSearchOk() (*KalturaSearchItem, bool)`

GetAdvancedSearchOk returns a tuple with the AdvancedSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSearch

`func (o *KalturaFilter) SetAdvancedSearch(v KalturaSearchItem)`

SetAdvancedSearch sets AdvancedSearch field to given value.

### HasAdvancedSearch

`func (o *KalturaFilter) HasAdvancedSearch() bool`

HasAdvancedSearch returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaFilter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaFilter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaFilter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaFilter) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetOrderBy

`func (o *KalturaFilter) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *KalturaFilter) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *KalturaFilter) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *KalturaFilter) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


