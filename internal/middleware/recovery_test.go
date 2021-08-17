package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_panicRecoveryMiddleware_nonPanic(t *testing.T) {
	handler := PanicRecovery()
	next := &panicRecTestHandler{}
	handlerFunc := handler(next)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handlerFunc.ServeHTTP(rr, req)

	assert.True(t, next.called)
	// panic did not happen
	assert.False(t, next.panicHappened)
}

func Test_panicRecoveryMiddleware_panic(t *testing.T) {
	handler := PanicRecovery()
	next := &panicRecTestHandler{panic: true}
	handlerFunc := handler(next)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	handlerFunc.ServeHTTP(rr, req)

	assert.True(t, next.called)
	// panic DID happen
	assert.True(t, next.panicHappened)
}

type panicRecTestHandler struct {
	panic         bool
	panicHappened bool
	called        bool
}

func (p *panicRecTestHandler) ServeHTTP(http.ResponseWriter, *http.Request) {
	p.called = true
	if p.panic {
		p.panicHappened = true
		panic("YOLO")
	}
}
