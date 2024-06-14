package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	val "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/resources"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

var hexRegExp = regexp.MustCompile("0[xX][0-9a-fA-F]+")

func NewTransferERC20Token(r *http.Request) (req resources.TransferErc20TokenRequest, err error) {
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		return req, newDecodeError("body", err)
	}

	attr := req.Data.Attributes

	err = val.Errors{
		"data/type":                val.Validate(req.Data.Type, val.Required, val.In(resources.TRANSFER_ERC20)),
		"data/attributes/sender":   val.Validate(attr.Sender.String(), val.Required, val.Match(ethAddrRegExp)),
		"data/attributes/receiver": val.Validate(attr.Receiver.String(), val.Required, val.Match(ethAddrRegExp)),
		"data/attributes/amount":   val.Validate(attr.Amount.Int64(), val.Required, val.Min(0)),
		"data/attributes/deadline": val.Validate(attr.Deadline.Int64(), val.Required, val.By(UnixTimestampRule)),
		"data/attributes/R":        val.Validate(attr.R, val.Required, val.Match(hexRegExp)),
		"data/attributes/S":        val.Validate(attr.S, val.Required, val.Match(hexRegExp)),
		"data/attributes/V":        val.Validate(attr.V, val.Required),
	}.Filter()
	if err != nil {
		return req, err
	}

	if err = VerifyPermitSignature(r, attr); err != nil {
		return req, val.Errors{
			"signature": errors.Wrap(err, "invalid permit signature"),
		}
	}

	return req, nil
}

func UnixTimestampRule(value interface{}) error {
	parsedTimestamp, ok := value.(int64)
	if !ok {
		return errors.From(errors.New("must be a valid integer"), logan.F{
			"value": value,
		})
	}

	timestamp := time.Unix(parsedTimestamp, 0)
	if timestamp.IsZero() {
		return errors.From(errors.New("timestamp is empty"), logan.F{
			"timestamp": timestamp,
		})
	}

	return nil
}

func VerifyPermitSignature(r *http.Request, attrs resources.TransferErc20TokenAttributes) error {
	sigHash, err := buildMessage(r, attrs)
	if err != nil {
		return errors.Wrap(err, "failed to build hash message")
	}

	rawSignature := make([]byte, 65)
	copy(rawSignature[:32], hexutil.MustDecode(attrs.R)[:])
	copy(rawSignature[32:64], hexutil.MustDecode(attrs.S)[:])
	rawSignature[64] = attrs.V - 27

	pubKey, err := crypto.SigToPub(sigHash, rawSignature)
	if err != nil {
		return errors.Wrap(err, "failed to recover public key from signature")
	}
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)

	if !bytes.Equal(recoveredAddr.Bytes(), attrs.Sender.Bytes()) {
		return errors.New("recovered pubkey is invalid")
	}

	return nil
}

func buildMessage(r *http.Request, attrs resources.TransferErc20TokenAttributes) ([]byte, error) {
	nonce, err := api.ERC20Permit(r).Nonces(&bind.CallOpts{}, attrs.Sender)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get nonce", logan.F{"addr": attrs.Sender})
	}

	domainSeparator, err := api.ERC20Permit(r).DOMAINSEPARATOR(&bind.CallOpts{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get domain separator")
	}

	uint256Ty, _ := abi.NewType("uint256", "uint256", nil)
	bytes32Ty, _ := abi.NewType("bytes32", "bytes32", nil)
	addressTy, _ := abi.NewType("address", "address", nil)

	args := abi.Arguments{
		{Type: bytes32Ty},
		{Type: addressTy},
		{Type: addressTy},
		{Type: uint256Ty},
		{Type: uint256Ty},
		{Type: uint256Ty},
	}

	permitTypeHash := [32]byte{}
	copy(
		permitTypeHash[:],
		crypto.Keccak256([]byte("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"))[:32],
	)

	packed, err := args.Pack(
		permitTypeHash,
		attrs.Sender,
		api.Broadcaster(r).ERC20PermitTransfer,
		attrs.Amount,
		nonce,
		attrs.Deadline,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack permit args")
	}

	structHash := crypto.Keccak256(packed)

	//keccak256(abi.encodePacked('x19x01', DOMAIN_SEPARATOR, hashed_args))
	hash := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator[:],
		structHash,
	)

	return hash, nil
}
