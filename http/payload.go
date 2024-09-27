package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"strconv"

	"howett.net/plist"
)

type Payload interface {
	data() ([]byte, string, error) // Returns data, content-type, and error
}

type XMLPayload struct {
	Content map[string]interface{}
}

type JSONPayload struct {
	Content any
}

type URLPayload struct {
	Content map[string]interface{}
}

type FormPayload struct {
	Fields map[string]interface{}
	Files  map[string]string // map field names to file paths
}

func (p *XMLPayload) data() ([]byte, string, error) {
	buffer := new(bytes.Buffer)

	// Encode the content as XML
	err := plist.NewEncoder(buffer).Encode(p.Content)
	if err != nil {
		return nil, "", fmt.Errorf("failed to encode XML: %w", err)
	}

	// Return the XML data and content type
	return buffer.Bytes(), "application/xml", nil
}

func (p *JSONPayload) data() ([]byte, string, error) {
	buffer := new(bytes.Buffer)

	// Encode the content as JSON
	err := json.NewEncoder(buffer).Encode(p.Content)
	if err != nil {
		return nil, "", fmt.Errorf("failed to encode JSON: %w", err)
	}

	// Return the JSON data and content type
	return buffer.Bytes(), "application/json", nil
}

func (p *URLPayload) data() ([]byte, string, error) {
	params := url.Values{}

	// Encode key-value pairs as URL-encoded form
	for key, val := range p.Content {
		switch t := val.(type) {
		case string:
			params.Add(key, val.(string))
		case int:
			params.Add(key, strconv.Itoa(val.(int)))
		default:
			return nil, "", fmt.Errorf("value type is not supported (%s)", t)
		}
	}

	// Return the URL-encoded data and content type
	return []byte(params.Encode()), "application/x-www-form-urlencoded", nil
}

func (p *FormPayload) data() ([]byte, string, error) {
	body := new(bytes.Buffer)

	// Create a new multipart writer with a boundary
	writer := multipart.NewWriter(body)

	// Add form fields
	for key, val := range p.Fields {
		err := writer.WriteField(key, fmt.Sprintf("%v", val))
		if err != nil {
			return nil, "", fmt.Errorf("failed to write field %s: %w", key, err)
		}
	}

	// Add files to the form
	for fieldName, filePath := range p.Files {
		file, err := os.Open(filePath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to open file %s: %w", filePath, err)
		}
		defer file.Close()

		part, err := writer.CreateFormFile(fieldName, filePath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to create form file %s: %w", fieldName, err)
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, "", fmt.Errorf("failed to copy file %s: %w", filePath, err)
		}
	}

	err := writer.Close()
	if err != nil {
		return nil, "", fmt.Errorf("failed to close writer: %w", err)
	}

	// Return the multipart form data and content type (including boundary)
	return body.Bytes(), writer.FormDataContentType(), nil
}
