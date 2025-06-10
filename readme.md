# 🧩 Dealls HR Microservices System

Sistem ini terdiri dari beberapa layanan microservice yang menangani berbagai fungsi seperti autentikasi, absensi, payroll, dan lainnya. Dibangun dengan arsitektur yang scalable, menggunakan Go, PostgreSQL, dan Kafka.

---

## 📚 Table of Contents

1. [Software Architecture](#software-architecture)
2. [Running Services](#running-services)
3. [API Usage](#api-usage)
4. [Limitation](#limitation)

---

## 🏗️ Software Architecture

### 🔧 Microservices Breakdown

- `api-gateway-service`: Gateway utama untuk mengarahkan request ke service yang sesuai.
- `authentication-service`: Otentikasi & otorisasi pengguna.
- `employee-service`: Pengelolaan data karyawan.
- `attendance-service`: Pencatatan absensi.
- `overtime-service`: Pengajuan dan persetujuan lembur.
- `reimbursement-service`: Pengajuan reimbursement.
- `payroll-service`: Perhitungan gaji.
- `audit-service`: Logging aktivitas penting untuk keperluan audit.

### 🗃️ Tech Stack

- **Language**: Go (Golang)
- **Database**: PostgreSQL
- **Broker**: Kafka (opsional untuk komunikasi async)
- **Communication**: REST API
- **Auth**: JWT

---

## 🚀 Running Services

### 2.1 🛠️ Create Database

Pastikan PostgreSQL sudah terinstall dan berjalan. 
Gunakan tools seperti psql, DBeaver, atau pgAdmin untuk mengeksekusi SQL yang tersedia 'create-database.sql'.

### 2.2 🛠️ Running Kafka (Optional)

Run kafka server dengan tutorial [di sini](https://kafka.apache.org/quickstart)

### 2.3 🛠️ Running Service 

Jalankan semua service dengan command `./run.sh` di bash terminal seperti gitbash.

---

## 📡 API Usage

Semua dokumentasi API dapat dilihat di folder `api-spec`. Tersedia postman-collection, dan api-spec dalam bentuk PDF.

### 3.1 🔐 Login

#### 👤 Sample Users
| Username | Password  | Role             |
| -------- | --------- | ---------------- |
| 3030300  | Dealls30! | Admin + Employee |
| 3030302  | Dealls30! | Employee         |
| 3030303  | Dealls30! | Employee         |
| 3030304  | Dealls30! | Employee         |


Untuk mengakses endpoint yang memerlukan otorisasi, gunakan JWT yang didapatkan dari login endpoint berikut:

#### Admin Login
**Endpoint:**
POST /api/auth/api/login

**Request Body:**
```json
{
  "username": "3030300",
  "password": "Dealls30!"
}
```

#### Employee Login
**Endpoint:**
POST /api/auth/api/login

**Request Body:**
```json
{
  "username": "3030302",
  "password": "Dealls30!"
}
```

### 3.2 Admin Add Attendance Period

#### 👤 Sample Existing Payroll
Secara default akan membuat 3 buah payroll dengan code `currentYear_currentMonth`,`currentYear_previousMonth`,`currentYear_nextMonth`. Jika hari sekarang tahun 2025 dan bulan sekarang adalah bulan Juni maka sample data payrollnya seperti ini
| Code      |
| --------- |
| 2025_JUNE |
| 2025_MAY  |
| 2025_JULY |

#### 👤 Sample Existing Attendance Period
Secara default akan membuat 3 buah payroll dengan code `currentYear_currentMonth`,`currentYear_previousMonth`,`currentYear_nextMonth`. Jika hari sekarang tahun 2025 dan bulan sekarang adalah bulan Juni maka sample data payrollnya seperti ini
| Code      | Start Date  |   End Date   |
| --------- | --------- |  --------- |
| 2025_JUNE | 2025-06-01 00:00:00.000 +0700 | 2025-06-30 23:59:59.000 +0700|
| 2025_MAY  | 2025-05-01 00:00:00.000 +0700 | 2025-05-31 23:59:59.000 +0700|
| 2025_JULY | 2025-07-01 00:00:00.000 +0700 | 2025-07-31 23:59:59.000 +0700|

Untuk menambahkan periode kehadiran (attendance period), login terlebih dahulu sebagai Admin untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian submit attendance period dengan endpoint berikut

#### 🧾 Submit Attendance Period
**Endpoint:**
POST /api/attendance/admin/attendance-period

**Request Body:**
```json
{
    "startDate": "2026-06-01",
    "endDate": "2026-06-30",
    "payrollPeriodCode": "2026_JUNE"
}
```

### 3.3 Employee Submit Attendance

Untuk submit kehadiran (attendance) sebagai employee, login terlebih dahulu sebagai Employee untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian submit attendance dengan endpoint berikut tanpa request body.

#### 🧾 Submit Attendance 
**Endpoint:**
POST /api/attendance/api/attendance

### 3.4 Employee Submit Overtime

Untuk submit lembur (overtime) sebagai employee, login terlebih dahulu sebagai Employee untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian submit overtime dengan endpoint berikut.

#### 🧾 Submit overtime 
**Endpoint:**
POST /api/overtime/api/overtime

**Request Body:**
```json
{
    "hours": 2.5
}
```

### 3.5 Employee Submit Reimbursement

Untuk submit reimbursement sebagai employee, login terlebih dahulu sebagai Employee untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian submit reimbursement dengan endpoint berikut.

#### 🧾 Submit overtime 
**Endpoint:**
POST /api/reimbursement/api/reimbursement

**Request Body:**
```json
{
    "amount": 100000,
    "description": "Parkir bulanan"
}
```

### 3.6 Admin Run Payroll

Untuk run payroll admin, login terlebih dahulu sebagai Admin untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian run payroll dengan endpoint berikut.

#### 🧾 Run Payroll 
**Endpoint:**
POST /api/payroll/admin/payroll/run

**Request Body:**
```json
{
    "payrollCode": "2025_JUNE"
}
```

### 3.7 Employee Generate Payslip

Untuk generate payslip sebagai employee, login terlebih dahulu sebagai Employee untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian generate payslip dengan endpoint berikut.

#### 🧾 Generate Payslip 
**Endpoint:**
POST /api/payroll/api/payslip

**Request Body:**
```json
{
    "payrollCode": "2025_JUNE"
}
```
Employee akan menerima payslip berupa file PDF dari response API tersebut.

### 3.8 Admin Generate Report Payslip All Employees

Untuk generate report payslip sebagai admin, login terlebih dahulu sebagai Admin untuk mendapatkan JWT Token. Simpan JWT Token dari response untuk digunakan sebagai Bearer Token.
Kemudian Generate Report Payslip All Employees dengan endpoint berikut.

#### 🧾 Generate Report Payslip All Employees
**Endpoint:**
POST /api/payroll/admin/payslip/report-excel

**Request Body:**
```json
{
    "payrollCode": "2025_JUNE"
}
```
Admin akan menerima report payslip semua employee berupa file Excel dari response API tersebut.

---

## 🏗️ Limitation

### 🔧 Audit Service Limitation

Saat ini audit service hanya menerima event dari kafka yang di hasilkan oleh `payroll-service`

---