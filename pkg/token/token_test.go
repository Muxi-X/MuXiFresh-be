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
		email string = "123@qq.com"
	)

	Convey("Test token", t, func() {
		Convey("Test token generation", func() {
			var err error
			token, err = GenerateToken(role, id, email, time.Hour*2)
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
