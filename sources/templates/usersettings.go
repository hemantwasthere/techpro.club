package templates

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"techpro.club/sources/common"
	"techpro.club/sources/users"
)

type SettingsStruct struct{
	UserProfile common.FetchUserStruct `json:"userprofile"`
	UserSocials common.FetchSocialStruct `json:"socials"`
}

// Display and edit user profile
func UserSettings(w http.ResponseWriter, r *http.Request) {

	// Check if the path is ok
	// Check if the session is ok

	// Fetch name and image from cookies
	// Fetch email, _id, location,repourl from database
	// Display the user's profile
	// Display the user's settings

	
	
	if r.URL.Path != "/users/settings" {
        ErrorHandler(w, r, http.StatusNotFound)
        return
    }
	
	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} 


	if r.Method == "POST"{
		// Update user profile

		errParse := r.ParseForm()
		if errParse != nil {
			log.Println(errParse.Error())
		} else {
			name := r.Form.Get("name")
			repoUrl := r.Form.Get("repourl")
			facebook := r.Form.Get("facebook")
			linkedin := r.Form.Get("linkedin")
			twitter := r.Form.Get("twitter")
			stackoverflow := r.Form.Get("stackoverflow")
			
			socialStatus, socialMsg := updateSocials(userID, facebook, linkedin, twitter, stackoverflow)
			profileStatus, profileMsg := UpdateUserProfile(userID, name, repoUrl)

			if (socialStatus && profileStatus){
				fmt.Println("ok",socialMsg, profileMsg)
			} else {
				fmt.Println("Wrong", socialMsg, profileMsg)
			}
		}
	}

	userprofile := fetchUserProfile(userID)
	socials := fetchSocials(userID)

	userSettingsStruct := SettingsStruct{userprofile, socials}

	tmpl, err := template.New("").ParseFiles("templates/app/settings.gohtml", "templates/app/contributors/common/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "contributorbase", userSettingsStruct) 
	}
}

// Fetch user profile
func fetchUserProfile(userID primitive.ObjectID)(userProfile common.FetchUserStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchUsers := client.Database(dbName).Collection(common.CONST_MO_USERS)
	err := fetchUsers.FindOne(context.TODO(),  bson.M{"_id": userID}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&userProfile)

	if err != nil {
		fmt.Println(err, userID)
	} 

	return userProfile
}

// Update user profile
func UpdateUserProfile(userID primitive.ObjectID, name, repoLink string)(status bool, msg string){

	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	updateUsers := client.Database(dbName).Collection(common.CONST_MO_USERS)
	_ , errUpdate := updateUsers.UpdateOne(context.TODO(), bson.M{"_id": userID}, bson.M{"$set": bson.M{"name": name, "repourl": repoLink}})

	if errUpdate != nil {
		status = false
		msg = errUpdate.Error()
	} else {
		status = true
		msg = "Success"
	}

	return status, msg
}

// Fetch socials
func fetchSocials(userID primitive.ObjectID)(socials common.FetchSocialStruct){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchSocials := client.Database(dbName).Collection(common.CONST_MO_SOCIALS)
	err := fetchSocials.FindOne(context.TODO(),  bson.M{"userid": userID}, options.FindOne().SetProjection(bson.M{"_id": 0})).Decode(&socials)

	if err != nil {
		fmt.Println(err, userID)
	} 

	return socials
}

// Update socials
func updateSocials(userID primitive.ObjectID, facebook, linkedin, twitter, stackoverflow string)(status bool, msg string){
	client, _ := common.Mongoconnect()
	defer client.Disconnect(context.TODO())

	dbName := common.GetMoDb()
	fetchSocials := client.Database(dbName).Collection(common.CONST_MO_SOCIALS)
	countUsers, errSearch := fetchSocials.CountDocuments(context.TODO(),  bson.M{"userid": userID})

	if errSearch != nil {
		status = false
		msg = errSearch.Error()
	} else {
		if (countUsers == 0){
			// Insert
			insertResult, errInsert := fetchSocials.InsertOne(context.TODO(), bson.M{"userid": userID, "facebook": facebook, "linkedin": linkedin, "twitter": twitter, "stackoverflow": stackoverflow})
			if errInsert != nil {
				status = false
				msg = errInsert.Error()
			} else {
				status = true
				msg = insertResult.InsertedID.(primitive.ObjectID).Hex()
			}

		} else {
			// Update
			_, errUpdate := fetchSocials.UpdateOne(context.TODO(),  bson.M{"userid": userID}, bson.M{"$set": bson.M{"facebook": facebook, "linkedin": linkedin, "twitter": twitter, "stackoverflow": stackoverflow}})
			if errUpdate != nil {
				status = false
				msg = errUpdate.Error()
			} else {
				status = true
				msg = "Success"
			}
		}
	}

	return status, msg
}