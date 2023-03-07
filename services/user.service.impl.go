package services

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"atm-machine.com/atm-apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"golang.org/x/text/number"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

// This should be in an env file in production
const MySecret string = "abc&1*~#^2^#s0^=)^^7!b34"

func Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// Encrypt method is to encrypt or hide any classified text
func Encrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

// Decrypt method is to extract back the encrypted text
func Decrypt(text, MySecret string) (string, error) {
	block, err := aes.NewCipher([]byte(MySecret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func generate_accno(low, hi int) int {
	return low + rand.Intn(hi-low)
}

func stringtoint(x string) int {
	number, e := strconv.Atoi(x)
	if e != nil {
		fmt.Println(e)
	}

	return number
}
func inttostring(x int) string {
	var num string = strconv.Itoa(x)
	return num
}

func stringaddsub(x1, x2 string) string {
	var rev string
	if strings.Contains(x1, "+") {
		var sum int = stringtoint(strings.ReplaceAll(x1, "+", "")) + stringtoint(x2)
		rev = inttostring(sum)
	} else if strings.Contains(x1, "-") {
		var sum int = stringtoint(x2) - stringtoint(strings.ReplaceAll(x1, "-", ""))
		rev = inttostring(sum)
	}
	return rev
}

func (u *UserServiceImpl) CreateUser(user *models.User) (int, error) {
	user.AccountNo = generate_accno(100000000, 999999999)
	user.Balance = "0"
	encText, er := Encrypt(user.Pin, MySecret)
	fmt.Println("encText", encText)
	user.Pin = encText
	if er != nil {
		fmt.Println("error encrypting your classified text: ", er)
	}
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return user.AccountNo, err
}

func (u *UserServiceImpl) DepositWithdraw(user []string) error {
	var users *models.User
	filter := bson.D{bson.E{Key: "user_no", Value: stringtoint(user[0])}}
	find := u.usercollection.FindOne(u.ctx, filter).Decode(&users)
	if find != nil {
		fmt.Println(find)
	}
	var val string = stringaddsub(user[1], users.Balance)
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_balance", Value: val}, bson.E{Key: "user_statement", Value: append(users.Statement, user[1]+" : "+time.Now().String())}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched Document found for update")
	}
	return nil
}

func (u *UserServiceImpl) ChangePin(user []string) error {
	filter := bson.D{bson.E{Key: "user_no", Value: stringtoint(user[0])}}
	encText, er := Encrypt(user[1], MySecret)
	if er != nil {
		fmt.Println("error encrypting your classified text: ", er)
	}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_pin", Value: stringtoint(encText)}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched Document found for update")
	}
	return nil
}

func (u *UserServiceImpl) GetTransacion(user string) (*models.User, error) {
	var users *models.User
	filter := bson.D{bson.E{Key: "user_no", Value: stringtoint(user)}}
	err := u.usercollection.FindOne(u.ctx, filter).Decode(&users)
	return users, err
}
