package orm

import (
    "fmt"
    "strconv"

    "gopkg.in/pg.v4"

    "model"
)

const dbName = "postgre"
const Liked = "liked"
const Disliked = "disliked"
const Matched = "matched"

func FindUsers() [] model.User{
	db := pg.Connect(&pg.Options{
        User: dbName,
    })

    var users []model.User
	err := db.Model(&users).
    	Columns("user.*").
    	Limit(10).
    	Select()

    if err != nil {
    	fmt.Println("find user err:", err)
    }

    return users
}

func CreateUser(name string) model.User{
	db := pg.Connect(&pg.Options{
        User: dbName,
    })

    user := model.User{
    	Name : name,
    	Type : "users",
	}

	err := db.Create(&user)
	if err != nil {
		fmt.Println("create user err:", err)
	}
	return user
}

func FindRelationships(userId string) []model.Relationship{

	uidInt, _ := strconv.Atoi(userId)

	db := pg.Connect(&pg.Options{
        User: dbName,
    })

    var relationships []model.Relationship
	err := db.Model(&relationships).
		Where("user_id = ?", uidInt).
    	Columns("relationship.*").
    	Limit(10).
    	Select()

    if err != nil {
    	fmt.Println("find relationship err:", err)
    }

    return relationships
}


func UpdateRelationship(userId, otherId, state string) model.Relationship {
	uidInt, _ := strconv.Atoi(userId)
	oidInt, _ := strconv.Atoi(otherId)

	db := pg.Connect(&pg.Options{
        User: dbName,
    })

    var count int
	err := db.Model(model.Relationship{}).Where("user_id = ? and other_id = ?",userId, otherId).Count(&count)
	if err != nil {
    	panic(err)
	}

	relationship := model.Relationship{
    	UserId : uidInt,
    	OtherId : oidInt,
    	State : state,
    	Type : "relationship",
	}

	fmt.Println("count:", count)
	if count > 0 {
		//already have relationship
		_ = db.Model(&relationship).
			Where("user_id = ? and other_id = ?", userId, otherId).
    		Columns("relationship.*").
    		Limit(1).
    		Select()

    	relationship.State = state
    	_ = db.Update(relationship)	
    } else {
    	//no result, create relationship
		_ = db.Create(&relationship)
    }

    var countOther int
    _ = db.Model(model.Relationship{}).Where("user_id = ? and other_id = ?", otherId, userId).Count(&countOther)
    fmt.Println("countOther,", countOther)
	if countOther > 0 {
		//they have relationships, update them
		var otherRelationship model.Relationship

		_ = db.Model(&otherRelationship).
			Where("user_id = ? and other_id = ?", otherId, userId).
    		Columns("relationship.*").
    		Limit(1).
    		Select()

    	//liked each other, matched
    	if relationship.State == Liked && (otherRelationship.State == Liked || otherRelationship.State == Matched) {
    		relationship.State = Matched
    		otherRelationship.State = Matched

    		_ = db.Update(relationship)	
			_ = db.Update(otherRelationship)	
    	} else if relationship.State == Disliked && otherRelationship.State == Matched {
    		otherRelationship.State = Liked

			_ = db.Update(otherRelationship)	
    	} 
    }
	
    return relationship
}
