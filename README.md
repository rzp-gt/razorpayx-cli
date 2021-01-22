# razorpayx-cli

A devoper friendly command-line tool to help you test an assortment of different APIs and webhooks on RazorpayX.

**Here are the features that we provide in our beta release:**

- Run custom `post` and `get` requests directly from your terminal with the option to change various parameters
- Run fixtures (which are a set of predefined APIs that help you understand the flow)
- Listen to webhooks triggered from the API (with the abilty to trigger webhooks straight from your terminal as well)
- A consolidated list of all documentation needed while onboarding
- A dynamic checklist that updates as you integrate with new APIs

All this to make your API onboarding onto RazorpayX much smoother.

## Installation

To install, firstly clone the repository:

`git clone https://github.com/rzp-gt/razorpayx-cli.git`

Next, we'll build the program using:

`go build -o /usr/local/bin/RazorpayX cmd/razorpayx/main.go`
