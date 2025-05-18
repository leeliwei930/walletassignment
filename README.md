## Plan and Ideas Intro

### Authentication
For this assignment, I will implement a simplified authentication approach. Since the focus is on demonstrating basic wallet functionality between users rather than security hardening, I'll use a header-based authentication method.

Authentication will be handled by including an `X-USER-PHONE` header in API requests. This header will identify which user is making the request, allowing the system to determine appropriate access permissions for wallet operations. This approach simulates user authentication without implementing complex security measures that would be required in a production environment.


### Entity relation
1. **User-Wallet Relationship**: Each user has a primary wallet. The system is designed with a one-to-many relationship between users and wallets to support future extensibility (such as multi-currency wallets).
   1. Currently, each user is limited to one primary wallet.

2. **Ledger System**: The `ledgers` table serves as a comprehensive audit trail for all wallet transaction activities, maintaining a complete history of financial movements.

3. **Wallet Structure**: The wallet table stores essential information for each user-associated wallet:
   - Current balance
   - Currency type
   - Rounding precision (digits)
   
   1. **Balance Management**: A hybrid approach is implemented for balance tracking, as computing wallet balances directly from the ledgers table would be computationally expensive.
   
   2. **Data Integrity**: While the hybrid approach optimizes performance, it introduces potential data integrity challenges. In a production environment, daily reconciliation processes would be necessary to ensure ledger-computed balances match the recorded wallet balances.

### Caching Strategy and performance tuning
At the moment caching will not be implemented part of this assignment as due to tight deadline of the assignment that require me to setup the entire MVC structure from scratch take sometimes to test and validate, therefore I acknowledge that the current assignment is not perfetch and the key area to improve the performance using caching are

1. Caching on phone number map to user id, reduce the query load towards users table just to lookup specific user ID TTL suggest to be 30 mins, as data is unlikely to change
2. Caching on recipient phone number to user id, this help when making transfer to specific user phone number twice result in cache hit, reduce the time to lookup recipient user id, TTL suggest to be 30 mins
as data is unlikely to change

## Project Structure
For code reviewers the key areas to review are
| Directory/File           | Description                                                                       |
| ------------------------ | --------------------------------------------------------------------------------- |
| `ent/schema`             | Contains each entity schema rules that generate the migrations in `db/migrations` |
| `services/wallet`        | Contains the logic of the wallet transaction                                      |
| `errors`                 | Contains error code for each error return in wallet service                       |
| `internal/app/routes.go` | Contains each endpoint and it specific handler                                    |

## Docs Structure
`docs/scenarios`
Including every stories sequence diagram written using Mermaid DSL, you can use Github to preview the diagram or a online mermaid live editor https://mermaid.live/

`docs/openapi.yaml`
Based on OAS standard v3.1, include the document of each API endpoint that is available from the server and also the response examples

## Dependencies use
| Dependency    | Description                                                                          |
| ------------- | ------------------------------------------------------------------------------------ |
| labstack/echo | Express like HTTP server that serve the API endpoint                                 |
| entgo.io/ent  | ORM that manage the relationship with each entity, managing the schema of each table |
| cobra         | CLI generator to bootstrap Go CLI app, to start server                               |
| viper         | Config management and loader via environment variable                                |
| atlas         | DB migration tools that compatible with ent, provide versioning migrations to ent    |
| validate      | First layer validation                                                               |

## How to run
1. Copy the environment file from example
```bash
cp .env.example .env
```

2. Start the services in container
```bash
docker compose up -d 
```

3. Build the server
```bash
CGO_ENABLED=0 go build -o build/bin/wallet_service
```

4. Run the migrations to create table on postgres db
```bash
./build/bin/wallet_service db migrate
```

6. Run the server
The API server will be running on http://localhost:8009
```bash
./build/bin/wallet_service serve
```

7. Prepare a new test users as below, suggest to create two user to simulate the transfer
```bash
./build/bin/wallet_service user create -f "Li Wei" -l "Lee" -p "+6018129033"
2025-05-18T22:16:03.070+0800    INFO    User created successfully       {"user": "User(id=2b2e38a8-0f60-4872-99c3-95463c34d120, first_name=Li Wei, last_name=Lee, phone_number=+6018129033, created_at=Sun May 18 22:16:03 2025, updated_at=Sun May 18 22:16:03 2025)"}
```

8. To test the API please copy the API docs in `docs/openapi.yml` to 
https://editor.swagger.io

> To impersonate a specific user please click on the Authorize button in the Swagger editor, and put in the phone number of the particular user that you created earlier and start making request~
>
> The provided user phone number will be injected to headers X-USER-PHONE as user impersonation

