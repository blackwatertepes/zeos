# Requirements
* eos (nodeos & cleos)
* MacOS

# Installation

`. setup`

# QuickStart Guide

```
eos_reset
eos_setup
eos_account_create $PROJECT_NAME
eos_deploy $PROJECT_NAME
cleos push action $PROJECT_NAME $METHOD $DATA -p $PROJECT_NAME
```

# Command Reference

## Starting a nodeos process

* `nodeos` The tradition way to start a EOS blockchain instance
* `eos_clean` Deletes ALL existing blockchain data
* `eos_reset` Deletes ALL existing blockchain data, and starts a nodeos instance

## Creating a default wallet, and deploying the bios

* `eos_setup` Creates a default wallet, and deploys the EOS bios
* `eos_wallet_create -n $NAME` Creates a wallet, and stores the password

## Creating a code account

* `eos_account_create $PROJECT_NAME` Creates an account for your project contract

## Building & Deploying your contract

`eos_deploy $PROJECT_NAME` Builds and deploys your contract code (**Make sure you're in the parent dir of your project**)
`eos_build` Builds your contract wast & ABI (**Make sure you're in the dir of your project**)
