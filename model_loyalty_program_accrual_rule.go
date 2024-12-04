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

// Represents an accrual rule, which defines how buyers can earn points from the base [loyalty program]($m/LoyaltyProgram).
type LoyaltyProgramAccrualRule struct {
	AccrualType *LoyaltyProgramAccrualRuleType `json:"accrual_type"`
	// The number of points that  buyers earn based on the `accrual_type`.
	Points int32 `json:"points,omitempty"`
	VisitData *LoyaltyProgramAccrualRuleVisitData `json:"visit_data,omitempty"`
	SpendData *LoyaltyProgramAccrualRuleSpendData `json:"spend_data,omitempty"`
	ItemVariationData *LoyaltyProgramAccrualRuleItemVariationData `json:"item_variation_data,omitempty"`
	CategoryData *LoyaltyProgramAccrualRuleCategoryData `json:"category_data,omitempty"`
}
