package main

import (
	"context"
	"flag"
	"fmt"
	"log"
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
	var daneZGry danezgry.DaneZGry

	// dołączamy do gry graID
	stanGry := dolaczDoGry(c, *graID, *nazwa)
	for {
		// wypisuję stan gry na ekranie
		drukujStatus(stanGry)
		if stanGry.CzyKoniec {
			return
		}

		daneZGry.PobierzDaneZeStanuGry(stanGry)
		for {
			// gracz podaje kartę na konsoli
			// karta = wczytajKarte()
			if daneZGry.NaszePole < PIERWSZA_FAZA {
				karta = wybierzKarte(daneZGry, stanGry)
			} else {
				karta = randomowaKarta(stanGry)
			}

			if _, ok := kartyDlaKtorychTrzebaPodacKolor[karta]; ok {
				// kolor = wczytajKolor()
				kolor = randomowyKolor()
			} else {
				kolor = proto.KolorZolwia_XXX
			}

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

func randomowyKolor() proto.KolorZolwia {
	return proto.KolorZolwia_BLUE
}

func randomowaKarta(stanGry *proto.StanGry) proto.Karta {
	return stanGry.TwojeKarty[0]
}

func wybierzKarte(dane danezgry.DaneZGry, stanGry *proto.StanGry) proto.Karta {
	for _, z := range dane.ZolwiePodNami {
		karta, ok := najlepszaKartaDla(z, stanGry)
		if ok {
			return karta
		}
	}
	karta, ok := najlepszaKartaDla(stanGry.TwojKolor, stanGry)
	if ok {
		return karta
	}
	return randomowaKarta(stanGry)
}

func najlepszaKartaDla(zolw proto.KolorZolwia, stanGry *proto.StanGry) (proto.Karta, bool) {
	kolor := proto.KolorZolwia_name[int32(zolw)]
	literaKoloru := kolor[:1]

	kartyDoPrzodu := []proto.Karta{}
	for _, k := range stanGry.TwojeKarty {
		if !strings.HasSuffix(k.String(), "B") {
			kartyDoPrzodu = append(kartyDoPrzodu, k)
		}
	}
	if len(kartyDoPrzodu) == 0 {
		return proto.Karta_XX, false
	}

	kartyNaMnie := []proto.Karta{}
	for _, k := range kartyDoPrzodu {
		if strings.HasPrefix(k.String(), literaKoloru) || strings.HasPrefix(k.String(), "A") {
			kartyNaMnie = append(kartyNaMnie, k)
		}
	}
	if len(kartyNaMnie) == 0 {
		return proto.Karta_XX, false
	}
	for _, k := range kartyNaMnie {
		if k.String()[1:2] == "2" {
			return k, true
		}
	}
	return kartyNaMnie[0], true
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
