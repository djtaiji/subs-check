package config

const DefaultConfigTemplate = `
# Whether to show progress
print-progress: false
# Port
port: 8080
# Check items
check-items:
  - openai
  - youtube
  - netflix
  - disney
# Number of concurrent threads
concurrent: 200

# Check interval (minutes)
check-interval: 30

# Timeout (milliseconds) (maximum delay of nodes)
timeout: 5000

# Quality level, the higher the level, the higher the stability; too high a level may result in too few nodes being filtered out
quality-level: 1

# Speed test address (Note: concurrent * node speed < maximum network speed, otherwise the speed test result is inaccurate)
# Try not to use Speedtest or download links provided by Cloudflare, as many nodes block speed test websites
# It is recommended to use files uploaded to Cloudflare R2
# If you do not need to set a speed test address, please set speed-test-url to ""
speed-test-url: https://github.com/AaronFeng753/Waifu2x-Extension-GUI/releases/download/v2.21.12/Waifu2x-Extension-GUI-v2.21.12-Portable.7z

# Minimum speed test result to discard (KB/s)
min-speed: 1024

# Download test time (s) (related to the size of the download link, default maximum test 10s)
download-timeout: 10

# mihomo api url (automatically update mihomo subscription after testing is complete)
mihomo-api-url: ""

# mihomo api secret
mihomo-api-secret: ""

# Save method
# Currently supported save methods: r2, local, gist, webdav
save-method: local

# webdav
webdav-url: "https://example.com/dav/"
webdav-username: "admin"
webdav-password: "admin"

# gist id
github-gist-id: ""

# github token
github-token: ""

# github api mirror
github-api-mirror: ""

# Address to push speed test results to Worker
worker-url: https://example.worker.dev

# Worker token
worker-token: 1234567890

# Retry count (number of retries after subscription retrieval fails)
sub-urls-retry: 3

# Proxy settings, this proxy is only used to obtain subscription links and upload/save (not used for node testing, etc.)
# Supports http and socks proxy
# If you do not need to set a proxy, please set type to ""
proxy:
  type: "http" # Optional values: http, socks; if you do not need to set a proxy, please set type to ""
  address: "http://localhost:8080" # Proxy address

# Subscription address supports clash/mihomo/v2ray/base64 format subscription links
sub-urls:
  - https://example.com/sub.txt
  - https://example.com/sub2.txt
`
