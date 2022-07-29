package enigma

import (
	"testing"
)

// Dönitz message:
// https://enigma.hoerenberg.com/index.php?cat=The%20U534%20messages&page=P1030681
// https://www.cryptomuseum.com/crypto/enigma/msg/p1030681.htm
func TestDonitzMessage(t *testing.T) {
	c := Config{
		Plugboard: "AE BF CM DQ HU JN LX PR SZ VW",
		Rotors: []RotorInfo{
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

	tp := TestPair{
		Ciphertext: LinesToText([]string{
			"LANO TCTO UARB BFPM HPHG CZXT DYGA HGUF XGEW KBLK GJWL QXXT",
			"GPJJ AVTO CKZF SLPP QIHZ FXOE BWII EKFZ LCLO AQJU LJOY HSSM",
			"BBGW HZAN VOII PYRB RTDJ QDJJ OQKC XWDN BBTY VXLY TAPG VEAT",
			"XSON PNYN QFUD BBHH VWEP YEYD OHNL XKZD NWRH DUWU JUMW WVII",
			"WZXI VIUQ DRHY MNCY EFUA PNHO TKHK GDNP SAKN UAGH JZSM JBMH",
			"VTRE QEDG XHLZ WIFU SKDQ VELN MIMI THBH DBWV HDFY HJOQ IHOR",
			"TDJD BWXE MEAY XGYQ XOHF DMYU XXNO JAZR SGHP LWML RECW WUTL",
			"RTTV LBHY OORG LGOW UXNX HMHY FAAC QEKT HSJW",
		}),
		Plaintext: LinesToText([]string{
			"KRKR ALLE XXFO LGEN DESI STSO FORT BEKA NNTZ UGEB ENXX ICHH",
			"ABEF OLGE LNBE BEFE HLER HALT ENXX JANS TERL EDES BISH ERIG",
			"XNRE ICHS MARS CHAL LSJG OERI NGJS ETZT DERF UEHR ERSI EYHV",
			"RRGR ZSSA DMIR ALYA LSSE INEN NACH FOLG EREI NXSC HRIF TLSC",
			"HEVO LLMA CHTU NTER WEGS XABS OFOR TSOL LENS IESA EMTL ICHE",
			"MASS NAHM ENVE RFUE GENY DIES ICHA USDE RGEG ENWA ERTI GENL",
			"AGEE RGEB ENXG EZXR EICH SLEI TEIK KTUL PEKK JBOR MANN JXXO",
			"BXDX MMMD URNH FKST XKOM XADM XUUU BOOI EXKP",
		}),
	}

	e, err := New(c)
	if err != nil {
		t.Fatal(err)
	}

	plaintext := e.FeedString(tp.Ciphertext)
	if plaintext != tp.Plaintext {
		t.Fatalf("wrong decrypt 'Dönitz message'")
	}
}
