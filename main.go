package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	//Connect to the database
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://mongodbcontainer"))

	//Defer the disconnection from the database
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	// Generate the collection name for this run prepended with "vehicles_"
	collection_name := "vehicles_" + strconv.FormatInt(time.Now().Unix(), 10)

	//Prepare to start iterating 100 vehicles at a time
	offset := 0

	//Iterate until we get an empty list of vehicles
	for {

		//Collect the data
		doc, err := collect_data(offset)
		if err != nil {
			panic(err)
		}

		log.Println("Collecting data for offset " + strconv.Itoa(offset) + " (" + strconv.Itoa(len(doc.AllListings)) + " vehicles))")

		//Break if the AllListings property is empty or missing
		if doc.AllListings == nil || len(doc.AllListings) == 0 {
			break
		}

		//Store the data
		store_data(client, context.Background(), collection_name, doc)

		//Increment the offset
		offset += 100

		//Wait 10 seconds before the next iteration
		time.Sleep(3 * time.Second)
	}
}

func store_data(client *mongo.Client, ctx context.Context, collection_name string, doc ListingsResponse) {
	//Iterate over the AllListings property
	for i := 0; i < len(doc.AllListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.AllListings[i].AdID}, bson.M{"$set": doc.AllListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.AllListings[i].AdID)
		}
	}

	//Iterate over the ProvincialPriorityListings property
	for i := 0; i < len(doc.ProvincialPriorityListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.ProvincialPriorityListings[i].AdID}, bson.M{"$set": doc.ProvincialPriorityListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id: %d", doc.ProvincialPriorityListings[i].AdID)
		}
	}

	//Iterate over the TopSpotListings property
	for i := 0; i < len(doc.TopSpotListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.TopSpotListings[i].AdID}, bson.M{"$set": doc.TopSpotListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.TopSpotListings[i].AdID)
		}
	}

	//Iterate over the PriorityListingsTop property
	for i := 0; i < len(doc.PriorityListingsTop); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.PriorityListingsTop[i].AdID}, bson.M{"$set": doc.PriorityListingsTop[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.PriorityListingsTop[i].AdID)
		}
	}

	//Iterate over the FeaturedListings property
	for i := 0; i < len(doc.FeaturedListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.FeaturedListings[i].AdID}, bson.M{"$set": doc.FeaturedListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.FeaturedListings[i].AdID)
		}
	}

	//Iterate over the PriorityListingsBottom property
	for i := 0; i < len(doc.PriorityListingsBottom); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.PriorityListingsBottom[i].AdID}, bson.M{"$set": doc.PriorityListingsBottom[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.PriorityListingsBottom[i].AdID)
		}
	}

	//Iterate over the PriorityListings property
	for i := 0; i < len(doc.PriorityListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.PriorityListings[i].AdID}, bson.M{"$set": doc.PriorityListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.PriorityListings[i].AdID)
		}
	}

	//Iterate over the OrganicListings property
	for i := 0; i < len(doc.OrganicListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.OrganicListings[i].AdID}, bson.M{"$set": doc.OrganicListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.OrganicListings[i].AdID)
		}
	}

	//Iterate over the OrganicProvincialPrioriyListings property
	for i := 0; i < len(doc.OrganicProvincialPrioriyListings); i++ {
		update_result, err := client.Database("autotrader").Collection(collection_name).UpdateOne(ctx, bson.M{"id": doc.OrganicProvincialPrioriyListings[i].AdID}, bson.M{"$set": doc.OrganicProvincialPrioriyListings[i]}, options.Update().SetUpsert(true))
		if update_result == nil && err != nil {
			log.Printf("Failed to update document with id %d", doc.OrganicProvincialPrioriyListings[i].AdID)
		}
	}

}

func collect_data(offset int) (ListingsResponse, error){
	url := fmt.Sprintf("https://www.autotrader.ca/cars/qc/quebec/?rcp=100&rcs=%s&prx=-1&prv=Quebec&loc=H9W%25205C3&hprc=True&wcp=True&iosp=True&sts=New-Used&inMarket=basicSearch", strconv.Itoa(offset))
	method := "GET"

	// Create request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return ListingsResponse{}, err
	}
	
	req.Header.Add("authority", "www.autotrader.ca")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("accept-language", "en-US,en;q=0.6")
	req.Header.Add("allowmvt", "true")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cookie", "atOptUser=5ba36b51-7476-4284-a1ee-ce0e57540475; visid_incap_820541=bS1vwk2nR8+gac6Fw+Ah0eWGzWMAAAAAQUIPAAAAAADwc+Wn3D/RcEd7NrPWWE+J; cbnr=1; visid_incap_2401829=CFJF4r4pQQSJ98HWLoTXXeaGzWMAAAAAQUIPAAAAAADqJUKsYj9jy6AiUw7/d/lI; incap_ses_532_2401829=dfljGUjQfw7VoTBU+wtiB+aGzWMAAAAAUkLY063Y3NBWiuKOUwyoxg==; _gcl_au=1.1.634441330.1674413800; __GTMADBLOCKER__=no; {E7ABF06F-D6A6-4c25-9558-3932D3B8A04D}=; PageSize=15; pCode=H9W5C3; searchState={\"isUniqueSearch\":false,\"make\":null,\"model\":null}; incap_ses_1291_820541=qyI1GJqnOnvs3bb8Uo7qEcWpzWMAAAAAcejOfIzLhl1l0X0I1gI9tA==; srchLocation=%7B%22Location%22:%7B%22Address%22:null,%22City%22:%22Beaconsfield%22,%22Latitude%22:45.422332763671875,%22Longitude%22:-73.870269775390625,%22Province%22:%22QC%22,%22PostalCode%22:%22H9W%205C3%22,%22Type%22:%22%22%7D,%22UnparsedAddress%22:%22H9W%205C3%22%7D; nlbi_820541_1646237=aXahRhhlxhUWCcx9cA+uiQAAAACIamlhCdzld+CTVxOjiDcv; gtm_inmarket_search=true; searchBreadcrumbs=%7B%22srpBreadcrumb%22%3A%5B%7B%22Text%22%3A%22Cars%2C%20Trucks%20%26%20SUVs%22%2C%22Url%22%3A%22%2Fcars%2F%3Fsrt%3D35%26loc%3DH9W%25205C3%26hprc%3DTrue%26wcp%3DTrue%22%7D%5D%2C%22isFromSRP%22%3Afalse%2C%22neighbouringIds%22%3Anull%7D; lastsrpurl=/cars/?rcp=15&rcs=0&srt=35&prx=-1&loc=H9W%205C3&hprc=True&wcp=True&inMarket=advancedSearch; atOptUser=5ba36b51-7476-4284-a1ee-ce0e57540475; incap_ses_1291_820541=eLhCJrOeCUzjbJD8Uo7qEReKzWMAAAAA87n3kqdgpiqUJ+Nr5aHTow==; nlbi_820541_1646235=Ve2KaRSXGSrs2CpqcA+uiQAAAADglEMG8ikZasTw5vT1cv2Z; nlbi_820541_1646237=E/UEb7MtyyUWgpLNcA+uiQAAAACo0pHBjQq5lJ95C5wKIhN2; searchBreadcrumbs=%7B%22srpBreadcrumb%22%3A%5B%7B%22Text%22%3A%22Cars%2C%20Trucks%20%26%20SUVs%22%2C%22Url%22%3A%22%2Fcars%2F%3Fsrt%3D35%26loc%3DH9W%25205C3%26hprc%3DTrue%26wcp%3DTrue%22%7D%5D%2C%22isFromSRP%22%3Afalse%2C%22neighbouringIds%22%3Anull%7D; srchLocation=%7b%22Location%22%3a%7b%22Type%22%3a%22Physical%22%2c%22Aliases%22%3a%5b%5d%2c%22Name%22%3a%22Toronto%22%2c%22Address%22%3anull%2c%22City%22%3a%22Toronto%22%2c%22Province%22%3a%22ON%22%2c%22PostalCode%22%3anull%2c%22Latitude%22%3a43.70011%2c%22Longitude%22%3a-79.4163%7d%2c%22UnparsedAddress%22%3a%22Toronto%2c+ON%22%7d; uag=037C12A993A2E849642EB988EFD26247558B9AB57D2B964CB028ABDF2668B02C; visid_incap_820541=+X00ccCjRrOmCK6RS1C6nheKzWMAAAAAQUIPAAAAAACqdLSSy8vNJ3KRoYhN7iyo; .ASPXANONYMOUS=QMwdYydl2QEkAAAAMDE3YjUxOWItMzE3Yi00NDdkLWFjMDMtOGQ3ZGVmZDAxNWNlcBaxtIi9NEDfFa8rODmfV9dh6qo1; ASP.NET_SessionId=wbrfvgqfaoo4u2z2ddeoafnf; InternalSignInComplete=False; InternalSignInCompleteNew=False; _vpc2=; culture=en-ca; {E7ABF06F-D6A6-4c25-9558-3932D3B8A04D}=atOptUser=5ba36b51-7476-4284-a1ee-ce0e57540475&searchBreadcrumbs=%257B%2522srpBreadcrumb%2522%253A%255B%257B%2522Text%2522%253A%2522Cars%252C%2520Trucks%2520%2526%2520SUVs%2522%252C%2522Url%2522%253A%2522%252Fcars%252F%253Fsrt%253D35%2526prv%253DOntario%2526loc%253DToronto%25252C%252520ON%2526hprc%253DTrue%2526wcp%253DTrue%2522%257D%252C%257B%2522Text%2522%253A%2522Ontario%2522%252C%2522Url%2522%253A%2522%252Fcars%252Fon%252F%253Fsrt%253D35%2526prv%253DOntario%2526loc%253DToronto%25252C%252520ON%2526hprc%253DTrue%2526wcp%253DTrue%2522%257D%252C%257B%2522Text%2522%253A%2522Toronto%2522%252C%2522Url%2522%253A%2522%252Fcars%252Fon%252Ftoronto%252F%253Fsrt%253D35%2526prv%253DOntario%2526loc%253DToronto%25252C%252520ON%2526hprc%253DTrue%2526wcp%253DTrue%2522%257D%255D%252C%2522isFromSRP%2522%253Afalse%252C%2522neighbouringIds%2522%253Anull%257D&cty=Toronto&prv=Ontario")
	req.Header.Add("isajax", "true")
	req.Header.Add("ms", "1")
	req.Header.Add("origin", "https://www.autotrader.ca")
	req.Header.Add("referer", "https://www.autotrader.ca/cars/?rcp=15&rcs=0&srt=35&prx=-1&loc=H9W%205C3&hprc=True&wcp=True&inMarket=advancedSearch")
	req.Header.Add("sec-ch-ua", "\"Not_A Brand\";v=\"99\", \"Brave\";v=\"109\", \"Chromium\";v=\"109\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ListingsResponse{}, err
	}

	// Unmarshall the response into SearchResponse json struct
	var searchResponse ListingsResponse
	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return ListingsResponse{}, err
	}

	return searchResponse, nil
}