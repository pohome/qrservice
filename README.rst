QR Service
=========================
QR service is a HTTP based QRCode API used to get a QRencoded image with the given text.

How to use
-------------------------
1. Get the QRcode image like this after get qrservice running:
curl -i "http://localhost:8001/qr?size=5&text=http://baidu.com"