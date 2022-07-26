package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gitchander/crypto/enigma"
)

func main() {
	testMarshalJSON()
	testDonitzMessage()
	testCompareStrings()
}

func testMarshalJSON() {
	c := enigma.Config{
		Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
		Rotors: []enigma.RotorInfo{
			{
				ID:       "I",
				Ring:     "A",
				Position: "A",
			},
			{
				ID:       "II",
				Ring:     "A",
				Position: "A",
			},
			{
				ID:       "III",
				Ring:     "A",
				Position: "A",
			},
		},
		ReflectorID: "A",
	}

	data, err := json.MarshalIndent(c, "", "\t")
	checkError(err)

	fmt.Println(string(data))

	e, err := enigma.New(c)
	checkError(err)

	plaintext := "ABCDEFGH"
	ciphertext := e.FeedString(plaintext)

	fmt.Println("plaintext: ", plaintext)
	fmt.Println("ciphertext:", ciphertext)
}

// Dönitz message:
// https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030681
// https://www.cryptomuseum.com/crypto/enigma/msg/p1030681.htm
func testDonitzMessage() {
	c := enigma.Config{
		Plugboard: "AE BF CM DQ HU JN LX PR SZ VW",
		Rotors: []enigma.RotorInfo{
			{
				ID:       "Beta",
				Ring:     "A",
				Position: "Y",
			},
			{
				ID:       "V",
				Ring:     "A",
				Position: "O",
			},
			{
				ID:       "VI",
				Ring:     "E",
				Position: "S",
			},
			{
				ID:       "VIII",
				Ring:     "L",
				Position: "Z",
			},
		},
		ReflectorID: "C-thin",
	}

	ciphertext := joinLines([]string{
		"DUHF TETO LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT",
		"GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM BBGW HZAN",
		"VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT XSON PNYN QFUD BBHH",
		"VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII WZXI VIUQ DRHY MNCY EFUA PNHO",
		"TKHK GDNP SAKN UAGH JZSM JBMH VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH",
		"DBWV HDFY HJOQ IHOR TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML",
		"RECW WUTL RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW DUHF TETO",
	})
	ciphertext = trimMessageIndicator(ciphertext)
	ciphertext = onlyLetters(ciphertext)

	e, err := enigma.New(c)
	checkError(err)

	plaintext := e.FeedString(ciphertext)

	fmt.Println("Dönitz message:")

	fmt.Println("ciphertext:", ciphertext)
	fmt.Println("plaintext: ", plaintext)
}

func joinLines(lines []string) string {
	var b strings.Builder
	for _, line := range lines {
		b.WriteString(line)
	}
	return b.String()
}

func trimMessageIndicator(s string) string {
	const messageIndicator = "DUHF TETO"
	s = strings.TrimPrefix(s, messageIndicator)
	s = strings.TrimSuffix(s, messageIndicator)
	return s
}

func onlyLetters(s string) string {
	as := []byte(s)
	bs := make([]byte, 0, len(as))
	for _, a := range as {
		if byteIsLetter(a) {
			bs = append(bs, a)
		}
	}
	return string(bs)
}

func byteIsLetter(b byte) bool {
	if ('A' <= b) && (b <= 'Z') {
		return true
	}
	if ('a' <= b) && (b <= 'z') {
		return true
	}
	return false
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testCompareStrings() {
	a := "KRKRALLEXXFOLGENDESISTSOFORTBEKANNTZUGEBENXXICHHABEFOLGELNBEBEFEHLERHALTENXXJANSTERLEDESBISHERIGXNREICHSMARSCHALLSJGOERINGJSETZTDERFUEHRERSIEYHVRRGRZSSADMIRALYALSSEINENNACHFOLGEREINXSCHRIFTLSCHEVOLLMACHTUNTERWEGSXABSOFORTSOLLENSIESAEMTLICHEMASSNAHMENVERFUEGENYDIESICHAUSDERGEGENWAERTIGENLAGEERGEBENXGEZXREICHSLEITEIKKTULPEKKJBORMANNJXXOBXDXMMMDURNHFKSTXKOMXADMXUUUBOOIEXKP"

	b := "KRKRALLEXXFOLGENDESISTSOFORTBEKANNTZUGEBENXXICHHABEFOLGENDENBEFEHLERHALTENXXJANSTERLEDESBISHERIGXNREICHSMARSCHALLSJGOERINGJSETZTDERFUEHRERSIEYHVRRGRZSSADMIRALYALSSEINENNACHFOLGEREINXSCHRIFTLSCHEVOLLMACHTUNTERWEGSXABSOFORTSOLLENSIESAEMTLICHEMASSNAHMENVERFUEGENYDIESICHAUSDERGEGENWAERTIGENLAGEERGEBENXGEZXREICHSLEITEIKKTULPEKKJBORMANNJXXOBXDXMMMDURNHFKSTXKOMXADMXUUUBOOIEXKP"
	//b := "KRKRALLEXXFOLGENDESISTSOFORTBEKANNTZUGEBENXXICHHABEFOLGELNBEBEFEHLERHALTENXXJANSTERLEDESBISHERIGXNREICHSMARSCHALLSJGOERINGJSETZTDERFUEHRERSIEYHVRRGRZSSADMIRALYALSSEINENNACHFOLGEREINXSCHRIFTLSCHEVOLLMACHTUNTERWEGSXABSOFORTSOLLENSIESAEMTLICHEMASSNAHMENVERFUEGENYDIESICHAUSDERGEGENWAERTIGENLAGEERGEBENXGEZXREICHSLEITEIKKTULPEKKJBORMANNJXXOBXDXMMMDURNHFKSTXKOMXADMXUUUBOOIEXKP"

	fmt.Println(a == b)

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			fmt.Printf("%d: (%#U) != (%#U)\n", i, a[i], b[i])
		}
	}
}
