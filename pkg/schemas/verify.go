package schemas

import (
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) VerifyObject(doc map[string]interface{}, list []*st.SchemaKindDefinition) error {
	fields := make(map[string]st.SchemaKind)
	for _, c := range list {
		fields[c.Name] = c.Field
	}

	for key, value := range doc {
		if _, ok := fields[key]; !ok {
			return errSchemaFieldsInvalid
		}
		if !CheckValueOfField(value, fields[key]) {
			return errSchemaFieldsInvalid
		}
	}

	return nil
}

// Current supported IPLD types, will be adding more once supporting of Links and Complex types (Object)
func CheckValueOfField(value interface{}, fieldType st.SchemaKind) bool {
	switch value.(type) {
	case int:
		return fieldType == st.SchemaKind_INT
	case float64:
		return fieldType == st.SchemaKind_FLOAT
	case bool:
		return fieldType == st.SchemaKind_BOOL
	case string:
		return fieldType == st.SchemaKind_STRING
	case []byte:
		return fieldType == st.SchemaKind_BYTES
	case interface{}:
		return fieldType == st.SchemaKind_ANY
	default:
		return false
	}
}