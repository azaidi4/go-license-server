package main

import (
	"fmt"
	"license-server/api"
	"math/rand"
	"strconv"

	_ "github.com/lib/pq"
)

func convertInt(input string) (bool, int) {
	i, err := strconv.Atoi(input)
	if err != nil {
		return false, 0
	}
	return true, i
}

func convertBool(input string) (bool, bool) {
	i, err := strconv.ParseBool(input)
	if err != nil {
		return false, false
	}
	return true, i
}

func randomString(n int) string {
	var letterRunes = []rune("1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	fmt.Println("Go Simple Licensing System")
	api.StartServer()

	// for {
	// Menu:
	// 	fmt.Println("[Total Licenses]", count())
	// 	fmt.Println("")
	// 	fmt.Print("Console: ")
	// 	scan := bufio.NewScanner(os.Stdin)
	// 	scan.Scan()
	// 	switch scan.Text() {
	// 	case "add":
	// 		var email string
	// 		var experation string
	// 		var license string

	// 		fmt.Print("License Email: ")
	// 		scan = bufio.NewScanner(os.Stdin)
	// 		scan.Scan()
	// 		email = scan.Text()
	// 	exp:
	// 		fmt.Print("License Experation (YYYY-MM-DD): ")
	// 		scan = bufio.NewScanner(os.Stdin)
	// 		scan.Scan()
	// 		_, err = time.Parse("2006-01-02", scan.Text())
	// 		if err != nil {
	// 			fmt.Println("Experation must be in the YYYY-MM-DD Format.")
	// 			goto exp
	// 		}
	// 		experation = scan.Text()
	// 		fmt.Println("")

	// 		license = randomString(4) + "-" + randomString(4) + "-" + randomString(4)

	// 		var tmpemail string
	// 		err := db.QueryRow("SELECT email FROM licenses WHERE license='" + license + "'").Scan(&tmpemail)
	// 		if err == sql.ErrNoRows {
	// 			_, err = db.Exec("INSERT INTO licenses(email, license, experation, ip) VALUES(?, ?, ?, ?)", email, license, experation, "none")
	// 			if err != nil {
	// 				fmt.Println("[!] ERROR: UNABLE TO INSERT INTO DATABASE [!]")
	// 				fmt.Println("")
	// 				goto Menu
	// 			}
	// 		} else {
	// 			fmt.Println("License already in database?")
	// 			fmt.Println("License:", licensing.Encrypt(KEY, license))
	// 			fmt.Println("Email:", tmpemail)
	// 			fmt.Println("")
	// 			goto Menu
	// 		}

	// 		fmt.Println("License Key Generated!")
	// 		fmt.Println("")
	// 		fmt.Println("License Email:", email)
	// 		fmt.Println("License Experation:", experation)
	// 		fmt.Println("Save this as license.dat")
	// 		fmt.Println("")
	// 		fmt.Println(licensing.Encrypt(KEY, license))
	// 		fmt.Println("")

	// 	case "remove":
	// 		fmt.Print("License Email: ")
	// 		scan = bufio.NewScanner(os.Stdin)
	// 		scan.Scan()
	// 		var tmp string
	// 		err = db.QueryRow("SELECT license FROM licenses WHERE email=?", scan.Text()).Scan(&tmp)
	// 		if err == sql.ErrNoRows {
	// 			fmt.Println("[!] ERROR: COULD NOT FIND LICENSE [!]")
	// 			fmt.Println("")
	// 			goto Menu
	// 		} else {
	// 			fmt.Println("License Found:", tmp)
	// 			_ = db.QueryRow("DELETE FROM licenses WHERE email=?", scan.Text())
	// 			fmt.Println("License removed from database.")
	// 			fmt.Println("")
	// 			goto Menu
	// 		}
	// 	case "exit":
	// 		os.Exit(0)
	// 	default:
	// 		fmt.Println("Unknown Command")
	// 	}
	// }
}
