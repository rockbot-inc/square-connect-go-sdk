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

// Defines the parameters for a `CreateDisputeEvidenceText` request.
type CreateDisputeEvidenceTextRequest struct {
	// A unique key identifying the request. For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
	EvidenceType *DisputeEvidenceType `json:"evidence_type,omitempty"`
	// The evidence string.
	EvidenceText string `json:"evidence_text"`
}
