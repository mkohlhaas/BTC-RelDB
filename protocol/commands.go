package protocol

type cmdString struct {
	cmd string
}

var (
	cmdVersion     = newCommand("version")
	cmdVersionAck  = newCommand("verack")
	cmdTransaction = newCommand("tx")
)

func newCommand(c string) [12]byte {
	res := append([]byte(c), 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
	return [12]byte(res)
}
