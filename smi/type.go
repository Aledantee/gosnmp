package smi

import "fmt"

// Type is a type that represents a type in the SMI.
type Type interface {
	// Name returns the name of the type.
	Name() string
	// Description returns the description of the type.
	Description() string
	// BaseType returns the base type of the type.
	// A base type returns itself.
	BaseType() BaseType
	// Units returns the units of the type.
	// Returns an empty string if the type has no units.
	Units() string
}

// BaseType is a type that represents the base type of any non-basic type.
type BaseType int

const (
	BaseTypeInteger32 BaseType = iota
	BaseTypeInteger
	BaseTypeOctetString
	BaseTypeObjectIdentifier
	BaseTypeBits
	BaseTypeIpAddress
	BaseTypeCounter32
	BaseTypeGauge32
	BaseTypeTimeTicks
	BaseTypeOpaque
	BaseTypeCounter64
	BaseTypeUnsigned32
)

func (b BaseType) Name() string {
	switch b {
	case BaseTypeInteger32:
		return "Integer32"
	case BaseTypeInteger:
		return "Integer"
	case BaseTypeOctetString:
		return "OctetString"
	case BaseTypeObjectIdentifier:
		return "ObjectIdentifier"
	case BaseTypeBits:
		return "Bits"
	case BaseTypeIpAddress:
		return "IpAddress"
	case BaseTypeCounter32:
		return "Counter32"
	case BaseTypeGauge32:
		return "Gauge32"
	case BaseTypeTimeTicks:
		return "TimeTicks"
	case BaseTypeOpaque:
		return "Opaque"
	case BaseTypeCounter64:
		return "Counter64"
	case BaseTypeUnsigned32:
		return "Unsigned32"
	default:
		panic(fmt.Sprintf("type %d not covered by BaseType.Name()", b))
	}
}

// Description returns the description of the base type.
// Values are taken from RFC 2578 section 7.
func (b BaseType) Description() string {
	switch b {
	case BaseTypeInteger32, BaseTypeInteger:
		return `The Integer32 type represents integer-valued information between
   -2^31 and 2^31-1 inclusive (-2147483648 to 2147483647 decimal).  This
   type is indistinguishable from the INTEGER type.  Both the INTEGER
   and Integer32 types may be sub-typed to be more constrained than the
   Integer32 type.

   The INTEGER type (but not the Integer32 type) may also be used to
   represent integer-valued information as named-number enumerations.
   In this case, only those named-numbers so enumerated may be present
   as a value.  Note that although it is recommended that enumerated
   values start at 1 and be numbered contiguously, any valid value for
   Integer32 is allowed for an enumerated value and, further, enumerated
   values needn't be contiguously assigned.

   Finally, a label for a named-number enumeration must consist of one
   or more letters or digits, up to a maximum of 64 characters, and the
   initial character must be a lower-case letter.  (However, labels
   longer than 32 characters are not recommended.)  Note that hyphens
   are not allowed by this specification (except for use by information
   modules converted from SMIv1 which did allow hyphens).`
	case BaseTypeOctetString:
		return `The OCTET STRING type represents arbitrary binary or textual data.
   Although the SMI-specified size limitation for this type is 65535
   octets, MIB designers should realize that there may be implementation
   and interoperability limitations for sizes in excess of 255 octets.`
	case BaseTypeObjectIdentifier:
		return `The OBJECT IDENTIFIER type represents administratively assigned
   names.  Any instance of this type may have at most 128 sub-
   identifiers.  Further, each sub-identifier must not exceed the value
   2^32-1 (4294967295 decimal).`
	case BaseTypeBits:
		return `The BITS construct represents an enumeration of named bits.  This
   collection is assigned non-negative, contiguous (but see below)
   values, starting at zero.  Only those named-bits so enumerated may be
   present in a value.  (Thus, enumerations must be assigned to
   consecutive bits; however, see Section 9 for refinements of an object
   with this syntax.)

   As part of updating an information module, for an object defined
   using the BITS construct, new enumerations can be added or existing
   enumerations can have new labels assigned to them.  After an
   enumeration is added, it might not be possible to distinguish between
   an implementation of the updated object for which the new enumeration
   is not asserted, and an implementation of the object prior to the
   addition.  Depending on the circumstances, such an ambiguity could
   either be desirable or could be undesirable.  The means to avoid such
   an ambiguity is dependent on the encoding of values on the wire;
   however, one possibility is to define new enumerations starting at
   the next multiple of eight bits.  (Of course, this can also result in
   the enumerations no longer being contiguous.)

   Although there is no SMI-specified limitation on the number of
   enumerations (and therefore on the length of a value), except as may
   be imposed by the limit on the length of an OCTET STRING, MIB
   designers should realize that there may be implementation and
   interoperability limitations for sizes in excess of 128 bits.

   Finally, a label for a named-number enumeration must consist of one
   or more letters or digits, up to a maximum of 64 characters, and the
   initial character must be a lower-case letter.  (However, labels
   longer than 32 characters are not recommended.)  Note that hyphens
   are not allowed by this specification.`
	case BaseTypeIpAddress:
		return `The IpAddress type represents a 32-bit internet address.  It is
   represented as an OCTET STRING of length 4, in network byte-order.

   Note that the IpAddress type is a tagged type for historical reasons.
   Network addresses should be represented using an invocation of the
   TEXTUAL-CONVENTION macro [3].`
	case BaseTypeCounter32:
		return `The Counter32 type represents a non-negative integer which
   monotonically increases until it reaches a maximum value of 2^32-1
   (4294967295 decimal), when it wraps around and starts increasing
   again from zero.

   Counters have no defined "initial" value, and thus, a single value of
   a Counter has (in general) no information content.  Discontinuities
   in the monotonically increasing value normally occur at re-
   initialization of the management system, and at other times as
   specified in the description of an object-type using this ASN.1 type.
   If such other times can occur, for example, the creation of an object
   instance at times other than re-initialization, then a corresponding
   object should be defined, with an appropriate SYNTAX clause, to
   indicate the last discontinuity.  Examples of appropriate SYNTAX
   clause include:  TimeStamp (a textual convention defined in [3]),
   DateAndTime (another textual convention from [3]) or TimeTicks.

   The value of the MAX-ACCESS clause for objects with a SYNTAX clause
   value of Counter32 is either "read-only" or "accessible-for-notify".

   A DEFVAL clause is not allowed for objects with a SYNTAX clause value
   of Counter32.`
	case BaseTypeGauge32:
		return `The Gauge32 type represents a non-negative integer, which may
   increase or decrease, but shall never exceed a maximum value, nor
   fall below a minimum value.  The maximum value can not be greater
   than 2^32-1 (4294967295 decimal), and the minimum value can not be
   smaller than 0.  The value of a Gauge32 has its maximum value
   whenever the information being modeled is greater than or equal to
   its maximum value, and has its minimum value whenever the information
   being modeled is smaller than or equal to its minimum value.  If the
   information being modeled subsequently decreases below (increases
   above) the maximum (minimum) value, the Gauge32 also decreases
   (increases).  (Note that despite of the use of the term "latched" in
   the original definition of this type, it does not become "stuck" at
   its maximum or minimum value.)`
	case BaseTypeTimeTicks:
		return `The TimeTicks type represents a non-negative integer which represents
   the time, modulo 2^32 (4294967296 decimal), in hundredths of a second
   between two epochs.  When objects are defined which use this ASN.1
   type, the description of the object identifies both of the reference
   epochs.

   For example, [3] defines the TimeStamp textual convention which is
   based on the TimeTicks type.  With a TimeStamp, the first reference
   epoch is defined as the time when sysUpTime [5] was zero, and the
   second reference epoch is defined as the current value of sysUpTime.

   The TimeTicks type may not be sub-typed.`
	case BaseTypeOpaque:
		return `The Opaque type is provided solely for backward-compatibility, and
   shall not be used for newly-defined object types.

   The Opaque type supports the capability to pass arbitrary ASN.1
   syntax.  A value is encoded using the ASN.1 Basic Encoding Rules [4]
   into a string of octets.  This, in turn, is encoded as an OCTET
   STRING, in effect "double-wrapping" the original ASN.1 value.

   Note that a conforming implementation need only be able to accept and
   recognize opaquely-encoded data.  It need not be able to unwrap the
   data and then interpret its contents.

   A requirement on "standard" MIB modules is that no object may have a
   SYNTAX clause value of Opaque.`
	case BaseTypeCounter64:
		return `The Counter64 type represents a non-negative integer which
   monotonically increases until it reaches a maximum value of 2^64-1
   (18446744073709551615 decimal), when it wraps around and starts
   increasing again from zero.

   Counters have no defined "initial" value, and thus, a single value of
   a Counter has (in general) no information content.  Discontinuities
   in the monotonically increasing value normally occur at re-
   initialization of the management system, and at other times as
   specified in the description of an object-type using this ASN.1 type.
   If such other times can occur, for example, the creation of an object
   instance at times other than re-initialization, then a corresponding
   object should be defined, with an appropriate SYNTAX clause, to
   indicate the last discontinuity.  Examples of appropriate SYNTAX
   clause are:  TimeStamp (a textual convention defined in [3]),
   DateAndTime (another textual convention from [3]) or TimeTicks.

   The value of the MAX-ACCESS clause for objects with a SYNTAX clause
   value of Counter64 is either "read-only" or "accessible-for-notify".

   A requirement on "standard" MIB modules is that the Counter64 type
   may be used only if the information being modeled would wrap in less
   than one hour if the Counter32 type was used instead.

   A DEFVAL clause is not allowed for objects with a SYNTAX clause value
   of Counter64.`
	case BaseTypeUnsigned32:
		return `The Unsigned32 type represents integer-valued information between 0
   and 2^32-1 inclusive (0 to 4294967295 decimal).`
	default:
		panic(fmt.Sprintf("type %d not covered by BaseType.Description()", b))
	}
}

func (b BaseType) BaseType() BaseType {
	return b
}

func (b BaseType) String() string {
	return b.Name()
}
