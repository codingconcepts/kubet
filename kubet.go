package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{}

	namespaceCmd := &cobra.Command{
		Use:   "namespace",
		Short: "Create a namespace manifest template.",
		Run:   runNamespace,
	}

	deploymentCmd := &cobra.Command{
		Use:   "deployment",
		Short: "Create a deployment manifest template.",
		Run:   runDeployment,
	}

	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Create a service manifest template.",
		Run:   runService,
	}

	rootCmd.AddCommand(
		namespaceCmd,
		deploymentCmd,
		serviceCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error executing command: %v", err)
	}
}

func runNamespace(cmd *cobra.Command, args []string) {
	run("templates/namespace.yaml.tmpl", args)
}

func runDeployment(cmd *cobra.Command, args []string) {
	run("templates/deployment.yaml.tmpl", args)
}

func runService(cmd *cobra.Command, args []string) {
	run("templates/service.yaml.tmpl", args)
}

func run(path string, args []string) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Fatalf("error parsing template: %v", err)
	}

	buf := &bytes.Buffer{}
	if err := t.Execute(buf, argsToMap(args)); err != nil {
		log.Fatalf("error running template: %v", err)
	}

	fmt.Println(buf.String())
}

func argsToMap(args []string) map[string]string {
	m := map[string]string{}
	for _, a := range args {
		parts := strings.Split(a, "=")
		m[parts[0]] = parts[1]
	}

	return m
}
