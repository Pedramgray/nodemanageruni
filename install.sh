#!/bin/bash

echo "[+] نصب وابستگی‌ها"
sudo apt update
sudo apt install golang -y

echo "[+] دریافت کتابخانه‌های Go"
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite

echo "[+] ساخت فایل اجرایی"
go build -o nodemanager

echo "[+] ساخت و کپی فایل سرویس"
cp nodemanager.service /etc/systemd/system/

echo "[+] راه‌اندازی سرویس"
sudo systemctl daemon-reexec
sudo systemctl enable nodemanager
sudo systemctl start nodemanager

echo "[✓] نصب کامل شد. باز کردن در مرورگر:"
hostname -I | awk '{print "http://"$1":8080/index.html"}'
