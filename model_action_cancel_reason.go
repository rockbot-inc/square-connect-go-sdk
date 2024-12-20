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

type ActionCancelReason string

// List of ActionCancelReason
const (
	BUYER_CANCELED_ActionCancelReason ActionCancelReason = "BUYER_CANCELED"
	SELLER_CANCELED_ActionCancelReason ActionCancelReason = "SELLER_CANCELED"
	TIMED_OUT_ActionCancelReason ActionCancelReason = "TIMED_OUT"
)
