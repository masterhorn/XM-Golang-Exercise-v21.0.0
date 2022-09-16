// For test exercise all data was hardcoded. All this data might be putting
// in some of secret registry Github Secrest for ex.
package cfg

import (
	"github.com/unrolled/render"
)

var Ren = render.New()

type Cfg struct {
	DbConnectionString string
	DbUser             string
	DbPassword         string
	DbName             string
}

var Config = Cfg{
	DbConnectionString: "localhost:5432",
	DbUser:             "postgres",
	DbPassword:         "CorranHorn2689",
	DbName:             "postgres",
}

// Use this token as value with key AccessToken in request Header
const JWT_Token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NjMzMTY2MzYsImV4cCI6MTY5NDg1MjYzNiwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkFkbWluIiwiU3VybmFtZSI6IkFkbWluIiwiRW1haWwiOiJ4bUBleGFtcGxlLmNvbSIsIlJvbGUiOlsiTWFuYWdlciIsIlByb2plY3QgQWRtaW5pc3RyYXRvciJdfQ.Xi46Y3ENr44300tljGRwP47cNDdS3ByjL45S61g6anA"
