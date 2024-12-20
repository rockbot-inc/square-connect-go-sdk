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

// Provides metadata when the event `type` is `OTHER`.
type LoyaltyEventOther struct {
	// The Square-assigned ID of the [loyalty program](entity:LoyaltyProgram).
	LoyaltyProgramId string `json:"loyalty_program_id"`
	// The number of points added or removed.
	Points int32 `json:"points"`
}
