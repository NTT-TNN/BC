docker exec -it cli bash

peer chaincode install -p chaincodedev/chaincode/rfid -n rfid -v 1
peer chaincode instantiate -n rfid -v 0  -c '{"Args":["a","10"]}' -C myc
peer chaincode invoke -n rfid -v 0 -c '{"Args":["queryUser","thao"]}' -C myc

old: 
    peer chaincode invoke -n rfid -v 0 -c '{"Args":["createUser","USER1","id","male","99","0","0","123"]}' -C myc
update: 
    peer chaincode invoke -n rfid -v 0 -c '{"Args":["createUser","thao","id","male","99","0","0"]}' -C myc


peer chaincode invoke -n rfid -v 0 -c '{"Args":["queryAllUsers"]}' -C myc
peer chaincode invoke -n rfid -v 0 -c '{"Args":["In","thao","123"]}' -C myc
peer chaincode invoke -n rfid -v 0 -c '{"Args":["Out","thao","234"]}' -C myc