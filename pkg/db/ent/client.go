// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/migrate"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponallocated"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/coupondiscount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponfixamount"
	"github.com/NpoolPlatform/inspire-manager/pkg/db/ent/couponspecialoffer"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CouponAllocated is the client for interacting with the CouponAllocated builders.
	CouponAllocated *CouponAllocatedClient
	// CouponDiscount is the client for interacting with the CouponDiscount builders.
	CouponDiscount *CouponDiscountClient
	// CouponFixAmount is the client for interacting with the CouponFixAmount builders.
	CouponFixAmount *CouponFixAmountClient
	// CouponSpecialOffer is the client for interacting with the CouponSpecialOffer builders.
	CouponSpecialOffer *CouponSpecialOfferClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CouponAllocated = NewCouponAllocatedClient(c.config)
	c.CouponDiscount = NewCouponDiscountClient(c.config)
	c.CouponFixAmount = NewCouponFixAmountClient(c.config)
	c.CouponSpecialOffer = NewCouponSpecialOfferClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		CouponAllocated:    NewCouponAllocatedClient(cfg),
		CouponDiscount:     NewCouponDiscountClient(cfg),
		CouponFixAmount:    NewCouponFixAmountClient(cfg),
		CouponSpecialOffer: NewCouponSpecialOfferClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                ctx,
		config:             cfg,
		CouponAllocated:    NewCouponAllocatedClient(cfg),
		CouponDiscount:     NewCouponDiscountClient(cfg),
		CouponFixAmount:    NewCouponFixAmountClient(cfg),
		CouponSpecialOffer: NewCouponSpecialOfferClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CouponAllocated.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.CouponAllocated.Use(hooks...)
	c.CouponDiscount.Use(hooks...)
	c.CouponFixAmount.Use(hooks...)
	c.CouponSpecialOffer.Use(hooks...)
}

// CouponAllocatedClient is a client for the CouponAllocated schema.
type CouponAllocatedClient struct {
	config
}

// NewCouponAllocatedClient returns a client for the CouponAllocated from the given config.
func NewCouponAllocatedClient(c config) *CouponAllocatedClient {
	return &CouponAllocatedClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `couponallocated.Hooks(f(g(h())))`.
func (c *CouponAllocatedClient) Use(hooks ...Hook) {
	c.hooks.CouponAllocated = append(c.hooks.CouponAllocated, hooks...)
}

// Create returns a builder for creating a CouponAllocated entity.
func (c *CouponAllocatedClient) Create() *CouponAllocatedCreate {
	mutation := newCouponAllocatedMutation(c.config, OpCreate)
	return &CouponAllocatedCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CouponAllocated entities.
func (c *CouponAllocatedClient) CreateBulk(builders ...*CouponAllocatedCreate) *CouponAllocatedCreateBulk {
	return &CouponAllocatedCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CouponAllocated.
func (c *CouponAllocatedClient) Update() *CouponAllocatedUpdate {
	mutation := newCouponAllocatedMutation(c.config, OpUpdate)
	return &CouponAllocatedUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponAllocatedClient) UpdateOne(ca *CouponAllocated) *CouponAllocatedUpdateOne {
	mutation := newCouponAllocatedMutation(c.config, OpUpdateOne, withCouponAllocated(ca))
	return &CouponAllocatedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponAllocatedClient) UpdateOneID(id uuid.UUID) *CouponAllocatedUpdateOne {
	mutation := newCouponAllocatedMutation(c.config, OpUpdateOne, withCouponAllocatedID(id))
	return &CouponAllocatedUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CouponAllocated.
func (c *CouponAllocatedClient) Delete() *CouponAllocatedDelete {
	mutation := newCouponAllocatedMutation(c.config, OpDelete)
	return &CouponAllocatedDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CouponAllocatedClient) DeleteOne(ca *CouponAllocated) *CouponAllocatedDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CouponAllocatedClient) DeleteOneID(id uuid.UUID) *CouponAllocatedDeleteOne {
	builder := c.Delete().Where(couponallocated.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponAllocatedDeleteOne{builder}
}

// Query returns a query builder for CouponAllocated.
func (c *CouponAllocatedClient) Query() *CouponAllocatedQuery {
	return &CouponAllocatedQuery{
		config: c.config,
	}
}

// Get returns a CouponAllocated entity by its id.
func (c *CouponAllocatedClient) Get(ctx context.Context, id uuid.UUID) (*CouponAllocated, error) {
	return c.Query().Where(couponallocated.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponAllocatedClient) GetX(ctx context.Context, id uuid.UUID) *CouponAllocated {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponAllocatedClient) Hooks() []Hook {
	hooks := c.hooks.CouponAllocated
	return append(hooks[:len(hooks):len(hooks)], couponallocated.Hooks[:]...)
}

// CouponDiscountClient is a client for the CouponDiscount schema.
type CouponDiscountClient struct {
	config
}

// NewCouponDiscountClient returns a client for the CouponDiscount from the given config.
func NewCouponDiscountClient(c config) *CouponDiscountClient {
	return &CouponDiscountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coupondiscount.Hooks(f(g(h())))`.
func (c *CouponDiscountClient) Use(hooks ...Hook) {
	c.hooks.CouponDiscount = append(c.hooks.CouponDiscount, hooks...)
}

// Create returns a builder for creating a CouponDiscount entity.
func (c *CouponDiscountClient) Create() *CouponDiscountCreate {
	mutation := newCouponDiscountMutation(c.config, OpCreate)
	return &CouponDiscountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CouponDiscount entities.
func (c *CouponDiscountClient) CreateBulk(builders ...*CouponDiscountCreate) *CouponDiscountCreateBulk {
	return &CouponDiscountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CouponDiscount.
func (c *CouponDiscountClient) Update() *CouponDiscountUpdate {
	mutation := newCouponDiscountMutation(c.config, OpUpdate)
	return &CouponDiscountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponDiscountClient) UpdateOne(cd *CouponDiscount) *CouponDiscountUpdateOne {
	mutation := newCouponDiscountMutation(c.config, OpUpdateOne, withCouponDiscount(cd))
	return &CouponDiscountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponDiscountClient) UpdateOneID(id uuid.UUID) *CouponDiscountUpdateOne {
	mutation := newCouponDiscountMutation(c.config, OpUpdateOne, withCouponDiscountID(id))
	return &CouponDiscountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CouponDiscount.
func (c *CouponDiscountClient) Delete() *CouponDiscountDelete {
	mutation := newCouponDiscountMutation(c.config, OpDelete)
	return &CouponDiscountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CouponDiscountClient) DeleteOne(cd *CouponDiscount) *CouponDiscountDeleteOne {
	return c.DeleteOneID(cd.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CouponDiscountClient) DeleteOneID(id uuid.UUID) *CouponDiscountDeleteOne {
	builder := c.Delete().Where(coupondiscount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponDiscountDeleteOne{builder}
}

// Query returns a query builder for CouponDiscount.
func (c *CouponDiscountClient) Query() *CouponDiscountQuery {
	return &CouponDiscountQuery{
		config: c.config,
	}
}

// Get returns a CouponDiscount entity by its id.
func (c *CouponDiscountClient) Get(ctx context.Context, id uuid.UUID) (*CouponDiscount, error) {
	return c.Query().Where(coupondiscount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponDiscountClient) GetX(ctx context.Context, id uuid.UUID) *CouponDiscount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponDiscountClient) Hooks() []Hook {
	hooks := c.hooks.CouponDiscount
	return append(hooks[:len(hooks):len(hooks)], coupondiscount.Hooks[:]...)
}

// CouponFixAmountClient is a client for the CouponFixAmount schema.
type CouponFixAmountClient struct {
	config
}

// NewCouponFixAmountClient returns a client for the CouponFixAmount from the given config.
func NewCouponFixAmountClient(c config) *CouponFixAmountClient {
	return &CouponFixAmountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `couponfixamount.Hooks(f(g(h())))`.
func (c *CouponFixAmountClient) Use(hooks ...Hook) {
	c.hooks.CouponFixAmount = append(c.hooks.CouponFixAmount, hooks...)
}

// Create returns a builder for creating a CouponFixAmount entity.
func (c *CouponFixAmountClient) Create() *CouponFixAmountCreate {
	mutation := newCouponFixAmountMutation(c.config, OpCreate)
	return &CouponFixAmountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CouponFixAmount entities.
func (c *CouponFixAmountClient) CreateBulk(builders ...*CouponFixAmountCreate) *CouponFixAmountCreateBulk {
	return &CouponFixAmountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CouponFixAmount.
func (c *CouponFixAmountClient) Update() *CouponFixAmountUpdate {
	mutation := newCouponFixAmountMutation(c.config, OpUpdate)
	return &CouponFixAmountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponFixAmountClient) UpdateOne(cfa *CouponFixAmount) *CouponFixAmountUpdateOne {
	mutation := newCouponFixAmountMutation(c.config, OpUpdateOne, withCouponFixAmount(cfa))
	return &CouponFixAmountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponFixAmountClient) UpdateOneID(id uuid.UUID) *CouponFixAmountUpdateOne {
	mutation := newCouponFixAmountMutation(c.config, OpUpdateOne, withCouponFixAmountID(id))
	return &CouponFixAmountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CouponFixAmount.
func (c *CouponFixAmountClient) Delete() *CouponFixAmountDelete {
	mutation := newCouponFixAmountMutation(c.config, OpDelete)
	return &CouponFixAmountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CouponFixAmountClient) DeleteOne(cfa *CouponFixAmount) *CouponFixAmountDeleteOne {
	return c.DeleteOneID(cfa.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CouponFixAmountClient) DeleteOneID(id uuid.UUID) *CouponFixAmountDeleteOne {
	builder := c.Delete().Where(couponfixamount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponFixAmountDeleteOne{builder}
}

// Query returns a query builder for CouponFixAmount.
func (c *CouponFixAmountClient) Query() *CouponFixAmountQuery {
	return &CouponFixAmountQuery{
		config: c.config,
	}
}

// Get returns a CouponFixAmount entity by its id.
func (c *CouponFixAmountClient) Get(ctx context.Context, id uuid.UUID) (*CouponFixAmount, error) {
	return c.Query().Where(couponfixamount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponFixAmountClient) GetX(ctx context.Context, id uuid.UUID) *CouponFixAmount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponFixAmountClient) Hooks() []Hook {
	hooks := c.hooks.CouponFixAmount
	return append(hooks[:len(hooks):len(hooks)], couponfixamount.Hooks[:]...)
}

// CouponSpecialOfferClient is a client for the CouponSpecialOffer schema.
type CouponSpecialOfferClient struct {
	config
}

// NewCouponSpecialOfferClient returns a client for the CouponSpecialOffer from the given config.
func NewCouponSpecialOfferClient(c config) *CouponSpecialOfferClient {
	return &CouponSpecialOfferClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `couponspecialoffer.Hooks(f(g(h())))`.
func (c *CouponSpecialOfferClient) Use(hooks ...Hook) {
	c.hooks.CouponSpecialOffer = append(c.hooks.CouponSpecialOffer, hooks...)
}

// Create returns a builder for creating a CouponSpecialOffer entity.
func (c *CouponSpecialOfferClient) Create() *CouponSpecialOfferCreate {
	mutation := newCouponSpecialOfferMutation(c.config, OpCreate)
	return &CouponSpecialOfferCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CouponSpecialOffer entities.
func (c *CouponSpecialOfferClient) CreateBulk(builders ...*CouponSpecialOfferCreate) *CouponSpecialOfferCreateBulk {
	return &CouponSpecialOfferCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CouponSpecialOffer.
func (c *CouponSpecialOfferClient) Update() *CouponSpecialOfferUpdate {
	mutation := newCouponSpecialOfferMutation(c.config, OpUpdate)
	return &CouponSpecialOfferUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CouponSpecialOfferClient) UpdateOne(cso *CouponSpecialOffer) *CouponSpecialOfferUpdateOne {
	mutation := newCouponSpecialOfferMutation(c.config, OpUpdateOne, withCouponSpecialOffer(cso))
	return &CouponSpecialOfferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CouponSpecialOfferClient) UpdateOneID(id uuid.UUID) *CouponSpecialOfferUpdateOne {
	mutation := newCouponSpecialOfferMutation(c.config, OpUpdateOne, withCouponSpecialOfferID(id))
	return &CouponSpecialOfferUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CouponSpecialOffer.
func (c *CouponSpecialOfferClient) Delete() *CouponSpecialOfferDelete {
	mutation := newCouponSpecialOfferMutation(c.config, OpDelete)
	return &CouponSpecialOfferDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CouponSpecialOfferClient) DeleteOne(cso *CouponSpecialOffer) *CouponSpecialOfferDeleteOne {
	return c.DeleteOneID(cso.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CouponSpecialOfferClient) DeleteOneID(id uuid.UUID) *CouponSpecialOfferDeleteOne {
	builder := c.Delete().Where(couponspecialoffer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CouponSpecialOfferDeleteOne{builder}
}

// Query returns a query builder for CouponSpecialOffer.
func (c *CouponSpecialOfferClient) Query() *CouponSpecialOfferQuery {
	return &CouponSpecialOfferQuery{
		config: c.config,
	}
}

// Get returns a CouponSpecialOffer entity by its id.
func (c *CouponSpecialOfferClient) Get(ctx context.Context, id uuid.UUID) (*CouponSpecialOffer, error) {
	return c.Query().Where(couponspecialoffer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CouponSpecialOfferClient) GetX(ctx context.Context, id uuid.UUID) *CouponSpecialOffer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CouponSpecialOfferClient) Hooks() []Hook {
	hooks := c.hooks.CouponSpecialOffer
	return append(hooks[:len(hooks):len(hooks)], couponspecialoffer.Hooks[:]...)
}