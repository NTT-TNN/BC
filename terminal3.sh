docker exec -it cli bash
peer chaincode install -p chaincodedev/chaincode/rfid -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"Args":["1234","Thao"]}' -C myc

peer chaincode invoke -n mycc -c '{"Args":["createUser", "Thao", "a","b","c"]}' -C myc


peer chaincode instantiate  -C myc -n marbles -v 1.0 -c '{"Args":["init"]}' 
