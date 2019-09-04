module zzz

go 1.12

require (
	github.com/JonathanMH/goClacks v0.0.0-20170325034831-aa5286893e3c
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	golang.org/x/crypto v0.0.0-20190820162420-60c769a6c586 // indirect
	util v0.0.0
)

replace util => ./util
