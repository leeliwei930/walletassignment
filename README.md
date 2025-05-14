## Plan and Ideas Intro

### Authentication
For this assignment, I will implement a simplified authentication approach. Since the focus is on demonstrating basic wallet functionality between users rather than security hardening, I'll use a header-based authentication method.

Authentication will be handled by including an `X-PHONE-NUMBER` header in API requests. This header will identify which user is making the request, allowing the system to determine appropriate access permissions for wallet operations. This approach simulates user authentication without implementing complex security measures that would be required in a production environment.


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

## Project Structure
T.B.A

## Docs Structure
`docs/scenarios`
Including every stories sequence diagram written using Mermaid DSL, you can use Github to preview the diagram or a online mermaid live editor https://mermaid.live/

`docs/openapi.yaml`
Based on OAS standard v3.1, include the document of each API endpoint that is available from the server and also the response examples

## Dependencies use
T.B.A
