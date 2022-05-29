package string

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_maxProduct(t *testing.T) {
	words := []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}
	assert.Equal(t, maxProduct(words), 16)

	words = []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}
	assert.Equal(t, maxProduct(words), 4)

	words = []string{"a", "aa", "aaa", "aaaa"}
	assert.Equal(t, maxProduct(words), 0)

}

func Test_noCommonLetter(t *testing.T) {
	totalMap = map[string]map[string]bool{}
	assert.True(t, noCommonLetter("abc", "fff"))
	assert.False(t, noCommonLetter("abcf", "fff"))
	assert.False(t, noCommonLetter("abc", "c"))
	assert.False(t, noCommonLetter("abc", "fefwefwa232323"))
}

func TestF(t *testing.T) {
	words := []string{
		"iebolahhjaneankkhioaaamnnfijijckhbokjjeelhfglbnhhf",
		"eglbhgcninjemdmljpf",
		"aijpcgmmdfidbbnnfgkigdejogkfabdgpfiljbhfcjnd",
		"cmidghlgidgpghalgknlhnlkjdaegbffjjihck",
		"mdjnmhhkkkdncilnnnjobmajieiffmdhahkaj",
		"gbpbkjknibddikom",
		"fffaofokncdegfabhpnmom",
		"lmjfgpfhhoohfncgakgmaecmkfbehddkmhmeigefbdamamgic",
		"dmeahegadlhpgnlchdcfgcmdchfigonlmk",
		"cgbiobhgbiimpbleljghdphmfadkiepfiddehhdhkkeiikc",
		"ljcmdlgfpiambhngjeppndgibphcemmcjckegofeogajfjo",
		"ngfhninlcmmpooejpcpklfpnheclmehhhf",
		"ncafhgpfi",
		"bffopidgdeopknmogdfkfhojomgmmdlalif",
		"bfkkeajieifmkalgoodjilonocddehbogmhlmikjfdoip",
		"filnboabacgkjkijemjejenfoeccakkpoelmbcibhflpdgagobd",
		"kiiogblmflmohfdhabmocaohjhjnlcmn",
		"kbjnbnikplm",
		"dbcpeca",
		"pejgcgjddlenagkbjpoemelapilnpckpmdcejognfebk",
		"doahlpbhpfbojbgnlldfchedicdldgbffclfeeaojc",
		"kmafbekjcgafcpbljhkakfpfgpmceikmabhmppgkmpdgdgjpomf",
		"gafeapbg",
		"lbhaafgkljaagjjhhkjloicjidemohmdi",
		"obmfhgfdjndgkgkgkedlmkdlpcggoiljl",
		"kninkifkpkonad",
		"lljhgekmbmpjn",
		"omehjdahfdpejgebhfcdldpojkbljhlddohcepjcgj",
		"lcabmfghibmb",
		"bifpiiakcnmjalkdip",
		"ahnnbgfcpgedghlaocpbk",
		"lajbpgjfmdkoifjhiij",
		"ijjagekhk",
		"ijfdmbe",
		"lejon",
		"bidaa",
		"ifjnlhogfmipmnccocdgdbebajg",
		"kodhafmeiakochkpjfhhcj",
		"nlcecadpbjhhiffhedlkgcmm",
		"glhlnigphkefoghapipdkjddpadpopofabgik",
		"cfbjmhfccalbijpgpilakmghh",
		"mmokmiiopnmnikka",
		"eoihgeomgjeppg",
		"fj",
		"mnkfnafblnpaahbfikajaekfpaacdjmkfjamnc",
		"dladihlljgafneeofipjlimamkimnekckgmmapgnkedm",
		"kflchckfgkhkacecokhhgiedogiimkdbpap",
		"bmnbfpbfagjfcmfmgchaipcnhacnpfgkeoln",
		"nfonehgfonhmjdoknk",
		"bbcceepmhchoglncimfmddfgmbodmpbojndipklppa",
		"feabhnlnhdloalcf",
		"ifkibdnbkfldenpgjkaakoennbjol",
		"nbmlha",
		"mgpnbfjfggahgokjf",
		"eafeagaojpebjjgkj",
		"jjpkjgieeojbglk",
		"eifmbajaphlaklkllbgfpfaoeaeemnnahdeijbnnbfp",
		"hjiihklgblfllljnkaakgcoinclnolahcgokbkp",
		"jie",
		"haicaadhdoldlggffimlkoeljiibcocpnpblbfnfpeipikehj",
		"icogdemcknoojpomkdnkchecikoehneoegeilfhl",
		"ieahkhmoapcipjoodcfnibahbaejlodnjgknngofmhmligfn",
		"ladbbfb",
		"lfddgdagfnodhehmaebflkhancbaiplbpkcgbmmi",
		"inanknpmhiadphhkodhgnbhakmfbddgkcc",
		"ojmi",
		"ppogobmjoadcddkdhe",
		"jbifcnmcgikmheieakolmeaplaca",
		"hkpidedccnelnnemmjbjijacdgpejfakghd",
		"kpkgbhpbmlnjopgdelmdkclgnloideop",
		"aeae",
		"jbinncimfbilmmm",
		"dbmajalmlcn",
		"cooollmddgajalkcne",
		"cgojkepipoec",
		"ocopi",
		"afdh",
		"nn",
		"kcecoinfpj",
		"kendcinbdadfoicjfadjlakmheojpdlppm",
		"bfeichjfmifofecgpiafdoa",
		"kdjekgkgclmfmkidadpogfa",
		"hhceabpgggmnoakpfafcabkgmnfmdc",
		"bkifmlo",
		"ninjiamaiemheeej",
		"admdndlnmokmckli",
		"pdopingcbjekjkibhapobjlekfpainkfimfjjidpeca",
		"caaak",
		"dgjocgidcffjegefiijcpdlonkajj",
		"ocecgchhhepeafnfohm",
		"ilonadpffiladaoapokkakifhhoolejplpppj",
		"deokcohmhcojmibnfhioadgglleeeiohbgooafcoo",
		"ohpcdielkiokgghpdjbleaofebepokahhppoc",
		"bldkfmpkbkie",
		"lbhcmjleejpaehombdlbannkokbkaaafmlmbafljbliedijhog",
		"jmcidohgdchalkanncpofejg",
		"pdhafbjlkldjp",
		"kbnojdmpeh",
		"cmoemi",
		"mpinlpalogdlofdd",
		"mcogdbdbnbpcpii",
		"bljigpmklejojgno",
		"glhga",
		"plboojpeednaiimdlif",
	}

	ti := time.Now()
	fmt.Println(maxProduct(words))
	fmt.Println(time.Since(ti))
}
