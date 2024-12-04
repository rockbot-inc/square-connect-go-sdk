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
// LoyaltyProgramAccrualRuleTaxMode : Indicates how taxes should be treated when calculating the purchase amount used for loyalty points accrual.  This setting applies only to `SPEND` accrual rules or `VISIT` accrual rules that have a minimum spend requirement.
type LoyaltyProgramAccrualRuleTaxMode string

// List of LoyaltyProgramAccrualRuleTaxMode
const (
	BEFORE_TAX_LoyaltyProgramAccrualRuleTaxMode LoyaltyProgramAccrualRuleTaxMode = "BEFORE_TAX"
	AFTER_TAX_LoyaltyProgramAccrualRuleTaxMode LoyaltyProgramAccrualRuleTaxMode = "AFTER_TAX"
)