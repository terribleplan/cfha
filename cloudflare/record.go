package cloudflare

import (
  "fmt"
)

//~Return the actual error when possible

type RecordsResponse struct {
  Response struct {
    Recs struct {
      Records []Record `json:"objs"`
    } `json:"recs"`
  } `json:"response"`
  Result  string `json:"result"`
  Message string `json:"msg"`
}

//-Removed wildcard option
func (r *RecordsResponse) FindRecordByName(name string) ([]Record, error) {
  if r.Result == "error" {
    return nil, fmt.Errorf("API Error: %s", r.Message)
  }

  objs := r.Response.Recs.Records

  var recs []Record

  //-Removed errror on no results
  //-Removed wildcard checking
  //~Checking against FullName, not just name

  for _, v := range objs {
    if v.FullName == name {
      recs = append(recs, v)
    }
  }

  return recs, nil
}

type RecordResponse struct {
  Response struct {
    Rec struct {
      Record Record `json:"obj"`
    } `json:"rec"`
  } `json:"response"`
  Result  string `json:"result"`
  Message string `json:"msg"`
}

func (r *RecordResponse) GetRecord() (*Record, error) {
  if r.Result == "error" {
    return nil, fmt.Errorf("API Error: %s", r.Message)
  }

  return &r.Response.Rec.Record, nil
}

// Record is used to represent a retrieved Record. All properties
// are set as strings.
type Record struct {
  Id       string `json:"rec_id"`
  Domain   string `json:"zone_name"`
  Name     string `json:"display_name"`
  FullName string `json:"name"`
  Value    string `json:"content"`
  Type     string `json:"type"`
  Priority string `json:"prio"`
  Ttl      string `json:"ttl"`
}

// CreateRecord contains the request parameters to create a new
// record.
type CreateRecord struct {
  Type     string
  Name     string
  Content  string
  Ttl      string
  Priority string
}

// CreateRecord creates a record from the parameters specified and
// returns an error if it fails. If no error and the name is returned,
// the Record was succesfully created.
func (c *Client) CreateRecord(domain string, opts *CreateRecord) (*Record, error) {
  // Make the request parameters
  params := make(map[string]string)
  //-No option checking, we set them all ourselves
  params["z"] = domain
  params["type"] = opts.Type
  params["name"] = opts.Name
  params["content"] = opts.Content
  params["prio"] = opts.Priority
  params["ttl"] = "1"

  req, err := c.NewRequest(params, "POST", "rec_new")
  if err != nil {
    return nil, err
  }

  resp, err := checkResp(c.Http.Do(req))

  if err != nil {
    return nil, err
  }

  recordResp := new(RecordResponse)

  err = decodeBody(resp, &recordResp)

  if err != nil {
    return nil, err
  }
  record, err := recordResp.GetRecord()
  if err != nil {
    return nil, err
  }

  // The request was successful
  return record, nil
}

// DestroyRecord destroys a record by the ID specified and
// returns an error if it fails. If no error is returned,
// the Record was succesfully destroyed.
func (c *Client) DestroyRecord(domain string, id string) error {
  params := make(map[string]string)
  params["z"] = domain
  params["id"] = id

  req, err := c.NewRequest(params, "POST", "rec_delete")
  if err != nil {
    return err
  }

  resp, err := checkResp(c.Http.Do(req))

  if err != nil {
    return err
  }

  recordResp := new(RecordResponse)

  err = decodeBody(resp, &recordResp)

  if err != nil {
    return err
  }
  _, err = recordResp.GetRecord()
  if err != nil {
    return err
  }

  // The request was successful
  return nil
}

func (c *Client) RetrieveRecordsByName(domain string, name string) ([]Record, error) {
  params := make(map[string]string)
  params["z"] = domain

  req, err := c.NewRequest(params, "GET", "rec_load_all")

  if err != nil {
    return nil, err
  }

  resp, err := checkResp(c.Http.Do(req))
  if err != nil {
    return nil, err
  }

  records := new(RecordsResponse)

  err = decodeBody(resp, records)

  if err != nil {
    return nil, err
  }

  record, err := records.FindRecordByName(name)
  if err != nil {
    return nil, err
  }

  // The request was successful
  return record, nil
}
