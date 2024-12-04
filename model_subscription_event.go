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

// Describes changes to a subscription and the subscription status.
type SubscriptionEvent struct {
	// The ID of the subscription event.
	Id string `json:"id"`
	SubscriptionEventType *SubscriptionEventSubscriptionEventType `json:"subscription_event_type"`
	// The `YYYY-MM-DD`-formatted date (for example, 2013-01-15) when the subscription event occurred.
	EffectiveDate string `json:"effective_date"`
	// The day-of-the-month the billing anchor date was changed to, if applicable.
	MonthlyBillingAnchorDate int32 `json:"monthly_billing_anchor_date,omitempty"`
	Info *SubscriptionEventInfo `json:"info,omitempty"`
	// A list of Phases, to pass phase-specific information used in the swap.
	Phases []Phase `json:"phases,omitempty"`
	// The ID of the subscription plan variation associated with the subscription.
	PlanVariationId string `json:"plan_variation_id"`
}
