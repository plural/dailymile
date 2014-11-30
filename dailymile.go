package dailymile

import (
	"fmt"

	"github.com/twpayne/gopolyline/polyline"
)

// Move these types to a daily_mile.go file
type AuthResponse struct {
	Access_Token string
	Token_Type   string
	Error        string
}

type MeResponse struct {
	Goal        string `json:"goal"`
	Location    string `json:"location"`
	TimeZone    string `json:"time_zone"`
	Url         string `json:"url"`
	DisplayName string `json:"display_name"`
	Username    string `json:"username"`
	PhotoUrl    string `json:"photo_url"`
}

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	PhotoUrl    string `json:"photo_url"`
	Url         string `json:"url"`
}

type Comment struct {
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Like struct {
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type Geo struct {
	Type        string  `json:"type,omitempty"`
	Coordinates LatLong `json:"coordinates,omitempty"`
}

type LatLong struct {
	Lat  float64 `json:",string"`
	Long float64 `json:",string"`
}

func (p *LatLong) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[\"%v\",\"%v\"]", p.Lat, p.Long)), nil
}

func (p *LatLong) UnmarshalJSON(data []byte) (err error) {
	_, err = fmt.Sscanf(string(data), "[\"%v\",\"%v\"]", &p.Lat, &p.Long)
	return
}

type Location struct {
	Name string `json:"name"`
}

type Distance struct {
	Value float64 `json:"value,omitempty"`
	Units string  `json:"units,omitempty"`
}

type Workout struct {
	ActivityType string   `json:"activity_type,omitempty"`
	Felt         string   `json:"felt,omitempty"`
	Duration     int64    `json:"duration,omitempty"`
	Distance     Distance `json:"distance,omitempty"`
	Title        string   `json:"title,omitempty"`
}

type Entry struct {
	Id       int       `json:"id"`
	Url      string    `json:"url"`
	At       string    `json:"at"`
	Message  string    `json:"message"`
	Comments []Comment `json:"comments"`
	Likes    []Like    `json:"likes"`
	Geo      Geo       `json:"geo,omitempty"`
	Location Location  `json:"location,omitempty"`
	User     User      `json:"user"`
	Workout  Workout   `json:"workout,omitempty"`
}

type EntryStream struct {
	Entries []Entry `json:"entries"`
}

type RouteStream struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Id             int      `json:"id"`
	ActivityType   string   `json:"activity_type"`
	Name           string   `json:"name"`
	EncodedSamples string   `json:"encoded_samples"`
	Distance       Distance `json:"distance,omitempty"`
	Geo            Geo      `json:"geo,omitempty"`
}

func (r Route) GetRoutePoints() []LatLong {
	var routePoints []LatLong
	decoded_points, err := polyline.Decode(r.EncodedSamples, 2)
	if err == nil {
		routePoints = make([]LatLong, len(decoded_points)/2)
		for i := 0; i < len(decoded_points); i += 2 {
			routePoints[i/2] = LatLong{Lat: decoded_points[i], Long: decoded_points[i+1]}
		}
	}
	return routePoints
}
