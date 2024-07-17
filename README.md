# bank
Golang 1.21

1. Ручка создания аккаунта
```bash
curl --location --request POST 'http://localhost:3000/accounts'
```
2. Ручка пополнения баланса
```bash
curl --location 'http://localhost:3000/accounts/1/deposit' \
--header 'Content-Type: application/json' \
--data '{
    "amount":10.2
}'
```
3. Ручка списание с баланса
```bash
curl --location 'http://localhost:3000/accounts/1/withdraw' \
--header 'Content-Type: application/json' \
--data '{
    "amount":10
}'
```

4. Ручка проверка  баланса
```bash
curl --location 'http://localhost:3000/accounts/1/balance'
```
