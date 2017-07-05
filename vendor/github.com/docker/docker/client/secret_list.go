package client

import (
	"encoding/json"
	"net/url"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

// SecretList returns the list of secrets.
func (cli *Client) SecretList(ctx context.Context, options types.SecretListOptions) ([]swarm.Secret, error) {
<<<<<<< HEAD
=======
	if err := cli.NewVersionError("1.25", "secret list"); err != nil {
		return nil, err
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	query := url.Values{}

	if options.Filters.Len() > 0 {
		filterJSON, err := filters.ToParam(options.Filters)
		if err != nil {
			return nil, err
		}

		query.Set("filters", filterJSON)
	}

	resp, err := cli.get(ctx, "/secrets", query, nil)
	if err != nil {
		return nil, err
	}

	var secrets []swarm.Secret
	err = json.NewDecoder(resp.body).Decode(&secrets)
	ensureReaderClosed(resp)
	return secrets, err
}
