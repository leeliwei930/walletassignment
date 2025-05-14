
### Scenario 1: Transfer to User B wallet account
```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/transfer<br/>(amount, recipient_phone_number, description)
    Backend->>Backend: Validate transfer amount
    Backend->>Backend: Validate sender and receiver phone number
    Backend->>PostgresDB: Begin Transaction
    Backend->>PostgresDB: Retrieve sender wallet account balance
    PostgresDB-->>Backend: Return sender wallet account record
    Backend->>PostgresDB: Retrieve receiver wallet account by recipient_phone_number
    PostgresDB-->>Backend: Destination wallet account balance
    Note over Backend: Check sufficient balance for transfer
    Backend->>Backend: Validate transfer amount is sufficient
    Backend->>Backend: Compute updated balance and balance to debit on receiver's account wallet
    Backend->>Backend: Compute updated balance and balance to credit from sender's account wallet
    Backend->>PostgresDB: Add new debit transaction to ledger<br/>(amount: 120, wallet_id: xxx,<br/>description: "Balance transfer from user A",<br/>type: "transfer_in")
    PostgresDB-->>Backend: Return created debit transaction record
    Backend->>PostgresDB: Update receiver's wallet account balance
    PostgresDB-->>Backend: Return updated receiver wallet account balance
    Backend->>PostgresDB: Add new credit transaction to ledger<br/>(amount: 120, wallet_id: xxx,<br/>description: "Balance transfer to user B",<br/>type: "transfer_out")
    Backend->>PostgresDB: Update sender's wallet account balance
    Backend-->>Client: Return updated source wallet account balance and transaction<br/>{<br/>wallet: {<br/>balance: 190000,<br/>formattedBalance: "MYR 1900"<br/>},<br/>transaction: {<br/>amount: 190,<br/>description: "Balance transfer to User B"<br/>}<br/>}
```


### Scenario 2: Transfer balance to recipient that is identical with the source user id

```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/transfer<br/>(amount, recipient_phone_number, description)
    Backend->>Backend: Validate transfer amount
    Backend->>Backend: Validate sender and receiver phone number
    Note over Backend: Self-transfer Validation Check
    Backend-->>Client: Return Bad Request<br/>Status: 400<br/>{<br/>error_code: "ERR_TRANSFER_10001",<br/>message: "Recipient phone number must not be identical as sender"<br/>}
```


### Scenario 3: Transfer to unknown destination account
```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/:wallet_id/transfer<br/>(amount, recipient_phone_number, description)
    Backend->>Backend: Validate transfer amount
    Backend->>Backend: Validate sender and receiver phone number
    Backend->>PostgresDB: Retrieve destination wallet account by user id
    PostgresDB-->>Backend: Return no wallet record found
    Backend-->>Client: Return Not Found<br/>Status: 404<br/>{<br/>error_code: "ERR_TRANSFER_10002",<br/>message: "Invalid recipient phone number"<br/>}
```

### Scenario 4: Transfer amount exceed the existing wallet balance
```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/transfer<br/>(amount, recipient_user_id, description)
    Backend->>Backend: Validate transfer amount
    Backend->>Backend: Validate sender and receiver phone number
    Backend->>PostgresDB: Begin Transaction
    Backend->>PostgresDB: Retrieve source wallet account balance
    PostgresDB-->>Backend: Return source wallet account record
    Note over Backend: Check sufficient balance for transfer
    Backend->>Backend: Validate transfer amount is sufficient
    Note over Backend: Insufficient Balance Check
    Backend->>PostgresDB: Rollback Transaction
    Backend-->>Client: Return Bad Request<br/>Status: 400<br/>{<br/>error_code: "ERR_TRANSFER_10003",<br/>message: "Insufficient balance for transfer"<br/>}
```
