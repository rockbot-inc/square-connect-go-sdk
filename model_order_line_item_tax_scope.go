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
// OrderLineItemTaxScope : Indicates whether this is a line-item or order-level tax.
type OrderLineItemTaxScope string

// List of OrderLineItemTaxScope
const (
	OTHER_TAX_SCOPE_OrderLineItemTaxScope OrderLineItemTaxScope = "OTHER_TAX_SCOPE"
	LINE_ITEM_OrderLineItemTaxScope OrderLineItemTaxScope = "LINE_ITEM"
	ORDER_OrderLineItemTaxScope OrderLineItemTaxScope = "ORDER"
)
