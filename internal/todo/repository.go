package todo

import "gorm.io/gorm"

// Repository defines the interface for todo data operations.
type Repository interface {
	FindAll() ([]Todo, error)
	FindByID(id uint) (*Todo, error)
	Create(todo *Todo) error
	Update(todo *Todo) error
	Delete(id uint) error
}

// repository is the GORM implementation of the Repository interface.
// HINT: gorm is 'injected' to the repository struct.
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new Repository with the given GORM DB instance.
// HINT: A 'constructor' function for the Repository struct.
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// FindAll retrieves all todo items from the database.
func (r *repository) FindAll() ([]Todo, error) {
	var todos []Todo
	if err := r.db.Order("id ASC").Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// FindByID retrieves a todo item by its ID.
func (r *repository) FindByID(id uint) (*Todo, error) {
	var todo Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &todo, nil
}

// Create adds a new todo item to the database.
func (r *repository) Create(todo *Todo) error {
	// HINT: This is how we CREATE resource and HANDLE errors in GORM at the same time.
	return r.db.Create(todo).Error
}

// Update modifies an existing todo item in the database.
func (r *repository) Update(todo *Todo) error {
	return r.db.Save(todo).Error
}

// Delete removes a todo item from the database by its ID.
func (r *repository) Delete(id uint) error {
	// HINT: This is how we DELETE a resource in GORM. We pass an empty struct with the relevant ID.
	res := r.db.Delete(&Todo{}, id)
	return res.Error
}
