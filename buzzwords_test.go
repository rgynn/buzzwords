package buzzwords

import (
	"strings"
	"testing"
)

func TestRandomAdjective(t *testing.T) {
	result := RandomAdjective()
	foundMatch := false
	for _, adjective := range adjectives {
		if result == adjective {
			foundMatch = true
		}
	}
	if foundMatch == false {
		t.Fatalf("Expected an adjective, got %s", result)
	}
}

func TestRandomNoun(t *testing.T) {
	result := RandomNoun()
	foundMatch := false
	for _, noun := range nouns {
		if result == noun {
			foundMatch = true
		}
	}
	if foundMatch == false {
		t.Fatalf("Expected a noun, got %s", result)
	}
}

func TestRandomVerb(t *testing.T) {
	result := RandomVerb()
	foundMatch := false
	for _, verb := range verbs {
		if result == verb {
			foundMatch = true
		}
	}
	if foundMatch == false {
		t.Fatalf("Expected a verb, got %s", result)
	}
}

func TestRandomSuffix(t *testing.T) {
	result := RandomSuffix()
	foundMatch := false
	for _, suffix := range suffixes {
		if result == suffix {
			foundMatch = true
		}
	}
	if foundMatch == false {
		t.Fatalf("Expected a suffix, got %s", result)
	}
}

func TestBuzzWords(t *testing.T) {
	result := BuzzWords()

	foundAdjective := false
	for _, adjective := range adjectives {
		if strings.Contains(result, adjective) {
			foundAdjective = true
		}
	}
	if foundAdjective == false {
		t.Fatalf("Expected an adjective in sentence, got %s", result)
	}

	foundNoun := false
	for _, noun := range nouns {
		if strings.Contains(result, noun) {
			foundNoun = true
		}
	}
	if foundNoun == false {
		t.Fatalf("Expected an noun in sentence, got %s", result)
	}
}

func TestWithSuffix(t *testing.T) {
	result := WithSuffix()

	foundAdjective := false
	for _, adjective := range adjectives {
		if strings.Contains(result, adjective) {
			foundAdjective = true
		}
	}
	if foundAdjective == false {
		t.Fatalf("Expected an adjective in sentence, got %s", result)
	}

	foundNoun := false
	for _, noun := range nouns {
		if strings.Contains(result, noun) {
			foundNoun = true
		}
	}
	if foundNoun == false {
		t.Fatalf("Expected an noun in sentence, got %s", result)
	}

	foundSuffix := false
	for _, suffix := range suffixes {
		if strings.Contains(result, suffix) {
			foundSuffix = true
		}
	}
	if foundSuffix == false {
		t.Fatalf("Expected an suffix in sentence, got %s", result)
	}
}

func TestWithVerb(t *testing.T) {
	result := WithVerb()

	foundVerb := false
	for _, verb := range verbs {
		if strings.Contains(result, verb) {
			foundVerb = true
		}
	}
	if foundVerb == false {
		t.Fatalf("Expected an verb in sentence, got %s", result)
	}

	foundAdjective := false
	for _, adjective := range adjectives {
		if strings.Contains(result, adjective) {
			foundAdjective = true
		}
	}
	if foundAdjective == false {
		t.Fatalf("Expected an adjective in sentence, got %s", result)
	}

	foundNoun := false
	for _, noun := range nouns {
		if strings.Contains(result, noun) {
			foundNoun = true
		}
	}
	if foundNoun == false {
		t.Fatalf("Expected an noun in sentence, got %s", result)
	}
}

func TestWithVerbAndSuffix(t *testing.T) {
	result := WithVerbAndSuffix()

	foundVerb := false
	for _, verb := range verbs {
		if strings.Contains(result, verb) {
			foundVerb = true
		}
	}
	if foundVerb == false {
		t.Fatalf("Expected an verb in sentence, got %s", result)
	}

	foundAdjective := false
	for _, adjective := range adjectives {
		if strings.Contains(result, adjective) {
			foundAdjective = true
		}
	}
	if foundAdjective == false {
		t.Fatalf("Expected an adjective in sentence, got %s", result)
	}

	foundNoun := false
	for _, noun := range nouns {
		if strings.Contains(result, noun) {
			foundNoun = true
		}
	}
	if foundNoun == false {
		t.Fatalf("Expected an noun in sentence, got %s", result)
	}

	foundSuffix := false
	for _, suffix := range suffixes {
		if strings.Contains(result, suffix) {
			foundSuffix = true
		}
	}
	if foundSuffix == false {
		t.Fatalf("Expected an suffix in sentence, got %s", result)
	}
}
