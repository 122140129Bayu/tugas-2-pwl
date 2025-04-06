# Sistem Create Order Saga

Repositori ini berisi implementasi pola **Saga** untuk mengelola transaksi terdistribusi di beberapa microservices. Sistem ini terdiri dari tiga microservices (**Order**, **Payment**, dan **Shipping**) serta **Saga Orchestrator** yang mengoordinasikan alur transaksi dan menangani tindakan kompensasi jika terjadi kegagalan.

---

## Struktur Proyek

```
.
── order-service/         # Layanan Order
── payment-service/       # Layanan Pembayaran
── shipping-service/      # Layanan Pengiriman
── orchestrator/          # Saga Orchestrator
── test-scenarios.go      # Skenario pengujian
── documentation.md       # Dokumentasi teknis
── Makefile               # Otomatisasi build & run
```

---

## Layanan

### Order Service (Port 8081)

- `POST /create-order`: Membuat pesanan baru dengan status PENDING
- `POST /cancel-order`: Membatalkan pesanan (tindakan kompensasi)
- `GET /order-status`: Mengembalikan status pesanan

### Payment Service (Port 8082)

- `POST /process-payment`: Memproses pembayaran
- `POST /refund-payment`: Mengembalikan pembayaran (tindakan kompensasi)
- `GET /payment-status`: Mengembalikan status pembayaran

### Shipping Service (Port 8083)

- `POST /start-shipping`: Memulai pengiriman
- `POST /cancel-shipping`: Membatalkan pengiriman (tindakan kompensasi)
- `GET /shipping-status`: Mengembalikan status pengiriman

### Saga Orchestrator (Port 8080)

- `POST /create-order-saga`: Memulai Saga Pembuatan Pesanan
- `GET /transaction-status`: Mengembalikan status transaksi saga

---

## Menjalankan Sistem

### Opsi Manual

1. Jalankan Order Service:

   ```bash
   cd order-service
   go run main.go
   ```

2. Jalankan Payment Service:

   ```bash
   cd payment-service
   go run main.go
   ```

3. Jalankan Shipping Service:

   ```bash
   cd shipping-service
   go run main.go
   ```

4. Jalankan Saga Orchestrator:

   ```bash
   cd orchestrator
   go run main.go
   ```

5. Jalankan test scenario:
   ```bash
   go run test-scenarios.go
   ```

## Implementasi Pola Saga

Sistem ini menggunakan pendekatan **Orchestration**, di mana seorang **orchestrator** mengarahkan urutan dan pengambilan keputusan terhadap layanan peserta.

### Alur Transaksi

1. **Create Order**: Orchestrator memanggil Order Service untuk membuat pesanan (status PENDING).
2. **Process Payment**: Jika sukses, Orchestrator memanggil Payment Service.
3. **Start Shipping**: Jika pembayaran sukses, Shipping Service akan dijalankan.
4. **Complete Transaction**: Semua langkah sukses → transaksi berstatus COMPLETED.

### Kompensasi Jika Gagal

Jika salah satu langkah gagal, orchestrator akan mengembalikan sistem ke kondisi konsisten:

- Gagal di Shipping:

  - Cancel Shipping
  - Refund Payment
  - Cancel Order

- Gagal di Payment:
  - Cancel Order

---
