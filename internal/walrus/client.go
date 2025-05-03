package walrus

import (
	"log"
	"fmt"
    "bytes"
    "bufio"
    "os/exec"
)

func Store(path string, epochs int, ctx string) (string, error) {
    cmd := exec.Command("walrus", "store", path, "--epochs", fmt.Sprint(epochs), "--context", ctx)
    stdout, _ := cmd.StdoutPipe()
    cmd.Stderr = cmd.Stdout

    var buf bytes.Buffer
    scanner := bufio.NewScanner(stdout)
    go func() {
        for scanner.Scan() {
			line := scanner.Text()
			log.Println("> "+line)
            buf.WriteString(line + "\n")
        }
    }()

    if err := cmd.Start(); err != nil {
        return buf.String(), err
    }
    err := cmd.Wait()
    return buf.String(), err
}
