package types

type MESSAGE_TYPES int32

const (
	MESSAGE_TYPES_UNSPECIFIED  MESSAGE_TYPES = 0
	MESSAGE_TYPES_UNICAST      MESSAGE_TYPES = 1
	MESSAGE_TYPES_BROADCAST    MESSAGE_TYPES = 2
	MESSAGE_TYPES_ECHO         MESSAGE_TYPES = 3
	MESSAGE_TYPES_HEALTH       MESSAGE_TYPES = 4
	MESSAGE_TYPES_TASK         MESSAGE_TYPES = 5
	MESSAGE_TYPES_CONNECTION   MESSAGE_TYPES = 6
	MESSAGE_TYPES_SUBSCRIPTION MESSAGE_TYPES = 7
	MESSAGE_TYPES_REGISTRATION MESSAGE_TYPES = 8
	MESSAGE_TYPES_SYSTEM       MESSAGE_TYPES = 9
	MESSAGE_TYPES_EVENT        MESSAGE_TYPES = 10
	MESSAGE_TYPES_UNRECOGNIZED MESSAGE_TYPES = 11
)

// Enum value maps for MESSAGE_TYPES.
var (
	MESSAGE_TYPES_name = map[int32]string{
		0:  "MESSAGE_TYPES_UNSPECIFIED",
		1:  "MESSAGE_TYPES_UNICAST",
		2:  "MESSAGE_TYPES_BROADCAST",
		3:  "MESSAGE_TYPES_ECHO",
		4:  "MESSAGE_TYPES_HEALTH",
		5:  "MESSAGE_TYPES_TASK",
		6:  "MESSAGE_TYPES_CONNECTION",
		7:  "MESSAGE_TYPES_SUBSCRIPTION",
		8:  "MESSAGE_TYPES_REGISTRATION",
		9:  "MESSAGE_TYPES_SYSTEM",
		10: "MESSAGE_TYPES_EVENT",
		11: "MESSAGE_TYPES_UNRECOGNIZED",
	}
	MESSAGE_TYPES_value = map[string]int32{
		"MESSAGE_TYPES_UNSPECIFIED":  0,
		"MESSAGE_TYPES_UNICAST":      1,
		"MESSAGE_TYPES_BROADCAST":    2,
		"MESSAGE_TYPES_ECHO":         3,
		"MESSAGE_TYPES_HEALTH":       4,
		"MESSAGE_TYPES_TASK":         5,
		"MESSAGE_TYPES_CONNECTION":   6,
		"MESSAGE_TYPES_SUBSCRIPTION": 7,
		"MESSAGE_TYPES_REGISTRATION": 8,
		"MESSAGE_TYPES_SYSTEM":       9,
		"MESSAGE_TYPES_EVENT":        10,
		"MESSAGE_TYPES_UNRECOGNIZED": 11,
	}
)

func (x MESSAGE_TYPES) Enum() *MESSAGE_TYPES {
	p := new(MESSAGE_TYPES)
	*p = x
	return p
}
