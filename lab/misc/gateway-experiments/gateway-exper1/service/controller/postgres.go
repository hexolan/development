package controller

import (
	"time"
	"context"
	"strconv"
	"encoding/json"

	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	// "github.com/jackc/pgx/v5/pgtype"

	"null.hexolan.dev/exp/service"
	"null.hexolan.dev/exp/utils/errors"
	pb "null.hexolan.dev/exp/protogen/testingv1"
)

// ---

const (
	itemBaseQuery string = "SELECT id, status, title, created_at FROM items"
)

// ---

type postgresController struct {
	pCl *pgxpool.Pool

	evtC service.EventController
}

func NewPostgresController(pCl *pgxpool.Pool, evtC service.EventController) service.StorageController {
	return postgresController{pCl: pCl, evtC: evtC}
}

// ---

func scanRowToItem(row pgx.Row) (*pb.Item, error) {
	var item pb.Item

	// todo: test enum conversion
	var itemId int64

	var createdAt time.Time
	err := row.Scan(
		&itemId, // todo: resolve scan error - may have to change from bigserial to serial for now?
		&item.Status, 
		&item.Title,
		&createdAt,
	)
	if err != nil {
		log.Error().Err(err).Msg("failed scan")
		return nil, err
	}

	item.Id = strconv.FormatInt(itemId, 10)
	item.CreatedAt = createdAt.Unix()
	
	return &item, nil
}

// ---

func (c postgresController) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	row := c.pCl.QueryRow(ctx, (itemBaseQuery + " WHERE id=$1"), req.GetItemId())
	item, err := scanRowToItem(row)
	if err != nil {
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to scan database row")
	}

	return &pb.GetItemResponse{Data: item}, nil
}

func (c postgresController) GetItems(ctx context.Context, req *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	// todo: implement query method
	/*
	rows, err := c.pCl.Query(ctx, QUERY, QUERYARGS)
	if err != nil {
		// todo: wrap in service error
		return nil, err
	}

	items := []*pb.Item{}
	for rows.Next() {
		item, err := scanRowToItem(rows)
		if err != nil {
			// todo: handling
			// return service error - something went wrong?
			continue
		}
		
		items = append(items, item)
	}

	if rows.Err() != nil {
		// todo: wrap with service error
		return nil, rows.Err()
	}

	return items, nil
	*/
	
	return &pb.GetItemsResponse{}, nil
}

func (c postgresController) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.UpdateItemResponse, error) {
	// todo: needs improving
	// really messy just for the exper atm

	// ======
	// messy patch data stuff
	patchData := goqu.Record{"updated_at": goqu.L("timezone('utc', now())")}

	// todo: ignoring bad keys - at validation stage
	marshalled, _ := json.Marshal(req.Data)
	_ = json.Unmarshal(marshalled, &patchData)
	// ======

	// actual update stuff
	statement, args, err := goqu.Dialect("postgres").Update("panels").Prepared(true).Set(patchData).Where(goqu.C("id").Eq(req.Data.Id)).ToSQL()
	if err != nil {
		return nil, errors.WrapServiceError(errors.ErrCodeUnknown, "failed to prepare update query", err)
	}

	result, err := c.pCl.Exec(ctx, statement, args...)
	if err != nil {
		// todo: actual error checking
		// not unique?
		return nil, errors.NewServiceError(errors.ErrCodeService, "failed to update item")
	}
	// ======
	
	// ensure a row was actually updated
	rows_affected := result.RowsAffected()
	if rows_affected != 1 {
		return nil, errors.NewServiceError(errors.ErrCodeNotFound, "item not found")
	}
	// ======

	// get the updated item
	// todo: improve by just merging patch locally
	itemReq, err := c.GetItem(ctx, &pb.GetItemRequest{ItemId: req.Data.GetId()})
	if err != nil {
		return nil, err
	}
	item := itemReq.Data

	// dispatch event
	c.evtC.DispatchUpdatedEvent(item)

	// return response
	return &pb.UpdateItemResponse{Data: item}, nil
}

func (c postgresController) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.DeleteItemResponse, error) {
	// database stuff
	result, err := c.pCl.Exec(ctx, "DELETE FROM items WHERE id=$1", req.ItemId)
	if err != nil {
		// todo: catching connection failures, etc
		return nil, errors.WrapServiceError(errors.ErrCodeUnknown, "failed to delete item", err)
	}

	// ensure rows were affected
	rows_affected := result.RowsAffected()
	if rows_affected != 1 {
		return nil, errors.NewServiceError(errors.ErrCodeNotFound, "item not found")
	}

	// dispatch event
	// todo: update to dispatch full item instead of just id
	c.evtC.DispatchDeletedEvent(&pb.Item{Id: req.ItemId})

	// return response
	return &pb.DeleteItemResponse{}, nil
}

func (c postgresController) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.CreateItemResponse, error)  {
	// database stuff
	var id int64
	err := c.pCl.QueryRow(
		ctx,
		"INSERT INTO items (status, title) VALUES ($1, $2) RETURNING id",
		req.Data.Status,
		req.Data.Title,
	).Scan(
		&id,
	)
	if err != nil {
		// todo: catching errors
		// e.g. conn error or not unique
		return nil, errors.WrapServiceError(errors.ErrCodeUnknown, "failed to create item", err)
	}

	// assemble item
	item := &pb.Item{
		Id: strconv.FormatInt(id, 10),
		Status: req.Data.Status,
		Title: req.Data.Title,
		CreatedAt: time.Now().Unix(),
	}

	// dispatch event
	c.evtC.DispatchCreatedEvent(item)

	// return response
	return &pb.CreateItemResponse{Data: item}, nil
}