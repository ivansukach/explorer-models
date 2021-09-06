package objects

import (
	sdk "github.com/ivansukach/modified-cosmos-sdk/types"
	"github.com/ivansukach/modified-cosmos-sdk/x/evidence/exported"
	govTypes "github.com/ivansukach/modified-cosmos-sdk/x/gov/types"
)

type CustomMsgSubmitEvidence struct {
	Submitter string            `json:"submitter"`
	Evidence  exported.Evidence `json:"evidence"`
}

type CustomMsgSubmitProposal struct {
	Content        govTypes.Content `json:"content"`
	InitialDeposit sdk.Coins        `json:"initial_deposit"`
	Proposer       string           `json:"proposer"`
}

type CustomMsgCreateValidator struct {
	Description       Description     `json:"description"`
	Commission        CommissionRates `json:"commission"`
	MinSelfDelegation sdk.Int         `json:"min_self_delegation"`
	DelegatorAddress  sdk.AccAddress  `json:"delegator_address"`
	ValidatorAddress  sdk.ValAddress  `json:"validator_address"`
	PubKey            string          `json:"pubkey"`
	Value             sdk.Coin        `json:"value"`
}
