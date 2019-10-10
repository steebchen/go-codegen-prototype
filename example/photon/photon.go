package photon

import (
	"context"
)

// helper methods
var StringNull = NullString{}

func String(str string) NullString {
	return NullString{
		Value: str,
		Valid: true,
	}
}

// NullString is a sql.NullString equivalent nullable string
type NullString struct {
	Value string
	Valid bool
}

// NullPost is a nullable Post
type NullPost struct {
	PostObject
	Valid bool
}

type IString struct {
	Value     string
	Specified bool
}

type INullString struct {
	Value     string
	Null      bool
	Specified bool
}

type PostObject struct {
	Title   string
	content NullString
	stuff   NullString
}

func (o PostObject) Content() (string, bool) {
	return o.content.Value, o.content.Valid
}

func (o PostObject) Stuff() (string, bool) {
	return o.stuff.Value, o.stuff.Valid
}

type StringField struct {
}

func (r StringField) Contains(str string) PostWhereOpts {
	return PostWhereOpts{}
}

func (r StringField) Equals(str string) PostWhereOpts {
	return PostWhereOpts{}
}

func (r StringField) EqualsOptional(str NullString) PostWhereOpts {
	return PostWhereOpts{}
}

func (r StringField) Null() PostWhereOpts {
	return PostWhereOpts{}
}

func (r StringField) EqualsPtr(str *string) PostWhereOpts {
	return PostWhereOpts{}
}

func NewClient() ClientStruct {
	return ClientStruct{}
}

var Post = PostQuery{}

type ClientStruct struct {
	Post PostMethods
}

type PostMethods struct {
	FindOne  PostMethodsFindOne
	FindMany PostMethodsFindMany
}

type PostMethodsFindOne struct {
}

type PostMethodsFindMany struct {
}

func (r PostMethodsFindOne) ID(id string) PostMethodsFindOne {
	return r
}

func (r PostMethodsFindOne) Where(query ...PostWhereOpts) PostMethodsFindOne {
	return r
}

func (r PostMethodsFindOne) Exec(ctx context.Context) (PostObject, error) {
	return PostObject{
		Title:   "John",
		content: NullString{Value: "f", Valid: true},
	}, nil
}

func (r PostMethodsFindMany) Where(query ...PostWhereOpts) PostMethodsFindMany {
	return r
}

func (r PostMethodsFindMany) OrderBy(query PostManyQuery) PostMethodsFindMany {
	return r
}

func (r PostMethodsFindMany) Exec(ctx context.Context) ([]PostObject, error) {
	return []PostObject{{
		Title:   "John",
		content: NullString{Value: "f", Valid: true},
	}}, nil
}

// CreateOne specifies options to create a user.
// This can be used with the fluent API:
//  CreateOne(photon.Post.CreateOne(user.Content("todo")))
func (r PostMethods) CreateOne(query ...UserCreate) (PostObject, error) {
	return PostObject{
		Title: "John",
	}, nil
}

func (r PostMethods) CreateMany(query []UserCreate) ([]PostObject, error) {
	return []PostObject{}, nil
}

type PostQuery struct {
	Title   StringField
	Content StringField

	Comments CommentQuery
}

type CommentQuery struct {
	Content StringField
}

func (r PostQuery) Where(where PostWhereOpts) PostOneQuery {
	return PostOneQuery{}
}

func (r PostQuery) Limit(count int) PostManyQuery {
	return PostManyQuery{
		Limit: count,
	}
}

func (r PostQuery) From(user PostObject) UserCreate {
	return UserCreate{
		User: user,
	}
}

func (r PostQuery) New() UserCreate {
	return UserCreate{}
}

func (r PostQuery) Relation(where UserCreateOpts) UserCreate {
	return UserCreate{
		nil,
		where,
	}
}

type UserCreate struct {
	User PostObject
	UserCreateOpts
}

// func (r UserCreate) Content(str NullString) UserCreate {
// 	r.Post.Content = str
// 	return r
// }

type UserCreateOpts struct {
	Posts UserCreateWithPosts
}

type UserCreateWithPosts struct {
	Create  []PostCreate
	Connect []UserCreatePostConnect
}

type PostCreate struct {
}

type UserCreatePostConnect struct {
}

type PostWhereOpts struct {
	Name         IString
	Email        INullString
	privateField string
}

type PostOneQuery struct {
	Where   PostWhereOpts
	OrderBy UserOrderByOpts
}

type PostManyQuery struct {
	Where   UserWhereManyOpts
	OrderBy UserOrderByOpts
	Limit   int
}

type UserWhereManyOpts struct {
	Name  IString
	Email INullString
}

type UserOrderByOpts struct {
}
