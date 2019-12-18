package keeper

import (
	"testing"

	"github.com/bandprotocol/d3n/chain/x/zoracle/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetterSetterCode(t *testing.T) {
	ctx, keeper := CreateTestInput(t, false)

	owner := sdk.AccAddress([]byte("owner"))

	code := []byte("This is code")
	codeHash := types.NewStoredCode(code, owner).GetCodeHash()

	_, err := keeper.GetCode(ctx, codeHash)
	require.NotNil(t, err)

	actualCodeHash := keeper.SetCode(ctx, code, owner)
	storedCode, err := keeper.GetCode(ctx, actualCodeHash)
	require.Nil(t, err)
	require.Equal(t, code, storedCode.Code)
	require.Equal(t, owner, storedCode.Owner)
	require.Equal(t, codeHash, actualCodeHash)
}
