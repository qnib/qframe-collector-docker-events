package client

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"golang.org/x/net/context"
)

// ServiceUpdate updates a Service.
func (cli *Client) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error) {
	var (
<<<<<<< HEAD
		headers map[string][]string
		query   = url.Values{}
	)

	if options.EncodedRegistryAuth != "" {
		headers = map[string][]string{
			"X-Registry-Auth": {options.EncodedRegistryAuth},
		}
=======
		query   = url.Values{}
		distErr error
	)

	headers := map[string][]string{
		"version": {cli.version},
	}

	if options.EncodedRegistryAuth != "" {
		headers["X-Registry-Auth"] = []string{options.EncodedRegistryAuth}
>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	}

	if options.RegistryAuthFrom != "" {
		query.Set("registryAuthFrom", options.RegistryAuthFrom)
	}

	if options.Rollback != "" {
		query.Set("rollback", options.Rollback)
	}

	query.Set("version", strconv.FormatUint(version.Index, 10))

<<<<<<< HEAD
=======
	// ensure that the image is tagged
	if taggedImg := imageWithTagString(service.TaskTemplate.ContainerSpec.Image); taggedImg != "" {
		service.TaskTemplate.ContainerSpec.Image = taggedImg
	}

	// Contact the registry to retrieve digest and platform information
	// This happens only when the image has changed
	if options.QueryRegistry {
		distributionInspect, err := cli.DistributionInspect(ctx, service.TaskTemplate.ContainerSpec.Image, options.EncodedRegistryAuth)
		distErr = err
		if err == nil {
			// now pin by digest if the image doesn't already contain a digest
			if img := imageWithDigestString(service.TaskTemplate.ContainerSpec.Image, distributionInspect.Descriptor.Digest); img != "" {
				service.TaskTemplate.ContainerSpec.Image = img
			}
			// add platforms that are compatible with the service
			service.TaskTemplate.Placement = setServicePlatforms(service.TaskTemplate.Placement, distributionInspect)
		}
	}

>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	var response types.ServiceUpdateResponse
	resp, err := cli.post(ctx, "/services/"+serviceID+"/update", query, service, headers)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.body).Decode(&response)
<<<<<<< HEAD
=======

	if distErr != nil {
		response.Warnings = append(response.Warnings, digestWarning(service.TaskTemplate.ContainerSpec.Image))
	}

>>>>>>> c22478687a5c584b3f2f3b5d68ca7552a70385b2
	ensureReaderClosed(resp)
	return response, err
}
