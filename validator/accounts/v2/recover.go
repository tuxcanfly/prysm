package v2

import (
	"fmt"
	"strings"

	"github.com/brianium/mnemonic"
	"github.com/pkg/errors"
	"github.com/prysmaticlabs/prysm/validator/flags"
	"github.com/urfave/cli/v2"
)

func RecoverWithdrawalKey(cliCtx *cli.Context) error {
	if !cliCtx.IsSet(flags.WithdrawalMnemonicFlag.Name) {
		return errors.New("Expected withdrawal mnemonic flag to be set")
	}

	enteredMnemonic := cliCtx.String(flags.WithdrawalMnemonicFlag.Name)
	words := strings.Split(enteredMnemonic, " ")

	// generate a random Mnemonic in English with 256 bits of entropy
	m := mnemonic.Mnemonic{
		Words:    words,
		Language: mnemonic.English,
	}

	// print the Mnemonic as a sentence
	fmt.Println(m.Sentence())

	// inspect underlying words
	fmt.Println(m.Words)

	// generate a seed from the Mnemonic
	seed := m.GenerateSeed(enteredMnemonic)

	// print the seed as a hex encoded string
	fmt.Println(seed)
	return nil
}
