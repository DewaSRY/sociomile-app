# sociomile-app

## Author 
Name : Dewa Surya Ariesta (Full stack Engineer)

# About 

Applikasi ini dibuat unutk mendata pesan tulisan yang di kirim oleh pelanggan kepada sebuah usaha (organisasi)

Cara kerja dari applikasi ini, jikalau ada sebuah pesan yang di terima oleh sebuah contact usaha, system akan mecatat pesan tersebt pada system 

kemudian untuk setiap pesan yang tampil pada dashboard pemilik usaha, pemilik usaha bisa menujuk staff/sales untuk menjawab pesan text dari customer

Jikalau ada masalah yang tidak terselesaikan melalui komunikasi staff bisa di buatkan sebuah tiket


## stack applikasi 

|  |                  |
|----------|----------|
| backend      | go     |
| client      | vue     |
| database      | MySql     |

## Cara menyiapkan database 


```bash
## ! pastikan docker terinstall di local computer 

## jalankan perintah 

docker-compose up -d

```
## Cara menjalankan applikasi backend 

```bash
## move to the backend directory 
cd backend 

## download the dependency
go mod tidy

## crete env variable 
cp .env.example .env

## run the seeder 
go run ./cmd/seeders

## run the application 
go run ./cmd/server

## if on linux or wsl run 
make run-dev
```

## Cara menjalankan applikasi front end  

```bash
## move to the backend directory 
cd client 

## download the dependency
npm install 

## crete env variable 
cp .env.example .env

## run the application 
npm run dev

```

doc link
[Application core](https://github.com/DewaSRY/sociomile-app/blob/main/doc/application-core.md)
[Entity ](https://github.com/DewaSRY/sociomile-app/blob/main/doc/application-entity.md)
[Business logic ](https://github.com/DewaSRY/sociomile-app/blob/main/doc/business-logic.md)