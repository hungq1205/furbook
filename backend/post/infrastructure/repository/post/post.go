package post

import (
	"context"
	"post/entity"
	"post/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	postCollection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{postCollection: db.Collection("post")}
}

func (p *Repository) GetPost(ctx context.Context, id string) (*entity.Post, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var post entity.Post
	opt := options.FindOne().SetProjection(bson.M{"comments": 0})
	err = p.postCollection.FindOne(ctx, bson.M{"_id": objectID}, opt).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Repository) GetNearLostPosts(ctx context.Context, latitude float64, longitude float64, pagination util.Pagination) ([]*entity.Post, error) {
	var posts []*entity.Post
	filter := bson.M{
		"type": bson.M{
			"$in": bson.A{"lost", "found"},
		},
		"lastSeen.location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{longitude, latitude},
				},
			},
		},
	}

	opts := options.Find().
		SetSkip(pagination.Offset()).
		SetLimit(pagination.Size)

	cursor, err := p.postCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Repository) GetPostsOfUser(ctx context.Context, username string, pagination util.Pagination) ([]*entity.Post, error) {
	var posts []*entity.Post

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(pagination.Offset()).SetLimit(pagination.Size)
	cursor, err := p.postCollection.Find(ctx, bson.M{"username": username}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Repository) GetPostsOfUsers(ctx context.Context, usernames []string, pagination util.Pagination) ([]*entity.Post, error) {
	var posts []*entity.Post

	opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: -1}}).SetSkip(pagination.Offset()).SetLimit(pagination.Size)
	cursor, err := p.postCollection.Find(ctx, bson.M{"username": bson.M{"$in": usernames}}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Repository) GetMediasOfPost(ctx context.Context, postID string) ([]entity.Media, error) {
	objectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return nil, err
	}

	var medias []entity.Media
	err = p.postCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&medias)
	if err != nil {
		return nil, err
	}
	return medias, nil
}

func (p *Repository) CheckOwnership(ctx context.Context, postID string, username string) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(postID)
	if err != nil {
		return false, err
	}

	count, err := p.postCollection.CountDocuments(ctx, bson.M{"_id": objectID, "username": username})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p *Repository) CreateBlogPost(ctx context.Context, username, content string, medias []entity.Media) (*entity.Post, error) {
	post := entity.Post{
		ID:           primitive.NewObjectID(),
		Username:     username,
		Type:         entity.Blog,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Comments:     []entity.Comment{},
		Interactions: []entity.Interaction{},
		Content:      content,
		Medias:       medias,
	}

	_, err := p.postCollection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Repository) CreateLostPetPost(ctx context.Context, username string, contactInfo string, postType entity.PostType, content string, medias []entity.Media, area, lastSeen *entity.Location, lostAt *time.Time) (*entity.Post, error) {
	var post entity.Post
	if postType == entity.Found {
		post = entity.Post{
			ID:           primitive.NewObjectID(),
			Username:     username,
			Type:         postType,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Comments:     []entity.Comment{},
			Interactions: []entity.Interaction{},
			Content:      content,
			Medias:       medias,

			Participants: []string{},
			ContactInfo:  contactInfo,
			LastSeen:     lastSeen,
			LostAt:       lostAt,
			IsResolved:   false,
		}
	} else {
		post = entity.Post{
			ID:           primitive.NewObjectID(),
			Username:     username,
			Type:         postType,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Comments:     []entity.Comment{},
			Interactions: []entity.Interaction{},
			Content:      content,
			Medias:       medias,

			Participants: []string{},
			ContactInfo:  contactInfo,
			Area:         area,
			LastSeen:     lastSeen,
			LostAt:       lostAt,
			IsResolved:   false,
		}
	}

	_, err := p.postCollection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *Repository) PatchContent(ctx context.Context, id, content string, medias []entity.Media) (bool, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	result, err := p.postCollection.UpdateOne(ctx, bson.M{"_id": objectID},
		bson.M{
			"$set": bson.M{
				"content":    content,
				"medias":     medias,
				"updated_at": time.Now(),
			},
		})
	if err != nil {
		return false, err
	}

	return result.MatchedCount == 0, nil
}

func (p *Repository) PatchFound(ctx context.Context, id string, found bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"found": found})
	return err
}

func (p *Repository) DeletePost(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.postCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// Comment

func (p *Repository) GetComments(ctx context.Context, postId string) ([]entity.Comment, error) {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	var result struct {
		Comments []entity.Comment `bson:"comments"`
	}
	err = p.postCollection.FindOne(ctx, bson.M{"_id": postOID}, options.FindOne().SetProjection(bson.M{"comments": 1})).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result.Comments, nil
}

func (p *Repository) CreateComment(ctx context.Context, postId, username, content string) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": postOID}, bson.M{
		"$push": bson.M{
			"comments": bson.M{
				"username":  username,
				"content":   content,
				"createdAt": time.Now(),
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *Repository) DeleteComment(ctx context.Context, postId, username string) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": postOID}, bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"username": username,
			},
		},
	})
	return err
}

// Interaction

func (p *Repository) UpsertInteraction(ctx context.Context, postId, username string, itype entity.InteractionType) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	exists, err := p.updateInteraction(ctx, postOID, username, itype)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": postOID}, bson.M{
		"$push": bson.M{
			"interactions": bson.M{
				"username": username,
				"type":     itype,
			},
		},
	})
	return err
}

func (p *Repository) updateInteraction(ctx context.Context, id primitive.ObjectID, username string, itype entity.InteractionType) (bool, error) {
	result, err := p.postCollection.UpdateOne(ctx, bson.M{"_id": id, "interactions.username": username}, bson.M{
		"$set": bson.M{"interactions.$.type": itype},
	})
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

func (p *Repository) DeleteInteraction(ctx context.Context, postId, username string) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": postOID}, bson.M{
		"$pull": bson.M{
			"interactions": bson.M{"username": username},
		},
	})

	return err
}

func (p *Repository) Participate(ctx context.Context, postId, username string) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": postOID}, bson.M{
		"$addToSet": bson.M{"participants": username},
	})
	return err
}

func (p *Repository) Unparticipate(ctx context.Context, postId, username string) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": postOID}, bson.M{
		"$pull": bson.M{"participants": username},
	})
	return err
}
