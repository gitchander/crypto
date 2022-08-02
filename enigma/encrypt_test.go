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

func TestSamples(t *testing.T) {

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
					Plaintext: LinesToText([]string{
						"THES TECK ERVE RBIN DUNG ENPL UGBO ARDI SANA DDED LAYE ROFS",
						"ECUR ITYW HICH CONS ISTS OFWI RESW HICH PLUG INTO INTO SOCK",
						"ETSO NTHE FRON TOFT HEEN IGMA MACH INEE ACHW IREC ONNE CTSL",
						"ETTE RSEG PTOO THES EPAI RING SARE SPEC IFIE DASP ARTO FTHE",
						"KEYM ATER IALW HENA LETT ERIS TYPE DBEF OREI TGOE SINT OTHE",
						"FIRS TROT ORIT UNDE RGOE STHE SUBS TITU TION ACCO RDIN GTOT",
						"HEPL UGBO ARDT HENA FTER THEL ETTE RCOM ESOU TITI SPUT THRO",
						"UGHT HEPL UGBO ARDS UBST ITUT IONA GAIN BEFO REBE INGO UTPU",
						"TANE XAMP LEPL UGBO ARDS ETTI NGIS ASFO LLOW SPOM LIUK JNHY",
						"TGBV FREA CTHI SMEA NSPA NDOA RESW APPE DMAN DLAR ESWA PPED",
						"ETCI FWEU SETH EEXA MPLE ABOV EWHE RETH ELET TERA WASE NCRY",
						"PTED WITH ROTO RSII IAND IIIW ITHT HEST ARTP OSIT IONS AAAW",
						"EHAD THEL ETTE RAEN CRYP TEDA SAUI FWEN OWTA KEIN TOAC COUN",
						"TTHE PLUG BOAR DUSI NGTH EPLU GBOA RDSE TTIN GSIN THEP REVI",
						"OUSP ARAG RAPH THEA ISFI RSTT RANS LATE DTOA CBEF OREE NCIP",
						"HERM ENTE NCIP HERM ENTC ONTI NUES ASUS UALT HIST IMET HECI",
						"SOUT PUTA SAJT HISL ETTE RIST HENR OUTE DTHR OUGH THEP LUGB",
						"OARD AGAI NTOB ESUB STIT UTED WITH KSON OWWE HAVE ANAB EING",
						"ENCI PHER EDAS AKWI THTH EPLU GBOA RDIN USET HEPL UGBO ARDS",
						"IGNI FICA NTLY INCR EASE STHE STRE NGTH OFTH EENI GMAC IPHE",
						"RASA WHOL EMOR ETHA NADD INGA NOTH ERRO TORC OULD",
					}),
					Ciphertext: LinesToText([]string{
						"JEFO JGZC XPAK WXLS ZHPP GKGW FWRZ VZNH FCJD YYQH SIUZ ZCUD",
						"TOGI TKJF IKEB WFGC UTVK KICE ODRT QVKW NYNU LALG UIZV LRHX",
						"TMUM CSUT QPFI ASMC DSNB RHSQ VKRX HGPZ OJMG VABO WIYP MCAF",
						"UDQV AMWC FIRY JLIO WGWY WSUF LLQQ FNJG ZRPN HFTG PPPH AAVV",
						"CMMV QMCI LQFN GCZC DOJP BSQZ BVYP ELQQ ACQP WDBJ AGKI CUMO",
						"HAGV CYBR ILTZ HUQS PXSV UDXR UZCL JWJF FWVF YBHY GCHD VUID",
						"JIEP MBHA TOII JITB JVBE DAUD JVQP OGNZ PTRC REQQ KDXM EQMM",
						"RJEQ ITFC SEZU PNUO KWCM BFKP TBKH TXEX ZOJY ENGC AYOU MLOM",
						"VMRH TCQU TRET BDZD DAMG OFKH TNGI JOXI MFFL FMIY RCLG NDUD",
						"GEDG XFGQ VDES IDIC WNVO THXT LSFN LQKR ILHT NMBV WZTR JMPK",
						"FGKZ ALAV AKGA JZTY BVBZ ECTA QDVN NPXJ QSQZ ENZD BBEZ URNO",
						"OZIK VJZV WRIU JHFL MKPS UKVI FKUF NAFO GGPK EOUD KKUH HXJI",
						"RIOQ AZWD SNDC SNJZ BYCV UFUS PMGY GRLO QUII EIMM HPYA VAQU",
						"JGSJ TJXD LYXL MVRQ PEHY JDUV CUTB YJFM NCBF UAUI VGZX SYEX",
						"RQXG PHLH EVBU XACK SCVL SBIR JTJB NZBH MLSI MGAJ POQD YLJS",
						"ESON YCIT IXWJ OLLI WMIQ XQRE ELTJ OWAC TUBR THEK GZYS XVEV",
						"IPVL GOOS DECO QZJA KDWI ZHZN XILS NTSD JRMJ RCEB NTHG XIEW",
						"YMGT CFOM RCGZ WHDW YMHY YHWZ NMCG MCMW FHYU ZYGB NRBZ TXPM",
						"CHIV YNRB VAPW GYTU JDQX VUOJ ZXMG SIUK MJBD RWFX QDDX CJZF",
						"JQPR RTUD AGDH PJSE LNYT EFLB HKCJ JNDF RZNB ZMFS PXQN TVSV",
						"NMBZ LICE FOJI LFSV LCIO GIUQ BVKM GSAE OKMI XWUO",
					}),
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
