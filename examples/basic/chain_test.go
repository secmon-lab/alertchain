package main_test

import (
	"testing"

	"github.com/m-mizutani/alertchain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/m-mizutani/alertchain/examples/basic"
)

func TestChain(t *testing.T) {
	chain, err := Chain()
	require.NoError(t, err)
	assert.NotNil(t, chain)

	// taskEx := chain.LookupTask(&TaskExample{}).(*TaskExample)
	alert, err := chain.TestInvokeTasks(t, &alertchain.Alert{})
	assert.Nil(t, alert)
	require.Error(t, err)
}
