# KalturaObjectIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtendedFeatures** | Pointer to **string** | Comma separated string of enum values denoting which features of the item need to be included in the MRSS | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaObjectIdentifier

`func NewKalturaObjectIdentifier() *KalturaObjectIdentifier`

NewKalturaObjectIdentifier instantiates a new KalturaObjectIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaObjectIdentifierWithDefaults

`func NewKalturaObjectIdentifierWithDefaults() *KalturaObjectIdentifier`

NewKalturaObjectIdentifierWithDefaults instantiates a new KalturaObjectIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtendedFeatures

`func (o *KalturaObjectIdentifier) GetExtendedFeatures() string`

GetExtendedFeatures returns the ExtendedFeatures field if non-nil, zero value otherwise.

### GetExtendedFeaturesOk

`func (o *KalturaObjectIdentifier) GetExtendedFeaturesOk() (*string, bool)`

GetExtendedFeaturesOk returns a tuple with the ExtendedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedFeatures

`func (o *KalturaObjectIdentifier) SetExtendedFeatures(v string)`

SetExtendedFeatures sets ExtendedFeatures field to given value.

### HasExtendedFeatures

`func (o *KalturaObjectIdentifier) HasExtendedFeatures() bool`

HasExtendedFeatures returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaObjectIdentifier) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaObjectIdentifier) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaObjectIdentifier) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaObjectIdentifier) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


