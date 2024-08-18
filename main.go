package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reset = "\033[0m"

func main() {
	//Verification de la longueur des arguments
	if len(os.Args) < 2 || len(os.Args) > 5 {
		fmt.Println("Nombre d'arguments non valides")
		os.Exit(0)
	}
	//colortxt = "\033[38;2;" + rgb[0] + ";" + rgb[1] + ";" + rgb[2] + "m"
	color := 0
	banner1 := 0
	condition := 0
	colortxt := ""
	if strings.HasPrefix(os.Args[1], "--color=rgb") {
		rgb := os.Args[1][12 : len(os.Args[1])-1]
		rgb2 := strings.Split(rgb, ", ")
		colortxt = "\033[38;2;" + rgb2[0] + ";" + rgb2[1] + ";" + rgb2[2] + "m"
		color++
	} else if strings.HasPrefix(os.Args[1], "--color=#") {
		colortxt = "\033[38;2;" + HexToDec(os.Args[1][9:11]) + ";" + HexToDec(os.Args[1][11:13]) + ";" + HexToDec(os.Args[1][13:15]) + "m"
		color++
	} else if strings.HasPrefix(os.Args[1], "--color=") {
		colortxt = os.Args[1][8:]
		color++
		switch colortxt {
		case "reset":
			colortxt = "\033[0m"
		case "red":
			colortxt = "\033[31m"
		case "green":
			colortxt = "\033[32m"
		case "yellow":
			colortxt = "\033[33m"
		case "blue":
			colortxt = "\033[34m"
		case "magenta":
			colortxt = "\033[35m"
		case "cyan":
			colortxt = "\033[36m"
		case "gray":
			colortxt = "\033[37m"
		case "white":
			colortxt = "\033[97m"
		case "orange":
			colortxt = "\033[38;2;" + HexToDec("FF") + ";" + HexToDec("75") + ";" + HexToDec("18") + "m"
		case "lightsalmon":
			colortxt = "\033[38;2;" + HexToDec("FF") + ";" + HexToDec("a0") + ";" + HexToDec("7a") + "m"
		case "indianred":
			colortxt = "\033[38;2;" + HexToDec("cd") + ";" + HexToDec("5c") + ";" + HexToDec("5c") + "m"
		case "crimson":
			colortxt = "\033[38;2;" + HexToDec("dc") + ";" + HexToDec("14") + ";" + HexToDec("3c") + "m"
		case "canard":
			colortxt = "\033[38;2;" + HexToDec("04") + ";" + HexToDec("8b") + ";" + HexToDec("9a") + "m"
		case "pupute":
			colortxt = "\033[38;2;" + HexToDec("FF") + ";" + HexToDec("66") + ";" + HexToDec("D1") + "m"
		case "whitecream":
			colortxt = "\033[38;2;" + HexToDec("Fd") + ";" + HexToDec("f1") + ";" + HexToDec("b8") + "m"
		case "cacadoie":
			colortxt = "\033[38;2;" + HexToDec("cd") + ";" + HexToDec("cd") + ";" + HexToDec("0d") + "m"
		case "lightorange":
			colortxt = "\033[38;2;" + HexToDec("FF") + ";" + HexToDec("AB") + ";" + HexToDec("91") + "m"
		default:
			fmt.Println("Not a valid color")
			return
		}
	}
	standardtxt := []string{}
	banner := string(os.Args[len(os.Args)-1])
	if banner == "standard" || banner == "shadow" || banner == "thinkertoy" || banner == "golden" {
		file, err := os.Open(os.Args[len(os.Args)-1] + ".txt")
		if err != nil {
			fmt.Println("The text file specified doesn't exist")
			os.Exit(2)
		}
		banner1++
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			standardtxt = append(standardtxt, fileScanner.Text())
		}
	} else {
		file, err := os.Open("standard.txt")
		if err != nil {
			fmt.Println("Cannot open the standard.txt file")
			os.Exit(2)
		}
		fileScanner := bufio.NewScanner(file)
		for fileScanner.Scan() {
			standardtxt = append(standardtxt, fileScanner.Text())
		}
	}
	standardtxt = append(standardtxt, " ")
	text := ""
	tabword := []string{}
	tab := []byte{}
	substring := 0
	/*fmt.Println(color)
	fmt.Println(banner1)
	fmt.Println(len(os.Args))*/
	if color == 0 && banner1 == 0 && len(os.Args) == 2 {
		condition++
	} else if color == 0 && banner1 == 0 && len(os.Args) == 3 {
		condition = condition + 2
	} else if color != 0 && banner1 != 0 && len(os.Args) == 4 {
		condition = condition + 2
	} else if color != 0 && banner1 == 0 && len(os.Args) == 3 {
		condition = condition + 2
	} else if color == 0 && banner1 != 0 && len(os.Args) == 3 {
		condition++
	} else if color != 0 && banner1 != 0 && len(os.Args) == 5 {
		condition = condition + 3
		substring++
	} else if color == 0 && banner1 != 0 && len(os.Args) == 4 {
		condition = condition + 2
		substring++
	} else if color != 0 && banner1 == 0 && len(os.Args) == 4 {
		condition = condition + 3
		substring++
	} else {
		fmt.Println("Wrong number of arguments")
		return
	}
	//fmt.Println(substring)
	text = string(os.Args[condition])
	tabword = strings.Split(os.Args[condition], "\\n")
	tab = []byte(os.Args[condition])
	tabascii := []string{}
	for i := 32; i <= 126; i++ {
		tabascii = append(tabascii, string(rune(i)))
	}
	const constant = 9
	finalstring := ""
	tabInt := []int{}
	if substring != 0 {
		re := regexp.MustCompile(os.Args[condition-1])
		AllIndex := re.FindAllIndex(tab, -1)
		if strings.Contains(text, os.Args[condition-1]) {
			for i := 0; i < len(AllIndex); i++ {
				for j := AllIndex[i][0]; j < AllIndex[i][0]+len(os.Args[condition-1]); j++ {
					tabInt = append(tabInt, j)
				}
			}
		}
	} else if substring == 0 && color != 0 {
		for i := 0; i < len(os.Args[condition]); i++ {
			tabInt = append(tabInt, i)
		}
	}
	compteur := 0
	tricks := 0
	for h := 0; h < len(tabword); h++ {
		for i := 0; i <= 8; i++ {
			if h == 1 {
				compteur = len(tabword[h-1]) + 2
			} else if h == 2 {
				compteur = len(tabword[h-1]) + len(tabword[h-2]) + 4
			} else {
				compteur = 0
			}
			for j := 0; j < len(tabword[h]); j++ {
				for k := 0; k < len(tabascii); k++ {
					for l := 0; l < len(tabInt); l++ {
						if string(tabword[h][j]) == tabascii[k] && compteur == tabInt[l] {
							finalstring = standardtxt[k*constant+i+1]
							fmt.Print(colortxt + finalstring + reset)
							tricks++
						}
					}
					if string(tabword[h][j]) == tabascii[k] && tricks == 0 {
						finalstring = standardtxt[k*constant+i+1]
						fmt.Print(finalstring)
					}
				}
				compteur++
				tricks = 0
			}
			if i != 8 {
				fmt.Println("")
			}
		}
	}
}

func HexToDec(s string) string {
	Dec, error := strconv.ParseInt(s, 16, 64)
	if error != nil {
		return "Convertion impossible"
	}
	return strconv.Itoa(int(Dec))
}
