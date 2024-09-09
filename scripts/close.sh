#!/bin/bash

source ./env
provider-services tx deployment close --from $AKASH_KEY_NAME -y

sed -i "/export AKASH_DSEQ/d" "./env"
sed -i "/export AKASH_PROVIDER/d" "./env"

export AKASH_DSEQ=""
export AKASH_PROVIDER=""
