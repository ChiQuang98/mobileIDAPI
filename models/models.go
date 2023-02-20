package models

type IdentityRequest struct {
	Msisdn  string `json:"msisdn"`
	Address string `json:"address"`
	Port    int    `json:"port,omitempty"`
}

type Identity struct {
	Timestamp       string
	Phone           string
	IPPrivate       string
	PortPrivate     string
	IPDestination   string
	PortDestination string
}
type ResponseHTTP struct {
	Status string       `json:"status,omitempty"`
	Time   string       `json:"time,omitempty"`
	Data   *MessageData `json:"data"`
}
type MessageData struct {
	Message string `json:"message"`
}
