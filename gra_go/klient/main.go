package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	danezgry "github.com/slaraz/turniej/gra_go/klient/DaneZGry"
	"github.com/slaraz/turniej/gra_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	IP_ADDR               = "localhost:50051"
	NOWY_MECZ_TIMEOUT     = time.Second * 5
	DOLACZ_DO_GRY_TIMEOUT = time.Second * 1000
	RUCH_GRACZA_TIMEOUT   = time.Second * 1000
)

const (
	PIERWSZA_FAZA = 2
)

var (
	addr  = flag.String("addr", IP_ADDR, "adres serwera gry")
	nazwa = flag.String("nazwa", "Ziutek", "nazwa gracza")
	nowa  = flag.Bool("nowa", false, "tworzy nową grę na serwerze")
	graID = flag.String("gra", "", "dołącza do gry o podanym id")
	lg    = flag.Int("lg", 2, "określa liczbę graczy")
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Koniec.")

	flag.Parse()

	// Utowrzenie połączenia z serwerem gry.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial: %v", err)
	}
	defer conn.Close()
	c := proto.NewGraClient(conn)

	conn.GetState()
	// Jeśli podano opcję -nowa, to utwórz nową grę.
	if *nowa {
		ctx, cancel := context.WithTimeout(context.Background(), NOWY_MECZ_TIMEOUT)
		defer cancel()

		nowaGraInfo, err := c.NowyMecz(ctx, &proto.KonfiguracjaGry{LiczbaGraczy: int32(*lg)})
		if err != nil {
			log.Fatalf("c.NowyMecz: %v", err)
		}
		log.Printf("Nowa gra %q\n", nowaGraInfo.GraID)

		*graID = nowaGraInfo.GraID
	}

	// Jeśli nie utworzono -nowa,
	// ani nie podano opcji -gra, to kończymy.
	if *graID == "" {
		flag.Usage()
		return
	}

	var (
		kartyDlaKtorychTrzebaPodacKolor = map[proto.Karta]bool{
			proto.Karta_L1:  true,
			proto.Karta_L2:  true,
			proto.Karta_A1:  true,
			proto.Karta_A1B: true,
		}
		karta proto.Karta
		kolor proto.KolorZolwia
	)

	// przebieg gry
	daneZGry := new(danezgry.DaneZGry)

	// dołączamy do gry graID
	stanGry := dolaczDoGry(c, *graID, *nazwa)
	for {
		// wypisuję stan gry na ekranie
		drukujStatus(stanGry)
		if stanGry.CzyKoniec {
			return
		}

		daneZGry.ZaktualizujDaneZeStanuGry(stanGry)
		for {
			if !czyKtorysWyszedl(stanGry) {
				stanGry.TwojeKarty = usunLasty(stanGry)
			}
			// gracz podaje kartę na konsoli
			// karta = wczytajKarte()
			// wybranie karty i ewentualnie koloru
			karta, kolor = wybierzRuch(stanGry, daneZGry)
			if _, ok := kartyDlaKtorychTrzebaPodacKolor[karta]; !ok {
				kolor = proto.KolorZolwia_XXX
			}

			log.Printf("wybrany ruch: (%v:%v)", karta, kolor)

			// wysyłam ruch do serwera
			nowyStan, err := wyslijRuch(c, &proto.RuchGracza{
				GraID:        stanGry.GraID,
				GraczID:      stanGry.GraczID,
				ZagranaKarta: karta,
				KolorWybrany: kolor,
			})
			if err != nil && status.Code(err) == codes.InvalidArgument {
				// zły ruch
				fmt.Printf("Błąd ruchu: %v\n", err)
				continue
			} else if err != nil {
				// inny błąd, np. połączenie z serwerem
				log.Fatalf("wyslijRuch: status: %v, err: %v", status.Code(err), err)
			}
			// ruch ok
			stanGry = nowyStan
			break
		}
	}
}

func wybierzRuch(stanGry *proto.StanGry, daneZGry *danezgry.DaneZGry) (proto.Karta, proto.KolorZolwia) {
	if daneZGry.NaszePole <= PIERWSZA_FAZA {
		return wybierzRuchPierwszaFazaGry(stanGry, daneZGry)
	} else {
		return wybierzRuchDrugaFazaGry(stanGry, daneZGry)
	}
}

func czyKtorysWyszedl(stanGry *proto.StanGry) bool {
	for _, p := range stanGry.Plansza {
		for _, z := range p.Zolwie {
			if stanGry.TwojKolor != z {
				return true
			}
		}
	}
	return false
}

func usunLasty(stanGry *proto.StanGry) []proto.Karta {
	zagrywalne := []proto.Karta{}
	for _, k := range stanGry.TwojeKarty {
		if !strings.HasPrefix(k.String(), "L") {
			zagrywalne = append(zagrywalne, k)
		}
	}
	return zagrywalne
}

func randomowaKarta(stanGry *proto.StanGry) proto.Karta {
	min := 0
	max := len(stanGry.TwojeKarty)
	return stanGry.TwojeKarty[rand.Intn(max-min)+min]
}

func wybierzRuchPierwszaFazaGry(stanGry *proto.StanGry, daneZGry *danezgry.DaneZGry) (proto.Karta, proto.KolorZolwia) {
	// sprobowac zagrac najlepsza karte dla zolwi pod nami od najdalszego
	for _, kz := range daneZGry.ZolwiePodNami {
		karta, kolor, ok := najlepszyRuchDla(kz, stanGry)
		if ok {
			return karta, kolor
		}
	}

	// sprobowac zagrac jak najlepsza karte dla naszego zolwia
	karta, kolor, ok := najlepszyRuchDla(stanGry.TwojKolor, stanGry)
	if ok {
		return karta, kolor
	}

	// random karta?
	return randomowyRuch(stanGry, daneZGry)
}

func wybierzRuchDrugaFazaGry(stanGry *proto.StanGry, daneZGry *danezgry.DaneZGry) (proto.Karta, proto.KolorZolwia) {
	kiedy := rand.Intn(10)
	if kiedy < 7 {
		if karta, ok := cofaj(daneZGry.Lider, stanGry); ok {
			return karta, daneZGry.Lider
		}
	}
	return wybierzRuchPierwszaFazaGry(stanGry, daneZGry)
}

func cofaj(kogo proto.KolorZolwia, stanGry *proto.StanGry) (proto.Karta, bool) {
	if kogo == stanGry.TwojKolor {
		return proto.Karta_XX, false
	}
	kartyCofajace := []proto.Karta{}
	for _, k := range stanGry.TwojeKarty {
		if strings.HasSuffix(k.String(), "B") {
			kartyCofajace = append(kartyCofajace, k)
		}
	}
	if len(kartyCofajace) == 0 {
		return proto.Karta_XX, false
	}
	for _, k := range kartyCofajace {
		if strings.HasPrefix(k.String(), kogo.String()[0:1]) {
			return k, true
		}
	}
	return proto.Karta_XX, false
}

func randomowyRuch(stanGry *proto.StanGry, daneZGry *danezgry.DaneZGry) (proto.Karta, proto.KolorZolwia) {
	var kolor proto.KolorZolwia
	var karta proto.Karta

	oryginalnaReka := stanGry.TwojeKarty

	for {
		if len(stanGry.TwojeKarty) == 0 {
			fmt.Printf("Oryginalna ręka: %v", oryginalnaReka)
			karta = proto.Karta_XX
			kolor = proto.KolorZolwia_XXX
		}
		indeksKarty := rand.Intn(len(stanGry.TwojeKarty))
		karta = stanGry.TwojeKarty[indeksKarty]
		if karta == proto.Karta_L1 || karta == proto.Karta_L2 {
			if len(daneZGry.OstatnieZolwie) == 0 {
				stanGry.TwojeKarty = usunKarte(stanGry.TwojeKarty, karta)
				continue
			}
			indeksKoloruOstatniego := rand.Intn(len(daneZGry.OstatnieZolwie))
			kolor = daneZGry.OstatnieZolwie[indeksKoloruOstatniego]
			break
		} else if karta == proto.Karta_A1B ||
			karta == proto.Karta_B1B ||
			karta == proto.Karta_G1B ||
			karta == proto.Karta_P1B ||
			karta == proto.Karta_R1B ||
			karta == proto.Karta_Y1B {
			if len(daneZGry.ZolwieKtoreMoznaCofac) == 0 {
				stanGry.TwojeKarty = usunKarte(stanGry.TwojeKarty, karta)
				continue
			}
			if strings.HasPrefix(karta.String(), stanGry.TwojKolor.String()[0:1]) {
				stanGry.TwojeKarty = usunKarte(stanGry.TwojeKarty, karta)
				continue
			}
			indeksKoloruDoCofania := rand.Intn(len(daneZGry.ZolwieKtoreMoznaCofac))
			kolor = daneZGry.ZolwieKtoreMoznaCofac[indeksKoloruDoCofania]
			break
		} else {
			const iloscKolorow = 5
			indeksKoloru := rand.Int31n(iloscKolorow)
			kolor = proto.KolorZolwia(indeksKoloru)
			break
		}
	}

	return karta, kolor
}

func usunKarte(karty []proto.Karta, doUsuniecia proto.Karta) []proto.Karta {
	res := []proto.Karta{}
	for _, k := range karty {
		if k != doUsuniecia {
			res = append(res, k)
		}
	}
	return res
}

func najlepszyRuchDla(zolw proto.KolorZolwia, stanGry *proto.StanGry) (proto.Karta, proto.KolorZolwia, bool) {
	kolor := proto.KolorZolwia_name[int32(zolw)]
	literaKoloru := kolor[:1]

	kartyDoPrzodu := []proto.Karta{}
	for _, k := range stanGry.TwojeKarty {
		if !strings.HasSuffix(k.String(), "B") {
			kartyDoPrzodu = append(kartyDoPrzodu, k)
		}
	}

	if len(kartyDoPrzodu) == 0 {
		return proto.Karta_XX, proto.KolorZolwia_XXX, false
	}

	kartyNaMnie := []proto.Karta{}
	for _, k := range kartyDoPrzodu {
		if strings.HasPrefix(k.String(), literaKoloru) || strings.HasPrefix(k.String(), "A") {
			kartyNaMnie = append(kartyNaMnie, k)
		}
	}

	if len(kartyNaMnie) == 0 {
		return proto.Karta_XX, proto.KolorZolwia_XXX, false
	}

	for _, k := range kartyNaMnie {
		if k.String()[1:2] == "2" {
			return k, zolw, true
		}
	}

	return kartyNaMnie[0], zolw, true
}

func dolaczDoGry(c proto.GraClient, graID, nazwa string) *proto.StanGry {
	log.Printf("Gracz %s dołącza do gry %q", nazwa, graID)
	ctx, cancel := context.WithTimeout(context.Background(), DOLACZ_DO_GRY_TIMEOUT)
	defer cancel()
	log.Println("Czekam na odpowiedź od serwera...")
	stanGry, err := c.DolaczDoGry(ctx, &proto.Dolaczanie{
		GraID:       graID,
		NazwaGracza: nazwa,
	})
	if err != nil {
		log.Fatalf("c.Dolacz: %v", err)
	}
	return stanGry
}

func wyslijRuch(c proto.GraClient, ruch *proto.RuchGracza) (*proto.StanGry, error) {
	log.Printf("Gracz %s-%s zagrywa kartę: %v", ruch.GraID, ruch.GraczID, ruch.ZagranaKarta)
	ctx, cancel := context.WithTimeout(context.Background(), RUCH_GRACZA_TIMEOUT)
	defer cancel()
	log.Println("Czekam na odpowiedź od serwera (ruch przeciwnika)...")

	return c.MojRuch(ctx, ruch)
}

func drukujStatus(stanGry *proto.StanGry) {
	if stanGry.CzyKoniec {
		fmt.Println("Koniec gry, wygrał gracz nr", stanGry.KtoWygral)
	} else {
		fmt.Printf("Twój kolor: %v, Pola:", stanGry.TwojKolor)
		for _, pole := range stanGry.Plansza {
			fmt.Printf(" %v", pole.Zolwie)
		}
		fmt.Printf(", Twoje karty: %v\n", stanGry.TwojeKarty)
	}
}
