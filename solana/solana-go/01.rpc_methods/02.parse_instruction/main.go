package main

import (
	"context"
	"encoding/base64"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/text"
)

func main() {
	exampleFromGetTransaction()
}

func exampleFromBase64() {
	encoded := "AfjEs3XhTc3hrxEvlnMPkm/cocvAUbFNbCl00qKnrFue6J53AhEqIFmcJJlJW3EDP5RmcMz+cNTTcZHW/WJYwAcBAAEDO8hh4VddzfcO5jbCt95jryl6y8ff65UcgukHNLWH+UQGgxCGGpgyfQVQV02EQYqm4QwzUt2qf9f1gVLM7rI4hwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA6ANIF55zOZWROWRkeh+lExxZBnKFqbvIxZDLE7EijjoBAgIAAQwCAAAAOTAAAAAAAAA="

	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}

	// parse transaction:
	tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(data))
	if err != nil {
		panic(err)
	}

	decodeSystemTransfer(tx)
}

func exampleFromGetTransaction() {
	endpoint := rpc.TestNet_RPC
	client := rpc.New(endpoint)

	txSig := solana.MustSignatureFromBase58("3hZorctJtD3QLCRV3zF6JM6FDbFR5kAvsuKEG1RH9rWdz8YgnDzAvMWZFjdJgoL8KSNzZnx7aiExm1JEMC8KHfyy")
	{
		out, err := client.GetTransaction(
			context.TODO(),
			txSig,
			&rpc.GetTransactionOpts{
				Encoding: solana.EncodingBase64,
			},
		)
		if err != nil {
			panic(err)
		}

		tx, err := solana.TransactionFromDecoder(bin.NewBinDecoder(out.Transaction.GetBinary()))
		if err != nil {
			panic(err)
		}

		decodeSystemTransfer(tx)
	}
}

func decodeSystemTransfer(tx *solana.Transaction) {
	spew.Dump(tx)

	// Get (for example) the first instruction of this transaction
	// which we know is a `system` program instruction:
	i0 := tx.Message.Instructions[0]

	// Find the program address of this instruction:
	progKey, err := tx.ResolveProgramIDIndex(i0.ProgramIDIndex)
	if err != nil {
		panic(err)
	}

	// Find the accounts of this instruction:
	accounts, err := i0.ResolveInstructionAccounts(&tx.Message)
	if err != nil {
		panic(err)
	}

	// Feed the accounts and data to the system program parser
	// OR see below for alternative parsing when you DON'T know
	// what program the instruction is for / you don't have a parser.
	inst, err := system.DecodeInstruction(accounts, i0.Data)
	if err != nil {
		panic(err)
	}

	// inst.Impl contains the specific instruction type (in this case, `inst.Impl` is a `*system.Transfer`)
	spew.Dump(inst)
	if _, ok := inst.Impl.(*system.Transfer); !ok {
		panic("the instruction is not a *system.Transfer")
	}

	// OR
	{
		// There is a more general instruction decoder: `solana.DecodeInstruction`.
		// But before you can use `solana.DecodeInstruction`,
		// you must register a decoder for each program ID beforehand
		// by using `solana.RegisterInstructionDecoder` (all solana-go program clients do it automatically with the default program IDs).
		decodedInstruction, err := solana.DecodeInstruction(
			progKey,
			accounts,
			i0.Data,
		)
		if err != nil {
			panic(err)
		}
		// The returned `decodedInstruction` is the decoded instruction.
		spew.Dump(decodedInstruction)

		// decodedInstruction == inst
		if !reflect.DeepEqual(inst, decodedInstruction) {
			panic("they are NOT equal (this would never happen)")
		}

		// To register other (not yet registered decoders), you can add them with
		// `solana.RegisterInstructionDecoder` function.
	}

	{
		// pretty-print whole transaction:
		_, err := tx.EncodeTree(text.NewTreeEncoder(os.Stdout, text.Bold("TEST TRANSACTION")))
		if err != nil {
			panic(err)
		}
	}
}
