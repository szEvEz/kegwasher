package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "print only the version",
	}
	app := &cli.App{
		Name:    "kegwasher",
		Version: "v0.0.2",
		Usage:   "Housekeeping for Homebrew",
		Action:  cleanupAction,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "noprune",
				Usage: "Do not remove all cached files",
			},
			&cli.BoolFlag{
				Name:  "noupdate",
				Usage: "Do not run 'brew update' and 'brew upgrade' before cleanup",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func cleanupAction(c *cli.Context) error {
	fmt.Println("Washing the kegs...")

	// Run Homebrew update and upgrade if flag is enabled
	if !c.Bool("noupdate") {
		fmt.Println("Updating Homebrew...")
		cmdUpdate := exec.Command("brew", "update")
		cmdUpdate.Stdout = os.Stdout
		cmdUpdate.Stderr = os.Stderr

		err := cmdUpdate.Run()
		if err != nil {
			return fmt.Errorf("Error running 'brew update': %v", err)
		}

		fmt.Println("Upgrading installed formulae...")
		cmdUpgrade := exec.Command("brew", "upgrade")
		cmdUpgrade.Stdout = os.Stdout
		cmdUpgrade.Stderr = os.Stderr

		err = cmdUpgrade.Run()
		if err != nil {
			return fmt.Errorf("Error running 'brew upgrade': %v", err)
		}
	}

	args := []string{"cleanup", "--prune=all"}
	if c.Bool("noprune") {
		args = []string{"cleanup"}
	}

	// Run Homebrew cleanup
	cmd := exec.Command("brew", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Error cleaning up Homebrew: %v", err)
	}

	// Run Homebrew autoremove
	fmt.Println("Scraping leftovers...")
	cmdAutoRemove := exec.Command("brew", "autoremove")
	cmdAutoRemove.Stdout = os.Stdout
	cmdAutoRemove.Stderr = os.Stderr

	err = cmdAutoRemove.Run()
	if err != nil {
		return fmt.Errorf("Error running 'brew autoremove': %v", err)
	}

	fmt.Println("Keg washing completed...")
	return nil
}
