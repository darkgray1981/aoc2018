package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	fmt.Println("Starting...")

	t := time.Now()

	double, triple := 0, 0

	arr := strings.Split(INPUT, "\n")
	if len(arr[0]) == 0 {
		arr = arr[1:]
	}
	if len(arr[len(arr)-1]) == 0 {
		arr = arr[:len(arr)-1]
	}

	for _, s := range arr {
		if len(s) == 0 {
			panic("Gyaah")
		}

		var count [30]int

		for _, c := range s {
			count[c-'a'] += 1
		}

		doubleFound, tripleFound := 0, 0

		for _, n := range count {
			if n == 2 {
				doubleFound = 1
			} else if n == 3 {
				tripleFound = 1
			}
		}

		double += doubleFound
		triple += tripleFound
	}

	fmt.Println("Done:", double*triple, time.Since(t))

	t = time.Now()

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if differing(arr[i], arr[j]) == 1 {
				fmt.Println("Done 2: "+same(arr[i], arr[j]), time.Since(t))
				return
			}
		}
	}
}

func same(a, b string) string {

	var output string

	if len(a) != len(b) {
		return output
	}

	for i := range a {
		if a[i] == b[i] {
			output += string(a[i])
		}
	}

	return output
}

func differing(a, b string) int {
	if len(a) != len(b) {
		return -1
	}

	count := 0

	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}

	return count
}

const INPUT = `
ymdrcyapvwfloiuktanxzjsieb
ymdrwhgznwfloiuktanxzjsqeb
ymdrchguvwfloiuktanxmjsleb
pmdrchgmvwfdoiuktanxzjsqeb
ymdrfegpvwfloiukjanxzjsqeb
ymdrchgpvwfloiukmanazjsdeb
ymdsnhgpvwflciuktanxzjsqeb
lmdrbhrpvwfloiuktanxzjsqeb
ymdrwhgpvwfloiukeanxzjsjeb
ymdrchgpvpfloihktanszjsqeb
omdrchgpvwflokuktanazjsqeb
kmsrchgpvwfloiuktanxqjsqeb
ymdrchopvwzloiustanxzjsqeb
omdrchgpvwfloiuktawxtjsqeb
ymdrchgpvwfroiukhanozjsqeb
ymdrchgpvwfloikktanyzosqeb
ymdrchgpvwfioiuktaexzjsqea
ymdrcngpvwfloiuktanxzjsamb
ymdrchgpqwfaoiuktanxxjsqeb
ymdrcmgpvwflziuktakxzjsqeb
ymdrchguvwsloiuktanxzjsqen
ymdrchppowfloiuvtanxzjsqeb
ymdrcngpvwflyiukkanxzjsqeb
ymdrcbgpvwfloiukjanxzjspeb
ymdrchgpvwflopuktanxzosseb
ygdrchgpvwfloiukxanxcjsqeb
ymdrchgpvwfloauktanuzjsqei
ymerchgpvwfloiumtanxzjsqet
ymdlcegpvwfloiuktbnxzjsqeb
ymdrclgpvwfloiukyanxzjlqeb
ymdrchgpvwhmoiuktanxijsqeb
ymdrchgpwrfloiuktanxdjsqeb
ymdbcwgpvwfloiuktanxzusqeb
ymgrchgphwfloiuktanxzjspeb
imdrchgpvwflmiuktanxzjsqib
ymdrihgpvwfloiiktanxzjsteb
ywdrchgpvwfloibkvanxzjsqeb
ymdrchgpxwfloiuktanezjsqes
ymdrchgpiwfloiukxanxzhsqeb
ymdrchgpveflokuktdnxzjsqeb
kmdrchgpvwfloviktanxzjsqeb
ymdrchgpvgfmoiuytanxzjsqeb
ymyrchgpvzfluiuktanxzjsqeb
ymdrchguvwfloiuktanxpjsqlb
ymerchgpvwfloiukthnxsjsqeb
hmdrchgpvwglfiuktanxzjsqeb
ymdrchgpvwfdoiuktanxzjsqaf
yudrchgdvwfloiuktaexzjsqeb
ymdbchgxvwfloiuktanxzjsqem
ymdrchgpvwfloiumjanxzjsqpb
ymdrchgpqwfloiuqtanxrjsqeb
ymdqchhpvwfloiuktanxzzsqeb
ymdryhgpfwfloiuktanxzjsyeb
xmdrchgpvwfloioitanxzjsqeb
ykdrchgpvwfloiuktcnxzisqeb
ymdrcpgprwfloiuktanqzjsqeb
yidrchgpvwfloiuktanxzjgleb
ymdrchgpxwfloiuktanxzjsxfb
ymdrchgfvwfloiuktanxzjiteb
ymdrchgvvwflqifktanxzjsqeb
ymdrchgplwfloiuktanizjnqeb
ymdrchgpvwfyfiuktafxzjsqeb
ymddchgpvwcloiuktanxzjsqeq
ymdrchgkvwflaiuktanxzjsqfb
yudrchgpvwfzoiuktanxzjsreb
ymdrdhgpvwfloiuktnnxqjsqeb
ymdrnhgpvwfloiuktauxzjdqeb
ymdrchgpvwflsiddtanxzjsqeb
ymdrchgpvwhloeuktanxzjsqek
ymdrchgpvjfioiuktawxzjsqeb
ycdrohgpvwfgoiuktanxzjsqeb
ymdrchgpvwflmifktanxzjsqel
yfdrchrpvwfloruktanxzjsqeb
ymdrchgjvwfloiuktanxzrsqeg
ymarchgpxwfloiukkanxzjsqeb
ymdrchgppwflghuktanxzjsqeb
ymdrchvpvwfloiuktanxpjrqeb
ymdlchgpqjfloiuktanxzjsqeb
ymdrchgpvwfofiuktandzjsqeb
ymdrcngpqwfloiuktanlzjsqeb
ymdrchgpvwfloiuiocnxzjsqeb
ymdrcogpvwfloizktanxzjcqeb
ymdrchgpvlfvoiuksanxzjsqeb
ymdrchgpvwflocpctanxzjsqeb
ymdrchgpvwfloiuktanlzjsejb
yndrchgpvwflzifktanxzjsqeb
ymdrcrgpvkfloiuktanxrjsqeb
ymdrchhpvwslocuktanxzjsqeb
ymdrxhgpvwfloiuwtazxzjsqeb
ymdrchgpvafloiuutanxzjsqxb
ymdrchppvhfloquktanxzjsqeb
ymprcugpvwtloiuktanxzjsqeb
ymdrchgpvvflyiuktanxzjsqvb
ymdrchgovwfloiuftanxzjwqeb
ymdrchrpvwflotyktanxzjsqeb
gmdrchgpvwfloauttanxzjsqeb
ymdrchmpvofloiukmanxzjsqeb
ymdrchgpvwflsiuktanxzjspkb
ymdrchgpvwfloluktajxijsqmb
ymdrcngpvwfloiukbanxzdsqeb
ymdrchgpvwploiuktnnxzmsqeb
ymdrcwgpvwfloiuktbnxhjsqeb
ymdrcngpvwfloiuktaaxbjsqeb
ykdrchgpvwfloiuktanxzgsqej
yuhrchgpvwfdoiuktanxzjsqeb
ymdrchgpvsfloiukbanxujsqeb
ymqrchgpvwfliiuktanxzjsteb
ysdqchgpvwfloiuktanxzjtqeb
ymdjchgpcwfloiuktanxzrsqeb
ymdkchgpvwfloiukfanlzjsqeb
ymdrchgpvxfloikktanxzjiqeb
smdrchgwewfloiuktanxzjsqeb
ymdrchgpvwfljiuktanxajsqer
ymdrchgpowflifuktanxzjsqeb
ymdrchgpvpzloiukoanxzjsqeb
yydrchgwvwfvoiuktanxzjsqeb
ymdgcdgpvwflobuktanxzjsqeb
ymdechgpvkfloiuktanxzjsjeb
ymdnchnpvwfloixktanxzjsqeb
ymdrchgpiefloiuktqnxzjsqeb
ymprchgpvwfloiuktjnxzjsxeb
ymdrjdgpzwfloiuktanxzjsqeb
ymsrchgpywfloiuktanxzjsueb
ymdrchgpvgoloiuktanxzcsqeb
ymdrphgpswflbiuktanxzjsqeb
ymqrchgpvnfloiumtanxzjsqeb
ymjrchgpvwyloiuktacxzjsqeb
ymdrchepvwmlqiuktanxzjsqeb
kmirchgpvwfloiuktanxzjsreb
ymdncygpvwfloiuktanuzjsqeb
ymdrzhgpvwploiuktanxzxsqeb
ymdrchkpvwfloiwkmanxzjsqeb
ywdrchgovwfloiuktanxzjsceb
amdrchgpvwfloiuktanrzjqqeb
ymdpshgpvwfloiuktanxzjyqeb
ymdrcegpvwfloijktcnxzjsqeb
ymdrcygpvwfloiuktanxztsqwb
ymdrchgpvufloiuvtabxzjsqeb
ymdrchgpvwflkiuktrnxzjsqmb
ymdrchgpvqfloiuktanxpjfqeb
ymdrclgpvkfloiyktanxzjsqeb
ybdxchgpvwfloiuktanxzjskeb
pmdrchgpvwfzoirktanxzjsqeb
ycdfchgpvwfloiuktanxzjtqeb
ymdrchgpdwfloiumtbnxzjsqeb
ymdrchgpqmfloiuktanxzjsqer
ymgrchgpvwfroiuktanxzjsqey
ymdrnhgpvwfloiuktanjzjsqlb
dmdrchgpvgfloiuktqnxzjsqeb
yudrchgnvwfloiukranxzjsqeb
ymdrxhgpvafloiuktanxzjsqeq
ymdrchgpvwfyofuktanxzjsueb
ymdrrhgpvwfloiuktavxzjsqpb
yvdrchgpvwfloiuktalxzhsqeb
ymdrchgpbwfloiuktanxzfnqeb
ymdrqhgpvwfloiuvtznxzjsqeb
ymdrchgpvbfloiuetanxzjsqeo
ymdrchjpvwfloiuktanxzjnqrb
ymdrchgpmwfqoiuknanxzjsqeb
ymdrchgpvwfuoiuktaqxzjtqeb
ymdrchgpvwfloiuktamxaosqeb
fmdrchgpvffloiuktanxzjsaeb
ymdrrhglvwfwoiuktanxzjsqeb
ymdrchgpvwflohuktanxzjcqei
ymdrcsgpvwfloiuktaexzjsqek
ymlrchfpvwfloiuktpnxzjsqeb
yxdrchgpvwfdoiuvtanxzjsqeb
ymdrchgrvwfloiuktadxzjsqew
ymdrchgpvwbloiyktandzjsqeb
ymdrchgpvsfloiyktanozjsqeb
ymdrchgpjwfloiuktanxibsqeb
ymdrchgjvyfloiuktanxzjsqeh
ymdrchgvvwfloiuktanzrjsqeb
ymdrchgpvwaloiuktynxzjsqev
ymdrccgpvwflonvktanxzjsqeb
ymdrchgqvffloiuktanxfjsqeb
ymdbchgpvwsloiudtanxzjsqeb
ymdachgpvwfloiuktanlzjsqwb
ymdrclgpvwwloiuktanxzjsjeb
ybdpchgpvwdloiuktanxzjsqeb
ymdtchgpvwfleijktanxzjsqeb
ymdrchgpvwfloiustanxzjsxep
ymdrcjypvwfloiuktanxnjsqeb
ymdrcdgpvwfloiuutanxkjsqeb
yhirchgpvufloiuktanxzjsqeb
ymdrlhgpvwfluigktanxzjsqeb
ywdrhhgpvwftoiuktanxzjsqeb
ymdrchgpvwflyiuktanozjsqtb
cmdrchgpuwfloiukmanxzjsqeb
ymdochgpvrfloiuktanvzjsqeb
ymdrcvgpvwfgoiuktfnxzjsqeb
ymdrchgpmufloiuktanxzssqeb
ymurchgrvwfloiuktanxzjsqep
bmdrchgpvwfloiukpanxzjsqmb
ymdrchgphwvloiuktanszjsqeb
ymdpkhgpvwfloiuktanxzjsqtb
ymdrchgpvwfloiuwtanxzjfqev
ymdrchgpvwfloguktqlxzjsqeb
ymkrshgpvwflgiuktanxzjsqeb
ymdrchgpzwfloizktanxznsqeb
ymdrchgpvxfloiuktegxzjsqeb
yydrchgpwwfloiuktanxzjsqqb
ymdrcngwvwfltiuktanxzjsqeb
ymdszhgwvwfloiuktanxzjsqeb
ymdrchguvwfjoiuktanxzxsqeb
ymdomhgpvwfloiuktanxgjsqeb
ymdrcvgpvwfloiuktanwzzsqeb
yydrchgpvwfloiuktanxzjmqtb
rmdrchgpvwfloiuktmnszjsqeb
ykdrchgpvwfloyuktmnxzjsqeb
ymcrchkpvwfloiuktanxzjsoeb
ymdrcrgpvwfloiukpanxzjsceb
yrdrchgpvwfloiukwanxzjsqhb
ymdrcfgpvwfloiurtanxojsqeb
ymdrchgpuwstoiuktanxzjsqeb
ymdrchgpvwflpxuktanxzjsqer
ymdrehgpvwfloiuktabxdjsqeb
yedrchgpvwfloiukqanxzjiqeb
ymdrthgpvyfloiuktanxzjsqen
cmdlchgpvwfloiuvtanxzjsqeb
ymdrchgpvwtloiuktanlpjsqeb
ymdrchgpvwfloiuktanyvjsqea
gmdrcogpvwfloiuktanxzjsqqb
ymmrchgpvwflosuktauxzjsqeb
ymgrchgjvwfloiuktavxzjsqeb
ymdbclgpvwfloeuktanxzjsqeb
ymdrchgpvwfloiuktaixzcsqfb
ymdrchgpvwflmiuktanxttsqeb
ymxrchgpvwfloiuktanxzfsqec
yqzrchgpcwfloiuktanxzjsqeb
yvdrchgpvwfloiukgvnxzjsqeb
ymdrchepvwfloiuktahxzosqeb
ymdlchgpvwfloiuktamizjsqeb
ymdrchgpcwflovuktanxzjsqzb
yvduchgpvwfloiukaanxzjsqeb
ymdrchgpvwfloiuktxmxzjsgeb
ymdrcrgpvwfloizktanbzjsqeb
amdrchgpvwfloiukhanxzjsqbb
ymdrchgpvwfloluktajxijsqeb
ymdrcfgpvwfloiubtanxznsqeb
ymdrchgpvwfleiuwtanxzjsweb
ymdrchgpvwfzdguktanxzjsqeb
ymdrchgwvwflosyktanxzjsqeb
ymrrchgpvwfloiultanxzjsqez
ymdpchgkvwfleiuktanxzjsqeb
ymdrchgpvwfloijktalxfjsqeb
ymdrchgpmwfloiuktanzzjsqfb
ymdrcsgpvwfljiukyanxzjsqeb
ymdrcarpvwfloiuktapxzjsqeb
ymdrchgpvwfloiuktanxzjcqvs
`
