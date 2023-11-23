package types

type ACTION_TYPES int32

const (
	ACTION_TYPES_UNSPECIFIED  ACTION_TYPES = 0
	ACTION_TYPES_LIST         ACTION_TYPES = 1
	ACTION_TYPES_GET          ACTION_TYPES = 2
	ACTION_TYPES_CREATE       ACTION_TYPES = 3
	ACTION_TYPES_UPDATE       ACTION_TYPES = 4
	ACTION_TYPES_DELETE       ACTION_TYPES = 5
	ACTION_TYPES_INIT         ACTION_TYPES = 6
	ACTION_TYPES_REQUIRED     ACTION_TYPES = 7
	ACTION_TYPES_OPTIONAL     ACTION_TYPES = 8
	ACTION_TYPES_CONNECT      ACTION_TYPES = 9
	ACTION_TYPES_DISCONNECT   ACTION_TYPES = 10
	ACTION_TYPES_REGISTER     ACTION_TYPES = 11
	ACTION_TYPES_UNREGISTER   ACTION_TYPES = 12
	ACTION_TYPES_SUBSCRIBE    ACTION_TYPES = 13
	ACTION_TYPES_UNSUBSCRIBE  ACTION_TYPES = 14
	ACTION_TYPES_UNRECOGNIZED ACTION_TYPES = 15
)

// Enum value maps for ACTION_TYPES.
var (
	ACTION_TYPES_name = map[int32]string{
		0:  "ACTION_TYPES_UNSPECIFIED",
		1:  "ACTION_TYPES_LIST",
		2:  "ACTION_TYPES_GET",
		3:  "ACTION_TYPES_CREATE",
		4:  "ACTION_TYPES_UPDATE",
		5:  "ACTION_TYPES_DELETE",
		6:  "ACTION_TYPES_INIT",
		7:  "ACTION_TYPES_REQUIRED",
		8:  "ACTION_TYPES_OPTIONAL",
		9:  "ACTION_TYPES_CONNECT",
		10: "ACTION_TYPES_DISCONNECT",
		11: "ACTION_TYPES_REGISTER",
		12: "ACTION_TYPES_UNREGISTER",
		13: "ACTION_TYPES_SUBSCRIBE",
		14: "ACTION_TYPES_UNSUBSCRIBE",
		15: "ACTION_TYPES_UNRECOGNIZED",
	}
	ACTION_TYPES_value = map[string]int32{
		"ACTION_TYPES_UNSPECIFIED":  0,
		"ACTION_TYPES_LIST":         1,
		"ACTION_TYPES_GET":          2,
		"ACTION_TYPES_CREATE":       3,
		"ACTION_TYPES_UPDATE":       4,
		"ACTION_TYPES_DELETE":       5,
		"ACTION_TYPES_INIT":         6,
		"ACTION_TYPES_REQUIRED":     7,
		"ACTION_TYPES_OPTIONAL":     8,
		"ACTION_TYPES_CONNECT":      9,
		"ACTION_TYPES_DISCONNECT":   10,
		"ACTION_TYPES_REGISTER":     11,
		"ACTION_TYPES_UNREGISTER":   12,
		"ACTION_TYPES_SUBSCRIBE":    13,
		"ACTION_TYPES_UNSUBSCRIBE":  14,
		"ACTION_TYPES_UNRECOGNIZED": 15,
	}
)

func (x ACTION_TYPES) Enum() *ACTION_TYPES {
	p := new(ACTION_TYPES)
	*p = x
	return p
}
