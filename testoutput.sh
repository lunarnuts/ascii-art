go build

./ascii-art "hello" standard --output=test0.txt
cat test0.txt
rm test0.txt
echo -------------------------------------------------
./ascii-art "Hello There!" shadow --output=test1.txt
cat test1.txt
rm test1.txt
echo -------------------------------------------------
./ascii-art "First\nTest" shadow --output=test00.txt
cat test00.txt
rm test00.txt
echo -------------------------------------------------
./ascii-art "hello" standard --output=test01.txt
cat test01.txt
rm test01.txt
echo -------------------------------------------------
./ascii-art "123 -> #$%" standard --output=test02.txt
cat test02.txt
rm test02.txt
echo -------------------------------------------------
./ascii-art "432 -> #$%&@" shadow --output=test03.txt
cat test03.txt
rm test03.txt
echo -------------------------------------------------
./ascii-art "There" shadow --output=test04.txt
cat test04.txt
rm test04.txt
echo -------------------------------------------------
./ascii-art "123 -> \"#$%@" thinkertoy --output=test05.txt
cat test05.txt
rm test05.txt
echo -------------------------------------------------
./ascii-art "2 you" thinkertoy --output=test06.txt
cat test06.txt
rm test06.txt
echo -------------------------------------------------
./ascii-art "Testing long output!" standard --output=test07.txt
cat test07.txt
rm test07.txt
echo -------------------------------------------------
./ascii-art "qwertyZXCVB" standard --output=test08.txt
cat test08.txt
rm test08.txt
echo -------------------------------------------------
./ascii-art "asdfg 123456789" standard --output=test09.txt
cat test09.txt
rm test09.txt
echo -------------------------------------------------
./ascii-art "POIU ,./;'<>?:[]{}\"" standard --output=test10.txt
cat test10.txt
rm test10.txt
echo -------------------------------------------------
./ascii-art "gh  12 ]'/ RTY" standard --output=test11.txt
cat test11.txt
rm test11.txt