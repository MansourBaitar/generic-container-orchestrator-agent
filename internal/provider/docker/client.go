package docker

import (
	"dsync.io/gco/agent/internal/log"
	docker "github.com/docker/docker/client"
	"os"
)

func newDockerClient() docker.APIClient {
	//client, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	//if err != nil {
	//	log.Errorf("Unable to create new docker client: %v", err)
	//	os.Exit(1)
	//}

	cli, err := docker.NewClientWithOpts(docker.FromEnv)
	if err != nil {
		log.Errorf("Unable to create new docker client: %v", err)
		os.Exit(1)
	}

	return cli
}