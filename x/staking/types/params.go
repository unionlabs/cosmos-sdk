package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Staking params default values
const (
	// DefaultUnbondingTime reflects three weeks in seconds as the default
	// unbonding time.
	// TODO: Justify our choice of default here.
	DefaultUnbondingTime time.Duration = time.Hour * 24 * 7 * 3

	// Default maximum number of bonded validators
	DefaultMaxValidators uint32 = 100

	// Default maximum entries in a UBD/RED pair
	DefaultMaxEntries uint32 = 7

	// Default threshold of jailed validators before triggering a forced validator set rotation. Processed as a percentage (50 => 50%)
	DefaultJailedValidorThreshold uint32 = 50

	// Default number of blocks between normal validator set rotations
	DefaultEpochLength int64 = 1
)

// Staking params bounds
const (
	// Maximum value for JailedValidorThreshold
	// JailedValidorThreshold should not be greater than 100 (100%)
	JailedValidorThresholdMax uint32 = 100

	// Minimum value for JailedValidorThreshold
	JailedValidorThresholdMin uint32 = 10

	// Maximum value for EpochLength
	EpochLengthMax int64 = 512

	// Minimum value for EpochLength
	// EpochLength should not be 0
	EpochLengthMin int64 = 1
)

var (
	// DefaultMinCommissionRate is set to 0%
	DefaultMinCommissionRate = math.LegacyZeroDec()

	// DefaultKeyRotationFee is fees used to rotate the ConsPubkey or Operator key
	DefaultKeyRotationFee = sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000000)
)

// NewParams creates a new Params instance
func NewParams(unbondingTime time.Duration,
	maxValidators, maxEntries uint32,
	bondDenom string, minCommissionRate math.LegacyDec,
	keyRotationFee sdk.Coin,
	jailedValidatorThreshold uint32, epochLength int64,
) Params {
	return Params{
		UnbondingTime:            unbondingTime,
		MaxValidators:            maxValidators,
		MaxEntries:               maxEntries,
		HistoricalEntries:        0,
		BondDenom:                bondDenom,
		MinCommissionRate:        minCommissionRate,
		KeyRotationFee:           keyRotationFee,
		JailedValidatorThreshold: jailedValidatorThreshold,
		EpochLength:              epochLength,
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		DefaultUnbondingTime,
		DefaultMaxValidators,
		DefaultMaxEntries,
		sdk.DefaultBondDenom,
		DefaultMinCommissionRate,
		DefaultKeyRotationFee,
		DefaultJailedValidorThreshold,
		DefaultEpochLength,
	)
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.LegacyAmino, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}

	return params
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.LegacyAmino, value []byte) (params Params, err error) {
	err = cdc.Unmarshal(value, &params)
	if err != nil {
		return
	}

	return
}

// validate a set of params
func (p Params) Validate() error {
	if err := validateUnbondingTime(p.UnbondingTime); err != nil {
		return err
	}

	if err := validateMaxValidators(p.MaxValidators); err != nil {
		return err
	}

	if err := validateMaxEntries(p.MaxEntries); err != nil {
		return err
	}

	if err := validateBondDenom(p.BondDenom); err != nil {
		return err
	}

	if err := validateMinCommissionRate(p.MinCommissionRate); err != nil {
		return err
	}

	if err := validateHistoricalEntries(p.HistoricalEntries); err != nil {
		return err
	}

	if err := validateKeyRotationFee(p.KeyRotationFee); err != nil {
		return err
	}

	if err := validateJailedValidorThreshold(p.JailedValidatorThreshold); err != nil {
		return err
	}

	if err := validateEpochLength(p.EpochLength); err != nil {
		return err
	}

	return nil
}

func validateUnbondingTime(i interface{}) error {
	v, ok := i.(time.Duration)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < 0 {
		return fmt.Errorf("unbonding time must not be negative: %d", v)
	}

	return nil
}

func validateMaxValidators(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateMaxEntries(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max entries must be positive: %d", v)
	}

	return nil
}

func validateHistoricalEntries(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateBondDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("bond denom cannot be blank")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func ValidatePowerReduction(i interface{}) error {
	v, ok := i.(math.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.LT(math.NewInt(1)) {
		return errors.New("power reduction cannot be lower than 1")
	}

	return nil
}

func validateMinCommissionRate(i interface{}) error {
	v, ok := i.(math.LegacyDec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("minimum commission rate cannot be nil: %s", v)
	}
	if v.IsNegative() {
		return fmt.Errorf("minimum commission rate cannot be negative: %s", v)
	}
	if v.GT(math.LegacyOneDec()) {
		return fmt.Errorf("minimum commission rate cannot be greater than 100%%: %s", v)
	}

	return nil
}

func validateKeyRotationFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("cons pubkey rotation fee cannot be nil: %s", v)
	}
	if v.IsLTE(sdk.NewInt64Coin(v.Denom, 0)) {
		return fmt.Errorf("cons pubkey rotation fee cannot be negative or zero: %s", v)
	}

	return nil
}

func validateJailedValidorThreshold(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < JailedValidorThresholdMin {
		return fmt.Errorf("jailed validor threshold cannot be less than %d: %d", JailedValidorThresholdMin, v)
	}
	if v > JailedValidorThresholdMax {
		return fmt.Errorf("jailed validor threshold cannot be greater than %d: %d", JailedValidorThresholdMax, v)
	}

	return nil
}

func validateEpochLength(i interface{}) error {
	v, ok := i.(int64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v < EpochLengthMin {
		return fmt.Errorf("epoch length cannot be less than %d: %d", EpochLengthMin, v)
	}
	if v > EpochLengthMax {
		return fmt.Errorf("epoch length cannot be greater than %d: %d", EpochLengthMax, v)
	}

	return nil
}
