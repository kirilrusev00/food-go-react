# Food Analyzer
Food Analyzer is a client-server application that gives information about the ingredients of food products by uploading a qr code of the product. It uses [FoodData Central](https://fdc.nal.usda.gov/) - a database for foods and their content by [U.S. Department of Agriculture](https://www.usda.gov/). The information is publicly accessible via the free REST API which is documented [here](https://fdc.nal.usda.gov/api-guide.html).
## Food Analyzer Server
The Food Analyzer Server:
- is written in go 1.17.
- can receive multiple requests from different clients at a time.
- is connected with a QR decoder server through a tcp socket.
- is connected with a local mysql database which is used to store the results from the requests to FoodData Central API.
- can make requests to FoodData Central API for a gtin upc code of a food and receive an HTTP response with status code 200. An example GET request is made to `https://api.nal.usda.gov/fdc/v1/foods/search?generalSearchInput=009800146130&requireAllWords=true&api_key=<API_KEY>` (where `API_KEY` is a valid API key) and the example response has this JSON body:
```json
{
  "foodSearchCriteria": {
    "generalSearchInput": "raffaello treat",
    "pageNumber": 1,
    "requireAllWords": true
  },
  "totalHits": 1,
  "currentPage": 1,
  "totalPages": 1,
  "foods": [
    {
      "fdcId": 415269,
      "description": "RAFFAELLO, ALMOND COCONUT TREAT",
      "dataType": "Branded",
      "gtinUpc": "009800146130",
      "publishedDate": "2019-04-01",
      "brandOwner": "Ferrero U.S.A., Incorporated",
      "ingredients": "VEGETABLE OILS (PALM AND SHEANUT). DRY COCONUT, SUGAR, ALMONDS, SKIM MILK POWDER, WHEY POWDER (MILK), WHEAT FLOUR, NATURAL AND ARTIFICIAL FLAVORS, LECITHIN AS EMULSIFIER (SOY), SALT, SODIUM BICARBONATE AS LEAVENING AGENT.",
      "allHighlightFields": "",
      "score": 247.10071
    }
  ]
}
```
The requests to the REST API require authentication with an API key which can be retrieved after registering [here](https://fdc.nal.usda.gov/api-key-signup.html).
- can receive a POST request with an image of a QR code in its body. Then it saves the image in a temporary file and sends the path to it to the QR decoder server. If the decoding of the image is successful, the server should receive the gtin upc code of the food. Then, it checks first in the local database if a food with this code exists. If not, it sends a request to the FoodData Central API. As a result the client will receive either information about the food or an error if there was some problem.
From the product data we use only its description field `description` (`RAFFAELLO, ALMOND COCONUT TREAT`), its unique id `fdcId` (`415269`). Some products with `data type Branded` have also GTIN or UPC код, `gtinUpc` (`009800146130`).

_Note:_ GTIN, or [Global Trade Item Number and UPC](https://en.wikipedia.org/wiki/Global_Trade_Item_Number), or [Universal Product Code](https://en.wikipedia.org/wiki/Universal_Product_Code), are identificators of products coded in a QR code like the one below:
![UPC QR code](../main/examples/raffaello.png)
To generate QR codes that can be decoded by the server, the [QR Code Generator from the ZXing Project](https://zxing.appspot.com/generator) is used.
## Food Analyzer Client
The client can make HTTP requests to the server after an image to decode is uploaded. Then it is sent in the body of a POST request to the server. If the decoding is successful, the client receives a response status 200 and a response body with information about the food. If there was a problem, the client receives another response status related to what has gone wrong.
## QR Decoder Server
It connects with the Food Analyzer Server via tcp socket. It receives the path to an image, then reads from the file and decodes it using the [gozxing](https://github.com/makiuchi-d/gozxing) library. The result should be a gtin upc code of a food that is sent back to the Food Analyzer Server.
## How to run the applications
Start the local database. For initial setup execute [this file](https://github.com/kirilrusev00/food-go-react/blob/main/backend/pkg/database/query/create.sql).

To start the __Food Analyzer Server__ and the __QR Decoder Server__ navigate to `backend/cmd` and run 
```
go run main.go
```

To start the __Food Analyzer Client__ navigate to `frontend` and run
```
npm start
```

To run the unit tests navigate to `backend/pkg` and run
```
go test ./…
```
or with coverage: 
```
go test -cover ./...
```

To view the documentation of the go module navigate to backend and run
```
godoc -http=:6060
```
