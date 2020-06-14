// package polo

// import (
// 	"testing"
// 	"net/http/httptest"
// 	"net/http"
// 	"github.com/stretchr/testify/assert"

// 	//"github.com/nao4869/go-mvc-demo/src/api/utils/test_utils"
// )

// func TestConstants(t *testing.T) {
// 	assert.EqualValues(t, "polo", polo)
// }

// func TestPolo(t *testing.T) {
// 	response := httptest.NewRecorder()
// 	request := httptest.NewRequest(http.MethodGet, "./marco", nil)
// 	c := test_utils.GetMockedContext(request, response)

// 	Marco(c)

// 	assert.EqualValues(t, http.StatusOK, response.Code)
// 	assert.EqualValues(t, "polo", response.Body.String())
// }
