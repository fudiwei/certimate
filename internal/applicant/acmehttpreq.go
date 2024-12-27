package applicant

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/go-acme/lego/v4/providers/dns/httpreq"

	"github.com/usual2970/certimate/internal/domain"
)

type acmeHttpReqApplicant struct {
	option *ApplyOption
}

func NewACMEHttpReqApplicant(option *ApplyOption) Applicant {
	return &acmeHttpReqApplicant{
		option: option,
	}
}

func (a *acmeHttpReqApplicant) Apply() (*Certificate, error) {
	access := &domain.ACMEHttpReqAccessConfig{}
	json.Unmarshal([]byte(a.option.AccessConfig), access)

	config := httpreq.NewDefaultConfig()
	endpoint, _ := url.Parse(access.Endpoint)
	config.Endpoint = endpoint
	config.Mode = access.Mode
	config.Username = access.Username
	config.Password = access.Password
	if a.option.PropagationTimeout != 0 {
		config.PropagationTimeout = time.Duration(a.option.PropagationTimeout) * time.Second
	}

	provider, err := httpreq.NewDNSProviderConfig(config)
	if err != nil {
		return nil, err
	}

	return apply(a.option, provider)
}
