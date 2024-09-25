## Installation
1. **Clone the repository:**

   ```bash
   git clone https://github.com/mohidex/voice-line.git
   cd voice-line
   ```

2. **Set up your environment variables:**

   Change and set your environment variable to your `compose.yml` file:

   ```env
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=vioceline
   FIREBASE_API_KEY=your_firebase_api_key
   FIREBASE_BASE_URL=your_firebase_base_url
   ```

3. **Build and Run the Application:**

    Using Docker Compose, you can build and run the application in one step.

   ```bash
    docker compose up --build
   ```

4. **Run the application in watch mode:**

   Start the application in watch mode:

   ```bash
   docker compose watch
   ```

## Usage

After starting the server, you can interact with the API endpoints outlined below.

## API Endpoints

### User Registration

- **Endpoint:** `POST /signup`
- **Request Body:**
    ```json
    {
        "name": "Test User",
        "email": "test@example.com",
        "password1": "your_password",
        "password2": "your_password"
    }
    ```

- **Response:**
    - Status: `201 Created`
    - Body:
    ```json
    {
        "user": {
            "id": "123",
            "name": "Test User",
            "email": "test@example.com",
            "is_active": true,
            "is_admin": false
        }
    }
    ```

### User Login

- **Endpoint:** `POST /login`
- **Request Body:**
    ```json
    {
        "email": "test@example.com",
        "password": "your_password"
    }
    ```

- **Response:**
    - Status: `200 OK`
    - Body:
    ```json
    {
        "status": "success",
        "data": {
            "idToken": "your_id_token",
            "refreshToken": "your_refresh_token",
            "expiresIn": "3600"
        }
    }
    ```

### User Info

- **Endpoint:** `GET /users/me`
- **Headers:**
    - `Authorization: Bearer <token>`

- **Response:**
    - Status: `200 OK`
    - Body:
    ```json
    {
        "user": {
            "id": "123",
            "name": "Test User",
            "email": "test@example.com",
            "is_active": true,
            "is_admin": false
        }
    }
    ```


## Test using a script


1. **Go to `client-script` directory:**

   ```bash
   cd client-script
   ```

2. **Install required packages::**
    Make sure you have the requests library installed. You can install it via pip if it's not already installed
   ```bash
    pip install requests
   ```

3. **Run the script:**
    ```bash
    python test_api.py
    ```
