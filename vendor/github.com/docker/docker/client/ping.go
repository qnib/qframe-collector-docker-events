package client

import (
<<<<<<< HEAD
	"fmt"

=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// Ping pings the server and returns the value of the "Docker-Experimental", "OS-Type" & "API-Version" headers
func (cli *Client) Ping(ctx context.Context) (types.Ping, error) {
	var ping types.Ping
<<<<<<< HEAD
	req, err := cli.buildRequest("GET", fmt.Sprintf("%s/_ping", cli.basePath), nil, nil)
=======
	req, err := cli.buildRequest("GET", cli.basePath+"/_ping", nil, nil)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	if err != nil {
		return ping, err
	}
	serverResp, err := cli.doRequest(ctx, req)
	if err != nil {
		return ping, err
	}
	defer ensureReaderClosed(serverResp)

<<<<<<< HEAD
	ping.APIVersion = serverResp.header.Get("API-Version")

	if serverResp.header.Get("Docker-Experimental") == "true" {
		ping.Experimental = true
	}

	ping.OSType = serverResp.header.Get("OSType")

	return ping, nil
=======
	if serverResp.header != nil {
		ping.APIVersion = serverResp.header.Get("API-Version")

		if serverResp.header.Get("Docker-Experimental") == "true" {
			ping.Experimental = true
		}
		ping.OSType = serverResp.header.Get("OSType")
	}

	err = cli.checkResponseErr(serverResp)
	return ping, err
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}
