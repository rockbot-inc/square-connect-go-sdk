/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
// OrderServiceChargeScope : Indicates whether this is a line-item or order-level apportioned service charge.
type OrderServiceChargeScope string

// List of OrderServiceChargeScope
const (
	OTHER_SERVICE_CHARGE_SCOPE_OrderServiceChargeScope OrderServiceChargeScope = "OTHER_SERVICE_CHARGE_SCOPE"
	LINE_ITEM_OrderServiceChargeScope OrderServiceChargeScope = "LINE_ITEM"
	ORDER_OrderServiceChargeScope OrderServiceChargeScope = "ORDER"
)
