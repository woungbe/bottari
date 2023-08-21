module woungbe.bottrai

go 1.19

// replace github.com/adshao/go-binance/v2 => /Users/gyeoungryeoungpark/project/myproject/bottari/bottari/modules/v2

replace woungbe.bottrai/modules/v2 => ./modules/v2

require (
	github.com/joho/godotenv v1.5.1
	woungbe.bottrai/modules/v2 v2.0.0-00010101000000-000000000000
)

require (
	github.com/adshao/go-binance/v2 v2.4.5 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
)
