package types

// https://github.com/Graylog2/graylog2-server/blob/master/graylog2-server/src/main/java/org/graylog2/rest/models/system/inputs/responses/InputCreated.java

type SystemInputCreateResponse struct {
	ID string `json:"id"`
}
