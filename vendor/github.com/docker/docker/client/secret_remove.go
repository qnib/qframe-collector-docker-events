package client

import "golang.org/x/net/context"

// SecretRemove removes a Secret.
func (cli *Client) SecretRemove(ctx context.Context, id string) error {
<<<<<<< HEAD
=======
	if err := cli.NewVersionError("1.25", "secret remove"); err != nil {
		return err
	}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	resp, err := cli.delete(ctx, "/secrets/"+id, nil, nil)
	ensureReaderClosed(resp)
	return err
}
