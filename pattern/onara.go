package pattern

// OhimesamaEmotion ... おじさんの感情列挙体
type OhimesamaEmotion int

const (
	// GREETING ... 挨拶
	GREETING OhimesamaEmotion = iota
	// COMMENT ... コメント
	COMMENT
	// JKGREETING ... JKの挨拶
	JKGREETING
	// JKCOMMENT ... JKのコメント
	JKCOMMENT
	// MENHERAGREETING ... メンヘラの挨拶
	MENHERAGREETING
	// MENHERACOMMENT ... メンヘラのコメント
	MENHERACOMMENT
)

// Onara ... Ohimesama NArikiri Randomized Algorithm: お姫様なりきり乱択アルゴリズム
// お姫様の感情表現の順番を表す。
// 近年の研究により (README.md 参考文献[1]) お姫様になりきるための効果的なアルゴリズムが提唱されている。
var Onara = [][]OhimesamaEmotion{
	// GC パターン
	[]OhimesamaEmotion{GREETING, COMMENT},
	// GC パターン2
	[]OhimesamaEmotion{GREETING, COMMENT, COMMENT},
	// JK パターン
	[]OhimesamaEmotion{JKGREETING, JKCOMMENT},
	// JK パターン2
	[]OhimesamaEmotion{JKGREETING, JKCOMMENT, JKCOMMENT},
	// MENHERA パターン
	[]OhimesamaEmotion{MENHERAGREETING, MENHERACOMMENT},
}

// OnaraMessages .. メッセージのテンプレート
// メッセージ中の{TARGET_NAME} などのタグについては tags.go 参照
var OnaraMessages = [][]string{
	GREETING: []string{
		"{TARGET_NAME}{EMOJI_POS}",
		"{TARGET_NAME}お疲れ様{EMOJI_POS}{EMOJI_POS}",
		"おはよう{TARGET_NAME}{EMOJI_POS}",
		"{TARGET_NAME}ヤッホー！{EMOJI_POS}なにしてるん{EMOJI_ASK}",
		"{TARGET_NAME}、今日もお仕事かな{EMOJI_POS}",
		"おはよー！{EMOJI_POS}",
		"{TARGET_NAME}久しぶり{EMOJI_POS}",
		"{FIRST_PERSON}はまだ起きてますよ〜{EMOJI_POS}",
	},
	COMMENT: []string{
		"顔面可愛いのに死ぬほどテンション高くておもしろいしめちゃめちゃらぶじゃん〜🥰🥰ずっ友や{TARGET_NAME}しか勝たん😢いい一年にしてね〜‼️",
		"私、看護学生{EMOJI_POS}医療系の学科ばかりの大学で毎日勉強{EMOJI_POS}でもある日「看護師は皆医者狙うためになるんだろ」とかいうクソ男が現れてもう大変{EMOJI_NEG}",
		"{FIRST_PERSON}は、近所に新しくできた{SPOT}に行ってきたよ。まぁまぁだったかな{EMOJI_POS}",
		"そういえば、昨日は例の{SPOT}に行ってきたよ。結構いい雰囲気だったから、オススメだよ{EMOJI_POS}",
		"{FIRST_PERSON}は今日から{LOCATION}へ〜{EMOJI_POS}",
		"{TARGET_NAME}の髪色めちゃくちゃすき{EMOJI_POS}{EMOJI_POS}わたしもしたい{EMOJI_POS}",
		"{TARGET_NAME}{EMOJI_POS}元気、ないのかなぁ{EMOJI_NEG}大丈夫{EMOJI_ASK}",
		"僕は、すごく心配だよ{EMOJI_NEG}そんなときは、美味しいもの食べて、元気出さなきゃだね{EMOJI_POS}",
		"今日も大変だったんだね{EMOJI_NEG}",
		"{FIRST_PERSON}は{TARGET_NAME}の味方だからね{EMOJI_POS}",
		"今日はよく休んでね{EMOJI_NEUT}",
		"くれぐれも体調に気をつけて{EMOJI_NEUT}",
		"{FIRST_PERSON}は{TARGET_NAME}一筋だよ{EMOJI_NEUT}",
		"よく頑張ったね{EMOJI_POS}えらいえらい{EMOJI_POS}",
		"風邪ひかないようにね{EMOJI_POS}",
		"寒いけど、頑張ってね{EMOJI_NEUT}",
		"ゆっくり、身体休めてね{EMOJI_POS}オヤスミナサイ{EMOJI_NEUT}",
		"今日も頑張ってね{EMOJI_POS}",
		"{TARGET_NAME}にとって素敵な1日になりますように{EMOJI_POS}",
		"今日は楽しい時間をありがとうね{EMOJI_POS}すごく、楽しかったよ{EMOJI_POS}",
		"早く会いたいな{EMOJI_POS}",
		"ありがとう{EMOJI_POS}",
		"今日はどんな一日だった{EMOJI_ASK}",
	},
	JKGREETING: []string{
		"{TARGET_NAME}{EMOJI_POS}",
		"おはよー！{EMOJI_POS}",
		"{TARGET_NAME}ヤッホー！{EMOJI_POS}なにしてるん{EMOJI_ASK}",
		"{TARGET_NAME}{EMOJI_POS}",
		"{TARGET_NAME}誕生日おめでとう🐶💖",
		"*⑅୨୧---------お友達探し----------୨୧⑅* ♡{TARGET_NAME}♡",
		"{EMOJI_POS}{TARGET_NAME}しか勝たん卍{EMOJI_POS}",
	},
	JKCOMMENT: []string{
		"ほんとこれおもろしすぎてしぬｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ",
		"ぴえん😢ぴえん😢ぴえん😢ぴえん😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ〜〜〜〜〜〜〜〜ん😢むりみ😭",
		"ふぁぼぼぼぼってちょぉーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーーだい{EMOJI_POS}！！！",
		"ほんまラブ会いたいんやけど？？？{EMOJI_POS}",
		"{SPOT}好きかな{EMOJI_ASK}",
		"もっともっと仲良くなる{EMOJI_POS}",
		"{TARGET_NAME}{EMOJI_POS}世界一ツインテールが似合うかわいあかわいい天使です{EMOJI_POS}",
		"よーーーくかんがえよーーー{TARGET_NAME}はあたおかよーーーー",
		"*⑅︎୨୧{TARGET_NAME}を自分なりに推してます୨୧⑅︎*",
		"拒否等ないのでどなたでも反応ください{EMOJI_POS}{EMOJI_NEUT}",
		"えっなにそれやばい沸いた好きすぎるんだけど？！？！？！",
		"{FIRST_PERSON}明日も学校だけどなかなか寝れないよ〜{EMOJI_NEG}早く{TARGET_NAME}会いたい{EMOJI_NEG}",
		"{DAY_OF_WEEK}曜日は学校{EMOJI_ASK}",
	},
	MENHERAGREETING: []string{
		"いっけなーい{EMOJI_NEG}殺意殺意{EMOJI_NEG}",
		"{TARGET_NAME}。",
		"{TARGET_NAME}と別れた。 辛ぃ。 どぅせぅちは遊びたったってことでしょ。",
		"「 任天堂がうどん店に 」 逆から読むと。。「 任天堂がうどん店に 」 ぃみゎかんなぃ。。。。",
	},
	MENHERACOMMENT: []string{
		"今日はもう寝ちゃったのかな{EMOJI_NEUT}",
		"ねられない{EMOJI_POS}どうしてくれるん{EMOJI_NEG}",
		"もぅﾏﾁﾞ無理。ﾏﾘｶしょ･･･ぇ。。ﾏﾘｶ発売停止。。。ｽﾞﾙﾙﾙﾙﾙﾙﾙｗｗｗｗｽﾞﾙｯｽﾞﾙﾙﾙﾙﾙｗｗｗｗﾋﾞﾁｬｯｗｗｗ",
		"どんどん分裂してぃまゎ32体になってる。影分身ぢゃなぃからぅちに勝ち目ゎなぃんだって。まぢ包囲されてる。っょぃ。勝てなぃ。",
		"ぴえん😢ぴえん😢ぴえん😢ぴえん😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ😢ぴえ〜〜〜〜〜〜〜〜ん😢むりみ😭",
	},
}
