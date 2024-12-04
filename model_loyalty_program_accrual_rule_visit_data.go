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

// Represents additional data for rules with the `VISIT` accrual type.
type LoyaltyProgramAccrualRuleVisitData struct {
	MinimumAmountMoney *Money `json:"minimum_amount_money,omitempty"`
	TaxMode *LoyaltyProgramAccrualRuleTaxMode `json:"tax_mode"`
}