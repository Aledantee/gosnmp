package smi

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ObjectIdentifier is an ASN.1 object identifier.
type ObjectIdentifier []int

// ParseObjectIdentifier parses a string representation of an object identifier.
func ParseObjectIdentifier(s string) (ObjectIdentifier, error) {
	if s == "" {
		return nil, nil
	}

	var (
		oid ObjectIdentifier
		sid int
	)
	for i, c := range s {
		switch {
		case c == '.':
			if i > 0 && s[i-1] == '.' {
				return nil, errors.New("cannot have consecutive periods")
			}

			if i == len(s)-1 {
				return nil, errors.New("cannot end with a period")
			}

			if i > 0 {
				oid = append(oid, sid)
				sid = 0
			}
		case c >= '0' && c <= '9':
			sid = sid*10 + int(c-'0')
		default:
			return nil, fmt.Errorf("invalid character at position %d: %c", i, c)
		}
	}

	oid = append(oid, sid)
	return oid, oid.Validate()
}

// Equals returns true if the object identifier is equal to the given object identifier.
// Required as Go slices cannot be compared directly.
func (oid ObjectIdentifier) Equals(o ObjectIdentifier) bool {
	if len(oid) != len(o) {
		return false
	}

	for i, v := range oid {
		if v != o[i] {
			return false
		}
	}

	return true
}

// IsBefore returns true if the object identifier is before the given object identifier.
func (oid ObjectIdentifier) IsBefore(o ObjectIdentifier) bool {
	for i, v := range oid {
		if i >= len(o) {
			return false
		}

		if v < o[i] {
			return true
		}

		if v > o[i] {
			return false
		}
	}

	return true
}

// IsAfter returns true if the object identifier is after the given object identifier.
// This is the inverse of IsBefore and is provided for convenience:
//
//	oid.IsAfter(o) == o.IsBefore(oid)
func (oid ObjectIdentifier) IsAfter(o ObjectIdentifier) bool {
	return o.IsBefore(oid)
}

// IsPrefixOf returns true if the object identifier is a prefix of the given object identifier.
func (oid ObjectIdentifier) IsPrefixOf(o ObjectIdentifier) bool {
	if len(oid) > len(o) {
		return false
	}

	for i, v := range oid {
		if v != o[i] {
			return false
		}
	}

	return true
}

// IsScalar returns true if the object identifier is of a scalar node.
// A scalar node must have a zero sub-identifier at the end of the OID.
func IsScalar(oid ObjectIdentifier) bool {
	if len(oid) == 0 {
		return false
	}

	return oid[len(oid)-1] == 0
}

// IsValid returns true if the object identifier is valid.
// This is a shorthand for oid.Validate() == nil and is provided for convenience.
func (oid ObjectIdentifier) IsValid() bool {
	return oid.Validate() == nil
}

// Validate returns an error if the object identifier is invalid.
// The error will describe the reason for the invalidity.
func (oid ObjectIdentifier) Validate() error {
	if len(oid) < 2 {
		return errors.New("must have at least two sub-identifiers")
	}

	for i, v := range oid {
		switch {
		case v < 0:
			return fmt.Errorf("sub-identifier at position %d is negative: %v", i+1, v)
		case v > math.MaxUint32:
			return fmt.Errorf("sub-identifier at position %d is too large: %v", i+1, v)
		case i == 0 && v > 2:
			return fmt.Errorf("first sub-identifier must be 0, 1, or 2: %v", v)
		}
	}

	return nil
}

// String returns the string representation of the object identifier.
func (oid ObjectIdentifier) String() string {
	if len(oid) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, v := range oid {
		if i > 0 {
			sb.WriteByte('.')
		}

		sb.WriteString(strconv.Itoa(v))
	}

	return sb.String()
}
