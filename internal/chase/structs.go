package chase

//RequestBody is the struct for a standard POST request
type RequestBody struct {
	APT string
}

//AliasGroup is the struct for storing a group of APT aliases
type AliasGroup struct {
	Groups []string
}

//APTMappings is the interface for unmarshalling apt_mappings.json
type APTMappings map[string]AliasGroup
