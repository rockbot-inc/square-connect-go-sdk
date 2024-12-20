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

// Represents a bulk upsert request for one or more order custom attributes.
type BulkUpsertOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual upsert operations for custom attributes.
	Values map[string]BulkUpsertOrderCustomAttributesRequestUpsertCustomAttribute `json:"values"`
}
