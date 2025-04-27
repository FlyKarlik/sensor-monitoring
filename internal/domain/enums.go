package domain

type AttestationEnum string

const (
	AttestationEnumA AttestationEnum = "A"
	AttestationEnumB AttestationEnum = "B"
	AttestationEnumC AttestationEnum = "C"
)

func (a AttestationEnum) String() string {
	return string(a)
}

func (a AttestationEnum) StringPtr() *string {
	if a == "" {
		return nil
	}
	attestationString := a.String()
	return &attestationString
}
