# araQR

## Getting Started

Build the application binary by running the following command:

```bash
make bin
```

Run the dependencies by running the following command:
```bash
docker-compose up -d
```

Run the following command to run the application in local:

```bash
go run . serve \
    --log.level debug \
    --db.sslmode disable \
    --auth.secret-key secret
```


**araQR** is a lightweight, privacy-aware platform that enables car owners to receive predefined messages from third
parties through a QR code sticker placed inside their vehicle. It helps people notify drivers (e.g. in case of improper
parking) without revealing personal contact information.

---

## 🚘 Why araQR?

In urban areas, especially densely populated cities, it's often difficult to reach vehicle owners when needed. Common
issues include:

- Improper parking
- Blocking exits or driveways
- Leaving lights on
- Forgotten items

araQR solves this by offering a one-way, privacy-preserving communication channel. Car owners control what information
is visible and are notified via mobile app when someone scans their QR sticker and selects a predefined message.

---

## 📱 How It Works

### For Car Owners (QR Code Holders):

1. Purchase a unique QR code sticker.
2. Download the mobile app (React Native).
3. Scan and link the QR code to their account.
4. Customize profile visibility (e.g., license plate, car color).
5. Get real-time push notifications when someone scans the sticker.

### For 3rd Parties (Message Senders):

1. Scan the QR code using their phone camera (no app required).
2. They are redirected to a browser page.
3. Choose from predefined messages (e.g., “You're blocking the garage.”).
4. Message is instantly delivered as a push notification to the car owner's app.

---

## 📦 Technical Stack

| Layer            | Technology                |
|------------------|---------------------------|
| Mobile App       | React Native (Expo)       |
| Backend          | Golang (RESTful APIs)     |
| Hosting          | AWS App Runner            |
| QR Code Storage  | AWS S3                    |
| Database         | AWS DynamoDB (NoSQL)      |
| Notifications    | Expo Push Notifications   |
| Domain & Routing | AWS Route 53 + CloudFront |
| CI/CD            | GitHub Actions            |

---

## 🧾 Technical Requirements

- Mobile app (React Native) with camera permission for QR scanning
- Secure user authentication and authorization
- QR code linking flow (one QR code per car, optional multiple cars per user)
- RESTful API endpoints for registration, message delivery, and QR linking
- DynamoDB tables for `Users`, `QrCodes`, `Messages`
- SVG QR code generation in backend using Go library (`github.com/skip2/go-qrcode`)
- S3 integration for storing QR codes
- HTTPS redirect logic for browser-side message sending
- Admin dashboard or CLI to track QR code generation and linking (optional)

---

## 📋 Non-Technical Requirements

- Physically printed QR stickers (unique UUID per sticker)
- Waterproof, windshield-compatible material
- Sticker generation system (batch printing in Izmir)
- Unique URL structure: `https://araqr.com/{uuid}`
- GDPR-compliant data handling (no PII exposed to third parties)
- Simple, intuitive UI/UX with localized content (TR / EN)
- Payment & order handling for QR sticker purchases (optional phase 2)

---

## 🧪 Future Ideas

- Anonymous reply feature (time-limited)
- Public "report abuse" form for misuse
- Statistics dashboard for users (how many times scanned)
- Admin moderation of predefined message list
- Integration with SMS for fallback notification

---

## ✅ MVP Scope

- Register/Login via mobile app
- Scan and claim QR code
- Predefined message delivery via browser
- Push notification via Expo
- Admin CLI for sticker/UUID generation
- Hosted backend on AWS App Runner
- SVG QR code generation and upload to S3

---

## 🚧 In Progress

- 🔧 QR Sticker purchase & printing workflow
- 🎨 Mobile UI polish
- 🔐 Security and abuse protection
- 🌍 Localization

---

## 🤝 Contributions

Not open-source (yet). If you're interested in helping, feel free to reach out.

---

## 🧠 License

Proprietary. All rights reserved by Ataberk.

