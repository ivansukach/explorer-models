package entities

import "time"

type Validator struct {
	ID               int64  `json:"id" sql:",pk"`
	AccountAddress   string `json:"account_address" sql:",notnull, unique"`
	OperatorAddress  string `json:"operator_address" sql:",notnull, unique"`
	ConsensusAddress string `json:"consensus_address" sql:",notnull"`
	ConsensusPubKey  string `json:"consensus_pubkey" sql:",notnull, unique"`

	Moniker         string `json:"moniker" `
	Identity        string `json:"identity"`         // optional identity signature (ex. UPort or Keybase)
	Website         string `json:"website"`          // optional website link
	SecurityContact string `json:"security_contact"` // optional security contact info
	Details         string `json:"details"`          // optional details

	Jailed          bool    `json:"jailed"`
	Status          string  `json:"status"`
	Tokens          float64 `json:"tokens"`
	DelegatorShares string  `json:"delegator_shares"`

	UnbondingHeight int64 `json:"unbonding_height" sql:"default:0"`
	UnbondingTime   int64 `json:"unbonding_time"`

	Rate          string `json:"rate"`
	MaxRate       string `json:"max_rate"`
	MaxChangeRate string `json:"max_change_rate"`
	UpdateTime    int64  `json:"update_time"`

	Timestamp         time.Time `json:"timestamp" sql:"default:now()"`
	MinSelfDelegation string    `json:"min_self_delegation"` // validator's self declared minimum self delegation

}

//// NewValidator returns a new Validator.
//func NewValidator(v Validator) *Validator {
//	return &Validator{
//		ID:                   v.ID,
//		Moniker:              v.Moniker,
//		OperatorAddress:      v.OperatorAddress,
//		ConsensusPubKey:      v.ConsensusPubKey,
//		Identity:             v.Identity,
//		Website:              v.Website,
//		SecurityContact:      v.SecurityContact,
//		Details:              v.Details,
//		Jailed:               v.Jailed,
//		Status:               v.Status,
//		Tokens:               v.Tokens,
//		DelegatorShares:      v.DelegatorShares,
//		UnbondingHeight:      v.UnbondingHeight,
//		UnbondingTime:        v.UnbondingTime,
//		Rate:                 v.Rate,
//		MaxRate:              v.MaxRate,
//		MaxChangeRate:        v.MaxChangeRate,
//		CommissionUpdateTime: v.CommissionUpdateTime,
//		Timestamp:            time.Now().UTC(),
//		MinSelfDelegation:    v.MinSelfDelegation,
//	}
//}
