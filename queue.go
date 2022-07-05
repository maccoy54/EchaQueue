package queue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type SortedKeys []string
var riakkey *MyQueue

func (s SortedKeys) Len() int {
	return len(s)
}

func (s SortedKeys) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortedKeys) Less(i, j int) bool {
	str1 := s[i]
	str2 := s[j]
	//p := fmt.Println
	//pf := fmt.Printf
	//YKY929401-89_2021-10-5_20_47_12_976819.zip
	// tab[0] =  YKY929401-89
	// tab[1] =  2021-10-5 (Date)
	// tab[2] = 20 Heure
	// tab[3] =47 Minutes
	// tab[4] =12 secondes
	// tab[5][0]  =976819 millisecondes . zip

	var an1, an2, j1, j2, h1, h2, m1, m2, s1, s2, ns1, ns2, tmp int
	var mois1, mois2 time.Month

	// Triatement de str1 et str2
	tab := strings.Split(str1, "_")
	tab2 := strings.Split(str2, "_")
	//p(str1 + " - " + str2)
	// cas ou disable dans le nom

	//taille := len(tab)
	//taille2 := len(tab2)
	//pf("%s : %d - %s : %d\n", str1, taille, str2, taille2)

	if len(tab) == 6 {
		an1, _ = strconv.Atoi(strings.Split(tab[1], "-")[0])
		tmp, _ = strconv.Atoi(strings.Split(tab[1], "-")[1])
		mois1 = time.Month(tmp)
		j1, _ = strconv.Atoi(strings.Split(tab[1], "-")[2])
		h1, _ = strconv.Atoi(tab[2])
		m1, _ = strconv.Atoi(tab[3])
		s1, _ = strconv.Atoi(tab[4])
		ns1, _ = strconv.Atoi(strings.Split(tab[5], ".")[0])
	} else {
		//p("split != 6 :" + str1)
		an1, _ = strconv.Atoi(strings.Split(tab[2], "-")[0])
		tmp, _ = strconv.Atoi(strings.Split(tab[2], "-")[1])
		mois1 = time.Month(tmp)
		j1, _ = strconv.Atoi(strings.Split(tab[2], "-")[2])
		h1, _ = strconv.Atoi(tab[3])
		m1, _ = strconv.Atoi(tab[4])
		s1, _ = strconv.Atoi(tab[5])
		ns1, _ = strconv.Atoi(strings.Split(tab[6], ".")[0])
	}
	// Triatement de str2
	if len(tab2) == 6 {
		an2, _ = strconv.Atoi(strings.Split(tab2[1], "-")[0])
		tmp, _ = strconv.Atoi(strings.Split(tab2[1], "-")[1])
		mois2 = time.Month(tmp)
		j2, _ = strconv.Atoi(strings.Split(tab2[1], "-")[2])
		h2, _ = strconv.Atoi(tab2[2])
		m2, _ = strconv.Atoi(tab2[3])
		s2, _ = strconv.Atoi(tab2[4])
		ns2, _ = strconv.Atoi(strings.Split(tab2[5], ".")[0])
	} else {
		//p("split != 6 :" + str2)
		an2, _ = strconv.Atoi(strings.Split(tab2[2], "-")[0])
		tmp, _ = strconv.Atoi(strings.Split(tab2[2], "-")[1])
		mois2 = time.Month(tmp)
		j2, _ = strconv.Atoi(strings.Split(tab2[2], "-")[2])
		h2, _ = strconv.Atoi(tab2[3])
		m2, _ = strconv.Atoi(tab2[4])
		s2, _ = strconv.Atoi(tab2[5])
		ns2, _ = strconv.Atoi(strings.Split(tab2[6], ".")[0])
	}

	ts1 := time.Date(an1, mois1, j1, h1, m1, s1, ns1, time.Local)
	ts2 := time.Date(an2, mois2, j2, h2, m2, s2, ns2, time.Local)

	return ts1.Before(ts2)
}

/*
	Transforme une key riak en date
*/
func getDateFromKey(key string) time.Time {
	var an1, j1, h1, m1, s1, ns1, tmp int
	var mois1 time.Month
	// Triatement de str1 et str2
	tab := strings.Split(key, "_")

	if len(tab) == 6 {
		an1, _ = strconv.Atoi(strings.Split(tab[1], "-")[0])
		tmp, _ = strconv.Atoi(strings.Split(tab[1], "-")[1])
		mois1 = time.Month(tmp)
		j1, _ = strconv.Atoi(strings.Split(tab[1], "-")[2])
		h1, _ = strconv.Atoi(tab[2])
		m1, _ = strconv.Atoi(tab[3])
		s1, _ = strconv.Atoi(tab[4])
		ns1, _ = strconv.Atoi(strings.Split(tab[5], ".")[0])
	} else {
		//p("split != 6 :" + str1)
		an1, _ = strconv.Atoi(strings.Split(tab[2], "-")[0])
		tmp, _ = strconv.Atoi(strings.Split(tab[2], "-")[1])
		mois1 = time.Month(tmp)
		j1, _ = strconv.Atoi(strings.Split(tab[2], "-")[2])
		h1, _ = strconv.Atoi(tab[3])
		m1, _ = strconv.Atoi(tab[4])
		s1, _ = strconv.Atoi(tab[5])
		ns1, _ = strconv.Atoi(strings.Split(tab[6], ".")[0])
	}
	return time.Date(an1, mois1, j1, h1, m1, s1, ns1, time.Local)

}

/*
	Transforme une chaine en date
*/
func getDate(str string) time.Time {
	var an1, j1, h1, m1, s1, ns1, tmp int
	var mois1 time.Month
	d1 := strings.Split(str, "-")
	if len(d1) == 3 {
		an1, _ = strconv.Atoi(d1[0])
		tmp, _ = strconv.Atoi(d1[1])
		mois1 = time.Month(tmp)
		j1, _ = strconv.Atoi(d1[2])
		ts1 := time.Date(an1, mois1, j1, h1, m1, s1, ns1, time.Local)
		return ts1
	} else {
		panic("Date " + str + " invalide")
	}

}

/*
 	Supprime les elements de la liste qui ne sont pas compris entre date1 et date2 inclu
	date1 : string au format aa-mm-jj
	date2 :  : string au format aa-mm-jj
*/
func (s *MyQueue) Tranche(date1 string, date2 string) {

	var tranche []string
	ts1 := getDate(date1)
	ts2 := getDate(date2)

	// Parcourir la liste
	for i := 0; i < len(s.Cle); i++ {
		tsKey := getDateFromKey(s.Cle[i])

		if !(tsKey.Before(ts1) || tsKey.After(ts2)) {
			// Supprime la cle
			tranche = append(tranche, s.Cle[i])
		}
	}
	s.Cle = tranche
}

type MyQueue struct {
	Cle []string `json:"keys"` //Recupere toutes les valeurs du json keys
}


func (q *MyQueue) loadRiak(url string) error {
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if r.Body != nil {
		defer r.Body.Close()
	}

	return json.NewDecoder(r.Body).Decode(q)
}

// Méthode de tri du slice
func (q *MyQueue) byDataArrivee() {
	sort.Sort(SortedKeys(q.Cle))
}

// On retourne le premier element de la queue et on le supprime
func (q *MyQueue) getKey() string {
	x := q.Cle[0]
	// on decale queue vers le haut
	q.Cle = q.Cle[1:]
	return x
}

func CreateQueue(url string, debut string, fin string) []string {
         riakkey = new(MyQueue)
         riakkey.loadRiak(url)
         riakkey.Tranche(debut, fin)
         sort.Sort(SortedKeys(riakkey.Cle))
         return riakkey.Cle
}

func GetKey() (string,int){
        return riakkey.getKey(),len(riakkey.Cle)
}


func main() {

	riakkey = new(MyQueue)
	riakkey.loadRiak()

	fmt.Println(len(riakkey.Cle))
	riakkey.Tranche("2021-01-01", "2021-12-31")
	fmt.Printf("%d", len(riakkey.Cle))
	//sort.Sort(SortedKeys(riakkey.Cle))
	sort.Sort(SortedKeys(riakkey.Cle))

	depile := len(riakkey.Cle)
	for depile > 0 {
		fmt.Println(riakkey.getKey())
		depile = len(riakkey.Cle)
	}
} 
