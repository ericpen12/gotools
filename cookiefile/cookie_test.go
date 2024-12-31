package cookiefile

import (
	"net/url"
	"testing"
)

func TestNew(t *testing.T) {
	s := New("./aa.cookie")
	//s.Set("ddd")
	ret, err := s.Get()
	t.Log(err)
	t.Log(ret)
}

func TestName(t *testing.T) {
	u, _ := url.Parse("https://uc.creditcard.ecitic.com/citiccard/ucweb/entry.do?channel=INNER_WEB_BSVC&rtnUrl=https%3A%2F%2Fe.creditcard.ecitic.com%2Fciticcard%2Febank-ocp%2Febankpc%2Fmyaccount.html%3Ffunc%3Dmainpage%26&loginSource=ucweb")
	t.Log(u.Host)

}
