package hotelclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/k33pup/Booking/hotel_svc/pkg/models"
	"net/http"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func NewClient(baseURL string) *Client {
	return &Client{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetHotelsList() ([]models.Hotel, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/hotels", c.baseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hotels list: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Hotels []models.Hotel `json:"hotels"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Hotels, nil
}

// CreateHotel создает новый отель.
func (c *Client) CreateHotel(hotel models.Hotel) (*models.Hotel, error) {
	payload, err := json.Marshal(hotel)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal hotel: %w", err)
	}

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/hotels", c.baseURL), "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to create hotel: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Hotel models.Hotel `json:"hotel"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.Hotel, nil
}

// DeleteHotel удаляет отель по ID.
func (c *Client) DeleteHotel(hotelID string) (bool, error) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/hotels?hotelId=%s", c.baseURL, hotelID), nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("failed to delete hotel: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Success, nil
}

// GetHotelByID получает информацию о конкретном отеле по ID.
func (c *Client) GetHotelByID(hotelID string) (*models.Hotel, error) {
	resp, err := c.httpClient.Get(fmt.Sprintf("%s/hotels/%s", c.baseURL, hotelID))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch hotel by ID: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result struct {
		Hotel models.Hotel `json:"hotel"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &result.Hotel, nil
}

// CreateRoom отправляет запрос на создание новой комнаты.
func (c *Client) CreateRoom(hotelID, name, description string, price uint64) (*models.Room, error) {
	url := fmt.Sprintf("%s/rooms", c.baseURL)

	requestBody := map[string]interface{}{
		"hotelId":     hotelID,
		"name":        name,
		"description": description,
		"price":       price,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Room *models.Room `json:"room"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Room, nil
}

// GetRoomsByHotelId возвращает список всех комнат в указанном отеле.
func (c *Client) GetRoomsByHotelId(hotelID string) ([]models.Room, error) {
	url := fmt.Sprintf("%s/hotels/%s/rooms", c.baseURL, hotelID)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Rooms []models.Room `json:"rooms"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Rooms, nil
}

// DeleteRoom удаляет существующую комнату по её ID.
func (c *Client) DeleteRoom(roomID string) (bool, error) {
	url := fmt.Sprintf("%s/rooms", c.baseURL)

	req, err := http.NewRequest(http.MethodDelete, url+"?roomId="+roomID, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return false, fmt.Errorf("failed to decode response: %w", err)
	}

	return response.Success, nil
}
