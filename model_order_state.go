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
// OrderState : The state of the order.
type OrderState string

// List of OrderState
const (
	OPEN_OrderState OrderState = "OPEN"
	COMPLETED_OrderState OrderState = "COMPLETED"
	CANCELED_OrderState OrderState = "CANCELED"
	DRAFT_OrderState OrderState = "DRAFT"
)
