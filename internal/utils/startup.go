// internal/startup/wait.go
package startup

import (
	"fmt"
	"log"
	"net"
	"time"
)

func WaitForService(host string, port string, timeout time.Duration) error {
	address := net.JoinHostPort(host, port)
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", address, 2*time.Second)
		if err == nil {
			_ = conn.Close()
			log.Printf("Service %s is available.\n", address)
			return nil
		}
		log.Printf("⏳ Waiting for %s ...\n", address)
		time.Sleep(2 * time.Second)
	}
	return fmt.Errorf("Timeout while waiting for %s", address)
}
