package types

import "time"

type (
	// Validator defines the structure for validator information
	Validator struct {
		ID                int64       `json:"id" sql:",pk"`
		OperatorAddress   string      `json:"operator_address" sql:",notnull"`
		ConsensusPubKey   string      `json:"consensus_pubkey" sql:",notnull"`
		Jailed            bool        `json:"jailed"`
		Status            int64       `json:"status"`
		Tokens            string      `json:"tokens"`
		DelegatorShares   string      `json:"delegator_shares"`
		Description       Description `json:"description"`
		UnbondingHeight   string      `json:"unbonding_height" sql:"default:0"`
		UnbondingTime     string      `json:"unbonding_time"`
		Commission        Commission  `json:"commission"`
		Timestamp         time.Time   `json:"timestamp" sql:"default:now()"`
		MinSelfDelegation string      `json:"min_self_delegation"`
	}

	// Description wraps description of a validator
	Description struct {
		Moniker         string `json:"moniker"`
		Identity        string `json:"identity"`
		Website         string `json:"website"`
		SecurityContact string `json:"security_contact"`
		Details         string `json:"details"`
	}

	// Commission wrpas general commission information about a validator
	Commission struct {
		CommissionRates CommissionRates `json:"commission_rates"`
		UpdateTime      string          `json:"update_time"`
	}
	CommissionRates struct {
		Rate          string `json:"rate"`
		MaxRate       string `json:"max_rate"`
		MaxChangeRate string `json:"max_change_rate"`
	}
)
