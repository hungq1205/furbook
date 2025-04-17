package repository

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

type PostRepository struct {
	postCollection *mongo.Collection
}

func NewPostRepository(db *mongo.Database) *PostRepository {
	return &PostRepository{postCollection: db.Collection("post")}
}

// Post

func (p *PostRepository) GetPost(ctx context.Context, id string) (*entity.Post, error) {
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

func (p *PostRepository) GetPostsOfUser(ctx context.Context, username string, pagination util.Pagination) ([]*entity.Post, error) {
	var posts []*entity.Post

	opts := options.Find().SetSort(bson.D{{"createdAt", -1}}).SetSkip(pagination.Offset()).SetLimit(pagination.Size)
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

func (p *PostRepository) GetPostsOfUsers(ctx context.Context, usernames []string, pagination util.Pagination) ([]*entity.Post, error) {
	var posts []*entity.Post

	sortOpts := options.Find().SetSort(bson.D{{"createdAt", -1}}).SetSkip(pagination.Offset()).SetLimit(pagination.Size)
	cursor, err := p.postCollection.Find(ctx, bson.M{"username": bson.M{"$in": usernames}}, sortOpts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *PostRepository) GetMediasOfPost(ctx context.Context, postID string) ([]entity.Media, error) {
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

func (p *PostRepository) CheckOwnership(ctx context.Context, postID string, username string) (bool, error) {
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

func (p *PostRepository) CreateBlogPost(ctx context.Context, username, content string, medias []entity.Media) (*entity.Post, error) {
	post := entity.Post{
		ID:           primitive.NewObjectID(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Comments:     []entity.Comment{},
		Interactions: []entity.Interaction{},
		Username:     username,
		Content:      content,
		Medias:       medias,
	}

	_, err := p.postCollection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *PostRepository) CreateLostPetPost(ctx context.Context, username, content string, medias []entity.Media, petId int, area, lastSeen *entity.Location, lostAt *time.Time) (*entity.Post, error) {
	post := entity.Post{
		ID:           primitive.NewObjectID(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Comments:     []entity.Comment{},
		Interactions: []entity.Interaction{},
		Username:     username,
		Content:      content,
		Medias:       medias,
		PetId:        &petId,
		Area:         area,
		LastSeen:     lastSeen,
		LostAt:       lostAt,
		Found:        false,
	}

	_, err := p.postCollection.InsertOne(ctx, post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (p *PostRepository) PatchContent(ctx context.Context, id, content string, medias []entity.Media) (bool, error) {
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

func (p *PostRepository) PatchFound(ctx context.Context, id string, found bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.postCollection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"found": found})
	return err
}

func (p *PostRepository) DeletePost(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = p.postCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}

// Comment

func (p *PostRepository) GetComments(ctx context.Context, postId string) ([]entity.Comment, error) {
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

func (p *PostRepository) CreateComment(ctx context.Context, postId, username, content string) error {
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

func (p *PostRepository) DeleteComment(ctx context.Context, postId, username string) error {
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

func (p *PostRepository) UpsertInteraction(ctx context.Context, postId, username string, itype entity.InteractionType) error {
	postOID, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}
	exists, err := p.updateInteraction(ctx, postOID, itype)
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

func (p *PostRepository) updateInteraction(ctx context.Context, id primitive.ObjectID, itype entity.InteractionType) (bool, error) {
	result, err := p.postCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{
		"$push": bson.M{
			"interactions": bson.M{
				"type": itype,
			},
		},
	})
	if err != nil {
		return false, err
	}
	return result.MatchedCount > 0, nil
}

func (p *PostRepository) DeleteInteraction(ctx context.Context, postId, username string) error {
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
