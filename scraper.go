package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"io/ioutil"
	"strings"
	"time"
)

type centro struct {
	Nome      string
	Provincia string
	Localita  string
	Cap       string
	Via       string
	Telefono  string
	Email     string
	Specifica []string
	Regione string
	Azienda string
	Url     string
}

type unisci struct {
	Regione string
	Centri  []string
}

type scanTotale struct {
	Url    string
	Totale []unisci
}

func main() {
	readFromJson()
	return
	elencoSiti := []string{"https://www.centriassistenza.org/asus.html", "https://www.centriassistenza.org/alcatel.html", "https://www.centriassistenza.org/oppo.html", "https://www.centriassistenza.org/sony-denon-yamaha-onkyo-focal-audio.html", "https://www.centriassistenza.org/tre-assistenza-e-store.html", "https://www.centriassistenza.org/xiaomi-assistenza-italia.html", "https://www.centriassistenza.org/tim-assistenza.html", "https://www.centriassistenza.org/meizu.html", "https://www.centriassistenza.org/apple.html", "https://www.centriassistenza.org/lg-smartphone.html", "https://www.centriassistenza.org/samsung-smartphone.html", "https://www.centriassistenza.org/archos.html", "https://www.centriassistenza.org/lumia-smartphone.html", "https://www.centriassistenza.org/wiko-smartphone.html", "https://www.centriassistenza.org/mediacom.html", "https://www.centriassistenza.org/neffos.html", "https://www.centriassistenza.org/blackberry-smartphone.html", "https://www.centriassistenza.org/huawei.html", "https://www.centriassistenza.org/brondi.html", "https://www.centriassistenza.org/garmin.html", "https://www.centriassistenza.org/zte.html", "https://www.centriassistenza.org/honor.html", "https://www.centriassistenza.org/goclever.html", "https://www.centriassistenza.org/polar.html", "https://www.centriassistenza.org/nodis.html", "https://www.centriassistenza.org/breil.html", "https://www.centriassistenza.org/citizen.html", "https://www.centriassistenza.org/seiko.html", "https://www.centriassistenza.org/trevi.html", "https://www.centriassistenza.org/lg.html", "https://www.centriassistenza.org/samsung.html", "https://www.centriassistenza.org/hisense.html", "https://www.centriassistenza.org/telefunken-tv.html", "https://www.centriassistenza.org/loewe.html", "https://www.centriassistenza.org/humax.html", "https://www.centriassistenza.org/panasonic.html", "https://www.centriassistenza.org/philips-tv.html", "https://www.centriassistenza.org/schaub-lorenz.html", "https://www.centriassistenza.org/sharp.html", "https://www.centriassistenza.org/sony.html", "https://www.centriassistenza.org/telesystem.html", "https://www.centriassistenza.org/toshiba.html", "https://www.centriassistenza.org/united.html", "https://www.centriassistenza.org/pioneer.html", "https://www.centriassistenza.org/sky-service.html", "https://www.centriassistenza.org/mediasetpremium-installatori.html", "https://www.centriassistenza.org/canon.html", "https://www.centriassistenza.org/changhong.html", "https://www.centriassistenza.org/funai.html", "https://www.centriassistenza.org/new-majestic-audiola.html", "https://www.centriassistenza.org/epson-stampanti-laser-videoproiettori.html", "https://www.centriassistenza.org/jvc-kenwood.html", "https://www.centriassistenza.org/hannspree.html", "https://www.centriassistenza.org/brother-stampanti.html", "https://www.centriassistenza.org/midland.html", "https://www.centriassistenza.org/exhibo.html", "https://www.centriassistenza.org/candy.html", "https://www.centriassistenza.org/girmi.html", "https://www.centriassistenza.org/bimby-vorwerk.html", "https://www.centriassistenza.org/sharp-lavatrici-frigoriferi.html", "https://www.centriassistenza.org/gorenje.html", "https://www.centriassistenza.org/beko.html", "https://www.centriassistenza.org/folletto-vorwerk.html", "https://www.centriassistenza.org/ardes.html", "https://www.centriassistenza.org/ariete.html", "https://www.centriassistenza.org/bialetti.html", "https://www.centriassistenza.org/termozeta.html", "https://www.centriassistenza.org/gaggia.html", "https://www.centriassistenza.org/polti.html", "https://www.centriassistenza.org/braun.html", "https://www.centriassistenza.org/elchim.html", "https://www.centriassistenza.org/caffitaly.html", "https://www.centriassistenza.org/imperia.html", "https://www.centriassistenza.org/smalvic.html", "https://www.centriassistenza.org/la-pavoni.html", "https://www.centriassistenza.org/g3-ferrari.html", "https://www.centriassistenza.org/krups.html", "https://www.centriassistenza.org/johnson-elettrodomestici.html", "https://www.centriassistenza.org/tefal.html", "https://www.centriassistenza.org/imetec.html", "https://www.centriassistenza.org/bosch-elettrodomestici.html", "https://www.centriassistenza.org/glem-cucine.html", "https://www.centriassistenza.org/miele-elettrodomestici.html", "https://www.centriassistenza.org/general-electric-elettrodomestici.html", "https://www.centriassistenza.org/whirlpool-indesit-ignis-hotpoint-bauknecht-scholtes-kitchenaid.html", "https://www.centriassistenza.org/electrolux-aeg-castor-kelvinator-philco-rex-zanussi-zoppas.html", "https://www.centriassistenza.org/sangiorgio-elettrodomestici.html", "https://www.centriassistenza.org/saeco.html", "https://www.centriassistenza.org/gammapiu.html", "https://www.centriassistenza.org/fogacci.html", "https://www.centriassistenza.org/daewoo-elettrodomestici.html", "https://www.centriassistenza.org/ocean.html", "https://www.centriassistenza.org/rowenta.html", "https://www.centriassistenza.org/smeg.html", "https://www.centriassistenza.org/moulinex.html", "https://www.centriassistenza.org/delonghi.html", "https://www.centriassistenza.org/peg-perego.html", "https://www.centriassistenza.org/tata-daitsu-fujielectric.html", "https://www.centriassistenza.org/samsung-climatizzatori.html", "https://www.centriassistenza.org/boschetti.html", "https://www.centriassistenza.org/accorroni.html", "https://www.centriassistenza.org/savio-caldaie.html", "https://www.centriassistenza.org/climatizzatori-daewoo.html", "https://www.centriassistenza.org/trane.html", "https://www.centriassistenza.org/climatizzatori-hermann-saunier-duval.html", "https://www.centriassistenza.org/haier-climatizzatori.html", "https://www.centriassistenza.org/bluebox.html", "https://www.centriassistenza.org/galletti-climatizzazione.html", "https://www.centriassistenza.org/climatizzatori-lg.html", "https://www.centriassistenza.org/toshiba-climatizzatori.html", "https://www.centriassistenza.org/buderus.html", "https://www.centriassistenza.org/de-dietrich.html", "https://www.centriassistenza.org/emmeti-caldaie-climatizzatori.html", "https://www.centriassistenza.org/climaveneta.html", "https://www.centriassistenza.org/clivet.html", "https://www.centriassistenza.org/mt-stufe.html", "https://www.centriassistenza.org/sabiana.html", "https://www.centriassistenza.org/atag-caldaie.html", "https://www.centriassistenza.org/pasian.html", "https://www.centriassistenza.org/lincar-stufe.html", "https://www.centriassistenza.org/lamborghini-caloreclima.html", "https://www.centriassistenza.org/italtherm.html", "https://www.centriassistenza.org/nova-florida.html", "https://www.centriassistenza.org/thermital.html", "https://www.centriassistenza.org/jolly-mec.html", "https://www.centriassistenza.org/ici-caldaie.html", "https://www.centriassistenza.org/castelmonte-stufe.html", "https://www.centriassistenza.org/laminox.html", "https://www.centriassistenza.org/joannes-caldaie.html", "https://www.centriassistenza.org/fondital.html", "https://www.centriassistenza.org/junkers.html", "https://www.centriassistenza.org/rinnai.html", "https://www.centriassistenza.org/innova.html", "https://www.centriassistenza.org/mcquay.html", "https://www.centriassistenza.org/maxa-climatizzatori.html", "https://www.centriassistenza.org/carrier.html", "https://www.centriassistenza.org/panasonic-climatizzatori.html", "https://www.centriassistenza.org/hitachi-climatizzatori.html", "https://www.centriassistenza.org/sharp-climatizzatori.html", "https://www.centriassistenza.org/mitsubishi-climatizzatori.html", "https://www.centriassistenza.org/daikin.html", "https://www.centriassistenza.org/fujitsu-climatizzatori.html", "https://www.centriassistenza.org/gree-argoclima.html", "https://www.centriassistenza.org/diloc.html", "https://www.centriassistenza.org/aermec.html", "https://www.centriassistenza.org/olimpia-splendid.html", "https://www.centriassistenza.org/comfee.html", "https://www.centriassistenza.org/midea.html", "https://www.centriassistenza.org/zephir-climatizzatori.html", "https://www.centriassistenza.org/radiant-caldaie-bruciatori.html", "https://www.centriassistenza.org/paradigma-stufe.html", "https://www.centriassistenza.org/kloben.html", "https://www.centriassistenza.org/superior-stufe.html", "https://www.centriassistenza.org/marocchi-stufe.html", "https://www.centriassistenza.org/rika-stufe.html", "https://www.centriassistenza.org/ravelli-stufe.html", "https://www.centriassistenza.org/freepoint-stufe.html", "https://www.centriassistenza.org/wolf.html", "https://www.centriassistenza.org/italkero.html", "https://www.centriassistenza.org/optima-stufe.html", "https://www.centriassistenza.org/sile-caldaie.html", "https://www.centriassistenza.org/beretta-riscaldamento.html", "https://www.centriassistenza.org/viessmann-caldaie.html", "https://www.centriassistenza.org/baxi-caldaie.html", "https://www.centriassistenza.org/immergas-caldaie.html", "https://www.centriassistenza.org/vaillant.html", "https://www.centriassistenza.org/vortice.html", "https://www.centriassistenza.org/ferroli-caldaie.html", "https://www.centriassistenza.org/elco-caldaie.html", "https://www.centriassistenza.org/edilkamin.html", "https://www.centriassistenza.org/cosmogas-caldaie.html", "https://www.centriassistenza.org/biasi-caldaie.html", "https://www.centriassistenza.org/chaffoteaux-caldaie.html", "https://www.centriassistenza.org/robur.html", "https://www.centriassistenza.org/la-nordica-extraflame.html", "https://www.centriassistenza.org/sime-caldaie.html", "https://www.centriassistenza.org/ariston-caldaie.html", "https://www.centriassistenza.org/hermann-caldaie.html", "https://www.centriassistenza.org/piazzetta.html", "https://www.centriassistenza.org/elledi-stufe.html", "https://www.centriassistenza.org/mcz-stufe.html", "https://www.centriassistenza.org/thermorossi.html", "https://www.centriassistenza.org/cs-thermos.html", "https://www.centriassistenza.org/vibrok-stufe-forni-barbeque.html", "https://www.centriassistenza.org/prisma-stufe.html", "https://www.centriassistenza.org/famar-stufe-caldaie.html", "https://www.centriassistenza.org/montegrappa-caminetti.html", "https://www.centriassistenza.org/klover-stufe.html", "https://www.centriassistenza.org/cola-stufe.html", "https://www.centriassistenza.org/palazzetti.html", "https://www.centriassistenza.org/karmek-one.html", "https://www.centriassistenza.org/lartistico-stufe-camini.html", "https://www.centriassistenza.org/cadel.html", "https://www.centriassistenza.org/ct-pasqualicchio.html", "https://www.centriassistenza.org/unical.html", "https://www.centriassistenza.org/riello-caldaie.html", "https://www.centriassistenza.org/truma.html", "https://www.centriassistenza.org/termet.html", "https://www.centriassistenza.org/ditec-entrematic.html", "https://www.centriassistenza.org/zodiac-poolcare.html", "https://www.centriassistenza.org/elkron-antifurti-antintrusione.html", "https://www.centriassistenza.org/rib-automazione.html", "https://www.centriassistenza.org/nice-automazione.html", "https://www.centriassistenza.org/lince-antifurti.html", "https://www.centriassistenza.org/serai.html", "https://www.centriassistenza.org/tecnoalarm.html", "https://www.centriassistenza.org/lofra.html", "https://www.centriassistenza.org/beghelli.html", "https://www.centriassistenza.org/gardena.html", "https://www.centriassistenza.org/faac.html", "https://www.centriassistenza.org/came-automazione.html", "https://www.centriassistenza.org/bft-automazione.html", "https://www.centriassistenza.org/sanitrit.html", "https://www.centriassistenza.org/aprimatic-automazione.html", "https://www.centriassistenza.org/bticino.html", "https://www.centriassistenza.org/comelit.html", "https://www.centriassistenza.org/honeywell.html", "https://www.centriassistenza.org/franke-cucine.html", "https://www.centriassistenza.org/urmet.html", "https://www.centriassistenza.org/jacuzzi.html", "https://www.centriassistenza.org/inda.html", "https://www.centriassistenza.org/grohe.html", "https://www.centriassistenza.org/faber.html", "https://www.centriassistenza.org/bompani.html", "https://www.centriassistenza.org/teuco.html", "https://www.centriassistenza.org/grandform.html", "https://www.centriassistenza.org/ideal-standard.html", "https://www.centriassistenza.org/barazza.html", "https://www.centriassistenza.org/combivox.html", "https://www.centriassistenza.org/vimar-elvox.html", "https://www.centriassistenza.org/ave.html", "https://www.centriassistenza.org/gardenitalia.html", "https://www.centriassistenza.org/shindaiwa.html", "https://www.centriassistenza.org/efco.html", "https://www.centriassistenza.org/mtd.html", "https://www.centriassistenza.org/kawasaki-motori.html", "https://www.centriassistenza.org/ryobi-elettroutensili-assistenza.html", "https://www.centriassistenza.org/bertolini.html", "https://www.centriassistenza.org/bcs-vendita-assistenza.html", "https://www.centriassistenza.org/alpina.html", "https://www.centriassistenza.org/mcculloch.html", "https://www.centriassistenza.org/valex.html", "https://www.centriassistenza.org/grundfos-pompe.html", "https://www.centriassistenza.org/karcher.html", "https://www.centriassistenza.org/einhell.html", "https://www.centriassistenza.org/singer.html", "https://www.centriassistenza.org/milwaukee.html", "https://www.centriassistenza.org/necchi.html", "https://www.centriassistenza.org/hitachi-utensili.html", "https://www.centriassistenza.org/bosch-elettroutensili.html", "https://www.centriassistenza.org/makita.html", "https://www.centriassistenza.org/stihl-viking.html", "https://www.centriassistenza.org/black-and-decker.html", "https://www.centriassistenza.org/rems.html", "https://www.centriassistenza.org/dewalt.html", "https://www.centriassistenza.org/husqvarna.html", "https://www.centriassistenza.org/abac-compressori.html", "https://www.centriassistenza.org/al-ko.html", "https://www.centriassistenza.org/stayer.html", "https://www.centriassistenza.org/echo.html", "https://www.centriassistenza.org/oleomac.html", "https://www.centriassistenza.org/landini-trattori.html", "https://www.centriassistenza.org/honda-rete-assistenza-motori.html", "https://www.centriassistenza.org/mercury-marine.html", "https://www.centriassistenza.org/yanmar-motori.html", "https://www.centriassistenza.org/beta-moto-vendita-assistenza.html", "https://www.centriassistenza.org/mvagusta-vendita-assistenza.html", "https://www.centriassistenza.org/isuzu-vendita-assistenza.html", "https://www.centriassistenza.org/moto-suzuki-assistenza-vendita.html", "https://www.centriassistenza.org/suzuki-auto-assistenza-vendita.html", "https://www.centriassistenza.org/subaru-assistenza-officine-autorizzate.html", "https://www.centriassistenza.org/assistenza-daihatsu-officine-autorizzate.html", "https://www.centriassistenza.org/honda-moto-vendita-officine-autorizzate.html", "https://www.centriassistenza.org/honda-auto-vendita-assistenza.html", "https://www.centriassistenza.org/honda-marine-vendita-assistenza.html", "https://www.centriassistenza.org/harley-davidson-concessionari-vendita-assistenza.html", "https://www.centriassistenza.org/aprilia-moto-scooter-vendita-assistenza.html", "https://www.centriassistenza.org/peugeot-scooter-assistenza.html", "https://www.centriassistenza.org/assistenza-e-bike-biciclette-a-pedalata-assistita.html", "https://www.centriassistenza.org/ktm-vendita-assistenza.html", "https://www.centriassistenza.org/sym-scooter-assistenza.html", "https://www.centriassistenza.org/tohatsu-fuoribordo.html", "https://www.centriassistenza.org/selva-marine-fuoribordo.html", "https://www.centriassistenza.org/piaggio.html", "https://www.centriassistenza.org/evinrude-fuoribordo.html", "https://www.centriassistenza.org/bmw-moto-vendita-assistenza.html", "https://www.centriassistenza.org/husqvarna-moto-vendita-assistenza.html", "https://www.centriassistenza.org/kia-officine-autorizzate.html", "https://www.centriassistenza.org/kawasaki-moto-vendita-assistenza.html", "https://www.centriassistenza.org/kymco-scooter-vendita-assistenza.html", "https://www.centriassistenza.org/ssangyong-officine-autorizzate.html", "https://www.centriassistenza.org/seat-officine-autorizzate.html", "https://www.centriassistenza.org/hyundai-officine-autorizzate.html", "https://www.centriassistenza.org/triumph-moto.html", "https://www.centriassistenza.org/ducati-moto.html", "https://www.centriassistenza.org/gilera-vendita-assistenza.html", "https://www.centriassistenza.org/fanticmotor-vendita-assistenza.html", "https://www.centriassistenza.org/assistenza-climatizzatori-condizionatori.html", "https://www.centriassistenza.org/assistenza-caldaie.html", "https://www.centriassistenza.orghttps://www.centriassistenza.org/lombardia/assistenza-daikin-milano.html", "https://www.centriassistenza.orghttps://www.centriassistenza.org/sardegna/assistenza-viessmann-cagliari.html", "https://www.centriassistenza.org/fiscale/sedi-caf-cna.html", "https://www.centriassistenza.org/fiscale/sedi-caf-acli.html", "https://www.centriassistenza.org/fiscale/sedi-caf-uil.html", "https://www.centriassistenza.org/fiscale/sedi-caf-cisl.html"}
	var sitiVisistati []string
	var lsScan []scanTotale

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.centriassistenza.org"),
		colly.Async(true),
	)

	c.OnHTML(".main", func(e *colly.HTMLElement) {
		var lsRegioni []string
		var centri []string
		var totale []unisci
		start := true

		e.ForEach("p", func(i int, element *colly.HTMLElement) {
			if i < 7 {
				return
			}

			isCentro := true
			element.ForEach("span[style='background:#53ac53;font-family: verdana; font-size: 16px; color: #ffffff;']", func(i int, element *colly.HTMLElement) {
				isCentro = false

				if start {
					lsRegioni = append(lsRegioni, strings.TrimSpace(element.Text))
					start = false
				} else {
					u := unisci{
						Regione: lsRegioni[len(lsRegioni)-1],
						Centri:  centri,
					}
					totale = append(totale, u)
					centri = nil
					lsRegioni = append(lsRegioni, strings.Trim(element.Text, " "))
				}
			})

			if isCentro {
				elm := strings.TrimSpace(element.Text)
				if !strings.HasPrefix(element.Text, "  >") && len(elm) > 0 {
					centri = append(centri, elm)
				}
			}

		})

		scan := scanTotale{
			Url:    e.Request.URL.String(),
			Totale: totale,
		}

		lsScan = append(lsScan, scan)
	})

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		Parallelism: 4,
		RandomDelay: 5 * time.Second,
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
		sitiVisistati = append(sitiVisistati, r.URL.String())
	})

	// Crawl all reddits the user passes in
	for _, sito := range elencoSiti {
		c.Visit(sito)
	}

	c.Wait()

	var sitiNonVisitati []string

	for i := 0; i < len(elencoSiti); i++ {
		visitato := false
		for j := 0; j < len(sitiVisistati); j++ {
			if elencoSiti[i] == sitiVisistati[j] {
				visitato = true
				break
			}
		}

		if !visitato {
			sitiNonVisitati = append(sitiNonVisitati, elencoSiti[i])
		}
	}

	fmt.Print(sitiNonVisitati)
	saveToJson(lsScan, "output.json")
	readFromJson()
}

func saveToJson(totale interface{}, filename string) {
	totJson, err := json.Marshal(totale)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile(filename, totJson, 0644)
}

func readFromJson() {
	// read file
	data, err := ioutil.ReadFile("./output.json")
	if err != nil {
		fmt.Print(err)
	}

	// json data
	var lsScans []scanTotale

	// unmarshall it
	err = json.Unmarshal(data, &lsScans)
	if err != nil {
		fmt.Println("error:", err)
	}

	var lsCentri []centro

	for _, element := range lsScans {
		for _, elTot := range element.Totale {
			for _, elCentro := range elTot.Centri {
				splittedInfo := strings.Split(elCentro, "\n")
				aziendaSplit := strings.Split(element.Url, "/")
				azienda := strings.TrimSuffix(aziendaSplit[len(aziendaSplit)-1], ".html")

				infoToAdd := guessFiledToParse(splittedInfo)

				infoCentro := centro{
					Nome:      splittedInfo[0],
					Via:       infoToAdd["via"],
					Telefono:  infoToAdd["tel"],
					Cap:       infoToAdd["cap"],
					Provincia: infoToAdd["prov"],
					Localita:  infoToAdd["loc"],
					Email:     infoToAdd["email"],
					Specifica: nil,
					Regione:   elTot.Regione,
					Azienda:   azienda,
					Url:       element.Url,
				}

				lsCentri = append(lsCentri, infoCentro)
			}
		}
	}
	saveToJson(lsCentri, "formatted.json")
	fmt.Println(lsCentri)
}

func guessFiledToParse(dati []string) map[string]string {
	campiCentro := make(map[string]string, 6)
	campiCentro["via"] = ""
	campiCentro["tel"] = ""
	campiCentro["cap"] = ""
	campiCentro["prov"] = ""
	campiCentro["loc"] = ""
	campiCentro["email"] = ""

	for _, val := range dati[1:]{
		if(strings.HasPrefix(val,"Via")){
			campiCentro["via"] = val
		} else if(strings.HasPrefix(val,"Tel")){
			campiCentro["tel"] = strings.TrimLeft(val,"Telefono ")
		} else if(strings.HasPrefix(val,"Cap")){
			campiCentro["cap"] = strings.TrimLeft(val,"Cap")
		} else if(strings.HasPrefix(val,"Provincia")){
			campiCentro["prov"] = strings.TrimLeft(val,"Provincia ")
		} else if(strings.HasPrefix(val,"Localita'")){
			campiCentro["loc"] = strings.TrimLeft(val,"Localita' ")
		} else if(strings.HasPrefix(val,"Email")){
			campiCentro["email"] = strings.TrimLeft(val,"Email ")
		}
	}

	return campiCentro
}
