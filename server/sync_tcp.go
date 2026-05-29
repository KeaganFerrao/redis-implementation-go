package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/KeaganFerrao/redis-implementation-go/config"
	"github.com/KeaganFerrao/redis-implementation-go/core"
)

// TODO: Max read in one shot is 512 bytes
// To allow input > 512 bytes, then repeated read until
// we get EOF or designated delimiter
func readCommand(c io.ReadWriter) (*core.RedisCmd, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return nil, err
	}

	tokens, err := core.DecodeArrayString(buf[:n])
	if err != nil {
		return nil, err
	}

	return &core.RedisCmd{
		Cmd:  strings.ToUpper(tokens[0]),
		Args: tokens[1:],
	}, nil
}

func respondError(err error, c io.ReadWriter) {
	fmt.Fprintf(c, "-%s\r\n", err)
}

func respond(cmd *core.RedisCmd, c io.ReadWriter) {
	err := core.EvalAndRespond(cmd, c)
	if err != nil {
		respondError(err, c)
	}
}

func RunSyncTCPServer() {
	log.Println("Starting a synchronous TCP server on", config.Host, config.Port)

	var con_clients int = 0

	// listening to the configured host:port
	lsnr, err := net.Listen("tcp", config.Host+":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Println("err", err)
		return
	}

	for {
		c, err := lsnr.Accept()
		if err != nil {
			log.Println("err", err)
		}

		con_clients += 1

		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1

				if err == io.EOF {
					break
				}
			}

			respond(cmd, c)
		}
	}
}
