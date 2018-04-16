# Why use this?
* **Quick Deploy** - Build & Deploy your contract code in 1 command
* **Wallet Password Storage** - Automatically stores the passwords for all the wallets you've created
* **Auto Setup & Account Creation** - Create a default wallet, deploy the BIOS, & create an account in a single command
* **Quick Restart** - Restart the EOS blockchain from the genesis block, again, in 1 command

# Requirements
* eos (nodeos & cleos)
* MacOS

# Installation

`. setup && touch ~/.zeos/zeos.sh`

# QuickStart Guide w/ Example

Start nodeos

`nodeos`

In another tab:
```
PROJECT_NAME=hello
eos_account_create $PROJECT_NAME
eos_deploy $PROJECT_NAME
cleos push action $PROJECT_NAME hi '["hello"]' -p $PROJECT_NAME
```

# Notable Features

## Quick Deploy

From within your projects parent directory...

`eos_deploy $PROJECT_NAME`

Normally, you'd have to...
* `eosiocpp -o hello.wast hello.cpp` Build the .wast
* `eosiocpp -g hello.abi hello.cpp` Build the .abi
* `cleos set contract hello hello` Deploys the contract

## Wallet Password Storage

`eos_keys` Print all your current passwords

## Auto Setup & Account Creation

`eos_account_create $PROJECT_NAME`

What this does for you...
* `cleos wallet create -n default` Creates a default wallet, if none already exists
* Stores your wallet password (~/.zeos/keys.yml)
* `cleos set contract eosio $EOS_PATH"/build/contracts/eosio.bios" -p eosio` Deploys the BIOS, if not already
* `cleos create account eosio hello $KEY $KEY` Creates a contract account

## Quick Restart

`eos_reset`

Most of the time, to start an EOS instance, it's best to simply use `nodeos`. But, sometimes, you want to start from scratch.
In which case, this is the command for you.

# Full Command Reference

## Starting a nodeos process

* `nodeos` The traditional way to start a EOS blockchain instance
* `eos_clean` Deletes ALL existing blockchain data
* `eos_reset` Deletes ALL existing blockchain data, and starts a nodeos instance

## Creating a default wallet, and deploying the bios

* `eos_setup` Creates a default wallet, and deploys the EOS bios
* `eos_wallet_create -n $NAME` Creates a wallet, and stores the password

## Creating a code account

* `eos_account_create $PROJECT_NAME` Creates an account for your project contract

## Building & Deploying your contract

* `eos_deploy $PROJECT_NAME` Builds and deploys your contract code (**Make sure you're in the parent dir of your project**)
* `eos_build` Builds your contract wast & ABI (**Make sure you're in the dir of your project**)

## Print your stored passwords

* `eos_keys` Prints the entire password file (~/.zeos/keys.yml)
