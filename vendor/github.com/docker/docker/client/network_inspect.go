package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
)

// NetworkInspect returns the information for a specific network configured in the docker host.
<<<<<<< HEAD
func (cli *Client) NetworkInspect(ctx context.Context, networkID string, verbose bool) (types.NetworkResource, error) {
	networkResource, _, err := cli.NetworkInspectWithRaw(ctx, networkID, verbose)
=======
func (cli *Client) NetworkInspect(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, error) {
	networkResource, _, err := cli.NetworkInspectWithRaw(ctx, networkID, options)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	return networkResource, err
}

// NetworkInspectWithRaw returns the information for a specific network configured in the docker host and its raw representation.
<<<<<<< HEAD
func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, verbose bool) (types.NetworkResource, []byte, error) {
=======
func (cli *Client) NetworkInspectWithRaw(ctx context.Context, networkID string, options types.NetworkInspectOptions) (types.NetworkResource, []byte, error) {
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	var (
		networkResource types.NetworkResource
		resp            serverResponse
		err             error
	)
	query := url.Values{}
<<<<<<< HEAD
	if verbose {
		query.Set("verbose", "true")
	}
=======
	if options.Verbose {
		query.Set("verbose", "true")
	}
	if options.Scope != "" {
		query.Set("scope", options.Scope)
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	resp, err = cli.get(ctx, "/networks/"+networkID, query, nil)
	if err != nil {
		if resp.statusCode == http.StatusNotFound {
			return networkResource, nil, networkNotFoundError{networkID}
		}
		return networkResource, nil, err
	}
	defer ensureReaderClosed(resp)

	body, err := ioutil.ReadAll(resp.body)
	if err != nil {
		return networkResource, nil, err
	}
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&networkResource)
	return networkResource, body, err
}
