import requests
import json

# Base URL of the API
BASE_URL = "http://localhost:8080/v1"

# User credentials
email = "test@example.com"
password = "your_password"
name = "Test User"

# Headers
headers = {
    "Content-Type": "application/json"
}

print("=" * 100)

# 1. User Registration
def test_signup():
    print("Testing User Registration...")
    
    payload = {
        "name": name,
        "email": email,
        "password1": password,
        "password2": password
    }
    
    response = requests.post(f"{BASE_URL}/signup", headers=headers, json=payload)
    
    if response.status_code == 201:
        print("Signup successful:")
        print(json.dumps(response.json(), indent=4))
    else:
        print(f"Signup failed: {response.status_code}")
        print(response.text)

    print("\n")
    return response

# 2. User Login
def test_login():
    print("Testing User Login...")
    
    payload = {
        "email": email,
        "password": password
    }
    
    response = requests.post(f"{BASE_URL}/login", headers=headers, json=payload)
    
    if response.status_code == 200:
        print("Login successful:")
        print(json.dumps(response.json(), indent=4))
        return response.json().get("data", {}).get("idToken")
    else:
        print(f"Login failed: {response.status_code}")
        print(response.text)

    print("\n")
    return None

# 3. Get User Info
def test_user_info(token):
    if token:
        print("Testing User Info...")
        
        headers_with_token = {
            "Content-Type": "application/json",
            "Authorization": f"Bearer {token}"
        }
        
        response = requests.get(f"{BASE_URL}/users/me", headers=headers_with_token)
        
        if response.status_code == 200:
            print("User info retrieved:")
            print(json.dumps(response.json(), indent=4))
        else:
            print(f"Failed to retrieve user info: {response.status_code}")
            print(response.text)

        print("\n")
    else:
        print("No token provided. Skipping user info test.")


if __name__ == "__main__":
    test_signup()

    print("=" * 100)

    token = test_login()

    print("=" * 100)

    test_user_info(token)
