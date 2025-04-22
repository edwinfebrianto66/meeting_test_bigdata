# Big Data Benchmark Project

Benchmark performa pengolahan dataset besar (1.000.000 baris CSV) menggunakan 4 bahasa backend:
- Node.js (Express)
- Golang (net/http)
- Python (Flask + Pandas + Psutil)
- PHP (Native CLI Server)

## 📁 Struktur Folder
```
bigdata-benchmark/
├── node-service/
├── go-service/
├── python-service/
├── php-service/
├── docker-compose.yml
└── .gitignore
```

---

## 🚀 Menjalankan dengan Docker
1. Masuk ke direktori:
   ```bash
   cd bigdata-benchmark
   ```

2. Jalankan semua service sekaligus:
   ```bash
   docker-compose up --build
   ```

3. Akses endpoint masing-masing:

| Backend  | Port | URL                                     |
|----------|------|-----------------------------------------|
| Node.js  | 4000 | [http://localhost:4000/test-bigdata](http://localhost:4000/test-bigdata) |
| Golang   | 4001 | [http://localhost:4001/test-bigdata](http://localhost:4001/test-bigdata) |
| Python   | 4002 | [http://localhost:4002/test-bigdata](http://localhost:4002/test-bigdata) |
| PHP      | 4003 | [http://localhost:4003/index.php](http://localhost:4003/index.php)       |


---

## ⚙️ Menjalankan Manual (Tanpa Docker)

### Node.js (port 4000)
```bash
cd node-service
npm install express csv-parser
node index.js
```

### Python (port 4002)
```bash
cd python-service
pip install -r requirements.txt
python app.py
```

### Golang (port 4001)
```bash
cd go-service
go run main.go
```

### PHP (port 4003)
```bash
cd php-service
php -S localhost:4003
```

---

## 📊 Output Endpoint
Contoh JSON hasil:
```json
{
  "total_rows": 1000000,
  "avg_age": 39.12,
  "total_salary": 5347123456,
  "execution_time_ms": 812,
  "memory_used_mb": 34.21,
  "cpu_percent": 18.5
}
```

---

## 🎯 Tujuan
- Mengukur performa pengolahan file besar
- Menentukan backend paling efisien dari sisi waktu dan resource (RAM, CPU)
