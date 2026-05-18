# Zadanie 6

Sklep demo z backendem Go, frontendem React oraz modułem Kotlin (Spring Boot).

## Struktura

- `backend/` — serwer HTTP w Go
- `frontend/` — aplikacja React + Vite
- `kotlin/` — serwis autoryzacji w Kotlin
- `.husky/` + `lint-staged` — lintowanie przed commitem

## Uruchomienie

Backend:

```
cd 6/backend
go run main.go
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
