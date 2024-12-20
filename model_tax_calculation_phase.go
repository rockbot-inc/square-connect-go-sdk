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
// TaxCalculationPhase : When to calculate the taxes due on a cart.
type TaxCalculationPhase string

// List of TaxCalculationPhase
const (
	SUBTOTAL_PHASE_TaxCalculationPhase TaxCalculationPhase = "TAX_SUBTOTAL_PHASE"
	TOTAL_PHASE_TaxCalculationPhase TaxCalculationPhase = "TAX_TOTAL_PHASE"
)
