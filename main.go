package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"

  "github.com/spf13/cobra"
)

func main() {
  var execCmdStream = func(cmd *exec.Cmd) {
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
  }

  var clean = func() {
    execCmdStream(exec.Command("rm", os.Getenv("NODEOS_PATH") + "/config/genesis.json"))
    execCmdStream(exec.Command("rm", "-rf", os.Getenv("NODEOS_PATH") + "/data"))
    fmt.Println("cleaned")
  }

  var start = func() {
    // TODO start multiple nodes in bg processes
    execCmdStream(exec.Command("nodeos"))
  }

  var cmdClean = &cobra.Command{
    Use:   "clean",
    Short: "deletes the blockchain data",
    Run: func(cmd *cobra.Command, args []string) {
      clean()
    },
  }

  var cmdStart = &cobra.Command{
    Use:   "start",
    Short:  "starts a # of nodeos block producers",
    Run: func(cmd *cobra.Command, args []string) {
      start()
    },
  }

  var cmdReset = &cobra.Command{
    Use:   "reset",
    Short: "deletes the blockchain data & starts nodeos",
    Run: func(cmd *cobra.Command, args []string) {
      clean()
      start()
    },
  }

  var wallet = func(name string) {
    // TODO Store the wallet password
    execCmdStream(exec.Command("$(cleos wallet create -n " + name))
    fmt.Println("wallet", name, "created")
  }

  var cmdWallet = &cobra.Command{
    Use:   "wallet",
    Short: "creates the default wallet",
    Args: cobra.MinimumNArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
      if (len(args) == 0) {
        wallet("default")
      } else {
        wallet(args[0])
      }
    },
  }

  var boot = func() {
    wallet("default")
    execCmdStream(exec.Command("cleos", "set", "contract", "eosio", os.Getenv("EOS_PATH") + "/build/contracts/eosio.bios", "-p", "eosio"))
    fmt.Println("booted")
  }

  var cmdBoot = &cobra.Command{
    Use:   "boot",
    Short: "loads the BIOS",
    // TODO Prompt for the EOS_PATH to be set, if it's not
    Run: func(cmd *cobra.Command, args []string) {
      boot()
    },
  }

  var rootCmd = &cobra.Command{Use: "zeos"}
  rootCmd.AddCommand(cmdClean, cmdStart, cmdReset, cmdWallet, cmdBoot, cmdAccount, cmdBuild, cmdDeploy, cmdProject)
  rootCmd.Execute()
}
