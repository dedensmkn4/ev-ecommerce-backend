package httpkit

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

// DumpHTTPRequest dump request without header, just body
func DumpHTTPRequest(req *http.Request) []byte {
	var err error
	save := req.Body
	if req.Body == nil {
		req.Body = nil
		return nil
	}

	save, req.Body, err = drainBody(req.Body)
	if err != nil {
		return nil
	}

	chunked := len(req.TransferEncoding) > 0 && req.TransferEncoding[0] == "chunked"
	var b bytes.Buffer

	if req.Body != nil {
		var dest io.Writer = &b
		if chunked {
			dest = httputil.NewChunkedWriter(dest)
		}
		_, err = io.Copy(dest, req.Body)
		if chunked {
			dest.(io.Closer).Close()
			io.WriteString(&b, "")
		}
	}

	req.Body = save
	if err != nil {
		return nil
	}

	return b.Bytes()
}

func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == http.NoBody {
		return http.NoBody, http.NoBody, nil
	}

	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}

	if err = b.Close(); err != nil {
		return nil, b, err
	}

	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}