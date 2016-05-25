package myradio

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Officership struct {
	OfficerId   uint   `json:"officerid,string"`
	OfficerName string `json:"officer_name"`
	TeamId      uint   `json:"teamid,string"`
	FromDateRaw string `json:"from_date,omitempty"`
	FromDate    time.Time
	TillDateRaw string `json:"till_date,omitempty"`
	TillDate    time.Time
}

type Photo struct {
	PhotoId      uint   `json:"photoid"`
	DateAddedRaw string `json:"date_added"`
	DateAdded    time.Time
	Format       string `json:"format"`
	Owner        uint   `json:"owner"`
	Url          string `json:"url"`
}

func (s *Session) GetUserBio(id int) (bio string, err error) {
	data, err := s.apiRequest(fmt.Sprintf("/user/%d/bio/", id))
	if err != nil {
		return
	}
	if data == nil {
		err = errors.New("No bio set")
		return
	}
	err = json.Unmarshal(*data, &bio)
	return
}

func (s *Session) GetUserName(id int) (name string, err error) {
	data, err := s.apiRequest(fmt.Sprintf("/user/%d/name/", id))
	if err != nil {
		return
	}
	err = json.Unmarshal(*data, &name)
	return
}

func (s *Session) GetUserProfilePhoto(id int) (profilephoto Photo, err error) {
	data, err := s.apiRequest(fmt.Sprintf("/user/%d/profilephoto/", id))
	if err != nil {
		return
	}
	if data == nil {
		err = errors.New("No profile picture set")
		return
	}
	err = json.Unmarshal(*data, &profilephoto)
	if err != nil {
		return
	}
	profilephoto.DateAdded, err = time.Parse("02/01/2006 15:04", profilephoto.DateAddedRaw)
	return
}

func (s *Session) GetUserOfficerships(id int) (officerships []Officership, err error) {
	data, err := s.apiRequest(fmt.Sprintf("/user/%d/officerships/", id))
	if err != nil {
		return
	}
	err = json.Unmarshal(*data, &officerships)
	if err != nil {
		return
	}
	for k, v := range officerships {
		if officerships[k].FromDateRaw != "" {
			officerships[k].FromDate, err = time.Parse("2006-01-02", v.FromDateRaw)
			if err != nil {
				return
			}
		}
		if officerships[k].TillDateRaw != "" {
			officerships[k].TillDate, err = time.Parse("2006-01-02", v.FromDateRaw)
			if err != nil {
				return
			}
		}
	}
	return
}

func (s *Session) GetUserShowCredits(id int) (shows []ShowMeta, err error) {
	data, err := s.apiRequest(fmt.Sprintf("/user/%d/shows/", id))
	if err != nil {
		return
	}
	err = json.Unmarshal(*data, &shows)
	return
}
