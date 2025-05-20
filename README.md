# Node Manager (Go + SQLite)

پروژه مدیریت نود با استفاده از ساختار داده‌های درخت جستجوی دودویی (BST) و صف اولویت (MaxHeap)، پیاده‌سازی شده با زبان Go و پایگاه‌داده SQLite.

## امکانات

- افزودن نود با ID، نام و اولویت
- جستجو و حذف نود بر اساس ID
- افزایش اولویت نود
- رسیدگی به نود با بیشترین اولویت
- مشاهده نودها بر اساس ID (درخت)
- مشاهده نودها بر اساس اولویت (صف)
- نمایش گرافیکی نتایج در HTML زیبا

## راه‌اندازی سریع روی Ubuntu

### 1. نصب Go

```bash
sudo apt update
sudo apt install golang -y
```

### 2. اجرای دستی پروژه

```bash
git clone https://github.com/YOUR_USERNAME/nodemanager.git
cd nodemanager
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get gorm.io/driver/sqlite
go run main.go
```

### 3. دسترسی از مرورگر

باز کردن:

```
http://YOUR_SERVER_IP:8080/index.html
```

در گوشی یا لپ‌تاپ، به شرط باز بودن پورت 8080:

```bash
sudo ufw allow 8080
```

---

## اجرای دائمی پروژه با systemd

### 1. ساخت فایل باینری

```bash
go build -o nodemanager
```

### 2. کپی فایل سرویس

```bash
sudo cp nodemanager.service /etc/systemd/system/
```

### 3. فعال‌سازی سرویس

```bash
sudo systemctl daemon-reexec
sudo systemctl enable nodemanager
sudo systemctl start nodemanager
```

### بررسی وضعیت:

```bash
sudo systemctl status nodemanager
```

---

## ساختار پوشه‌ها

- `main.go`: کد اصلی پروژه
- `static/index.html`: رابط کاربری
- `nodes.db`: پایگاه‌داده SQLite
- `nodemanager.service`: فایل اجرای دائمی
- `install.sh`: اسکریپت نصب خودکار

# nodemanageruni
# nodemanageruni
