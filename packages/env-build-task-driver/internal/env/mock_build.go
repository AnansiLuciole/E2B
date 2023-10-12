package env

import (
	"context"

	"github.com/docker/docker/client"
	docker "github.com/fsouza/go-dockerclient"
	"go.opentelemetry.io/otel"
)

func MockBuild(envID, buildID string) {
	ctx := context.Background()

	tracer := otel.Tracer("test")

	client, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	legacyClient, err := docker.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	contextsPath := "/mnt/disks/docker-contexts/v1"
	registry := "us-central1-docker.pkg.dev/e2b-prod/custom-environments"
	envsDisk := "/mnt/disks/fc-envs/v1"
	kernelImagePath := "/fc-vm/vmlinux.bin"
	firecrackerBinaryPath := "/usr/bin/firecracker"
	envdPath := "/fc-vm/envd"
	contextFileName := "context.tar.gz"
	vCPUCount := int64(1)
	memoryMB := int64(512)
	diskSizeMB := int64(512)

	e := Env{
		BuildID:               buildID,
		EnvID:                 envID,
		EnvsDiskPath:          envsDisk,
		VCpuCount:             vCPUCount,
		MemoryMB:              memoryMB,
		DockerContextsPath:    contextsPath,
		DockerRegistry:        registry,
		KernelImagePath:       kernelImagePath,
		DiskSizeMB:            diskSizeMB,
		FirecrackerBinaryPath: firecrackerBinaryPath,
		EnvdPath:              envdPath,
		ContextFileName:       contextFileName,
	}

	err = e.Build(ctx, tracer, client, legacyClient)
	if err != nil {
		panic(err)
	}
}