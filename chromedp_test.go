package JinJi_Service
//
//import (
//	"context"
//	"github.com/chromedp/cdproto/cdp"
//	"os"
//
//	"github.com/chromedp/chromedp"
//	"github.com/kr/pretty"
//	"github.com/rs/zerolog/log"
//
//	"testing"
//	"time"
//	"github.com/chromedp/cdproto/network"
//)
//
//type Cookie struct {
//	Name 				string
//	Value 				string
//	Domain				string
//	Path				string
//	Expires				int
//	Size				string
//	HTTPOnly			bool
//	Secure				bool
//	Session				string
//	SameSite			string
//	Priority			string
//	SourceScheme		string
//	SourcePort			int64
//	PartitionKey		string
//	PartitionKeyOpaque	string
//}
//
//
//func TestStart(t *testing.T) {
//	var err error
//	var opts = []chromedp.ExecAllocatorOption{
//		chromedp.NoFirstRun,
//		chromedp.NoDefaultBrowserCheck,
//		chromedp.NoSandbox,
//		chromedp.UserAgent("Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36"),
//	}
//
//	//var cookies_file *os.File
//	//if cookies_file, err = os.OpenFile("./cookies.txt", os.O_WRONLY | os.O_APPEND, 0755); err != nil {pretty.Println(err)}
//	var cookies_bytes []byte
//	if cookies_bytes, err = os.ReadFile("./cookies.txt",); err != nil {pretty.Println(err)}
//
//	var cookies []*Cookie
//	for _, cookie_str := range string(c)
//
//
//
//	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
//	defer cancel()
//
//	// create context
//	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
//	ctx, cancel = context.WithTimeout(ctx, time.Minute * 2)
//	defer cancel()
//
//	if err = chromedp.Run(ctx,
//		chromedp.Navigate("https://www.youtube.com/"),
//		chromedp.Sleep(time.Minute * 1),
//		chromedp.ActionFunc(func(ctx context.Context) error {
//			//var cookies []*network.Cookie
//			//if cookies, err = network.GetCookies().Do(ctx); err != nil {pretty.Println(err)}
//
//			//for i, cookie := range cookies {
//			//	log.Info().Msgf("set chrome cookie %d: %+v", i, cookie)
//			//	if _, err = cookies_file.Write([]byte(fmt.Sprintf("Name:%s Value:%s \n", cookie.Name, cookie.Value))); err != nil {pretty.Println(err)}
//			//}
//			//return nil
//
//			for i, cookie := range cookies {
//				exp := cdp.TimeSinceEpoch(time.Unix(int64(cookie.Expires), int64(1000)))
//				if err = network.SetCookie(cookie.Name, cookie.Value).
//					WithExpires(&exp).
//					WithPath(cookie.Path).
//					WithDomain(cookie.Domain).
//					WithHTTPOnly(cookie.HTTPOnly).
//					WithSameSite(network.CookieSameSite(cookie.SameSite)).
//					WithPriority(network.CookiePriority(cookie.Priority)).
//					//WithSameParty(cookie.SameParty).
//					WithSecure(cookie.Secure).
//					WithSourceScheme(network.CookieSourceScheme(cookie.SourceScheme)).
//					WithSourcePort(cookie.SourcePort).
//					Do(ctx); err != nil {pretty.Println(err)}
//				log.Info().Msgf("set chrome cookie %d: %+v", i, cookie)
//			}
//			return nil
//		}),
//	); err != nil {pretty.Println(err)}
//}
