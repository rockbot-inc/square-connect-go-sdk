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

// Represents a response for an individual upsert request in a [BulkUpsertMerchantCustomAttributes]($e/MerchantCustomAttributes/BulkUpsertMerchantCustomAttributes) operation.
type BulkUpsertMerchantCustomAttributesResponseMerchantCustomAttributeUpsertResponse struct {
	// The ID of the merchant associated with the custom attribute.
	MerchantId string `json:"merchant_id,omitempty"`
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty"`
	// Any errors that occurred while processing the individual request.
	Errors []ModelError `json:"errors,omitempty"`
}
