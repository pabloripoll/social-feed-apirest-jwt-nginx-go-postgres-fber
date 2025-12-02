package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"time"

	"example.com/fiber-pg-rmq-jwt/internal/db"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"github.com/jackc/pgx/v5/pgxpool"
	"example.com/fiber-pg-rmq-jwt/internal/mq"
)

type Handler struct {
	pool *pgxpool.Pool
	rmq  *mq.MQ
}

func NewHandler(pool *pgxpool.Pool, rmq *mq.MQ) *Handler {
	return &Handler{pool: pool, rmq: rmq}
}

func (h *Handler) RegisterRoutes(app *fiber.App) {
	app.Post("/signup", h.SignUp)
	app.Post("/login", h.Login)

	// JWT protected group
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret" // fallback for quick testing; don't use in prod
	}
	app.Group("/", jwtware.New(jwtware.Config{
		SigningKey:   []byte(jwtSecret),
		ContextKey:   "user", // store token in ctx.Locals("user")
		AuthScheme:   "Bearer",
		TokenLookup:  "header:Authorization",
		ErrorHandler: jwtError,
	})).Get("/profile", h.Profile).
		Post("/publish", h.Publish)
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"error": "invalid or missing token",
	})
}

// SignUp creates a new user with bcrypt hashed password
func (h *Handler) SignUp(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.Email == "" || req.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "email and password required"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "unable to hash password"})
	}
	id, err := db.CreateUser(context.Background(), h.pool, req.Email, string(hash))
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "email already exists"})
	}
	return c.JSON(fiber.Map{"id": id})
}

// Login validates credentials and returns a signed JWT
func (h *Handler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	u, err := db.GetUserByEmail(context.Background(), h.pool, req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid credentials"})
	}

	// Create JWT
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev-secret"
	}
	ttlMinutes := 60
	if v := os.Getenv("JWT_TTL_MINUTES"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			ttlMinutes = n
		}
	}
	exp := time.Now().Add(time.Duration(ttlMinutes) * time.Minute)
	claims := jwt.MapClaims{
		"sub": u.ID,
		"exp": jwt.NewNumericDate(exp),
		"iat": jwt.NewNumericDate(time.Now()),
		"email": u.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not sign token"})
	}
	return c.JSON(fiber.Map{"token": s, "expires_at": exp})
}

// Profile returns the current user's info using the token sub claim
func (h *Handler) Profile(c *fiber.Ctx) error {
	// token has been verified by middleware, extract claims
	user := c.Locals("user")
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "no token data"})
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	sub := claims["sub"]
	var id int64
	switch v := sub.(type) {
	case float64:
		id = int64(v)
	case int64:
		id = v
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid token sub"})
	}
	u, err := db.GetUserByID(context.Background(), h.pool, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(fiber.Map{"id": u.ID, "email": u.Email})
}

func (h *Handler) Publish(c *fiber.Ctx) error {
	var req struct {
		Topic   string `json:"topic"`
		Payload any    `json:"payload"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}
	if req.Topic == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "topic required"})
	}
	b, err := json.Marshal(req.Payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid payload"})
	}
	// For example publish to queue named req.Topic
	if err := h.rmq.Publish("", req.Topic, b); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "publish failed"})
	}
	return c.JSON(fiber.Map{"ok": true})
}