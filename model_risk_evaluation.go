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

// Represents fraud risk information for the associated payment.  When you take a payment through Square's Payments API (using the `CreatePayment` endpoint), Square evaluates it and assigns a risk level to the payment. Sellers can use this information to determine the course of action (for example, provide the goods/services or refund the payment).
type RiskEvaluation struct {
	// The timestamp when payment risk was evaluated, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	RiskLevel *RiskEvaluationRiskLevel `json:"risk_level,omitempty"`
}
