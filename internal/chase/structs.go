package chase

//Request is the struct that retrieves the feedly access token
type Request struct {
	Token string
}

//TRAMRequest is the structure that stores the necessary data to be sent to TRAM
type TRAMRequest struct {
	index string //insert_report
	title string
	url   string
}

//FeedlyRequest is the struct for building the request body for a feedly API call
type FeedlyRequest struct {
	url   string
	token string
}

//FeedlyResponse is the struct for unmarshalling a feedly API call
type FeedlyResponse struct {
	FeedTitle string  `json:"title"`
	Items     []Blogs `json:"items"`
}

//Blogs is used in Response to store individual Blog objects
type Blogs struct {
	OriginID string `json:"originId"`
	Title    string `json:"title"`
}

//Feeds is used to unmarshal feeds.json
type Feeds struct {
	Feeds []string `json:"feeds"`
}
