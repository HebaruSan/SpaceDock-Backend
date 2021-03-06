/*
 SpaceDock Backend
 API Backend for the SpaceDock infrastructure to host modfiles for various games

 SpaceDock Backend is licensed under the Terms of the MIT License.
 Copyright (c) 2017 Dorian Stoll (StollD), RockyTV
 */

package objects

import (
    "crypto/md5"
    "crypto/rand"
    "encoding/hex"
)

type Token struct {
    Model

    Token  string `gorm:"size:32" spacedock:"lock" json:"token"`
}

func NewToken() *Token {
    key := make([]byte, 64)
    _, err := rand.Read(key)
    if err != nil {
        panic(err) // aahhh
    }
    hash := md5.New()
    hash.Write(key)
    t := &Token{
        Token:string(hex.EncodeToString(hash.Sum(nil))),
    }
    t.Meta = "{}"
    return t
}