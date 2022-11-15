package handler

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/cateiru/cateiru.com/ent"
	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/cateiru/cateiru.com/src/db"
	"github.com/labstack/echo/v4"
)

type BioResponse struct {
	Biography ent.Biography
	Location  ent.Location
}

// Get my bio
func (h *Handler) BioHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	bio, err := h.DB.Client.Biography.
		Query().
		Where(biography.UserID(h.User.ID)).
		First(ctx)
	if err != nil {
		return err
	}

	location, err := h.DB.Client.Location.Get(ctx, bio.LocationID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, BioResponse{
		*bio,
		*location,
	})
}

// Set a new bio
//
// require form is
// - is_public: boolean
// - location_id: uint32
// - position: string
// - position_ja: string
// - join_date:  type date
// - leave_date: Optional RFC3339 type date
func (h *Handler) CreateBioHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	isPublic, err := strconv.ParseBool(e.FormValue("is_public"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: is_public")
	}
	locationId, err := strconv.Atoi(e.FormValue("location_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: location_id")
	}

	position := e.FormValue("position")
	positionJa := e.FormValue("position_ja")

	joinDateStr := e.FormValue("join_date")
	leaveDateStr := e.FormValue("leave_date") // optional

	// form checks
	if position == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: position")
	}
	if positionJa == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: position_ja")
	}
	if joinDateStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: join_date")
	}

	if err := existLocationId(ctx, h.DB, uint32(locationId)); err != nil {
		return err
	}

	// RFC3339
	joinDate, err := time.Parse(time.RFC3339, joinDateStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: join_date")
	}

	bio := h.DB.Client.Biography.Create()

	if leaveDateStr != "" {
		// RFC3339
		leaveDate, err := time.Parse(time.RFC3339, leaveDateStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: leave_date")
		}
		bio = bio.SetLeave(leaveDate)
	}

	bio = bio.SetUserID(h.User.ID)

	bio = bio.SetIsPublic(isPublic)
	bio = bio.SetLocationID(uint32(locationId))
	bio = bio.SetPosition(position)
	bio = bio.SetPositionJa(positionJa)
	bio = bio.SetJoin(joinDate)

	bioDB, err := bio.Save(ctx)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusCreated, bioDB)
}

// Set a new bio
//
// changeable form is
// - is_public: boolean
// - location_id: uint32
// - position: string
// - position_ja: string
// - join_date: RFC3339 type date
// - leave_date: RFC3339 type date
func (h *Handler) UpdateBioHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	bioId, err := strconv.Atoi(e.FormValue("bio_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: bio_id")
	}

	changes := false
	u := h.DB.Client.Biography.
		Update().
		Where(
			biography.And(
				biography.ID(uint32(bioId)),
				biography.UserID(h.User.ID),
			),
		)

	isPublicStr := e.FormValue("is_public")
	if isPublicStr != "" {
		changes = true
		isPublic, err := strconv.ParseBool(isPublicStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: is_public")
		}
		u = u.SetIsPublic(isPublic)
	}

	locationIdStr := e.FormValue("location_id")
	if locationIdStr != "" {
		changes = true
		locationId, err := strconv.Atoi(locationIdStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: location_id")
		}
		if err := existLocationId(ctx, h.DB, uint32(locationId)); err != nil {
			return err
		}
		u = u.SetLocationID(uint32(locationId))
	}

	position := e.FormValue("position")
	if position != "" {
		changes = true
		u = u.SetPosition(position)
	}

	positionJa := e.FormValue("position_ja")
	if positionJa != "" {
		changes = true
		u = u.SetPositionJa(positionJa)
	}

	joinDateStr := e.FormValue("join_date")
	if joinDateStr != "" {
		changes = true
		joinDate, err := time.Parse(time.RFC3339, joinDateStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: join_date")
		}
		u = u.SetJoin(joinDate)
	}

	leaveDateStr := e.FormValue("leave_date")
	if leaveDateStr != "" {
		changes = true
		leaveDate, err := time.Parse(time.RFC3339, leaveDateStr)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: leave_date")
		}
		u = u.SetLeave(leaveDate)
	} else {
		u = u.ClearLeave()
	}

	if !changes {
		return echo.NewHTTPError(http.StatusBadRequest, "no changes")
	}

	if err := u.Exec(ctx); err != nil {
		return err
	}

	bioDB, err := h.DB.Client.Biography.
		Query().
		Where(
			biography.And(
				biography.ID(uint32(bioId)),
				biography.UserID(h.User.ID),
			),
		).
		First(ctx)
	if err != nil {
		return err
	}
	location, err := h.DB.Client.Location.Get(ctx, bioDB.LocationID)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, BioResponse{
		Biography: *bioDB,
		Location:  *location,
	})
}

func (h *Handler) DeleteBioHandler(e echo.Context) error {
	ctx := context.Background()

	if err := h.Session(ctx, e); err != nil {
		return err
	}

	idStr := e.QueryParam("bio_id")
	if idStr == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: bio_id")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid form: bio_id")
	}

	_, err = h.DB.Client.Biography.
		Delete().
		Where(biography.And(
			biography.ID(uint32(id)),
			biography.UserID(h.User.ID),
		)).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// check location id is exist in locations db
func existLocationId(ctx context.Context, db *db.DB, locationId uint32) error {
	exist, err := db.Client.Location.
		Query().
		Where(location.ID(locationId)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if !exist {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid forms: not exist location")
	}
	return nil
}
