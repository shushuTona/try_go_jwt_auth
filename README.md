# try_go_jwt_auth


- [https://jwt.io/](https://jwt.io/)

## `github.com/golang-jwt/jwt` でのトークン生成の流れ

### 1. 下記のどちらかでトークン（ `*jwt.Token` ）を生成する。

- `jwt.New` : 署名で使用するアルゴリズムだけを指定してトークン生成
- `jwt.NewWithClaims` : `claims` とアルゴリズムを指定してトークン生成
    - `jwt.StandardClaims` : `sub` や `exp` など [JWTの仕様](https://datatracker.ietf.org/doc/html/rfc7519#section-4.1) に定義されているclaimが構造体のフィールドとして定義されている。
    - `jwt.MapClaims` : `map` のキーとしてclaim名を定義する。

### 2. SignedStringで署名する

`*jwt.Token.SignedString` を実行することでJWTトークン文字列を生成することができる。
