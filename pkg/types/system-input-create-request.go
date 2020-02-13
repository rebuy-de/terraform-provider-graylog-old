package types

// https://github.com/Graylog2/graylog2-server/blob/master/graylog2-server/src/main/java/org/graylog2/rest/models/system/inputs/requests/InputCreateRequest.java
type SystemInputCreateRequest struct {
	Title         string                 `json:"title"`
	Type          string                 `json:"type"`
	Global        bool                   `json:"global"`
	Configuration map[string]interface{} `json:"configuration"`
	Node          string                 `json:"node"`
}
