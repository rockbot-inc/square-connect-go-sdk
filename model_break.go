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

// A record of an employee's break during a shift.
type ModelBreak struct {
	// The UUID for this object.
	Id string `json:"id,omitempty"`
	// RFC 3339; follows the same timezone information as `Shift`. Precision up to the minute is respected; seconds are truncated.
	StartAt string `json:"start_at"`
	// RFC 3339; follows the same timezone information as `Shift`. Precision up to the minute is respected; seconds are truncated.
	EndAt string `json:"end_at,omitempty"`
	// The `BreakType` that this `Break` was templated on.
	BreakTypeId string `json:"break_type_id"`
	// A human-readable name.
	Name string `json:"name"`
	// Format: RFC-3339 P[n]Y[n]M[n]DT[n]H[n]M[n]S. The expected length of the break.
	ExpectedDuration string `json:"expected_duration"`
	// Whether this break counts towards time worked for compensation purposes.
	IsPaid bool `json:"is_paid"`
}
