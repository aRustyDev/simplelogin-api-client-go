//go:generate mockgen -source=account.go -destination=mock_account.go.
package simplelogin_api_client_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test function.
func TestAccountEndpoint_GetUserInfo_WithMock(t *testing.T) {
	t.Run("Test GetUserInfo", func(t *testing.T) {

		object := &AccountOptions{}

		// assert equality.
		assert.Equal(t, 123, 123, "they should be equal")

		// assert inequality.
		assert.NotEqual(t, 123, 456, "they should not be equal")

		// assert for nil (good for errors).
		assert.Nil(t, object)

		// assert for not nil (good when you expect something).
		if assert.NotNil(t, object) {

			// now we know that object isn't nil, we are safe to make.
			// further assertions without causing any errors.
			assert.Equal(t, "Something", object)

		}

	})
}
