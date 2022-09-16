# KalturaCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Not** | Pointer to **bool** |  | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaConditionType&#x60;  The type of the access control condition | [optional] [readonly] 

## Methods

### NewKalturaCondition

`func NewKalturaCondition() *KalturaCondition`

NewKalturaCondition instantiates a new KalturaCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaConditionWithDefaults

`func NewKalturaConditionWithDefaults() *KalturaCondition`

NewKalturaConditionWithDefaults instantiates a new KalturaCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KalturaCondition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaCondition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaCondition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaCondition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNot

`func (o *KalturaCondition) GetNot() bool`

GetNot returns the Not field if non-nil, zero value otherwise.

### GetNotOk

`func (o *KalturaCondition) GetNotOk() (*bool, bool)`

GetNotOk returns a tuple with the Not field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNot

`func (o *KalturaCondition) SetNot(v bool)`

SetNot sets Not field to given value.

### HasNot

`func (o *KalturaCondition) HasNot() bool`

HasNot returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaCondition) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaCondition) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaCondition) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaCondition) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetType

`func (o *KalturaCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


