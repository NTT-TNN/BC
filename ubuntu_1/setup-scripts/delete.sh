
#!/usr/bin/env bash

# Copyright IBM Corp., All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#comment this if you don't want to use the default CHANNEL_NAME "mychannel"
#to delete your channel, mannually run:

# bash generateArtifacts.sh <CHANNEL_NAME>
#bash driving-files/generateArtifacts.sh

#bash driving-files/prepare-files.sh


echo "Deploying ca"
kubectl delete -f ./manifests/fabric-1.0/local/ca.yaml
sleep 5

echo "Deploying orderer"
kubectl delete -f ./manifests/fabric-1.0/local/orderer.yaml
sleep 5


echo "Deploying Peer0"
kubectl delete -f ./manifests/fabric-1.0/local/peer0.yaml
sleep 5

echo "Deploying rest of the Peers"
kubectl delete -f./manifests/fabric-1.0/local/peer1.yaml -f ./manifests/fabric-1.0/local/peer2.yaml -f ./manifests/fabric-1.0/local/peer3.yaml

sleep 5

echo "Deploying Cli"
kubectl delete -f ./manifests/fabric-1.0/local/cli.yaml

echo "**********Deployment done successfully**********"
