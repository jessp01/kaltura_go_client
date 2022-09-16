# KalturaPartnerUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percent** | Pointer to **float32** | &#x60;readOnly&#x60;  percent of usage out of partner&#39;s package. if usageGB is 5 and package is 10GB, this value will be 50 | [optional] [readonly] 
**HostingGB** | Pointer to **float32** | &#x60;readOnly&#x60;  Partner total hosting in GB on the disk | [optional] [readonly] 
**PackageBW** | Pointer to **int32** | &#x60;readOnly&#x60;  package total BW - actually this is usage, which represents BW+storage | [optional] [readonly] 
**ReachedLimitDate** | Pointer to **int32** | &#x60;readOnly&#x60;  date when partner reached the limit of his package (timestamp) | [optional] [readonly] 
**UsageGB** | Pointer to **float32** | &#x60;readOnly&#x60;  total usage in GB - including bandwidth and storage | [optional] [readonly] 
**UsageGraph** | Pointer to **string** | &#x60;readOnly&#x60;  a semi-colon separated list of comma-separated key-values to represent a usage graph.  keys could be 1-12 for a year view (1,1.2;2,1.1;3,0.9;...;12,1.4;)  keys could be 1-[28,29,30,31] depending on the requested month, for a daily view in a given month (1,0.4;2,0.2;...;31,0.1;) | [optional] [readonly] 

## Methods

### NewKalturaPartnerUsage

`func NewKalturaPartnerUsage() *KalturaPartnerUsage`

NewKalturaPartnerUsage instantiates a new KalturaPartnerUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaPartnerUsageWithDefaults

`func NewKalturaPartnerUsageWithDefaults() *KalturaPartnerUsage`

NewKalturaPartnerUsageWithDefaults instantiates a new KalturaPartnerUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercent

`func (o *KalturaPartnerUsage) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *KalturaPartnerUsage) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *KalturaPartnerUsage) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *KalturaPartnerUsage) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetHostingGB

`func (o *KalturaPartnerUsage) GetHostingGB() float32`

GetHostingGB returns the HostingGB field if non-nil, zero value otherwise.

### GetHostingGBOk

`func (o *KalturaPartnerUsage) GetHostingGBOk() (*float32, bool)`

GetHostingGBOk returns a tuple with the HostingGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostingGB

`func (o *KalturaPartnerUsage) SetHostingGB(v float32)`

SetHostingGB sets HostingGB field to given value.

### HasHostingGB

`func (o *KalturaPartnerUsage) HasHostingGB() bool`

HasHostingGB returns a boolean if a field has been set.

### GetPackageBW

`func (o *KalturaPartnerUsage) GetPackageBW() int32`

GetPackageBW returns the PackageBW field if non-nil, zero value otherwise.

### GetPackageBWOk

`func (o *KalturaPartnerUsage) GetPackageBWOk() (*int32, bool)`

GetPackageBWOk returns a tuple with the PackageBW field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageBW

`func (o *KalturaPartnerUsage) SetPackageBW(v int32)`

SetPackageBW sets PackageBW field to given value.

### HasPackageBW

`func (o *KalturaPartnerUsage) HasPackageBW() bool`

HasPackageBW returns a boolean if a field has been set.

### GetReachedLimitDate

`func (o *KalturaPartnerUsage) GetReachedLimitDate() int32`

GetReachedLimitDate returns the ReachedLimitDate field if non-nil, zero value otherwise.

### GetReachedLimitDateOk

`func (o *KalturaPartnerUsage) GetReachedLimitDateOk() (*int32, bool)`

GetReachedLimitDateOk returns a tuple with the ReachedLimitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachedLimitDate

`func (o *KalturaPartnerUsage) SetReachedLimitDate(v int32)`

SetReachedLimitDate sets ReachedLimitDate field to given value.

### HasReachedLimitDate

`func (o *KalturaPartnerUsage) HasReachedLimitDate() bool`

HasReachedLimitDate returns a boolean if a field has been set.

### GetUsageGB

`func (o *KalturaPartnerUsage) GetUsageGB() float32`

GetUsageGB returns the UsageGB field if non-nil, zero value otherwise.

### GetUsageGBOk

`func (o *KalturaPartnerUsage) GetUsageGBOk() (*float32, bool)`

GetUsageGBOk returns a tuple with the UsageGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageGB

`func (o *KalturaPartnerUsage) SetUsageGB(v float32)`

SetUsageGB sets UsageGB field to given value.

### HasUsageGB

`func (o *KalturaPartnerUsage) HasUsageGB() bool`

HasUsageGB returns a boolean if a field has been set.

### GetUsageGraph

`func (o *KalturaPartnerUsage) GetUsageGraph() string`

GetUsageGraph returns the UsageGraph field if non-nil, zero value otherwise.

### GetUsageGraphOk

`func (o *KalturaPartnerUsage) GetUsageGraphOk() (*string, bool)`

GetUsageGraphOk returns a tuple with the UsageGraph field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageGraph

`func (o *KalturaPartnerUsage) SetUsageGraph(v string)`

SetUsageGraph sets UsageGraph field to given value.

### HasUsageGraph

`func (o *KalturaPartnerUsage) HasUsageGraph() bool`

HasUsageGraph returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


