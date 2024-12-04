/*
 * Square
 *
 * Use Square APIs to manage and run business including payment, customer, product, inventory, and employee management.
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Represents a [BulkUpsertLocationCustomAttributes]($e/LocationCustomAttributes/BulkUpsertLocationCustomAttributes) response, which contains a map of responses that each corresponds to an individual upsert request.
type BulkUpsertLocationCustomAttributesResponse struct {
	// A map of responses that correspond to individual upsert requests. Each response has the same ID as the corresponding request and contains either a `location_id` and `custom_attribute` or an `errors` field.
	Values map[string]BulkUpsertLocationCustomAttributesResponseLocationCustomAttributeUpsertResponse `json:"values,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}