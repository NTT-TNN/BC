# BlockChain Overview

## 1.Introdution to blockchain

### 1.1. Definition of blockchain

Blockchain bản chất là một hệ thống phân tán peer to peer của ledger cái mà được mã hoá, chỉ có thể mở rộng mà không thể chỉnh sửa, cực kỳ khó để thay đổi(gần như không thể) và chỉ có thể cập nhật thông quả giải thuật đồng thuận. 

Từ góc nhìn kinh doanh, blockchain được định nghĩa như một platform nơi mà các peer có thể tiến hành trao đổi giá trị mà không cần một người tin cậy trung gian. 

#### 1.1.1. Generic elements of a blockchain(Ví dụ trong ethereum blockchain)

- Address: là một giá trị duy nhất được sử dụng trong 1 giáo dịch để xác định người gửi là người nhận.
- Transaction: là đợn vị cơ bản của một block. Thể hiện một quá trình chuyền giá trị từ một địa chỉ đến một địa chỉ khác.
- Block: Một block là một tập của các giao dịch(transaction) được nhóm lại với nhau để tổ chức chúng một cách logic, ngoài ra còn có một số thông tin khác như timestamp, nonce. Kích thước của một block thì được định nghĩa tuỳ theo mục địch sử dụng của blockchain. 

Cấu trúc của một block:

    - BlockHeader: Lưu trữ metadata về block. 
        - The previous block hash.
        - Mining competition. Mỗi block cần có một vaid hash. Bao gồm: timestamp, noncev và độ khó( với bitcoin)
        - Merkle tree root: Cấu trúc dữ liệu để tổng kết các transaction trong block.
        - Block identifiers: Mỗi một block sẽ có một mã băm để xác định block đó.
        - Merkle tree root: Là một cáu trúc dữ liệu có được bằng cách băm một cách đệ quy một cặp các transaction cho đến khi chỉ còn một giá trị băm gọi là merkle root. Giá trị root này sẽ được lưu tại blockchain header. 
    - Danh sách các transactions.

[image]

- Blockchain: Là một tập các block được nối với nhau theo một thứ tự. Mỗi một block sẽ tham chiếu đến một block trước nó ngoại trừ genesis block. Genesis block là block đầu tiền trong blockchain cái mà được tạo ra bằng tay tại thời điểm blockchain bắt đầu.

[image]

#### 1.1.2. Features of blockchain

#### 1.1.3.Tiers of blockchain

## 2. Type of blockchain

### 2.1. Public blockchain

### 2.2. Private blockchain

### 2.3. Permissioned ledger

### 2.4. Shared ledger

### 2.5. Consensus in blockchain

Consensus là quá trình đồng thuận giữa các node không tin cậy về trạng thái cuối cùng của dữ liệu. Có nhiều giải thuật được sử dụng để đạt được sự động thuận này. Quá trình đồng thuận có thể dễ dàng thực hiện trên 2 node nhưng khi có nhiều node trong hệ thống phân tán việc đồng thuận trở nên khó khăn. Giải thuật giúp đạt được sự đồng thuận giữa các node trong môi trường phân tán gọi là distributed consensus.

Cơ chế đồng thuận: Một cơ chế để đồng thuận là một tập các bước được thực hiện bởi tất cả hoặc hầu hết các node để đạt được sự đồng ý về một đề xuất về trạng thái hoặc giá trị. 

Các yêu cầu để đạt được kết quả mong muốn trong cơ chế đồng thuận:
- Agreegreement: Tất cả các nodes thống nhất về một giá trị
- Termination: Tất cả các nút chấm dứt quá trình đồng thuận và cuối cùng đạt được một quyết định.
- Validity: Giá trị được đồng ý bởi tất cả các nodes phải giống với giá trị được đề xuất bởi ít nhất 1 node.
- Fault tolerant: Giải thuật đồng thuận có thể thực hiện được trong trường hợp có sự xuất hiện của node bị lỗi hay node không tin cậy(byzantine nodes)
- Integrity: Không một node nào có quyển quyết định lớn hơn 1 node nào khác. Mỗi một node quyết định đúng một lần trong một chu kỳ đồng thuận.

#### 2.5.1. Proof of work

Thuật toán hash là thuần toán nhận ra trị input là một sâu có độ dài bắt kỳ và trả về output một sâu có độ dài cố định. Với cùng một sâu ịnput ta luôn xác định được duy nhất 1 sâu output và kết quả này có thể dễ dàng tính toán và verify bởi bất kỳ một người nào implement thuật toán hash. Điểm đặc biệt của thuật toán cryptographic hash là hầu như không thể tìm được 2 giá trị input cho ra cùng một giá trị output. 

Ví dụ dưới đây là một dãy các input và output:

```
I am Satoshi Nakamoto0 => a80a81401765c8eddee25df36728d732...
I am Satoshi Nakamoto1 => f7bc9a6304a4647bb41241a677b5345f...
I am Satoshi Nakamoto2 => ea758a8134b115298a1583ffb80ae629...
I am Satoshi Nakamoto3 => bfa9779618ff072c903d773de30c99bd...
I am Satoshi Nakamoto4 => bce8564de9a83c18c31944a66bde992f...
I am Satoshi Nakamoto5 => eb362c3cf3479be0a97a20163589038e...
I am Satoshi Nakamoto6 => 4a2fd48e3be420d0d28e202360cfbaba...
I am Satoshi Nakamoto7 => 790b5a1349a5f2b909bf74d0d166b17a...
I am Satoshi Nakamoto8 => 702c45e5b15aa54b625d68dd947f1597...
I am Satoshi Nakamoto9 => 7007cf7dd40f5e933cd89fff5b791ff0...
I am Satoshi Nakamoto10 => c2f38c81992f4614206a21537bd634a...
I am Satoshi Nakamoto11 => 7045da6ed8a914690f087690e1e8d66...
I am Satoshi Nakamoto12 => 60f01db30c1a0d4cbce2b4b22e88b9b...
I am Satoshi Nakamoto13 => 0ebc56d59a34f5082aaef3d66b37a66...
I am Satoshi Nakamoto14 => 27ead1ca85da66981fd9da01a8c6816...
I am Satoshi Nakamoto15 => 394809fb809c5f83ce97ab554a2812c...
I am Satoshi Nakamoto16 => 8fa4992219df33f50834465d3047429...
I am Satoshi Nakamoto17 => dca9b8b4f8d8e1521fa4eaa46f4f0cd...
I am Satoshi Nakamoto18 => 9989a401b2a3a318b01e9ca9a22b0f3...
I am Satoshi Nakamoto19 => cda56022ecb5b67b2bc93a2d764e75f...

```

Ta thấy các input chỉ giống nhau đoạn đầu và chỉ khác nhau các số cuối. Số này được gọi là số `nonce`. Thuật toán proof of work sẽ cho trước một output của thuật toán hash gọi là target ví dụ: `0x1000000000000000000000000000000000000000000000000000000000000000` và các node trong hệ thống sẽ tiến hành chạy thuật toán hash để tìm ra giá trị `nonce` sao cho giá trị hash tương ứng nhỏ hơn so với giá trị target theo hệ cơ số 16. Ta dễ dàng có thể thấy trong ví dụ trên khi nonce băng 13 ta tìm được output nhỏ hơn target. Tưởng tưởng khi giảm target đi thì việc tìm kiếm sẽ trở nên khó khăn do phải thử nhiều `nonce` hơn. Việc tìm ra được `nonce` chính là giải xong bài toán proof of work

#### 2.5.2. Proof of stake

Thuật toán proof of stake sẽ phân phối việc xác minh giao dịch cho các node dựa vào cổ phần node đó đóng ghóp vào. Ví dụ một node cổ phần vào 10% thì node đó có quyền xác minh 10% giao dịch.  

#### 2.5.3. Delegated proof of stake

Các node trong hệ thống sẽ tiến hành bỏ phiếu cho các node khác để các node đó có thể tiến hành tạo block mới. Top 100 node gọi là Witness được bỏ phiếu nhiều nhất sẽ được tiền hành tạo block mới. Quyền biểu quyết của một node được xác định dựa vào số token mà họ dữ. Có nghĩa là người dữ nhiều token sẽ có quyền biểu quyết lớn hơn người khác. Các node cũng có quyền bỏ phiếu để loại bỏ một node gọi Witness nếu node đó làm việc không hiệu quả. 

#### 2.5.4. Proof of Elapsed Time

Thuật toán dựa trên 2 chiến lược:
- Mỗi node trong mạng lưới blockchain chờ đợi một khoảng thời gian ngẫu nhiên.
- Node đầu tiên để hoàn thành thời gian chờ sẽ được tạo cho block mới

#### 2.5.4. Practical Byzentine Fault tolerance

Sử dụng giải thuật byzentine để xác định node lỗi.

## 3. Benefits and limitations of blockchain

- Decentralization: Hệ thống blockchain không cần một bên thứ ba thực sự tin tưởng hoặc một người trung gian để xác thực giao dịch. 
- Transparency and trust: Blockchain được lưu trữ một cách công khai và bất kỳ người nào cũng có thể xem được những thứ ở trong blockchain. Điều này làm cho hệ thống trở nên minh bạch và một kết quả có thể tin tưởng khi đã được lưu vào blockchain. 
- Immutability: Dữ liệu một khi được khi vào blockchain gần như không thể thay đổi. Mặc dù không phải là không thể thay đổi nhưng khẳ năng thay đổi là cực kỳ khó. 
- High avaibility: Blockchian được lưu trữ toàn bộ trên tất cả các node trong hệ thống peer to peer, dữ liệu được sao lưu và update trên một node và trên mọi node do đó hệ thống có tính khả dụng cao ngay cả khi một số nodes bị lỗi hoặc không tin cậy. 



## TLTK

https://hackernoon.com/explain-delegated-proof-of-stake-like-im-5-888b2a74897d
Mastering Blockchain
Mastering bitcoin
https://medium.com/kokster/understanding-hyperledger-sawtooth-proof-of-elapsed-time-e0c303577ec1
https://medium.com/@chrshmmmr/consensus-in-blockchain-systems-in-short-691fc7d1fefe