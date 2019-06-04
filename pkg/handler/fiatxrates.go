package handler

import (
	"encoding/json"
	"fmt"

	"github.com/horizontalsystems/xrates-kit/pkg/config"
	dtutil "github.com/horizontalsystems/xrates-kit/pkg/util/datetime"
	httputil "github.com/horizontalsystems/xrates-kit/pkg/util/http"
)

// -------------------------------
var FiatXRatesHandler *FiatXRates

// -------------------------------

//FiatXRates handler object
type FiatXRates struct {
	Conf *config.FiatXRatesConfig
}

type FiatXRatesResponse struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
	Time  int64              `json:"time"`
}

func init() {
	FiatXRatesHandler = &FiatXRates{&config.Get().FiatXRates}
}

func (fXRates *FiatXRates) GetLatestXRatesAsJSON(sCcy string, tCcy string, exchange string) (string, error) {

	respStr, err := httputil.DoGet(
		6, fXRates.Conf.APIURL, "/latest?base="+sCcy+"&symbols="+tCcy, "")

	if err != nil {

	}

	return respStr, err
}

func (fXRates *FiatXRates) GetXRatesAsJSON(sCcy string, tCcy string,
	exchange string, epochSec *int64) (string, error) {

	respStr, err := httputil.DoGet(
		6, fXRates.Conf.APIURL, dtutil.EpochToStr("", epochSec), "base="+sCcy+"&symbols="+tCcy)

	if err != nil {

	}

	return respStr, err
}

func (fXRates *FiatXRates) GetXRates(sCcy string, tCcy string,
	exchange string, epochSec *int64) (FiatXRatesResponse, error) {

	respStr, err := httputil.DoGet(
		6, fXRates.Conf.APIURL, dtutil.EpochToStr("", epochSec), "base="+sCcy+"&symbols="+tCcy)

	if err != nil {

	}
	res := FiatXRatesResponse{}
	json.Unmarshal([]byte(respStr), &res)

	fmt.Println("FXrate", res.Rates)

	return res, err
}
