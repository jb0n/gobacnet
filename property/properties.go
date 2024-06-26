package property

import "fmt"

const (
	SoftwareRevision uint32 = 12
	Description      uint32 = 28
	FileSize         uint32 = 42
	FileType         uint32 = 43
	FirmwareRevision uint32 = 44
	ModelName        uint32 = 70
	ObjectIdentifier uint32 = 75
	ObjectList       uint32 = 76
	ObjectName       uint32 = 77
	ObjectReference  uint32 = 78
	ObjectType       uint32 = 79
	PresentValue     uint32 = 85
	Units            uint32 = 117
	VendorID         uint32 = 120
	VendorName       uint32 = 121
)

const (
	DescriptionStr = "Description"
	ObjectNameStr  = "ObjectName"
)

// enumMapping should be treated as read only.
var enumMapping = map[string]uint32{
	DescriptionStr:     Description,
	"FileSize":         FileSize,
	"FileType":         FileType,
	"FirmwareRevision": FirmwareRevision,
	"ModelName":        ModelName,
	"ObjectIdentifier": ObjectIdentifier,
	"ObjectList":       ObjectList,
	ObjectNameStr:      ObjectName,
	"ObjectReference":  ObjectReference,
	"ObjectType":       ObjectType,
	"PresentValue":     PresentValue,
	"Units":            Units,
	"VendorID":         VendorID,
	"VendorName":       VendorName,
}

var strMapping = map[uint32]string{}

// listOfKeys should be treated as read only after init
var listOfKeys []string

func init() {
	listOfKeys = make([]string, len(enumMapping))
	i := 0
	for k := range enumMapping {
		listOfKeys[i] = k
		i++
	}

	for k, v := range enumMapping {
		strMapping[v] = k
	}
}

func Keys() map[string]uint32 {
	// A copy is made since we do not want outside packages editing our keys by
	// accident
	keys := make(map[string]uint32)
	for k, v := range enumMapping {
		keys[k] = v
	}
	return keys
}

func Get(s string) (uint32, error) {
	if v, ok := enumMapping[s]; ok {
		return v, nil
	}
	err := fmt.Errorf("%s is not a valid property.", s)
	return 0, err
}

// String returns a human readible string of the given property
func String(prop uint32) string {
	s, ok := strMapping[prop]
	if !ok {
		return "Unknown"
	}
	return fmt.Sprintf("%s (%d)", s, prop)
}

// The bool in the map doesn't actually matter since it won't be used.
var deviceProperties = map[uint32]bool{
	ObjectList: true,
}

func IsDeviceProperty(id uint32) bool {
	_, ok := deviceProperties[id]
	return ok
}
