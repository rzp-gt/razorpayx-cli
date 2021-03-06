package requests

import (
	"bufio"
	"context"
	"fmt"
	"github.com/rzp-gt/razorpayx-cli/internal/ansi"
	"github.com/rzp-gt/razorpayx-cli/internal/client"
	"github.com/rzp-gt/razorpayx-cli/internal/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// RequestParameters captures the structure of the parameters that can be sent to RazorpayX
type RequestParameters struct {
	data        []string
	idempotency string
}

// AppendData appends data to the request parameters.
func (r *RequestParameters) AppendData(data []string) {
	r.data = append(r.data, data...)
}

// Base encapsulates the required information needed to make requests to the API
type Base struct {
	Cmd *cobra.Command

	Method  string
	Profile *config.Profile

	Parameters RequestParameters

	// SuppressOutput is used by `trigger` to hide output
	SuppressOutput bool

	DarkStyle bool

	APIBaseURL string

	Livemode bool

	autoConfirm bool
	showHeaders bool
}

var confirmationCommands = map[string]bool{http.MethodDelete: true}

// RunRequestsCmd is the interface exposed for the CLI to run network requests through
func (rb *Base) RunRequestsCmd(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("this command only supports one argument. Run with the --help flag to see usage and examples")
	}

	if len(args) == 0 {
		return nil
	}

	confirmed, err := rb.confirmCommand()
	if err != nil {
		return err
	} else if !confirmed {
		fmt.Println("Exiting without execution. User did not confirm the command.")
		return nil
	}

	apiKey, err := rb.Profile.GetAPIKey(rb.Livemode)
	if err != nil {
		return err
	}
	apiSecret, err := rb.Profile.GetAPISecret(rb.Livemode)
	if err != nil {
		return err
	}

	path, err := createOrNormalizePath(args[0])
	if err != nil {
		return err
	}

	_, err = rb.MakeRequest(apiKey, apiSecret, path, &rb.Parameters, false)

	return err
}

// InitFlags initialize shared flags for all requests commands
func (rb *Base) InitFlags() {

	rb.Cmd.Flags().StringArrayVarP(&rb.Parameters.data, "data", "d", []string{}, "Data for the API request")
	rb.Cmd.Flags().BoolVar(&rb.Livemode, "live", false, "Make a live request (default: test)")

	// Hidden configuration flags, useful for dev/debugging
	rb.Cmd.Flags().StringVar(&rb.APIBaseURL, "api-base", client.DefaultAPIBaseURL, "Sets the API base URL")
	rb.Cmd.Flags().MarkHidden("api-base") // #nosec G104
}

// MakeRequest will make a request to the RazorpayX API with the specific variables given to it
func (rb *Base) MakeRequest(apiKey, apiSecret, path string, params *RequestParameters, errOnStatus bool) ([]byte, error) {
	parsedBaseURL, err := url.Parse(rb.APIBaseURL)
	if err != nil {
		return []byte{}, err
	}

	client := &client.Client{
		BaseURL:   parsedBaseURL,
		APIKey:    apiKey,
		APISecret: apiSecret,
		Verbose:   rb.showHeaders,
	}

	data, err := rb.buildDataForRequest(params)
	fmt.Println("Request body")
	fmt.Println(ansi.ColorizeJSON(data, false,os.Stdout ))

	if err != nil {
		return []byte{}, err
	}

	configureReq := func(req *http.Request) {
		rb.setIdempotencyHeader(req, params)
	}

	resp, err := client.PerformRequest(context.TODO(), rb.Method, path, data, configureReq)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if !rb.SuppressOutput {
		if err != nil {
			return []byte{}, err
		}

		result := ansi.ColorizeJSON(string(body), rb.DarkStyle, os.Stdout)
		fmt.Print(result)
	}

	if errOnStatus && resp.StatusCode >= 300 {
		return nil, fmt.Errorf("Request failed, status=%d, body=%s", resp.StatusCode, string(body))
	}
	return body, nil
}

// Confirm calls the confirmCommand() function, triggering the confirmation process
func (rb *Base) Confirm() (bool, error) {
	return rb.confirmCommand()
}

// Note: We converted to using two arrays to track keys and values, with our own
// implementation of Go's url.Values Encode function due to our query parameters being
// order sensitive for API requests involving arrays like `items` for `/v1/orders`.
// Go's url.Values uses Go's map, which jumbles the key ordering, and their Encode
// implementation sorts keys by alphabetical order, but this doesn't work for us since
// some API endpoints have required parameter ordering. Yes, this is hacky, but it works.
func (rb *Base) buildDataForRequest(params *RequestParameters) (string, error) {
	keys := []string{}
	values := []string{}

	if len(params.data) > 0 {
		for _, datum := range params.data {
			splitDatum := strings.SplitN(datum, "=", 2)

			if len(splitDatum) < 2 {
				return "", fmt.Errorf("Invalid data argument: %s", datum)
			}

			keys = append(keys, splitDatum[0])
			values = append(values, splitDatum[1])
		}
	}

	return encode(keys, values), nil
}

// encode creates a url encoded string with the request parameters
func encode(keys []string, values []string) string {
	var buf strings.Builder

	for i := range keys {
		key := keys[i]
		value := values[i]

		keyEscaped := url.QueryEscape(key)

		// Don't use strict form encoding by changing the square bracket
		// control characters back to their literals. This is fine by the
		// server, and makes these parameter strings easier to read.
		keyEscaped = strings.ReplaceAll(keyEscaped, "%5B", "[")
		keyEscaped = strings.ReplaceAll(keyEscaped, "%5D", "]")

		if buf.Len() > 0 {
			buf.WriteByte('&')
		}

		buf.WriteString(keyEscaped)
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(value))
	}

	return buf.String()
}

func (rb *Base) setIdempotencyHeader(request *http.Request, params *RequestParameters) {
	if params.idempotency != "" {
		request.Header.Set("X-Idempotency-Key", params.idempotency)

		if rb.Method == http.MethodGet || rb.Method == http.MethodDelete {
			warning := fmt.Sprintf(
				"Warning: sending an idempotency key with a %s request has no effect and should be avoided, as %s requests are idempotent by definition.",
				rb.Method,
				rb.Method,
			)
			fmt.Println(warning)
		}
	}
}

func (rb *Base) confirmCommand() (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	return rb.getUserConfirmation(reader)
}

func (rb *Base) getUserConfirmation(reader *bufio.Reader) (bool, error) {
	if _, needsConfirmation := confirmationCommands[rb.Method]; needsConfirmation && !rb.autoConfirm {
		confirmationPrompt := fmt.Sprintf("Are you sure you want to perform the command: %s?\nEnter 'yes' to confirm: ", rb.Method)
		fmt.Print(confirmationPrompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}

		// remove whitespace from either side of the input, as ReadString returns with \n at the end
		input = strings.ToLower(strings.Trim(input, " \r\n"))

		return strings.Compare(input, "yes") == 0, nil
	}

	// Always confirm the command if it does not require explicit user confirmation
	return true, nil
}

func createOrNormalizePath(arg string) (string, error) {
	fmt.Println(arg)
	if idRegex.Match([]byte(arg)) {
		matches := idRegex.FindStringSubmatch(arg)

		if path, ok := idURLMap[matches[1]]; ok {
			return path + arg, nil
		}

		return "", fmt.Errorf("Unrecognized object id: %s", arg)
	}

	return normalizePath(arg), nil
}

func normalizePath(path string) string {
	if strings.HasPrefix(path, "/v1/") {
		return path
	}

	if strings.HasPrefix(path, "v1/") {
		return "/" + path
	}

	if strings.HasPrefix(path, "/") {
		return "/v1" + path
	}

	return "/v1/" + path
}
