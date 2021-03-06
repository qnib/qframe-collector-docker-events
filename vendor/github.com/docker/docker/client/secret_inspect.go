package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

// SecretInspectWithRaw returns the secret information with raw data
func (cli *Client) SecretInspectWithRaw(ctx context.Context, id string) (swarm.Secret, []byte, error) {
<<<<<<< HEAD
=======
	if err := cli.NewVersionError("1.25", "secret inspect"); err != nil {
		return swarm.Secret{}, nil, err
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	resp, err := cli.get(ctx, "/secrets/"+id, nil, nil)
	if err != nil {
		if resp.statusCode == http.StatusNotFound {
			return swarm.Secret{}, nil, secretNotFoundError{id}
		}
		return swarm.Secret{}, nil, err
	}
	defer ensureReaderClosed(resp)

	body, err := ioutil.ReadAll(resp.body)
	if err != nil {
		return swarm.Secret{}, nil, err
	}

	var secret swarm.Secret
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&secret)

	return secret, body, err
}
