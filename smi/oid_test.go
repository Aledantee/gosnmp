package smi

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestObjectIdentifier_Equals(t *testing.T) {
	tests := []struct {
		name string
		oid1 ObjectIdentifier
		oid2 ObjectIdentifier
		want bool
	}{
		{"Same length and values", ObjectIdentifier{1, 2, 3}, ObjectIdentifier{1, 2, 3}, true},
		{"Different length", ObjectIdentifier{1, 2, 3}, ObjectIdentifier{1, 2, 3, 4}, false},
		{"Same length and different values", ObjectIdentifier{1, 2, 3}, ObjectIdentifier{1, 2, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oid1.Equals(tt.oid2); got != tt.want {
				t.Errorf("ObjectIdentifier.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectIdentifier_IsBefore(t *testing.T) {
	tests := []struct {
		name string
		oid1 ObjectIdentifier
		oid2 ObjectIdentifier
		want bool
	}{
		{"OID1 is before OID2", ObjectIdentifier{1, 2, 3}, ObjectIdentifier{2, 3, 4}, true},
		{"OID1 is not before OID2", ObjectIdentifier{2, 3, 4}, ObjectIdentifier{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oid1.IsBefore(tt.oid2); got != tt.want {
				t.Errorf("ObjectIdentifier.IsBefore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectIdentifier_IsAfter(t *testing.T) {
	tests := []struct {
		name string
		oid1 ObjectIdentifier
		oid2 ObjectIdentifier
		want bool
	}{
		{"OID1 is after OID2", ObjectIdentifier{2, 3, 4}, ObjectIdentifier{1, 2, 3}, true},
		{"OID1 is not after OID2", ObjectIdentifier{1, 2, 3}, ObjectIdentifier{2, 3, 4}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oid1.IsAfter(tt.oid2); got != tt.want {
				t.Errorf("ObjectIdentifier.IsAfter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectIdentifier_IsPrefixOf(t *testing.T) {
	tests := []struct {
		name string
		oid1 ObjectIdentifier
		oid2 ObjectIdentifier
		want bool
	}{
		{"OID1 is prefix of OID2", ObjectIdentifier{1, 2}, ObjectIdentifier{1, 2, 3}, true},
		{"OID1 is not prefix of OID2", ObjectIdentifier{1, 3}, ObjectIdentifier{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oid1.IsPrefixOf(tt.oid2); got != tt.want {
				t.Errorf("ObjectIdentifier.IsPrefixOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestObjectIdentifier_Validate(t *testing.T) {
	tests := []struct {
		name    string
		oid     ObjectIdentifier
		wantErr error
	}{
		{"Valid OID", ObjectIdentifier{1, 2}, nil},
		{"Invalid OID - less than two sub-identifiers", ObjectIdentifier{1}, errors.New("must have at least two sub-identifiers")},
		{"Invalid OID - negative sub-identifier", ObjectIdentifier{-1, 2}, errors.New("sub-identifier at position 1 is negative: -1")},
		{"Invalid OID - first sub-identifier greater than 2", ObjectIdentifier{3, 2}, errors.New("first sub-identifier must be 0, 1, or 2: 3")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.oid.Validate(); err != nil {
				if tt.wantErr == nil || err.Error() != tt.wantErr.Error() {
					t.Errorf("ObjectIdentifier.Validate() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestObjectIdentifier_String(t *testing.T) {
	tests := []struct {
		name string
		oid  ObjectIdentifier
		want string
	}{
		{"OID to string", ObjectIdentifier{1, 2, 3}, "1.2.3"},
		{"Empty OID to string", ObjectIdentifier{}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.oid.String(); got != tt.want {
				t.Errorf("ObjectIdentifier.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsScalar(t *testing.T) {
	tests := []struct {
		name string
		oid  ObjectIdentifier
		want bool
	}{
		{"OID is scalar", ObjectIdentifier{1, 0}, true},
		{"OID is not scalar", ObjectIdentifier{1, 2}, false},
		{"Empty OID is not scalar", ObjectIdentifier{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsScalar(tt.oid); got != tt.want {
				t.Errorf("IsScalar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseObjectIdentifier(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    ObjectIdentifier
		wantErr error
	}{
		{"Valid OID", "1.2.3", ObjectIdentifier{1, 2, 3}, nil},
		{"Invalid OID - consecutive periods", "1..3", nil, errors.New("cannot have consecutive periods")},
		{"Invalid OID - ends with a period", "1.2.", nil, errors.New("cannot end with a period")},
		{"Invalid OID - invalid character", "1.2a.3", nil, fmt.Errorf("invalid character at position 3: a")},
		{"Valid OID - empty string", "", nil, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseObjectIdentifier(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseObjectIdentifier() = %v, want %v", got, tt.want)
			}
			if tt.wantErr == nil && err != nil {
				t.Errorf("ParseObjectIdentifier() unexpected error = %v", err)
			}
			if tt.wantErr != nil && (err == nil || err.Error() != tt.wantErr.Error()) {
				t.Errorf("ParseObjectIdentifier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
