---
title: Dealls Case Study
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.30"

---

# Dealls Case Study

Base URLs:

* <a href="http://localhost:8079">Develop Env: http://localhost:8079</a>

# Authentication

- HTTP Authentication, scheme: bearer

# authentication-service

## POST Login

POST /api/auth/api/login

> Body Parameters

```json
{
  "username": "3030302",
  "password": "Dealls10!"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Successfully logged in",
  "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDk2MTgxNDgsInJvbGVzIjpbIkFETUlOX0hSX1BBWVJPTEwiLCJFTVBMT1lFRSJdLCJ1c2VyIjoiMzAzMDMwMCJ9.KNpF5Lr3qgbQBJYVrnNjW9PzHWT6hUn-PSojaqgWWD4"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# employee-service/microservice

## POST Get Employee By Nip

POST /services/employees/services/employee/get-by-nip

> Body Parameters

```json
{
  "nip": "3030302"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Employee retrieved successfully",
  "data": {
    "ID": 1016,
    "CreatedAt": "2025-06-08T07:08:28.708083+07:00",
    "UpdatedAt": "2025-06-08T07:08:28.708083+07:00",
    "DeletedAt": null,
    "CreatedBy": "SYSTEM",
    "UpdatedBy": "SYSTEM",
    "DeletedBy": "",
    "name": "Lady Shakira Hackett",
    "nip": "3030302",
    "salary": 7985870
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## GET Get All Employees

GET /services/employees/services/employee/get-all-employee

> Body Parameters

```json
{}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Employee list retrieved successfully",
  "data": [
    {
      "nip": "3030300",
      "name": "Miss Sandra Streich",
      "salary": 5901288
    },
    {
      "nip": "3030302",
      "name": "Lady Shakira Hackett",
      "salary": 7985870
    },
    {
      "nip": "3030303",
      "name": "Mrs. Elena Beer",
      "salary": 7285724
    },
    {
      "nip": "3030304",
      "name": "Lady Cecilia Cormier",
      "salary": 6248426
    },
    {
      "nip": "344825934584232",
      "name": "Mrs. Marcia Kerluke",
      "salary": 7262942
    },
    {
      "nip": "372158902901335",
      "name": "Queen Ashley Okuneva",
      "salary": 4920956
    },
    {
      "nip": "378015674983431",
      "name": "Mrs. Avis Kovacek",
      "salary": 4563419
    },
    {
      "nip": "340374359941019",
      "name": "Mrs. Millie Miller",
      "salary": 7496093
    },
    {
      "nip": "375874233313400",
      "name": "Mrs. Oma Schulist",
      "salary": 7445181
    },
    {
      "nip": "349516509566153",
      "name": "Dr. Alexandrine Parisian",
      "salary": 4738746
    },
    {
      "nip": "347779246252952",
      "name": "Lady Yoshiko Schmeler",
      "salary": 4816440
    },
    {
      "nip": "371381224529469",
      "name": "Miss Wilhelmine Leannon",
      "salary": 5130829
    },
    {
      "nip": "346661836095614",
      "name": "Miss Roxanne Rohan",
      "salary": 7733408
    },
    {
      "nip": "373569409207119",
      "name": "Mrs. Ashlee Harris",
      "salary": 5269033
    },
    {
      "nip": "345572997799347",
      "name": "Ms. Michele Friesen",
      "salary": 4130790
    },
    {
      "nip": "375536874663994",
      "name": "Lady Loren Stamm",
      "salary": 3402660
    },
    {
      "nip": "377891546550930",
      "name": "Queen Stephany Walsh",
      "salary": 3771816
    },
    {
      "nip": "373650588750044",
      "name": "Mrs. Christiana Spinka",
      "salary": 3306389
    },
    {
      "nip": "345868163612952",
      "name": "Miss Maddison Schinner",
      "salary": 7807150
    },
    {
      "nip": "374970793743730",
      "name": "Prof. Mckayla Kuhic",
      "salary": 6155191
    },
    {
      "nip": "340636614246700",
      "name": "Miss Grace Rosenbaum",
      "salary": 3536866
    },
    {
      "nip": "340044046626459",
      "name": "Queen Candace Sanford",
      "salary": 3382353
    },
    {
      "nip": "370135330497530",
      "name": "Dr. Dortha Towne",
      "salary": 4704115
    },
    {
      "nip": "343411567008656",
      "name": "Prof. Linnie Gottlieb",
      "salary": 6498986
    },
    {
      "nip": "343698769049697",
      "name": "Queen Gloria Grady",
      "salary": 4185995
    },
    {
      "nip": "372308970292461",
      "name": "Lady Josiane Schinner",
      "salary": 3516574
    },
    {
      "nip": "379627758583783",
      "name": "Miss Yessenia Russel",
      "salary": 6718737
    },
    {
      "nip": "379745415086086",
      "name": "Princess Nyah Runolfsson",
      "salary": 3452637
    },
    {
      "nip": "375214107640131",
      "name": "Lady Breanne Von",
      "salary": 5123848
    },
    {
      "nip": "379851671132702",
      "name": "Prof. Bettie Murray",
      "salary": 7779702
    },
    {
      "nip": "348137877930355",
      "name": "Princess Angelita Johnson",
      "salary": 7693248
    },
    {
      "nip": "344630635862276",
      "name": "Princess Isabella O\"Conner",
      "salary": 3532236
    },
    {
      "nip": "379686065162906",
      "name": "Dr. Flo Schowalter",
      "salary": 4144573
    },
    {
      "nip": "378529489793705",
      "name": "Miss Maymie Dickinson",
      "salary": 6083492
    },
    {
      "nip": "345541847207492",
      "name": "Prof. Marisa Considine",
      "salary": 4316053
    },
    {
      "nip": "348853986721592",
      "name": "Miss Margarett Wiza",
      "salary": 5556449
    },
    {
      "nip": "342138086993928",
      "name": "Mrs. Laurence Hirthe",
      "salary": 6573446
    },
    {
      "nip": "340011877509857",
      "name": "Queen Aryanna Schmidt",
      "salary": 6376367
    },
    {
      "nip": "377742409868922",
      "name": "Ms. Pink Wiza",
      "salary": 7785716
    },
    {
      "nip": "344613482669814",
      "name": "Queen Allison Rice",
      "salary": 3053574
    },
    {
      "nip": "341334251161893",
      "name": "Queen Piper Schmeler",
      "salary": 5025722
    },
    {
      "nip": "379433572477027",
      "name": "Dr. Bernadette Hauck",
      "salary": 5779763
    },
    {
      "nip": "376134801955841",
      "name": "Mrs. Rosalinda Bruen",
      "salary": 7997070
    },
    {
      "nip": "373807710829489",
      "name": "Miss Mattie Breitenberg",
      "salary": 7585316
    },
    {
      "nip": "347881240530567",
      "name": "Prof. Velda O\"Keefe",
      "salary": 3133898
    },
    {
      "nip": "343687942294558",
      "name": "Dr. Aniyah Moore",
      "salary": 6989768
    },
    {
      "nip": "343559971057013",
      "name": "Mrs. Eloise Hirthe",
      "salary": 6120951
    },
    {
      "nip": "373608411129108",
      "name": "Dr. Providenci Zieme",
      "salary": 7507342
    },
    {
      "nip": "341348510049489",
      "name": "Dr. Princess Schulist",
      "salary": 5101704
    },
    {
      "nip": "345893257733419",
      "name": "Mrs. Annamarie Deckow",
      "salary": 3198606
    },
    {
      "nip": "349992961312830",
      "name": "Ms. Jennie Towne",
      "salary": 7815097
    },
    {
      "nip": "370884835802397",
      "name": "Queen Joelle Considine",
      "salary": 7059488
    },
    {
      "nip": "344532900681657",
      "name": "Ms. Sallie Schimmel",
      "salary": 5196722
    },
    {
      "nip": "342084905883916",
      "name": "Miss Esther Rowe",
      "salary": 3193629
    },
    {
      "nip": "345557021734626",
      "name": "Ms. Adela Nader",
      "salary": 7954464
    },
    {
      "nip": "348718955749905",
      "name": "Princess Otha Dicki",
      "salary": 3902455
    },
    {
      "nip": "344208689092552",
      "name": "Lady Lacey Kuhic",
      "salary": 7659511
    },
    {
      "nip": "371229062241214",
      "name": "Mrs. Jaquelin Carroll",
      "salary": 7754609
    },
    {
      "nip": "348876928595789",
      "name": "Prof. Treva Pfannerstill",
      "salary": 6722194
    },
    {
      "nip": "371104994388771",
      "name": "Miss Billie Schumm",
      "salary": 5171002
    },
    {
      "nip": "374615859750275",
      "name": "Prof. Allison Osinski",
      "salary": 6592625
    },
    {
      "nip": "373264560950005",
      "name": "Miss Esperanza Zulauf",
      "salary": 3707729
    },
    {
      "nip": "371570004858764",
      "name": "Lady Caitlyn Wyman",
      "salary": 5470547
    },
    {
      "nip": "373244653066886",
      "name": "Lady Jackie O\"Keefe",
      "salary": 5518548
    },
    {
      "nip": "342810175844006",
      "name": "Mrs. Marielle Schinner",
      "salary": 7008411
    },
    {
      "nip": "372450880802794",
      "name": "Miss Celestine Marvin",
      "salary": 6057659
    },
    {
      "nip": "344665666360659",
      "name": "Dr. Jaqueline Osinski",
      "salary": 7505004
    },
    {
      "nip": "341530939916602",
      "name": "Prof. Gracie Reichert",
      "salary": 4911782
    },
    {
      "nip": "377209801453160",
      "name": "Mrs. Lilyan Wiza",
      "salary": 3855018
    },
    {
      "nip": "375411087835299",
      "name": "Mrs. Gloria Brekke",
      "salary": 3618192
    },
    {
      "nip": "373911982701903",
      "name": "Princess Aliza Schimmel",
      "salary": 3522039
    },
    {
      "nip": "343165329297962",
      "name": "Princess Martina Mosciski",
      "salary": 3353873
    },
    {
      "nip": "376461154976017",
      "name": "Dr. Virginie Bechtelar",
      "salary": 4559753
    },
    {
      "nip": "371173083176942",
      "name": "Prof. Clarissa Kunze",
      "salary": 4305481
    },
    {
      "nip": "343691435868138",
      "name": "Ms. Sabrina Parisian",
      "salary": 6520572
    },
    {
      "nip": "340897581967457",
      "name": "Ms. Ayana Goodwin",
      "salary": 5070415
    },
    {
      "nip": "346523829767172",
      "name": "Ms. Blanca Hegmann",
      "salary": 6461688
    },
    {
      "nip": "346961704055344",
      "name": "Queen Muriel Dickinson",
      "salary": 6069340
    },
    {
      "nip": "377996815712345",
      "name": "Ms. Barbara Lesch",
      "salary": 5056187
    },
    {
      "nip": "371765844178191",
      "name": "Miss May White",
      "salary": 3071229
    },
    {
      "nip": "345137203379052",
      "name": "Queen Cleora Streich",
      "salary": 5702141
    },
    {
      "nip": "370495688443644",
      "name": "Lady Casandra Dare",
      "salary": 3763474
    },
    {
      "nip": "378603695133128",
      "name": "Prof. Estelle Hegmann",
      "salary": 4257651
    },
    {
      "nip": "373274465418883",
      "name": "Miss Laury Kerluke",
      "salary": 7989283
    },
    {
      "nip": "375821313941531",
      "name": "Lady Kylie Graham",
      "salary": 6518745
    },
    {
      "nip": "376949802508791",
      "name": "Lady Sabrina Roberts",
      "salary": 4171854
    },
    {
      "nip": "371257885196841",
      "name": "Lady Nicolette Sawayn",
      "salary": 6641448
    },
    {
      "nip": "375713391943454",
      "name": "Princess Myriam Witting",
      "salary": 3014368
    },
    {
      "nip": "376199318198964",
      "name": "Ms. Earlene Moore",
      "salary": 7964754
    },
    {
      "nip": "348481991335134",
      "name": "Mrs. Maegan Shields",
      "salary": 5450532
    },
    {
      "nip": "346783279167482",
      "name": "Ms. Verna Morar",
      "salary": 7628443
    },
    {
      "nip": "374178146363472",
      "name": "Princess Vernice Lehner",
      "salary": 6156151
    },
    {
      "nip": "342506313886713",
      "name": "Ms. Peggie Hackett",
      "salary": 3377458
    },
    {
      "nip": "343591743281267",
      "name": "Mrs. Lorine Haley",
      "salary": 5901153
    },
    {
      "nip": "376875185090789",
      "name": "Prof. Janelle Hansen",
      "salary": 4039156
    },
    {
      "nip": "343430361333370",
      "name": "Prof. Bianka Walker",
      "salary": 4940163
    },
    {
      "nip": "375797635043425",
      "name": "Miss Heloise Robel",
      "salary": 4215006
    },
    {
      "nip": "344053371411342",
      "name": "Dr. Erica Casper",
      "salary": 4688925
    },
    {
      "nip": "346590617449413",
      "name": "Princess Hillary Kutch",
      "salary": 7635158
    },
    {
      "nip": "343235748252211",
      "name": "Dr. Helen Osinski",
      "salary": 4880837
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# audit-service

## POST Get All Employees

POST /audit/test/publish

> Body Parameters

```json
{
  "event": "UserLoggedIn",
  "service": "auth-service",
  "request_id": "req-123456",
  "ip": "192.168.1.10",
  "timestamp": "2025-06-08T12:34:56Z"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# attendance-service/admin

## POST Add Attendance Period

POST /api/attendance/admin/attendance-period

> Body Parameters

```json
{
  "startDate": "2025-01-01",
  "endDate": "2025-01-31",
  "payrollPeriodCode": "2025_JUNE"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Attendance period created successfully",
  "data": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# attendance-service/api

## POST Submit Attendance

POST /api/attendance/api/attendance

> Body Parameters

```json
{}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Attendance submitted successfully",
  "data": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# attendance-service/microservice

## POST Get Attendance Period by Payroll

POST /services/attendance/services/get-attendance-period

> Body Parameters

```json
{
  "payrollCode": "2025_JUNE"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Attendance period created successfully",
  "data": {
    "startDate": "2025-06-01T07:00:00+07:00",
    "endDate": "2025-06-30T07:00:00+07:00",
    "payrollPeriodCode": "2025_JUNE"
  }
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## POST Get Attendance Summary

POST /services/attendance/services/get-attendance-summary

> Body Parameters

```json
{
  "periodStart": "2025-05-01T00:00:00Z07:00",
  "periodEnd": "2025-05-30T23:59:59Z07:00"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Attendance summaries retrieved successfully",
  "data": [
    {
      "nip": "3030300",
      "attendanceDetails": [
        {
          "date": "2025-06-09"
        },
        {
          "date": "2025-06-10"
        }
      ]
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# overtime-service/api

## POST Submit Overtime

POST /api/overtime/api/overtime

> Body Parameters

```json
{
  "hours": 2.5
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Overtime submitted successfully",
  "data": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# overtime-service/microservice

## POST Get Overtime Summary

POST /services/overtime/services/get-overtime-summary

> Body Parameters

```json
{
  "periodStart": "2025-06-01T00:00:00Z07:00",
  "periodEnd": "2025-06-30T23:59:59Z07:00"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Attendance summaries retrieved successfully",
  "data": [
    {
      "nip": "3030300",
      "overtimeDetails": [
        {
          "date": "2025-06-10",
          "hours": 2.5
        }
      ]
    },
    {
      "nip": "3030302",
      "overtimeDetails": [
        {
          "date": "2025-06-08",
          "hours": 2.5
        }
      ]
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# reimbursement-service/api

## POST Submit Reimbursement

POST /api/reimbursement/api/reimbursement

> Body Parameters

```json
{
  "amount": 100000,
  "description": "Parkir bulanan"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Reimbursement submitted successfully",
  "data": null
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# reimbursement-service/microservice

## POST Get Reimbursement Summary

POST /services/reimbursement/services/get-reimbursement-summary

> Body Parameters

```json
{
  "periodStart": "2025-06-01T00:00:00Z07:00",
  "periodEnd": "2025-06-30T23:59:59Z07:00"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Attendance summaries retrieved successfully",
  "data": [
    {
      "nip": "3030300",
      "reimbursementDetails": [
        {
          "amount": 100000,
          "description": "Parkir bulanan",
          "date": "2025-06-09"
        },
        {
          "amount": 100000,
          "description": "Parkir bulanan",
          "date": "2025-06-10"
        }
      ]
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# payroll-service/admin

## POST Run Payroll

POST /api/payroll/admin/payroll/run

> Body Parameters

```json
{
  "payrollCode": "2025_JUNE"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{
  "status": "OK",
  "code": 200,
  "message": "Payroll ran successfully",
  "data": [
    {
      "nip": "3030300",
      "name": "Miss Sandra Streich",
      "salary": 5901288,
      "attendanceDetails": [
        {
          "date": "2025-06-09"
        },
        {
          "date": "2025-06-10"
        }
      ],
      "overtimeDetails": [
        {
          "date": "2025-06-10",
          "hours": 2.5
        }
      ],
      "reimbursementDetails": [
        {
          "amount": 100000,
          "description": "Parkir bulanan",
          "date": "2025-06-09"
        },
        {
          "amount": 100000,
          "description": "Parkir bulanan",
          "date": "2025-06-10"
        }
      ]
    },
    {
      "nip": "3030302",
      "name": "Lady Shakira Hackett",
      "salary": 7985870,
      "attendanceDetails": null,
      "overtimeDetails": [
        {
          "date": "2025-06-08",
          "hours": 2.5
        }
      ],
      "reimbursementDetails": null
    },
    {
      "nip": "3030303",
      "name": "Mrs. Elena Beer",
      "salary": 7285724,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "3030304",
      "name": "Lady Cecilia Cormier",
      "salary": 6248426,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344825934584232",
      "name": "Mrs. Marcia Kerluke",
      "salary": 7262942,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "372158902901335",
      "name": "Queen Ashley Okuneva",
      "salary": 4920956,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "378015674983431",
      "name": "Mrs. Avis Kovacek",
      "salary": 4563419,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "340374359941019",
      "name": "Mrs. Millie Miller",
      "salary": 7496093,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375874233313400",
      "name": "Mrs. Oma Schulist",
      "salary": 7445181,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "349516509566153",
      "name": "Dr. Alexandrine Parisian",
      "salary": 4738746,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "347779246252952",
      "name": "Lady Yoshiko Schmeler",
      "salary": 4816440,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371381224529469",
      "name": "Miss Wilhelmine Leannon",
      "salary": 5130829,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "346661836095614",
      "name": "Miss Roxanne Rohan",
      "salary": 7733408,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373569409207119",
      "name": "Mrs. Ashlee Harris",
      "salary": 5269033,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "345572997799347",
      "name": "Ms. Michele Friesen",
      "salary": 4130790,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375536874663994",
      "name": "Lady Loren Stamm",
      "salary": 3402660,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "377891546550930",
      "name": "Queen Stephany Walsh",
      "salary": 3771816,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373650588750044",
      "name": "Mrs. Christiana Spinka",
      "salary": 3306389,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "345868163612952",
      "name": "Miss Maddison Schinner",
      "salary": 7807150,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "374970793743730",
      "name": "Prof. Mckayla Kuhic",
      "salary": 6155191,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "340636614246700",
      "name": "Miss Grace Rosenbaum",
      "salary": 3536866,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "340044046626459",
      "name": "Queen Candace Sanford",
      "salary": 3382353,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "370135330497530",
      "name": "Dr. Dortha Towne",
      "salary": 4704115,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343411567008656",
      "name": "Prof. Linnie Gottlieb",
      "salary": 6498986,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343698769049697",
      "name": "Queen Gloria Grady",
      "salary": 4185995,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "372308970292461",
      "name": "Lady Josiane Schinner",
      "salary": 3516574,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "379627758583783",
      "name": "Miss Yessenia Russel",
      "salary": 6718737,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "379745415086086",
      "name": "Princess Nyah Runolfsson",
      "salary": 3452637,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375214107640131",
      "name": "Lady Breanne Von",
      "salary": 5123848,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "379851671132702",
      "name": "Prof. Bettie Murray",
      "salary": 7779702,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "348137877930355",
      "name": "Princess Angelita Johnson",
      "salary": 7693248,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344630635862276",
      "name": "Princess Isabella O\"Conner",
      "salary": 3532236,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "379686065162906",
      "name": "Dr. Flo Schowalter",
      "salary": 4144573,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "378529489793705",
      "name": "Miss Maymie Dickinson",
      "salary": 6083492,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "345541847207492",
      "name": "Prof. Marisa Considine",
      "salary": 4316053,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "348853986721592",
      "name": "Miss Margarett Wiza",
      "salary": 5556449,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "342138086993928",
      "name": "Mrs. Laurence Hirthe",
      "salary": 6573446,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "340011877509857",
      "name": "Queen Aryanna Schmidt",
      "salary": 6376367,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "377742409868922",
      "name": "Ms. Pink Wiza",
      "salary": 7785716,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344613482669814",
      "name": "Queen Allison Rice",
      "salary": 3053574,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "341334251161893",
      "name": "Queen Piper Schmeler",
      "salary": 5025722,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "379433572477027",
      "name": "Dr. Bernadette Hauck",
      "salary": 5779763,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "376134801955841",
      "name": "Mrs. Rosalinda Bruen",
      "salary": 7997070,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373807710829489",
      "name": "Miss Mattie Breitenberg",
      "salary": 7585316,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "347881240530567",
      "name": "Prof. Velda O\"Keefe",
      "salary": 3133898,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343687942294558",
      "name": "Dr. Aniyah Moore",
      "salary": 6989768,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343559971057013",
      "name": "Mrs. Eloise Hirthe",
      "salary": 6120951,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373608411129108",
      "name": "Dr. Providenci Zieme",
      "salary": 7507342,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "341348510049489",
      "name": "Dr. Princess Schulist",
      "salary": 5101704,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "345893257733419",
      "name": "Mrs. Annamarie Deckow",
      "salary": 3198606,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "349992961312830",
      "name": "Ms. Jennie Towne",
      "salary": 7815097,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "370884835802397",
      "name": "Queen Joelle Considine",
      "salary": 7059488,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344532900681657",
      "name": "Ms. Sallie Schimmel",
      "salary": 5196722,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "342084905883916",
      "name": "Miss Esther Rowe",
      "salary": 3193629,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "345557021734626",
      "name": "Ms. Adela Nader",
      "salary": 7954464,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "348718955749905",
      "name": "Princess Otha Dicki",
      "salary": 3902455,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344208689092552",
      "name": "Lady Lacey Kuhic",
      "salary": 7659511,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371229062241214",
      "name": "Mrs. Jaquelin Carroll",
      "salary": 7754609,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "348876928595789",
      "name": "Prof. Treva Pfannerstill",
      "salary": 6722194,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371104994388771",
      "name": "Miss Billie Schumm",
      "salary": 5171002,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "374615859750275",
      "name": "Prof. Allison Osinski",
      "salary": 6592625,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373264560950005",
      "name": "Miss Esperanza Zulauf",
      "salary": 3707729,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371570004858764",
      "name": "Lady Caitlyn Wyman",
      "salary": 5470547,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373244653066886",
      "name": "Lady Jackie O\"Keefe",
      "salary": 5518548,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "342810175844006",
      "name": "Mrs. Marielle Schinner",
      "salary": 7008411,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "372450880802794",
      "name": "Miss Celestine Marvin",
      "salary": 6057659,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344665666360659",
      "name": "Dr. Jaqueline Osinski",
      "salary": 7505004,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "341530939916602",
      "name": "Prof. Gracie Reichert",
      "salary": 4911782,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "377209801453160",
      "name": "Mrs. Lilyan Wiza",
      "salary": 3855018,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375411087835299",
      "name": "Mrs. Gloria Brekke",
      "salary": 3618192,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373911982701903",
      "name": "Princess Aliza Schimmel",
      "salary": 3522039,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343165329297962",
      "name": "Princess Martina Mosciski",
      "salary": 3353873,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "376461154976017",
      "name": "Dr. Virginie Bechtelar",
      "salary": 4559753,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371173083176942",
      "name": "Prof. Clarissa Kunze",
      "salary": 4305481,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343691435868138",
      "name": "Ms. Sabrina Parisian",
      "salary": 6520572,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "340897581967457",
      "name": "Ms. Ayana Goodwin",
      "salary": 5070415,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "346523829767172",
      "name": "Ms. Blanca Hegmann",
      "salary": 6461688,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "346961704055344",
      "name": "Queen Muriel Dickinson",
      "salary": 6069340,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "377996815712345",
      "name": "Ms. Barbara Lesch",
      "salary": 5056187,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371765844178191",
      "name": "Miss May White",
      "salary": 3071229,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "345137203379052",
      "name": "Queen Cleora Streich",
      "salary": 5702141,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "370495688443644",
      "name": "Lady Casandra Dare",
      "salary": 3763474,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "378603695133128",
      "name": "Prof. Estelle Hegmann",
      "salary": 4257651,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "373274465418883",
      "name": "Miss Laury Kerluke",
      "salary": 7989283,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375821313941531",
      "name": "Lady Kylie Graham",
      "salary": 6518745,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "376949802508791",
      "name": "Lady Sabrina Roberts",
      "salary": 4171854,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "371257885196841",
      "name": "Lady Nicolette Sawayn",
      "salary": 6641448,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375713391943454",
      "name": "Princess Myriam Witting",
      "salary": 3014368,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "376199318198964",
      "name": "Ms. Earlene Moore",
      "salary": 7964754,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "348481991335134",
      "name": "Mrs. Maegan Shields",
      "salary": 5450532,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "346783279167482",
      "name": "Ms. Verna Morar",
      "salary": 7628443,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "374178146363472",
      "name": "Princess Vernice Lehner",
      "salary": 6156151,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "342506313886713",
      "name": "Ms. Peggie Hackett",
      "salary": 3377458,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343591743281267",
      "name": "Mrs. Lorine Haley",
      "salary": 5901153,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "376875185090789",
      "name": "Prof. Janelle Hansen",
      "salary": 4039156,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343430361333370",
      "name": "Prof. Bianka Walker",
      "salary": 4940163,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "375797635043425",
      "name": "Miss Heloise Robel",
      "salary": 4215006,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "344053371411342",
      "name": "Dr. Erica Casper",
      "salary": 4688925,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "346590617449413",
      "name": "Princess Hillary Kutch",
      "salary": 7635158,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    },
    {
      "nip": "343235748252211",
      "name": "Dr. Helen Osinski",
      "salary": 4880837,
      "attendanceDetails": null,
      "overtimeDetails": null,
      "reimbursementDetails": null
    }
  ]
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

## POST Get Payslip All Employee Excel

POST /api/payroll/admin/payslip/report-excel

> Body Parameters

```json
{
  "payrollCode": "2025_JUNE"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
"PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0018\u0000\u0000\u0000xl/worksheets/sheet2.xml|�Ks\\�����+\u0014܏\b�\u000b��$�ٷ�3�y�e�v+,�\u001d������D����J��y��[<�] ���_������\u000fO�>=~}{S_��W\u000f_?<~���\u001foo��~��~��w����ǧ~������?�|����ͯ߿��\u001f���>������׏�=|�����?>}y����ǧ�~������\u001fM_>�B)r�����7��|������\u000f|�����7�7���^�������O\u000f|;���������\u000f\u001f�?||{����\u001fn^=��o���|����ۛ�|�k����J����߽yz����ۛ�����\u001f�7�����v�����ʛ��߽�����Y�V�鬁Վ��V�����~9kl���&V�ϳ֮����\u001fW�p5\f��sw���p�ᲬQ*t�w�&P���g���r�\"W\u0010y\r��E���?g�\"�Z\u0011\u0006�;�[����-��׈�\u0006���9:��\u001a\u000f�^.?cj�R��H�u�V�e�\u0017IB�qJW�t6�^�wFt/ɟ��\u0014:7p/ɃR��:�T\tQ\u001aF���Ӌ$��8��S>�q\u000b�3�[�O|�A@��\u000b���)�N/Ulu�k���m��l��(W�r���zg��<��M�A>�$�(�Gѻ\tX�k3�e��7��\u0019��곝��82�ϣvY߀2��˖�l��K\u0015\u0006���j��F�\u000bp��i�:��7\f�DF�I�/?�\u0005��\u0016��:���K\u0015Z\u0015�1t��Nu\u0001��8\u001dW��l�'�\u0011}\u0012��k�����c�NG��R%�:����͋$��x�eqC9\u001b�)dU\u001fC?.�쓸vw�\u000f\u0015�Q-{�Z&�*\u0001\u0019��^U\u0013�g͞ �Ў�#��<�q��W\u000f�F�\u000fU�`��f/e�P$f�TS��\b�i\r/H�g\u0012\u0002�KV��TAV\u0015\"O\u000e*\u0006�9&i\u0019�Ȁhx�I�E�NkxqR=��S��j@a\u0005������QsV�9,i\u0019\u000bSI��npi^1�Z�\u000b��\u0019����U}PUe��H��p͡��Ԥe����r�p�j\u0012Z��EN�LG\u0018B˨!�\u0014_@F񿏣��Ts|�2V\u0014L^�\u001b��\u0017����¨zF%\f�e�\u0010\\J@\u0015K\u001b��8G�������w\u000f�u�Rs\t�Ϛ],Uϼ�!��\u001aB��k��H0��T́�Nf\u0002�䉧n�J5\n�����z\u0006'\u001f�wV\r������=b\u001f5\u0007�����\u0001�k��ޠ�j�;���V��O����\u001aB�\u00028�E������急�:���\u001eo(k�\"����EZpf)��eU\u001fZ���K��_֐�\u0016䤥e\u0012j5�i���j\u0012Z���[`v�|lY��\u0016\\PG*s\u001d\u001e.!�-�yK�(�!�`C[�r�Ϛ=mH�i��Qæ��HB����!g-xaKJi�\u0016�\b��ۓR�����b-8ӔU�YՇ\u0016( a\u0007d�.\f9kA�Z0a�b�ŀ\rh�F�Ϛ]�\u0005g�b\u001fXV��\u0005�F�P�{9\u0007��\u00059ei\u0019z)\u0019X��*B���(\u000b�\u001c�!��\u001a\u0002K��h���\u000b:�,�)K�إS��\u0001\u001b̚W\f����,8��ߎ��j\b��Y��\\? �,�9\u000b&g���+oPk�\"�Zˋ���S\u001cb˨!����\n7��\n9mAN[0qj`MމaC[0i˷Zˋ���S�\u0012��\u001a��\u0002:�jo\u001e�\u000f�i\u000br��2�1 {aopk^1�Zˋ��LT\u0012�˨!��t�A�圷 �-�@Ŝ�R�T��jOL\u0016oᙨ�ǗU}|��S\u0005�������0�--c�៾�q\u0003[�r�Ϛ]���\\�\u0007�U}p�\u0005sZk�\u0015�Վ9iaNZZ&��\"|���T\u0013�i\r/��3O��-����\u000b�4\u0019\b��洅9mi���l3\u001a7����Vk�t\u0006x&*�wg�p\n���\b��\u000f\u000e|�\u001c���@e�Z�3˻�@E��j-/��3W5\u001f[V����;���x`N]�S����yO��A��\u0004�g�.��3Q���;���R�*\u001di��᜷0�--�@M6\u0001p�[����\u001a^��g�j!��\u001aBk��U)�\u0001\u0002s���\f,�$�����\b�i\r/��3M�\u0010ZF\r����,��\u001f\u000f�Y\u000bs��2bi~��\u001e7�5/g���EYx�(��}g�\u0010WJYܐ�a\u000b攅9e᤬\n-ai�Pּbh��\u0017eᙣz�+����c;l��-����0�,-c\u001b�w��q�X�r�Ϙ��Wt&���ʪ>�H��Z�\u0016�,)�+����~\u0016��뿧\r_���2}���+:\u0013������*����\u0018��(�+��J˵���[�iCWs\t�Ϛ]lEf��ǔU}L�rR\u0001n������(g+-�\u0018a���6`�\u001a�>kvQ\u0015���\u0013��U}L�\u0012�39�Ǧ�r������D%9K�\rT�\u000b�Nk�4`u�&�\nwV\r#Vz^8F+a��^\u0018�za�jn�I�{b��\u001b���B�g�.��35�\u0010RF\r!u��Ɲ1�T�T�3���+g�v��*�$�Zˋ���M#D�QCT]�\u0006+b\u001f�Ɉr�����\\�0%�\t������j-/��39�u�Y5\u0004�\u0005jd���s�A9WQ�UZ�g$� I\u001b�������b+:�S���sr�-��\ne����r������B���\r]�+�Vkz�\u0015\r�*$��Ct�P\u0016��a��r�����:��ð\u0012m����Y��a%^|���\t\u0003�V���s\u0007���{\u000f�\t�s�⹃���[\u0013o\u0018k^�wZˋ��L�\u0017�]N���\n9ut��\u0007���9gi�����\u0012oHk.�wZˋ�\u0018�'�]N��ź��+\u0017\u001f�\u0007���9mi��kO�fy\u0003\\�Ih��\u0017s��X���d\u001f_|\u0001�VxP\u0018�眺8�.-\u0013\u0014H\u001e�yC]��ﴖ\u0017u�\u0019]\u000f��N����PuH�jι�s��2����0o�k^�wZˋ��̰�\tw+�\u0000��U\u0003%L�qN_�ӗ��(e$��\u001b���\b����/6C�a���!�.���\u0000N������8�/-SC\nC\u000f��/���Y����̶��w'��j��\u000b䏐\u000f�ًs��2 @O�\u000fy\u0003_�Qh��\u0017}���0\u0004��\u0010_\u0017�i}�\u001d��\u0017���eb\u001e��Mm���\u0006�봖\u0017{���0\r��\u0010_�^�I�{��9{q�^Z&���Mm�K5\t�ƴ,\u0002\u0013�Xa8��>�D7�\u001a@�O���\u0004&9�i�J�a�E6�5�`���E_b�*��;9|^GI��Z�)���%9}��0km$����׼�ﴖ\u0017}���0 �d\u001f_2O�\u0006������%9}i\u0019\u0007P�1�\r|�\u000b�Nky��\u0018�\n#�N��%�M�Jk�]Yr�����\fP$��\u001b��+���./�\u0012��A\u001f^N��%�=�\n�)�Cr����������l�k^�wZˋ���U\u0018�wr�.�(���\u0001���%9{i\u0019��?\u0005��\rx���>kwQ�\u0018�\n\u0003�N\u000eѥ�UJ���,9uIN]2��Z:�%\u001b�RMB�5��K\f\\�Ay'���-���$�M��%9{i\u0019����'ˆ��\u0015C�5��K\f\\�Qy'��R���%��H�^������Av�7𥚄Vkzї\u0018�\n��N\u000e\u0001��4jm�P򐜾$�/-�\u0000L>�&\u001b��.�u\u001a�m�W3p\u0015����\u0003���Ug.���h9{����\f0(l��\r{�F���]��\f\\�Qy'�\u0000ks���\u001a�?Z�^-g/-C��\u000e�\r{�F��Z^��\f\\��y'��M+01@��v���Z�^m�Րp\\�6��\u001a�>kwqW3`\u0015���샫)w!c\u000f�L-箖s���\u0018\u0012\u000e�ۆ��\u0012l�����\u0019�\n\u0003�N����p�y�0��r�j9si\u0019Z����\u001b蚫\b�����f��!���Ch�!!\u0016��q��SW˩K�?h1�+�p�\\�ﴖ\u0017y5�Vad��!��y\u0018\n�������r�j�аWH6�چ��\u0015C�5�ȫ\u0019�\nC�N\u000eѥ�SZ��q���W��K�Xj����m�]�%�>kw1W3P\u0015\u0006��\u001c��J;U�v}˙��̥e�$��|�0׼bh��\u0017s5\u0003Uad��!�t��������3W˙K�(\u0015���}�\u0010׼��3v���n�*��;�GW�\u0007},��?\u001d=筞�V���eH���\u001b��.�uZˋ����04�d\u001f]��IU!�����z�\\}NuA�>��7�5W\u0011Z��E]�`U\u0018�w����\u0004�\n�������s��2��d̸o�k^�wZˋ��A�08����5s�K���^�ɫ�����\u0017��<2�\r|����[���_�\u0000V\u0018�w�����\u0003�$L����z�_Z\u0006FJ��g�_s\u0015�՚^��\r`�\u0011z'�\u0010�y��\u0007��ܣ���s��25fN^�\u001b�RM|������/�\n!f�\u0010b��jo����9�������?\u0011��\u001b��W�:�Y����A�0B��\u0010_s@������9y����\f�1C�������jM/��\u0006��(��C�)\tU\n��=篞���qpnz�_���՚^��\r`�az'�\u0000Sh*�0L����z�_Z�mP���7\b6W\u0011Z���(l\u0018�\nC�N�\u00016.\u0000ԆP\u000b/��S��)L�Ԩ$\u0003PcCac\u001e{�NkyQ�0�\u0015F���\u0003l(Q\u0011\u0017�\u0007;��)l�\u0014�e\u0004\"L���Pؼbh��\u0017�\r�Ya���>ƆN}\tt�����)l�\u0014�eb*����`ؘ���՚^\u001c6\fh�Q{'�\u0018\u001b:pUY���1r\u000e\u001b9�i\u0019E\b����`ؼ�ﴖ\u0017�\r�Ya����k\u0004��\u0010[#��p���FNacbV-ٗ\u0019�\r��F�՚^\u00146\ff��{'�\u0010S�\u001a��}�c�\u00146r\n�2r\u0005�Lo0l\\O@]�5�8l\u0018�\nc�N\u000e1�ǆ\u0005G�C;F�a#�0-\u0003\u0015�\u0019�᰹\u0004�g�.\u000e\u001b�;MC�Y9\u0004�\u0012�x~�\t�8簑s��aP�,�7\u001cvZ\u0005oR{qذ�}\u001a\u0002��!���\u0010*�0�<r\u000e\u001b9�i\u0019����y�as\u0011��Z^\u00146�����{'�\u0000��Z��0 3r\n\u001b9�i\u0019ڨ%�����\b�i,ײ ���gW�KQ��\u0013�r���g��9|L���uo|։����)�֧(��y�'��KP�\b��×\r\u0016=\u0005����e!S��s$�u\u0018�ß�U̝�:l���L0�tŕ������#\u0014\u0006\u0010;W��.q{������ݿ\u0002\u0000\u0000��PK\u0007\b�\bb<\u000f\u0011\u0000\u0000�`\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0018\u0000\u0000\u0000xl/worksheets/sheet1.xml��AK�0\u0010����0w�VPdi�\bˢ7\u0011�\u001e�I\u00136ɔL���K�۲Go/y�=�L���\u0017���QTP\u0015%\b�-u.�\n�>��ϰo���҉-b\u0016S��\u0015؜����Z\f�\u000b\u001a0N�\u001bJAg.(�����[���CY>ɠ]�K�.�'��q-\u001e�=\u0007��\u0012����(�u\u0003CSw.`�닄F�K\u0005M-�Ϧ^\u001a|;\u001c�F�y�\u001f���x�\u0014�3�ٷzŎK��$:4���\u0007���z�\u0015T�\u001b�N]���z��z;e�\u0017\u0000\u0000��PK\u0007\b\u001a��\u0011�\u0000\u0000\u0000}\u0001\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u000f\u0000\u0000\u0000xl/workbook.xml��O��0\u0010���\u0014��7`�DQ�Y�j��RE�t���!��?�6\t|�\n(�D{��\fc�ߛy�ϝ���>��\u0002�*\u0005�V���E��ӯ�\u001d<\u0017����׳sW�\u0019m��:�f�$A�hdX�\u0006mgt弑1���$��(�P#F��,M���da\"��W\u0018��H�\u000f�Z�6N\u0010�ZFr6�Ԅ�f�WpF�k�<)g\u001a\u0019�L�b?B�\u0019�?\\����Q@�73���OhCʻસR��_�ӽ<M8�N.��4�N!3�4�\u0019\\40-C�YR�R�\u0016�vw|\u0018���ޒ.\u0005��:K�ȓ\u0005�x�/G�*�\u0011���M�^@�-\u0002+�����\u001a�,\u0012�������!/���\u0012��\u0007u�d�\u001b���\u0005���E\u001f�7*c=l�K7��7ҥ�\u0002v)O�I\u0015��'y\u0016��9��<Y؎�͕�1�������a�\u0001��S)�\u001fʑ7�<���\u000f�\u001a��\u001a#}�Pg\u000b�z�Nfo%�:z6��-˲)��a��/\u0000\u0000��PK\u0007\b���z�\u0001\u0000\u0000<\u0003\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0013\u0000\u0000\u0000xl/theme/theme1.xml���o�6\u0014���+\b�W�?d�A�\"��vk�\u0016Mڡ�g���P�@�I|\u001b��.\u0003\u0006t�.\u0003v�a\u0018V`\u0005V��?&@���#\u0006�HL�\u0012\u0013�.V\fM\u000e���y��\u001e)�X�~�4a��HE\u0005�q뚋\u0011��\b)�|��p��\u0006����u��1I\b:M\u0018W���X�t�qT\u0010�\u0004�5�\u0012~����\thuM��\t%�P\u001e%�i�n�I�r\\��*��Ni@F\"�%��$\f4\u0015\\�4U\u0018qH������0\u0013��Υ�2�q*\u001b\b�<\br�&�ۆG�쏚�!��\u0018��O(\u000f��!9�\u00181PzȤ���\u0007o]w΍��\u0012g����a���x(����{����E��u{�%QL�A�E�Ud������\"�-\u0004\u0001��V\u0013�v\u0006;#o\u00153��˚�����jB���\u0015t��~���\u0002�����p�\u001c+hw�z5U����&�[��\u0015��n���&4��\u0019�G+���:Ú\"]XO\u0005�UK\u000e����^%\u0017�c���\u0015�M�=��B�\u0005���\u0001M9���L! >\u001e\u0002�\u0013I�\u001e�b�Q\n\\(�c���ݎ����UY��A所ᦘ#�?xje\"P�\"�\n$M��?O�c���˗gO^�=����ӳ'���ʴ\u000b�\u0015�\u0016�������??|�����7Ͼ�s��^����?��J8]�����/������~~f��%LL��&D�;�\u0004�\u0017\tp[@2�oG\u001e�@+$�\"\u0001\u000b���\npg\u000e�f�C��()\u000fm����J.\u0007��ij\u0001n�I\u0005�\u0017��\biM�v��L{�#�(93��\u0003\u001c�4\r�6��,�IBm!�1��q�\u0001�\u0010\u0011N4���\u0011!\u0016�\u0011��u٧�\u0014JL5zD�\u000ePk\t\u000f�D�÷h\u0002\f�6��1Tj��\u0010�\bf\u000b7\"�U\u0002x\u0004�\u0016��J�o�LCb�\b\u0012f\u0012{�c[\u0012\u0007s\u0019T\u0016Li\t<\"L�ݐ(ec��y%����};��yR%��G6b\u000f�0��8\u001aƐ�֜(�M�3u$\u0004\u0003tOh�8Q}b�{�(�K��CJ��\u001dC\u000fh\u0014�o�lf&m�&\u0011��bΦ@*�����P~i�[jo�����f���\u000e���mKj}З\u001b�e����5�\u0019�GxlC>v����c��\u0017���u����{֢M���g���#ܔ2v�����<�\u0012��c�X~�C\u0017\u001f\u001e�x�$v�\u0000\u0015�HB~���_P\u001d\u001fĐ\u0012\u001f��\b�*]G\n�B��ō��\t6K�EX��Z�k\u0012'\u0007@/�]�b\\S���^�\u001ct\f��]�L\u0001^����\"�`U\u0011�\u001a\u0011���D��u�\u0018Ԩ�h�T8ƪ0�\u0011���^�P�T\u0000���:\u0015����}���YM�]�ޠ{�\"_a�+\"��V\u0015al�\u0018B�<���\u001e\fꗺ]+���>��Y=\u001b\u0018�ޡ\u0013\u001f�:��Q\u0000���\f4FA��>Vٹ\n,�>\u000etY��9YR��\bT\\��SE�\t�D\"F\u0013\u001fo����B[��w?\\q\u0003�ë����d:%�n\u0018Y��)]8��}G��F�4�\u0007qx�&l&�C�c���\n\u0018R�/�\u0019Ril�E\u0015����Q��R]<���\u0018ʎb\u001e����\u000b9F\u001e��嬜�\u0012N��:����V��lh ��S��5yCU�^�W{�\r6.�\u0012��\u0010\fi\u001b��:�Қz�\u001a�!0��\u001a�ֶ��w�\u0006˻�1����V�'\u0013��$�#2�\u0019�E��!�$�Z���\u001b����v4���o\u0000\u0000\u0000��PK\u0007\b���{}\u0005\u0000\u0000e\u001c\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\r\u0000\u0000\u0000xl/styles.xml��]��(\u0017���_�|?c�&�Ld�zժR��j�v��\u0012�\u00134|X������+p��qf�8�b�c�\u0003\u001c·I��r\u0006�Di*E���,\u0001D`YQ�/��~|��$\u001f���ڜ\u0018�~ �\u0000`9\u0013�H\u000e�4�4��@8���!�rVKő��R�S�(�*��8K\u0017Y�N9�\"�\b[��@d]SL>K�r\"LG!�\u0010Q���Q�!�P�{��\rŭ6�_A�\u0018b���T쯭\u0013���\u0007�L\u0000Tj.�s'\u000b$ɣ��\u0012%y�\f�QFͩgU{>\u0003UQ�W��\u00106�8L�gR}B���&5t\u000e��ش��\u0010;��#7a�(�5_T�!C��\u0007ڄ%VQA��k���GE��5\u0012G�\u0010�\u0015\u0003�M\u000e�,\n�\u001f}�;��i\n�Q[ϑzn��\u000b_���\u001eA&\u001cN��Z��\u001eK~�HJ,&W�µ����\u0019<oKKC��Z�u}�\r=��!��f�m�kcy$�\t�ɓ�����N��\u001b\u0010?��A����T�'$H�5lށ�Ġ\n\u0019�b)\f\u0011�ǩ\t�)̕�\u0013���_��b�Yd�0)�O�����!�5\u0014\u0019/��1��Z3�j*�C��9]d0�,6p\u00197�s\u0012s\u0013{H7�bRr�\u000f��բ�f�Ƌj��e�l��c!����>^L�ƛ�O\u00160{�PX�D�$\u0016?7��\"\u0011�ܣ�߼��\\����\td\u000f�\\V�y,�b�^���U\u001c��Y`�Bx�gp1��\u0010�pdc\u000bW�I~Z#TdҞp��\"G���\u0007�b&k\u0015X�\u0001�a&l\u001d`\u001f\u0006X�'�\u0001[\u000e��\u001e1�V\u0003l}3lxP\u000f7�\u001e\u0006��f�f�=�\f{\u001c`0��\u0006�\u0011\u000eގ�#��0\u0018�Fq\u0000o\u000f\u00048��ؤ�\u0016n\u0014\u000b���\b7D\u0003�\u001b\r�i�]ߐ���\u0019����^H�v�\u0014\t�K�_L\u0000�K�_\n�u_/�5\u0014T�\u0001�|\u0001_���+�3?�Y\u001bX\u0005�Z\u0000�>\u0000����\u0015�j\r�z\u0000Vm�U��9��@wq]���X���w\r�\n\u0001��[��@�c����\u0000����n.�\u0012�G\u0001�\bǑ\u0017�c��Jl��wA�\u001d��\u001c��q8\u0010ʷ�vc�o�ƛ􎬶-�������˲\f������Mʼ��h�e+L������/pD�H�\u0013��W�cɤ\u0002�@8�:�^R�\u0002q�u��\u0018�)��9a�׈Sv�\u001e.��\u0013��\u001f��t�ה�0\t��\tʼA�\u0010%�P�����&E\"� \u000e7��`��\u001bսB'�X���o��wRUD]\u0018�\u0013�9#��zwStp��w#�2O�u'����Ӿ�޲�@n�Q���\u001b��1a��91���\u0018��@��\u000b7_�\"�\u0012�l�7)c�f��~�yjk�Q#bǿ\u0019\rl=\u001d#��`\u0017#\u0004)p�Q$�c+\u001b0`�Rf�\bР0n�2��0u��\t�ܸlu9l���Ԩe��\u001e��\u000f�dhs[\u0006ס׏�(���\u0007�h˽���p1\u0011��-�\u000b\u0000\u0000��PK\u0007\b��\t;�\u0004\u0000\u0000�\u0015\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0014\u0000\u0000\u0000xl/sharedStrings.xml|Y�O\u001c9\u0012~�����{��r�.�.�\n\b\u001b6a\u0002\u0007\\����1�5\u001e7��!���'���)v#\u001e\u0010�r����}U����&�'��Ч7\u0007��\u001f0��~\u0019�Û�����Os��ۿ�\u001e��}��4�9X����\u000e\u000f�n�7nx�?��u\u0013���q����\u000f��c�n9��\u001f7�Pp�\u000f7.�\u0003���4�9\u0010�\u001e�m\nn��_O޾\u001e�����O�]�>\u001c߾>,>?r\u001b���c7xv��˻�_]�>��/g^_<�<��gg�6\u000f�o/]u�ʇ��6\u000f~�����\u001b�.&�'��\u001f���E\u0018\u0006v��2;v=f\u001f�U������[���ʭCv��uk?V7�\u001f���yx�N�O�\u001d{�ۧd��'�\u000b18v��Mh\u001c��\beQ*#\u0005Vw�>�p�\u000b�}�9n�UV�\u0004(c��\u001c\u0010�ϯ���>��a\u0015��]���?�ڄ��4IkP\"4/q�\u0014\u0006��r�_�Np$��Z\t\u001clۉ\u0010c�ӯF\u0014H\u0019�\u0002\u0011\u0001e#����Ʊ�n��a��&�\u0002��UZ����.�bG�-U\u0013R��\u001c��Rm����Z(aU��~��UX��*\u001b�t\u0005Ѐ\u0010R\t+u\u001d�R��!�|ܔ��{�R߸��\u001a\f��\u0012T�5Y�꿺�<��W\rW\b���[�\tfR2\u0015�gg.�P55J�HXKd-J�\fL)�V>z�k\u000e~��+(�ڐ�\u001a�mw�y�}bף�l��d,(���\u0016�������\u001fW.�ح�C\r\u0004�Zqe\f)�e\u001d�\u0012��U\u000e�\u0018\\r��1�u�\u001cR\u0019m@��VQL�X��2\f}*U\u0011RjU���ɖ6�ڙ��߿b�n�vѱ��U�\u001a=�Qk�Bj�����u�\u001e\u0006��ܶ��,Q�Rk���jb\u001f�\u0013���ʵK�}^֎p@�ȥ%U;R:�]�Ǖc7��T��D\t\u0005j87Z�v\u001c�CJ����8���\u001a6�5�-�V۪2�n��}\u000e�DdY3\u000b\t��\u0012\u0017��+��\n�C_\u0000¿�P�\u0005�2� ��0J:~���Sp�j;\f>6L�T\u0012\u00147��F,B��0�O;�bW����ah@\u0005)\u0001\u00128iɡ\u0006�ɛ��'���:n�\u0002M\u0000(�&�}F��8\u0016���� ji\u0000�\u0010Y��*\u0002���Qz�1��}�W�刔\u001a�Fe�\u00104\u001f��\u0006w�ct��\u001fG���~.?Fs�@\u000b�+c�N�\u0013��_\\\u001c[\u0006LApcK��ʫ���m�g�B�\u000em���`$\tN���vQȨ��4�eh��1\n�х�k#�\u0017�\u000f.�qd��[\u0003�\u0004��F[�V��B��s��>u���<�\u001a��\u001c�\u0010)n��鹣�s)M,�\t˚��H\nɭ�Ɗڑ�\u0015�\fi=��Ԁ�\b�������1N |\u0015��\u0003�(�\u0002�`lկ{\u0003����\u0017��JDEB\u0012qQŠ�Ա��-�8\u0016V�v�P��\u000b\u000eV)#�B��\u001f\\\fi��q޶X\u0015\r'\u0002n���)��]��\u000f�Ow>?�� c@H��+]�����>.�w�G��[h�\rY)��JU�5ɮ\u0014\n|-�>��+e-\u0001W�U��8�}\u0018f��Ps#\u0001@X���_��),}�\u0002�#�z&*u!�\u0002�e+�{\u0013��3�>��(\u0014\u0011���Z)���\u0005-|����6���j@\u0010����\u001c\u001f�D�m^%n�4�\f\u00178ǉ\u001fz\u001f�\ti�Ba9/r�n�r�k7���M�Ҥ�F�\"��B��Sq�\u000e��gv�i\\@)E\\\u0000�Ԣ>^���ѱOn�\u001a�\f��J���F�\u001fi�(�d��F\u0000\u0004/�ĭP3���u~7#�\bD!\u001b!$��h/u�����1\u0014��s\u001f\u001b\u00114E�\b����\u001e��y���c�����0��\u0019\u0002��Z��\u0010�@S\u0012q\u001c��r�R�R�2ʒ��\u001a�,��\u000e�\u0017CHC#��Eh�\u0019��\u0019\n=\u001d\u001e}v��cl���7\u001cQ�9�F\u0019�3#�\u000bc�%v��4\u0007!!�Vȵ6��ڋ=׭�\u0019��\u0001\u000e�����\u0013���05ؼf\u0014Rqc�����h�ǉ�~\u0018�d�p�)4E�VZk,⦉W�K���E�Bn�Z�zN��Y\"xv�ː�[d.�5\u001c�B����<ĝKm>'%\u0001�!�J�6d>K���׭�\u0007Z\u0000k\u0004q����/�\u0019�7�\u0002^!��y�\u0016=?ke��\u0018�c�~�B�ҵ�\u0000JZ�\u001c���s�\u000f�����V����9\u0000Bn\u0010H[9���X��P����I�\u0016$N\u0013+�JoB��\\\\�߁��H\u0019��d�\u0004�veZ~���/�\u0002�J�\u0011�4\u00015U�qt�s��?l\\j\u0019�\u001a�K�\u0014�C���\u0016��k/(o\"k�\u0001E P��Sn���ˎ����B$ ]�\u001dȀm������*�-\u001a\u0003$�\u0011���F�����\u0017V�ĥU�\u0018)Q�Ax��a�\u0006}�\u001aڊ�樭\u0002,��]L��8\u0001�\\&\b\u0005\u0015��`�̌[����jR\u0019\u0001\bh\u000b��Gӏ�BB��[�\u0006\ti+��BqCu\u0016�\u000b��z���|\u001e�5\u0016�Pd�\u0002�\u001b�{2�)t}�D����v�\t�\u0000т�(U���Pb��۰�0�!Ւ�4X�`�\u001a[3Y)�S��/��V�F\u001a�\u0016\u0010\u0015`[],�p�]�����JO�AA\u00164I�l��e�a�>��ij\u0006�Q��\u001b�G\u0014���yv�W-\u000e�Bq���hjh�2\u0004����g7�\u0012�\u0005�(\f�zr�\u0013O?�y�\\����6��(n�����Ҿ1\\j-6%J��\u0001\u0011�f�y����c�.���n��QqY\u0002���3�\u001f}Ja7���nH \u0001���B8�9t���c+\u000bZY������\\��\u0018]�����(�\u0002U�$%\u0004T�U�p棟��7\u00177G������wu8\f���\u0005\u0000\u0000��PK\u0007\b��\u0002�\u0015\t\u0000\u0000a\u001b\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u001a\u0000\u0000\u0000xl/_rels/workbook.xml.rels���N�0\u0010���O��]\u0006�#�\u0000{1&{U|�\u0006\u0006�l?H�*��&�\b�\u001e0�Ҧ����)\u000e�5��\u0002k�JȒ\u0014\u0004�Ʒ��%����\u000fp���\u001722j�X���h��\u0012T��#\"7����\u000f�Fk:\u001f������ ���\t�4�ǰl@�j�c[B8�\u0019�Z��b\t�>�X\u0011E��ʒ�\u001a\u0010�y�-�}�醞|�n��\u000b\u0002�\u0019\u0000U�K�eZ>�8�\r�ޞ��6�͌��,�t��ESu��v\u0016�h�����i���ݚ�J\u0006j_cЮ����_����W\u0000\u0000\u0000��PK\u0007\b��^��\u0000\u0000\u0000�\u0003\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0011\u0000\u0000\u0000docProps/core.xml��MK\u00031\u0010����%��ɶP4��A�IA��x\u000bɴ\rn>HR7�{���V�M�i�g\u001e�$|QlW}aLƻ9i\u001bJ*t�k�6s��Zַd!n�\nL����\u0007��`���\\b*��6��\u0000�ڢ���\u0001]���G+sj|�@��Sn\u0010&���b�Zf\t{a\u001dF#9*�\u001a�a\u0017�A�\u0015`�\u0016]N�6-�ٌѦ�\u000bCrAZ��\u0003^EO�H�dF�����\u000e���\u0016ޟ\u001e_���ƥ,�B\"�VLE��GQv�p�\u0018�c��\u0000uU�a�.��mz��Z\u0012������lE)\u001b����k�,�^�����$\u0010\u001c����\t\u0000\u0000��PK\u0007\bE�@�\u0010\u0001\u0000\u0000\u001c\u0002\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0010\u0000\u0000\u0000docProps/app.xml��?K\u00041\u0010\u0005��O�����B���\u0010��Z��K2{7�̄�xD?���Y[>\u001e�x�\u001fz��\u0019�������\r�Q\u0012�qv�����\u001diR�\u0019�\u000e�d�ٝ��\u001e@�\t˪�T�^�&�����#ȶQ�\u0007��\u0005��f�n\u0001�!'L��\u000b�\u001fq���I��>}[>*�\u000b~\u0011[�B\u0005���\u0012�}���j$\u001c�ex�\u00113}��������\u0015\u0000\u0000��PK\u0007\b�8 \t�\u0000\u0000\u0000 \u0001\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u000b\u0000\u0000\u0000_rels/.rels���J�0\u0010\u0006���)B��t�\u000f\"��^D؛H}��L��$\u0013&�Ʒ\u0017��Ŋ�80|ߏ�x�1�\u0017��1i�oZ) Yt>�Z>�w�ky��\u001d\u001f \u0018����s\u00115�T�����R�N\u0010Mi0C�1\fH�pi�F����\b�жW�.3d��\u0014g�%��)zC#��\u000e�=a.����\u0018���2��\u0015��[�E�\u001c!�J��ʐ\u001c�]&�@��\u0003�.E��Ê�\"�6�ϳ�\bl�a󙺙���ՠ^��'�y\u001b����\u001f�e��t�\u0001\u0000\u0000��PK\u0007\b��{��\u0000\u0000\u0000d\u0002\u0000\u0000PK\u0003\u0004\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0013\u0000\u0000\u0000[Content_Types].xml��ͮ�0\u0010��<E�-��{\u0017\b�$w��\u0012�Dy\u0000cO\u0012���\u001aOK��(I[AiQ�f\u0013o&��t�q�2\u0004_쀲�X�'�\u0012\u0005D��Ů\u0016?�_����yS��\tr1\u0004\u001fs-z��A�lz\b:KL\u0010��[��9K�N%m6�\u0003��Z�S\u0006#C���\f�T���[����!�\\\u0002�E�q\u001e\u001cY��)yg4;�j\u0017�\u0019�<\u0010$��fr�R~;\u0004/�J\u001d\b\u0017Q��u�y��\u001d\u00109\u000bū&��\u0003�B\r^q\u000f\u0001�����\u0005wl[g���\u0006�,����\u0011x\u0015�y�!?\f͉@��\u0003p�r\u000e���\u0017��'�fi���A�x��E�J���)=�\u0002�\u0015�`�D������$��t<~'�.���ѩ\u0017�\u0004��\u001cWk�{�6��WX��qgzM`�3��-�:f�k�����\u001d\u0000\u0000��PK\u0007\b��~�[\u0001\u0000\u0000h\u0005\u0000\u0000PK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000�\bb<\u000f\u0011\u0000\u0000�`\u0000\u0000\u0018\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000xl/worksheets/sheet2.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000\u001a��\u0011�\u0000\u0000\u0000}\u0001\u0000\u0000\u0018\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000U\u0011\u0000\u0000xl/worksheets/sheet1.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000���z�\u0001\u0000\u0000<\u0003\u0000\u0000\u000f\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000}\u0012\u0000\u0000xl/workbook.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000���{}\u0005\u0000\u0000e\u001c\u0000\u0000\u0013\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000a\u0014\u0000\u0000xl/theme/theme1.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000��\t;�\u0004\u0000\u0000�\u0015\u0000\u0000\r\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u001f\u001a\u0000\u0000xl/styles.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000��\u0002�\u0015\t\u0000\u0000a\u001b\u0000\u0000\u0014\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u000e\u001f\u0000\u0000xl/sharedStrings.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000��^��\u0000\u0000\u0000�\u0003\u0000\u0000\u001a\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000e(\u0000\u0000xl/_rels/workbook.xml.relsPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000E�@�\u0010\u0001\u0000\u0000\u001c\u0002\u0000\u0000\u0011\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000�)\u0000\u0000docProps/core.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000�8 \t�\u0000\u0000\u0000 \u0001\u0000\u0000\u0010\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000�*\u0000\u0000docProps/app.xmlPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000��{��\u0000\u0000\u0000d\u0002\u0000\u0000\u000b\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000�+\u0000\u0000_rels/.relsPK\u0001\u0002\u0014\u0000\u0014\u0000\b\u0000\b\u0000\u0000\u0000\u0000\u0000��~�[\u0001\u0000\u0000h\u0005\u0000\u0000\u0013\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0000\u0004-\u0000\u0000[Content_Types].xmlPK\u0005\u0006\u0000\u0000\u0000\u0000\u000b\u0000\u000b\u0000�\u0002\u0000\u0000�.\u0000\u0000\u0000\u0000"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# payroll-service/api

## POST Get Payslip Employee

POST /api/payroll/api/payslip

> Body Parameters

```json
{
  "payrollCode": "2025_JUNE"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
"%PDF-1.3\n3 0 obj\n<</Type /Page\n/Parent 1 0 R\n/Resources 2 0 R\n/Contents 4 0 R>>\nendobj\n4 0 obj\n<</Filter /FlateDecode /Length 492>>\nstream\nx\u0001��Os�L\f��|\n\u001d�̋\"�_/�7��ig�2��\\���&`2�i'߾�\u001aB\u001384m\n\u0007_$=�����Ӏ`1 �\u001e~\f\b�\b>��_\u0007\u00179�_�m��h�Z<�\u0010$��*3�>D�J\n��\u000b��\u0001\u001b$�|\u000eW��S3r\u0000\u001f,������<���q�/��8�&�Γ-9�d�)Ke���:Z-B�\u0015�)\u0016!\u0000+$�|����v�ަٷ\u001f'cд��A`W�\u0003Ro�VVq\f������e#0��X\u0015ߎzء�4�B�\bSYJ�4��G��Xe\u0019�*�\u0016J��K\\�PgI����u)u\u0011�2vR-ۃ��\u0000�S2���1�K��\u0018\u0014);\"7�p���9C�.yz]�G�m��L*�4�F�X�`�\u0014)��ȬӨuR��=6]��\u0007��ΐ9�\u001c3�\u000fn֛�\u001d�B{\u0004�\u0004�6��g{\u0013�\u0013�-+�P�\u0003��\r�q��\u000f�]�V�M��U���ϐ�nM#�H�P50�,��:�\u0007��\u0007�g��u?<<�O�ဴO���b:\u001e��C��/_w��\u0017\u0014�1��E�A}➴a\f*-�W��!�n֫��kf�*(Tv�/�*\u001f�\f\u0000\u0000����`�\nendstream\nendobj\n1 0 obj\n<</Type /Pages\n/Kids [3 0 R ]\n/Count 1\n/MediaBox [0 0 595.28 841.89]\n>>\nendobj\n5 0 obj\n<</Type /Font\n/BaseFont /Helvetica-Bold\n/Subtype /Type1\n/Encoding /WinAnsiEncoding\n>>\nendobj\n6 0 obj\n<</Type /Font\n/BaseFont /Helvetica\n/Subtype /Type1\n/Encoding /WinAnsiEncoding\n>>\nendobj\n2 0 obj\n<<\n/ProcSet [/PDF /Text /ImageB /ImageC /ImageI]\n/Font <<\n/Ff5d2de5f3a71699ae4b2d83179e62d09e6fc4126 5 0 R\n/F0a76705d18e0494dd24cb573e53aa0a8c710ec99 6 0 R\n>>\n/XObject <<\n>>\n/ColorSpace <<\n>>\n>>\nendobj\n7 0 obj\n<<\n/Producer (��\u0000F\u0000P\u0000D\u0000F\u0000 \u00001\u0000.\u00007)\n/CreationDate (D:20250610120744)\n/ModDate (D:20250610120744)\n>>\nendobj\n8 0 obj\n<<\n/Type /Catalog\n/Pages 1 0 R\n/Names <<\n/EmbeddedFiles << /Names [\n  \n] >>\n>>\n>>\nendobj\nxref\n0 9\n0000000000 65535 f \n0000000649 00000 n \n0000000933 00000 n \n0000000009 00000 n \n0000000087 00000 n \n0000000736 00000 n \n0000000837 00000 n \n0000001143 00000 n \n0000001256 00000 n \ntrailer\n<<\n/Size 9\n/Root 8 0 R\n/Info 7 0 R\n>>\nstartxref\n1353\n%%EOF\n"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### Responses Data Schema

# Data Schema

