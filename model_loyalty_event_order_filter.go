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

// Filter events by the order associated with the event.
type LoyaltyEventOrderFilter struct {
	// The ID of the [order](entity:Order) associated with the event.
	OrderId string `json:"order_id"`
}
