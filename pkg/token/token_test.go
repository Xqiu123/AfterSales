package token

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestToken(t *testing.T) {
	var (
		token string
		id    int = 2020
		role  int = 3
	)

	Convey("Test token", t, func() {
		Convey("Test token generation", func() {
			var err error
			token, err = GenerateToken(&TokenPayload{
				Id:      id,
				Role:    role,
				Expired: time.Hour * 2,
			})
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
