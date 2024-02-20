package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PostListResponse 
type PostListResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]any
    // The posts property
    posts []Postable
}
// NewPostListResponse instantiates a new PostListResponse and sets the default values.
func NewPostListResponse()(*PostListResponse) {
    m := &PostListResponse{
    }
    m.SetAdditionalData(make(map[string]any))
    return m
}
// CreatePostListResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePostListResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPostListResponse(), nil
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostListResponse) GetAdditionalData()(map[string]any) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PostListResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["posts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreatePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Postable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(Postable)
                }
            }
            m.SetPosts(res)
        }
        return nil
    }
    return res
}
// GetPosts gets the posts property value. The posts property
func (m *PostListResponse) GetPosts()([]Postable) {
    return m.posts
}
// Serialize serializes information the current object
func (m *PostListResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetPosts() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetPosts()))
        for i, v := range m.GetPosts() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err := writer.WriteCollectionOfObjectValues("posts", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PostListResponse) SetAdditionalData(value map[string]any)() {
    m.additionalData = value
}
// SetPosts sets the posts property value. The posts property
func (m *PostListResponse) SetPosts(value []Postable)() {
    m.posts = value
}
// PostListResponseable 
type PostListResponseable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPosts()([]Postable)
    SetPosts(value []Postable)()
}
