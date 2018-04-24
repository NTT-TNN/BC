
#!/usr/bin/env bash

# Copyright IBM Corp., All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

#This is a one step script to delete all the deployment and
#services executed during the execution of the cluster
echo "Delete ca"
kubectl delete -f ./manifests/fabric-1.0/local/ca.yaml
sleep 5

echo "Delete orderer"
kubectl delete -f ./manifests/fabric-1.0/local/orderer.yaml
sleep 5


echo "Delete Peer0"
kubectl delete -f ./manifests/fabric-1.0/local/peer0.yaml
sleep 5

echo "Delete rest of the Peers"
kubectl delete -f./manifests/fabric-1.0/local/peer1.yaml -f ./manifests/fabric-1.0/local/peer2.yaml -f ./manifests/fabric-1.0/local/peer3.yaml

sleep 5

echo "Delete Cli"
kubectl delete -f ./manifests/fabric-1.0/local/cli.yaml

echo "**********Deployment done successfully**********"
