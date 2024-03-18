package rotate

var (
	// OutputFlag for rotate.
	OutputFlag string

	// Admins to be rotated.
	Admins bool

	// Services to be rotated.
	Services bool
)

func isAll() bool {
	return !Admins && !Services
}
