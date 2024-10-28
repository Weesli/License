****Hi everyone,****

This is a simple license checker for product with licenses.
Everyone will use for free but you don't get a support!


# API Endpoints

## Public Endpoints

### Check License
**Endpoint:** `GET /api/v1/public/checkLicense`  
**Query Parameters:**
- `key=<license-key>` - The license key to check for validity.

---

## Private Endpoints

> **Note:** All private endpoints require an authorization header with `Auth: <admin-secret>`.

### Create License
**Endpoint:** `POST /api/v1/private/createLicense`  
**Request Body:**
- `Owner` - The owner of the license.
- `Name` - The name associated with the license.
- `Key` - The license key to create.
- `Status` - The status of the license.

### Remove License
**Endpoint:** `POST /api/v1/private/removeLicense`  
**Request Body:**
- `Key` - The license key to remove.

### Get All Licenses
**Endpoint:** `POST /api/v1/private/getAll`  
**Request Body:** None

# Usage

>Please note that you cannot request any support on this project!
You are fully responsible for your use and any errors!

# Contact
You want to contact me?

[Discord](https://discord.com/users/509803473106239528)
