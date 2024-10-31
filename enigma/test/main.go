package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gitchander/crypto/enigma"
	"github.com/gitchander/crypto/enigma/base16"
	"github.com/gitchander/crypto/enigma/base26"
	"github.com/gitchander/crypto/utils/random"
)

func main() {
	// testCompareStrings()
	// testJoinLines()

	//------------------------------------------

	fs := []func() error{
		testMarshalJSON,
		testDonitzMessage,
		testDonitzMessageV2,
		genCodeLines,
		testUtf8Base16,
		testUtf8Base26,
		testAnyBytes,
		testAnyBytes2,
		testFormatText,
	}

	for _, f := range fs {
		checkError(f())
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func testMarshalJSON() error {
	c := enigma.Config{
		Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
		Rotors: enigma.RotorsConfig{
			IDs:       "I II III",
			Rings:     "AAA",
			Positions: "AAA",
		},
		Reflector: "A",
	}

	data, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	e, err := enigma.New(c)
	if err != nil {
		return err
	}
	tf := enigma.TrueTextFeeder(e)

	plaintext := "ABCDEFGH"
	ciphertext, err := tf.FeedText(plaintext)
	if err != nil {
		return err
	}

	fmt.Println("plaintext: ", plaintext)
	fmt.Println("ciphertext:", ciphertext)

	return nil
}

// Dönitz message:
// https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030681
// https://www.cryptomuseum.com/crypto/enigma/msg/p1030681.htm
func testDonitzMessage() error {
	c := enigma.Config{
		Plugboard: "AE BF CM DQ HU JN LX PR SZ VW",
		Rotors: enigma.RotorsConfig{
			IDs:       "Beta V VI VIII",
			Rings:     "AAEL",
			Positions: "YOSZ",
		},
		Reflector: "C-thin",
	}

	ciphertext := enigma.JoinStrings(
		"DUHF TETO LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT",
		"GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM BBGW HZAN",
		"VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT XSON PNYN QFUD BBHH",
		"VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII WZXI VIUQ DRHY MNCY EFUA PNHO",
		"TKHK GDNP SAKN UAGH JZSM JBMH VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH",
		"DBWV HDFY HJOQ IHOR TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML",
		"RECW WUTL RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW DUHF TETO",
	)
	ciphertext = trimMessageIndicator(ciphertext)
	ciphertext = enigma.OnlyLetters(ciphertext)

	e, err := enigma.New(c)
	if err != nil {
		return err
	}
	tf := enigma.TrueTextFeeder(e)

	plaintext, err := tf.FeedText(ciphertext)
	if err != nil {
		return err
	}

	fmt.Println("Dönitz message:")

	fmt.Println("ciphertext:", ciphertext)
	fmt.Println("plaintext: ", plaintext)

	return nil
}

func testDonitzMessageV2() error {
	c := enigma.Config{
		Plugboard: "AE BF CM DQ HU JN LX PR SZ VW",
		Rotors: enigma.RotorsConfig{
			IDs:       "Beta V VI VIII",
			Rings:     "AAEL",
			Positions: "YOSZ",
		},
		Reflector: "C-thin",
	}

	ciphertext := `
	DUHF TETO LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT
	GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM BBGW HZAN
	VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT XSON PNYN QFUD BBHH
	VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII WZXI VIUQ DRHY MNCY EFUA PNHO
	TKHK GDNP SAKN UAGH JZSM JBMH VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH
	DBWV HDFY HJOQ IHOR TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML
	RECW WUTL RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW DUHF TETO
`
	ciphertext = trimMessageIndicator(ciphertext)

	e, err := enigma.New(c)
	if err != nil {
		return err
	}
	tf := enigma.IncludeForeign(e)

	plaintext, err := tf.FeedText(ciphertext)
	if err != nil {
		return err
	}

	fmt.Println("Dönitz message:")

	fmt.Println("ciphertext:", ciphertext)
	fmt.Println("plaintext: ", plaintext)

	ciphertext2, err := tf.FeedText(plaintext)
	if err != nil {
		return err
	}

	fmt.Println("ciphertext2:", ciphertext2)

	fmt.Println(ciphertext == ciphertext2)

	return nil
}

func trimMessageIndicator(s string) string {
	const messageIndicator = "DUHF TETO"
	s = strings.Trim(s, "\n\t ")
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

func genCodeLines() error {
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

	outputText := enigma.OnlyLetters(enigma.JoinLines("", lines...))
	if inputText != outputText {
		err := fmt.Errorf("%q != %q\n", inputText, outputText)
		if err != nil {
			return err
		}
	}
	return nil
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
	text := enigma.JoinLines("\t", lines...)
	fmt.Print(text)
}

func testUtf8Base16() error {
	text := "Hello, 世界"
	//text := "Привіт, світ!"
	plaintext := base16.EncodeToString([]byte(text))

	c := enigma.Config{
		Plugboard: "AE BF CM DQ HU JN LX PR SZ VW",
		Rotors: enigma.RotorsConfig{
			IDs:       "Beta V VI VIII",
			Rings:     "AAEL",
			Positions: "YOSZ",
		},
		Reflector: "C-thin",
	}

	e, err := enigma.New(c)
	if err != nil {
		return err
	}
	tf := enigma.TrueTextFeeder(e)

	ciphertext, err := tf.FeedText(plaintext)
	if err != nil {
		return err
	}

	fmt.Println("plaintext:", plaintext)
	fmt.Println("ciphertext:", ciphertext)

	plaintextDecrypted, err := tf.FeedText(ciphertext)
	if err != nil {
		return err
	}

	bs, err := base16.DecodeString(plaintextDecrypted)
	if err != nil {
		return err
	}

	resultText := string(bs)

	fmt.Println("resultText:", resultText)
	return nil
}

func testUtf8Base26() error {
	text := "Hello, 世界"
	plaintext := base26.EncodeToString([]byte(text))

	c := enigma.Config{
		Plugboard: "AE BF CM DQ HU JN LX PR SZ VW",
		Rotors: enigma.RotorsConfig{
			IDs:       "Beta V VI VIII",
			Rings:     "AAEL",
			Positions: "YOSZ",
		},
		Reflector: "C-thin",
	}

	e, err := enigma.New(c)
	if err != nil {
		return err
	}
	tf := enigma.TrueTextFeeder(e)

	ciphertext, err := tf.FeedText(plaintext)
	if err != nil {
		return err
	}

	fmt.Println("plaintext:", plaintext)
	fmt.Println("ciphertext:", ciphertext)

	plaintextDecrypted, err := tf.FeedText(ciphertext)
	if err != nil {
		return err
	}

	bs, err := base26.DecodeString(plaintextDecrypted)
	if err != nil {
		return err
	}

	resultText := string(bs)

	fmt.Println("resultText:", resultText)
	return nil
}

func testAnyBytes() error {

	c := enigma.Config{
		Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
		Rotors: enigma.RotorsConfig{
			IDs:       "I II III",
			Rings:     "AAA",
			Positions: "AAA",
		},
		Reflector: "A",
	}

	r := random.NewRandNow()
	data := make([]byte, r.Intn(50))
	for i := 0; i < 100; i++ {
		plainBytesInput := data[:r.Intn(len(data)+1)]
		random.FillBytes(r, plainBytesInput)

		plainTextInput := base26.EncodeToString(plainBytesInput)

		e, err := enigma.New(c)
		if err != nil {
			return err
		}
		tf := enigma.TrueTextFeeder(e)

		cipherText, err := tf.FeedText(plainTextInput)
		if err != nil {
			return err
		}

		plainTextOutput, err := tf.FeedText(cipherText)
		if err != nil {
			return err
		}

		plainBytesOutput, err := base26.DecodeString(plainTextOutput)
		if err != nil {
			return err
		}

		if !(bytes.Equal(plainBytesInput, plainBytesOutput)) {
			err := fmt.Errorf("[%x] != [%x]\n", plainBytesInput, plainBytesOutput)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func testAnyBytes2() error {

	config := enigma.Config{
		Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
		Rotors: enigma.RotorsConfig{
			IDs:       "I II III",
			Rings:     "AAA",
			Positions: "AAA",
		},
		Reflector: "A",
	}
	e, err := enigma.New(config)
	if err != nil {
		return err
	}
	c := enigma.NewBytesCrypt(e)

	r := random.NewRandNow()
	data := make([]byte, r.Intn(1024))
	for i := 0; i < 1000; i++ {
		plainData1 := data[:r.Intn(len(data)+1)]
		random.FillBytes(r, plainData1)

		cipherData, err := c.Encrypt(plainData1)
		if err != nil {
			return err
		}

		//fmt.Printf("cipherData: %s\n", cipherData)

		plainData2, err := c.Decrypt(cipherData)
		if err != nil {
			return err
		}

		if !(bytes.Equal(plainData1, plainData2)) {
			err := fmt.Errorf("[%x] != [%x]\n", plainData1, plainData2)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func testFormatText() error {

	text := `
	DUHF TETO LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT
	GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM BBGW HZAN
	VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT XSON PNYN QFUD BBHH
	VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII WZXI VIUQ DRHY MNCY EFUA PNHO
	TKHK GDNP SAKN UAGH JZSM JBMH VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH
	DBWV HDFY HJOQ IHOR TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML
	RECW WUTL RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW DUHF TETO
`

	text = enigma.OnlyLetters(text)

	fmt.Println(text)

	ft := enigma.DefaultTextFormatter.FormatText(text)
	fmt.Println(ft)

	return nil
}
