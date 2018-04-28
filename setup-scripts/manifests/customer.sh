export KUBECONFIG=/etc/kubernetes/admin.conf

CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/fabric/msp
peer chaincode install -p . -n mycc -v 0


peer chaincode install -n mycc -v 1.0 -p .

peer chaincode instantiate -o orderer.example.com:7050 --tls $CORE_PEER_TLS_ENABLED --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n mycc -v 1.0 -c '{"Args":["init","a", "100", "b","200"]}' -P "OR ('Org1MSP.member','Org2MSP.member')"

ssh -i /Volumes/HOC_TAP/HPC/CLOUD/thao-ubuntu-1.pem ubuntu@25.29.133.43
ssh -i /Volumes/HOC_TAP/HPC/CLOUD/thao-ubuntu-1.pem ubuntu@25.29.138.183
ssh -i /Volumes/HOC_TAP/HPC/CLOUD/thao-ubuntu-1.pem ubuntu@25.29.139.169


  kubeadm join 192.168.160.15:6443 --token 59i4bp.pebf8ay1xy4sbdhw --discovery-token-ca-cert-hash sha256:47381189016819ac2ce2aa53ff577b23d44afcb5f9c0722012f92d3d8aaaa3e9


CORE_PEER_ADDRESS=peer:7051 CORE_CHAINCODE_ID_NAME=mycc:0 ./chaincode_example02
