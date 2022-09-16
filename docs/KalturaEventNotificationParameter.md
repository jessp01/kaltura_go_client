# KalturaEventNotificationParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** | The key in the subject and body to be replaced with the dynamic value | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**KalturaStringValue**](KalturaStringValue.md) |  | [optional] 

## Methods

### NewKalturaEventNotificationParameter

`func NewKalturaEventNotificationParameter() *KalturaEventNotificationParameter`

NewKalturaEventNotificationParameter instantiates a new KalturaEventNotificationParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaEventNotificationParameterWithDefaults

`func NewKalturaEventNotificationParameterWithDefaults() *KalturaEventNotificationParameter`

NewKalturaEventNotificationParameterWithDefaults instantiates a new KalturaEventNotificationParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KalturaEventNotificationParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaEventNotificationParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaEventNotificationParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaEventNotificationParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *KalturaEventNotificationParameter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KalturaEventNotificationParameter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KalturaEventNotificationParameter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KalturaEventNotificationParameter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetObjectType

`func (o *KalturaEventNotificationParameter) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *KalturaEventNotificationParameter) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *KalturaEventNotificationParameter) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *KalturaEventNotificationParameter) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetValue

`func (o *KalturaEventNotificationParameter) GetValue() KalturaStringValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *KalturaEventNotificationParameter) GetValueOk() (*KalturaStringValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *KalturaEventNotificationParameter) SetValue(v KalturaStringValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *KalturaEventNotificationParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


