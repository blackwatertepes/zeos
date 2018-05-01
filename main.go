package main

import (
  "fmt"
  "os"
  "os/exec"

  "github.com/spf13/cobra"
)

func main() {
  var clean = func() {
    exec.Command("rm", os.Getenv("NODEOS_PATH") + "/config/genesis.json").Output()
    exec.Command("rm", "-rf", os.Getenv("NODEOS_PATH") + "/data").Output()
    fmt.Println("cleaned")
  }

  var start = func() {
    // TODO start multiple nodes in bg processes
    // TODO show nodeos output
    exec.Command("nodeos").Output()
    fmt.Println("started")
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

  var cmdWallet = &cobra.Command{
    Use:   "wallet",
    Short: "creates the default wallet",
    Args: cobra.MinimumNArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
      // TODO use the standard eos wallet
      if (len(args) == 0) {
        //kleos "default" $(cleos wallet create | grep PW)
        fmt.Println("wallet default created")
      } else {
        //kleos $1, $(cleos wallet create -n $1 | grep PW)
        fmt.Println("wallet", args[0], "created")
      }
    },
  }

  var cmdBoot = &cobra.Command{
    Use:   "boot",
    Short: "loads the BIOS",
    // TODO Remove the boot command in v2? So we no longer require the EOS_PATH
    Run: func(cmd *cobra.Command, args []string) {
      cmdWallet.Execute()
      //exec.Command("cleos", "set", "contract", "eosio", "$EOS_PATH/build/contracts/eosio.bios" "-p", "eosio").Output()
      fmt.Println("boot")
    },
  }

  // TODO Add to projects as cmdProjectCreate
  var cmdAccount = &cobra.Command{
    Use:   "account",
    Short: "creates a contract account",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      /*
      if ! (cleos wallet keys | grep EOS); then
        //cmdBoot.Execute()
      fi

      KEY=$(cleos wallet keys | grep EOS | cut -d '"' -f 2)
      cleos create account eosio $1 $KEY $KEY
      */
      fmt.Println("account", args[0], "created")
    },
  }

  // TODO Add to projects as cmdProjectBuild
  var cmdBuild = &cobra.Command{
    Use:   "build",
    Short: "builds a contract (wast & abi)",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      exec.Command("eosiocpp", "-o", "${args[0]}.wast", "${args[0]}.cpp").Output()
      exec.Command("eosiocpp", "-g", "${args[0]}.abi", "${args[0]}.cpp").Output()
      fmt.Println("built", args[0])
    },
  }

  // TODO add to projects as cmdProjectDeploy
  var cmdDeploy = &cobra.Command{
    Use:   "deploy",
    Short: "builds & deploys a contract",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      exec.Command("cd", args[0]).Output()
      cmdBuild.Execute()
      exec.Command("cd", "..").Output()
      exec.Command("cleos", "set", "contract", args[0], args[0]).Output()
      fmt.Println("deployed", args[0])
    },
  }

  // TODO Store a list of all the projects
  // TODO Store a list of all contracts & accounts for a given projects
  var cmdProject = &cobra.Command{
    Use:   "project",
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("project")
    },
  }

  var rootCmd = &cobra.Command{Use: "zeos"}
  rootCmd.AddCommand(cmdClean, cmdStart, cmdReset, cmdWallet, cmdBoot, cmdAccount, cmdBuild, cmdDeploy, cmdProject)
  //cmdProject.AddCommand(cmdProjectCreate, cmdProjectList, cmdProjectDelete, cmdProjectClean, cmdProjectBuild, cmdProjectDeploy)
  rootCmd.Execute()
}
