package apisix_test

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	provider "github.com/usual2970/certimate/internal/pkg/core/deployer/providers/apisix"
)

var (
	fInputCertPath string
	fInputKeyPath  string
	fServerUrl     string
	fApiKey        string
	fCertificateId string
)

func init() {
	argsPrefix := "CERTIMATE_DEPLOYER_APISIX_"

	flag.StringVar(&fInputCertPath, argsPrefix+"INPUTCERTPATH", "", "")
	flag.StringVar(&fInputKeyPath, argsPrefix+"INPUTKEYPATH", "", "")
	flag.StringVar(&fServerUrl, argsPrefix+"SERVERURL", "", "")
	flag.StringVar(&fApiKey, argsPrefix+"APIKEY", "", "")
	flag.StringVar(&fCertificateId, argsPrefix+"CERTIFICATEID", "", "")
}

/*
Shell command to run this test:

	go test -v ./apisix_test.go -args \
	--CERTIMATE_DEPLOYER_APISIX_INPUTCERTPATH="/path/to/your-input-cert.pem" \
	--CERTIMATE_DEPLOYER_APISIX_INPUTKEYPATH="/path/to/your-input-key.pem" \
	--CERTIMATE_DEPLOYER_APISIX_SERVERURL="http://127.0.0.1:9080" \
	--CERTIMATE_DEPLOYER_APISIX_APIKEY="your-api-key" \
	--CERTIMATE_DEPLOYER_APISIX_CERTIFICATEID="your-cerficiate-id"
*/
func TestDeploy(t *testing.T) {
	flag.Parse()

	t.Run("Deploy", func(t *testing.T) {
		t.Log(strings.Join([]string{
			"args:",
			fmt.Sprintf("INPUTCERTPATH: %v", fInputCertPath),
			fmt.Sprintf("INPUTKEYPATH: %v", fInputKeyPath),
			fmt.Sprintf("SERVERURL: %v", fServerUrl),
			fmt.Sprintf("APIKEY: %v", fApiKey),
			fmt.Sprintf("CERTIFICATEID: %v", fCertificateId),
		}, "\n"))

		deployer, err := provider.NewDeployer(&provider.DeployerConfig{
			ServerUrl:                fServerUrl,
			ApiKey:                   fApiKey,
			AllowInsecureConnections: true,
			ResourceType:             provider.RESOURCE_TYPE_CERTIFICATE,
			CertificateId:            fCertificateId,
		})
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		fInputCertData, _ := os.ReadFile(fInputCertPath)
		fInputKeyData, _ := os.ReadFile(fInputKeyPath)
		res, err := deployer.Deploy(context.Background(), string(fInputCertData), string(fInputKeyData))
		if err != nil {
			t.Errorf("err: %+v", err)
			return
		}

		t.Logf("ok: %v", res)
	})
}
