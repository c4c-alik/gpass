package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"gpass/internal/database" // 引入内部包路径

	"github.com/wailsapp/wails/v2"
)

//go:embed assets/*
var assets embed.FS

// App struct
type App struct {
	ctx             context.Context
	db              *sql.DB
	passwordService *database.PasswordService
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// GetAllPasswords retrieves all passwords from the database
func (a *App) GetAllPasswords() ([]*database.Password, error) {
	return a.passwordService.GetAllPasswords()
}

// CreatePassword creates a new password entry
func (a *App) CreatePassword(account, username, password string) error {
	return a.passwordService.CreatePassword(account, username, password)
}

// UpdatePassword updates an existing password entry
func (a *App) UpdatePassword(id int, account, username, password string) error {
	return a.passwordService.UpdatePassword(id, account, username, password)
}

// DeletePassword deletes a password entry
func (a *App) DeletePassword(id int) error {
	return a.passwordService.DeletePassword(id)
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app staadsfasdfrts. The context is saved so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// 初始化数据库连接
	a.initDB()
	// 初始化密码服务
	a.passwordService = database.NewPasswordService(a.db)
}

func (a *App) initDB() {
	var err error
	a.db, err = sql.Open("sqlite3", "./gpass.db")
	if err != nil {
		log.Fatal(err)
	}

	// 创建密码表
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS passwords (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		account TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = a.db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "gpass - Password Manager",
		Width:     1024,
		Height:    768,
		MinWidth:  800,
		MinHeight: 600,
		AssetServer: &assetserver.Options{
			Assets: Assets(),
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

// Assets returns the embedded assets
func Assets() fs.FS {
	return assets
}
