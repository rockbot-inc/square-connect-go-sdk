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

// Filter events by event type.
type LoyaltyEventTypeFilter struct {
	// The loyalty event types used to filter the result. If multiple values are specified, the endpoint uses a  logical OR to combine them. See [LoyaltyEventType](#type-loyaltyeventtype) for possible values
	Types []LoyaltyEventType `json:"types"`
}
