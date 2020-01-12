package main

import (
	"encoding/json"
	"fmt"
)

type MxRecord struct {
	Value    string `json:"value"`
	Ttl      int64  `json:"ttl"`
	Priority int64  `json:"priority"`
	HostName string `json:"hostName"`
}

type CnameRecord struct {
	AliasHost     string `json:"aliasHost"`
	CanonicalHost string `json:"canonicalHost"`
}

type DkimTxtRecord struct {
	Domainname string `json:"domainname"`
	Value      string `json:"value"`
	Ttl        int64  `json:"ttl"`
}

type SpfTxtRecord struct {
	Domainname string `json:"domainname"`
	Value      string
	Ttl        int64
}

type Data struct {
	MxRecords     []MxRecord
	CnameRecords  []CnameRecord
	DkimTxtRecord DkimTxtRecord
	SpfTxtRecord  SpfTxtRecord
	LoginURL      string `json:"loginUrl"`
}

type Response struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

type apiR struct {
	Response Response `json:"response"`
}

func main() {
	body := `
	  {"response": {
		  "status": "SUCCESS",
		  "data": {
			"mxRecords": [
			  {
				"value": "us2.mx3.mailhostbox.com.",
				"ttl": 1,
				"priority": 100,
				"hostName": "@"
			  },
			  {
				"value": "us2.mx1.mailhostbox.com.",
				"ttl": 1,
				"priority": 100,
				"hostName": "@"
			  },
			  {
				"value": "us2.mx2.mailhostbox.com.",
				"ttl": 1,
				"priority": 100,
				"hostName": "@"
			  }
			],
			"cnameRecords": [
			  {
				"aliasHost": "pop.a.co.uk.",
				"canonicalHost": "us2.pop.mailhostbox.com."
			  },
			  {
				"aliasHost": "webmail.a.co.uk.",
				"canonicalHost": "us2.webmail.mailhostbox.com."
			  },
			  {
				"aliasHost": "smtp.a.co.uk.",
				"canonicalHost": "us2.smtp.mailhostbox.com."
			  },
			  {
				"aliasHost": "imap.a.co.uk.",
				"canonicalHost": "us2.imap.mailhostbox.com."
			  }
			],
			"dkimTxtRecord": {
			  "domainname": "20a19._domainkey.a.co.uk",
			  "value": "\"v=DKIM1; g=*; k=rsa; p=DkfbhO8Oyy0E1WyUWwIDAQAB\"",
			  "ttl": 1
			},
			"spfTxtRecord": {
			  "domainname": "a.co.uk",
			  "value": "\"v=spf1 redirect=_spf.mailhostbox.com\"",
			  "ttl": 1
			},
			"loginUrl": "us2.cp.mailhostbox.com"
		  }
	}}`

	var api apiR
	err := json.Unmarshal([]byte(body), &api)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%v", api)
	fmt.Println(api.Response.Status)

	mxrecords := api.Response.Data.MxRecords

	for i, v := range mxrecords {
		fmt.Println(i, v.HostName, v.Priority, v.Ttl, v.Value)
	}
}
