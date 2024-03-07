# Airline Booking System

## Overview

The Airline Booking System is a robust and user-friendly console application designed to facilitate the booking of airline seats. The application is built with Go, leveraging its concurrent features to handle multiple bookings simultaneously and ensure a smooth user experience. It is designed to provide a quick and efficient means of reserving seats, managing passenger details, and issuing boarding passes.

## Architecture Overview

The system is composed of the following components:

- **User Input Collection**: A simple, interactive console interface that collects user input, including personal details and seat requirements.
- **Input Validation**: Utilizes a helper package to ensure that all user input is valid, including name checks, email format verification, and passport number validity.
- **Seat Booking and Management**: Core functionality that manages seat availability and books seats according to user requests.
- **Boarding Pass Generation**: A concurrent process that simulates the generation and sending of boarding passes to the passenger's email.
- **Data Persistence**: Currently, the application stores booking information in memory. Persistence to a database can be implemented for long-term storage.



