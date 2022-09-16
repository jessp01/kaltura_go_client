# KalturaExtendingItemMrssParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionMode** | Pointer to **int32** | Enum Type: &#x60;KalturaMrssExtensionMode&#x60;  Mode of extension - append to MRSS or replace the xpath content. | [optional] 
**Identifier** | Pointer to [**KalturaObjectIdentifier**](KalturaObjectIdentifier.md) |  | [optional] 
**Xpath** | Pointer to **string** | XPath for the extending item | [optional] 

## Methods

### NewKalturaExtendingItemMrssParameter

`func NewKalturaExtendingItemMrssParameter() *KalturaExtendingItemMrssParameter`

NewKalturaExtendingItemMrssParameter instantiates a new KalturaExtendingItemMrssParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaExtendingItemMrssParameterWithDefaults

`func NewKalturaExtendingItemMrssParameterWithDefaults() *KalturaExtendingItemMrssParameter`

NewKalturaExtendingItemMrssParameterWithDefaults instantiates a new KalturaExtendingItemMrssParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionMode

`func (o *KalturaExtendingItemMrssParameter) GetExtensionMode() int32`

GetExtensionMode returns the ExtensionMode field if non-nil, zero value otherwise.

### GetExtensionModeOk

`func (o *KalturaExtendingItemMrssParameter) GetExtensionModeOk() (*int32, bool)`

GetExtensionModeOk returns a tuple with the ExtensionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionMode

`func (o *KalturaExtendingItemMrssParameter) SetExtensionMode(v int32)`

SetExtensionMode sets ExtensionMode field to given value.

### HasExtensionMode

`func (o *KalturaExtendingItemMrssParameter) HasExtensionMode() bool`

HasExtensionMode returns a boolean if a field has been set.

### GetIdentifier

`func (o *KalturaExtendingItemMrssParameter) GetIdentifier() KalturaObjectIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *KalturaExtendingItemMrssParameter) GetIdentifierOk() (*KalturaObjectIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *KalturaExtendingItemMrssParameter) SetIdentifier(v KalturaObjectIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *KalturaExtendingItemMrssParameter) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetXpath

`func (o *KalturaExtendingItemMrssParameter) GetXpath() string`

GetXpath returns the Xpath field if non-nil, zero value otherwise.

### GetXpathOk

`func (o *KalturaExtendingItemMrssParameter) GetXpathOk() (*string, bool)`

GetXpathOk returns a tuple with the Xpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXpath

`func (o *KalturaExtendingItemMrssParameter) SetXpath(v string)`

SetXpath sets Xpath field to given value.

### HasXpath

`func (o *KalturaExtendingItemMrssParameter) HasXpath() bool`

HasXpath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


