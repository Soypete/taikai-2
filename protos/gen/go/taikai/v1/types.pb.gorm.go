// Code generated by protoc-gen-go-gorm. DO NOT EDIT.
// source: taikai/v1/types.proto

package taikaiv1

import (
	context "context"
	uuid "github.com/google/uuid"
	pq "github.com/lib/pq"
	lo "github.com/samber/lo"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	gorm "gorm.io/gorm"
	clause "gorm.io/gorm/clause"
	sync "sync"
	time "time"
)

// cockroachdb doesn't support nanosecond timestamp columns so use microsecond instead
const TimestampFormat = "2006-01-02T15:04:05.999999Z07:00"

type HelloGormModels []*HelloGormModel
type HelloProtos []*Hello
type HelloGormModel struct {

	// @gotags: fake:"skip"
	Id *string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"created_at" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updated_at" fake:"skip"`

	// @gotags: fake:"{number:1,1}"
	HelloType int `json:"hello_type" fake:"{number:1,1}"`

	// @gotags: fake:"{beername}"
	PersonName *string `json:"person_name" fake:"{beername}"`
}

func (m *HelloGormModel) TableName() string {
	return "hellos"
}

func (m HelloGormModels) ToProtos() (protos HelloProtos, err error) {
	protos = HelloProtos{}
	for _, model := range m {
		var proto *Hello
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p HelloProtos) ToModels() (models HelloGormModels, err error) {
	models = HelloGormModels{}
	for _, proto := range p {
		var model *HelloGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *HelloGormModel) ToProto() (theProto *Hello, err error) {
	if m == nil {
		return
	}
	theProto = &Hello{}

	theProto.Id = m.Id

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	theProto.HelloType = HelloType(m.HelloType)

	theProto.PersonName = m.PersonName

	return
}

func (p *Hello) GetProtoId() *string {
	return p.Id
}

func (p *Hello) SetProtoId(id string) {
	p.Id = lo.ToPtr(id)
}

func (m *HelloGormModel) New() interface{} {
	return &HelloGormModel{}
}

func (m *HelloGormModel) GetModelId() *string {
	return m.Id
}

func (m *HelloGormModel) SetModelId(id string) {
	if m == nil {
		m = &HelloGormModel{}
	}
	m.Id = lo.ToPtr(id)
}

func (p *Hello) ToModel() (theModel *HelloGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &HelloGormModel{}

	theModel.Id = p.Id

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	theModel.HelloType = int(p.HelloType)

	theModel.PersonName = p.PersonName

	return
}

func (m HelloGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		if model.Id != nil {
			ids = append(ids, *model.Id)
		}
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("id in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *HelloProtos) Upsert(ctx context.Context, tx *gorm.DB) (models HelloGormModels, err error) {
	if p != nil {
		for _, proto := range *p {
			if proto.Id == nil {
				proto.Id = lo.ToPtr(uuid.New().String())
			}
		}
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *HelloProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models HelloGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = HelloProtos{}
		}
	}
	return
}

func (p *HelloProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models HelloGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("id in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = HelloProtos{}
		}
	}
	return
}

func DeleteHelloGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("id in ?", ids)
	return statement.Delete(&HelloGormModel{}).Error
}

type OrgGormModels []*OrgGormModel
type OrgProtos []*Org
type OrgGormModel struct {

	// @gotags: fake:"skip"
	Name string `json:"name" fake:"skip"`

	// @gotags: fake:"skip"
	Id string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	Description string `json:"description" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"created_at" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updated_at" fake:"skip"`

	// @gotags: fake:"skip"
	DeletedAt *time.Time `gorm:"type:timestamp;" json:"deleted_at" fake:"skip"`
}

func (m *OrgGormModel) TableName() string {
	return "orgs"
}

func (m OrgGormModels) ToProtos() (protos OrgProtos, err error) {
	protos = OrgProtos{}
	for _, model := range m {
		var proto *Org
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p OrgProtos) ToModels() (models OrgGormModels, err error) {
	models = OrgGormModels{}
	for _, proto := range p {
		var model *OrgGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *OrgGormModel) ToProto() (theProto *Org, err error) {
	if m == nil {
		return
	}
	theProto = &Org{}

	theProto.Name = m.Name

	theProto.Id = m.Id

	theProto.Description = m.Description

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	if m.DeletedAt != nil {
		theProto.DeletedAt = timestamppb.New(*m.DeletedAt)
	}

	return
}

func (p *Org) GetProtoId() *string {
	return p.Id
}

func (p *Org) SetProtoId(id string) {
	p.Id = lo.ToPtr(id)
}

func (m *OrgGormModel) New() interface{} {
	return &OrgGormModel{}
}

func (m *OrgGormModel) GetModelId() *string {
	return m.Id
}

func (m *OrgGormModel) SetModelId(id string) {
	if m == nil {
		m = &OrgGormModel{}
	}
	m.Id = lo.ToPtr(id)
}

func (p *Org) ToModel() (theModel *OrgGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &OrgGormModel{}

	theModel.Name = p.Name

	theModel.Id = p.Id

	theModel.Description = p.Description

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	if p.DeletedAt != nil {
		theModel.DeletedAt = lo.ToPtr(p.DeletedAt.AsTime())
	}

	return
}

func (m OrgGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		if model.Id != nil {
			ids = append(ids, *model.Id)
		}
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("id in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *OrgProtos) Upsert(ctx context.Context, tx *gorm.DB) (models OrgGormModels, err error) {
	if p != nil {
		for _, proto := range *p {
			if proto.Id == nil {
				proto.Id = lo.ToPtr(uuid.New().String())
			}
		}
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *OrgProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models OrgGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = OrgProtos{}
		}
	}
	return
}

func (p *OrgProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models OrgGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("id in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = OrgProtos{}
		}
	}
	return
}

func DeleteOrgGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("id in ?", ids)
	return statement.Delete(&OrgGormModel{}).Error
}

type UserGormModels []*UserGormModel
type UserProtos []*User
type UserGormModel struct {

	// @gotags: fake:"skip"
	Name string `json:"name" fake:"skip"`

	// @gotags: fake:"skip"
	Id string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	Email string `json:"email" fake:"skip"`

	// @gotags: fake:"skip"
	GroupIds pq.StringArray `gorm:"type:text[];" json:"group_ids" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"created_at" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updated_at" fake:"skip"`

	// @gotags: fake:"skip"
	DeletedAt *time.Time `gorm:"type:timestamp;" json:"deleted_at" fake:"skip"`
}

func (m *UserGormModel) TableName() string {
	return "users"
}

func (m UserGormModels) ToProtos() (protos UserProtos, err error) {
	protos = UserProtos{}
	for _, model := range m {
		var proto *User
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p UserProtos) ToModels() (models UserGormModels, err error) {
	models = UserGormModels{}
	for _, proto := range p {
		var model *UserGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *UserGormModel) ToProto() (theProto *User, err error) {
	if m == nil {
		return
	}
	theProto = &User{}

	theProto.Name = m.Name

	theProto.Id = m.Id

	theProto.Email = m.Email

	theProto.GroupIds = m.GroupIds

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	if m.DeletedAt != nil {
		theProto.DeletedAt = timestamppb.New(*m.DeletedAt)
	}

	return
}

func (p *User) GetProtoId() *string {
	return p.Id
}

func (p *User) SetProtoId(id string) {
	p.Id = lo.ToPtr(id)
}

func (m *UserGormModel) New() interface{} {
	return &UserGormModel{}
}

func (m *UserGormModel) GetModelId() *string {
	return m.Id
}

func (m *UserGormModel) SetModelId(id string) {
	if m == nil {
		m = &UserGormModel{}
	}
	m.Id = lo.ToPtr(id)
}

func (p *User) ToModel() (theModel *UserGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &UserGormModel{}

	theModel.Name = p.Name

	theModel.Id = p.Id

	theModel.Email = p.Email

	theModel.GroupIds = p.GroupIds

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	if p.DeletedAt != nil {
		theModel.DeletedAt = lo.ToPtr(p.DeletedAt.AsTime())
	}

	return
}

func (m UserGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		if model.Id != nil {
			ids = append(ids, *model.Id)
		}
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("id in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *UserProtos) Upsert(ctx context.Context, tx *gorm.DB) (models UserGormModels, err error) {
	if p != nil {
		for _, proto := range *p {
			if proto.Id == nil {
				proto.Id = lo.ToPtr(uuid.New().String())
			}
		}
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *UserProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models UserGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = UserProtos{}
		}
	}
	return
}

func (p *UserProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models UserGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("id in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = UserProtos{}
		}
	}
	return
}

func DeleteUserGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("id in ?", ids)
	return statement.Delete(&UserGormModel{}).Error
}

type GroupGormModels []*GroupGormModel
type GroupProtos []*Group
type GroupGormModel struct {

	// @gotags: fake:"skip"
	Name string `json:"name" fake:"skip"`

	// @gotags: fake:"skip"
	Id string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	OrgId string `json:"org_id" fake:"skip"`

	// @gotags: fake:"skip"
	OwnerIds pq.StringArray `gorm:"type:text[];" json:"owner_ids" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"created_at" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updated_at" fake:"skip"`

	// @gotags: fake:"skip"
	DeletedAt *time.Time `gorm:"type:timestamp;" json:"deleted_at" fake:"skip"`
}

func (m *GroupGormModel) TableName() string {
	return "groups"
}

func (m GroupGormModels) ToProtos() (protos GroupProtos, err error) {
	protos = GroupProtos{}
	for _, model := range m {
		var proto *Group
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p GroupProtos) ToModels() (models GroupGormModels, err error) {
	models = GroupGormModels{}
	for _, proto := range p {
		var model *GroupGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *GroupGormModel) ToProto() (theProto *Group, err error) {
	if m == nil {
		return
	}
	theProto = &Group{}

	theProto.Name = m.Name

	theProto.Id = m.Id

	theProto.OrgId = m.OrgId

	theProto.OwnerIds = m.OwnerIds

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	if m.DeletedAt != nil {
		theProto.DeletedAt = timestamppb.New(*m.DeletedAt)
	}

	return
}

func (p *Group) GetProtoId() *string {
	return p.Id
}

func (p *Group) SetProtoId(id string) {
	p.Id = lo.ToPtr(id)
}

func (m *GroupGormModel) New() interface{} {
	return &GroupGormModel{}
}

func (m *GroupGormModel) GetModelId() *string {
	return m.Id
}

func (m *GroupGormModel) SetModelId(id string) {
	if m == nil {
		m = &GroupGormModel{}
	}
	m.Id = lo.ToPtr(id)
}

func (p *Group) ToModel() (theModel *GroupGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &GroupGormModel{}

	theModel.Name = p.Name

	theModel.Id = p.Id

	theModel.OrgId = p.OrgId

	theModel.OwnerIds = p.OwnerIds

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	if p.DeletedAt != nil {
		theModel.DeletedAt = lo.ToPtr(p.DeletedAt.AsTime())
	}

	return
}

func (m GroupGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		if model.Id != nil {
			ids = append(ids, *model.Id)
		}
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("id in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *GroupProtos) Upsert(ctx context.Context, tx *gorm.DB) (models GroupGormModels, err error) {
	if p != nil {
		for _, proto := range *p {
			if proto.Id == nil {
				proto.Id = lo.ToPtr(uuid.New().String())
			}
		}
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *GroupProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models GroupGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = GroupProtos{}
		}
	}
	return
}

func (p *GroupProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models GroupGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("id in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = GroupProtos{}
		}
	}
	return
}

func DeleteGroupGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("id in ?", ids)
	return statement.Delete(&GroupGormModel{}).Error
}

type EventGormModels []*EventGormModel
type EventProtos []*Event
type EventGormModel struct {

	// @gotags: fake:"skip"
	Name string `json:"name" fake:"skip"`

	// @gotags: fake:"skip"
	Id string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();" json:"id" fake:"skip"`

	// @gotags: fake:"skip"
	GroupId string `json:"group_id" fake:"skip"`

	// @gotags: fake:"skip"
	Title string `json:"title" fake:"skip"`

	// @gotags: fake:"skip"
	Location string `json:"location" fake:"skip"`

	// @gotags: fake:"skip"
	Description string `json:"description" fake:"skip"`

	// @gotags: fake:"skip"
	UserIds []*UserGormModel `json:"user_ids" fake:"skip"`

	// @gotags: fake:"skip"
	StartTime *time.Time `gorm:"type:timestamp;" json:"start_time" fake:"skip"`

	// @gotags: fake:"skip"
	EndTime *time.Time `gorm:"type:timestamp;" json:"end_time" fake:"skip"`

	// @gotags: fake:"skip"
	CreatedAt *time.Time `gorm:"type:timestamp;" json:"created_at" fake:"skip"`

	// @gotags: fake:"skip"
	UpdatedAt *time.Time `gorm:"type:timestamp;" json:"updated_at" fake:"skip"`

	// @gotags: fake:"skip"
	DeletedAt *time.Time `gorm:"type:timestamp;" json:"deleted_at" fake:"skip"`
}

func (m *EventGormModel) TableName() string {
	return "events"
}

func (m EventGormModels) ToProtos() (protos EventProtos, err error) {
	protos = EventProtos{}
	for _, model := range m {
		var proto *Event
		if proto, err = model.ToProto(); err != nil {
			return
		}
		protos = append(protos, proto)
	}
	return
}

func (p EventProtos) ToModels() (models EventGormModels, err error) {
	models = EventGormModels{}
	for _, proto := range p {
		var model *EventGormModel
		if model, err = proto.ToModel(); err != nil {
			return
		}
		models = append(models, model)
	}
	return
}

func (m *EventGormModel) ToProto() (theProto *Event, err error) {
	if m == nil {
		return
	}
	theProto = &Event{}

	theProto.Name = m.Name

	theProto.Id = m.Id

	theProto.GroupId = m.GroupId

	theProto.Title = m.Title

	theProto.Location = m.Location

	theProto.Description = m.Description

	if len(m.UserIds) > 0 {
		theProto.UserIds = []*User{}
		for _, item := range m.UserIds {
			var UserIdsProto *User
			if UserIdsProto, err = item.ToProto(); err != nil {
				return
			} else {
				theProto.UserIds = append(theProto.UserIds, UserIdsProto)
			}
		}
	}

	if m.StartTime != nil {
		theProto.StartTime = timestamppb.New(*m.StartTime)
	}

	if m.EndTime != nil {
		theProto.EndTime = timestamppb.New(*m.EndTime)
	}

	if m.CreatedAt != nil {
		theProto.CreatedAt = timestamppb.New(*m.CreatedAt)
	}

	if m.UpdatedAt != nil {
		theProto.UpdatedAt = timestamppb.New(*m.UpdatedAt)
	}

	if m.DeletedAt != nil {
		theProto.DeletedAt = timestamppb.New(*m.DeletedAt)
	}

	return
}

func (p *Event) GetProtoId() *string {
	return p.Id
}

func (p *Event) SetProtoId(id string) {
	p.Id = lo.ToPtr(id)
}

func (m *EventGormModel) New() interface{} {
	return &EventGormModel{}
}

func (m *EventGormModel) GetModelId() *string {
	return m.Id
}

func (m *EventGormModel) SetModelId(id string) {
	if m == nil {
		m = &EventGormModel{}
	}
	m.Id = lo.ToPtr(id)
}

func (p *Event) ToModel() (theModel *EventGormModel, err error) {
	if p == nil {
		return
	}
	theModel = &EventGormModel{}

	theModel.Name = p.Name

	theModel.Id = p.Id

	theModel.GroupId = p.GroupId

	theModel.Title = p.Title

	theModel.Location = p.Location

	theModel.Description = p.Description

	if len(p.UserIds) > 0 {
		theModel.UserIds = []*UserGormModel{}
		for _, item := range p.UserIds {
			var UserIdsModel *UserGormModel
			if UserIdsModel, err = item.ToModel(); err != nil {
				return
			} else {
				theModel.UserIds = append(theModel.UserIds, UserIdsModel)
			}
		}
	}

	if p.StartTime != nil {
		theModel.StartTime = lo.ToPtr(p.StartTime.AsTime())
	}

	if p.EndTime != nil {
		theModel.EndTime = lo.ToPtr(p.EndTime.AsTime())
	}

	if p.CreatedAt != nil {
		theModel.CreatedAt = lo.ToPtr(p.CreatedAt.AsTime())
	}

	if p.UpdatedAt != nil {
		theModel.UpdatedAt = lo.ToPtr(p.UpdatedAt.AsTime())
	}

	if p.DeletedAt != nil {
		theModel.DeletedAt = lo.ToPtr(p.DeletedAt.AsTime())
	}

	return
}

func (m EventGormModels) GetByModelIds(ctx context.Context, tx *gorm.DB, preloads ...string) (err error) {
	ids := []string{}
	for _, model := range m {
		if model.Id != nil {
			ids = append(ids, *model.Id)
		}
	}
	if len(ids) > 0 {
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		err = statement.Where("id in ?", ids).Find(&m).Error
	}
	return
}

// Upsert creates the protos using an on conflict clause to do updates. This function does not update *any* associations
// use gorm's association mode functions to update associations as you see fit after calling upsert. See https://gorm.io/docs/associations.html#Replace-Associations
func (p *EventProtos) Upsert(ctx context.Context, tx *gorm.DB) (models EventGormModels, err error) {
	if p != nil {
		for _, proto := range *p {
			if proto.Id == nil {
				proto.Id = lo.ToPtr(uuid.New().String())
			}
		}
		models, err = p.ToModels()
		if err != nil {
			return
		}
		// create new session so the tx isn't modified
		session := tx.Session(&gorm.Session{})
		err = session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error
	}
	return
}

func (p *EventProtos) List(ctx context.Context, tx *gorm.DB, limit, offset int, order interface{}, preloads ...string) (err error) {
	if p != nil {
		var models EventGormModels
		statement := tx.Preload(clause.Associations).Limit(limit).Offset(offset)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if order != nil {
			statement = statement.Order(order)
		}
		if err = statement.Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = EventProtos{}
		}
	}
	return
}

func (p *EventProtos) GetByIds(ctx context.Context, tx *gorm.DB, ids []string, preloads ...string) (err error) {
	if p != nil {
		var models EventGormModels
		statement := tx.Preload(clause.Associations)
		for _, preload := range preloads {
			statement = statement.Preload(preload)
		}
		if err = statement.Where("id in ?", ids).Find(&models).Error; err != nil {
			return
		}
		if len(models) > 0 {
			*p, err = models.ToProtos()
		} else {
			*p = EventProtos{}
		}
	}
	return
}

func DeleteEventGormModels(ctx context.Context, tx *gorm.DB, ids []string) error {
	statement := tx.Where("id in ?", ids)
	return statement.Delete(&EventGormModel{}).Error
}

// Protos is a union of other types that defines which types may be used in generic functions
type Protos interface {
	*Hello | *Org | *User | *Group | *Event
	GetProtoId() *string
	SetProtoId(string)
}

// Models is a union of other types that defines which types may be used in generic functions
type Models interface {
	*HelloGormModel | *OrgGormModel | *UserGormModel | *GroupGormModel | *EventGormModel
	GetModelId() *string
	SetModelId(string)
	New() interface{}
	TableName() string // tabler interface for gorm model, gives us access to the table name that gorm will use, see https://gorm.io/docs/conventions.html#TableName
}

// Proto[M Models] is an interface type that defines behavior for the implementer of a given Models type
type Proto[M Models] interface {
	GetProtoId() *string
	SetProtoId(string)
	ToModel() (M, error)
}

// Model[P Protos] is an interface type that defines behavior for the implementer of a given Protos type
type Model[P Protos] interface {
	ToProto() (P, error)
}

// ToModels converts an array of protos to an array of gorm db models by calling the proto's ToModel method
func ToModels[P Protos, M Models](protos interface{}) ([]M, error) {
	converted := ConvertProtosToProtosM[P, M](protos)
	models := []M{}
	for _, proto := range converted {
		model, err := proto.ToModel()
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}
	return models, nil
}

// ConvertProtosToProtosM converts a given slice of protos into an array of the Proto interface type, which can then
// leverage the interface methods
func ConvertProtosToProtosM[P Protos, M Models](protos interface{}) []Proto[M] {
	assertedProtos := protos.([]P)
	things := make([]Proto[M], len(assertedProtos))
	for i, v := range assertedProtos {
		things[i] = ConvertProtoToProtosM[P, M](v)
	}
	return things
}

// ConvertProtoToProtosM converts a single proto to a Proto[M]
func ConvertProtoToProtosM[P Protos, M Models](proto interface{}) Proto[M] {
	return any(proto).(Proto[M])
}

// ConvertProtosToProtosM converts a given slice of protos into an array of the Proto interface type, which can then
// leverage the interface methods
func ConvertModelsToModelsP[P Protos, M Models](models interface{}) []Model[P] {
	assertedModels := models.([]M)
	things := make([]Model[P], len(assertedModels))
	for i, v := range assertedModels {
		things[i] = ConvertModelToModelP[P, M](v)
	}
	return things
}

// ConvertProtoToProtosM converts a single proto to a Proto[M]
func ConvertModelToModelP[P Protos, M Models](model interface{}) Model[P] {
	return any(model).(Model[P])
}

// ToProtos converts an array of models into an array of protos by calling the model's ToProto method
func ToProtos[P Protos, M Models](models interface{}) ([]P, error) {
	converted := ConvertModelsToModelsP[P, M](models)
	protos := []P{}
	for _, model := range converted {
		proto, err := model.ToProto()
		if err != nil {
			return nil, err
		}
		protos = append(protos, proto)
	}
	return protos, nil
}

// Upsert is a generic function that will upsert any of the generated protos, returning the upserted models. Upsert
// excludes all associations, and uses an on conflict clause to handle upsert. A function may be provided to be executed
// during the transaction. The function is executed after the upsert. If the function returns an error, the transaction
// will be rolled back.
func Upsert[P Protos, M Models](ctx context.Context, db *gorm.DB, protos interface{}) ([]M, error) {
	converted := ConvertProtosToProtosM[P, M](protos)
	if len(converted) > 0 {
		models := []M{}
		for _, proto := range converted {
			if proto.GetProtoId() == nil {
				proto.SetProtoId(uuid.New().String())
			}
			model, err := proto.ToModel()
			if err != nil {
				return nil, err
			}
			models = append(models, model)
		}
		session := db.Session(&gorm.Session{})
		err := session.
			// on conflict, update all fields
			Clauses(clause.OnConflict{
				UpdateAll: true,
			}).
			// exclude associations from upsert
			Omit(clause.Associations).
			Create(&models).Error

		return models, err
	}
	return nil, nil
}

// Delete is a generic function that will delete any of the generated protos. A function may be provided to be executed
// during the transaction. The function is executed after the delete. If the function returns an error, the transaction
// will be rolled back.
func Delete[M Models](ctx context.Context, db *gorm.DB, ids []string) ([]M, error) {
	if len(ids) > 0 {
		session := db.Session(&gorm.Session{})
		models := []M{}
		err := session.Where("id in ?", ids).Delete(&models).Error
		return models, err
	}
	return nil, nil
}

// List lists the given model type
func List[M Models](ctx context.Context, db *gorm.DB, limit, offset int, orderBy string, preloads map[string][]interface{}) ([]M, error) {
	session := db.Session(&gorm.Session{}).WithContext(ctx)
	// set limit
	if limit > 0 {
		session = session.Limit(limit)
	}
	// set offset
	if offset > 0 {
		session = session.Offset(offset)
	}
	// set preloads
	for preload, args := range preloads {
		session = session.Preload(preload, args...)
	}
	// set order by
	if orderBy != "" {
		session = session.Order(orderBy)
	}
	// execute
	var models []M
	err := session.Find(&models).Error
	return models, err
}

// GetByIds gets the given model type by id
func GetByIds[M Models](ctx context.Context, db *gorm.DB, ids []string, preloads map[string][]interface{}) ([]M, error) {
	session := db.Session(&gorm.Session{}).WithContext(ctx)
	// set preloads
	for preload, args := range preloads {
		session = session.Preload(preload, args...)
	}
	models := []M{}
	err := session.Where("id in ?", ids).Find(&models).Error
	return models, err
}

// ManyToManyAssociations is a sync map with helper functions. I'm using a sync.map so that it's thread safe, and
// a struct to allow us to easily define behavior we can use elsewhere
type ManyToManyAssociations struct {
	data sync.Map
}

func (m *ManyToManyAssociations) Associations() map[string][]string {
	associations := map[string][]string{}
	m.data.Range(func(key, value any) bool {
		associations[key.(string)] = value.([]string)
		return true
	})
	return associations
}

func (m *ManyToManyAssociations) AddAssociation(modelId, associatedId string) {
	var associations []string
	val, ok := m.data.Load(modelId)
	if ok {
		associations = val.([]string)
		associations = append(associations, associatedId)
	} else {
		associations = []string{associatedId}
	}
	m.data.Store(modelId, associations)
}

func ReplaceManyToMany[L Models, R Models](ctx context.Context, db *gorm.DB, associations *ManyToManyAssociations, associationName string) error {
	session := db.Session(&gorm.Session{})
	session = session.Clauses(clause.OnConflict{DoNothing: true})
	for id, associatedIds := range associations.Associations() {
		var associations []R
		var temp L
		model := temp.New().(L)
		model.SetModelId(id)
		for _, id := range associatedIds {
			var associatedTemp R
			associatedModel := associatedTemp.New().(R)
			associatedModel.SetModelId(id)
			associations = append(associations, associatedModel)
		}
		// omit is required otherwise it generates some weird sql and tries to update other columns that don't exist
		err := session.Model(&model).Omit("*").Association(associationName).Replace(&associations)
		if err != nil {
			return err
		}
	}
	return nil
}

func AssociateManyToMany[L Models, R Models](ctx context.Context, db *gorm.DB, associations *ManyToManyAssociations, associationName string) error {
	session := db.Session(&gorm.Session{})
	session = session.Clauses(clause.OnConflict{DoNothing: true})
	for id, associatedIds := range associations.Associations() {
		var associations []R
		var temp L
		model := temp.New().(L)
		model.SetModelId(id)
		for _, id := range associatedIds {
			var associatedTemp R
			associatedModel := associatedTemp.New().(R)
			associatedModel.SetModelId(id)
			associations = append(associations, associatedModel)
		}
		// omit is required otherwise it generates some weird sql and tries to update other columns that don't exist
		err := session.Model(&model).Omit("*").Association(associationName).Append(&associations)
		if err != nil {
			return err
		}
	}
	return nil
}

func DissociateManyToMany[L Models, R Models](ctx context.Context, db *gorm.DB, associations *ManyToManyAssociations, associationName string) error {
	session := db.Session(&gorm.Session{})
	for id, associatedIds := range associations.Associations() {
		var associations []R
		var temp L
		model := temp.New().(L)
		model.SetModelId(id)
		for _, id := range associatedIds {
			var associatedTemp R
			associatedModel := associatedTemp.New().(R)
			associatedModel.SetModelId(id)
			associations = append(associations, associatedModel)
		}
		// omit is required otherwise it generates some weird sql and tries to update other columns that don't exist
		txErr := session.Model(&model).Omit("*").Association(associationName).Delete(&associations)
		if txErr != nil {
			return txErr
		}
	}
	return nil
}
