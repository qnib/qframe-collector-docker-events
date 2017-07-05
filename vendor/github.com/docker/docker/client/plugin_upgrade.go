package client

import (
<<<<<<< HEAD
	"fmt"
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	"io"
	"net/url"

	"github.com/docker/distribution/reference"
	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

// PluginUpgrade upgrades a plugin
func (cli *Client) PluginUpgrade(ctx context.Context, name string, options types.PluginInstallOptions) (rc io.ReadCloser, err error) {
<<<<<<< HEAD
=======
	if err := cli.NewVersionError("1.26", "plugin upgrade"); err != nil {
		return nil, err
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	query := url.Values{}
	if _, err := reference.ParseNormalizedNamed(options.RemoteRef); err != nil {
		return nil, errors.Wrap(err, "invalid remote reference")
	}
	query.Set("remote", options.RemoteRef)

	privileges, err := cli.checkPluginPermissions(ctx, query, options)
	if err != nil {
		return nil, err
	}

	resp, err := cli.tryPluginUpgrade(ctx, query, privileges, name, options.RegistryAuth)
	if err != nil {
		return nil, err
	}
	return resp.body, nil
}

func (cli *Client) tryPluginUpgrade(ctx context.Context, query url.Values, privileges types.PluginPrivileges, name, registryAuth string) (serverResponse, error) {
	headers := map[string][]string{"X-Registry-Auth": {registryAuth}}
<<<<<<< HEAD
	return cli.post(ctx, fmt.Sprintf("/plugins/%s/upgrade", name), query, privileges, headers)
=======
	return cli.post(ctx, "/plugins/"+name+"/upgrade", query, privileges, headers)
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
}
