// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/cateiru/cateiru.com/ent/migrate"
	"github.com/google/uuid"

	"github.com/cateiru/cateiru.com/ent/biography"
	"github.com/cateiru/cateiru.com/ent/category"
	"github.com/cateiru/cateiru.com/ent/contact"
	"github.com/cateiru/cateiru.com/ent/link"
	"github.com/cateiru/cateiru.com/ent/location"
	"github.com/cateiru/cateiru.com/ent/notice"
	"github.com/cateiru/cateiru.com/ent/product"
	"github.com/cateiru/cateiru.com/ent/session"
	"github.com/cateiru/cateiru.com/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Biography is the client for interacting with the Biography builders.
	Biography *BiographyClient
	// Category is the client for interacting with the Category builders.
	Category *CategoryClient
	// Contact is the client for interacting with the Contact builders.
	Contact *ContactClient
	// Link is the client for interacting with the Link builders.
	Link *LinkClient
	// Location is the client for interacting with the Location builders.
	Location *LocationClient
	// Notice is the client for interacting with the Notice builders.
	Notice *NoticeClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Biography = NewBiographyClient(c.config)
	c.Category = NewCategoryClient(c.config)
	c.Contact = NewContactClient(c.config)
	c.Link = NewLinkClient(c.config)
	c.Location = NewLocationClient(c.config)
	c.Notice = NewNoticeClient(c.config)
	c.Product = NewProductClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:       ctx,
		config:    cfg,
		Biography: NewBiographyClient(cfg),
		Category:  NewCategoryClient(cfg),
		Contact:   NewContactClient(cfg),
		Link:      NewLinkClient(cfg),
		Location:  NewLocationClient(cfg),
		Notice:    NewNoticeClient(cfg),
		Product:   NewProductClient(cfg),
		Session:   NewSessionClient(cfg),
		User:      NewUserClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		Biography: NewBiographyClient(cfg),
		Category:  NewCategoryClient(cfg),
		Contact:   NewContactClient(cfg),
		Link:      NewLinkClient(cfg),
		Location:  NewLocationClient(cfg),
		Notice:    NewNoticeClient(cfg),
		Product:   NewProductClient(cfg),
		Session:   NewSessionClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Biography.
//		Query().
//		Count(ctx)
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
	c.Biography.Use(hooks...)
	c.Category.Use(hooks...)
	c.Contact.Use(hooks...)
	c.Link.Use(hooks...)
	c.Location.Use(hooks...)
	c.Notice.Use(hooks...)
	c.Product.Use(hooks...)
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
}

// BiographyClient is a client for the Biography schema.
type BiographyClient struct {
	config
}

// NewBiographyClient returns a client for the Biography from the given config.
func NewBiographyClient(c config) *BiographyClient {
	return &BiographyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `biography.Hooks(f(g(h())))`.
func (c *BiographyClient) Use(hooks ...Hook) {
	c.hooks.Biography = append(c.hooks.Biography, hooks...)
}

// Create returns a builder for creating a Biography entity.
func (c *BiographyClient) Create() *BiographyCreate {
	mutation := newBiographyMutation(c.config, OpCreate)
	return &BiographyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Biography entities.
func (c *BiographyClient) CreateBulk(builders ...*BiographyCreate) *BiographyCreateBulk {
	return &BiographyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Biography.
func (c *BiographyClient) Update() *BiographyUpdate {
	mutation := newBiographyMutation(c.config, OpUpdate)
	return &BiographyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BiographyClient) UpdateOne(b *Biography) *BiographyUpdateOne {
	mutation := newBiographyMutation(c.config, OpUpdateOne, withBiography(b))
	return &BiographyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BiographyClient) UpdateOneID(id uint32) *BiographyUpdateOne {
	mutation := newBiographyMutation(c.config, OpUpdateOne, withBiographyID(id))
	return &BiographyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Biography.
func (c *BiographyClient) Delete() *BiographyDelete {
	mutation := newBiographyMutation(c.config, OpDelete)
	return &BiographyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BiographyClient) DeleteOne(b *Biography) *BiographyDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BiographyClient) DeleteOneID(id uint32) *BiographyDeleteOne {
	builder := c.Delete().Where(biography.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BiographyDeleteOne{builder}
}

// Query returns a query builder for Biography.
func (c *BiographyClient) Query() *BiographyQuery {
	return &BiographyQuery{
		config: c.config,
	}
}

// Get returns a Biography entity by its id.
func (c *BiographyClient) Get(ctx context.Context, id uint32) (*Biography, error) {
	return c.Query().Where(biography.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BiographyClient) GetX(ctx context.Context, id uint32) *Biography {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BiographyClient) Hooks() []Hook {
	return c.hooks.Biography
}

// CategoryClient is a client for the Category schema.
type CategoryClient struct {
	config
}

// NewCategoryClient returns a client for the Category from the given config.
func NewCategoryClient(c config) *CategoryClient {
	return &CategoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `category.Hooks(f(g(h())))`.
func (c *CategoryClient) Use(hooks ...Hook) {
	c.hooks.Category = append(c.hooks.Category, hooks...)
}

// Create returns a builder for creating a Category entity.
func (c *CategoryClient) Create() *CategoryCreate {
	mutation := newCategoryMutation(c.config, OpCreate)
	return &CategoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Category entities.
func (c *CategoryClient) CreateBulk(builders ...*CategoryCreate) *CategoryCreateBulk {
	return &CategoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Category.
func (c *CategoryClient) Update() *CategoryUpdate {
	mutation := newCategoryMutation(c.config, OpUpdate)
	return &CategoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CategoryClient) UpdateOne(ca *Category) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategory(ca))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CategoryClient) UpdateOneID(id uint32) *CategoryUpdateOne {
	mutation := newCategoryMutation(c.config, OpUpdateOne, withCategoryID(id))
	return &CategoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Category.
func (c *CategoryClient) Delete() *CategoryDelete {
	mutation := newCategoryMutation(c.config, OpDelete)
	return &CategoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CategoryClient) DeleteOne(ca *Category) *CategoryDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CategoryClient) DeleteOneID(id uint32) *CategoryDeleteOne {
	builder := c.Delete().Where(category.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CategoryDeleteOne{builder}
}

// Query returns a query builder for Category.
func (c *CategoryClient) Query() *CategoryQuery {
	return &CategoryQuery{
		config: c.config,
	}
}

// Get returns a Category entity by its id.
func (c *CategoryClient) Get(ctx context.Context, id uint32) (*Category, error) {
	return c.Query().Where(category.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CategoryClient) GetX(ctx context.Context, id uint32) *Category {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CategoryClient) Hooks() []Hook {
	return c.hooks.Category
}

// ContactClient is a client for the Contact schema.
type ContactClient struct {
	config
}

// NewContactClient returns a client for the Contact from the given config.
func NewContactClient(c config) *ContactClient {
	return &ContactClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `contact.Hooks(f(g(h())))`.
func (c *ContactClient) Use(hooks ...Hook) {
	c.hooks.Contact = append(c.hooks.Contact, hooks...)
}

// Create returns a builder for creating a Contact entity.
func (c *ContactClient) Create() *ContactCreate {
	mutation := newContactMutation(c.config, OpCreate)
	return &ContactCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Contact entities.
func (c *ContactClient) CreateBulk(builders ...*ContactCreate) *ContactCreateBulk {
	return &ContactCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Contact.
func (c *ContactClient) Update() *ContactUpdate {
	mutation := newContactMutation(c.config, OpUpdate)
	return &ContactUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ContactClient) UpdateOne(co *Contact) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContact(co))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ContactClient) UpdateOneID(id uint32) *ContactUpdateOne {
	mutation := newContactMutation(c.config, OpUpdateOne, withContactID(id))
	return &ContactUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Contact.
func (c *ContactClient) Delete() *ContactDelete {
	mutation := newContactMutation(c.config, OpDelete)
	return &ContactDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ContactClient) DeleteOne(co *Contact) *ContactDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ContactClient) DeleteOneID(id uint32) *ContactDeleteOne {
	builder := c.Delete().Where(contact.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ContactDeleteOne{builder}
}

// Query returns a query builder for Contact.
func (c *ContactClient) Query() *ContactQuery {
	return &ContactQuery{
		config: c.config,
	}
}

// Get returns a Contact entity by its id.
func (c *ContactClient) Get(ctx context.Context, id uint32) (*Contact, error) {
	return c.Query().Where(contact.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ContactClient) GetX(ctx context.Context, id uint32) *Contact {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ContactClient) Hooks() []Hook {
	return c.hooks.Contact
}

// LinkClient is a client for the Link schema.
type LinkClient struct {
	config
}

// NewLinkClient returns a client for the Link from the given config.
func NewLinkClient(c config) *LinkClient {
	return &LinkClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `link.Hooks(f(g(h())))`.
func (c *LinkClient) Use(hooks ...Hook) {
	c.hooks.Link = append(c.hooks.Link, hooks...)
}

// Create returns a builder for creating a Link entity.
func (c *LinkClient) Create() *LinkCreate {
	mutation := newLinkMutation(c.config, OpCreate)
	return &LinkCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Link entities.
func (c *LinkClient) CreateBulk(builders ...*LinkCreate) *LinkCreateBulk {
	return &LinkCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Link.
func (c *LinkClient) Update() *LinkUpdate {
	mutation := newLinkMutation(c.config, OpUpdate)
	return &LinkUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LinkClient) UpdateOne(l *Link) *LinkUpdateOne {
	mutation := newLinkMutation(c.config, OpUpdateOne, withLink(l))
	return &LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LinkClient) UpdateOneID(id uint32) *LinkUpdateOne {
	mutation := newLinkMutation(c.config, OpUpdateOne, withLinkID(id))
	return &LinkUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Link.
func (c *LinkClient) Delete() *LinkDelete {
	mutation := newLinkMutation(c.config, OpDelete)
	return &LinkDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LinkClient) DeleteOne(l *Link) *LinkDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *LinkClient) DeleteOneID(id uint32) *LinkDeleteOne {
	builder := c.Delete().Where(link.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LinkDeleteOne{builder}
}

// Query returns a query builder for Link.
func (c *LinkClient) Query() *LinkQuery {
	return &LinkQuery{
		config: c.config,
	}
}

// Get returns a Link entity by its id.
func (c *LinkClient) Get(ctx context.Context, id uint32) (*Link, error) {
	return c.Query().Where(link.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LinkClient) GetX(ctx context.Context, id uint32) *Link {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LinkClient) Hooks() []Hook {
	return c.hooks.Link
}

// LocationClient is a client for the Location schema.
type LocationClient struct {
	config
}

// NewLocationClient returns a client for the Location from the given config.
func NewLocationClient(c config) *LocationClient {
	return &LocationClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `location.Hooks(f(g(h())))`.
func (c *LocationClient) Use(hooks ...Hook) {
	c.hooks.Location = append(c.hooks.Location, hooks...)
}

// Create returns a builder for creating a Location entity.
func (c *LocationClient) Create() *LocationCreate {
	mutation := newLocationMutation(c.config, OpCreate)
	return &LocationCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Location entities.
func (c *LocationClient) CreateBulk(builders ...*LocationCreate) *LocationCreateBulk {
	return &LocationCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Location.
func (c *LocationClient) Update() *LocationUpdate {
	mutation := newLocationMutation(c.config, OpUpdate)
	return &LocationUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LocationClient) UpdateOne(l *Location) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocation(l))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LocationClient) UpdateOneID(id uint32) *LocationUpdateOne {
	mutation := newLocationMutation(c.config, OpUpdateOne, withLocationID(id))
	return &LocationUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Location.
func (c *LocationClient) Delete() *LocationDelete {
	mutation := newLocationMutation(c.config, OpDelete)
	return &LocationDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LocationClient) DeleteOne(l *Location) *LocationDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *LocationClient) DeleteOneID(id uint32) *LocationDeleteOne {
	builder := c.Delete().Where(location.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LocationDeleteOne{builder}
}

// Query returns a query builder for Location.
func (c *LocationClient) Query() *LocationQuery {
	return &LocationQuery{
		config: c.config,
	}
}

// Get returns a Location entity by its id.
func (c *LocationClient) Get(ctx context.Context, id uint32) (*Location, error) {
	return c.Query().Where(location.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LocationClient) GetX(ctx context.Context, id uint32) *Location {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LocationClient) Hooks() []Hook {
	return c.hooks.Location
}

// NoticeClient is a client for the Notice schema.
type NoticeClient struct {
	config
}

// NewNoticeClient returns a client for the Notice from the given config.
func NewNoticeClient(c config) *NoticeClient {
	return &NoticeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `notice.Hooks(f(g(h())))`.
func (c *NoticeClient) Use(hooks ...Hook) {
	c.hooks.Notice = append(c.hooks.Notice, hooks...)
}

// Create returns a builder for creating a Notice entity.
func (c *NoticeClient) Create() *NoticeCreate {
	mutation := newNoticeMutation(c.config, OpCreate)
	return &NoticeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Notice entities.
func (c *NoticeClient) CreateBulk(builders ...*NoticeCreate) *NoticeCreateBulk {
	return &NoticeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Notice.
func (c *NoticeClient) Update() *NoticeUpdate {
	mutation := newNoticeMutation(c.config, OpUpdate)
	return &NoticeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NoticeClient) UpdateOne(n *Notice) *NoticeUpdateOne {
	mutation := newNoticeMutation(c.config, OpUpdateOne, withNotice(n))
	return &NoticeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NoticeClient) UpdateOneID(id uint32) *NoticeUpdateOne {
	mutation := newNoticeMutation(c.config, OpUpdateOne, withNoticeID(id))
	return &NoticeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Notice.
func (c *NoticeClient) Delete() *NoticeDelete {
	mutation := newNoticeMutation(c.config, OpDelete)
	return &NoticeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NoticeClient) DeleteOne(n *Notice) *NoticeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NoticeClient) DeleteOneID(id uint32) *NoticeDeleteOne {
	builder := c.Delete().Where(notice.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NoticeDeleteOne{builder}
}

// Query returns a query builder for Notice.
func (c *NoticeClient) Query() *NoticeQuery {
	return &NoticeQuery{
		config: c.config,
	}
}

// Get returns a Notice entity by its id.
func (c *NoticeClient) Get(ctx context.Context, id uint32) (*Notice, error) {
	return c.Query().Where(notice.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NoticeClient) GetX(ctx context.Context, id uint32) *Notice {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *NoticeClient) Hooks() []Hook {
	return c.hooks.Notice
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Create returns a builder for creating a Product entity.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id uint32) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ProductClient) DeleteOneID(id uint32) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{
		config: c.config,
	}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id uint32) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id uint32) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a builder for creating a Session entity.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id uuid.UUID) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SessionClient) DeleteOneID(id uuid.UUID) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id uuid.UUID) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id uuid.UUID) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uint32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uint32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uint32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uint32) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}