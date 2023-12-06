# ENIGMA CAMP 2.0 [BOOKING ROOM]

### Deskripsi

Projek ini adalah Final Project Golang Batch 12, di projek ini kita membuat aplikasi Booking Room untuk menggantikan sistem pemesanan ruangan yang saat ini masih dilakukan secara lisan dengan aplikasi berbasis digital.

---

## API Spec

### Employe

#### Create Employe

Request :

- Method : `POST`
- Endpoint : `/employees`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "name": "string",
  "divisi": "string",
  "jabatan": "string",
  "email": "string",
  "password": "string",
  "role": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### Get Employe By Id

Request :

- Method : `GET`
- Endpoint : `/employees/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### Update Employee

Request :

- Method : `PUT`
- Endpoint : `/employees/:id`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "name": "string",
  "divisi": "string",
  "jabatan": "string",
  "email": "string",
  "password": "string",
  "role": "string"
}
```

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

#### Delete Employee

Request :

- Method : `DELETE`
- Endpoint : `/employees/:id`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": "OK"
}
```

#### Get List Employee

Request :

- Method : `GET`
- Endpoint : `/employees`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
{
  "message": "string",
  "data": {
    "id": "string",
    "name": "string",
    "divisi": "string",
    "jabatan": "string",
    "email": "string",
    "role": "string"
  }
}
```

### Booking

#### Create Booking

Request :

- Method : `POST`
- Endpoint : `/booking`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "bookingDetails": [
    {
      "description": "ini desciption",
      "rooms": {
        "id": "92681558-650e-44fa-95ce-7ee92e7220b8"
      }
    }
  ]
}
```

Response :

- Status : 201 Created
- Body :

```json
{
  "status": {
    "code": 201,
    "description": "Ok"
  },
  "data": {
    "bookingId": "bdcaecf0-6ff5-4f9f-9225-49de1d584212",
    "employe": {
      "id": "8f568c4a-95b4-46ec-ba29-de66b2e84f86",
      "name": "ricky",
      "divisi": "enigma",
      "jabatan": "karyawan",
      "email": "ricky@mail.com",
      "role": "GA",
      "createdAt": "2023-11-15T11:32:48.532463Z",
      "updatedAt": "2023-11-15T11:32:48.466881Z"
    },
    "bookingDetails": [
      {
        "id": "0dd409b8-f8c2-402a-9133-7dcae0c7cada",
        "bookingId": "bdcaecf0-6ff5-4f9f-9225-49de1d584212",
        "rooms": {
          "id": "92681558-650e-44fa-95ce-7ee92e7220b8",
          "roomType": "kolam_renang",
          "maxcapacity": 12,
          "facility": {
            "id": "83cca32f-45da-41ac-a35d-31e53c6952bd",
            "description": "anjay",
            "wifi": "ada tapi indihumu",
            "soundSystem": "salon adanya",
            "projector": "ada",
            "screenProjector": "green screen",
            "chairs": "ada",
            "tables": "ada",
            "soundProof": "ada",
            "smokingArea": "ada",
            "television": "ada",
            "ac": "ac alam",
            "bathroom": "JEDING KEBON",
            "coffeMaker": "ada",
            "createdAt": "2023-11-15T11:50:15.314491Z",
            "updatedAt": "2023-11-15T11:50:15.156612Z"
          },
          "status": "available",
          "createdAt": "2023-11-15T11:50:15.326955Z",
          "updatedAt": "2023-11-15T11:50:15.326477Z"
        },
        "description": "ini desciption",
        "status": "pending",
        "bookingDate": "2023-11-21T08:45:00.954394Z",
        "bookingDateEnd": "2023-11-21T11:45:00.954394Z",
        "createdAt": "2023-11-21T08:45:00.928153Z",
        "updatedAt": "2023-11-21T08:45:00.954394Z"
      }
    ],
    "createdAt": "2023-11-21T08:45:00.928153Z",
    "updatedAt": "2023-11-21T08:45:00.929702Z"
  }
}
```

#### Get Booking By Id

Request :

- Method : `GET`
- Endpoint : `/booking/:id`
- Header :
  - Accept : application/json

Response :

- Status : 200 OK
- Body :

```json
{
  "status": {
    "code": 200,
    "description": "Ok"
  },
  "data": {
    "bookingId": "bdcaecf0-6ff5-4f9f-9225-49de1d584212",
    "employe": {
      "id": "8f568c4a-95b4-46ec-ba29-de66b2e84f86",
      "name": "ricky",
      "divisi": "enigma",
      "jabatan": "karyawan",
      "email": "ricky@mail.com",
      "role": "GA",
      "createdAt": "2023-11-15T11:32:48.532463Z",
      "updatedAt": "2023-11-15T11:32:48.466881Z"
    },
    "bookingDetails": [
      {
        "id": "0dd409b8-f8c2-402a-9133-7dcae0c7cada",
        "bookingId": "",
        "rooms": {
          "id": "92681558-650e-44fa-95ce-7ee92e7220b8",
          "roomType": "kolam_renang",
          "maxcapacity": 12,
          "facility": {
            "id": "83cca32f-45da-41ac-a35d-31e53c6952bd",
            "description": "anjay",
            "wifi": "ada tapi indihumu",
            "soundSystem": "salon adanya",
            "projector": "ada",
            "screenProjector": "",
            "chairs": "ada",
            "tables": "ada",
            "soundProof": "ada",
            "smokingArea": "ada",
            "television": "ada",
            "ac": "ac alam",
            "bathroom": "JEDING KEBON",
            "coffeMaker": "ada",
            "createdAt": "2023-11-15T11:50:15.156612Z",
            "updatedAt": "2023-11-15T11:50:15.314491Z"
          },
          "status": "available",
          "createdAt": "2023-11-15T11:50:15.326955Z",
          "updatedAt": "2023-11-15T11:50:15.326477Z"
        },
        "description": "ini desciption",
        "status": "pending",
        "bookingDate": "2023-11-21T08:45:00.954394Z",
        "bookingDateEnd": "2023-11-21T11:45:00.954394Z",
        "createdAt": "2023-11-21T08:45:00.928153Z",
        "updatedAt": "2023-11-21T08:45:00.954394Z"
      }
    ],
    "createdAt": "2023-11-21T08:45:00.928153Z",
    "updatedAt": "2023-11-21T08:45:00.929702Z"
  }
}
```

#### Get All Booking

Request :

- Method : `GET`
- Endpoint : `/booking`
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
 "status": {
    "code": 200,
    "description": "Ok"
  },
  "data": []booking

```

#### Get Booking By Status

Request :

- Method : `GET`
- Endpoint : `/booking/status/:status` ("pending", "accept", "decline")
- Header :
  - Accept : application/json
- Body :

Response :

- Status : 200 OK
- Body :

```json
 "status": {
    "code": 200,
    "description": "Ok"
  },
  "data": []booking

```

#### Change Status Booking By Booking Details ID

Request :

- Method : `GET`
- Endpoint : `/booking/approval`
- Header :
  - Accept : application/json
- Body :

```json
{
  "approval": "accept",  ("accept", "decline")
  "bookingDetailId": "a543db66-6df7-4cde-a14d-8bf2be6edd2f"
}
```

Response :

- Status : 201
- Body :

```json
{
  "status": {
    "code": 200,
    "description": "Ok"
  },
  "data": {
    "bookingId": "bdcaecf0-6ff5-4f9f-9225-49de1d584212",
    "employe": {
      "id": "8f568c4a-95b4-46ec-ba29-de66b2e84f86",
      "name": "ricky",
      "divisi": "enigma",
      "jabatan": "karyawan",
      "email": "ricky@mail.com",
      "role": "GA",
      "createdAt": "2023-11-15T11:32:48.532463Z",
      "updatedAt": "2023-11-15T11:32:48.466881Z"
    },
    "bookingDetails": [
      {
        "id": "0dd409b8-f8c2-402a-9133-7dcae0c7cada",
        "bookingId": "",
        "rooms": {
          "id": "92681558-650e-44fa-95ce-7ee92e7220b8",
          "roomType": "kolam_renang",
          "maxcapacity": 12,
          "facility": {
            "id": "83cca32f-45da-41ac-a35d-31e53c6952bd",
            "description": "anjay",
            "wifi": "ada tapi indihumu",
            "soundSystem": "salon adanya",
            "projector": "ada",
            "screenProjector": "",
            "chairs": "ada",
            "tables": "ada",
            "soundProof": "ada",
            "smokingArea": "ada",
            "television": "ada",
            "ac": "ac alam",
            "bathroom": "JEDING KEBON",
            "coffeMaker": "ada",
            "createdAt": "2023-11-15T11:50:15.156612Z",
            "updatedAt": "2023-11-15T11:50:15.314491Z"
          },
          "status": "available",
          "createdAt": "2023-11-15T11:50:15.326955Z",
          "updatedAt": "2023-11-15T11:50:15.326477Z"
        },
        "description": "ini desciption",
        "status": "decline",
        "bookingDate": "2023-11-21T08:45:00.954394Z",
        "bookingDateEnd": "2023-11-21T11:45:00.954394Z",
        "createdAt": "2023-11-21T08:45:00.928153Z",
        "updatedAt": "2023-11-21T08:45:00.954394Z"
      }
    ],
    "createdAt": "2023-11-21T08:45:00.928153Z",
    "updatedAt": "2023-11-21T08:45:00.929702Z"
  }
}
```
