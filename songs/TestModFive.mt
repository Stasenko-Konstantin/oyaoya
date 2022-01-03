
(Test module.)
Module "Test Mod Five"
	Channels 8
	Sequence 0,0,1,1,2,2,3,3,4,4,4,4,5,6,4,4,4,4,6,7
	Instrument 1 Name "Str1(Micromod)" Volume 64
		Waveform Sawtooth Octave -1 Chorus 512
	Instrument 5 Name "Bass(Micromod)" Volume 64
		Waveform Square   Octave -4 Chorus 255
	Instrument 7 Name "BassDrum(HammerHead)" Volume 64
		WaveFile "samples/BassDrum.wav" Gain 192 Pitch 46
	Instrument 8 Name "HiHat(HammerHead)" Volume 64
		WaveFile "samples/HiHat.wav" Gain 128 Pitch 46
	Instrument 13 Name "=:)===============v2="
	Instrument 14 Name "2015 mumart@gmail.com"
	Instrument 15 Name "==Micromod Compiler!="
	Macro 32
		Note "C-2-1C-0"
		Note "-----A1-" Repeat Begin
		Note "-----A3F" Repeat 14
		Note "-----AF8" Repeat 32
	Macro 33 Scale C-D-EF-G-A-B Root C-2
		Note "C-2-1840"
		Note "E-2-1---"
		Note "G-2-1---"
		Note "A-3-1---"
		Repeat 8
	Macro 34 Scale C-D-EF-G-A-B Root C-2
		Note "C-2-1-47"
		Note "E-2-1-37"
		Note "G-2-1---"
		Note "A-3-1-37"
		Repeat 8
	Macro 40 Root B-4
		Note "B-4-7---"
		Note "--------"
		Note "B-4-8---"
		Note "--------"
		Note "B-4-7---"
		Note "--------"
		Note "B-4-8---"
		Note "--------"
		Note "B-4-7---"
		Note "--------"
		Note "B-4-8---"
		Note "--------"
		Note "B-4-7---"
		Note "--------"
		Note "B-4-8---"
		Note "B-4-8---"
		Repeat 2
	Pattern 0
		Row "00 F-332--- A-432--- C-432--- -------- -------- -------- -----840 -----840"
		Row "16 -------- -------- -------- G-332--- C-432--- E-432---"
		Row "32 F-332--- C-432--- A-432--- -------- -------- --------"
		Row "48 -------- -------- -------- G-332--- B-432--- D-432---"
	Pattern 1
		Row "00 F-332--- A-432--- C-432--- -------- -------- -------- C-4-5---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- G-332--- C-432--- E-432--- --------"
		Row "32 F-332--- C-432--- A-432--- -------- -------- -------- G-3-5---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 -------- -------- -------- G-332--- B-432--- D-432--- --------"
	Pattern 2
		Row "00 F-332--- A-432--- C-432--- -------- -------- -------- C-4-5--- B-440---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- G-332--- C-432--- E-432--- --------"
		Row "32 F-332--- C-432--- A-432--- -------- -------- -------- G-3-5--- B-440---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 -------- -------- -------- G-332--- B-432--- D-432--- --------"
	Pattern 3
		Row "00 F-332--- A-433--- C-432--- -------- -------- -------- C-4-5--- B-440---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- G-332--- C-432--- E-432--- --------"
		Row "32 F-332--- C-433--- A-432--- -------- -------- -------- G-3-5--- B-440---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 -------- -------- -------- G-332--- B-432--- D-432--- --------"
	Pattern 4
		Row "00 F-332--- A-434--- C-432--- -------- -------- -------- C-4-5--- B-440---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- G-432--- C-432--- E-432--- --------"
		Row "32 F-332--- C-434--- A-432--- -------- -------- -------- G-3-5--- B-440---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 -------- -------- -------- G-332--- B-432--- D-432--- --------"
	Pattern 5
		Row "00 -------- A-433--- -------- -------- -------- -------- C-4-5---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- -------- E-432--- -------- --------"
		Row "32 -------- A-434--- -------- -------- -------- -------- G-3-5---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 C-231--- -------- C-231--- C-231--- C-231--- C-231--- --------"
	Pattern 6
		Row "00 -------- A-433--- -------- -------- -------- -------- C-4-5---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- -------- E-432--- -------- --------"
		Row "32 -------- A-434--- -------- -------- -------- -------- G-3-5---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 -------- -------- -------- -------- D-432--- G-332---"
	Pattern 7
		Row "00 -------- A-433--- -------- -------- -------- -------- C-4-5---"
		Row "04 -------- -------- -------- -------- -------- -------- F-3-5---"
		Row "16 -------- -------- -------- -------- G-432--- -------- --------"
		Row "32 -------- A-434--- -------- -------- -------- -------- G-3-5---"
		Row "36 -------- -------- -------- -------- -------- -------- E-3-5---"
		Row "48 -------- -------- -------- -------- B-432--- G-332---"
		Row "63 -----B0E"
(End.)
