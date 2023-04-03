# Go Web Application with Sentry

This is a simple web application written in Go that demonstrates how to integrate Sentry for performance monitoring and error tracking for a web application. It allows users to enter their name and age and displays a personalized greeting based on their age.

## Requirements

- Go 1.16 or higher
- A Sentry account and project with an API key

## Installation

1. Clone this repository.
2. Set your Sentry DSN by replacing `YOUR_KEY` in the `main()` function with your own DSN.
3. Run the main function
4. Open your web browser and go to `http://localhost:8085` to access the web application.

## Usage

1. Enter your name and age in the web form.
2. Click the "Submit" button.
3. The application will display a personalized greeting based on your age.

## Sentry Integration

This application includes Sentry error tracking to help you monitor and debug errors that occur in production. To view error events in your Sentry dashboard, follow these steps:
1. Perform an action that triggers an error, such as entering an invalid age in the web form.
2. Go to your Sentry dashboard and select your project.
3. Click on "Issues" in the sidebar.
4. You should see a list of error events that have occurred in your application. Click on an event to view more details, such as the error message, stack trace, and user information.
