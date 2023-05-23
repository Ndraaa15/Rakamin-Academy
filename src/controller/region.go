package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rakamin-academy/sdk/message"

	"github.com/gin-gonic/gin"
)

func (c *Controller) GetAllProvinces(ctx *gin.Context) {

	provinces, err := http.Get("https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json")

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all provinces from API", err.Error())
		return
	}

	defer provinces.Body.Close()

	if provinces.StatusCode != http.StatusOK {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all provinces from API", err.Error())
		return
	}

	var datas interface{}
	if err := json.NewDecoder(provinces.Body).Decode(&datas); err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to decode all provinces from API", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch all province from API", &datas)

}

func (c *Controller) GetProvinceByID(ctx *gin.Context) {

	pID := ctx.Param("provinceID")

	province, err := http.Get(fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/provinces/%s.json", pID))

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch  province from API", err.Error())
		return
	}

	defer province.Body.Close()

	if province.StatusCode != http.StatusOK {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all province from API", err.Error())
		return
	}

	var data interface{}
	if err := json.NewDecoder(province.Body).Decode(&data); err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to decode all province from API", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch all province from API", &data)

}

func (c *Controller) GetAllCitiesByProvinceID(ctx *gin.Context) {

	pID := ctx.Param("provinceID")

	cities, err := http.Get(fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regencies/%s.json", pID))

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all cities from API", err.Error())
		return
	}

	defer cities.Body.Close()

	if cities.StatusCode != http.StatusOK {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all cities from API", err.Error())
		return
	}

	var datas interface{}
	if err := json.NewDecoder(cities.Body).Decode(&datas); err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to decode all cities from API", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch all cities from API", &datas)

}

func (c *Controller) GetDistrictsByCityID(ctx *gin.Context) {

	cID := ctx.Param("cityID")

	districts, err := http.Get(fmt.Sprintf("https: //emsifa.github.io/api-wilayah-indonesia/api/districts/%s.json", cID))

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all cities from API", err.Error())
		return
	}

	defer districts.Body.Close()

	if districts.StatusCode != http.StatusOK {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all cities from API", err.Error())
		return
	}

	var datas interface{}
	if err := json.NewDecoder(districts.Body).Decode(&datas); err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to decode all cities from API", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch all cities from API", &datas)

}

func (c *Controller) GetVillagesByDistrictsID(ctx *gin.Context) {

	dID := ctx.Param("districtID")

	villages, err := http.Get(fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/villages/%s.json", dID))

	if err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all cities from API", err.Error())
		return
	}

	defer villages.Body.Close()

	if villages.StatusCode != http.StatusOK {
		message.Error(ctx, http.StatusInternalServerError, "Failed to fetch all cities from API", err.Error())
		return
	}

	var datas interface{}
	if err := json.NewDecoder(villages.Body).Decode(&datas); err != nil {
		message.Error(ctx, http.StatusInternalServerError, "Failed to decode all cities from API", err.Error())
		return
	}

	message.Success(ctx, http.StatusOK, "Success to fetch all cities from API", &datas)

}
