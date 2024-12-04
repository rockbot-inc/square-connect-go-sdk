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

// Represents a response from deleting one or more order custom attributes.
type BulkDeleteOrderCustomAttributesResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	//  A map of responses that correspond to individual delete requests. Each response has the same ID  as the corresponding request and contains either a `custom_attribute` or an `errors` field.
	Values map[string]DeleteOrderCustomAttributeResponse `json:"values"`
}
