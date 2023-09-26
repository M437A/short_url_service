# My Personal URL Shortener

Welcome to my personal URL shortener project! This small yet impactful project showcases my knowledge and skills in web development. It's a simple tool that takes long, unwieldy URLs and turns them into concise and easy-to-share links.

## Project Overview
![image](https://github.com/M437A/short_url_service/assets/105558638/df1f72d4-4604-4428-9471-b85d9bbc06f1)

### URL Shortening Controller:

1. Receive user requests to shorten URLs.
2. Validate incoming data to ensure accuracy.
3. If data is valid, retrieve pre-generated short URLs and provide them to users. Simultaneously, asynchronously store the data in the database and cache it for 1 day.

### Short Link Handling Controller:

1. Receive short link requests from users.
2. Check the cache for corresponding long URLs. If a long URL is not found in the cache, asynchronously record it, along with the count of monthly calls. If the call count is high, cache the data as well.
3. Query the database to fetch the corresponding long URL. Concurrently, send data to Amplitude analytics through a message broker (e.g., Kafka).

This dual-controller approach ensures efficient management of shortened URL addresses, enabling rapid resolution for users and data collection and analysis to enhance the project.

Implementing URL shortening mechanisms with comprehensive analytics and caching can be highly beneficial for tracking popular links and improving performance in handling user requests.

## Technology Stack

- **Go**: The backend of this project is powered by Go, a programming language known for its simplicity and efficiency.
- **PostgreSQL**: To store and manage link data, I use PostgreSQL, a robust open-source relational database.
- **Kafka**: Kafka helps me handle real-time data streams efficiently, providing seamless communication.
- **Redis**: Redis serves as the high-speed, in-memory cache for quick data retrieval.
- **Chi**: Chi is the lightning-fast HTTP router for Go, optimizing routing and enhancing user experience.

## Getting Started

To see my personal URL shortener in action, follow these simple steps:

1. Run the migrations script to create the necessary database tables by executing `./up.sh`.
2. Start the project by running `./run.sh`.

## Analytics Integration

For added functionality, I've integrated the service with Amplitude, an analytics platform. Feel free to create and configure your Amplitude account to gain insights into user behavior.

## Contributions

While this project is primarily a personal endeavor, I'm open to feedback and suggestions. If you have any ideas or spot improvements, please don't hesitate to reach out or create issues.

Thank you for visiting my personal URL shortener project!
