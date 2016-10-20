/*
Package rest wraps httpClient to provide a platform for additional services.
*/
package rest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

// Configuration arguments for the rest package.  The caller must
// create and supply this struct when calling rest.New()
type Config struct {
	Server  string            // REST server, in form http[s]://hostname:port (required)
	Headers map[string]string // Headers for requests issued by this client
}

type Client struct {
	httpClient *http.Client
	server     string
	Headers    map[string]string
}

// New creates a new REST Client
func New(config Config) (c *Client) {
	c = new(Client)
	c.server = config.Server
	c.httpClient = &http.Client{}
	c.Headers = config.Headers
	if c.Headers == nil {
		c.Headers = map[string]string{}
	}
	return
}

/*
NewRequest creates a new HTTP request
Parameters:
   method	HTTP method
   path		URL path
   params	Map of strings to be added to the URL
			(e.g. { "foo":  "bar" } -> ?foo=bar)
   body		Pointer to data structure converted to JSON as the body
   resp		Data structure returned by request
*/
func (c *Client) NewRequest(method string, path string, params map[string]string, body, resp interface{}) (e error) {

	// Create URI for request

	endpoint, e := url.Parse(c.server + path)
	if e != nil {
		return
	}

	if endpoint.Path == "" {
		endpoint.Path = "/"
	}

	if len(endpoint.Scheme) == 0 {
		endpoint.Scheme = "http"
	}

	// Build URL query from token in client config and any parms passed to this call

	parms := endpoint.Query()

	if params != nil {
		for key, value := range params {
			parms.Add(key, value)
		}
	}
	if len(parms) > 0 {
		endpoint.RawQuery = parms.Encode()
	}

	// Create HTTP request, adding body if supplied by caller

	//log.Printf("%s %s", method, endpoint.String())
	var req *http.Request
	if body == nil {
		req, e = http.NewRequest(method, endpoint.String(), nil)
		if e != nil {
			return
		}
		req.ContentLength = 0
	} else {
		var payload []byte
		payload, e = json.Marshal(body)
		if e != nil {
			return
		}
		log.Printf("Payload: %v", string(payload))
		req, e = http.NewRequest(method, endpoint.String(), bytes.NewReader(payload))
		if e != nil {
			return
		}
		req.ContentLength = int64(len(payload))
		req.Header.Set("Content-Length", strconv.Itoa(len(payload)))
		if len(c.Headers["Content-Type"]) == 0 {
			req.Header.Set("Content-Type", "application/json")
		}
	}
	for header, value := range c.Headers {
		//log.Printf("Setting header %s: %s", header, value)
		req.Header.Set(header, value)
	}

	req.ProtoAtLeast(1, 1)
	req.Close = true
	if len(req.Header) > 0 {
		//log.Printf("Headers: %v", req.Header)
	}

	// Make the request and decode the response

	r, e := c.httpClient.Do(req)
	if e != nil {
		return e
	}
	defer r.Body.Close()
	//log.Printf("Request status code %d (%s)", r.StatusCode, r.Status)
	respBody, e := ioutil.ReadAll(r.Body)
	if e != nil {
		return
	}

	// Handle error conditions and return response
	//  text associated with the status code

	if r.StatusCode > 299 {
		errorText := fmt.Sprintf("Error %d", r.StatusCode)
		httpErrorText := http.StatusText(r.StatusCode)
		if len(httpErrorText) > 0 {
			errorText += " " + httpErrorText
		}
		e = errors.New(errorText)
	} else if resp != nil {
		if _, ok := resp.(*[]byte); ok {
			// Caller wants a byte array - return raw response body
			*resp.(*[]byte) = make([]byte, len(respBody))
			copy(*resp.(*[]byte), respBody)
		} else if e = json.Unmarshal(respBody, resp); e != nil {
			//log.Printf("Error decoding JSON response to REST request: %s", e.Error())
			if len(respBody) == 0 {
				//log.Printf("Response body empty")
			} else {
				//log.Printf("Response body:\n%s", respBody)
			}
		}
	}

	return
}

// Delete is a convenience function for the DELETE method
func (c *Client) Delete(path string, params map[string]string, body interface{}, resp interface{}) error {
	return c.NewRequest("DELETE", path, params, body, resp)
}

// Get is a convenience function for the GET method
func (c *Client) Get(path string, params map[string]string, body interface{}, resp interface{}) error {
	return c.NewRequest("GET", path, params, body, resp)
}

// Post is a convenience function for the POST method
func (c *Client) Post(path string, params map[string]string, body interface{}, resp interface{}) error {
	return c.NewRequest("POST", path, params, body, resp)
}

// Put is a convenience function for the PUT method
func (c *Client) Put(path string, params map[string]string, body interface{}, resp interface{}) error {
	return c.NewRequest("PUT", path, params, body, resp)
}

// EncodeWithPP writes to ResponseWriter with optional pretty printing.
// The caller requests pretty printing with pp=yes in the URL
func EncodeWithPP(data interface{}, w http.ResponseWriter, r *http.Request) (e error) {
	b, e := MarshalWithPP(data, r)
	if e == nil {
		w.Write(b)
	}
	return
}

// MarshalWithPP marshalls data to JSON with optional pretty printing.
// The caller requests pretty printing with pp=yes in the URL
func MarshalWithPP(data interface{}, r *http.Request) ([]byte, error) {
	pp := r.URL.Query()["pp"]
	if len(pp) > 0 && (pp[0] == "true" || pp[0] == "yes") {
		return json.MarshalIndent(data, "", "    ")
	}
	return json.Marshal(data)
}
