# Golang RESTAPI
## Flow
1. Input : Menerima input dari user
2. Handler : Mapping input dari user ke *struct input*
3. Service : Mapping dari *struct input* ke *struct entity*
4. Repository : Handler query from *struct entity* to *database*
## Tech Stack
- Golang Gin
- Golang Gorm
- Gorm Driver Mysql
## Database Structure
**Table users**
Field|Type|Length
:-----|:----:|:--:
ID|int|11
Name|varchar|100
Occupation|varchar|50
Email|varchar|100
PasswordHash|varchar|255
Avatar|varchar|100
Role|varchar|50
Token|varchar|255
CreatedAt|datetime|-
UpdatedAt|datetime|-

**Table campaigns**
Field|Type|Length
:-----|:----:|:--:
ID|int|11
UserID|int|11
Name|varchar|100
ShortDescription|varchar|100
Description|varchar|255
GoalAmount|int|15
CurrentAmount|int|15
Perks|text|-
BackerCount|int|20
Slug|varchar|100
CreatedAt|datetime|-
UpdatedAt|datetime|-

**Table campaigns_images**
Field|Type|Length
:-----|:----:|:--:
ID|int|11
CampaignID|int|11
FileName|varchar|100
IsPrimary|tinyint|2
CreatedAt|datetime|-
UpdatedAt|datetime|-

**Table transactions**
Field|Type|Length
:-----|:----:|:--:
ID|int|11
CampaignID|int|11
UserID|int|11
Amount|int|100
Status|int|50
Code|varchar|50
CreatedAt|datetime|-
UpdatedAt|datetime|-