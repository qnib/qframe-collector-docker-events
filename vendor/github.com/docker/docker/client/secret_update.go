package client

import (
	"net/url"
	"strconv"

	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

<<<<<<< HEAD
// SecretUpdate attempts to updates a Secret
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error {
=======
// SecretUpdate attempts to update a Secret
func (cli *Client) SecretUpdate(ctx context.Context, id string, version swarm.Version, secret swarm.SecretSpec) error {
	if err := cli.NewVersionError("1.25", "secret update"); err != nil {
		return err
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	query := url.Values{}
	query.Set("version", strconv.FormatUint(version.Index, 10))
	resp, err := cli.post(ctx, "/secrets/"+id+"/update", query, secret, nil)
	ensureReaderClosed(resp)
	return err
}
