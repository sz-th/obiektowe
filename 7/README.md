# Zadanie 7

Sklep demo z prostym backendem w Go i frontem w React (Vite).

## Struktura

- `backend/` — serwer HTTP w Go obslugujacy `/api/products`, `/api/cart`, `/api/payments`.
- `frontend/` — aplikacja React + Vite z trzema widokami: produkty, koszyk, platnosci.
- `.githooks/` — lokalne hooki gita uruchamiajace lintery przed kazdym commitem.

## Uruchomienie lokalne

Backend:

```
cd 7/backend
go run main.go
```

Frontend:

```
cd 7/frontend
npm install
npm run dev
```

Albo razem przez Docker Compose w katalogu `7/`:

```
docker compose up --build
```

## Lintery i hooki

Aby aktywowac hooki gita lokalnie:

```
git config core.hooksPath 7/.githooks
```

Hook `pre-commit` (POSIX `sh`) uruchamia `gofmt -l backend` oraz `go vet ./...` w katalogu `7/backend/`. Konfiguracja `golangci-lint` znajduje sie w `7/backend/.golangci.yml` i obejmuje `govet`, `staticcheck`, `errcheck`, `gosimple`, `ineffassign`, `unused`, `gosec`. W razie potrzeby:

```
cd 7/backend
golangci-lint run
```

W systemach Windows hook dziala dzieki `sh` z Git for Windows. Wariant PowerShell znajduje sie w `7/.githooks/pre-commit.ps1`.

## SonarCloud

Po utworzeniu projektow w SonarCloud nalezy zaktualizowac ponizsze odznaki, podstawiajac wlasne klucze projektow w miejsce `<server-project-key>` oraz `<client-project-key>`.

### Serwer (Go)

[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=<server-project-key>&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=<server-project-key>)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=<server-project-key>&metric=bugs)](https://sonarcloud.io/summary/new_code?id=<server-project-key>)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=<server-project-key>&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=<server-project-key>)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=<server-project-key>&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=<server-project-key>)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=<server-project-key>&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=<server-project-key>)

### Klient (React)

[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=<client-project-key>&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=<client-project-key>)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=<client-project-key>&metric=bugs)](https://sonarcloud.io/summary/new_code?id=<client-project-key>)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=<client-project-key>&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=<client-project-key>)
