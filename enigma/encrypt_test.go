package enigma

import (
	"testing"
)

type TestPair struct {
	Plaintext  string
	Ciphertext string
}

type TestSample struct {
	Config Config
	Pairs  []TestPair
}

var samples = []TestSample{
	{
		Config: Config{
			Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
			Rotors: []RotorInfo{
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
		},
		Pairs: []TestPair{
			{
				Plaintext:  "AAAAAAAAAAAA",
				Ciphertext: "OOWKOENOJSKL",
			},
		},
	},
	{
		Config: Config{
			Plugboard: "AV BS CG DL FU HZ IN KM OW RX",
			Rotors: []RotorInfo{
				{
					ID:       "I",
					Ring:     "A",
					Position: "A",
				},
				{
					ID:       "II",
					Ring:     "A",
					Position: "B",
				},
				{
					ID:       "III",
					Ring:     "A",
					Position: "C",
				},
			},
			ReflectorID: "A",
		},
		Pairs: []TestPair{
			{
				Plaintext:  "HelloWorld",
				Ciphertext: "MFBXQLVGHA",
			},
		},
	},
	{
		Config: Config{
			Plugboard: "PO ML IU KJ NH YT GB VF RE DC",
			Rotors: []RotorInfo{
				{
					ID:       "I",
					Ring:     "C",
					Position: "S",
				},
				{
					ID:       "II",
					Ring:     "D",
					Position: "H",
				},
				{
					ID:       "III",
					Ring:     "W",
					Position: "J",
				},
			},
			ReflectorID: "B",
		},
		Pairs: []TestPair{
			{
				Plaintext:  "ABCDEF",
				Ciphertext: "DFRHYR",
			},
			{
				Plaintext:  "THESTECKERVERBINDUNGENPLUGBOARDISANADDEDLAYEROFSECURITYWHICHCONSISTSOFWIRESWHICHPLUGINTOINTOSOCKETSONTHEFRONTOFTHEENIGMAMACHINEEACHWIRECONNECTSLETTERSEGPTOOTHESEPAIRINGSARESPECIFIEDASPARTOFTHEKEYMATERIALWHENALETTERISTYPEDBEFOREITGOESINTOTHEFIRSTROTORITUNDERGOESTHESUBSTITUTIONACCORDINGTOTHEPLUGBOARDTHENAFTERTHELETTERCOMESOUTITISPUTTHROUGHTHEPLUGBOARDSUBSTITUTIONAGAINBEFOREBEINGOUTPUTANEXAMPLEPLUGBOARDSETTINGISASFOLLOWSPOMLIUKJNHYTGBVFREACTHISMEANSPANDOARESWAPPEDMANDLARESWAPPEDETCIFWEUSETHEEXAMPLEABOVEWHERETHELETTERAWASENCRYPTEDWITHROTORSIIIANDIIIWITHTHESTARTPOSITIONSAAAWEHADTHELETTERAENCRYPTEDASAUIFWENOWTAKEINTOACCOUNTTHEPLUGBOARDUSINGTHEPLUGBOARDSETTINGSINTHEPREVIOUSPARAGRAPHTHEAISFIRSTTRANSLATEDTOACBEFOREENCIPHERMENTENCIPHERMENTCONTINUESASUSUALTHISTIMETHECISOUTPUTASAJTHISLETTERISTHENROUTEDTHROUGHTHEPLUGBOARDAGAINTOBESUBSTITUTEDWITHKSONOWWEHAVEANABEINGENCIPHEREDASAKWITHTHEPLUGBOARDINUSETHEPLUGBOARDSIGNIFICANTLYINCREASESTHESTRENGTHOFTHEENIGMACIPHERASAWHOLEMORETHANADDINGANOTHERROTORCOULD",
				Ciphertext: "JEFOJGZCXPAKWXLSZHPPGKGWFWRZVZNHFCJDYYQHSIUZZCUDTOGITKJFIKEBWFGCUTVKKICEODRTQVKWNYNULALGUIZVLRHXTMUMCSUTQPFIASMCDSNBRHSQVKRXHGPZOJMGVABOWIYPMCAFUDQVAMWCFIRYJLIOWGWYWSUFLLQQFNJGZRPNHFTGPPPHAAVVCMMVQMCILQFNGCZCDOJPBSQZBVYPELQQACQPWDBJAGKICUMOHAGVCYBRILTZHUQSPXSVUDXRUZCLJWJFFWVFYBHYGCHDVUIDJIEPMBHATOIIJITBJVBEDAUDJVQPOGNZPTRCREQQKDXMEQMMRJEQITFCSEZUPNUOKWCMBFKPTBKHTXEXZOJYENGCAYOUMLOMVMRHTCQUTRETBDZDDAMGOFKHTNGIJOXIMFFLFMIYRCLGNDUDGEDGXFGQVDESIDICWNVOTHXTLSFNLQKRILHTNMBVWZTRJMPKFGKZALAVAKGAJZTYBVBZECTAQDVNNPXJQSQZENZDBBEZURNOOZIKVJZVWRIUJHFLMKPSUKVIFKUFNAFOGGPKEOUDKKUHHXJIRIOQAZWDSNDCSNJZBYCVUFUSPMGYGRLOQUIIEIMMHPYAVAQUJGSJTJXDLYXLMVRQPEHYJDUVCUTBYJFMNCBFUAUIVGZXSYEXRQXGPHLHEVBUXACKSCVLSBIRJTJBNZBHMLSIMGAJPOQDYLJSESONYCITIXWJOLLIWMIQXQREELTJOWACTUBRTHEKGZYSXVEVIPVLGOOSDECOQZJAKDWIZHZNXILSNTSDJRMJRCEBNTHGXIEWYMGTCFOMRCGZWHDWYMHYYHWZNMCGMCMWFHYUZYGBNRBZTXPMCHIVYNRBVAPWGYTUJDQXVUOJZXMGSIUKMJBDRWFXQDDXCJZFJQPRRTUDAGDHPJSELNYTEFLBHKCJJNDFRZNBZMFSPXQNTVSVNMBZLICEFOJILFSVLCIOGIUQBVKMGSAEOKMIXWUO",
			},
		},
	},
	{
		Config: Config{
			Plugboard: "AB CD EF GH IJ KL",
			Rotors: []RotorInfo{
				{
					ID:       "I",
					Ring:     "D",
					Position: "A",
				},
				{
					ID:       "II",
					Ring:     "E",
					Position: "B",
				},
				{
					ID:       "III",
					Ring:     "F",
					Position: "C",
				},
			},
			ReflectorID: "B",
		},
		Pairs: []TestPair{
			{
				Plaintext:  "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
				Ciphertext: "TCNXDQYWKDEKCQKLLDNCEEIAJI",
			},
		},
	},
}

func TestSamples(t *testing.T) {
	for i, sample := range samples {
		for j, pair := range sample.Pairs {
			e, err := New(sample.Config)
			if err != nil {
				t.Fatal(err)
			}
			ciphertext := e.FeedString(pair.Plaintext)
			if ciphertext != pair.Ciphertext {
				t.Fatalf("(sample %d, pair %d) invalid ciphertext: have %q, want %q",
					i, j, ciphertext, pair.Ciphertext)
			}
		}
	}
}

func TestIncludeForeign(t *testing.T) {
	var samples = []TestSample{
		{
			Config: Config{
				Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
				Rotors: []RotorInfo{
					{
						ID:       "VI",
						Ring:     "A",
						Position: "A",
					},
					{
						ID:       "I",
						Ring:     "A",
						Position: "Q",
					},
					{
						ID:       "III",
						Ring:     "A",
						Position: "L",
					},
				},
				ReflectorID: "B",
			},
			Pairs: []TestPair{
				{
					Plaintext:  "Hello, World!",
					Ciphertext: "GKTWX, GEGZQ!",
				},
				{
					Plaintext:  "The Enigma cipher machine is well known for the vital role it played during WWII. Alan Turing and his attempts to crack the Enigma machine code changed history. Nevertheless, many messages could not be decrypted until today.",
					Ciphertext: "XAS FHVLAQ AWTDIH WGLYCQT GO JFIR RPSQC JPQ KAU NPCYW MEAH OG GKBGKW WSZMVI XTWX. MGKS GTFVJY NJZ YTB FBUHOUWQ UE FBIFN HEY UMUPHW EMJZBYV KJSW ZRCQZYF JZNXMNA. OFSUHNFDMXQV, CHUH GQFMYEVE VEWCC JHI FB YXBZCIEIC CHZEU UWEOL.",
				},
			},
		},
	}
	for i, sample := range samples {
		for j, pair := range sample.Pairs {
			e, err := New(sample.Config)
			if err != nil {
				t.Fatal(err)
			}
			ciphertext := e.FeedIncludeForeign(pair.Plaintext)
			if ciphertext != pair.Ciphertext {
				t.Fatalf("(sample %d, pair %d) invalid ciphertext: have %q, want %q",
					i, j, ciphertext, pair.Ciphertext)
			}
		}
	}
}

func TestIgnoreForeign(t *testing.T) {
	var samples = []TestSample{
		{
			Config: Config{
				Plugboard: "BQ CR DI EJ KW MT OS PX UZ GH",
				Rotors: []RotorInfo{
					{
						ID:       "VI",
						Ring:     "A",
						Position: "A",
					},
					{
						ID:       "I",
						Ring:     "A",
						Position: "Q",
					},
					{
						ID:       "III",
						Ring:     "A",
						Position: "L",
					},
				},
				ReflectorID: "B",
			},
			Pairs: []TestPair{
				{
					Plaintext:  "Hello, World!",
					Ciphertext: "GKTWX GEGZQ",
				},
				{
					Plaintext:  "The Enigma cipher machine is well known for the vital role it played during WWII. Alan Turing and his attempts to crack the Enigma machine code changed history. Nevertheless, many messages could not be decrypted until today.",
					Ciphertext: "XASFH VLAQA WTDIH WGLYC QTGOJ FIRRP SQCJP QKAUN PCYWM EAHOG GKBGK WWSZM VIXTW XMGKS GTFVJ YNJZY TBFBU HOUWQ UEFBI FNHEY UMUPH WEMJZ BYVKJ SWZRC QZYFJ ZNXMN AOFSU HNFDM XQVCH UHGQF MYEVE VEWCC JHIFB YXBZC IEICC HZEUU WEOL",
				},
			},
		},
	}
	for i, sample := range samples {
		for j, pair := range sample.Pairs {
			e, err := New(sample.Config)
			if err != nil {
				t.Fatal(err)
			}
			ciphertext := e.FeedIgnoreForeign(pair.Plaintext)
			if ciphertext != pair.Ciphertext {
				t.Fatalf("(sample %d, pair %d) invalid ciphertext: have %q, want %q",
					i, j, ciphertext, pair.Ciphertext)
			}
		}
	}
}
