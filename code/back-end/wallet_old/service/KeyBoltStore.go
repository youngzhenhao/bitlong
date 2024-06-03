package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wallet/base"
	"io"
	"path/filepath"
	"time"
)

const passphrase = "wlt1F8@6Lz#dX9rTq4!Ko1Mv"

type KeyInfo struct {
	ID         string `json:"id"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

type KeyStore struct {
	DB *sql.DB
}

func InitDB() (*sql.DB, error) {
	dbPath := filepath.Join(base.QueryConfigByKey("dirpath"), "keyInfo.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	createStmt := `
    CREATE TABLE IF NOT EXISTS keys (
        id TEXT PRIMARY KEY,
        private_key TEXT,
        public_key TEXT
    );`
	_, err = db.Exec(createStmt)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *KeyStore) CreateOrUpdateKey(i *KeyInfo) error {
	encryptedKey, err := encrypt([]byte(i.PrivateKey), passphrase)
	if err != nil {
		return err
	}
	i.PrivateKey = string(encryptedKey)

	_, err = s.DB.Exec("REPLACE INTO keys (id, private_key, public_key) VALUES (?, ?, ?)", i.ID, i.PrivateKey, i.PublicKey)
	return err
}

func (s *KeyStore) ReadKey(ID string) (*KeyInfo, error) {
	var i KeyInfo

	// 使用 SQL 的 SELECT 语句从 keys 表中查询特定 ID 的密钥信息
	err := s.DB.QueryRow("SELECT id, private_key, public_key FROM keys WHERE id = ?", ID).Scan(&i.ID, &i.PrivateKey, &i.PublicKey)
	if err != nil {
		if err == sql.ErrNoRows {
			// 没有找到记录
			return nil, fmt.Errorf("no key found with ID: %s", ID)
		}
		// 数据库查询出错
		return nil, err
	}
	privateKeyEncrypted := []byte(i.PrivateKey)
	// 对查询到的加密私钥进行解密
	decryptedKey, err := decrypt(privateKeyEncrypted, passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt private key: %v", err)
	}
	i.PrivateKey = string(decryptedKey)

	return &i, nil
}

func getTimeNow() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

// Encrypt data using AES
func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return ciphertext, nil
}

// Decrypt data using AES
func decrypt(data []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {

		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)
	return data, nil
}
