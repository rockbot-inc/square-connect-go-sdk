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

// Represents a [DeleteMerchantCustomAttribute]($e/MerchantCustomAttributes/DeleteMerchantCustomAttribute) response. Either an empty object `{}` (for a successful deletion) or `errors` is present in the response.
type DeleteMerchantCustomAttributeResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
