# Build 단계: 필요한 의존성과 빌드를 수행
FROM golang:1.23-alpine AS builder

# SQLite를 위한 C 라이브러리 설치
RUN apk add --no-cache gcc libc-dev sqlite-dev

# 작업 디렉토리 설정
WORKDIR /app

# air 설치 (Hot-reload 툴)
RUN go install github.com/air-verse/air@latest

# 의존성 관리
COPY go.mod ./
RUN go mod download

# 애플리케이션 소스 복사
COPY . .

# CGO 활성화 설정 및 애플리케이션 빌드
ENV CGO_ENABLED=1
RUN go build -o app .

# 실행 단계: 가벼운 이미지를 사용하여 실행
FROM alpine:latest

# SQLite C 라이브러리 설치 (런타임 필요)
RUN apk add --no-cache sqlite-libs

# 애플리케이션 실행 디렉토리 설정
WORKDIR /app

# 빌드된 애플리케이션 복사
COPY --from=builder /app/app .
COPY .air.toml ./

# 네트워크 포트 노출
EXPOSE 3000

# Air로 애플리케이션 실행
CMD ["./app"]
