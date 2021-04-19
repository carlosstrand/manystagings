package seeds

import "github.com/carlosstrand/manystagings/models"

func createHelloWorldApp() models.Application {
	return models.Application{
		Name:            "node-api",
		DockerImageName: "dockercloud/hello-world",
		DockerImageTag:  "latest",
		Port:            80,
		ContainerPort:   80,
		ApplicationEnvVars: []models.ApplicationEnvVar{
			{
				Key:   "NODE_ENV",
				Value: "development",
			},
			{
				Key:   "SITE_URL",
				Value: "http://staging.mysite.com",
			},
		},
	}
}

var DEFAULT = SeedsData{
	Users: []models.User{
		{
			FirstName:    "Clark",
			LastName:     "Kent",
			Username:     "clark.kent",
			Email:        "clark.kent@test.com",
			PasswordHash: CreatePasswordHash("clark123"),
		},
		{
			FirstName:    "Bruce",
			LastName:     "Wayne",
			Username:     "bruce.wayne",
			Email:        "bruce.wayne@test.com",
			PasswordHash: CreatePasswordHash("bruce123"),
		},
	},
	Environment: []models.Environment{
		{
			Name:        "QA",
			Namespace:   "qa",
			Description: "Environment used for for QA",
			Applications: []models.Application{
				createHelloWorldApp(),
				{
					Name:            "db",
					DockerImageName: "postgres",
					DockerImageTag:  "13.2-alpine",
					Port:            5432,
					ContainerPort:   5432,
					ApplicationEnvVars: []models.ApplicationEnvVar{
						{
							Key:   "POSTGRES_USER",
							Value: "staging",
						},
						{
							Key:   "POSTGRES_PASSWORD",
							Value: "staging",
						},
					},
				},
			},
		},
		{
			Name:        "Clark's Environment",
			Namespace:   "clark",
			Description: "Environment used by Clark Kent",
			Applications: []models.Application{
				createHelloWorldApp(),
			},
		},
		{
			Name:        "Bruce's Environment",
			Namespace:   "bruce",
			Description: "Environment used by Bruce Wayne",
			Applications: []models.Application{
				createHelloWorldApp(),
			},
		},
	},
}
