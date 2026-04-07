package pg_err

const (
	// Class 23 — Integrity Constraint Violation
	CodeUniqueViolation     = "23505"
	CodeForeignKeyViolation = "23503"
	CodeNotNullViolation    = "23502"
	CodeCheckViolation      = "23514"
	CodeExclusionViolation  = "23P01"

	// Class 22 — Data Exception
	CodeInvalidTextRepresentation = "22P02" // invalid input syntax (UUID, int, etc)

	// Class 42 — Syntax Error or Access Rule Violation
	CodeSyntaxError       = "42601"
	CodeUndefinedFunction = "42883"

	// Class 08 — Connection Exception
	CodeConnectionException = "08000"

	// Class 40 — Transaction Rollback
	CodeSerializationFailure = "40001" // deadlock, retry
)