package proof

type Proof interface {
	Generate()
	ValidateBlock() bool
}
