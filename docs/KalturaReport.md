# KalturaReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**Description** | Pointer to **string** | Report description | [optional] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  Report id | [optional] [readonly] 
**Name** | Pointer to **string** | Report name | [optional] 
**PartnerId** | Pointer to **int32** | Partner id associated with the report | [optional] 
**Query** | Pointer to **string** | Report query | [optional] 
**SystemName** | Pointer to **string** | Used to identify system reports in a friendly way | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Last update date as Unix timestamp (In seconds) | [optional] [readonly] 

## Methods

### NewKalturaReport

`func NewKalturaReport() *KalturaReport`

NewKalturaReport instantiates a new KalturaReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaReportWithDefaults

`func NewKalturaReportWithDefaults() *KalturaReport`

NewKalturaReportWithDefaults instantiates a new KalturaReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KalturaReport) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaReport) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaReport) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaReport) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaReport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaReport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaReport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaReport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *KalturaReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaReport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KalturaReport) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaReport) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaReport) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaReport) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaReport) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaReport) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaReport) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaReport) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetQuery

`func (o *KalturaReport) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *KalturaReport) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *KalturaReport) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *KalturaReport) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSystemName

`func (o *KalturaReport) GetSystemName() string`

GetSystemName returns the SystemName field if non-nil, zero value otherwise.

### GetSystemNameOk

`func (o *KalturaReport) GetSystemNameOk() (*string, bool)`

GetSystemNameOk returns a tuple with the SystemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemName

`func (o *KalturaReport) SetSystemName(v string)`

SetSystemName sets SystemName field to given value.

### HasSystemName

`func (o *KalturaReport) HasSystemName() bool`

HasSystemName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaReport) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaReport) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaReport) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaReport) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


