package token

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToken(t *testing.T) {
	var (
		token string
		id    uint32 = 2020
		role  uint32 = 3
	)

	Convey("Test token", t, func() {
		Convey("Test token generation", func() {
			var err error
			var tokenPayload = &TokenPayload{
				Id:      id,
				Role:    role,
				Expired: time.Hour * 3,
			}
			token, err = tokenPayload.GenerateToken()
			So(err, ShouldBeNil)
		})

		Convey("Test token resolution", func() {
			t, err := ResolveToken(token)
			So(err, ShouldBeNil)
			So(t.Id, ShouldEqual, id)
			So(t.Role, ShouldEqual, role)
		})
	})
}
