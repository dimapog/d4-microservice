package ai

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

type Service interface {
	CalculatePersonalStatistics(ctx context.Context, req *PersonalCalculationRequest) (interface{}, error)
}

type service struct {
	apiKey string
}

func NewService() Service {
	return &service{
		apiKey: os.Getenv("OPENAI_API_KEY"),
	}
}

func (s *service) CalculatePersonalStatistics(ctx context.Context, req *PersonalCalculationRequest) (interface{}, error) {
	if s.apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY is not configured")
	}

	client := openai.NewClient(
		option.WithAPIKey(s.apiKey),
	)

	prompt := s.buildPrompt(req)

	resp, err := client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{OfString: openai.String(prompt)},
		Model: openai.ChatModelGPT5_4Nano,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to call OpenAI API: %w", err)
	}

	return resp, nil
}

func (s *service) buildPrompt(req *PersonalCalculationRequest) string {
	return fmt.Sprintf(`You are a professional fitness and health assistant.

	Your task is to calculate BMI and heart rate training zones based on the provided user data.

	USER DATA:
	- Age: %d
	- Gender: %s
	- Weight: %.1f kg
	- Height: %.1f cm
	- Resting Heart Rate: %d bpm
	- Units: %s

	INSTRUCTIONS:
	1. Validate the input data (ensure all values are reasonable).
	2. Calculate BMI using:
	BMI = weight (kg) / (height (m))^2
	3. Round BMI to 1 decimal place.
	4. Determine BMI category:
	- Underweight (<18.5)
	- Normal (18.5–24.9)
	- Overweight (25–29.9)
	- Obese (30+)

	5. Calculate Max Heart Rate:
	maxHR = 220 - age

	6. Calculate heart rate zones using Karvonen formula:
	Target HR = ((maxHR - restingHR) * intensity) + restingHR

	7. Use these intensity ranges:
	- Zone 1: 50–60%%
	- Zone 2: 60–70%%
	- Zone 3: 70–80%%
	- Zone 4: 80–90%%
	- Zone 5: 90–100%%

	8. Return heart rate zones as BPM ranges (min and max).
	9. Round all heart rate values to whole numbers.

	OUTPUT REQUIREMENTS:
	- Respond ONLY in valid JSON
	- No explanations, no text outside JSON

	OUTPUT FORMAT:
	{
	"bmi": {
		"value": number,
		"category": string
	},
	"heart_rate": {
		"max_hr": number,
		"zones": {
		"zone_1": {"min": number, "max": number},
		"zone_2": {"min": number, "max": number},
		"zone_3": {"min": number, "max": number},
		"zone_4": {"min": number, "max": number},
		"zone_5": {"min": number, "max": number}
		}
	},
	"meta": {
		"input_valid": boolean,
		"notes": string
	}
	}

	If the output is not valid JSON, fix it before responding.`,
		req.Age,
		req.Gender,
		req.Weight,
		req.Height,
		req.RestingHeartRate,
		req.Units,
	)
}
