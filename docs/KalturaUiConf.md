# KalturaUiConf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfFile** | Pointer to **string** |  | [optional] 
**ConfFileFeatures** | Pointer to **string** |  | [optional] 
**ConfFilePath** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ConfVars** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**CreationMode** | Pointer to **int32** | Enum Type: &#x60;KalturaUiConfCreationMode&#x60; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Html5Url** | Pointer to **string** |  | [optional] 
**HtmlParams** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the uiConf, this is not a primary key | [optional] 
**ObjType** | Pointer to **int32** | Enum Type: &#x60;KalturaUiConfObjType&#x60; | [optional] 
**ObjTypeAsString** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerTags** | Pointer to **string** |  | [optional] 
**SwfUrl** | Pointer to **string** |  | [optional] 
**SwfUrlVersion** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Entry creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**UseCdn** | Pointer to **bool** |  | [optional] 
**Version** | Pointer to **string** | &#x60;readOnly&#x60;  UiConf version | [optional] [readonly] 
**Width** | Pointer to **int32** |  | [optional] 

## Methods

### NewKalturaUiConf

`func NewKalturaUiConf() *KalturaUiConf`

NewKalturaUiConf instantiates a new KalturaUiConf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaUiConfWithDefaults

`func NewKalturaUiConfWithDefaults() *KalturaUiConf`

NewKalturaUiConfWithDefaults instantiates a new KalturaUiConf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfFile

`func (o *KalturaUiConf) GetConfFile() string`

GetConfFile returns the ConfFile field if non-nil, zero value otherwise.

### GetConfFileOk

`func (o *KalturaUiConf) GetConfFileOk() (*string, bool)`

GetConfFileOk returns a tuple with the ConfFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfFile

`func (o *KalturaUiConf) SetConfFile(v string)`

SetConfFile sets ConfFile field to given value.

### HasConfFile

`func (o *KalturaUiConf) HasConfFile() bool`

HasConfFile returns a boolean if a field has been set.

### GetConfFileFeatures

`func (o *KalturaUiConf) GetConfFileFeatures() string`

GetConfFileFeatures returns the ConfFileFeatures field if non-nil, zero value otherwise.

### GetConfFileFeaturesOk

`func (o *KalturaUiConf) GetConfFileFeaturesOk() (*string, bool)`

GetConfFileFeaturesOk returns a tuple with the ConfFileFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfFileFeatures

`func (o *KalturaUiConf) SetConfFileFeatures(v string)`

SetConfFileFeatures sets ConfFileFeatures field to given value.

### HasConfFileFeatures

`func (o *KalturaUiConf) HasConfFileFeatures() bool`

HasConfFileFeatures returns a boolean if a field has been set.

### GetConfFilePath

`func (o *KalturaUiConf) GetConfFilePath() string`

GetConfFilePath returns the ConfFilePath field if non-nil, zero value otherwise.

### GetConfFilePathOk

`func (o *KalturaUiConf) GetConfFilePathOk() (*string, bool)`

GetConfFilePathOk returns a tuple with the ConfFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfFilePath

`func (o *KalturaUiConf) SetConfFilePath(v string)`

SetConfFilePath sets ConfFilePath field to given value.

### HasConfFilePath

`func (o *KalturaUiConf) HasConfFilePath() bool`

HasConfFilePath returns a boolean if a field has been set.

### GetConfVars

`func (o *KalturaUiConf) GetConfVars() string`

GetConfVars returns the ConfVars field if non-nil, zero value otherwise.

### GetConfVarsOk

`func (o *KalturaUiConf) GetConfVarsOk() (*string, bool)`

GetConfVarsOk returns a tuple with the ConfVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfVars

`func (o *KalturaUiConf) SetConfVars(v string)`

SetConfVars sets ConfVars field to given value.

### HasConfVars

`func (o *KalturaUiConf) HasConfVars() bool`

HasConfVars returns a boolean if a field has been set.

### GetConfig

`func (o *KalturaUiConf) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *KalturaUiConf) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *KalturaUiConf) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *KalturaUiConf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaUiConf) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaUiConf) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaUiConf) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaUiConf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreationMode

`func (o *KalturaUiConf) GetCreationMode() int32`

GetCreationMode returns the CreationMode field if non-nil, zero value otherwise.

### GetCreationModeOk

`func (o *KalturaUiConf) GetCreationModeOk() (*int32, bool)`

GetCreationModeOk returns a tuple with the CreationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationMode

`func (o *KalturaUiConf) SetCreationMode(v int32)`

SetCreationMode sets CreationMode field to given value.

### HasCreationMode

`func (o *KalturaUiConf) HasCreationMode() bool`

HasCreationMode returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaUiConf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaUiConf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaUiConf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaUiConf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHeight

`func (o *KalturaUiConf) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *KalturaUiConf) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *KalturaUiConf) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *KalturaUiConf) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetHtml5Url

`func (o *KalturaUiConf) GetHtml5Url() string`

GetHtml5Url returns the Html5Url field if non-nil, zero value otherwise.

### GetHtml5UrlOk

`func (o *KalturaUiConf) GetHtml5UrlOk() (*string, bool)`

GetHtml5UrlOk returns a tuple with the Html5Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml5Url

`func (o *KalturaUiConf) SetHtml5Url(v string)`

SetHtml5Url sets Html5Url field to given value.

### HasHtml5Url

`func (o *KalturaUiConf) HasHtml5Url() bool`

HasHtml5Url returns a boolean if a field has been set.

### GetHtmlParams

`func (o *KalturaUiConf) GetHtmlParams() string`

GetHtmlParams returns the HtmlParams field if non-nil, zero value otherwise.

### GetHtmlParamsOk

`func (o *KalturaUiConf) GetHtmlParamsOk() (*string, bool)`

GetHtmlParamsOk returns a tuple with the HtmlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlParams

`func (o *KalturaUiConf) SetHtmlParams(v string)`

SetHtmlParams sets HtmlParams field to given value.

### HasHtmlParams

`func (o *KalturaUiConf) HasHtmlParams() bool`

HasHtmlParams returns a boolean if a field has been set.

### GetId

`func (o *KalturaUiConf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaUiConf) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaUiConf) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaUiConf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaUiConf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaUiConf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaUiConf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaUiConf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObjType

`func (o *KalturaUiConf) GetObjType() int32`

GetObjType returns the ObjType field if non-nil, zero value otherwise.

### GetObjTypeOk

`func (o *KalturaUiConf) GetObjTypeOk() (*int32, bool)`

GetObjTypeOk returns a tuple with the ObjType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjType

`func (o *KalturaUiConf) SetObjType(v int32)`

SetObjType sets ObjType field to given value.

### HasObjType

`func (o *KalturaUiConf) HasObjType() bool`

HasObjType returns a boolean if a field has been set.

### GetObjTypeAsString

`func (o *KalturaUiConf) GetObjTypeAsString() string`

GetObjTypeAsString returns the ObjTypeAsString field if non-nil, zero value otherwise.

### GetObjTypeAsStringOk

`func (o *KalturaUiConf) GetObjTypeAsStringOk() (*string, bool)`

GetObjTypeAsStringOk returns a tuple with the ObjTypeAsString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjTypeAsString

`func (o *KalturaUiConf) SetObjTypeAsString(v string)`

SetObjTypeAsString sets ObjTypeAsString field to given value.

### HasObjTypeAsString

`func (o *KalturaUiConf) HasObjTypeAsString() bool`

HasObjTypeAsString returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaUiConf) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaUiConf) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaUiConf) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaUiConf) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerTags

`func (o *KalturaUiConf) GetPartnerTags() string`

GetPartnerTags returns the PartnerTags field if non-nil, zero value otherwise.

### GetPartnerTagsOk

`func (o *KalturaUiConf) GetPartnerTagsOk() (*string, bool)`

GetPartnerTagsOk returns a tuple with the PartnerTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerTags

`func (o *KalturaUiConf) SetPartnerTags(v string)`

SetPartnerTags sets PartnerTags field to given value.

### HasPartnerTags

`func (o *KalturaUiConf) HasPartnerTags() bool`

HasPartnerTags returns a boolean if a field has been set.

### GetSwfUrl

`func (o *KalturaUiConf) GetSwfUrl() string`

GetSwfUrl returns the SwfUrl field if non-nil, zero value otherwise.

### GetSwfUrlOk

`func (o *KalturaUiConf) GetSwfUrlOk() (*string, bool)`

GetSwfUrlOk returns a tuple with the SwfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwfUrl

`func (o *KalturaUiConf) SetSwfUrl(v string)`

SetSwfUrl sets SwfUrl field to given value.

### HasSwfUrl

`func (o *KalturaUiConf) HasSwfUrl() bool`

HasSwfUrl returns a boolean if a field has been set.

### GetSwfUrlVersion

`func (o *KalturaUiConf) GetSwfUrlVersion() string`

GetSwfUrlVersion returns the SwfUrlVersion field if non-nil, zero value otherwise.

### GetSwfUrlVersionOk

`func (o *KalturaUiConf) GetSwfUrlVersionOk() (*string, bool)`

GetSwfUrlVersionOk returns a tuple with the SwfUrlVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwfUrlVersion

`func (o *KalturaUiConf) SetSwfUrlVersion(v string)`

SetSwfUrlVersion sets SwfUrlVersion field to given value.

### HasSwfUrlVersion

`func (o *KalturaUiConf) HasSwfUrlVersion() bool`

HasSwfUrlVersion returns a boolean if a field has been set.

### GetTags

`func (o *KalturaUiConf) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaUiConf) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaUiConf) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaUiConf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaUiConf) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaUiConf) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaUiConf) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaUiConf) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUseCdn

`func (o *KalturaUiConf) GetUseCdn() bool`

GetUseCdn returns the UseCdn field if non-nil, zero value otherwise.

### GetUseCdnOk

`func (o *KalturaUiConf) GetUseCdnOk() (*bool, bool)`

GetUseCdnOk returns a tuple with the UseCdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCdn

`func (o *KalturaUiConf) SetUseCdn(v bool)`

SetUseCdn sets UseCdn field to given value.

### HasUseCdn

`func (o *KalturaUiConf) HasUseCdn() bool`

HasUseCdn returns a boolean if a field has been set.

### GetVersion

`func (o *KalturaUiConf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *KalturaUiConf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *KalturaUiConf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *KalturaUiConf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWidth

`func (o *KalturaUiConf) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *KalturaUiConf) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *KalturaUiConf) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *KalturaUiConf) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


