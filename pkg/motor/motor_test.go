package motor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/sonr-io/sonr/pkg/client"
	"github.com/sonr-io/sonr/pkg/crypto"
	rt "github.com/sonr-io/sonr/x/registry/types"
	"github.com/stretchr/testify/assert"
	prt "go.buf.build/grpc/go/sonr-io/motor/api/v1"
)

func Test_CreateAccount(t *testing.T) {
	aesKey := loadKey("aes.key")
	if aesKey == nil || len(aesKey) != 32 {
		key, err := crypto.NewAesKey()
		assert.NoError(t, err, "generates aes key")
		aesKey = key

		// store the key
		fmt.Printf("stored key? %v\n", storeKey("aes.key", key))
	} else {
		fmt.Println("loaded key")
	}

	req := prt.CreateAccountRequest{
		Password:  "password123",
		AesDscKey: aesKey,
	}

	m := EmptyMotor("test_device")
	res, err := m.CreateAccount(req)
	assert.NoError(t, err, "wallet generation succeeds")

	// write the PSK to local file system for later use
	if err == nil {
		fmt.Printf("stored psk? %v\n", storeKey(fmt.Sprintf("psk%s", m.Address), res.AesPsk))
	}

	b := m.Balance()
	log.Println("balance:", b)

	// Print the address of the wallet
	log.Println("address:", m.Address)
}

func Test_Login(t *testing.T) {
	did := "snr1q34xcsdj9hp04akdvyz427w0eytxps7xy4gtkt"
	t.Run("with password", func(t *testing.T) {
		pskKey := loadKey(fmt.Sprintf("psk%s", did))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := prt.LoginRequest{
			Did:       did,
			Password:  "password123",
			AesPskKey: pskKey,
		}

		m := EmptyMotor("test_device")
		_, err := m.Login(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.Balance())
			fmt.Println("address: ", m.Address)
		}
	})

	t.Run("with DSC", func(t *testing.T) {
		aesKey := loadKey("aes.key")
		fmt.Printf("aes: %x\n", aesKey)
		if aesKey == nil || len(aesKey) != 32 {
			t.Errorf("could not load key.")
			return
		}

		pskKey := loadKey(fmt.Sprintf("psk%s", did))
		if pskKey == nil || len(pskKey) != 32 {
			t.Errorf("could not load psk key")
			return
		}

		req := prt.LoginRequest{
			Did:       did,
			AesDscKey: aesKey,
			AesPskKey: pskKey,
		}

		m := EmptyMotor("test_device")
		_, err := m.Login(req)
		assert.NoError(t, err, "login succeeds")

		if err == nil {
			fmt.Println("balance: ", m.Balance())
			fmt.Println("address: ", m.Address)
		}
	})
}

func Test_LoginAndMakeRequest(t *testing.T) {
	did := "snr1q34xcsdj9hp04akdvyz427w0eytxps7xy4gtkt"
	pskKey := loadKey(fmt.Sprintf("psk%s", did))
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := prt.LoginRequest{
		Did:       did,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor("test_device")
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// do something with the logged in account
	m.DIDDocument.AddAlias("gotest.snr")
	_, err = updateWhoIs(m)
	assert.NoError(t, err, "updates successfully")
}

func Test_CreateSchema(t *testing.T) {
	did := "snr1q34xcsdj9hp04akdvyz427w0eytxps7xy4gtkt"
	pskKey := loadKey(fmt.Sprintf("psk%s", did))
	fmt.Printf("psk: %x\n", pskKey)
	if pskKey == nil || len(pskKey) != 32 {
		t.Errorf("could not load psk key")
		return
	}

	req := prt.LoginRequest{
		Did:       did,
		Password:  "password123",
		AesPskKey: pskKey,
	}

	m := EmptyMotor("test_device")
	_, err := m.Login(req)
	assert.NoError(t, err, "login succeeds")

	// LOGIN DONE, TRY TO CREATE SCHEMA
	createSchemaRequest, err := json.Marshal(prt.CreateSchemaRequest{
		Label: "TestUser",
		Fields: map[string]prt.CreateSchemaRequest_SchemaKind{
			"email":     prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
			"firstName": prt.CreateSchemaRequest_SCHEMA_KIND_STRING,
			"age":       prt.CreateSchemaRequest_SCHEMA_KIND_INT,
		},
	})
	assert.NoError(t, err, "request marshals")
	resp, err := m.CreateSchema(createSchemaRequest)
	assert.NoError(t, err, "schema created successfully")

	fmt.Printf("success: %s\n", resp.Did)
}

func Test_DecodeTxData(t *testing.T) {
	data := "0A91010A242F736F6E72696F2E736F6E722E72656769737472792E4D736743726561746557686F497312691267122A736E723134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A35371A31122F6469643A736E723A3134373071366D3476776D6537346A376D3573326364773939357A35796E6B747A726D377A353730BC8FA197063801"

	mcr := &rt.MsgCreateWhoIsResponse{}
	err := client.DecodeTxResponseData(data, mcr)
	assert.NoError(t, err, "decodes tx data successfully")
	assert.Equal(t, "snr1470q6m4vwme74j7m5s2cdw995z5ynktzrm7z57", mcr.WhoIs.Owner)
}

func storeKey(n string, aesKey []byte) bool {
	name := fmt.Sprintf("./test_keys/%s", n)
	file, err := os.Create(name)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.Write(aesKey)
	return err == nil
}

func loadKey(n string) []byte {
	name := fmt.Sprintf("./test_keys/%s", n)
	var file *os.File
	if _, err := os.Stat(name); os.IsNotExist(err) {
		file, err = os.Create(name)
		if err != nil {
			return nil
		}
	} else if err != nil {
		fmt.Printf("load err: %s\n", err)
	} else {
		file, err = os.Open(name)
		if err != nil {
			return nil
		}
	}
	defer file.Close()

	key, err := ioutil.ReadAll(file)
	if err != nil {
		return nil
	}
	return key
}
