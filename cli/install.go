package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mholt/archiver"
	"github.com/urfave/cli"
)

func install(ctx *cli.Context) (err error) {
	vname := ctx.Args().First()
	if vname == "" {
		return cli.ShowSubcommandHelp(ctx)
	}

	// 下载安装包
	download(ctx)
	// 解压安装包
	if err = archiver.Unarchive(filename, versionsDir); err != nil {
		return cli.NewExitError(fmt.Sprintf("[g] %s", err.Error()), 1)
	}
	// 目录重命名
	if err = os.Rename(filepath.Join(versionsDir, "go"), targetV); err != nil {
		return cli.NewExitError(fmt.Sprintf("[g] %s", err.Error()), 1)
	}
	// 重新建立软链接
	_ = os.Remove(goroot)

	if err := os.Symlink(targetV, goroot); err != nil {
		return cli.NewExitError(fmt.Sprintf("[g] %s", err.Error()), 1)
	}
	fmt.Println("Installed successfully")
	return nil
}
