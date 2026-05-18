# Zadanie 6

Sklep demo z backendem Go, frontendem React oraz modułem Kotlin (Spring Boot).

## Struktura

- `backend/` — serwer HTTP w Go
- `frontend/` — aplikacja React + Vite
- `kotlin/` — serwis autoryzacji w Kotlin
- `oss/gocat/` — mini projekt open source z poprawka CodeQL
- `.husky/` + `lint-staged` — lintowanie przed commitem

## Uruchomienie

Backend:

```
cd 6/backend
go run .
```

Frontend:

```
cd 6/frontend
npm install
npm run dev
```

Kotlin:

```
cd 6/kotlin
./gradlew bootRun
```

## Husky + lint-staged

```
cd 6
npm install
```

Hook `pre-commit` uruchamia `lint-staged`, które lintuje zmienione pliki JS/JSX (ESLint) oraz Go (`gofmt`, `go vet`).

## SonarCloud

Po utworzeniu projektów w SonarCloud podmień `<project-key>` na własne klucze.

### Backend (Go)

[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-backend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=zadanie6-backend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-backend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=zadanie6-backend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-backend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=zadanie6-backend)

### Frontend (React)

[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-frontend&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=zadanie6-frontend)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-frontend&metric=bugs)](https://sonarcloud.io/summary/new_code?id=zadanie6-frontend)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-frontend&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=zadanie6-frontend)

### Kotlin

[![Quality Gate](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-kotlin&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=zadanie6-kotlin)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-kotlin&metric=bugs)](https://sonarcloud.io/summary/new_code?id=zadanie6-kotlin)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=zadanie6-kotlin&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=zadanie6-kotlin)
