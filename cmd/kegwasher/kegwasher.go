package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"

    "github.com/urfave/cli/v2"
)

func main() {
    cli.VersionFlag = &cli.BoolFlag{
        Name:    "version",
        Aliases: []string{"v"},
        Usage:   "print only the version",
    }
    app := &cli.App{
        Name:  "kegwasher",
        Version: "v0.0.1",
        Usage: "Housekeeping for Homebrew",
		Action:  cleanupAction,
        Flags: []cli.Flag{
            &cli.BoolFlag{
                Name:  "prune",
                Usage: "Remove all cache files",
            },
            &cli.BoolFlag{
                Name:  "update",
                Usage: "Run 'brew update' and 'brew upgrade' before cleanup",
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
	if c.Bool("update") {
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

	args := []string{"cleanup"}
	if c.Bool("prune") {
		args = append(args, "--prune=all")
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
    fmt.Println("Autoremove...")
	cmdAutoRemove := exec.Command("brew", "autoremove")
	cmdAutoRemove.Stdout = os.Stdout
	cmdAutoRemove.Stderr = os.Stderr

	err = cmdAutoRemove.Run()
	if err != nil {
		return fmt.Errorf("Error running 'brew autoremove': %v", err)
	}

	fmt.Println("Keg washing completed..")
	return nil
}
