package offer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/personal_proj/hotel_management/datasources/mysql/offer_db"
	"github.com/personal_proj/hotel_management/utils/errors"
)

const (
	queryInsertOfferTable    = "INSERT INTO offer(cm_offer_id,original_data,capacity,number,price,currency,check_in,check_out,fees) VALUES(?,?,?,?,?,?,?,?,?);"
	queryInsertHotelTable    = "INSERT INTO hotel(cm_offer_id,hotel_id,name,country,address,latitude,longitude,telephone,amenities,description,room_count,currency) VALUES(?,?,?,?,?,?,?,?,?,?,?,?);"
	queryInsertRoomTable     = "INSERT INTO room(cm_offer_id,hotel_id,room_id,description,name,capacity) VALUES(?,?,?,?,?,?);"
	queryInsertRatePlanTable = "INSERT INTO rate_plan(cm_offer_id,hotel_id,rate_plan_id,cancellation_policy,name,other_conditions,meal_plan) VALUES(?,?,?,?,?,?,?);"
)

//Save ...
func Save(offer map[string]interface{}) *errors.RestErr {
	var offr Offer
	b, err := json.Marshal(offer)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	if err := json.Unmarshal(b, &offr); err != nil {
		fmt.Println(err)
	}
	ctx := context.Background()
	tx, err := offer_db.Client.BeginTx(ctx, nil)
	if err != nil {
		return errors.NewInternalServerError(err.Error())

	}
	//cm_offer_id,original_data,capacity,number,price,currency,check_in,check_out,fees
	fmt.Println("==================")
	fmt.Println(offr)
	fmt.Println("==================")
	//fmt.Println(offr.OfferDetails)
	for i := 0; i < len(offr.Offers); i++ {
		originalDataBytes, err := json.Marshal(offr.Offers[i].OriginalData)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}

		capacityBytes, err := json.Marshal(offr.Offers[i].Capacity)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}

		feesBytes, err := json.Marshal(offr.Offers[i].Fees)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}

		_, err = tx.ExecContext(ctx, queryInsertOfferTable, offr.Offers[i].CmOfferID, originalDataBytes,
			capacityBytes, offr.Offers[i].Number, offr.Offers[i].Price, offr.Offers[i].Currency,
			offr.Offers[i].CheckIn, offr.Offers[i].CheckOut, feesBytes)

		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		//cm_offer_id,hotel_id,name,country,address,latitude,longitude,telephone,amenities,description,room_count,currency

		amenitiesBytes, err := json.Marshal(offr.Offers[i].Hotel.Amenities)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		_, err = tx.ExecContext(ctx, queryInsertHotelTable, offr.Offers[i].CmOfferID, offr.Offers[i].Hotel.HotelId,
			offr.Offers[i].Hotel.Name, offr.Offers[i].Hotel.Country, offr.Offers[i].Hotel.Address,
			offr.Offers[i].Hotel.Latitude, offr.Offers[i].Hotel.Longitude, offr.Offers[i].Hotel.Telephone,
			amenitiesBytes, offr.Offers[i].Hotel.Description, offr.Offers[i].Hotel.RoomCount,
			offr.Offers[i].Hotel.Currency)

		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		// 	//cm_offer_id,hotel_id,rate_plan_id,cancellation_policy,name,other_conditions,meal_plan
		cancellationPolicyBytes, err := json.Marshal(offr.Offers[i].RatePlan.CancellationPolicy)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}

		otherConditionBytes, err := json.Marshal(offr.Offers[i].RatePlan.OtherConditions)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}

		_, err = tx.ExecContext(ctx, queryInsertRatePlanTable, offr.Offers[i].CmOfferID, offr.Offers[i].RatePlan.HotelId,
			offr.Offers[i].RatePlan.RatePlanId, cancellationPolicyBytes,
			offr.Offers[i].RatePlan.Name, otherConditionBytes, offr.Offers[i].RatePlan.MealPlan)

		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
		//cm_offer_id,hotel_id,room_id,description,name,capacity

		capacityBytes, err = json.Marshal(offr.Offers[i].Room.Capacity)
		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}

		_, err = tx.ExecContext(ctx, queryInsertRoomTable, offr.Offers[i].CmOfferID, offr.Offers[i].Room.HotelId,
			offr.Offers[i].Room.RroomId, offr.Offers[i].Room.Description, offr.Offers[i].Room.Name,
			capacityBytes)

		if err != nil {
			return errors.NewInternalServerError(err.Error())
		}
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
