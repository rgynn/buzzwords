package buzzwords

import (
	"fmt"
	"math/rand"
)

var verbs = []string{
	"restarting",
	"lazy loading",
	"rerouting",
	"deploying",
	"testing",
	"reloading",
	"initiating",
	"drawing",
	"tailing",
	"securing",
	"printing",
	"downloading",
	"hacking",
	"transcoding",
	"translating",
	"encoding",
	"decoding",
	"connecting",
	"disconnecting",
	"converting",
	"diverting",
	"stopping",
	"populating",
	"scaling",
	"automating",
	"reticulating",
	"injesting",
	"spamming",
	"truncating",
	"reverse engineering",
	"calling",
	"benchmarking",
	"federating",
	"calculating",
}

var adjectives = []string{
	"agile",
	"all flash",
	"auto-scaling",
	"automation centric",
	"apache licensed",
	"application aware",
	"artificially intelligent",
	"backwards compatible",
	"blockchain",
	"cloud agnostic",
	"cloud native",
	"concurrent",
	"customer centric",
	"crowdfunded",
	"data aware",
	"deep learning",
	"developer focused",
	"disruptive",
	"distributed",
	"eagerly loaded",
	"elastic",
	"event driven",
	"everything on-demand",
	"gesture based",
	"globally accessible",
	"greenwashing",
	"haptic",
	"hassle free",
	"holographic",
	"hybrid cloud",
	"hyperconverged",
	"immersive",
	"lazy loaded",
	"location based",
	"managed",
	"mobile first",
	"native",
	"on-demand",
	"open source",
	"highly resilient",
	"high performant",
	"petabyte",
	"programmatic",
	"quantum",
	"reactive",
	"real-time",
	"responsive",
	"robust and scalable",
	"seamless",
	"self-healing",
	"self-service",
	"serverless",
	"service-oriented",
	"software defined",
	"social media driven",
	"stateless",
	"viral",
	"wearable",
}

var nouns = []string{
	"3d printing",
	"a/b testing",
	"artifical intelligence",
	"augmented reality",
	"big data",
	"wetware",
	"biowearables",
	"blockchains",
	"business intelligence",
	"chaosmonkey",
	"chatbots",
	"chatops",
	"computing",
	"configuration management",
	"containerization",
	"containers",
	"continous deployment",
	"continous integration",
	"content delivery network",
	"crypto currency",
	"cyberattack",
	"database clusters",
	"datacenter",
	"data visualization",
	"datamining",
	"data pipelines",
	"data transactions",
	"datawarehousing",
	"deep learning",
	"deep linking",
	"deep web",
	"device mesh",
	"devops",
	"devsec",
	"job dispatchers",
	"embedded devices",
	"faas",
	"filesystem",
	"gamification",
	"growth hacking",
	"home automation",
	"iaas",
	"information",
	"infrastructure appliance",
	"infrastructure",
	"iot",
	"l3 cache",
	"machine learning",
	"microservices",
	"monitoring",
	"net neutrality",
	"npm packages",
	"networking",
	"neural networks",
	"nosql database",
	"infrastructure orchestration",
	"paas",
	"real-time data",
	"saas",
	"container scheduling",
	"server migration",
	"service discovery",
	"service integration",
	"site reliability management",
	"slack integration",
	"slack bots",
	"software development",
	"surveillance",
	"thought leaders",
	"timeseries database",
	"transaction brokers",
	"ux design",
	"video transcoding",
	"virtual assistants",
	"virtual reality",
}

var suffixes = []string{
	"as a service",
	"as code",
	"of everything",
	"of things",
}

// BuzzWords returns a sentence with adjective and noun
func BuzzWords() string {
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("%v %v", adjective, noun)
}

// WithSuffix returns a sentence with adjective, noun and suffix
func WithSuffix() string {
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	return fmt.Sprintf("%v %v %v", adjective, noun, suffix)
}

// WithVerb returns a sentence with a verb, adjective and noun
func WithVerb() string {
	verb := verbs[rand.Intn(len(verbs))]
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("%v %v %v", verb, adjective, noun)
}

// WithVerbAndSuffix returns a sentence with a verb, adjective, noun and suffix
func WithVerbAndSuffix() string {
	verb := verbs[rand.Intn(len(verbs))]
	adjective := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	suffix := suffixes[rand.Intn(len(suffixes))]
	return fmt.Sprintf("%v %v %v %v", verb, adjective, noun, suffix)
}
