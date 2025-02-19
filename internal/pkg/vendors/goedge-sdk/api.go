package goedgesdk

func (c *Client) ListSSLCerts(req *ListSSLCertsRequest) (*ListSSLCertsResponse, error) {
	resp := ListSSLCertsResponse{}
	err := c.sendRequestWithResult("/SSLCertService/listSSLCerts", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) CreateSSLCert(req *CreateSSLCertRequest) (*CreateSSLCertResponse, error) {
	resp := CreateSSLCertResponse{}
	err := c.sendRequestWithResult("/SSLCertService/createSSLCert", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UpdateSSLCert(req *UpdateSSLCertRequest) (*UpdateSSLCertResponse, error) {
	resp := UpdateSSLCertResponse{}
	err := c.sendRequestWithResult("/SSLCertService/updateSSLCert", req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
