#!/bin/bash
echo ASCII ART COLOR
echo "----------------------------------------------------------"

echo Test 1: Color: "red" text: "ahello world"
go run main.go --color=red "hello world"

echo "----------------------------------------------------------"

echo Test 2: Color: "green" text: "1 + 1 = 2"
go run main.go --color=green "1 + 1 = 2"

echo "----------------------------------------------------------"

echo Test 3: Color: "yellow" text: "(%&) ??"
go run main.go --color=yellow "(%&) ??"

echo "----------------------------------------------------------"

echo Test 4: Color: "orange" subtext: "GuYs" text: "HeY GuYs"
go run main.go --color=orange GuYs "HeY GuYs"

echo "----------------------------------------------------------"

echo Test 5: Color: "blue" subtext: "B" text: "RGB()"
go run main.go --color=blue B 'RGB()'

echo "----------------------------------------------------------"

echo Test 6: Color: "Yellow" subtext: "kit" text: "a king kitten have kit"
go run main.go --color=yellow kit "a king kitten have kit"

echo "----------------------------------------------------------"

echo Test 7: Color: "Green" subtext: "klmno" text: "abcdefghijklmnopqrstuvwxyz"
go run main.go --color=green klmno "abcdefghijklmnopqrstuvwxyz"

echo "----------------------------------------------------------"

echo Test 8: Color: "Red" subtext: "Color" text: "Ascii Art Color" banner: "shadow"
go run main.go --color=red Color "Ascii Art Color" shadow

echo "----------------------------------------------------------"

echo Test 9: Color: "Cyan" text: "Hello There" banner: "shadow"
go run main.go --color=cyan "Hello There" shadow

echo "----------------------------------------------------------"

echo Test 10: Color: "Magenta" subtext: "01" text: "Zone 01" banner: "thinkertoy"
go run main.go --color=magenta 01 "Zone01 Code 0110" thinkertoy

echo "----------------------------------------------------------"

echo Test 11: Color: "pupute" subtext: "O" text: "~(OwO)~" banner: "thinkertoy"
go run main.go --color=pupute O "~(OwO)~ ~(UwU)~ ~(OwO)~" thinkertoy

echo "----------------------------------------------------------"

echo Test 12: Color: "Green" subtext: "1a" text: '1a\#FdwHywR&/()=' banner: "standard"
go run main.go --color=green 1a "'1a\#FdwHy1a\wR&/()='" standard

echo "----------------------------------------------------------"

echo Test 13: Color: "lightsalmon" subtext: "LOWER" text: "upper and LOWER" banner: "shadow"
go run main.go --color=lightsalmon LOWER "upper and LOWER" shadow

echo "----------------------------------------------------------"

echo Test 14: Color: "Cyan" subtext: "e" text: "Hello There" banner: "standard"
go run main.go --color=cyan e "Hello There" standard

echo "----------------------------------------------------------"

echo Test 15: Color: "Magenta" subtext: "is" text: "This is a test" banner: "standard"
go run main.go --color=magenta is "This is a test" 

echo "----------------------------------------------------------"

echo Test 16: Color: "Canard" subtext: "ed" text: "Fred fed Ted bread" banner: "thinkertoy"
go run main.go --color=canard ed "Fred fed Ted bread" thinkertoy 

echo "----------------------------------------------------------"

echo Test 17: Color: "Red" subtext: "sn" text: "A snake sneaks to seek a snack." banner: "standard"
go run main.go --color=red sn "A snake sneaks to seek a snack." 

echo "----------------------------------------------------------"

echo Test 18: Color: "rgb(84, 186, 227)" subtext: "The" text: "Hello There" banner: "shadow"
go run main.go '--color=rgb(84, 186, 227)' The "Hello There" shadow

echo "----------------------------------------------------------"

echo Test 19: Color: "#7e3dd4" subtext: "The" text: "She sells seashells on the seashore." banner: "standard"
go run main.go '--color=#7e3dd4' sea "She sells seashells on the seashore." standard

echo "----------------------------------------------------------"

echo Test 20: Color: "#4affab" subtext: "FGHIJKL" text: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
go run main.go '--color=#4affab' FGHIJKL "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

echo "----------------------------------------------------------"

echo Test 21: Color: "#4affab" subtext: "HER" text: "HELLO\nTHERE" banner: "golden"
go run main.go '--color=#ff82b0' HER "HELLO\nTHERE" golden

echo "----------------------------------------------------------"

echo Test 22: Color: "caca d'oie" subtext: "Caca" text: "Caca d'oie" banner: "standard"
go run main.go '--color=cacadoie' Caca "Caca d'oie" standard