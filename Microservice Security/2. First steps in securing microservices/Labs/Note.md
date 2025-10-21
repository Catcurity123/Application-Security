#### 1. To send GET request to API

##### 1.1 Without Header for authorization
```
curl -v http://localhost:8080/orders \
  -H 'Content-Type: application/json' \
  --data-binary @- <<EOF | jq
{
  "items": [
    {
      "itemCode": "IT0001",
      "quantity": 3
    },
    {
      "itemCode": "IT0004",
      "quantity": 1
    }
  ],
  "shippingAddress": "No 4, Castro Street, Mountain View, CA, USA"
}
EOF
```

##### 1.2 With Header for authorization
```
curl -v http://localhost:8080/orders \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer 73b540c3-98bd-4f18-8b03-4dac6950d17f" \
  --data-binary @- <<EOF | jq
{
  "items": [
    {
      "itemCode": "IT0001",
      "quantity": 3
    },
    {
      "itemCode": "IT0004",
      "quantity": 1
    }
  ],
  "shippingAddress": "No 4, Castro Street, Mountain View, CA, USA"
}
EOF
```

##### 1.3 To get token from AS 

```
curl -u orderprocessingapp:orderprocessingappsecret \
  -H "Content-Type: application/json" \
  -d '{"grant_type": "client_credentials", "scope": "read write"}' \
  http://localhost:8085/oauth/token | jq
```

#### 2. To send POST request to API

##### 2.1 Without Header for authorization

```
curl -v http://localhost:8080/orders/8214c381-8627-469d-86e4-55191070cd10 | jq
```

##### 2.2 With Header for authorization
```
curl -v http://localhost:8080/orders/ebd5ed77-9106-48f3-9c6d-b50b44c81c63 -H "Authorization: Bearer 16bc2a38-e783-46df-986a-eab57259e003" | jq
```

#### 3. To send requests to API Gateway

```
curl -u application1:application1secret \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "grant_type=client_credentials" \
  http://localhost:9090/token/oauth/token | jq
```

```
curl -v http://localhost:9090/retail \
  -H 'Content-Type: application/json' \
  -H "Authorization: Bearer 73b540c3-98bd-4f18-8b03-4dac6950d17f" \
  --data-binary @- <<EOF | jq
{
  "items": [
    {
      "itemCode": "IT0001",
      "quantity": 3
    },
    {
      "itemCode": "IT0004",
      "quantity": 1
    }
  ],
  "shippingAddress": "No 4, Castro Street, Mountain View, CA, USA"
}
EOF
```
```
curl http://localhost:9090/retail/orders/9955ad7b-ee81-4381-b161-33f707475e1a \
  -H "Authorization: Bearer 73b540c3-98bd-4f18-8b03-4dac6950d17f"\ | jq
```