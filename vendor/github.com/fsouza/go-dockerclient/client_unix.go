<<<<<<< HEAD
// +build !windows
=======
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
// Copyright 2016 go-dockerclient authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
=======
// +build !windows

>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
package docker

import (
	"context"
	"net"
	"net/http"

	"github.com/hashicorp/go-cleanhttp"
)

// initializeNativeClient initializes the native Unix domain socket client on
// Unix-style operating systems
func (c *Client) initializeNativeClient() {
	if c.endpointURL.Scheme != unixProtocol {
		return
	}
	socketPath := c.endpointURL.Path
	tr := cleanhttp.DefaultTransport()
	tr.Dial = func(network, addr string) (net.Conn, error) {
		return c.Dialer.Dial(network, addr)
	}
	tr.DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		return c.Dialer.Dial(unixProtocol, socketPath)
	}
	c.nativeHTTPClient = &http.Client{Transport: tr}
}
