package cli

import "github.com/urfave/cli"

var (
	commands = []cli.Command{
		{
			Name:      "ls",
			Aliases:   []string{"list", "ll", "l"},
			Usage:     "List installed versions",
			UsageText: "g ls",
			Action:    list,
		},
		{
			Name:      "ls-remote",
			Usage:     "List remote versions available for install",
			UsageText: "g ls-remote [stable|archived]",
			Action:    listRemote,
		},
		{
			Name:      "use",
			Aliases:   []string{"switch"},
			Usage:     "Switch to specified version",
			UsageText: "g use <version>",
			Action:    use,
		},
		{
			Name:      "install",
			Aliases:   []string{"i"},
			Usage:     "Download and install a <version>",
			UsageText: "g install <version>",
			Action:    install,
		},
		{
			Name:      "uninstall",
			Aliases:   []string{"remove", "erase"},
			Usage:     "Uninstall a version",
			UsageText: "g uninstall <version>",
			Action:    uninstall,
		},
		{
			Name:      "download",
			Aliases:   []string{"dl"},
			Usage:     "Download a version",
			UsageText: "g download <version>",
			Action:    download,
		},
	}
)
