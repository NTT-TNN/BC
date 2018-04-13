docker exec -it chaincode bash
cd rfid
wget https://raw.githubusercontent.com/NTT-TNN/BC/master/rfid_chaincode.go
go build
CORE_PEER_ADDRESS=peer:7052 CORE_CHAINCODE_ID_NAME=mycc:0 ./rfid
