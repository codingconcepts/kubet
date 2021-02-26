package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	//go:embed templates/deployment.yaml.tmpl
	deployment         string
	deploymentTemplate = template.Must(template.New("deployment").Parse(deployment))

	//go:embed templates/service.yaml.tmpl
	service         string
	serviceTemplate = template.Must(template.New("service").Parse(service))

	//go:embed templates/namespace.yaml.tmpl
	namespace         string
	namespaceTemplate = template.Must(template.New("namespace").Parse(namespace))
)

func main() {
	rootCmd := &cobra.Command{}

	namespaceCmd := &cobra.Command{
		Use:   "namespace",
		Short: "Create a namespace manifest template.",
		Run: func(cmd *cobra.Command, args []string) {
			run(namespaceTemplate, args)
		},
	}

	deploymentCmd := &cobra.Command{
		Use:   "deployment",
		Short: "Create a deployment manifest template.",
		Run: func(cmd *cobra.Command, args []string) {
			run(deploymentTemplate, args)
		},
	}

	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "Create a service manifest template.",
		Run: func(cmd *cobra.Command, args []string) {
			run(serviceTemplate, args)
		},
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

func run(t *template.Template, args []string) {
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
