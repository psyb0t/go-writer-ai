package writer_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-writer-ai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testAPIKey = "test-key"

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *writer.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := writer.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)

			return nil
		},
	)

	client, err := writer.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func assertBearerAuth(t *testing.T, r *http.Request) {
	t.Helper()

	assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))
}

func TestModels(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/v1/models", r.URL.Path)
		assertBearerAuth(t, r)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"models":[{"id":"model-a","name":"Model A"}]}`,
		))
	})

	res, err := client.ModelsWithResponse(context.Background())
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Models, 1)

	assert.Equal(t, "model-a", res.JSON200.Models[0].Id)
	assert.Equal(t, "Model A", res.JSON200.Models[0].Name)
}

func TestChat(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/v1/chat", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assertBearerAuth(t, r)

		var body struct {
			Model    string `json:"model"`
			Messages []struct {
				Role    string `json:"role"`
				Content string `json:"content"`
			} `json:"messages"`
			Temperature float64 `json:"temperature"`
		}

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "model-a", body.Model)
		assert.InDelta(t, 0.25, body.Temperature, 0)
		assert.Len(t, body.Messages, 1)

		if len(body.Messages) == 1 {
			assert.Equal(t, "user", body.Messages[0].Role)
			assert.Equal(t, "hello", body.Messages[0].Content)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"00000000-0000-0000-0000-000000000001",` +
				`"object":"chat.completion","created":1700000000,` +
				`"model":"model-a","choices":[{"index":0,` +
				`"finish_reason":"stop","message":{"role":"assistant",` +
				`"content":"hi there","refusal":null}}]}`,
		))
	})

	var content writer.ChatMessage_Content
	require.NoError(t, content.FromChatMessageContent0("hello"))

	temperature := 0.25

	res, err := client.ChatWithResponse(
		context.Background(),
		writer.ChatRequest{
			Model: "model-a",
			Messages: []writer.ChatMessage{{
				Content: &content,
				Role:    writer.ChatMessageRequestRoleUser,
			}},
			Temperature: &temperature,
		},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)
	require.Len(t, res.JSON200.Choices, 1)

	assert.Equal(t,
		"00000000-0000-0000-0000-000000000001",
		res.JSON200.Id.String(),
	)
	assert.Equal(t, int64(1700000000), res.JSON200.Created)

	choice := res.JSON200.Choices[0]
	assert.Equal(t, writer.ChatCompletionFinishReasonStop, choice.FinishReason)
	assert.Equal(t, "hi there", choice.Message.Content)
	assert.Nil(t, choice.Message.Refusal)
}
