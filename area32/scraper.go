package area32

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"net/http/cookiejar"
	"strings"
	"time"
)

type Area32User struct {
	Id         string
	ImageUrl   string
	Fullname   string
	ExpireDate time.Time
}

func (u *Area32User) IsExpired() bool {
	return time.Now().After(u.ExpireDate)
}

type ScraperApi struct {
	client *resty.Client
}

func NewAPI() *ScraperApi {
	cookieJar, _ := cookiejar.New(nil)
	client := resty.New().SetCookieJar(cookieJar).SetDoNotParseResponse(true)
	return &ScraperApi{client: client}
}

func (api *ScraperApi) DoLoginAndRetrieveMain(email, password string) (*Area32User, error) {
	resp, err := api.client.R().
		Get("https://www.cloud32.it/Associazioni/utenti/login?codass=170734")
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.RawBody())

	if err != nil {
		print("2")
		return nil, err
	}

	var token string
	doc.Find("input").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "_token" {
			token, _ = s.Attr("value")
		}
	})

	formData := map[string]string{
		"email":    email,
		"password": password,
		"_token":   token,
	}
	_, err = api.client.R().
		SetFormData(formData).
		Post("https://www.cloud32.it/Associazioni/utenti/login")
	if err != nil {
		return nil, err
	}

	resp, err = api.client.R().
		Get("https://www.cloud32.it/Associazioni/utenti/home")
	if err != nil {
		print("4")
		return nil, err
	}

	doc, err = goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	imageUrl := retrieveImageUrl(doc)
	userId := retrieveID(doc)
	expireDate := retrieveExpireDate(doc)
	fullName := retrieveFullName(doc)
	if userId == "" {
		return nil, errors.New("Invalid credentials")
	}
	return &Area32User{
		Id:         userId,
		ImageUrl:   imageUrl,
		ExpireDate: expireDate,
		Fullname:   fullName,
	}, nil
}

func retrieveImageUrl(doc *goquery.Document) string {
	foundImage := false
	imageUrl := ""
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if alt, _ := s.Attr("alt"); alt == "Foto" {
			if altImage, _ := s.Attr("src"); altImage != "" {
				foundImage = true
				imageUrl, _ = s.Attr("src")
			}
		}
	})

	if !foundImage {
		return ""
	}

	return "https://www.cloud32.it" + imageUrl
}

func retrieveID(doc *goquery.Document) string {
	foundID := false
	id := ""
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		if class, _ := s.Attr("class"); class == "col-sm-12" {
			if strings.Contains(s.Text(), "Tessera:") {
				s.Find("label").Each(func(i int, s *goquery.Selection) {
					id = s.Text()
					foundID = true
				})
			}
		}
	})

	if !foundID {
		return ""
	}

	return id
}

func retrieveExpireDate(doc *goquery.Document) time.Time {
	expireDate := time.Now().Add(time.Hour * 24 * 365 * 10)
	doc.Find("div").Each(func(i int, s *goquery.Selection) {
		if class, _ := s.Attr("class"); class == "col-sm-12" {
			if strings.Contains(s.Text(), "Scadenza:") {
				s.Find("label").Each(func(i int, s *goquery.Selection) {
					loc, _ := time.LoadLocation("Europe/Rome")
					expireDate, _ = time.ParseInLocation("02/01/2006", s.Text(), loc)
				})
			}
		}
	})
	return expireDate
}

func retrieveFullName(doc *goquery.Document) string {
	fullName := ""
	doc.Find("span").Each(func(i int, s *goquery.Selection) {
		if class, _ := s.Attr("class"); class == "itemless nomeprofilo" {
			fullName = s.Text()
		}
	})
	return fullName
}
