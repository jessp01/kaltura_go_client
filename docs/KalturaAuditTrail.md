# KalturaAuditTrail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Enum Type: &#x60;KalturaAuditTrailAction&#x60; | [optional] 
**AuditObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaAuditTrailObjectType&#x60; | [optional] 
**ClientTag** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaAuditTrailContext&#x60; | [optional] [readonly] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Data** | Pointer to [**KalturaAuditTrailInfo**](KalturaAuditTrailInfo.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EntryId** | Pointer to **string** |  | [optional] 
**EntryPoint** | Pointer to **string** | &#x60;readOnly&#x60;  The API service and action that called and caused this audit | [optional] [readonly] 
**ErrorDescription** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**IpAddress** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Ks** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**MasterPartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ObjectId** | Pointer to **string** |  | [optional] 
**ParsedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Indicates when the data was parsed | [optional] [readonly] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**RelatedObjectId** | Pointer to **string** |  | [optional] 
**RelatedObjectType** | Pointer to **string** | Enum Type: &#x60;KalturaAuditTrailObjectType&#x60; | [optional] 
**RequestId** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**ServerName** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaAuditTrailStatus&#x60; | [optional] [readonly] 
**UserAgent** | Pointer to **string** | &#x60;readOnly&#x60; | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewKalturaAuditTrail

`func NewKalturaAuditTrail() *KalturaAuditTrail`

NewKalturaAuditTrail instantiates a new KalturaAuditTrail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaAuditTrailWithDefaults

`func NewKalturaAuditTrailWithDefaults() *KalturaAuditTrail`

NewKalturaAuditTrailWithDefaults instantiates a new KalturaAuditTrail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *KalturaAuditTrail) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *KalturaAuditTrail) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *KalturaAuditTrail) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *KalturaAuditTrail) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAuditObjectType

`func (o *KalturaAuditTrail) GetAuditObjectType() string`

GetAuditObjectType returns the AuditObjectType field if non-nil, zero value otherwise.

### GetAuditObjectTypeOk

`func (o *KalturaAuditTrail) GetAuditObjectTypeOk() (*string, bool)`

GetAuditObjectTypeOk returns a tuple with the AuditObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditObjectType

`func (o *KalturaAuditTrail) SetAuditObjectType(v string)`

SetAuditObjectType sets AuditObjectType field to given value.

### HasAuditObjectType

`func (o *KalturaAuditTrail) HasAuditObjectType() bool`

HasAuditObjectType returns a boolean if a field has been set.

### GetClientTag

`func (o *KalturaAuditTrail) GetClientTag() string`

GetClientTag returns the ClientTag field if non-nil, zero value otherwise.

### GetClientTagOk

`func (o *KalturaAuditTrail) GetClientTagOk() (*string, bool)`

GetClientTagOk returns a tuple with the ClientTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTag

`func (o *KalturaAuditTrail) SetClientTag(v string)`

SetClientTag sets ClientTag field to given value.

### HasClientTag

`func (o *KalturaAuditTrail) HasClientTag() bool`

HasClientTag returns a boolean if a field has been set.

### GetContext

`func (o *KalturaAuditTrail) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *KalturaAuditTrail) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *KalturaAuditTrail) SetContext(v int32)`

SetContext sets Context field to given value.

### HasContext

`func (o *KalturaAuditTrail) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaAuditTrail) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaAuditTrail) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaAuditTrail) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaAuditTrail) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *KalturaAuditTrail) GetData() KalturaAuditTrailInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KalturaAuditTrail) GetDataOk() (*KalturaAuditTrailInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KalturaAuditTrail) SetData(v KalturaAuditTrailInfo)`

SetData sets Data field to given value.

### HasData

`func (o *KalturaAuditTrail) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaAuditTrail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaAuditTrail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaAuditTrail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaAuditTrail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntryId

`func (o *KalturaAuditTrail) GetEntryId() string`

GetEntryId returns the EntryId field if non-nil, zero value otherwise.

### GetEntryIdOk

`func (o *KalturaAuditTrail) GetEntryIdOk() (*string, bool)`

GetEntryIdOk returns a tuple with the EntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryId

`func (o *KalturaAuditTrail) SetEntryId(v string)`

SetEntryId sets EntryId field to given value.

### HasEntryId

`func (o *KalturaAuditTrail) HasEntryId() bool`

HasEntryId returns a boolean if a field has been set.

### GetEntryPoint

`func (o *KalturaAuditTrail) GetEntryPoint() string`

GetEntryPoint returns the EntryPoint field if non-nil, zero value otherwise.

### GetEntryPointOk

`func (o *KalturaAuditTrail) GetEntryPointOk() (*string, bool)`

GetEntryPointOk returns a tuple with the EntryPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoint

`func (o *KalturaAuditTrail) SetEntryPoint(v string)`

SetEntryPoint sets EntryPoint field to given value.

### HasEntryPoint

`func (o *KalturaAuditTrail) HasEntryPoint() bool`

HasEntryPoint returns a boolean if a field has been set.

### GetErrorDescription

`func (o *KalturaAuditTrail) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *KalturaAuditTrail) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *KalturaAuditTrail) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *KalturaAuditTrail) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaAuditTrail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaAuditTrail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaAuditTrail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaAuditTrail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *KalturaAuditTrail) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *KalturaAuditTrail) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *KalturaAuditTrail) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *KalturaAuditTrail) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetKs

`func (o *KalturaAuditTrail) GetKs() string`

GetKs returns the Ks field if non-nil, zero value otherwise.

### GetKsOk

`func (o *KalturaAuditTrail) GetKsOk() (*string, bool)`

GetKsOk returns a tuple with the Ks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKs

`func (o *KalturaAuditTrail) SetKs(v string)`

SetKs sets Ks field to given value.

### HasKs

`func (o *KalturaAuditTrail) HasKs() bool`

HasKs returns a boolean if a field has been set.

### GetMasterPartnerId

`func (o *KalturaAuditTrail) GetMasterPartnerId() int32`

GetMasterPartnerId returns the MasterPartnerId field if non-nil, zero value otherwise.

### GetMasterPartnerIdOk

`func (o *KalturaAuditTrail) GetMasterPartnerIdOk() (*int32, bool)`

GetMasterPartnerIdOk returns a tuple with the MasterPartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterPartnerId

`func (o *KalturaAuditTrail) SetMasterPartnerId(v int32)`

SetMasterPartnerId sets MasterPartnerId field to given value.

### HasMasterPartnerId

`func (o *KalturaAuditTrail) HasMasterPartnerId() bool`

HasMasterPartnerId returns a boolean if a field has been set.

### GetObjectId

`func (o *KalturaAuditTrail) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *KalturaAuditTrail) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *KalturaAuditTrail) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *KalturaAuditTrail) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetParsedAt

`func (o *KalturaAuditTrail) GetParsedAt() int32`

GetParsedAt returns the ParsedAt field if non-nil, zero value otherwise.

### GetParsedAtOk

`func (o *KalturaAuditTrail) GetParsedAtOk() (*int32, bool)`

GetParsedAtOk returns a tuple with the ParsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsedAt

`func (o *KalturaAuditTrail) SetParsedAt(v int32)`

SetParsedAt sets ParsedAt field to given value.

### HasParsedAt

`func (o *KalturaAuditTrail) HasParsedAt() bool`

HasParsedAt returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaAuditTrail) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaAuditTrail) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaAuditTrail) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaAuditTrail) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetRelatedObjectId

`func (o *KalturaAuditTrail) GetRelatedObjectId() string`

GetRelatedObjectId returns the RelatedObjectId field if non-nil, zero value otherwise.

### GetRelatedObjectIdOk

`func (o *KalturaAuditTrail) GetRelatedObjectIdOk() (*string, bool)`

GetRelatedObjectIdOk returns a tuple with the RelatedObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectId

`func (o *KalturaAuditTrail) SetRelatedObjectId(v string)`

SetRelatedObjectId sets RelatedObjectId field to given value.

### HasRelatedObjectId

`func (o *KalturaAuditTrail) HasRelatedObjectId() bool`

HasRelatedObjectId returns a boolean if a field has been set.

### GetRelatedObjectType

`func (o *KalturaAuditTrail) GetRelatedObjectType() string`

GetRelatedObjectType returns the RelatedObjectType field if non-nil, zero value otherwise.

### GetRelatedObjectTypeOk

`func (o *KalturaAuditTrail) GetRelatedObjectTypeOk() (*string, bool)`

GetRelatedObjectTypeOk returns a tuple with the RelatedObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedObjectType

`func (o *KalturaAuditTrail) SetRelatedObjectType(v string)`

SetRelatedObjectType sets RelatedObjectType field to given value.

### HasRelatedObjectType

`func (o *KalturaAuditTrail) HasRelatedObjectType() bool`

HasRelatedObjectType returns a boolean if a field has been set.

### GetRequestId

`func (o *KalturaAuditTrail) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *KalturaAuditTrail) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *KalturaAuditTrail) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *KalturaAuditTrail) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetServerName

`func (o *KalturaAuditTrail) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *KalturaAuditTrail) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *KalturaAuditTrail) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *KalturaAuditTrail) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaAuditTrail) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaAuditTrail) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaAuditTrail) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaAuditTrail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserAgent

`func (o *KalturaAuditTrail) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *KalturaAuditTrail) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *KalturaAuditTrail) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *KalturaAuditTrail) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserId

`func (o *KalturaAuditTrail) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *KalturaAuditTrail) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *KalturaAuditTrail) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *KalturaAuditTrail) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


