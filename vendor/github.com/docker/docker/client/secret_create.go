package client

import (
	"encoding/json"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

// SecretCreate creates a new Secret.
func (cli *Client) SecretCreate(ctx context.Context, secret swarm.SecretSpec) (types.SecretCreateResponse, error) {
	var response types.SecretCreateResponse
<<<<<<< HEAD
=======
	if err := cli.NewVersionError("1.25", "secret create"); err != nil {
		return response, err
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	resp, err := cli.post(ctx, "/secrets/create", nil, secret, nil)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.body).Decode(&response)
	ensureReaderClosed(resp)
	return response, err
}
