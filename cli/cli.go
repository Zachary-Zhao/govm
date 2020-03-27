package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
	"govm/build"
)

var (
	rootDir      string
	downloadsDir string
	versionsDir  string
	goroot       string
)

// Run 运行g命令行
func Run() {
	app := cli.NewApp()
	app.Name = "govm"
	app.Usage = "Golang Version Manager"
	app.Version = build.Version()
	app.Copyright = "Copyright (c) 2019, Zachary. All rights reserved."
	app.Authors = []cli.Author{
		{
			Name:  "Zachary",
			Email: "gg19881120@126.com",
		},
	}
	app.Before = func(ctx *cli.Context) (err error) {
		homeDir, _ := os.UserHomeDir()
		rootDir = filepath.Join(homeDir, ".g")
		goroot = filepath.Join(rootDir, "go")
		downloadsDir = filepath.Join(rootDir, "downloads")
		if err = os.MkdirAll(downloadsDir, 0755); err != nil {
			return err
		}
		versionsDir = filepath.Join(rootDir, "versions")
		if err = os.MkdirAll(versionsDir, 0755); err != nil {
			return err
		}
		return nil
	}
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "[g] %s\n", err.Error())
		os.Exit(1)
	}
}

func init() {
	cli.AppHelpTemplate = fmt.Sprintf(`NAME:
	{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}
 
 USAGE:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .Commands}} command{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}
 
 VERSION:
	%s{{end}}{{end}}{{if .Description}}
 
 DESCRIPTION:
	{{.Description}}{{end}}{{if len .Authors}}
 
 AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
	{{range $index, $author := .Authors}}{{if $index}}
	{{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}
 
 COMMANDS:{{range .VisibleCategories}}{{if .Name}}
 
	{{.Name}}:{{end}}{{range .VisibleCommands}}
	  {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
 
 GLOBAL OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}
 
 COPYRIGHT:
	{{.Copyright}}{{end}}
`, build.ShortVersion)
}
