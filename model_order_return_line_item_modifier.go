/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// A line item modifier being returned.
type OrderReturnLineItemModifier struct {
	// A unique ID that identifies the return modifier only within this order.
	Uid string `json:"uid,omitempty"`
	// The modifier `uid` from the order's line item that contains the original sale of this line item modifier.
	SourceModifierUid string `json:"source_modifier_uid,omitempty"`
	// The catalog object ID referencing [CatalogModifier](entity:CatalogModifier).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this line item modifier references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The name of the item modifier.
	Name string `json:"name,omitempty"`
	BasePriceMoney *Money `json:"base_price_money,omitempty"`
	TotalPriceMoney *Money `json:"total_price_money,omitempty"`
	// The quantity of the line item modifier. The modifier quantity can be 0 or more. For example, suppose a restaurant offers a cheeseburger on the menu. When a buyer orders this item, the restaurant records the purchase by creating an `Order` object with a line item for a burger. The line item includes a line item modifier: the name is cheese and the quantity is 1. The buyer has the option to order extra cheese (or no cheese). If the buyer chooses the extra cheese option, the modifier quantity increases to 2. If the buyer does not want any cheese, the modifier quantity is set to 0.
	Quantity string `json:"quantity,omitempty"`
}
