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

// Represents an input to a call to [CreateVendor]($e/Vendors/CreateVendor).
type CreateVendorRequest struct {
	// A client-supplied, universally unique identifier (UUID) to make this [CreateVendor](api-endpoint:Vendors-CreateVendor) call idempotent.  See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) in the [API Development 101](https://developer.squareup.com/docs/buildbasics) section for more information.
	IdempotencyKey string `json:"idempotency_key"`
	Vendor *Vendor `json:"vendor,omitempty"`
}