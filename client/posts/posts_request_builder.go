package posts

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703 "github.com/pistatium/kiota_example/client/models"
)

// PostsRequestBuilder builds and executes requests for operations under \posts
type PostsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PostsRequestBuilderGetQueryParameters get posts
type PostsRequestBuilderGetQueryParameters struct {
    // limit
    Limit *int32 `uriparametername:"limit"`
    // offset
    Offset *int32 `uriparametername:"offset"`
}
// PostsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PostsRequestBuilderGetQueryParameters
}
// PostsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PostsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPostsRequestBuilderInternal instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostsRequestBuilder) {
    m := &PostsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/posts{?limit*,offset*}", pathParameters),
    }
    return m
}
// NewPostsRequestBuilder instantiates a new PostsRequestBuilder and sets the default values.
func NewPostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PostsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPostsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get posts
func (m *PostsRequestBuilder) Get(ctx context.Context, requestConfiguration *PostsRequestBuilderGetRequestConfiguration)(i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.PostListResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "500": i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.CreateErrResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.CreatePostListResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.PostListResponseable), nil
}
// Post add post
func (m *PostsRequestBuilder) Post(ctx context.Context, body i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.Postable, requestConfiguration *PostsRequestBuilderPostRequestConfiguration)(i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.Postable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "400": i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.CreateErrResponseFromDiscriminatorValue,
        "500": i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.CreateErrResponseFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.CreatePostFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.Postable), nil
}
// ToGetRequestInformation get posts
func (m *PostsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PostsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation add post
func (m *PostsRequestBuilder) ToPostRequestInformation(ctx context.Context, body i8b557d3a2c13771573de98180b4f1b333ef2258a4830a42c4f070fa32a07f703.Postable, requestConfiguration *PostsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *PostsRequestBuilder) WithUrl(rawUrl string)(*PostsRequestBuilder) {
    return NewPostsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
