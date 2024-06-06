package handlers

import (
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/models"
	"github.com/rarimo/evm-airdrop-svc/internal/service/api/requests"
	"github.com/rarimo/evm-airdrop-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func BuildPermitHash(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewBuildPermitHashRequest(r)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to parse request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nonce, err := api.ERC20Permit(r).Nonces(&bind.CallOpts{}, req.Data.Attributes.Sender)
	if err != nil {
		api.Log(r).WithFields(logan.F{
			"sender": req.Data.Attributes.Sender,
		}).WithError(err).Error("failed to get nonce")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	domainSeparator, err := api.ERC20Permit(r).DOMAINSEPARATOR(&bind.CallOpts{})
	if err != nil {
		api.Log(r).WithError(err).Error("failed to get domain separator")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	rawMsgHash, err := buildMessageHash(
		nonce,
		api.Broadcaster(r).ERC20PermitTransfer,
		domainSeparator,
		req.Data.Attributes,
	)
	if err != nil {
		api.Log(r).WithError(err).Error("failed to get build permit message hash")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, models.NewPermitHashResponse(hexutil.Encode(rawMsgHash)))
}

func buildMessageHash(
	nonce *big.Int,
	spender common.Address,
	domainSeparator [32]byte,
	attr resources.BuildPermitHashAttributes,
) ([]byte, error) {
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

	rawPermitTypeHash := crypto.Keccak256([]byte("Permit(address owner,address spender,uint256 value,uint256 nonce,uint256 deadline)"))
	permitTypeHash := [32]byte{}
	copy(permitTypeHash[:], rawPermitTypeHash[:32])

	//abi.encode(_PERMIT_TYPEHASH, owner, spender, value, _useNonce(owner), deadline)
	packed, err := args.Pack(
		permitTypeHash,
		attr.Sender,
		spender,
		attr.Amount,
		nonce,
		attr.Deadline,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to pack permit signature")
	}

	structHash := crypto.Keccak256(packed)

	//keccak256(abi.encodePacked('x19x01', DOMAIN_SEPARATOR, hashed_args))
	msgHash := crypto.Keccak256(
		[]byte("\x19\x01"),
		domainSeparator[:],
		structHash,
	)

	return msgHash, nil
}
