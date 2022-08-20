package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gitchander/crypto/enigma"
	"github.com/gitchander/crypto/enigma/base16"
	"github.com/gitchander/crypto/enigma/base26"
)

func main() {
	testMarshalJSON()
	testDonitzMessage()
	testCompareStrings()
	testValidate()
	genCodeLines()
	testJoinLines()
	testUtf8Base16()
	testUtf8Base26()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
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

	ciphertext := enigma.JoinStrings([]string{
		"DUHF TETO LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT",
		"GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM BBGW HZAN",
		"VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT XSON PNYN QFUD BBHH",
		"VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII WZXI VIUQ DRHY MNCY EFUA PNHO",
		"TKHK GDNP SAKN UAGH JZSM JBMH VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH",
		"DBWV HDFY HJOQ IHOR TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML",
		"RECW WUTL RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW DUHF TETO",
	})
	ciphertext = trimMessageIndicator(ciphertext)
	ciphertext = enigma.OnlyLetters(ciphertext)

	e, err := enigma.New(c)
	checkError(err)

	plaintext := e.FeedString(ciphertext)

	fmt.Println("Dönitz message:")

	fmt.Println("ciphertext:", ciphertext)
	fmt.Println("plaintext: ", plaintext)
}

func trimMessageIndicator(s string) string {
	const messageIndicator = "DUHF TETO"
	s = strings.TrimPrefix(s, messageIndicator)
	s = strings.TrimSuffix(s, messageIndicator)
	return s
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

func testValidate() {
	err := enigma.ValidatePosition("Y")
	checkError(err)

	err = enigma.ValidateWiring("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	checkError(err)

	err = enigma.ValidatePlugboard("AB CD FG XR")
	checkError(err)
}

func genCodeLines() {
	st := "JEFOJGZCXPAKWXLSZHPPGKGWFWRZVZNHFCJDYYQHSIUZZCUDTOGITKJFIKEBWFGCUTVKKICEODRTQVKWNYNULALGUIZVLRHXTMUMCSUTQPFIASMCDSNBRHSQVKRXHGPZOJMGVABOWIYPMCAFUDQVAMWCFIRYJLIOWGWYWSUFLLQQFNJGZRPNHFTGPPPHAAVVCMMVQMCILQFNGCZCDOJPBSQZBVYPELQQACQPWDBJAGKICUMOHAGVCYBRILTZHUQSPXSVUDXRUZCLJWJFFWVFYBHYGCHDVUIDJIEPMBHATOIIJITBJVBEDAUDJVQPOGNZPTRCREQQKDXMEQMMRJEQITFCSEZUPNUOKWCMBFKPTBKHTXEXZOJYENGCAYOUMLOMVMRHTCQUTRETBDZDDAMGOFKHTNGIJOXIMFFLFMIYRCLGNDUDGEDGXFGQVDESIDICWNVOTHXTLSFNLQKRILHTNMBVWZTRJMPKFGKZALAVAKGAJZTYBVBZECTAQDVNNPXJQSQZENZDBBEZURNOOZIKVJZVWRIUJHFLMKPSUKVIFKUFNAFOGGPKEOUDKKUHHXJIRIOQAZWDSNDCSNJZBYCVUFUSPMGYGRLOQUIIEIMMHPYAVAQUJGSJTJXDLYXLMVRQPEHYJDUVCUTBYJFMNCBFUAUIVGZXSYEXRQXGPHLHEVBUXACKSCVLSBIRJTJBNZBHMLSIMGAJPOQDYLJSESONYCITIXWJOLLIWMIQXQREELTJOWACTUBRTHEKGZYSXVEVIPVLGOOSDECOQZJAKDWIZHZNXILSNTSDJRMJRCEBNTHGXIEWYMGTCFOMRCGZWHDWYMHYYHWZNMCGMCMWFHYUZYGBNRBZTXPMCHIVYNRBVAPWGYTUJDQXVUOJZXMGSIUKMJBDRWFXQDDXCJZFJQPRRTUDAGDHPJSELNYTEFLBHKCJJNDFRZNBZMFSPXQNTVSVNMBZLICEFOJILFSVLCIOGIUQBVKMGSAEOKMIXWUO"
	st = strings.ToUpper(st)
	inputText := enigma.OnlyLetters(st)
	var b strings.Builder
	const (
		lettersPerGroup = 4
		groupsPerLine   = 12

		lettersPerLine = lettersPerGroup * groupsPerLine
	)
	for i := 0; i < len(inputText); i++ {
		b.WriteByte(inputText[i])
		j := (i + 1)
		if (j % lettersPerLine) == 0 {
			b.WriteByte('\n')
		} else if (j % lettersPerGroup) == 0 {
			b.WriteByte(' ')
		}
	}
	result := b.String()
	fmt.Println(result)
	lines := strings.Split(result, "\n")

	fmt.Println("var lines = []string{")
	for _, line := range lines {
		fmt.Printf("%s%q,\n", "\t", line)
	}
	fmt.Println("}")

	outputText := enigma.OnlyLetters(enigma.JoinLines("", lines))
	if inputText != outputText {
		err := fmt.Errorf("%q != %q\n", inputText, outputText)
		checkError(err)
	}
}

func testJoinLines() {
	lines := []string{
		"DUHF TETO LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT",
		"GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM BBGW HZAN",
		"VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT XSON PNYN QFUD BBHH",
		"VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII WZXI VIUQ DRHY MNCY EFUA PNHO",
		"TKHK GDNP SAKN UAGH JZSM JBMH VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH",
		"DBWV HDFY HJOQ IHOR TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML",
		"RECW WUTL RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW DUHF TETO",
	}
	text := enigma.JoinLines("\t", lines)
	fmt.Print(text)
}

func testUtf8Base16() {
	text := "Hello, 世界"
	//text := "Привіт, світ!"
	plaintext := base16.EncodeToString([]byte(text))

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

	e, err := enigma.New(c)
	checkError(err)

	ciphertext := e.FeedString(plaintext)

	fmt.Println("plaintext:", plaintext)
	fmt.Println("ciphertext:", ciphertext)

	d, err := enigma.New(c)
	checkError(err)

	plaintextDecrypted := d.FeedString(ciphertext)
	bs, err := base16.DecodeString(plaintextDecrypted)
	checkError(err)
	resultText := string(bs)

	fmt.Println("resultText:", resultText)
}

func testUtf8Base26() {
	text := "Hello, 世界"
	plaintext := base26.EncodeToString([]byte(text))

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

	e, err := enigma.New(c)
	checkError(err)

	ciphertext := e.FeedString(plaintext)

	fmt.Println("plaintext:", plaintext)
	fmt.Println("ciphertext:", ciphertext)

	d, err := enigma.New(c)
	checkError(err)

	plaintextDecrypted := d.FeedString(ciphertext)
	bs, err := base26.DecodeString(plaintextDecrypted)
	checkError(err)
	resultText := string(bs)

	fmt.Println("resultText:", resultText)
}
