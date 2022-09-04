package kalyna

import (
	"fmt"
	"testing"
)

type roundKeySample struct {
	id        string
	config    Config
	key       []uint64
	roundKeys [][]uint64
}

var roundKeySamples = []roundKeySample{
	{
		id: "B.2.1",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize128,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
		},
		roundKeys: [][]uint64{
			{
				0xe6b13a9b6b5e5016, 0xf4a082e0dc775b86,
			},
			{
				0xa082e0dc775b86e6, 0xb13a9b6b5e5016f4,
			},
			{
				0x768449ae6e87707e, 0x42ec937c0aa0aa8a,
			},
			{
				0xec937c0aa0aa8a76, 0x8449ae6e87707e42,
			},
			{
				0xf540911ec5d4ce45, 0xfed90b0f8276723e,
			},
			{
				0xd90b0f8276723ef5, 0x40911ec5d4ce45fe,
			},
			{
				0x62c4007922ee778c, 0xb1c4600532665f51,
			},
			{
				0xc4600532665f5162, 0xc4007922ee778cb1,
			},
			{
				0xb8b0d25ce272980a, 0xd86da686209a87aa,
			},
			{
				0x6da686209a87aab8, 0xb0d25ce272980ad8,
			},
			{
				0x18c4db94a8b12657, 0x6148d7e8d5f30bf6,
			},
		},
	},
	{
		id: "B.2.2",
		config: Config{
			BlockSize: BlockSize128,
			KeySize:   KeySize256,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
		},
		roundKeys: [][]uint64{
			{
				0xde127e3feb16c857, 0x1abeb5e6566b2ced,
			},
			{
				0xbeb5e6566b2cedde, 0x127e3feb16c8571a,
			},
			{
				0x80cd9a887d9a06d8, 0x6faecc6c458431cd,
			},
			{
				0xaecc6c458431cd80, 0xcd9a887d9a06d86f,
			},
			{
				0x1a41133597cc61c3, 0xfef342672b4d3282,
			},
			{
				0xf342672b4d32821a, 0x41133597cc61c3fe,
			},
			{
				0xa480cf658c691083, 0x560fb8beaa6fef09,
			},
			{
				0x0fb8beaa6fef09a4, 0x80cf658c69108356,
			},
			{
				0x7b1a4681c3c4d5c6, 0x1015904218694d03,
			},
			{
				0x15904218694d037b, 0x1a4681c3c4d5c610,
			},
			{
				0xf9bdc84621f8d084, 0x7e38494d7b70b3b2,
			},
			{
				0x38494d7b70b3b2f9, 0xbdc84621f8d0847e,
			},
			{
				0x2bd4d1a028dbfa43, 0xb3464579f92ff9bf,
			},
			{
				0x464579f92ff9bf2b, 0xd4d1a028dbfa43b3,
			},
			{
				0x24ed2c7ea8e81ec3, 0x925bb2fd35a4215a,
			},
		},
	},
	{
		id: "B.2.3",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize256,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
		},
		roundKeys: [][]uint64{
			{
				0x355bd5df4726daf7, 0xa1cb0fe30852082f,
				0x80d70dc8dcc9b369, 0x362e946cc12c071f,
			},
			{
				0xc9b369a1cb0fe308, 0x2c071f80d70dc8dc,
				0x26daf7362e946cc1, 0x52082f355bd5df47,
			},
			{
				0xa5ee74fd58111fdf, 0xdb389023c93155c1,
				0xcb1f0300919a8326, 0xc2734ff135fd7cb7,
			},
			{
				0x9a8326db389023c9, 0xfd7cb7cb1f030091,
				0x111fdfc2734ff135, 0x3155c1a5ee74fd58,
			},
			{
				0x0a6c655c8ae52aca, 0xfca9fb28a6f00bce,
				0x5df711fc10e77631, 0x12fd225cf90196eb,
			},
			{
				0xe77631fca9fb28a6, 0x0196eb5df711fc10,
				0xe52aca12fd225cf9, 0xf00bce0a6c655c8a,
			},
			{
				0x4e4010c5ecfc5751, 0x80cbf2f26ce7f544,
				0x7dbe10e353ee9bc2, 0x605117eb9aa416f8,
			},
			{
				0xee9bc280cbf2f26c, 0xa416f87dbe10e353,
				0xfc5751605117eb9a, 0xe7f5444e4010c5ec,
			},
			{
				0x2a86519b88151c5e, 0xd1b8815787d21da1,
				0xb06946aa70a2d00b, 0xeb673447b2497a6b,
			},
			{
				0xa2d00bd1b8815787, 0x497a6bb06946aa70,
				0x151c5eeb673447b2, 0xd21da12a86519b88,
			},
			{
				0xaae6a91f5e265237, 0x7bb6c83199a92a08,
				0xef4f6f94e7df6408, 0x02db12927cad5b7c,
			},
			{
				0xdf64087bb6c83199, 0xad5b7cef4f6f94e7,
				0x26523702db12927c, 0xa92a08aae6a91f5e,
			},
			{
				0xbe38887fd643a738, 0xc5e734176c7ed174,
				0x06b60644987dc82b, 0x3b8fbaf9171ca779,
			},
			{
				0x7dc82bc5e734176c, 0x1ca77906b6064498,
				0x43a7383b8fbaf917, 0x7ed174be38887fd6,
			},
			{
				0x49cac135a7b169fc, 0xacfcd688bbeb5018,
				0x319c105c1616765d, 0x02ea25b8c54431f1,
			},
		},
	},
	{
		id: "B.2.4",
		config: Config{
			BlockSize: BlockSize256,
			KeySize:   KeySize512,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
			0x2726252423222120, 0x2f2e2d2c2b2a2928,
			0x3736353433323130, 0x3f3e3d3c3b3a3938,
		},
		roundKeys: [][]uint64{
			{
				0xa8dd49ce3897bdf7, 0x21e81e8079bd9a0b,
				0x569f5c4742fe6088, 0xc489c9b433f4d85c,
			},
			{
				0xfe608821e81e8079, 0xf4d85c569f5c4742,
				0x97bdf7c489c9b433, 0xbd9a0ba8dd49ce38,
			},
			{
				0x6410eee5a3800b40, 0x43c0a845647e0302,
				0xb09442fb83bc25c0, 0xd2ca0bf2a19233a0,
			},
			{
				0xbc25c043c0a84564, 0x9233a0b09442fb83,
				0x800b40d2ca0bf2a1, 0x7e03026410eee5a3,
			},
			{
				0x783ee6e165d811d0, 0xe631e1f6e9dc3c35,
				0x871401a350af7a6f, 0x04d94c00a452cd98,
			},
			{
				0xaf7a6fe631e1f6e9, 0x52cd98871401a350,
				0xd811d004d94c00a4, 0xdc3c35783ee6e165,
			},
			{
				0x507cdcaef58799b2, 0xd00dd7b4922a8749,
				0x304677d36652cb6e, 0x451cb2bb1bf230f3,
			},
			{
				0x52cb6ed00dd7b492, 0xf230f3304677d366,
				0x8799b2451cb2bb1b, 0x2a8749507cdcaef5,
			},
			{
				0xdf3808ce65c55b53, 0xc71cd1d7b46afc30,
				0xac96c7a3b10ef01d, 0x8d2380cec6e86914,
			},
			{
				0x0ef01dc71cd1d7b4, 0xe86914ac96c7a3b1,
				0xc55b538d2380cec6, 0x6afc30df3808ce65,
			},
			{
				0x9c7db42e55585457, 0x7f075ef11af04662,
				0x5f36c85bc5d897cd, 0x987894b837fe9837,
			},
			{
				0xd897cd7f075ef11a, 0xfe98375f36c85bc5,
				0x585457987894b837, 0xf046629c7db42e55,
			},
			{
				0xdd6fa5af55a610ae, 0xc9db7b23adde9f36,
				0x6277a7149db6cfb9, 0x752602b29a967447,
			},
			{
				0xb6cfb9c9db7b23ad, 0x9674476277a7149d,
				0xa610ae752602b29a, 0xde9f36dd6fa5af55,
			},
			{
				0xc541f8cadc63a73b, 0x957c4457798a139b,
				0x4aa0dca9656102cb, 0xf5248b87bb92705d,
			},
			{
				0x6102cb957c445779, 0x92705d4aa0dca965,
				0x63a73bf5248b87bb, 0x8a139bc541f8cadc,
			},
			{
				0x3e473f211a45be19, 0xde5c2d9613875d9d,
				0xc9377f3e3c7b36e7, 0x790e79f307a3ab6e,
			},
			{
				0x7b36e7de5c2d9613, 0xa3ab6ec9377f3e3c,
				0x45be19790e79f307, 0x875d9d3e473f211a,
			},
			{
				0xe1451a023f12ca5b, 0x72345e2d09126115,
				0x9b918979a5ebede9, 0x23ca59eeff86acde,
			},
		},
	},
	{
		id: "B.2.5",
		config: Config{
			BlockSize: BlockSize512,
			KeySize:   KeySize512,
		},
		key: []uint64{
			0x0706050403020100, 0x0f0e0d0c0b0a0908,
			0x1716151413121110, 0x1f1e1d1c1b1a1918,
			0x2726252423222120, 0x2f2e2d2c2b2a2928,
			0x3736353433323130, 0x3f3e3d3c3b3a3938,
		},
		roundKeys: [][]uint64{
			{
				0x4cd6c8f2fcc04719, 0x41e2a82d72fb56b2,
				0x4f06306102bfea88, 0xc954a84cbdf2e003,
				0x00cffe4564be48a8, 0xf1434d1c0cde0a51,
				0xdf01df63b1533d07, 0xe562f252957e3101,
			},
			{
				0xf2e0034f06306102, 0xbe48a8c954a84cbd,
				0xde0a5100cffe4564, 0x533d07f1434d1c0c,
				0x7e3101df01df63b1, 0xc04719e562f25295,
				0xfb56b24cd6c8f2fc, 0xbfea8841e2a82d72,
			},
			{
				0x74932a31faf20824, 0x577cbc063c7700a3,
				0x79d8e1ec36a169e3, 0xb8f54903a3a2dfc2,
				0x5c03ab0117d25c23, 0x33e28d0ff31bbf23,
				0x2701f84dc93af8e0, 0x560dda078a4bbe03,
			},
			{
				0xa2dfc279d8e1ec36, 0xd25c23b8f54903a3,
				0x1bbf235c03ab0117, 0x3af8e033e28d0ff3,
				0x4bbe032701f84dc9, 0xf20824560dda078a,
				0x7700a374932a31fa, 0xa169e3577cbc063c,
			},
			{
				0xfd8b7ca96739dfb8, 0xb4ee6851ecec4d5c,
				0xd67ad9802d56ee0b, 0x5d4d08cd9c659809,
				0xfa7abf00f04134c3, 0xde676df931e4ebee,
				0xaab0f8ff5b5ae0a0, 0x23ea917a85c9389a,
			},
			{
				0x659809d67ad9802d, 0x4134c35d4d08cd9c,
				0xe4ebeefa7abf00f0, 0x5ae0a0de676df931,
				0xc9389aaab0f8ff5b, 0x39dfb823ea917a85,
				0xec4d5cfd8b7ca967, 0x56ee0bb4ee6851ec,
			},
			{
				0xc35aa0aa7726b67a, 0x8083a01c8694f95b,
				0xfc514382fb518605, 0xeb829f22ddb584d6,
				0xee980666fdcdcf3a, 0x68687d092dc622b9,
				0xeae660692c517503, 0xa3db59c1a7534a97,
			},
			{
				0xb584d6fc514382fb, 0xcdcf3aeb829f22dd,
				0xc622b9ee980666fd, 0x51750368687d092d,
				0x534a97eae660692c, 0x26b67aa3db59c1a7,
				0x94f95bc35aa0aa77, 0x5186058083a01c86,
			},
			{
				0x103e37223289410b, 0xa6aaec2ae52c8553,
				0x1f77d287bbd46ad5, 0x0776bab4b29efeb7,
				0x02f8e93197dbf039, 0xe1600093bde121e8,
				0xc74340d83a2538fc, 0x826285a59a6afe8f,
			},
			{
				0x9efeb71f77d287bb, 0xdbf0390776bab4b2,
				0xe121e802f8e93197, 0x2538fce1600093bd,
				0x6afe8fc74340d83a, 0x89410b826285a59a,
				0x2c8553103e372232, 0xd46ad5a6aaec2ae5,
			},
			{
				0xbb81978c95f753b2, 0xfaa4b29431927fc9,
				0x9ce3d2e3410195f3, 0xfffd0a8b6c84d727,
				0x411545e3a1ede010, 0xf109c4ccd7e54458,
				0x0841d72804b308af, 0x52ee1289743913ae,
			},
			{
				0x84d7279ce3d2e341, 0xede010fffd0a8b6c,
				0xe54458411545e3a1, 0xb308aff109c4ccd7,
				0x3913ae0841d72804, 0xf753b252ee128974,
				0x927fc9bb81978c95, 0x0195f3faa4b29431,
			},
			{
				0xc63b22ecb9873b9c, 0xdaa76f3b15a4f15b,
				0x5f37eafc2ab7b962, 0x4f2957486c65cd34,
				0x270fb88847c00106, 0x10b781fbe17e9dc1,
				0x107f4466bd617fcf, 0x0421e3630eda4ce3,
			},
			{
				0x65cd345f37eafc2a, 0xc001064f2957486c,
				0x7e9dc1270fb88847, 0x617fcf10b781fbe1,
				0xda4ce3107f4466bd, 0x873b9c0421e3630e,
				0xa4f15bc63b22ecb9, 0xb7b962daa76f3b15,
			},
			{
				0x94e6b336e687e039, 0x4ecfddb4ac4e749c,
				0x49d5312765953c60, 0x9550d9f100520ce7,
				0xf9565b674e404eb0, 0x382c51f74ac4be3c,
				0x308cbcfbde3d1a8e, 0xf42bb1fddc3bdefa,
			},
			{
				0x520ce749d5312765, 0x404eb09550d9f100,
				0xc4be3cf9565b674e, 0x3d1a8e382c51f74a,
				0x3bdefa308cbcfbde, 0x87e039f42bb1fddc,
				0x4e749c94e6b336e6, 0x953c604ecfddb4ac,
			},
			{
				0xc8fbe43f1b0dfce0, 0x96cb2b11b53dac7f,
				0xccc1853f593ef2d8, 0x054ddc23aad2d945,
				0x8b1f34c308eb6f6c, 0x43592b9bf8e27aaf,
				0xb81b28039f109e54, 0x2e01e4c101d21eb9,
			},
			{
				0xd2d945ccc1853f59, 0xeb6f6c054ddc23aa,
				0xe27aaf8b1f34c308, 0x109e5443592b9bf8,
				0xd21eb9b81b28039f, 0x0dfce02e01e4c101,
				0x3dac7fc8fbe43f1b, 0x3ef2d896cb2b11b5,
			},
			{
				0x91d0099e98d97a83, 0x57a083d8a45e2eab,
				0xdb3302c97514a86e, 0x5ae82297303752bd,
				0x4a67418b202851a3, 0xa8c451812f286790,
				0x7da3e7b3bcbd5a68, 0x246ba8adb6ba6218,
			},
		},
	},
}

func TestRoundKeySamples(t *testing.T) {
	for _, v := range roundKeySamples {
		err := testRoundKeySample(t, v)
		if err != nil {
			t.Fatalf("sample %q error: %s", v.id, err)
		} else {
			t.Logf("sample %q success", v.id)
		}
	}
}

func testRoundKeySample(t *testing.T, v roundKeySample) error {

	k, err := newKalynaContext(v.config.BlockSize, v.config.KeySize)
	if err != nil {
		return err
	}

	k.KeyExpand(v.key)
	if err != nil {
		return err
	}

	var (
		haveSize = len(k.roundKeys)
		wantSize = len(v.roundKeys)
	)

	if haveSize != wantSize {
		return fmt.Errorf("invalid roundKeys size: have %d, want %d",
			haveSize, wantSize)
	}

	for i := 0; i < haveSize; i++ {
		if !compareWords(k.roundKeys[i], v.roundKeys[i]) {
			return fmt.Errorf("failed roundKey %d", i)
		}
	}

	return nil
}
