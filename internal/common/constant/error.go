package constant

const (
	// Error for character range

	ErrMustBeFilled = "Must be filled"
	ErrMax100Chars  = "Must be 100 characters or less"
	ErrMax500Chars  = "Must be 500 characters or less"

	// Error for number range

	ErrMin1 = "Must be 1 or more"

	// Error for date range

	ErrNotZeroDate  = "Must not be zero date"
	ErrNotBeforeNow = "Must not be before now"

	// Error for invalid domain type

	ErrInvalidSkytraxType     = "Must be 'Full Service' or 'Low Cost'"
	ErrInvalidRating          = "Must be between 1 and 5"
	ErrInvalidPackageCategory = "Must be 'Silver', 'Gold', 'Platinum', or 'Luxury'"
	ErrInvalidPackageType     = "Must be 'Reguler' or 'Plus Wisata'"
	ErrInvalidAirportCode     = "Must be 3 character"
)
