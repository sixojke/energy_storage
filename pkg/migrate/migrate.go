package migrate

import (
	"database/sql"
	"energy_storage/pkg/logger"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrate struct {
	*migrate.Migrate
}

// NewMigratorMysql создает новый объект Migrate для работы с миграциями в базе данных MySQL.
//
// Параметры:
//   - sourseURL: URL, указывающий на каталог с файлами миграций.
//   - db: Соединение с базой данных MySQL, полученное из sqlx.DB.
//
// Возвращает:
//   - *migrate.Migrate: Объект Migrate для управления миграциями.
//   - error: Возникает, если есть ошибка при создании объекта Migrate.
//
// Пример использования:
//
//	m, _ := NewMigratorMysql("file:///migrations", db)
func NewMigratorMysql(sourseURL string, db *sql.DB) (*Migrate, error) {
	if db == nil {
		return nil, fmt.Errorf("no database")
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create MySQL instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourseURL, "mysql", driver)
	if err != nil {
		return nil, fmt.Errorf("error creating migrator: %v", err)
	}

	return &Migrate{Migrate: m}, nil
}

// NewMigratorPostgres создает новый объект Migrate для работы с миграциями в базе данных MySQL.
//
// Параметры:
//   - sourseURL: URL, указывающий на каталог с файлами миграций.
//   - db: Соединение с базой данных PostgreSQL, полученное из sqlx.DB.
//
// Возвращает:
//   - *migrate.Migrate: Объект Migrate для управления миграциями.
//   - error: Возникает, если есть ошибка при создании объекта Migrate.
//
// Пример использования:
//
//	m, _ := NewMigratorPostgres("file:///migrations", db)
func NewMigratorPostgres(sourseURL string, db *sql.DB) (*Migrate, error) {
	if db == nil {
		return nil, fmt.Errorf("no database")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create MySQL instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(sourseURL, "postgres", driver)
	if err != nil {
		return nil, fmt.Errorf("error creating migrator: %v", err)
	}

	return &Migrate{Migrate: m}, nil
}

// Up выполняет миграции до последней версии.
//
// Примечание:
//   - Если миграций нет, функция завершается без ошибки.
//   - Если миграция уже выполнена, функция завершается без ошибки.
//   - Если возникает ошибка при выполнении миграций, функция возвращает ошибку.
func (m *Migrate) Up() error {
	err := m.Migrate.Up()
	if err != nil && !(err == migrate.ErrNoChange || (strings.Contains(err.Error(), "no migration found") && (strings.Contains(err.Error(), "read down")))) {
		_ = m.ResetDirtyVersion("up")
		return fmt.Errorf("error migrate up: %v", err)
	}

	return nil
}

// Down откатывает миграции до предыдущей версии.
//
// Примечание:
//   - Если миграций нет, функция завершается без ошибки.
//   - Если миграция уже откатана, функция завершается без ошибки.
//   - Если возникает ошибка при выполнении отката миграций, функция возвращает ошибку.
func (m *Migrate) Down() error {
	err := m.Migrate.Down()
	if err != nil && !(err == migrate.ErrNoChange || (strings.Contains(err.Error(), "no migration found") && (strings.Contains(err.Error(), "read down")))) {
		if err := m.ResetDirtyVersion("down"); err != nil {
			logger.Warn(err)
		}
		return fmt.Errorf("error migrate down: %v", err)
	}

	return nil
}

// MigrateToVersion мигрирует базу данных до заданной версии.
//
// Аргументы:
//   - version: Целевая версия миграции.
//
// Примечание:
//   - Если миграций нет, функция завершается без ошибки.
//   - Если миграция уже выполнена, функция завершается без ошибки.
//   - Если возникает ошибка при выполнении миграции, функция возвращает ошибку.
func (m *Migrate) MigrateToVersion(version int) error {
	vCurrent, _, err := m.Migrate.Version()
	if err != nil {
		return fmt.Errorf("failed to get current version: %v", err)
	}

	err = m.Migrate.Migrate(uint(version))
	if err != nil && !(err == migrate.ErrNoChange || (strings.Contains(err.Error(), "no migration found") && (strings.Contains(err.Error(), "read down")))) {
		if vCurrent > uint(version) {
			_ = m.ResetDirtyVersion("down")
		} else if vCurrent < uint(version) {
			_ = m.ResetDirtyVersion("up")
		}
		return fmt.Errorf("error migrate to version: %v", err)
	}

	return nil
}

// SetVersion устанавливает версию миграции,
// сбрасывая грязное состояние базы данных
//
// Аргументы:
//   - version: Целевая версия миграции.
//
// Примечание:
//   - Эта функция устанавливает версию миграции принудительно,
//     не выполняя миграционные операции.
func (m *Migrate) SetVersion(version int) error {
	if err := m.Migrate.Force(version); err != nil {
		return fmt.Errorf("error set version: %v", err)
	}

	return nil
}

// ResetDirtyVersion сбрасывает "грязную" версию миграции.
//
// Примечание:
//   - "Грязная" версия возникает, когда миграция была прервана
//     или завершилась с ошибкой.
//   - Эта функция сбрасывает версию до предыдущей версии,
//     если она была "грязной".
func (m *Migrate) ResetDirtyVersion(cmd string) error {
	version, dirty, err := m.Version()
	if err != nil {
		return err
	}
	logger.Info("мы тут", cmd)
	logger.Info(dirty)

	if dirty {
		logger.Info("мы тут", cmd)
		if cmd == "up" {
			logger.Info("upp")
			if err := m.SetVersion(int(version) - 1); err != nil {
				return fmt.Errorf("error reset dirty version: %v", err)
			}
		} else if cmd == "down" {
			logger.Info("downn")
			logger.Info(int(version) + 1)
			if err := m.SetVersion(int(version) + 1); err != nil {
				return fmt.Errorf("error reset dirty version: %v", err)
			}
			logger.Info(int(version) + 1)
		}
	}

	return nil
}

// Version возвращает текущую версию миграции.
//
// Возвращаемое значение:
//   - version: Текущая версия миграции.
//   - dirty:  Флаг, указывающий, была ли миграция прервана
//     или завершилась с ошибкой.
//   - err:    Ошибка, если произошла ошибка при получении версии.
func (m *Migrate) Version() (version int, dirty bool, err error) {
	versionUint, dirty, err := m.Migrate.Version()
	if !errors.Is(err, migrate.ErrNilVersion) && err != nil {
		return 0, false, fmt.Errorf("error get version: %v", err)
	}

	return int(versionUint), dirty, nil
}
