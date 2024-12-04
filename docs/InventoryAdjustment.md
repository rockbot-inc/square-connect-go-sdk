# InventoryAdjustment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID generated by Square for the &#x60;InventoryAdjustment&#x60;. | [optional] [default to null]
**ReferenceId** | **string** | An optional ID provided by the application to tie the &#x60;InventoryAdjustment&#x60; to an external system. | [optional] [default to null]
**FromState** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**ToState** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**LocationId** | **string** | The Square-generated ID of the [Location](entity:Location) where the related quantity of items is being tracked. | [optional] [default to null]
**CatalogObjectId** | **string** | The Square-generated ID of the [CatalogObject](entity:CatalogObject) being tracked. | [optional] [default to null]
**CatalogObjectType** | **string** | The [type](entity:CatalogObjectType) of the [CatalogObject](entity:CatalogObject) being tracked.   The Inventory API supports setting and reading the &#x60;\&quot;catalog_object_type\&quot;: \&quot;ITEM_VARIATION\&quot;&#x60; field value.  In addition, it can also read the &#x60;\&quot;catalog_object_type\&quot;: \&quot;ITEM\&quot;&#x60; field value that is set by the Square Restaurants app. | [optional] [default to null]
**Quantity** | **string** | The number of items affected by the adjustment as a decimal string. Can support up to 5 digits after the decimal point. | [optional] [default to null]
**TotalPriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**OccurredAt** | **string** | A client-generated RFC 3339-formatted timestamp that indicates when the inventory adjustment took place. For inventory adjustment updates, the &#x60;occurred_at&#x60; timestamp cannot be older than 24 hours or in the future relative to the time of the request. | [optional] [default to null]
**CreatedAt** | **string** | An RFC 3339-formatted timestamp that indicates when the inventory adjustment is received. | [optional] [default to null]
**Source** | [***SourceApplication**](SourceApplication.md) |  | [optional] [default to null]
**EmployeeId** | **string** | The Square-generated ID of the [Employee](entity:Employee) responsible for the inventory adjustment. | [optional] [default to null]
**TeamMemberId** | **string** | The Square-generated ID of the [Team Member](entity:TeamMember) responsible for the inventory adjustment. | [optional] [default to null]
**TransactionId** | **string** | The Square-generated ID of the [Transaction](entity:Transaction) that caused the adjustment. Only relevant for payment-related state transitions. | [optional] [default to null]
**RefundId** | **string** | The Square-generated ID of the [Refund](entity:Refund) that caused the adjustment. Only relevant for refund-related state transitions. | [optional] [default to null]
**PurchaseOrderId** | **string** | The Square-generated ID of the purchase order that caused the adjustment. Only relevant for state transitions from the Square for Retail app. | [optional] [default to null]
**GoodsReceiptId** | **string** | The Square-generated ID of the goods receipt that caused the adjustment. Only relevant for state transitions from the Square for Retail app. | [optional] [default to null]
**AdjustmentGroup** | [***InventoryAdjustmentGroup**](InventoryAdjustmentGroup.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
