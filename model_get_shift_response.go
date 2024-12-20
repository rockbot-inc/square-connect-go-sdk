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

// A response to a request to get a `Shift`. The response contains the requested `Shift` object and might contain a set of `Error` objects if the request resulted in errors.
type GetShiftResponse struct {
	Shift *Shift `json:"shift,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
