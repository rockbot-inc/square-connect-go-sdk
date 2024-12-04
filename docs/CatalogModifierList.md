# CatalogModifierList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the &#x60;CatalogModifierList&#x60; instance. This is a searchable attribute for use in applicable query filters, and its value length is of  Unicode code points. | [optional] [default to null]
**Ordinal** | **int32** | The position of this &#x60;CatalogModifierList&#x60; within a list of &#x60;CatalogModifierList&#x60; instances. | [optional] [default to null]
**SelectionType** | [***CatalogModifierListSelectionType**](CatalogModifierListSelectionType.md) |  | [optional] [default to null]
**Modifiers** | [**[]CatalogObject**](CatalogObject.md) | A non-empty list of &#x60;CatalogModifier&#x60; objects to be included in the &#x60;CatalogModifierList&#x60;,  for non text-based modifiers when the &#x60;modifier_type&#x60; attribute is &#x60;LIST&#x60;. Each element of this list  is a &#x60;CatalogObject&#x60; instance of the &#x60;MODIFIER&#x60; type, containing the following attributes: &#x60;&#x60;&#x60; { \&quot;id\&quot;: \&quot;{{catalog_modifier_id}}\&quot;, \&quot;type\&quot;: \&quot;MODIFIER\&quot;,  \&quot;modifier_data\&quot;: {{a CatalogModifier instance&gt;}}  } &#x60;&#x60;&#x60; | [optional] [default to null]
**ImageIds** | **[]string** | The IDs of images associated with this &#x60;CatalogModifierList&#x60; instance. Currently these images are not displayed on Square products, but may be displayed in 3rd-party applications. | [optional] [default to null]
**ModifierType** | [***CatalogModifierListModifierType**](CatalogModifierListModifierType.md) |  | [optional] [default to null]
**MaxLength** | **int32** | The maximum length, in Unicode points, of the text string of the text-based modifier as represented by  this &#x60;CatalogModifierList&#x60; object with the &#x60;modifier_type&#x60; set to &#x60;TEXT&#x60;. | [optional] [default to null]
**TextRequired** | **bool** | Whether the text string must be a non-empty string (&#x60;true&#x60;) or not (&#x60;false&#x60;) for a text-based modifier as represented by this &#x60;CatalogModifierList&#x60; object with the &#x60;modifier_type&#x60; set to &#x60;TEXT&#x60;. | [optional] [default to null]
**InternalName** | **string** | A note for internal use by the business.     For example, for a text-based modifier applied to a T-shirt item, if the buyer-supplied text of \&quot;Hello, Kitty!\&quot;   is to be printed on the T-shirt, this &#x60;internal_name&#x60; attribute can be \&quot;Use italic face\&quot; as  an instruction for the business to follow.    For non text-based modifiers, this &#x60;internal_name&#x60; attribute can be  used to include SKUs, internal codes, or supplemental descriptions for internal use. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

