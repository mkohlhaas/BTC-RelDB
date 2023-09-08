package protocol

var (
	cmdVersion     = newCommand("version")
	cmdVersionAck  = newCommand("verack")
	cmdTransaction = newCommand("tx")
)

func newCommand(cmd string) [12]byte {
	res := append([]byte(cmd), 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0)
	return *(*[12]byte)(res) // the same but gopls doesn't like it: return [12]byte(res)
}
