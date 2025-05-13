## Withdraw Money

### Scenario 1: User withdraw amount with sufficient wallet balance
```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/withdraw<br/>(amount)
    Backend->>PostgresDB: Begin Transaction
    Backend->>PostgresDB: Retrieve wallet belongs to user
    PostgresDB-->>Backend: Return wallet record
    Backend->>Backend: Validate withdrawal eligibility
    Backend->>PostgresDB: Create ledger record<br/>(description: "Withdrawal from wallet",<br/>type: "withdrawal")
    Backend->>PostgresDB: Deduct balance from wallet
    Backend->>PostgresDB: Commit Transaction
    Backend-->>Client: Return updated wallet balance<br/>{<br/>wallet: {<br/>amount: 12000,<br/>formattedAmount: "MYR 120"<br/>},<br/>transaction: {<br/>id: "xxx",<br/>amount: 200,<br/>formattedAmount: "MYR 2.00",<br/>description: "Balance withdrawal from account x"<br/>}<br/>}
```


### Scenario 2: User withdraw amount that is exceed the wallet balance
```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/:wallet_id/withdraw<br/>(amount)
    Backend->>PostgresDB: Begin Transaction
    Backend->>PostgresDB: Lookup wallet_id belongs to user
    PostgresDB-->>Backend: Return wallet record
    Backend->>Backend: Validate withdrawal eligibility
    Note over Backend: Insufficient Balance Check
    Backend->>PostgresDB: Rollback Transaction
    Backend-->>Client: Return Bad Request<br/>Status: 400<br/>{<br/>error_code: "ERR_WITHDRAW_10001",<br/>message: "Insufficient balance"<br/>}
```

### Scenario 3: User withdraw amount from an invalid wallet account id

```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/withdraw<br/>(amount)
    Backend->>PostgresDB: Begin Transaction
    Backend->>PostgresDB: Return wallet record belongs to user
    PostgresDB-->>Backend: Return no wallet record
    Backend->>PostgresDB: Rollback Transaction
    Backend-->>Client: Return Unauthorized<br/>Status: 401<br/>{<br/>error_code: "ERR_WITHDRAW_10000",<br/>message: "Unauthorized access to wallet"<br/>}
```

### Scenario 4: User withdraw with incorrect amount
```mermaid
sequenceDiagram
    actor Client
    participant Backend
    participant PostgresDB

    Client->>Backend: POST /api/v1/wallet/withdraw<br/>(amount)
    Backend->>Backend: Validate withdraw amount
    Backend-->>Client: Return Bad Request<br/>Status: 400<br/>{<br/>error_code: "ERR_WITHDRAW_10003",<br/>message: "Invalid withdrawal amount"<br/>}
```
