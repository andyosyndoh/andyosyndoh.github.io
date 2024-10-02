package utils

import (
	"reflect"
	"testing"
)

func TestGetsingleartist(t *testing.T) {
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		want    Artists
		wantErr bool
	}{
		{name: "Queen", args: args{ID: 1}, want: Artists{ID: 1, ImageURL: "https://groupietrackers.herokuapp.com/api/images/queen.jpeg", Name: "Queen", Members: []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"}, CreationDate: 1970, FirstAlbum: "14-12-1973", LocationsURL: "https://groupietrackers.herokuapp.com/api/locations/1", ConcertDatesURL: "https://groupietrackers.herokuapp.com/api/dates/1", RelationsURL: "https://groupietrackers.herokuapp.com/api/relation/1"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Getsingleartist(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Getsingleartist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Getsingleartist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetLocations(t *testing.T) {
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		want    Location
		wantErr bool
	}{
		{
			name: "Queen",
			args: args{
				ID: 1,
			},
			want: Location{
				ID: 1,
				Locations: []string{
					"north_carolina-usa",
					"georgia-usa",
					"los_angeles-usa",
					"saitama-japan",
					"osaka-japan",
					"nagoya-japan",
					"penrose-new_zealand",
					"dunedin-new_zealand",
				},
				Dates: "https://groupietrackers.herokuapp.com/api/dates/1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLocations(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLocations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDates(t *testing.T) {
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		want    Date
		wantErr bool
	}{
		{
			name: "Queen",
			args: args{ID: 1},
			want: Date{
				ID: 1,
				Dates: []string{
					"*23-08-2019",
					"*22-08-2019",
					"*20-08-2019",
					"*26-01-2020",
					"*28-01-2020",
					"*30-01-2019",
					"*07-02-2020",
					"*10-02-2020",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDates(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRelation(t *testing.T) {
	type args struct {
		ID int
	}
	tests := []struct {
		name    string
		args    args
		want    ArtistDetails
		wantErr bool
	}{
		{
			name: "Queen",
			args: args{ID: 1},
			want: ArtistDetails{
				ID: 1,
				DatesLocations: map[string][]string{
					"dunedin-new_zealand": {"10-02-2020"},
					"georgia-usa":         {"22-08-2019"},
					"los_angeles-usa":     {"20-08-2019"},
					"nagoya-japan":        {"30-01-2019"},
					"north_carolina-usa":  {"23-08-2019"},
					"osaka-japan":         {"28-01-2020"},
					"penrose-new_zealand": {"07-02-2020"},
					"saitama-japan":       {"26-01-2020"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRelation(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRelation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRelation() = %v, want %v", got, tt.want)
			}
		})
	}
}
