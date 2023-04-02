package main

import (
        "bufio"
        "fmt"
        "net/url"
        "os"

        "golang.org/x/net/publicsuffix"
)

func main() {
        scanner := bufio.NewScanner(os.Stdin)

        for scanner.Scan() {
                rawurl := scanner.Text()
                domain, err := extractDomain(rawurl)
                if err != nil {
                        fmt.Printf("Error: %s\n", err)
                } else {
                        fmt.Printf("%s\n",domain)
                }
        }

        if err := scanner.Err(); err != nil {
                fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
        }
}

func extractDomain(rawurl string) (string, error) {
        parsedURL, err := url.Parse(rawurl)
        if err != nil {
                return "", err
        }

        host := parsedURL.Hostname()
        domain, err := publicsuffix.EffectiveTLDPlusOne(host)
        if err != nil {
                return "", err
        }

        return domain, nil
}
