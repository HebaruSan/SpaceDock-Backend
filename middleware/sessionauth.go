/*
 SpaceDock Backend
 API Backend for the SpaceDock infrastructure to host modfiles for various games

 SpaceDock Backend is licensed under the Terms of the MIT License.
 Copyright (c) 2017 Dorian Stoll (StollD), RockyTV
 */

package middleware

import (
    "github.com/KSP-SpaceDock/SpaceDock-Backend/objects"
    "github.com/spf13/cast"
    "gopkg.in/kataras/iris.v6"
)

func LoginRequired(ctx *iris.Context) {
    userID, err := ctx.Session().GetInt("SessionID")
    if err == nil {
        var user objects.User
        err = user.GetById(cast.ToUint(userID))
        if !(err == nil && user.IsAuthenticated()) {
            ctx.SetStatusCode(iris.StatusUnauthorized)
            return
        }
    } else {
        ctx.SetStatusCode(iris.StatusUnauthorized)
        return
    }
    ctx.Next()
}

func LoginUser(ctx *iris.Context, user *objects.User) {
    user.Login()
    ctx.Session().Set("SessionID", cast.ToInt(user.ID))
}

func LogoutUser(ctx *iris.Context) {
    CurrentUser(ctx).Logout()
    ctx.Session().Delete("SessionID")
}

func CurrentUser(ctx *iris.Context) *objects.User {
    userID, err := ctx.Session().GetInt("SessionID")
    if err == nil {
        user := &objects.User{}
        err = user.GetById(cast.ToUint(userID))
        if err == nil {
            return user
        } else {
            return nil
        }
    }
    return nil
}
