package helpers

import (
	"context"

	"github.com/BharathKumarRavichandran/cloud99/utils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func CreateContainer(image string) (string, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		utils.Logger.Errorf("Unable to create docker client: %s", err.Error())
		panic(err)
	}

	/* Type signature function for ContainerCreate
	ContainerCreate func(
		ctx context.Context,
		config *container.Config,
		hostConfig *container.HostConfig,
		networkingConfig *network.NetworkingConfig,
		containerName string) (container.ContainerCreateCreatedBody, error)

	*/
	cont, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: image,
			Cmd:   []string{"/bin/bash"},
		},
		nil, nil, "")
	if err != nil {
		utils.Logger.Errorf("Unable to create container: %s", err.Error())
		panic(err)
	}

	utils.Logger.Infof("Created container: %s", cont.ID)
	return cont.ID, nil
}

func StartContainer(containerID string) (string, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		utils.Logger.Errorf("Unable to create docker client: %s", err.Error())
		return containerID, err
	}

	cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	utils.Logger.Infof("Started container: %s", containerID)
	return containerID, nil
}

func StopContainer(containerID string) (string, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		utils.Logger.Errorf("Unable to create docker client: %s", err.Error())
		panic(err)
	}

	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		utils.Logger.Errorf("Cannot stop container: %s", err.Error())
		panic(err)
	}

	utils.Logger.Infof("Stopped container: %s", containerID)
	return containerID, err
}

func DeleteContainer(containerID string) (string, error) {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		utils.Logger.Errorf("Unable to create docker client: %s", err.Error())
		panic(err)
	}

	err = cli.ContainerStop(ctx, containerID, nil)
	if err != nil {
		utils.Logger.Errorf("Cannot stop container: %s", err.Error())
		panic(err)
	}

	utils.Logger.Infof("Deleted container: %s", containerID)
	return containerID, err
}
