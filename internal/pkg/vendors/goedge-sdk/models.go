package goedgesdk

type BaseResponse interface{}

type ServerConfigData struct{}

type SSLCertData struct {
	IsOn        bool     `json:"isOn"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ServerName  string   `json:"serverName"`
	IsCA        bool     `json:"isCA"`
	CertData    []byte   `json:"certData"`
	KeyData     []byte   `json:"keyData"`
	TimeBeginAt int64    `json:"timeBeginAt"`
	TimeEndAt   int64    `json:"timeEndAt"`
	DNSNames    []string `json:"dnsNames"`
	CommonNames []string `json:"commonNames"`
}

type FindEnabledServerConfigRequest struct {
	ServerId int64 `json:"serverId"`
}

type FindEnabledServerConfigResponse struct {
	ServerConfigJSON []byte `json:"serverJSON"`
}

type ListSSLCertsRequest struct {
	IsCA         *bool     `json:"isCA,omitempty"`
	IsAvailable  *bool     `json:"isAvailable,omitempty"`
	IsExpired    *bool     `json:"isExpired,omitempty"`
	ExpiringDays *int32    `json:"expiringDays,omitempty"`
	Keyword      *string   `json:"keyword,omitempty"`
	UserId       *int64    `json:"userId,omitempty"`
	Domains      *[]string `json:"domains,omitempty"`
	Offset       *int64    `json:"offset,omitempty"`
	Size         *int64    `json:"size,omitempty"`
	UserOnly     *bool     `json:"userOnly,omitempty"`
}

type ListSSLCertsResponse struct {
	SSLCertsJSON []byte `json:"sslCertsJSON"`
}

type CreateSSLCertRequest struct {
	SSLCertData
}

type CreateSSLCertResponse struct {
	SSLCertId int64 `json:"sslCertId"`
}

type UpdateSSLCertRequest struct {
	SSLCertData
	SSLCertId int64 `json:"sslCertId"`
}

type UpdateSSLCertResponse struct{}
