# PROJECT MOSES

## How to use

```bash
docker compose up -d
## open http://localhost:8080 on your browser.
```

or 

Terminal 1
```bash
cd backend
cargo run
```

Terminal2
```bash
cd html_front
go run .
```

## Architecture

### TOBE(?)
```mermaid
graph LR
    BU[Browser] <--"HTTP"--> F[HTML Frontend]
    F <--GraphQL--> BE[Backend]
    BE <--> D[Database]

    CB[Cli Browser] <--"(?)"--> CF[CLI Frontend]
    CF<--GraphQL-->BE
    CB<--"GraphQL"-->BE
```

- Brower: Chrome, Firefox, Safari, etc.
- HTML Frontend: Create HTML using moses-frontend-lib.
- moses-frontend-lib: Create a request of GraphQL and convert the response to ts object(?).
- Backend: Return blog content as JSON.
- Database: Store blog contents.
- Cli Browser: Blog viewer for CLI.
- CLI Frontend: Respond to requests from Cli Browser using the new protocol.



### Current
```mermaid
graph LR
    BU[Browser] <--"HTTP"--> F[HTML Frontend]
    F --REST API--> BE[Backend]
    BE--JSON--> F

```