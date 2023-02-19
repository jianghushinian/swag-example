package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jianghushinian/swag-example/pkg/httputil"
)

// PingExample godoc
//
//	@Summary		ping example
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Response		200	{string}	string	"pong"
//	@Router			/examples/ping [get]
func (c *Controller) PingExample(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// CalcExample godoc
//
//	@Summary		calc example
//	@Description	plus
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Param			val1	query		int		true	"used for calc"
//	@Param			val2	query		int		true	"used for calc"
//	@Success		200		{string}	string	"ok"
//	@Failure		400		{string}	string	"fail"
//	@Failure		404		{string}	string	"fail"
//	@Failure		500		{string}	string	"fail"
//	@Router			/examples/calc [get]
func (c *Controller) CalcExample(ctx *gin.Context) {
	val1, err := strconv.Atoi(ctx.Query("val1"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	val2, err := strconv.Atoi(ctx.Query("val2"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ans := val1 + val2
	ctx.String(http.StatusOK, "%d", ans)
}

// PathParamsExample godoc
//
//	@Summary		path params example
//	@Description	path params
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Param			group_id	path		int		true	"Group ID"
//	@Param			account_id	path		int		true	"Account ID"
//	@Success		200			{string}	string	"ok"
//	@Failure		400			{string}	string	"fail"
//	@Failure		404			{string}	string	"fail"
//	@Failure		500			{string}	string	"fail"
//	@Router			/examples/groups/{group_id}/accounts/{account_id} [get]
func (c *Controller) PathParamsExample(ctx *gin.Context) {
	groupID, err := strconv.Atoi(ctx.Param("group_id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	accountID, err := strconv.Atoi(ctx.Param("account_id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.String(http.StatusOK, "group_id=%d account_id=%d", groupID, accountID)
}

// HeaderExample godoc
//
//	@Summary		custom header example
//	@Description	custom header
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Param			Authorization	header		string	true	"Authentication header"
//	@Success		200				{string}	string	"ok"
//	@Failure		400				{string}	string	"fail"
//	@Failure		404				{string}	string	"fail"
//	@Failure		500				{string}	string	"fail"
//	@Router			/examples/header [get]
func (c *Controller) HeaderExample(ctx *gin.Context) {
	ctx.String(http.StatusOK, ctx.GetHeader("Authorization"))
}

// SecurityBasicAuthExample godoc
//
//	@Summary		security basic auth example
//	@Description	custom header
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"ok"
//	@Failure		400	{string}	string	"fail"
//	@Failure		404	{string}	string	"fail"
//	@Failure		500	{string}	string	"fail"
//	@Security		BasicAuth
//	@Router			/examples/security/basicauth [post]
func (c *Controller) SecurityBasicAuthExample(ctx *gin.Context) {
	user := ctx.MustGet(gin.AuthUserKey).(string)
	ctx.String(http.StatusOK, "basic auth user: %s", user)
}

// AttributeExample godoc
//
//	@Summary		attribute example
//	@Description	attribute
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Param			enumstring	query		string	false	"string enums"		Enums(A, B, C)
//	@Param			enumint		query		int		false	"int enums"			Enums(1, 2, 3)
//	@Param			enumnumber	query		number	false	"int enums"			Enums(1.1, 1.2, 1.3)
//	@Param			string		query		string	false	"string valid"		minlength(5)	maxlength(10)
//	@Param			int			query		int		false	"int valid"			minimum(1)		maximum(10)
//	@Param			default		query		string	false	"string default"	default(A)
//	@Success		200			{string}	string	"ok"
//	@Failure		400			{string}	string	"fail"
//	@Failure		404			{string}	string	"fail"
//	@Failure		500			{string}	string	"fail"
//	@Router			/examples/attribute [get]
func (c *Controller) AttributeExample(ctx *gin.Context) {
	ctx.String(http.StatusOK, fmt.Sprintf("enumstring=%s enumint=%s enumnumber=%s string=%s int=%s default=%s",
		ctx.Query("enumstring"),
		ctx.Query("enumint"),
		ctx.Query("enumnumber"),
		ctx.Query("string"),
		ctx.Query("int"),
		ctx.Query("default"),
	))
}

// PostExample godoc
//
//	@Summary		post request example
//	@Description	post request example
//	@Tags			example
//	@Accept			json
//	@Produce		plain
//	@Param			message	body		model.Account	true	"Account Info"
//	@Success		200		{string}	string			"success"
//	@Failure		500		{string}	string			"fail"
//	@Router			/examples/post [post]
func (c *Controller) PostExample(ctx *gin.Context) {
	// do something
	ctx.JSON(http.StatusOK, nil)
}
