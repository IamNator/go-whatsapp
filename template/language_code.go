package template

type LanguageCode string

func (l LanguageCode) String() string {
	return string(l)
}

// Short codes for supported languages
const (
	// Afrikaans
	AF LanguageCode = "af"

	// Albanian
	SQ LanguageCode = "sq"

	// Arabic
	AR LanguageCode = "ar"

	// Azerbaijani
	AZ LanguageCode = "az"

	// Bengali
	BN LanguageCode = "bn"

	// Bulgarian
	BG LanguageCode = "bg"

	// Catalan
	CA LanguageCode = "ca"

	// Chinese (China)
	ZH_CN LanguageCode = "zh_CN"

	// Chinese (Hong Kong)
	ZH_HK LanguageCode = "zh_HK"

	// Chinese (Taiwan)
	ZH_TW LanguageCode = "zh_TW"

	// Croatian
	HR LanguageCode = "hr"

	// Czech
	CS LanguageCode = "cs"

	// Danish
	DA LanguageCode = "da"

	// Dutch
	NL LanguageCode = "nl"

	// English
	EN LanguageCode = "en"

	// English (UK)
	EN_GB LanguageCode = "en_GB"

	// English (US)
	EN_US LanguageCode = "en_US"

	// Estonian
	ET LanguageCode = "et"

	// Filipino
	FIL LanguageCode = "fil"

	// Finnish
	FI LanguageCode = "fi"

	// French
	FR LanguageCode = "fr"

	// German
	DE LanguageCode = "de"

	// Greek
	EL LanguageCode = "el"

	// Gujarati
	GU LanguageCode = "gu"

	// Hausa
	HA LanguageCode = "ha"

	// Hebrew
	HE LanguageCode = "he"

	// Hindi
	HI LanguageCode = "hi"

	// Hungarian
	HU LanguageCode = "hu"

	// Indonesian
	ID LanguageCode = "id"

	// Irish
	GA LanguageCode = "ga"

	// Italian
	IT LanguageCode = "it"

	// Japanese
	JA LanguageCode = "ja"

	// Kannada
	KN LanguageCode = "kn"

	// Kazakh
	KK LanguageCode = "kk"

	// Korean
	KO LanguageCode = "ko"

	// Lao
	LO LanguageCode = "lo"

	// Latvian
	LV LanguageCode = "lv"

	// Lithuanian
	LT LanguageCode = "lt"

	// Macedonian
	MK LanguageCode = "mk"

	// Malay
	MS LanguageCode = "ms"

	// Malayalam
	ML LanguageCode = "ml"

	// Marathi
	MR LanguageCode = "mr"

	// Norwegian
	NB LanguageCode = "nb"

	// Persian
	FA LanguageCode = "fa"

	// Polish
	PL LanguageCode = "pl"

	// Portuguese (Brazil)
	PT_BR LanguageCode = "pt_BR"

	// Portuguese (Portugal)
	PT_PT LanguageCode = "pt_PT"

	// Punjabi
	PA LanguageCode = "pa"

	// Romanian
	RO LanguageCode = "ro"

	// Russian
	RU LanguageCode = "ru"

	// Serbian
	SR LanguageCode = "sr"

	// Slovak
	SK LanguageCode = "sk"

	// Slovenian
	SL LanguageCode = "sl"

	// Spanish
	ES LanguageCode = "es"

	// Spanish (Argentina)
	ES_AR LanguageCode = "es_AR"

	// Spanish (Spain)
	ES_ES LanguageCode = "es_ES"

	// Spanish (Mexico)
	ES_MX LanguageCode = "es_MX"

	// Swahili
	SW LanguageCode = "sw"

	// Swedish
	SV LanguageCode = "sv"

	// Tamil
	TA LanguageCode = "ta"

	// Telugu
	TE LanguageCode = "te"

	// Thai
	TH LanguageCode = "th"

	// Turkish
	TR LanguageCode = "tr"

	// Ukrainian
	UK LanguageCode = "uk"

	// Urdu
	UR LanguageCode = "ur"

	// Uzbek
	UZ LanguageCode = "uz"

	// Vietnamese
	VI LanguageCode = "vi"

	// Zulu
	ZU LanguageCode = "zu"
)

// Long codes for supported languages

const (
	// Afrikaans
	Afrikaans LanguageCode = "af"

	// Albanian
	Albanian LanguageCode = "sq"

	// Arabic
	Arabic LanguageCode = "ar"

	// Azerbaijani
	Azerbaijani LanguageCode = "az"

	// Bengali
	Bengali LanguageCode = "bn"

	// Bulgarian
	Bulgarian LanguageCode = "bg"

	// Catalan
	Catalan LanguageCode = "ca"

	// Chinese (China)
	ChineseChina LanguageCode = "zh_CN"

	// Chinese (Hong Kong)
	ChineseHongKong LanguageCode = "zh_HK"

	// Chinese (Taiwan)
	ChineseTaiwan LanguageCode = "zh_TW"

	// Croatian
	Croatian LanguageCode = "hr"

	// Czech
	Czech LanguageCode = "cs"

	// Danish
	Danish LanguageCode = "da"

	// Dutch
	Dutch LanguageCode = "nl"

	// English
	English LanguageCode = "en"

	// English (UK)
	EnglishUK LanguageCode = "en_GB"

	// English (US)
	EnglishUS LanguageCode = "en_US"

	// Estonian
	Estonian LanguageCode = "et"

	// Filipino
	Filipino LanguageCode = "fil"

	// Finnish
	Finnish LanguageCode = "fi"

	// French
	French LanguageCode = "fr"

	// German
	German LanguageCode = "de"

	// Greek
	Greek LanguageCode = "el"

	// Gujarati
	Gujarati LanguageCode = "gu"

	// Hausa
	Hausa LanguageCode = "ha"

	// Hebrew
	Hebrew LanguageCode = "he"

	// Hindi
	Hindi LanguageCode = "hi"

	// Hungarian
	Hungarian LanguageCode = "hu"

	// Indonesian
	Indonesian LanguageCode = "id"

	// Irish
	Irish LanguageCode = "ga"

	// Italian
	Italian LanguageCode = "it"

	// Japanese
	Japanese LanguageCode = "ja"

	// Kannada
	Kannada LanguageCode = "kn"

	// Kazakh
	Kazakh LanguageCode = "kk"

	// Korean
	Korean LanguageCode = "ko"

	// Lao
	Lao LanguageCode = "lo"

	// Latvian
	Latvian LanguageCode = "lv"

	// Lithuanian
	Lithuanian LanguageCode = "lt"

	// Macedonian
	Macedonian LanguageCode = "mk"

	// Malay
	Malay LanguageCode = "ms"

	// Malayalam
	Malayalam LanguageCode = "ml"

	// Marathi
	Marathi LanguageCode = "mr"

	// Norwegian
	Norwegian LanguageCode = "nb"

	// Persian
	Persian LanguageCode = "fa"

	// Polish
	Polish LanguageCode = "pl"

	// Portuguese (Brazil)
	PortugueseBrazil LanguageCode = "pt_BR"

	// Portuguese (Portugal)
	PortuguesePortugal LanguageCode = "pt_PT"

	// Punjabi
	Punjabi LanguageCode = "pa"

	// Romanian
	Romanian LanguageCode = "ro"

	// Russian
	Russian LanguageCode = "ru"

	// Serbian
	Serbian LanguageCode = "sr"

	// Slovak
	Slovak LanguageCode = "sk"

	// Slovenian
	Slovenian LanguageCode = "sl"

	// Spanish
	Spanish LanguageCode = "es"

	// Spanish (Argentina)
	SpanishArgentina LanguageCode = "es_AR"

	// Spanish (Spain)
	SpanishSpain LanguageCode = "es_ES"

	// Spanish (Mexico)
	SpanishMexico LanguageCode = "es_MX"

	// Swahili
	Swahili LanguageCode = "sw"

	// Swedish
	Swedish LanguageCode = "sv"

	// Tamil
	Tamil LanguageCode = "ta"

	// Telugu
	Telugu LanguageCode = "te"

	// Thai
	Thai LanguageCode = "th"

	// Turkish
	Turkish LanguageCode = "tr"

	// Ukrainian
	Ukrainian LanguageCode = "uk"

	// Urdu
	Urdu LanguageCode = "ur"

	// Uzbek
	Uzbek LanguageCode = "uz"

	// Vietnamese
	Vietnamese LanguageCode = "vi"

	// Zulu
	Zulu LanguageCode = "zu"
)
