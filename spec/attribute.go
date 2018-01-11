package spec

// AttributeFormat represent the allowed format for an attribute.
type AttributeFormat string

// Various values for an AttributeFormat.
const (
	AttributeFormatFree  AttributeFormat = "free"
	AttributeFormatEmail AttributeFormat = "email"
	AttributeFormatPhone AttributeFormat = "phone"
	AttributeFormatIPv4  AttributeFormat = "ipv4"
	AttributeFormatIPv6  AttributeFormat = "ipv6"
	AttributeFormatCIDR  AttributeFormat = "cidr"
	AttributeFormatDate  AttributeFormat = "date"
)

// AttributeUniqueScope represent the unique scope for an attribute.
type AttributeUniqueScope string

// Various values for an AttributeUniqueScope.
const (
	AttributeUniqueScopeLocal  AttributeUniqueScope = "local"
	AttributeUniqueScopeGlobal AttributeUniqueScope = "global"
)

// AttributeType represents the various type for an attribute.
type AttributeType string

// Various values for AttributeType.
const (
	AttributeTypeString AttributeType = "string"
	AttributeTypeInt    AttributeType = "integer"
	AttributeTypeFloat  AttributeType = "float"
	AttributeTypeBool   AttributeType = "boolean"
	AttributeTypeEnum   AttributeType = "enum"
	AttributeTypeList   AttributeType = "list"
	AttributeTypeObject AttributeType = "object"
	AttributeTypeTime   AttributeType = "time"
	AttributeTypeExt    AttributeType = "external"
)

// AttributeNameConverterFunc is the type of a attribute name conveter.
type AttributeNameConverterFunc func(name string) string

// AttributeTypeConverterFunc is the type of a attribute type conveter.
type AttributeTypeConverterFunc func(typ AttributeType, subtype string) (converted string, provider string)

// An Attribute represents a monolithe specification attribute.
type Attribute struct {
	AllowedChars   string               `json:"allowed_chars"`
	AllowedChoices []string             `json:"allowed_choices"`
	Autogenerated  bool                 `json:"autogenerated"`
	Availability   bool                 `json:"availability"`
	CreationOnly   bool                 `json:"creation_only"`
	DefaultOrder   bool                 `json:"default_order"`
	DefaultValue   interface{}          `json:"default_value"`
	Deprecated     bool                 `json:"deprecated"`
	Description    string               `json:"description"`
	Exposed        bool                 `json:"exposed"`
	Filterable     bool                 `json:"filterable"`
	ForeignKey     bool                 `json:"foreign_key"`
	Format         AttributeFormat      `json:"format"`
	Getter         bool                 `json:"getter"`
	Identifier     bool                 `json:"identifier"`
	Index          bool                 `json:"index"`
	MaxLength      uint16               `json:"max_length"`
	MaxValue       float64              `json:"max_value"`
	MinLength      uint16               `json:"min_length"`
	MinValue       float64              `json:"min_value"`
	Name           string               `json:"name"`
	Orderable      bool                 `json:"orderable"`
	PrimaryKey     bool                 `json:"primary_key"`
	ReadOnly       bool                 `json:"read_only"`
	Required       bool                 `json:"required"`
	Secret         bool                 `json:"secret"`
	Setter         bool                 `json:"setter"`
	Stored         bool                 `json:"stored"`
	SubType        string               `json:"subtype"`
	Transient      bool                 `json:"transient"`
	Type           AttributeType        `json:"type"`
	Unique         bool                 `json:"unique"`
	UniqueScope    AttributeUniqueScope `json:"unique_scope"`

	ConvertedName string `json:"_"`
	ConvertedType string `json:"_"`
	TypeProvider  string `json:"_"`
	Initializer   string `json:"_"`
}
