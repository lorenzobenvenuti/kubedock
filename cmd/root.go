package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubedock",
	Short: "Kubedock is a docker api to orchestrate containers on kubernetes.",
	Long: `Kubedock is a docker api to orchestrate containers on kubernetes.

  Kubedock is an minimal implementation of the docker api that will orchestrate
  containers on a kubernetes cluster, rather than running containers locally.
  Each container is considered to be short-lived and emphemeral, with the
  intention they are solely used for running CI tests that require docker
  containers inside a container. This enables running the tests without the
  requirement of running docker-in-docker within resource heavy containers.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
