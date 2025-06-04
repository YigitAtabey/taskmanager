# ✅ Task Manager API (Go + MySQL + JWT)

Merhaba! Bu proje, kullanıcıların görevlerini oluşturabildiği, listeleyebildiği ve silebildiği basit bir **Görev Yönetim API'sidir**.  
Backend tarafı tamamen Go (Golang) ile yazılmıştır ve MySQL veritabanı kullanmaktadır.  
JWT (JSON Web Token) ile kullanıcı doğrulama yapılmaktadır.

---

## 🚀 Kullanılan Teknolojiler

- 🧠 **Go (Golang)** – Backend geliştirme dili  
- 🔐 **JWT (JSON Web Token)** – Kimlik doğrulama  
- 🔒 **bcrypt** – Şifreleri hash'lemek için  
- 🛠️ **Gin** – Go için hızlı ve hafif web framework  
- 🐘 **MySQL** – Veritabanı  
- ⚙️ **GORM** – Go için ORM kütüphanesi  

---

## 📦 Kurulum ve Çalıştırma

> Aşağıdaki adımlar, projeyi lokal olarak çalıştırmak isteyenler içindir.

### 1. Bağımlılıkları yükle
```bash
go mod tidy
```

### 2. `.env` dosyası oluştur
Proje dizinine bir `.env` dosyası ekle:

```
DB_USER=root
DB_PASSWORD=seninsifren
DB_NAME=taskmanager
DB_PORT=5151
JWT_SECRET=senintokenkeyin
```

### 3. MySQL veritabanı oluştur
```sql
CREATE DATABASE taskmanager;
```

### 4. Projeyi çalıştır
```bash
go run main.go
```

---

## 🔑 API Endpointleri

### 👤 Kullanıcı Kayıt
```http
POST /register
{
  "name": "Yiğit",
  "email": "yigit@example.com",
  "password": "123456"
}
```

### 🔐 Giriş (Login)
```http
POST /login
{
  "email": "yigit@example.com",
  "password": "123456"
}
```
> Başarılı giriş sonrası JWT token döner.

### 👤 Profil Bilgisi
```http
GET /profile
Authorization: Bearer <JWT_TOKEN>
```

---

## 📝 Görevler

### 📌 Yeni Görev Oluştur
```http
POST /tasks
Authorization: Bearer <JWT_TOKEN>
{
  "title": "Vue.js öğren",
  "description": "Frontend tarafını Vue.js ile yap"
}
```

### 📋 Tüm Görevleri Listele
```http
GET /tasks
Authorization: Bearer <JWT_TOKEN>
```

### ❌ Görev Sil
```http
DELETE /tasks/:id
Authorization: Bearer <JWT_TOKEN>
```

---

## 📁 Proje Yapısı

```
taskmanager/
├── main.go
├── config/
│   └── db.go
├── controllers/
│   ├── auth_controller.go
│   └── task_controller.go
├── models/
│   ├── user.go
│   └── task.go
├── routes/
│   ├── user_routes.go
│   └── task_routes.go
├── utils/
│   └── jwt.go
```

---

## ✍️ Geliştirici

👤 **Yiğit Atabey**  
📧 [github.com/YigitAtabey](https://github.com/YigitAtabey)

---

## 📌 Notlar

- Token süresi, şifre güvenliği, veri doğrulama gibi konular öğrenme amaçlı uygulanmıştır.  
- Proje ilerleyen zamanda Vue.js frontend ile genişletilecektir.
