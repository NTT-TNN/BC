## Hướng dẫn sử dụng thingsboard theo kịch bản tại HPC ZONE 2

Hệ thống quản lý các devices: Chuyển động, ánh sáng
Cả 2 devices được gán vào customer(HPCC_ZONE2)


- Ở mỗi device 
    - phần attribute -> client attribute sử dụng để lưu trạng thái của device. Trạng thái được cập nhật thông qua MQTT theo topic v1/devices/me/attributes dưới dạng json có dạng 
{"state":"ON"} or {"state":"OFF"} và sau sẽ sử dụng để quản lý trạng thái của các device qua API
    - Phần sẻver attribute sử dụng để lưu các thông tin không thay đổi của device như tên,... Sẽ sử dụng ở phần rule -> filter sẽ trình bày ở dưới phần Rules
    - Phần telemetry dùng để nhận dữ liệu và lưu dữ liệu dưới dạng json thông qua giao thức MQTT tại topic v1/devices/me/telemetry định dạng message: {"key1":"value1", "key2":"value2"}


### Note: Tất cả các device trên thingsBoard đều sử dụng cùng topic như đã nói ở trên. Để phần biệt các device với nhau khi publist lên MQTT ta thêm username là devicedevice acces token. Cụ thể trong kịch bản này để dễ quản lý ta sẻ tạo access token theo kịch bản HPCC_xxxxx với xxxxx là 5 số đầu trong deviceID. Việc này sẽ làm cho việc viết API đơn giản và có thể tự động.

- Phần customer dùng để nhóm các device thành một tập có thể dụng API để quản lý các device theo customers.


- Phần Rules: Dùng để điều khiển action lại thông qua MQTT. Phần rules có các phần sau
    - Filter: Dùng để lọc các điều kiện thoả mãn
        - Ta có thể thêm các filter như Message Type filter(lọc các message theo loạiloại: get hoặc post telemetry,...), Device attribute filter ( để lọc các thuộc tính server attribute như đã nói trên), Device telemetry filter( dùng để lọc dữ liệu theo điều kiênj nào đó ví dụ humi >10),...
    - Phần Plugin dùng để đưa ra action. Cụ thể trong này ta sẽ sử dụng MQTT plugin. Cách tạo sẽ trình bày trong phần plugin. Tại phần này ta có thể định nghĩa topic và message  mà message sẽ gửi khi thoả mãn các điều kiện filter.

- Phần plugin: Phần này cho phép tạo các plugin sủ dụng trong rules. Cụ thể ta sẽ tạo plugin MQTT. Ta có thể định nghĩa host và port cho MQTT tại phần này.



## Notes: Khi sử dụng Arduino để pushlish message lên thingsboard cần thêm username vàovào.

- Cần viết service để forward messge cho việc quản lý devices
