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
    // TODO use the standard eos wallet
    execCmdStream(exec.Command("kleos", name, "$(cleos wallet create -n " + name + " | grep PW)"))
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
    // TODO Remove the boot command in v2? So we no longer require the EOS_PATH
    Run: func(cmd *cobra.Command, args []string) {
      boot()
    },
  }

  var account = func(name string) {
    /*
    if ! (cleos wallet keys | grep EOS); then
      boot()
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

  var cwd = func() string {
    out, err := exec.Command("pwd").Output()
    if (err != nil) {
      fmt.Println(err)
    }
    pwd := string(out)
    dirs := strings.Split(pwd, "/")
    return dirs[len(dirs) - 1]
  }

  var build = func() {
    cwd := cwd()
    execCmdStream(exec.Command("eosiocpp", "-o", cwd + ".wast", cwd + ".cpp"))
    execCmdStream(exec.Command("eosiocpp", "-g", cwd + ".abi", cwd + ".cpp"))
    fmt.Println("built", cwd)
  }

  // TODO Add to projects as cmdProjectBuild
  var cmdBuild = &cobra.Command{
    Use:   "build",
    Short: "builds a contract (wast & abi)",
    Args: cobra.MinimumNArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
      build()
    },
  }

  var deploy = func(name string) {
    exec.Command("cd", name).Run()
    build()
    exec.Command("cd", "..").Run()
    execCmdStream(exec.Command("cleos", "set", "contract", name, name))
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
