export KUBECONFIG=/etc/kubernetes/admin.conf

CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/msp
peer chaincode install -p . -n mycc -v 0


peer chaincode install -n mycc -v 1.0 -p .

peer chaincode instantiate -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"
