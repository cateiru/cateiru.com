package test

import (
	"context"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/cateiru/cateiru.com/src/db"
)

const UNDEFINED_LOCATION_ID uint32 = 1000000000

type TestBio struct {
	UserId     uint32
	IsPublic   bool
	LocationId uint32
	Position   string
	PositionJa string
	Join       time.Time
	Leave      time.Time

	LocationType      location.Type
	LocationName      string
	LocationNameJa    string
	LocationAddress   string
	LocationAddressJa string

	Biography *ent.Biography
	Location  *ent.Location
}

func NewTestBio(userId uint32) (*TestBio, error) {
	position, err := MakeRandomStr(10)
	if err != nil {
		return nil, err
	}
	locationName, err := MakeRandomStr(10)
	if err != nil {
		return nil, err
	}
	locationAddress, err := MakeRandomStr(10)
	if err != nil {
		return nil, err
	}

	join := time.Now()
	join.AddDate(-5, 0, 0)
	leave := time.Now()

	return &TestBio{
		UserId:     userId,
		IsPublic:   false,
		LocationId: UNDEFINED_LOCATION_ID,
		Position:   position,
		PositionJa: position,

		Join:  join,
		Leave: leave,

		LocationType:      "univ",
		LocationName:      locationName,
		LocationNameJa:    locationName,
		LocationAddress:   locationAddress,
		LocationAddressJa: locationAddress,

		Biography: nil,
		Location:  nil,
	}, nil
}

func (c *TestBio) CreateDB(ctx context.Context, db *db.DB) (*ent.Biography, error) {
	if c.Location != nil {
		c.LocationId = c.Location.ID
	} else if c.LocationId == UNDEFINED_LOCATION_ID {
		_, err := c.CreateLocationDB(ctx, db)
		if err != nil {
			return nil, err
		}
	}

	bio, err := db.Client.Biography.
		Create().
		SetUserID(c.UserId).
		SetIsPublic(c.IsPublic).
		SetLocationID(c.LocationId).
		SetPosition(c.Position).
		SetPositionJa(c.PositionJa).
		SetJoin(c.Join).
		SetLeave(c.Leave).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	c.Biography = bio

	return bio, nil
}

func (c *TestBio) CreateLocationDB(ctx context.Context, db *db.DB) (*ent.Location, error) {
	location, err := db.Client.Location.
		Create().
		SetType(c.LocationType).
		SetName(c.LocationName).
		SetNameJa(c.LocationNameJa).
		SetAddress(c.LocationAddress).
		SetAddressJa(c.LocationAddressJa).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	c.LocationId = location.ID
	c.Location = location

	return location, nil
}
