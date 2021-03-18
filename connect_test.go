package github.com/gofiliate/db

import (
	"testing"
	"errors"
	)

func TestConnectMysqlRead(t *testing.T) {

	 doNotWant := errors.New("Test Error")

	if got := ConnectMysqlRead(); got == doNotWant {
		t.Errorf("ConnectMysqlRead() = %q, doNotWant %q", got, doNotWant)
	}
}
