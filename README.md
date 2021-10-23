# find-me-a-car

//todo: readme

## Usage Example
```bash
docker run -e "TELEGRAM_TOKEN=YOUR_TOKEN" \
    -e "TELEGRAM_CHAT_ID=YOUR_CHAT_ID" \
    -e "DELAY_RANDOM=10" \
    -e "DELAY_CONSTANT=50" \
    -e "BRAND=.*(Mercedes).*" \
    -e "MODEL=.*(C 180|E 180).*" \
    -e "MINIMUM_YEAR=2015" \
    -e "MAXIMUM_MILEAGE=100000" \
    -e "MAXIMUM_PRICE=500000" \
    -e "SELLER_DOMAINS=koluman,mengerler"\
    -d yunussandikci/findmeacar:0.0.1
```
