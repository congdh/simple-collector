# Simple log collector
Example app  log đơn giản, nhận message từ HTTP request và ghi ra file và console.

App được cải tiến dần qua các version từ đơn giản đến phức tạp.

## Version 1
App ghi 2 string ra file, mỗi string trên 1 dòng

## Version 2
Ghi 2 dòng log append vào file log, mỗi dòng log có format

```
message := date + random string 10 kí tự
```

## Version 3
HTTP web server nhận log từ tham số message trong query string và ghi ra file

## Version 4
Web server nhận message từ request và ghi log ra file nhưng giao tiếp qua channel

## Version 5
Load cấu hình từ  file
 
## Version 6
Multi output: Ghi log đồng thời ra file và console 