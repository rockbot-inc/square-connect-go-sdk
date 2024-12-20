# OrderReturnLineItemModifier

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID that identifies the return modifier only within this order. | [optional] [default to null]
**SourceModifierUid** | **string** | The modifier &#x60;uid&#x60; from the order&#x27;s line item that contains the original sale of this line item modifier. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object ID referencing [CatalogModifier](entity:CatalogModifier). | [optional] [default to null]
**CatalogVersion** | **int64** | The version of the catalog object that this line item modifier references. | [optional] [default to null]
**Name** | **string** | The name of the item modifier. | [optional] [default to null]
**BasePriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalPriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Quantity** | **string** | The quantity of the line item modifier. The modifier quantity can be 0 or more. For example, suppose a restaurant offers a cheeseburger on the menu. When a buyer orders this item, the restaurant records the purchase by creating an &#x60;Order&#x60; object with a line item for a burger. The line item includes a line item modifier: the name is cheese and the quantity is 1. The buyer has the option to order extra cheese (or no cheese). If the buyer chooses the extra cheese option, the modifier quantity increases to 2. If the buyer does not want any cheese, the modifier quantity is set to 0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

