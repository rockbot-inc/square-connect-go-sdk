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

// Represents an individual upsert request in a [BulkUpsertCustomerCustomAttributes]($e/CustomerCustomAttributes/BulkUpsertCustomerCustomAttributes) request. An individual request contains a customer ID, the custom attribute to create or update, and an optional idempotency key.
type BulkUpsertCustomerCustomAttributesRequestCustomerCustomAttributeUpsertRequest struct {
	// The ID of the target [customer profile](entity:Customer).
	CustomerId string `json:"customer_id"`
	CustomAttribute *CustomAttribute `json:"custom_attribute"`
	// A unique identifier for this individual upsert request, used to ensure idempotency. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
