eos_clean() {
  rm "$NODEOS_PATH/config/genesis.json"
  rm -rf "$NODEOS_PATH/data"
  echo "" > ~/.zeos/keys.yml
}

eos_reset() {
  eos_clean
  nodeos
}

eos_wallet_save() {
  echo "$1=$2" >> ~/.zeos/keys.yml
  echo "Wallet Saved: $1 : $2"
}

eos_wallet_create() {
  if [ $1 ]
  then
    echo "Creating Wallet: $1"
    eos_wallet_save $1, $(cleos wallet create -n $1 | grep PW)
  else
    echo "Creating Default Wallet"
    eos_wallet_save "default" $(cleos wallet create | grep PW)
  fi
}

eos_setup() {
  eos_wallet_create
  cleos set contract eosio "$EOS_PATH/build/contracts/eosio.bios" -p eosio
}

eos_account_create() {
  if ! (cleos wallet keys | grep EOS); then
    eos_setup
  fi

  KEY=$(cleos wallet keys | grep EOS | cut -d '"' -f 2)
  cleos create account eosio $1 $KEY $KEY
}

eos_build() {
  eosiocpp -o $1.wast $1.cpp
  eosiocpp -g $1.abi $1.cpp
}

eos_deploy() {
  cd $1
  eos_build $1
  cd ..
  cleos set contract $1 $1
}

eos_keys() {
  cat ~/.zeos/keys.yml
}
