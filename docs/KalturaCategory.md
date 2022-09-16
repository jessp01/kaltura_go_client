# KalturaCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminTags** | Pointer to **string** |  | [optional] 
**AggregationCategories** | Pointer to **string** | List of aggregation channels the category belongs to | [optional] 
**AppearInList** | Pointer to **int32** | Enum Type: &#x60;KalturaAppearInListType&#x60;  If category will be returned for list action. | [optional] 
**ContributionPolicy** | Pointer to **int32** | Enum Type: &#x60;KalturaContributionPolicyType&#x60;  who can assign entries to this category | [optional] 
**CreatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Creation date as Unix timestamp (In seconds) | [optional] [readonly] 
**DefaultOrderBy** | Pointer to **string** | Enum Type: &#x60;KalturaCategoryOrderBy&#x60;  Enable client side applications to define how to sort the category child categories | [optional] 
**DefaultPermissionLevel** | Pointer to **int32** | Enum Type: &#x60;KalturaCategoryUserPermissionLevel&#x60;  Default permissionLevel for new users | [optional] 
**Depth** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**Description** | Pointer to **string** | Category description | [optional] 
**DirectEntriesCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Number of entries that belong to this category directly | [optional] [readonly] 
**DirectSubCategoriesCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Number of direct children categories | [optional] [readonly] 
**EntriesCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Number of entries in this Category (including child categories) | [optional] [readonly] 
**FullIds** | Pointer to **string** | &#x60;readOnly&#x60;  The full ids of the Category | [optional] [readonly] 
**FullName** | Pointer to **string** | &#x60;readOnly&#x60;  The full name of the Category | [optional] [readonly] 
**Id** | Pointer to **int32** | &#x60;readOnly&#x60;  The id of the Category | [optional] [readonly] 
**InheritanceType** | Pointer to **int32** | Enum Type: &#x60;KalturaInheritanceType&#x60;  If Category members are inherited from parent category or set manualy. | [optional] 
**InheritedParentId** | Pointer to **int32** | &#x60;readOnly&#x60;  The category id that this category inherit its members and members permission (for contribution and join) | [optional] [readonly] 
**IsAggregationCategory** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  Flag indicating that the category is an aggregation category | [optional] 
**MembersCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Number of active members for this category | [optional] [readonly] 
**Moderation** | Pointer to **int32** | Enum Type: &#x60;KalturaNullableBoolean&#x60;  Moderation to add entries to this category by users that are not of permission level Manager or Moderator. | [optional] 
**Name** | Pointer to **string** | The name of the Category.   The following characters are not allowed: &#39;&lt;&#39;, &#39;&gt;&#39;, &#39;,&#39; | [optional] 
**Owner** | Pointer to **string** | Category Owner (User id) | [optional] 
**ParentId** | Pointer to **int32** |  | [optional] 
**PartnerData** | Pointer to **string** | Can be used to store various partner related data as a string | [optional] 
**PartnerId** | Pointer to **int32** | &#x60;readOnly&#x60; | [optional] [readonly] 
**PartnerSortValue** | Pointer to **int32** | Can be used to store various partner related data as a numeric value | [optional] 
**PendingEntriesCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Nunber of pending moderation entries | [optional] [readonly] 
**PendingMembersCount** | Pointer to **int32** | &#x60;readOnly&#x60;  Number of pending members for this category | [optional] [readonly] 
**Privacy** | Pointer to **int32** | Enum Type: &#x60;KalturaPrivacyType&#x60;  defines the privacy of the entries that assigned to this category | [optional] 
**PrivacyContext** | Pointer to **string** | Set privacy context for search entries that assiged to private and public categories. the entries will be private if the search context is set with those categories. | [optional] 
**PrivacyContexts** | Pointer to **string** | &#x60;readOnly&#x60;  comma separated parents that defines a privacyContext for search | [optional] [readonly] 
**ReferenceId** | Pointer to **string** | Category external id, controlled and managed by the partner. | [optional] 
**Status** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaCategoryStatus&#x60;  Status | [optional] [readonly] 
**Tags** | Pointer to **string** | Category tags | [optional] 
**UpdatedAt** | Pointer to **int32** | &#x60;readOnly&#x60;  Update date as Unix timestamp (In seconds) | [optional] [readonly] 
**UserJoinPolicy** | Pointer to **int32** | &#x60;readOnly&#x60;  Enum Type: &#x60;KalturaUserJoinPolicyType&#x60;  Who can ask to join this category | [optional] [readonly] 

## Methods

### NewKalturaCategory

`func NewKalturaCategory() *KalturaCategory`

NewKalturaCategory instantiates a new KalturaCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKalturaCategoryWithDefaults

`func NewKalturaCategoryWithDefaults() *KalturaCategory`

NewKalturaCategoryWithDefaults instantiates a new KalturaCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminTags

`func (o *KalturaCategory) GetAdminTags() string`

GetAdminTags returns the AdminTags field if non-nil, zero value otherwise.

### GetAdminTagsOk

`func (o *KalturaCategory) GetAdminTagsOk() (*string, bool)`

GetAdminTagsOk returns a tuple with the AdminTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminTags

`func (o *KalturaCategory) SetAdminTags(v string)`

SetAdminTags sets AdminTags field to given value.

### HasAdminTags

`func (o *KalturaCategory) HasAdminTags() bool`

HasAdminTags returns a boolean if a field has been set.

### GetAggregationCategories

`func (o *KalturaCategory) GetAggregationCategories() string`

GetAggregationCategories returns the AggregationCategories field if non-nil, zero value otherwise.

### GetAggregationCategoriesOk

`func (o *KalturaCategory) GetAggregationCategoriesOk() (*string, bool)`

GetAggregationCategoriesOk returns a tuple with the AggregationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationCategories

`func (o *KalturaCategory) SetAggregationCategories(v string)`

SetAggregationCategories sets AggregationCategories field to given value.

### HasAggregationCategories

`func (o *KalturaCategory) HasAggregationCategories() bool`

HasAggregationCategories returns a boolean if a field has been set.

### GetAppearInList

`func (o *KalturaCategory) GetAppearInList() int32`

GetAppearInList returns the AppearInList field if non-nil, zero value otherwise.

### GetAppearInListOk

`func (o *KalturaCategory) GetAppearInListOk() (*int32, bool)`

GetAppearInListOk returns a tuple with the AppearInList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearInList

`func (o *KalturaCategory) SetAppearInList(v int32)`

SetAppearInList sets AppearInList field to given value.

### HasAppearInList

`func (o *KalturaCategory) HasAppearInList() bool`

HasAppearInList returns a boolean if a field has been set.

### GetContributionPolicy

`func (o *KalturaCategory) GetContributionPolicy() int32`

GetContributionPolicy returns the ContributionPolicy field if non-nil, zero value otherwise.

### GetContributionPolicyOk

`func (o *KalturaCategory) GetContributionPolicyOk() (*int32, bool)`

GetContributionPolicyOk returns a tuple with the ContributionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributionPolicy

`func (o *KalturaCategory) SetContributionPolicy(v int32)`

SetContributionPolicy sets ContributionPolicy field to given value.

### HasContributionPolicy

`func (o *KalturaCategory) HasContributionPolicy() bool`

HasContributionPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *KalturaCategory) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KalturaCategory) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KalturaCategory) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KalturaCategory) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefaultOrderBy

`func (o *KalturaCategory) GetDefaultOrderBy() string`

GetDefaultOrderBy returns the DefaultOrderBy field if non-nil, zero value otherwise.

### GetDefaultOrderByOk

`func (o *KalturaCategory) GetDefaultOrderByOk() (*string, bool)`

GetDefaultOrderByOk returns a tuple with the DefaultOrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOrderBy

`func (o *KalturaCategory) SetDefaultOrderBy(v string)`

SetDefaultOrderBy sets DefaultOrderBy field to given value.

### HasDefaultOrderBy

`func (o *KalturaCategory) HasDefaultOrderBy() bool`

HasDefaultOrderBy returns a boolean if a field has been set.

### GetDefaultPermissionLevel

`func (o *KalturaCategory) GetDefaultPermissionLevel() int32`

GetDefaultPermissionLevel returns the DefaultPermissionLevel field if non-nil, zero value otherwise.

### GetDefaultPermissionLevelOk

`func (o *KalturaCategory) GetDefaultPermissionLevelOk() (*int32, bool)`

GetDefaultPermissionLevelOk returns a tuple with the DefaultPermissionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPermissionLevel

`func (o *KalturaCategory) SetDefaultPermissionLevel(v int32)`

SetDefaultPermissionLevel sets DefaultPermissionLevel field to given value.

### HasDefaultPermissionLevel

`func (o *KalturaCategory) HasDefaultPermissionLevel() bool`

HasDefaultPermissionLevel returns a boolean if a field has been set.

### GetDepth

`func (o *KalturaCategory) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *KalturaCategory) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *KalturaCategory) SetDepth(v int32)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *KalturaCategory) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetDescription

`func (o *KalturaCategory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KalturaCategory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KalturaCategory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KalturaCategory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirectEntriesCount

`func (o *KalturaCategory) GetDirectEntriesCount() int32`

GetDirectEntriesCount returns the DirectEntriesCount field if non-nil, zero value otherwise.

### GetDirectEntriesCountOk

`func (o *KalturaCategory) GetDirectEntriesCountOk() (*int32, bool)`

GetDirectEntriesCountOk returns a tuple with the DirectEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectEntriesCount

`func (o *KalturaCategory) SetDirectEntriesCount(v int32)`

SetDirectEntriesCount sets DirectEntriesCount field to given value.

### HasDirectEntriesCount

`func (o *KalturaCategory) HasDirectEntriesCount() bool`

HasDirectEntriesCount returns a boolean if a field has been set.

### GetDirectSubCategoriesCount

`func (o *KalturaCategory) GetDirectSubCategoriesCount() int32`

GetDirectSubCategoriesCount returns the DirectSubCategoriesCount field if non-nil, zero value otherwise.

### GetDirectSubCategoriesCountOk

`func (o *KalturaCategory) GetDirectSubCategoriesCountOk() (*int32, bool)`

GetDirectSubCategoriesCountOk returns a tuple with the DirectSubCategoriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectSubCategoriesCount

`func (o *KalturaCategory) SetDirectSubCategoriesCount(v int32)`

SetDirectSubCategoriesCount sets DirectSubCategoriesCount field to given value.

### HasDirectSubCategoriesCount

`func (o *KalturaCategory) HasDirectSubCategoriesCount() bool`

HasDirectSubCategoriesCount returns a boolean if a field has been set.

### GetEntriesCount

`func (o *KalturaCategory) GetEntriesCount() int32`

GetEntriesCount returns the EntriesCount field if non-nil, zero value otherwise.

### GetEntriesCountOk

`func (o *KalturaCategory) GetEntriesCountOk() (*int32, bool)`

GetEntriesCountOk returns a tuple with the EntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntriesCount

`func (o *KalturaCategory) SetEntriesCount(v int32)`

SetEntriesCount sets EntriesCount field to given value.

### HasEntriesCount

`func (o *KalturaCategory) HasEntriesCount() bool`

HasEntriesCount returns a boolean if a field has been set.

### GetFullIds

`func (o *KalturaCategory) GetFullIds() string`

GetFullIds returns the FullIds field if non-nil, zero value otherwise.

### GetFullIdsOk

`func (o *KalturaCategory) GetFullIdsOk() (*string, bool)`

GetFullIdsOk returns a tuple with the FullIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullIds

`func (o *KalturaCategory) SetFullIds(v string)`

SetFullIds sets FullIds field to given value.

### HasFullIds

`func (o *KalturaCategory) HasFullIds() bool`

HasFullIds returns a boolean if a field has been set.

### GetFullName

`func (o *KalturaCategory) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *KalturaCategory) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *KalturaCategory) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *KalturaCategory) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetId

`func (o *KalturaCategory) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KalturaCategory) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KalturaCategory) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KalturaCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceType

`func (o *KalturaCategory) GetInheritanceType() int32`

GetInheritanceType returns the InheritanceType field if non-nil, zero value otherwise.

### GetInheritanceTypeOk

`func (o *KalturaCategory) GetInheritanceTypeOk() (*int32, bool)`

GetInheritanceTypeOk returns a tuple with the InheritanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceType

`func (o *KalturaCategory) SetInheritanceType(v int32)`

SetInheritanceType sets InheritanceType field to given value.

### HasInheritanceType

`func (o *KalturaCategory) HasInheritanceType() bool`

HasInheritanceType returns a boolean if a field has been set.

### GetInheritedParentId

`func (o *KalturaCategory) GetInheritedParentId() int32`

GetInheritedParentId returns the InheritedParentId field if non-nil, zero value otherwise.

### GetInheritedParentIdOk

`func (o *KalturaCategory) GetInheritedParentIdOk() (*int32, bool)`

GetInheritedParentIdOk returns a tuple with the InheritedParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedParentId

`func (o *KalturaCategory) SetInheritedParentId(v int32)`

SetInheritedParentId sets InheritedParentId field to given value.

### HasInheritedParentId

`func (o *KalturaCategory) HasInheritedParentId() bool`

HasInheritedParentId returns a boolean if a field has been set.

### GetIsAggregationCategory

`func (o *KalturaCategory) GetIsAggregationCategory() int32`

GetIsAggregationCategory returns the IsAggregationCategory field if non-nil, zero value otherwise.

### GetIsAggregationCategoryOk

`func (o *KalturaCategory) GetIsAggregationCategoryOk() (*int32, bool)`

GetIsAggregationCategoryOk returns a tuple with the IsAggregationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAggregationCategory

`func (o *KalturaCategory) SetIsAggregationCategory(v int32)`

SetIsAggregationCategory sets IsAggregationCategory field to given value.

### HasIsAggregationCategory

`func (o *KalturaCategory) HasIsAggregationCategory() bool`

HasIsAggregationCategory returns a boolean if a field has been set.

### GetMembersCount

`func (o *KalturaCategory) GetMembersCount() int32`

GetMembersCount returns the MembersCount field if non-nil, zero value otherwise.

### GetMembersCountOk

`func (o *KalturaCategory) GetMembersCountOk() (*int32, bool)`

GetMembersCountOk returns a tuple with the MembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersCount

`func (o *KalturaCategory) SetMembersCount(v int32)`

SetMembersCount sets MembersCount field to given value.

### HasMembersCount

`func (o *KalturaCategory) HasMembersCount() bool`

HasMembersCount returns a boolean if a field has been set.

### GetModeration

`func (o *KalturaCategory) GetModeration() int32`

GetModeration returns the Moderation field if non-nil, zero value otherwise.

### GetModerationOk

`func (o *KalturaCategory) GetModerationOk() (*int32, bool)`

GetModerationOk returns a tuple with the Moderation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeration

`func (o *KalturaCategory) SetModeration(v int32)`

SetModeration sets Moderation field to given value.

### HasModeration

`func (o *KalturaCategory) HasModeration() bool`

HasModeration returns a boolean if a field has been set.

### GetName

`func (o *KalturaCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KalturaCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KalturaCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KalturaCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *KalturaCategory) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *KalturaCategory) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *KalturaCategory) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *KalturaCategory) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetParentId

`func (o *KalturaCategory) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *KalturaCategory) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *KalturaCategory) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *KalturaCategory) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPartnerData

`func (o *KalturaCategory) GetPartnerData() string`

GetPartnerData returns the PartnerData field if non-nil, zero value otherwise.

### GetPartnerDataOk

`func (o *KalturaCategory) GetPartnerDataOk() (*string, bool)`

GetPartnerDataOk returns a tuple with the PartnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerData

`func (o *KalturaCategory) SetPartnerData(v string)`

SetPartnerData sets PartnerData field to given value.

### HasPartnerData

`func (o *KalturaCategory) HasPartnerData() bool`

HasPartnerData returns a boolean if a field has been set.

### GetPartnerId

`func (o *KalturaCategory) GetPartnerId() int32`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *KalturaCategory) GetPartnerIdOk() (*int32, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *KalturaCategory) SetPartnerId(v int32)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *KalturaCategory) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetPartnerSortValue

`func (o *KalturaCategory) GetPartnerSortValue() int32`

GetPartnerSortValue returns the PartnerSortValue field if non-nil, zero value otherwise.

### GetPartnerSortValueOk

`func (o *KalturaCategory) GetPartnerSortValueOk() (*int32, bool)`

GetPartnerSortValueOk returns a tuple with the PartnerSortValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerSortValue

`func (o *KalturaCategory) SetPartnerSortValue(v int32)`

SetPartnerSortValue sets PartnerSortValue field to given value.

### HasPartnerSortValue

`func (o *KalturaCategory) HasPartnerSortValue() bool`

HasPartnerSortValue returns a boolean if a field has been set.

### GetPendingEntriesCount

`func (o *KalturaCategory) GetPendingEntriesCount() int32`

GetPendingEntriesCount returns the PendingEntriesCount field if non-nil, zero value otherwise.

### GetPendingEntriesCountOk

`func (o *KalturaCategory) GetPendingEntriesCountOk() (*int32, bool)`

GetPendingEntriesCountOk returns a tuple with the PendingEntriesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingEntriesCount

`func (o *KalturaCategory) SetPendingEntriesCount(v int32)`

SetPendingEntriesCount sets PendingEntriesCount field to given value.

### HasPendingEntriesCount

`func (o *KalturaCategory) HasPendingEntriesCount() bool`

HasPendingEntriesCount returns a boolean if a field has been set.

### GetPendingMembersCount

`func (o *KalturaCategory) GetPendingMembersCount() int32`

GetPendingMembersCount returns the PendingMembersCount field if non-nil, zero value otherwise.

### GetPendingMembersCountOk

`func (o *KalturaCategory) GetPendingMembersCountOk() (*int32, bool)`

GetPendingMembersCountOk returns a tuple with the PendingMembersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingMembersCount

`func (o *KalturaCategory) SetPendingMembersCount(v int32)`

SetPendingMembersCount sets PendingMembersCount field to given value.

### HasPendingMembersCount

`func (o *KalturaCategory) HasPendingMembersCount() bool`

HasPendingMembersCount returns a boolean if a field has been set.

### GetPrivacy

`func (o *KalturaCategory) GetPrivacy() int32`

GetPrivacy returns the Privacy field if non-nil, zero value otherwise.

### GetPrivacyOk

`func (o *KalturaCategory) GetPrivacyOk() (*int32, bool)`

GetPrivacyOk returns a tuple with the Privacy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacy

`func (o *KalturaCategory) SetPrivacy(v int32)`

SetPrivacy sets Privacy field to given value.

### HasPrivacy

`func (o *KalturaCategory) HasPrivacy() bool`

HasPrivacy returns a boolean if a field has been set.

### GetPrivacyContext

`func (o *KalturaCategory) GetPrivacyContext() string`

GetPrivacyContext returns the PrivacyContext field if non-nil, zero value otherwise.

### GetPrivacyContextOk

`func (o *KalturaCategory) GetPrivacyContextOk() (*string, bool)`

GetPrivacyContextOk returns a tuple with the PrivacyContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyContext

`func (o *KalturaCategory) SetPrivacyContext(v string)`

SetPrivacyContext sets PrivacyContext field to given value.

### HasPrivacyContext

`func (o *KalturaCategory) HasPrivacyContext() bool`

HasPrivacyContext returns a boolean if a field has been set.

### GetPrivacyContexts

`func (o *KalturaCategory) GetPrivacyContexts() string`

GetPrivacyContexts returns the PrivacyContexts field if non-nil, zero value otherwise.

### GetPrivacyContextsOk

`func (o *KalturaCategory) GetPrivacyContextsOk() (*string, bool)`

GetPrivacyContextsOk returns a tuple with the PrivacyContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyContexts

`func (o *KalturaCategory) SetPrivacyContexts(v string)`

SetPrivacyContexts sets PrivacyContexts field to given value.

### HasPrivacyContexts

`func (o *KalturaCategory) HasPrivacyContexts() bool`

HasPrivacyContexts returns a boolean if a field has been set.

### GetReferenceId

`func (o *KalturaCategory) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *KalturaCategory) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *KalturaCategory) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *KalturaCategory) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetStatus

`func (o *KalturaCategory) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *KalturaCategory) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *KalturaCategory) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *KalturaCategory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *KalturaCategory) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *KalturaCategory) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *KalturaCategory) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *KalturaCategory) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *KalturaCategory) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *KalturaCategory) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *KalturaCategory) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *KalturaCategory) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserJoinPolicy

`func (o *KalturaCategory) GetUserJoinPolicy() int32`

GetUserJoinPolicy returns the UserJoinPolicy field if non-nil, zero value otherwise.

### GetUserJoinPolicyOk

`func (o *KalturaCategory) GetUserJoinPolicyOk() (*int32, bool)`

GetUserJoinPolicyOk returns a tuple with the UserJoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserJoinPolicy

`func (o *KalturaCategory) SetUserJoinPolicy(v int32)`

SetUserJoinPolicy sets UserJoinPolicy field to given value.

### HasUserJoinPolicy

`func (o *KalturaCategory) HasUserJoinPolicy() bool`

HasUserJoinPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


