# octiq-multistack-demo-2

A small multi-stack monorepo with one service per language Octiq supports. It is a sandbox for
running tickets: every service is a tiny, independent web app that keeps its data in memory.

| Service | Language / framework | Endpoints | Default port |
| --- | --- | --- | --- |
| `services/gateway-ts` | TypeScript / Express | `GET /health`, `GET /services` | 3000 |
| `services/catalog-go` | Go / net/http | `GET /health`, `GET /products`, `GET /products/{id}` | 3001 |
| `services/orders-java` | Java 17 / Spring Boot | `GET /health`, `GET /orders`, `POST /orders` | 3002 |
| `services/pricing-python` | Python / FastAPI | `GET /health`, `POST /quote` | 3003 |
| `services/notify-ruby` | Ruby / Sinatra | `GET /health`, `GET /notifications`, `POST /notifications` | 3004 |

Every service reads its port from the `PORT` environment variable. The services do not call each
other.

## Build, test, run

Run these inside the service directory.

| Service | Install | Test | Lint | Run |
| --- | --- | --- | --- | --- |
| gateway-ts | `npm ci` | `npm test` | `npm run lint` | `npm run build && npm start` |
| catalog-go | — | `go test ./...` | `go vet ./...` | `go run .` |
| orders-java | — | `./mvnw test` | — | `./mvnw spring-boot:run` |
| pricing-python | `python3 -m venv .venv && .venv/bin/pip install -r requirements-dev.txt` | `.venv/bin/pytest` | `.venv/bin/ruff check .` | `.venv/bin/uvicorn app.main:app --port ${PORT:-3003}` |
| notify-ruby | `bundle install` | `bundle exec rspec` | `bundle exec rubocop` | `bundle exec rackup -p ${PORT:-3004}` |
