package config

import (
	"io"
	"text/tabwriter"

	"github.com/kelseyhightower/envconfig"
)

var Cfg Server

type Server struct {
	HttpAddr string `envconfig:"HTTP_ADDR" default:":8077"`
	GrpcAddr string `envconfig:"GRPC_ADDR" default:":8078"`
}

func Config() (err error) {
	if err = envconfig.Process("", &Cfg); err != nil {
		return err
	}
	return nil
}

func PrintUsage(output io.Writer) {
	printUsageHelper(&Cfg, output)
}

func printUsageHelper(spec interface{}, output io.Writer) {
	var DefaultTableFormat = `KEY	TYPE	DEFAULT	REQUIRED	DESCRIPTION
{{range .}}{{usage_key .}}	{{usage_type .}}	{{usage_default .}}	{{usage_required .}}	{{usage_description .}}
{{end}}`

	tabs := tabwriter.NewWriter(output, 1, 0, 4, ' ', 0)
	envconfig.Usagef("", spec, tabs, DefaultTableFormat)
	tabs.Flush()
}
