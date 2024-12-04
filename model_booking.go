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

// Represents a booking as a time-bound service contract for a seller's staff member to provide a specified service at a given location to a requesting customer in one or more appointment segments.
type Booking struct {
	// A unique ID of this object representing a booking.
	Id string `json:"id,omitempty"`
	// The revision number for the booking used for optimistic concurrency.
	Version int32 `json:"version,omitempty"`
	Status *BookingStatus `json:"status,omitempty"`
	// The RFC 3339 timestamp specifying the creation time of this booking.
	CreatedAt string `json:"created_at,omitempty"`
	// The RFC 3339 timestamp specifying the most recent update time of this booking.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The RFC 3339 timestamp specifying the starting time of this booking.
	StartAt string `json:"start_at,omitempty"`
	// The ID of the [Location](entity:Location) object representing the location where the booked service is provided. Once set when the booking is created, its value cannot be changed.
	LocationId string `json:"location_id,omitempty"`
	// The ID of the [Customer](entity:Customer) object representing the customer receiving the booked service.
	CustomerId string `json:"customer_id,omitempty"`
	// The free-text field for the customer to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a relevant [CatalogObject](entity:CatalogObject) instance.
	CustomerNote string `json:"customer_note,omitempty"`
	// The free-text field for the seller to supply notes about the booking. For example, the note can be preferences that cannot be expressed by supported attributes of a specific [CatalogObject](entity:CatalogObject) instance. This field should not be visible to customers.
	SellerNote string `json:"seller_note,omitempty"`
	// A list of appointment segments for this booking.
	AppointmentSegments []AppointmentSegment `json:"appointment_segments,omitempty"`
	// Additional time at the end of a booking. Applications should not make this field visible to customers of a seller.
	TransitionTimeMinutes int32 `json:"transition_time_minutes,omitempty"`
	// Whether the booking is of a full business day.
	AllDay bool `json:"all_day,omitempty"`
	LocationType *BusinessAppointmentSettingsBookingLocationType `json:"location_type,omitempty"`
	CreatorDetails *BookingCreatorDetails `json:"creator_details,omitempty"`
	Source *BookingBookingSource `json:"source,omitempty"`
	Address *Address `json:"address,omitempty"`
}
