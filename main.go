package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "github-release plugin"
	app.Usage = "github-release plugin"
	app.Action = run
	app.Version = version
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "api-key",
			Usage:  "api key to access github api",
			EnvVar: "PLUGIN_API_KEY,GITHUB_RELEASE_API_KEY,GITHUB_TOKEN,DRONE_NETRC_USERNAME",
		},
		cli.StringSliceFlag{
			Name:   "files",
			Usage:  "list of files to upload",
			EnvVar: "PLUGIN_FILES,GITHUB_RELEASE_FILES",
		},
		cli.StringFlag{
			Name:   "file-exists",
			Value:  "overwrite",
			Usage:  "what to do if file already exist",
			EnvVar: "PLUGIN_FILE_EXISTS,GITHUB_RELEASE_FILE_EXISTS",
		},
		cli.StringSliceFlag{
			Name:   "checksum",
			Usage:  "generate specific checksums",
			EnvVar: "PLUGIN_CHECKSUM,GITHUB_RELEASE_CHECKSUM",
		},
		cli.StringFlag{
			Name:   "checksum-file",
			Usage:  "name used for checksum file. \"CHECKSUM\" is replaced with the chosen method",
			EnvVar: "PLUGIN_CHECKSUM_FILE",
			Value:  "CHECKSUMsum.txt",
		},
		cli.BoolFlag{
			Name:   "checksum-flatten",
			Usage:  "include only the basename of the file in the checksum file",
			EnvVar: "PLUGIN_CHECKSUM_FLATTEN",
		},
		cli.BoolFlag{
			Name:   "draft",
			Usage:  "create a draft release",
			EnvVar: "PLUGIN_DRAFT,GITHUB_RELEASE_DRAFT",
		},
		cli.BoolFlag{
			Name:   "prerelease",
			Usage:  "set the release as prerelease",
			EnvVar: "PLUGIN_PRERELEASE,GITHUB_RELEASE_PRERELEASE",
		},
		cli.StringFlag{
			Name:   "base-url",
			Value:  "https://api.github.com/",
			Usage:  "api url, needs to be changed for ghe",
			EnvVar: "PLUGIN_BASE_URL,GITHUB_RELEASE_BASE_URL",
		},
		cli.StringFlag{
			Name:   "upload-url",
			Value:  "https://uploads.github.com/",
			Usage:  "upload url, needs to be changed for ghe",
			EnvVar: "PLUGIN_UPLOAD_URL,GITHUB_RELEASE_UPLOAD_URL",
		},
		cli.StringFlag{
			Name:   "note",
			Value:  "",
			Usage:  "file or string with notes for the release (example: changelog)",
			EnvVar: "PLUGIN_NOTE,GITHUB_RELEASE_NOTE",
		},
		cli.StringFlag{
			Name:   "title",
			Value:  "",
			Usage:  "file or string for the title shown in the github release",
			EnvVar: "PLUGIN_TITLE,GITHUB_RELEASE_TITLE",
		},
		cli.BoolFlag{
			Name:   "overwrite",
			Usage:  "force overwrite existing release informations e.g. title or note",
			EnvVar: "PLUGIN_OVERWRITE,GITHUB_RELEASE_OVERWRIDE",
		},
		cli.StringFlag{
			Name:   "repo.owner",
			Usage:  "repository owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "repo.name",
			Usage:  "repository name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "build.event",
			Value:  "push",
			Usage:  "build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.StringFlag{
			Name:   "commit.ref",
			Value:  "refs/heads/master",
			Usage:  "git commit ref",
			EnvVar: "DRONE_COMMIT_REF",
		},
		cli.StringFlag{
			Name:  "env-file",
			Usage: "source env file",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	if c.String("env-file") != "" {
		_ = godotenv.Load(c.String("env-file"))
	}

	plugin := Plugin{
		Repo: Repo{
			Owner: c.String("repo.owner"),
			Name:  c.String("repo.name"),
		},
		Build: Build{
			Event: c.String("build.event"),
		},
		Commit: Commit{
			Ref: c.String("commit.ref"),
		},
		Config: Config{
			APIKey:          c.String("api-key"),
			Files:           c.StringSlice("files"),
			FileExists:      c.String("file-exists"),
			Checksum:        c.StringSlice("checksum"),
			ChecksumFile:    c.String("checksum-file"),
			ChecksumFlatten: c.Bool("checksum-flatten"),
			Draft:           c.Bool("draft"),
			Prerelease:      c.Bool("prerelease"),
			BaseURL:         c.String("base-url"),
			UploadURL:       c.String("upload-url"),
			Title:           c.String("title"),
			Note:            c.String("note"),
			Overwrite:       c.Bool("overwrite"),
		},
	}

	return plugin.Exec()
}
