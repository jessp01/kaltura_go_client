# KalturaGenericDistributionProviderAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **int32** | &#x60;insertOnly&#x60;  Enum Type: &#x60;KalturaDistributionAction&#x60; | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Generic distribution provider action creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**EditableFields** | Pointer to **string** |  | [optional] 
**GenericDistributionProviderId** | Pointer to **int32** | &#x60;insertOnly&#x60; | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Auto generated | [optional] [readonly] 
**MandatoryFields** | Pointer to **string** |  | [optional] 
**MrssTransformer** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MrssValidator** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Protocol** | Pointer to **int32** | Enum Type: &#x60;KalturaDistributionProtocol&#x60; | [optional] 
**RemotePassword** | Pointer to **string** |  | [optional] 
**RemotePath** | Pointer to **string** |  | [optional] 
**RemoteUsername** | Pointer to **string** |  | [optional] 
**ResultsParser** | Pointer to **int32** | Enum Type: &#x60;KalturaGenericDistributionProviderParser&#x60; | [optional] 
**ResultsTransformer** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ServerAddress** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaGenericDistributionProviderStatus&#x60; | [optional] [readonly] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Generic distribution provider action last update date as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaGenericDistributionProviderAction

`func NewKalturaGenericDistributionProviderAction() *KalturaGenericDistributionProviderAction`

NewKalturaGenericDistributionProviderAction instantiates a new KalturaGenericDistributionProviderAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaGenericDistributionProviderActionWithDefaults

`func NewKalturaGenericDistributionProviderActionWithDefaults() *KalturaGenericDistributionProviderAction`

NewKalturaGenericDistributionProviderActionWithDefaults instantiates a new KalturaGenericDistributionProviderAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *KalturaGenericDistributionProviderAction) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *KalturaGenericDistributionProviderAction) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *KalturaGenericDistributionProviderAction) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *KalturaGenericDistributionProviderAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaGenericDistributionProviderAction) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaGenericDistributionProviderAction) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaGenericDistributionProviderAction) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaGenericDistributionProviderAction) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEditableFields

`func (o *KalturaGenericDistributionProviderAction) GetEditableFields() string`

GetEditableFields returns the EditableFields field if non-nil, zero value otherwise.

### GetEditableFieldsOk

`func (o *KalturaGenericDistributionProviderAction) GetEditableFieldsOk() (*string, bool)`

GetEditableFieldsOk returns a tuple with the EditableFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditableFields

`func (o *KalturaGenericDistributionProviderAction) SetEditableFields(v string)`

SetEditableFields sets EditableFields field to given value.

### HasEditableFields

`func (o *KalturaGenericDistributionProviderAction) HasEditableFields() bool`

HasEditableFields returns a boolean if a field has been set.

### GetGenericDistributionProviderId

`func (o *KalturaGenericDistributionProviderAction) GetGenericDistributionProviderId() int32`

GetGenericDistributionProviderId returns the GenericDistributionProviderId field if non-nil, zero value otherwise.

### GetGenericDistributionProviderIdOk

`func (o *KalturaGenericDistributionProviderAction) GetGenericDistributionProviderIdOk() (*int32, bool)`

GetGenericDistributionProviderIdOk returns a tuple with the GenericDistributionProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenericDistributionProviderId

`func (o *KalturaGenericDistributionProviderAction) SetGenericDistributionProviderId(v int32)`

SetGenericDistributionProviderId sets GenericDistributionProviderId field to given value.

### HasGenericDistributionProviderId

`func (o *KalturaGenericDistributionProviderAction) HasGenericDistributionProviderId() bool`

HasGenericDistributionProviderId returns a boolean if a field has been set.

### GetId

`func (o *KalturaGenericDistributionProviderAction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaGenericDistributionProviderAction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaGenericDistributionProviderAction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaGenericDistributionProviderAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMandatoryFields

`func (o *KalturaGenericDistributionProviderAction) GetMandatoryFields() string`

GetMandatoryFields returns the MandatoryFields field if non-nil, zero value otherwise.

### GetMandatoryFieldsOk

`func (o *KalturaGenericDistributionProviderAction) GetMandatoryFieldsOk() (*string, bool)`

GetMandatoryFieldsOk returns a tuple with the MandatoryFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatoryFields

`func (o *KalturaGenericDistributionProviderAction) SetMandatoryFields(v string)`

SetMandatoryFields sets MandatoryFields field to given value.

### HasMandatoryFields

`func (o *KalturaGenericDistributionProviderAction) HasMandatoryFields() bool`

HasMandatoryFields returns a boolean if a field has been set.

### GetMrssTransformer

`func (o *KalturaGenericDistributionProviderAction) GetMrssTransformer() string`

GetMrssTransformer returns the MrssTransformer field if non-nil, zero value otherwise.

### GetMrssTransformerOk

`func (o *KalturaGenericDistributionProviderAction) GetMrssTransformerOk() (*string, bool)`

GetMrssTransformerOk returns a tuple with the MrssTransformer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrssTransformer

`func (o *KalturaGenericDistributionProviderAction) SetMrssTransformer(v string)`

SetMrssTransformer sets MrssTransformer field to given value.

### HasMrssTransformer

`func (o *KalturaGenericDistributionProviderAction) HasMrssTransformer() bool`

HasMrssTransformer returns a boolean if a field has been set.

### GetMrssValidator

`func (o *KalturaGenericDistributionProviderAction) GetMrssValidator() string`

GetMrssValidator returns the MrssValidator field if non-nil, zero value otherwise.

### GetMrssValidatorOk

`func (o *KalturaGenericDistributionProviderAction) GetMrssValidatorOk() (*string, bool)`

GetMrssValidatorOk returns a tuple with the MrssValidator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMrssValidator

`func (o *KalturaGenericDistributionProviderAction) SetMrssValidator(v string)`

SetMrssValidator sets MrssValidator field to given value.

### HasMrssValidator

`func (o *KalturaGenericDistributionProviderAction) HasMrssValidator() bool`

HasMrssValidator returns a boolean if a field has been set.

### GetProtocol

`func (o *KalturaGenericDistributionProviderAction) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *KalturaGenericDistributionProviderAction) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *KalturaGenericDistributionProviderAction) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *KalturaGenericDistributionProviderAction) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemotePassword

`func (o *KalturaGenericDistributionProviderAction) GetRemotePassword() string`

GetRemotePassword returns the RemotePassword field if non-nil, zero value otherwise.

### GetRemotePasswordOk

`func (o *KalturaGenericDistributionProviderAction) GetRemotePasswordOk() (*string, bool)`

GetRemotePasswordOk returns a tuple with the RemotePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePassword

`func (o *KalturaGenericDistributionProviderAction) SetRemotePassword(v string)`

SetRemotePassword sets RemotePassword field to given value.

### HasRemotePassword

`func (o *KalturaGenericDistributionProviderAction) HasRemotePassword() bool`

HasRemotePassword returns a boolean if a field has been set.

### GetRemotePath

`func (o *KalturaGenericDistributionProviderAction) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *KalturaGenericDistributionProviderAction) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *KalturaGenericDistributionProviderAction) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *KalturaGenericDistributionProviderAction) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetRemoteUsername

`func (o *KalturaGenericDistributionProviderAction) GetRemoteUsername() string`

GetRemoteUsername returns the RemoteUsername field if non-nil, zero value otherwise.

### GetRemoteUsernameOk

`func (o *KalturaGenericDistributionProviderAction) GetRemoteUsernameOk() (*string, bool)`

GetRemoteUsernameOk returns a tuple with the RemoteUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteUsername

`func (o *KalturaGenericDistributionProviderAction) SetRemoteUsername(v string)`

SetRemoteUsername sets RemoteUsername field to given value.

### HasRemoteUsername

`func (o *KalturaGenericDistributionProviderAction) HasRemoteUsername() bool`

HasRemoteUsername returns a boolean if a field has been set.

### GetResultsParser

`func (o *KalturaGenericDistributionProviderAction) GetResultsParser() int32`

GetResultsParser returns the ResultsParser field if non-nil, zero value otherwise.

### GetResultsParserOk

`func (o *KalturaGenericDistributionProviderAction) GetResultsParserOk() (*int32, bool)`

GetResultsParserOk returns a tuple with the ResultsParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsParser

`func (o *KalturaGenericDistributionProviderAction) SetResultsParser(v int32)`

SetResultsParser sets ResultsParser field to given value.

### HasResultsParser

`func (o *KalturaGenericDistributionProviderAction) HasResultsParser() bool`

HasResultsParser returns a boolean if a field has been set.

### GetResultsTransformer

`func (o *KalturaGenericDistributionProviderAction) GetResultsTransformer() string`

GetResultsTransformer returns the ResultsTransformer field if non-nil, zero value otherwise.

### GetResultsTransformerOk

`func (o *KalturaGenericDistributionProviderAction) GetResultsTransformerOk() (*string, bool)`

GetResultsTransformerOk returns a tuple with the ResultsTransformer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsTransformer

`func (o *KalturaGenericDistributionProviderAction) SetResultsTransformer(v string)`

SetResultsTransformer sets ResultsTransformer field to given value.

### HasResultsTransformer

`func (o *KalturaGenericDistributionProviderAction) HasResultsTransformer() bool`

HasResultsTransformer returns a boolean if a field has been set.

### GetServerAddress

`func (o *KalturaGenericDistributionProviderAction) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *KalturaGenericDistributionProviderAction) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *KalturaGenericDistributionProviderAction) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *KalturaGenericDistributionProviderAction) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaGenericDistributionProviderAction) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaGenericDistributionProviderAction) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaGenericDistributionProviderAction) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaGenericDistributionProviderAction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaGenericDistributionProviderAction) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaGenericDistributionProviderAction) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaGenericDistributionProviderAction) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaGenericDistributionProviderAction) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


