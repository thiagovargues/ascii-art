#!/usr/bin/env bash

# Tests ascii-art - audit Functional

echo 'Test 1: "hello"'
go run . "hello" | cat -e
echo

echo 'Test 2: "HELLO"'
go run . "HELLO" | cat -e
echo

echo 'Test 3: "HeLlo HuMaN"'
go run . "HeLlo HuMaN" | cat -e
echo

echo 'Test 4: "1Hello 2There"'
go run . "1Hello 2There" | cat -e
echo

echo 'Test 5: "Hello\nThere"'
go run . 'Hello\nThere' | cat -e
echo

echo 'Test 6: "Hello\n\nThere"'
go run . 'Hello\n\nThere' | cat -e
echo

echo 'Test 7: "{Hello & There #}"'
go run . "{Hello & There #}" | cat -e
echo

echo "Test 8: 'hello There 1 to 2!'"
go run . 'hello There 1 to 2!' | cat -e
echo

echo 'Test 9: "MaD3IrA&LiSboN"'
go run . "MaD3IrA&LiSboN" | cat -e
echo

echo 'Test 10: "1a\"#FdwHywR&/()="'
go run . '1a"#FdwHywR&/()=' | cat -e
echo

echo 'Test 11: "{|}~"'
go run . "{|}~" | cat -e
echo

echo 'Test 12: "[\]^_ '\''a"'
go run . "[\\]^_ 'a" | cat -e
echo

echo 'Test 13: "RGB"'
go run . "RGB" | cat -e
echo

echo 'Test 14: ":;<=>?@"'
go run . ":;<=>?@" | cat -e
echo

echo "Test 15: '\\!\" #\$%&'\"'\"'()*+,-./'"
go run . '\!" #$%&'"'"'()*+,-./' | cat -e
echo

echo 'Test 16: "ABCDEFGHIJKLMNOPQRSTUVWXYZ"'
go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e
echo

echo 'Test 17: "abcdefghijklmnopqrstuvwxyz"'
go run . "abcdefghijklmnopqrstuvwxyz" | cat -e
echo

# Tests random demandés par l’audit

echo 'Random 1 (>=4 lower, >=3 upper): "abcdnOPQe"'
go run . "abcdnOPQe" | cat -e
echo

echo 'Random 2 (>=5 lower, space, 2 numbers): "abcde 12f"'
go run . "abcde 12f" | cat -e
echo

echo 'Random 3 (>=1 upper, 3 specials): "Ab!@#cd"'
go run . "Ab!@#cd" | cat -e
echo

echo 'Random 4 (>=2 lower, 2 spaces, 1 number, 2 specials, 3 upper): "ab  3!@ABC"'
go run . "ab  3!@ABC" | cat -e
echo

echo "Tous les tests Functional de l'audit ont été lancés."
