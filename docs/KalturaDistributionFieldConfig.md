# KalturaDistributionFieldConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntryMrssXslt** | Pointer to **string** | An XSLT string that extracts the right value from the Kaltura entry MRSS XML.      The value of the current connector field will be the one that is returned from transforming the Kaltura entry MRSS XML using this XSLT string. | [optional] 
**FieldName** | Pointer to **string** | A value taken from a connector field enum which associates the current configuration to that connector field      Field enum class should be returned by the provider&#39;s getFieldEnumClass function. | [optional] 
**IsDefault** | Pointer to **bool** | &#x60;readOnly&#x60;  Is this field config is the default for the distribution provider? | [optional] [readonly] 
**IsRequired** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionFieldRequiredStatus&#x60;  Is the field required to have a value for submission ? | [optional] 
**TriggerDeleteOnError** | Pointer to **bool** | Is an error on this field going to trigger deletion of distributed content? | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdateOnChange** | Pointer to **bool** | Trigger distribution update when this field changes or not ? | [optional] 
**UpdateParams** | Pointer to [**[]KalturaString**](KalturaString.md) |  | [optional] 
**UserFriendlyFieldName** | Pointer to **string** | A string that will be shown to the user as the field name in error messages related to the current field | [optional] 

## Methods

### NewKalturaDistributionFieldConfig

`func NewKalturaDistributionFieldConfig() *KalturaDistributionFieldConfig`

NewKalturaDistributionFieldConfig instantiates a new KalturaDistributionFieldConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaDistributionFieldConfigWithDefaults

`func NewKalturaDistributionFieldConfigWithDefaults() *KalturaDistributionFieldConfig`

NewKalturaDistributionFieldConfigWithDefaults instantiates a new KalturaDistributionFieldConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntryMrssXslt

`func (o *KalturaDistributionFieldConfig) GetEntryMrssXslt() string`

GetEntryMrssXslt returns the EntryMrssXslt field if non-nil, zero value otherwise.

### GetEntryMrssXsltOk

`func (o *KalturaDistributionFieldConfig) GetEntryMrssXsltOk() (*string, bool)`

GetEntryMrssXsltOk returns a tuple with the EntryMrssXslt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryMrssXslt

`func (o *KalturaDistributionFieldConfig) SetEntryMrssXslt(v string)`

SetEntryMrssXslt sets EntryMrssXslt field to given value.

### HasEntryMrssXslt

`func (o *KalturaDistributionFieldConfig) HasEntryMrssXslt() bool`

HasEntryMrssXslt returns a boolean if a field has been set.

### GetFieldName

`func (o *KalturaDistributionFieldConfig) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *KalturaDistributionFieldConfig) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *KalturaDistributionFieldConfig) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *KalturaDistributionFieldConfig) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetIsDefault

`func (o *KalturaDistributionFieldConfig) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *KalturaDistributionFieldConfig) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *KalturaDistributionFieldConfig) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *KalturaDistributionFieldConfig) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsRequired

`func (o *KalturaDistributionFieldConfig) GetIsRequired() int32`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *KalturaDistributionFieldConfig) GetIsRequiredOk() (*int32, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *KalturaDistributionFieldConfig) SetIsRequired(v int32)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *KalturaDistributionFieldConfig) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetTriggerDeleteOnError

`func (o *KalturaDistributionFieldConfig) GetTriggerDeleteOnError() bool`

GetTriggerDeleteOnError returns the TriggerDeleteOnError field if non-nil, zero value otherwise.

### GetTriggerDeleteOnErrorOk

`func (o *KalturaDistributionFieldConfig) GetTriggerDeleteOnErrorOk() (*bool, bool)`

GetTriggerDeleteOnErrorOk returns a tuple with the TriggerDeleteOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerDeleteOnError

`func (o *KalturaDistributionFieldConfig) SetTriggerDeleteOnError(v bool)`

SetTriggerDeleteOnError sets TriggerDeleteOnError field to given value.

### HasTriggerDeleteOnError

`func (o *KalturaDistributionFieldConfig) HasTriggerDeleteOnError() bool`

HasTriggerDeleteOnError returns a boolean if a field has been set.

### GetType

`func (o *KalturaDistributionFieldConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KalturaDistributionFieldConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KalturaDistributionFieldConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KalturaDistributionFieldConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdateOnChange

`func (o *KalturaDistributionFieldConfig) GetUpdateOnChange() bool`

GetUpdateOnChange returns the UpdateOnChange field if non-nil, zero value otherwise.

### GetUpdateOnChangeOk

`func (o *KalturaDistributionFieldConfig) GetUpdateOnChangeOk() (*bool, bool)`

GetUpdateOnChangeOk returns a tuple with the UpdateOnChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateOnChange

`func (o *KalturaDistributionFieldConfig) SetUpdateOnChange(v bool)`

SetUpdateOnChange sets UpdateOnChange field to given value.

### HasUpdateOnChange

`func (o *KalturaDistributionFieldConfig) HasUpdateOnChange() bool`

HasUpdateOnChange returns a boolean if a field has been set.

### GetUpdateParams

`func (o *KalturaDistributionFieldConfig) GetUpdateParams() []KalturaString`

GetUpdateParams returns the UpdateParams field if non-nil, zero value otherwise.

### GetUpdateParamsOk

`func (o *KalturaDistributionFieldConfig) GetUpdateParamsOk() (*[]KalturaString, bool)`

GetUpdateParamsOk returns a tuple with the UpdateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateParams

`func (o *KalturaDistributionFieldConfig) SetUpdateParams(v []KalturaString)`

SetUpdateParams sets UpdateParams field to given value.

### HasUpdateParams

`func (o *KalturaDistributionFieldConfig) HasUpdateParams() bool`

HasUpdateParams returns a boolean if a field has been set.

### GetUserFriendlyFieldName

`func (o *KalturaDistributionFieldConfig) GetUserFriendlyFieldName() string`

GetUserFriendlyFieldName returns the UserFriendlyFieldName field if non-nil, zero value otherwise.

### GetUserFriendlyFieldNameOk

`func (o *KalturaDistributionFieldConfig) GetUserFriendlyFieldNameOk() (*string, bool)`

GetUserFriendlyFieldNameOk returns a tuple with the UserFriendlyFieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFriendlyFieldName

`func (o *KalturaDistributionFieldConfig) SetUserFriendlyFieldName(v string)`

SetUserFriendlyFieldName sets UserFriendlyFieldName field to given value.

### HasUserFriendlyFieldName

`func (o *KalturaDistributionFieldConfig) HasUserFriendlyFieldName() bool`

HasUserFriendlyFieldName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


