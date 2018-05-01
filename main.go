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

  var wallet = func(name string) {
    // TODO use the standard eos wallet
    //kleos $1, $(cleos wallet create -n $1 | grep PW)
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
    exec.Command("cleos", "set", "contract", "eosio", os.Getenv("EOS_PATH") + "/build/contracts/eosio.bios", "-p", "eosio").Output()
    fmt.Println("boot")
  }

  var cmdBoot = &cobra.Command{
    Use:   "boot",
    Short: "loads the BIOS",
    // TODO Remove the boot command in v2? So we no longer require the EOS_PATH
    Run: func(cmd *cobra.Command, args []string) {
      boot()
    },
  }

  var account = func(name string) {
    /*
    if ! (cleos wallet keys | grep EOS); then
      //cmdBoot.Execute()
    fi

    KEY=$(cleos wallet keys | grep EOS | cut -d '"' -f 2)
    cleos create account eosio $1 $KEY $KEY
    */
    fmt.Println("account", name, "created")
  }

  // TODO Add to projects as cmdProjectCreate
  var cmdAccount = &cobra.Command{
    Use:   "account",
    Short: "creates a contract account",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      account(args[0])
    },
  }

  var build = func(name string) {
    exec.Command("eosiocpp", "-o", "${name}.wast", "${name}.cpp").Output()
    exec.Command("eosiocpp", "-g", "${name}.abi", "${name}.cpp").Output()
    fmt.Println("built", name)
  }

  // TODO Add to projects as cmdProjectBuild
  var cmdBuild = &cobra.Command{
    Use:   "build",
    Short: "builds a contract (wast & abi)",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      build(args[0])
    },
  }

  var deploy = func(name string) {
    exec.Command("cd", name).Output()
    cmdBuild.Execute()
    exec.Command("cd", "..").Output()
    exec.Command("cleos", "set", "contract", name, name).Output()
    fmt.Println("deployed", name)
  }

  // TODO add to projects as cmdProjectDeploy
  var cmdDeploy = &cobra.Command{
    Use:   "deploy",
    Short: "builds & deploys a contract",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      deploy(args[0])
    },
  }

  var project = func() {
    fmt.Println("project")
  }

  // TODO Store a list of all the projects
  // TODO Store a list of all contracts & accounts for a given projects
  var cmdProject = &cobra.Command{
    Use:   "project",
    Run: func(cmd *cobra.Command, args []string) {
      project()
    },
  }

  var rootCmd = &cobra.Command{Use: "zeos"}
  rootCmd.AddCommand(cmdClean, cmdStart, cmdReset, cmdWallet, cmdBoot, cmdAccount, cmdBuild, cmdDeploy, cmdProject)
  //cmdProject.AddCommand(cmdProjectCreate, cmdProjectList, cmdProjectDelete, cmdProjectClean, cmdProjectBuild, cmdProjectDeploy)
  rootCmd.Execute()
}
