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

// Filter events by location.
type LoyaltyEventLocationFilter struct {
	// The [location](entity:Location) IDs for loyalty events to query. If multiple values are specified, the endpoint uses  a logical OR to combine them.
	LocationIds []string `json:"location_ids"`
}
