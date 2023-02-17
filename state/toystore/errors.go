package toystore

import "errors"

var (
	ErrNoToys                = errors.New("no toys")
	ErrSelectToyFirst        = errors.New("select toy first")
	ErrToyDispenseInProgress = errors.New("toy dispense in progress")
	ErrToyAlreadyRequested   = errors.New("toy already requested")
	ErrReceivedMoneyIsLess   = errors.New("received money is less")
	ErrGiveMoneyFirst        = errors.New("give money first")
)
