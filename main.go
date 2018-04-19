package main

import (
  "fmt"
  //"strings"

  "github.com/spf13/cobra"
)

func main() {
  var cmdClean = &cobra.Command{
    Use:   "clean",
    Short: "deletes the blockchain data",
    //Long: `print is for printing anything back to the screen.`
    Run: func(cmd *cobra.Command, args []string) {
      //rm "$NODEOS_PATH/config/genesis.json"
      //rm -rf "$NODEOS_PATH/data"
      fmt.Println("cleaned")
    },
  }

  var cmdStart = &cobra.Command{
    User:   "start",
    Short:  "starts a # of nodeos block producers",
    Run: func(cmd *cobra.Command, args []string) {
      //nodeos
      // TODO start multiple nodes in bg processes
      fmt.Println("start")
    },
  }

  var cmdReset = &cobra.Command{
    Use:   "reset",
    Short: "deletes the blockchain data & starts nodeos",
    Run: func(cmd *cobra.Command, args []string) {
      //cmdClean.Execute()
      //cmdStart.Execute()
      fmt.Println("reset")
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
    Run: func(cmd *cobra.Command, args []string) {
      //cmdWallet.Execute()
      //cleos set contract eosio "$EOS_PATH/build/contracts/eosio.bios" -p eosio
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
      //eosiocpp -o $1.wast $1.cpp
      //eosiocpp -g $1.abi $1.cpp
      fmt.Println("built", args[0])
    },
  }

  // TODO add to projects as cmdProjectDeploy
  var cmdDeploy = &cobra.Command{
    Use:   "deploy",
    Short: "builds & deploys a contract",
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      /*
      cd $1
      //cmdBuild.Execute()
      cd ..
      cleos set contract $1 $1
      */
      fmt.Println("deployed", args[0])
    },
  }

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
