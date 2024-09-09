#!/bin/bash

sudo apt install jq -y

sudo apt install unzip -y

sudo curl -sfL https://raw.githubusercontent.com/akash-network/provider/main/install.sh | bash

echo "export PATH=$PATH:~/bin" >> ~/.bashrc
export PATH=$PATH:~/bin

rm ./env
export AKASH_KEY_NAME=test_sandbox
export AKASH_KEYRING_BACKEND=os
provider-services keys add $AKASH_KEY_NAME
export AKASH_ACCOUNT_ADDRESS="$(provider-services keys show $AKASH_KEY_NAME -a)"

echo $AKASH_ACCOUNT_ADDRESS

export AKASH_NET="https://raw.githubusercontent.com/akash-network/net/main/sandbox"

export AKASH_VERSION="$(curl -s https://api.github.com/repos/akash-network/provider/releases/latest | jq -r '.tag_name')"

export AKASH_CHAIN_ID="$(curl -s "$AKASH_NET/chain-id.txt")"

export AKASH_NODE="$(curl -s "$AKASH_NET/rpc-nodes.txt" | shuf -n 1)"

echo $AKASH_NODE $AKASH_CHAIN_ID $AKASH_KEYRING_BACKEND

export AKASH_GAS=1000000
export AKASH_GAS_ADJUSTMENT=1.25
export AKASH_GAS_PRICES=0.025uakt
export AKASH_SIGN_MODE=amino-json

provider-services query bank balances --node $AKASH_NODE $AKASH_ACCOUNT_ADDRESS

provider-services tx cert generate client --from $AKASH_KEY_NAME
provider-services tx cert publish client --from $AKASH_KEY_NAME

echo "export AKASH_KEY_NAME="$AKASH_KEY_NAME";" >> ./env
echo "export AKASH_KEYRING_BACKEND="$AKASH_KEYRING_BACKEND";" >> ./env
echo "export AKASH_ACCOUNT_ADDRESS="$AKASH_ACCOUNT_ADDRESS";" >> ./env
echo "export AKASH_NET="$AKASH_NET";" >> ./env
echo "export AKASH_VERSION="$AKASH_VERSION";" >> ./env
echo "export AKASH_CHAIN_ID="$AKASH_CHAIN_ID";" >> ./env
echo "export AKASH_NODE="$AKASH_NODE";" >> ./env
echo "export AKASH_GAS="$AKASH_GAS";" >> ./env
echo "export AKASH_GAS_ADJUSTMENT="$AKASH_GAS_ADJUSTMENT";" >> ./env
echo "export AKASH_GAS_PRICES="$AKASH_GAS_PRICES";" >> ./env
echo "export AKASH_SIGN_MODE="$AKASH_SIGN_MODE";" >> ./env


echo ""
echo "top up your balance at https://faucet.sandbox-01.aksh.pw/"
echo "AKASH_ACCOUNT_ADDRESS: " $AKASH_ACCOUNT_ADDRESS
echo "press any key"

stty -echo
read -n 1
stty echo
provider-services query bank balances --node $AKASH_NODE $AKASH_ACCOUNT_ADDRESS

