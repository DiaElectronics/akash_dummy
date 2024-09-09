#!/bin/bash

source ./env
echo create...
export AKASH_DSEQ=`provider-services tx deployment create deploy.yml --from $AKASH_KEY_NAME -y | jq -r '[.logs[] | .events[]| .attributes[] | select(.key == "dseq")| .value][0]'`
echo "dseg: "$AKASH_DSEQ
AKASH_OSEQ=1
AKASH_GSEQ=1

echo ""
echo "sleep 30"
sleep 30

echo bid...
export PROVIDER=`provider-services query market bid list --owner=$AKASH_ACCOUNT_ADDRESS --node $AKASH_NODE --dseq $AKASH_DSEQ --state=open`
echo "provider: " $PROVIDER
export AKASH_PROVIDER=`echo $PROVIDER | grep -oP '(provider: )\K.*'| cut -d' ' -f1`
echo "AKASH_PROVIDER: "$AKASH_PROVIDER

echo ""
echo create lease...
provider-services tx market lease create --dseq $AKASH_DSEQ --provider $AKASH_PROVIDER --from $AKASH_KEY_NAME -y

echo ""
echo lease list...
provider-services query market lease list --owner $AKASH_ACCOUNT_ADDRESS --node $AKASH_NODE --dseq $AKASH_DSEQ

echo ""
echo send manifest...
provider-services send-manifest deploy.yml --dseq $AKASH_DSEQ --provider $AKASH_PROVIDER --from $AKASH_KEY_NAME

echo ""
echo lease-status...
provider-services lease-status --dseq $AKASH_DSEQ --from $AKASH_KEY_NAME --provider $AKASH_PROVIDER

echo "export AKASH_DSEQ="$AKASH_DSEQ";" >> ./env
echo "export AKASH_PROVIDER="$AKASH_PROVIDER";" >> ./env

