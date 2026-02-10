
# FurBook - Mạng xã hội thú cưng

Report: [pdf](https://github.com/hungq1205/furbook/blob/master/Report.pdf)

## 1. Khởi động Frontend

```bash
cd frontend
npm run dev
```

## 2. Khởi động Backend bới Docker

### Dùng Docker Compose (dùng để testing, triển khai nhanh)

```bash
cd backend
docker-compose up
```

### Dùng Kubernetes (Yêu cầu: kind - Kubernetes in Docker)

1. **Tải image Docker vào cluster `kind`**  
   Nếu chưa build, hãy chạy `docker build` trước
   
```bash
kind load docker-image backend-gateway:latest
kind load docker-image backend-user:latest
kind load docker-image backend-post:latest
kind load docker-image backend-mesage:latest
kind load docker-image backend-noti:latest
```

2. **Triển khai lên Kubernetes**

```bash
cd backend/k8s
kubectl apply -f .
```
