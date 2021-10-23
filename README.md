# find-me-a-car

## Build
```bash
docker build -t yunussandikci/findmeacar:0.6 .
docker push yunussandikci/findmeacar:0.6
```

## Usage
```bash
docker run -e "TELEGRAM_TOKEN=YOUR_TOKEN" \
    -e "TELEGRAM_CHAT_ID=YOUR_CHAT_ID" \
    -e "DELAY_RANDOM=10" \
    -e "DELAY_CONSTANT=50" \
    -e "BRAND=Mercedes" \
    -e "MODEL=C 180" \
    -e "MINIMUM_YEAR=2015" \
    -e "MAXIMUM_MILEAGE=100000" \
    -e "SELLER_DOMAINS=hasistanbul,koluman,mengerler,otomol,otomolizmir,kolumantarsus,westmotorluaraclar,mercedeshta,mengerlerankara,kolumanankara,yilmazlarotomativas,kolumanmotorluaraclar,birollarotomotiv,mengerlerbursa,gulsoyotomotiv,hastalyaotomotiv,mengerleregemer,gelecekotomotiv,batiotokozyatagi,mengerleradana,gulsoyanadolu,mercedesbenzyigitleras,hasmerotomotivankara,hasmerotomotiv,merkay"\
    -d yunussandikci/findmeacar:0.6
```
